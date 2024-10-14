package assist

import (
	"context"
	"slices"

	"xddq/pkg/game/client"
	"xddq/pkg/game/db"
	"xddq/pkg/game/protocol"
	"xddq/pkg/game/protocol/pb"
	"xddq/pkg/utils"

	"github.com/samber/lo"
	"go.uber.org/zap"
)

type TowerManager struct {
	a         *Assistant
	jobId     int
	onlyMax   bool
	curPassId int32
	passMaxId int32

	selectedBuff []*Skill
}

func NewTowerManager(a *Assistant) *TowerManager {
	t := &TowerManager{
		a:       a,
		onlyMax: true,
	}

	if !a.Config().Challenge.Tower {
		return t
	}

	a.addRunner(t)

	a.Client().OnTowerDataMsg(t.onTowerDataSync)
	a.Client().OnTowerChallengeResp(t.onTowerChallengeResp)
	a.Client().OnSelectBuffReq(t.onSelectBuffResp)

	return t
}

func (t *TowerManager) doChallenge(ctx context.Context) error {
	if t.onlyMax && t.curPassId == t.passMaxId {
		t.Stop()
		return nil
	}

	return t.challengeWithSeparation(ctx)
}

func (t *TowerManager) Run(ctx context.Context) error {
	t.challengeWithSeparation(ctx)
	t.jobId = t.a.addJob("tower", "@every 10s", t.doChallenge)
	return nil
}

func (t *TowerManager) Stop() error {
	t.a.L().Info("tower stop")

	t.a.removeJob(t.jobId)
	return nil
}

func (t *TowerManager) challengeWithSeparation(ctx context.Context) (err error) {
	return t.a.Role().SwitchSeparation(ctx, t.a.Config().Challenge.TowerIdx, func() error {
		return t.challenge(ctx)
	})
}

func (t *TowerManager) challenge(ctx context.Context) error {
	if t.curPassId == 0 && !t.a.Palace().HasMiracle() {
		return nil
	}

	if t.curPassId == 0 {
		t.a.Client().SendSaveSelectOptionReq(ctx, &pb.SaveSelectOptionReq{
			MarkPreference: []*pb.MarkupPreference{
				{Priority: utils.Ptr(int32(1)), SkillType: utils.Ptr(int32(1017))},
				{Priority: utils.Ptr(int32(2)), SkillType: utils.Ptr(int32(1018))},
				{Priority: utils.Ptr(int32(3)), SkillType: utils.Ptr(int32(1023))},
				{Priority: utils.Ptr(int32(4)), SkillType: utils.Ptr(int32(1001))},
				{Priority: utils.Ptr(int32(5)), SkillType: utils.Ptr(int32(1024))},
			},
		})
		t.quickChallenge(ctx)
	}

	return t.a.sendWithMsgId(ctx, protocol.TOWER_CHALLENGE, nil)
}

func (t *TowerManager) onTowerChallengeResp(ctx client.Context, msg *pb.TowerChallengeResp) error {
	if msg.GetRet() != 0 {
		return db.GameError(msg)
	}

	towerData := msg.TowerDataSync
	lg := t.a.L().With(
		zap.Int64("sepration", int64(t.a.Role().SeparationIdx())),
		zap.Int64("cur", int64(towerData.GetCurPassId())),
		zap.Int64("max", int64(towerData.GetPassMaxId())),
	)

	t.onTowerDataSync(ctx, msg.TowerDataSync)

	if msg.GetChallengeSuccess() {
		lg.Info("✅ 镇妖塔挑战成功", zap.Any("rewards", db.ParseRewards(msg.GetRewards())))
		return nil
	}

	lg.Info("❌ 镇妖塔挑战失败")

	return nil
}

func (t *TowerManager) onTowerDataSync(ctx client.Context, msg *pb.TowerDataMsg) error {
	if msg == nil {
		return nil
	}

	t.syncState(ctx, msg)

	return nil
}

// onQuickChallengeResp
func (t *TowerManager) onQuickChallengeResp(ctx client.Context, msg *pb.QuickChallengeResp) error {
	if msg.GetRet() != 0 {
		return db.GameError(msg)
	}
	if msg.GetReward() != "" {
		t.a.L().Info(
			"✅ 镇妖塔快速挑战成功",
			zap.Stringer("rewards", db.ParseRewards(msg.GetReward())),
		)
	}
	if msg.GetMiracleReward() != "" {
		t.a.L().Info(
			"✅ 镇妖塔奇迹奖励",
			zap.Stringer("rewards", db.ParseRewards(msg.GetMiracleReward())),
		)
	}

	t.syncState(ctx, msg.GetTowerDataSync())

	return nil
}

func (t *TowerManager) quickChallenge(ctx context.Context) error {
	return t.a.Client().Send(ctx, protocol.TOWER_QUICK_CHANLLENGE, nil)
}

var selectBuffs = []string{"最终减伤", "强化灵兽", "攻击", "最终增伤"}

func (t *TowerManager) selectBuff(ctx context.Context, buffs []int32) error {
	if len(buffs) == 0 {
		return nil
	}

	skills := lo.Map(buffs, func(item int32, i int) lo.Tuple2[int, *db.GameSkill] {
		return lo.T2(i, db.GetGameSkill(int(item)))
	})

	slices.SortFunc(skills, func(a, b lo.Tuple2[int, *db.GameSkill]) int {
		askill := a.B
		bSkill := b.B

		return slices.Index(selectBuffs, bSkill.GetName()) - slices.Index(selectBuffs, askill.GetName())
	})

	for _, tuple := range skills {
		i, skill := tuple.Unpack()
		if skill.Star > 2 {
			continue
		}

		if !slices.Contains(selectBuffs, skill.GetName()) {
			continue
		}

		return t.a.Client().SendSelectBuffReq(ctx, &pb.SelectBuffReq{
			Index: utils.Ptr(int32(i)),
		})
	}

	return t.a.Client().SendSelectBuffReq(ctx, &pb.SelectBuffReq{
		Index: utils.Ptr(int32(-1)),
	})
}

func (t *TowerManager) syncState(ctx context.Context, state *pb.TowerDataMsg) {
	t.curPassId = state.GetCurPassId()
	t.passMaxId = state.GetPassMaxId()
	t.selectedBuff = lo.Map(state.GetBuffData(), func(item *pb.BuffDataMsg, _ int) *Skill {
		skill := db.GetGameSkill(int(*item.SkillId))
		return &Skill{
			Id:   item.GetSkillId(),
			Name: skill.GetName(),
			Lv:   item.GetLv(),
		}
	})
	t.selectBuff(ctx, state.GetPendingBuffPool().GetBuffSkillId())
}

func (t *TowerManager) onSelectBuffResp(ctx client.Context, msg *pb.SelectBuffResp) error {
	if msg.GetRet() != 0 {
		return db.GameError(msg)
	}

	t.syncState(ctx, msg.GetTowerDataSync())

	return nil
}

type Skill struct {
	Id   int32
	Name string
	Lv   int32
}
