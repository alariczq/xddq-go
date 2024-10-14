package assist

import (
	"context"
	"xddq/pkg/game/client"
	"xddq/pkg/game/db"
	"xddq/pkg/game/protocol"
	"xddq/pkg/game/protocol/pb"
	"xddq/pkg/utils"

	"go.uber.org/zap"
)

type ChallengeManager struct {
	a *Assistant

	// 挑战妖王
	willBossPassId int32
	willBossTimes  int32

	// 冒险章节
	stagePassId int32

	// StarTrial
	starTrialBossId int32
}

func NewChallengeManager(a *Assistant) *ChallengeManager {
	cm := &ChallengeManager{
		a: a,
	}

	if !a.Config().Challenge.Enabled {
		return cm
	}

	a.addRunner(cm)

	a.Client().OnInvadeDataMsg(cm.onInvadeDataMsg)
	a.Client().OnPlayerStageData(cm.onPlayerStageData)
	a.Client().OnChallengeRspMsg(cm.onChallengeRspMsg)
	a.Client().OnSynSecretTowerInfo(cm.onSynSecretTowerInfo)
	a.Client().OnSynHeroRankPlayerInfo(cm.onSynHeroRankPlayerInfo)

	// 秘境
	a.Client().OnSecretTowerFightReq(cm.onSecretTowerFightResp)

	// StarTrial
	a.Client().OnStarTrialDataMsg(cm.onStarTrialDataMsg)

	// 妖王
	a.Client().OnWildBossDataSync(cm.onWildBossDataSync)

	// 法则试练
	a.Client().OnRuleTrialDataSync(cm.onRuleTrialDataSync)

	return cm
}

func (cm *ChallengeManager) Run(ctx context.Context) error {
	if cm.a.Config().Challenge.StageChallenge {
		cm.a.addJob("chapterChallenge", "@every 10s", cm.stageChallenge)
	}

	if cm.a.Config().Challenge.SecretTower {
		cm.a.addJob("secretTower", "@every 10s", cm.secretTowerChallenge)
	}

	return nil
}

func (cm *ChallengeManager) Stop() error {
	return nil
}

// 冒险章节挑战
func (cm *ChallengeManager) stageChallenge(ctx context.Context) error {
	return cm.a.Role().SwitchSeparation(ctx, cm.a.Config().Challenge.StageIdx, func() error {
		return cm.a.Client().Send(ctx, protocol.STAGE_CHALLENGE, nil)
	})
}

// 异兽入侵

// invadeChallenge
func (cm *ChallengeManager) invadeChallenge(ctx context.Context) {
	cm.a.sendWithMsgId(ctx, protocol.INVADE_CHALLENGE, nil)
}

// onInvadeDataMsg
func (cm *ChallengeManager) onInvadeDataMsg(ctx client.Context, msg *pb.InvadeDataMsg) error {
	return nil
}

// onInvadeChallengeResp
func (cm *ChallengeManager) onInvadeChallengeResp(ctx client.Context, msg *pb.InvadeChallengeResp) error {
	return nil
}

// onPlayerStageData
func (cm *ChallengeManager) onPlayerStageData(ctx client.Context, msg *pb.PlayerStageData) error {
	cm.stagePassId = msg.GetPassStageId()
	cm.a.L().Info("冒险进度", zap.Int32("chapterStage", cm.stagePassId))
	return nil
}

// onChallengeRspMsg
func (cm *ChallengeManager) onChallengeRspMsg(ctx client.Context, msg *pb.ChallengeRspMsg) error {
	if msg.GetRet() != 0 {
		return db.GameError(msg)
	}
	lg := cm.a.L().With(
		zap.Int64("sepration", int64(cm.a.Role().SeparationIdx())),
	)

	if msg.GetChallengeSuccess() {
		lg.Info(
			"✅ 冒险挑战成功",
			zap.String("rewards", db.ParseRewards(msg.GetRewards()).String()),
		)
		return nil
	}

	lg.Info("❌ 冒险挑战失败")

	return nil
}

// 秘境

func (cm *ChallengeManager) secretTowerChallenge(ctx context.Context) error {
	return cm.a.Role().SwitchSeparation(ctx, cm.a.Config().Challenge.SecretTowerIdx, func() error {
		return cm.a.Client().SendSecretTowerFightReq(ctx, &pb.SecretTowerFightReq{
			Type: utils.Ptr(int32(1)),
		})
	})
}

func (cm *ChallengeManager) onSecretTowerFightResp(_ client.Context, msg *pb.SecretTowerFightResp) error {
	if msg.GetRet() != 0 {
		return db.GameError(msg)
	}
	lg := cm.a.L().With(
		zap.Int64("sepration", int64(cm.a.Role().SeparationIdx())),
	)

	if msg.GetRewards() != "" {
		lg.Info(
			"✅ 真火秘境挑战成功",
			zap.String("rewards", db.ParseRewards(msg.GetRewards()).String()),
		)
		return nil
	}

	lg.Info("❌ 真火秘境挑战失败")

	return nil
}

// onSynSecretTowerInfo
func (cm *ChallengeManager) onSynSecretTowerInfo(ctx client.Context, msg *pb.SynSecretTowerInfo) error {
	return nil
}

// 群英榜

// onSynHeroRankPlayerInfo
func (cm *ChallengeManager) onSynHeroRankPlayerInfo(ctx client.Context, msg *pb.SynHeroRankPlayerInfo) error {
	return nil
}

// 星宿试炼

// onStarTrialDataMsg
func (cm *ChallengeManager) onStarTrialDataMsg(ctx client.Context, msg *pb.StarTrialDataMsg) error {
	if msg.GetRewardState() == 1 {
		cm.a.Client().SendGetDailyFightRewardReq(ctx, &pb.GetDailyFightRewardReq{})
	}

	if msg.GetChallengeTimes() <= 20 {
		return nil
	}

	if msg.GetBossId() == msg.GetPrevBossId() {
		return nil
	}

	cm.a.Client().SendStarTrialChallengeReq(ctx, &pb.StarTrialChallengeReq{
		BossId: msg.BossId,
	})

	return nil
}

// onRuleTrialDataSync 法则试练速战数据同步
func (dt *ChallengeManager) onRuleTrialDataSync(ctx client.Context, msg *pb.RuleTrialDataSync) error {
	if !msg.GetIsRepeated() {
		return dt.a.Client().SendRuleTrialRepeatAllReq(ctx, &pb.RuleTrialRepeatAllReq{})
	}
	return nil
}
