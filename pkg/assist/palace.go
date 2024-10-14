package assist

import (
	"context"

	"xddq/pkg/game/client"
	"xddq/pkg/game/db"
	"xddq/pkg/game/protocol/pb"
	"xddq/pkg/utils"

	"github.com/samber/lo"
	"go.uber.org/zap"
)

type PalaceManager struct {
	a          *Assistant
	hasMiracle bool
}

func NewPalaceManager(a *Assistant) *PalaceManager {
	p := &PalaceManager{
		a: a,
	}

	if a.Config().Palace {
		a.addRunner(p)

		a.Client().OnSendGiftSyncMsg(p.onSendGiftSyncMsg)
		a.Client().OnGetSendGiftRewardReq(p.onGetSendGiftRewardRsp)
		a.Client().OnPalaceSyncMsg(p.onPalaceSyncMsg)
		a.Client().OnEnterPalaceReq(p.onEnterPalaceRsp)
		a.Client().OnPalaceMiracleDataMsg(p.onPalaceMiracleDataMsg)
		a.Client().OnPalaceWorshipReq(p.onPalaceWorshipRsp)
	}

	return p
}

func (pm *PalaceManager) Run(ctx context.Context) error {
	pm.a.L().Debug("palace manager running")

	pm.a.addJob("check_enter_palace_0_hour", "0 0 0 * * *", func(ctx context.Context) error {
		return pm.a.Client().SendEnterPalaceReq(ctx, &pb.EnterPalaceReq{})
	})

	pm.a.addJob("check_enter_palace_22_hour", "0 0 22 * * *", func(ctx context.Context) error {
		return pm.a.Client().SendEnterPalaceReq(ctx, &pb.EnterPalaceReq{})
	})

	return nil
}

func (pm *PalaceManager) Stop() error {
	pm.a.L().Debug("palace manager stop")

	return nil
}

func (pm *PalaceManager) addRecivedGiftTimes(ctx context.Context) {
	n := pm.getRecivedGiftTimes()
	pm.a.Storage().Cache().Set(ctx, "palace_worship_times", n+1)
}

func (pm *PalaceManager) getRecivedGiftTimes() int64 {
	return pm.a.Storage().Cache().GetInt64("palace_worship_times")
}

func (pm *PalaceManager) onSendGiftSyncMsg(ctx client.Context, resp *pb.SendGiftSyncMsg) error {
	for _, v := range resp.GetData() {
		p := v.GetPlayerData()

		pm.addRecivedGiftTimes(ctx)

		pm.a.L().Info(
			"🎁 收到礼物",
			zap.String("nickname", p.GetNickName()),
			zap.String("giftType", v.GetType().String()),
			zap.Int64("times", pm.getRecivedGiftTimes()),
		)

		return pm.a.Client().SendGetSendGiftRewardReq(ctx, &pb.GetSendGiftRewardReq{
			Id:        utils.Ptr(v.GetId()),
			GetReward: utils.Ptr(true),
			Type:      utils.Ptr(v.GetType()),
		})
	}

	return nil
}

// GetSendGiftRewardRsp
func (pm *PalaceManager) onGetSendGiftRewardRsp(ctx client.Context, resp *pb.GetSendGiftRewardRsp) error {
	rewards := db.ParseRewards(resp.GetReward())

	if len(rewards) > 0 {
		reward := rewards[0]
		n := pm.a.Storage().Cache().GetInt64("gift_num")
		pm.a.Storage().Cache().Set(ctx, "gift_num", n+int64(reward.Num))
	}

	pm.a.L().Info(
		"获取送礼奖励",
		zap.Any("reward", rewards),
		zap.Int64("total", pm.a.Storage().Cache().GetInt64("gift_num")),
	)

	return nil
}

func (pm *PalaceManager) onPalaceSyncMsg(ctx client.Context, resp *pb.PalaceSyncMsg) error {
	pm.a.L().Info("同步神殿数据", zap.Any("resp", resp))
	lo.ForEach(resp.GetCanWorshipTitle(), func(id int32, _ int) {
		pm.placeWorship(ctx, id, 0)
	})

	return nil
}

// PalaceWorshipRsp
func (pm *PalaceManager) onPalaceWorshipRsp(ctx client.Context, resp *pb.PalaceWorshipRsp) error {
	pm.a.L().Info("神殿朝拜", zap.Any("resp", resp))
	return nil
}

// EnterPalaceRsp
func (pm *PalaceManager) onEnterPalaceRsp(ctx client.Context, resp *pb.EnterPalaceRsp) error {
	if resp.GetWorship() && resp.GetWorshipRandom() {
		return pm.placeWorship(ctx, 0, 1)
	}

	lo.ForEach(resp.GetPalaceData(), func(palace *pb.PalaceDataMsg, _ int) {
		if palace.GetWorship() {
			pm.placeWorship(ctx, palace.GetTitleId(), 0)
		}
	})

	return nil
}

// ctx.Send(ctx, &pb.PalaceWorshipReq{ TitleId:  0, IsRandom: 1, })
func (pm *PalaceManager) placeWorship(ctx client.Context, titleId int32, isRandom int32) error {
	return pm.a.Client().SendPalaceWorshipReq(ctx, &pb.PalaceWorshipReq{TitleId: &titleId, IsRandom: &isRandom})
}

func (pm *PalaceManager) HasMiracle() bool {
	return pm.hasMiracle
}

// PalaceMiracleDataMsg
func (pm *PalaceManager) onPalaceMiracleDataMsg(ctx client.Context, msg *pb.PalaceMiracleDataMsg) error {
	pm.hasMiracle = msg.GetMiracleId() != 0

	if msg.GetMiracleId() == 0 {
		pm.a.L().Info("尝试今日点赞")
		return pm.placeWorship(ctx, 0, 1)
	}

	return nil
}
