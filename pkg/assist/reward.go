package assist

import (
	"context"
	"time"

	"xddq/pkg/game/client"
	"xddq/pkg/game/db"
	"xddq/pkg/game/protocol/pb"
	"xddq/pkg/utils"

	"go.uber.org/zap"
)

type RewardManager struct {
	a                *Assistant
	claimRewardJobId int

	lastAdRewardTime int64
	getAdRewardTimes int32

	getAdRewardC chan func()
}

func NewRewardManager(a *Assistant) *RewardManager {
	am := &RewardManager{
		a:            a,
		getAdRewardC: make(chan func(), 100),
	}

	a.addRunner(am)

	a.Client().OnGetAdRewardReq(am.onGetAdRewardResp)
	a.Client().OnPlayerAdRewardDataMsg(am.onPlayerAdRewardDataMsg)
	a.Client().OnReqBuyActivityGoods(am.onReqBuyActivityGoods)

	return am
}

const (
	AdRewardDailyMaxNum = 6
	AdRewardCd          = 5 * 60 * 1000
)

func (am *RewardManager) Run(ctx context.Context) error {
	am.a.L().Debug("reward manager running")

	am.claimRewardJobId = am.a.addJob("reward", "@every 5m", am.claimAdReward)
	go am.getAdRewardLoop(ctx)

	return nil
}

func (am *RewardManager) Stop() error {
	return nil
}

func (am *RewardManager) PushDelayFunc(fn func()) {
	am.getAdRewardC <- fn
}

func (am *RewardManager) getAdRewardLoop(ctx context.Context) error {

	for {
		select {
		case fn := <-am.getAdRewardC:
			fn()

			interval := 30 * time.Second
			if am.a.PlayerInfo().IsVip() {
				interval = 1 * time.Second
			}
			time.Sleep(interval)
		case <-ctx.Done():
			return nil
		}
	}
}

func (am *RewardManager) claimAdReward(ctx context.Context) error {
	if am.getAdRewardTimes >= AdRewardDailyMaxNum {
		am.a.L().Debug("reach max num, stop job", zap.Int32("getAdRewardTimes", am.getAdRewardTimes))
		am.a.removeJob(am.claimRewardJobId)
		return nil
	}

	am.a.Client().SendGetAdRewardReq(ctx, &pb.GetAdRewardReq{IsUseADTime: utils.Ptr(false)})

	return nil
}

func (am *RewardManager) onGetAdRewardResp(_ client.Context, msg *pb.GetAdRewardResp) error {
	am.a.L().Info(
		"onGetAdRewardResp",
		zap.Int32("getAdRewardTimes", am.getAdRewardTimes),
		zap.Int64("lastAdRewardTime", am.lastAdRewardTime),
		zap.Any("reward", db.ParseRewards(msg.GetReward())),
	)

	am.getAdRewardTimes++
	am.lastAdRewardTime = time.Now().Unix()

	return nil
}

// onPlayerAdRewardDataMsg
func (am *RewardManager) onPlayerAdRewardDataMsg(ctx client.Context, msg *pb.PlayerAdRewardDataMsg) error {
	am.lastAdRewardTime = msg.GetLastAdRewardTime()
	am.getAdRewardTimes = msg.GetGetAdRewardTimes()

	return nil
}

// onReqBuyActivityGoods
func (am *RewardManager) onReqBuyActivityGoods(_ client.Context, msg *pb.RspBuyActivityGoods) error {
	if msg.GetRet() != 0 {
		return db.GameError(msg)
	}

	am.a.L().Info("购买活动商品成功", zap.Any("rewards", db.ParseRewards(msg.GetRewards())))

	return nil
}
