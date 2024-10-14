package assist

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"xddq/pkg/game/client"
	"xddq/pkg/game/protocol/pb"
	"xddq/pkg/logutil"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

var (
	ErrRestart = errors.New("restart")
	ErrStop    = errors.New("stop")
)

type Assistant struct {
	c *client.Client

	conf *AssistantConfig

	ctx    context.Context
	cancel context.CancelCauseFunc

	logger atomic.Value

	cron    *cron.Cron
	runners []Runner

	loginedHandler func(ctx client.Context) error

	stopMu  sync.Mutex
	stopped atomic.Bool
	stopC   chan struct{}

	role         *RoleManager
	bag          *BagManager
	mail         *MailManager
	playerInfo   *PlayerInfoManager
	systemUnlock *SystemUnlockManager
	activity     *ActivityManager
	storage      *PlayerDataStorageManager
	reward       *RewardManager
	homeland     *HomelandManager
	union        *UnionManager

	yard      *YardManager
	palace    *PalaceManager
	tower     *TowerManager
	destiny   *DestinyManager
	universe  *UniverseManager
	yuebao    *YuebaoManager
	gather    *GatherManager
	askWay    *AskWayManager
	dailyTask *DailyTaskManager
	challenge *ChallengeManager
	pupil     *PupilManager
	helper    *HelperManager
}

func NewAssistant(conf *AssistantConfig) *Assistant {
	if conf.Username == "" || conf.Password == "" {
		panic("username and password is required")
	}

	a := &Assistant{
		conf:  conf,
		ctx:   context.Background(),
		stopC: make(chan struct{}),
	}

	a.logger.Store(logutil.DefaultLogger)

	return a
}

func (a *Assistant) init() {
	a.cron = cron.New(
		cron.WithParser(cron.NewParser(
			cron.SecondOptional|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.Dow|cron.Descriptor,
		)),
		cron.WithLogger(cron.DefaultLogger),
		cron.WithChain(
			cron.Recover(cron.DefaultLogger),
		),
	)

	a.ctx, a.cancel = context.WithCancelCause(context.Background())

	a.c = client.NewClient(a.conf.Username, a.conf.Password, client.WithContext(a.ctx))

	a.c.OnLoginOverMsg(a.onLoginOverMsg)
	a.c.OnOtherLoginMsg(a.onOtherLoginMsg)

	a.initManager()
}

func (a *Assistant) initManager() {
	a.bag = NewBagManager(a)
	a.role = NewRoleManager(a)
	a.mail = NewMailManager(a)
	a.playerInfo = NewPlayerInfoManager(a)
	a.systemUnlock = NewSystemUnlockManager(a)
	a.activity = NewActivityManager(a)
	a.reward = NewRewardManager(a)
	a.storage = NewPlayerDataStorageManager(a)

	a.homeland = NewHomelandManager(a)
	a.yard = NewYardManager(a)
	a.palace = NewPalaceManager(a)
	a.tower = NewTowerManager(a)
	a.destiny = NewDestinyManager(a)
	a.universe = NewUniverseManager(a)
	a.yuebao = NewYuebaoManager(a)
	a.gather = NewGatherManager(a)
	a.union = NewUnionManager(a)
	a.askWay = NewAskWayManager(a)
	a.dailyTask = NewDailyTaskManager(a)
	a.challenge = NewChallengeManager(a)
	a.pupil = NewPupilManager(a)
	a.helper = NewHelperManager(a)
}

func (a *Assistant) sendWithMsgId(ctx context.Context, msgId int32, data []byte) (err error) {
	return a.c.Send(ctx, msgId, data)
}

func (a *Assistant) Context() context.Context {
	return a.ctx
}

func (a *Assistant) Client() *client.Client {
	return a.c
}

func (a *Assistant) Config() *AssistantConfig {
	return a.conf
}

func (a *Assistant) reset() error {
	for _, runner := range a.runners {
		if err := runner.Stop(); err != nil {
			return err
		}
	}

	return nil
}

func (a *Assistant) addRunner(runner ...Runner) {
	for _, r := range runner {
		if r == nil {
			continue
		}

		a.runners = append(a.runners, r)
	}
}

func (a *Assistant) onLogined(ctx client.Context) error {
	a.L().Info("🎉 Login success")

	for _, key := range a.conf.ExchangeKey {
		ctx.Send(ctx, &pb.UseExchangeKeyReq{
			Key: &key,
		})
	}

	if a.loginedHandler != nil {
		return a.loginedHandler(ctx)
	}

	return nil
}

func (a *Assistant) onLoginOverMsg(_ client.Context, _ *pb.LoginOverMsg) error {
	a.L().Info("Synchronize data successfully")

	a.cron.Start()

	for _, runner := range a.runners {
		go runner.Run(a.ctx)
	}

	return nil
}

func (a *Assistant) onOtherLoginMsg(_ client.Context, _ *pb.OtherLoginMsg) error {
	a.L().Warn("其他玩家登录")
	return nil
}

func (a *Assistant) WithLoginedHandler(fn func(ctx client.Context) error) {
	a.loginedHandler = fn
}

func (a *Assistant) stop() error {
	a.cancel(ErrStop)

	for _, runner := range a.runners {
		runner.Stop()
	}
	a.runners = nil

	a.L().Info("runners stopped")

	<-a.cron.Stop().Done()

	a.L().Info("cron stopped")

	return a.c.Close()
}

func (a *Assistant) Stop() error {
	a.stopMu.Lock()
	defer a.stopMu.Unlock()

	if a.stopped.Load() {
		return nil
	}

	close(a.stopC)

	a.stopped.Store(true)

	a.L().Info("stopped", zap.String("username", a.conf.Username), zap.String("nickname", a.playerInfo.Nickname()))

	return a.stop()
}

func (a *Assistant) removeJob(id int) {
	if a.cron == nil {
		return
	}

	a.cron.Remove(cron.EntryID(id))
}

func (a *Assistant) addJob(name, spec string, job func(ctx context.Context) error) int {
	a.L().Debug("add job", zap.String("name", name), zap.String("spec", spec))

	id, err := a.cron.AddFunc(spec, func() {
		if err := job(a.ctx); err != nil {
			a.L().Error("execute job error", zap.String("name", name), zap.Error(err))
		}
	})
	if err != nil {
		panic(err)
	}
	return int(id)
}

func (a *Assistant) run() error {
	defer a.stop()

	a.init()

	a.c.WithOnLogined(a.onLogined)

	// 每天0点重启
	a.cron.AddFunc("0 0 0 * * *", func() {
		a.cancel(ErrRestart)
	})

	return a.c.Run()
}

func (a *Assistant) Run() error {
	for {
		select {
		case <-a.stopC:
			a.L().Info(
				"assistant stopped",
				zap.String("username", a.conf.Username),
				zap.String("nickname", a.playerInfo.Nickname()),
			)
			return nil

		default:
			err := a.run()
			if errors.Is(err, context.Canceled) && context.Cause(a.ctx) == ErrRestart {
				continue
			}

			if err != nil {
				a.L().Info(
					"🚨 Exited",
					zap.String("username", a.conf.Username),
					zap.String("nickname", a.playerInfo.Nickname()),
					zap.Error(err),
				)
			}

			a.L().Info(
				"🔁 Reconnecting",
				zap.String("username", a.conf.Username),
				zap.String("nickname", a.playerInfo.Nickname()),
				zap.Duration("duration", a.conf.ReconnectDuration),
				zap.String("restartAt", time.Now().Add(a.conf.ReconnectDuration).Format("2006-01-02 15:04:05")),
			)

			select {
			case <-a.stopC:
				return nil
			case <-time.After(a.conf.ReconnectDuration):
			}
		}
	}
}

func (a *Assistant) Activity() *ActivityManager {
	return a.activity
}

func (a *Assistant) Storage() *PlayerDataStorageManager {
	return a.storage
}

func (a *Assistant) Bag() *BagManager {
	return a.bag
}

func (a *Assistant) Role() *RoleManager {
	return a.role
}

func (a *Assistant) Yard() *YardManager {
	return a.yard
}

func (a *Assistant) Palace() *PalaceManager {
	return a.palace
}

func (a *Assistant) Tower() *TowerManager {
	return a.tower
}

func (a *Assistant) Destiny() *DestinyManager {
	return a.destiny
}

func (a *Assistant) Union() *UnionManager {
	return a.union
}

func (a *Assistant) PlayerInfo() *PlayerInfoManager {
	return a.playerInfo
}
