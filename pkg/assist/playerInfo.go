package assist

import (
	"time"
	"xddq/pkg/game/client"
	"xddq/pkg/game/protocol/pb"
	"xddq/pkg/logutil"

	"go.uber.org/zap"
)

type PlayerInfoManager struct {
	a          *Assistant
	playerInfo *pb.PlayerDataMsg

	isMonthCardVip bool
	isYearCardVip  bool
}

func NewPlayerInfoManager(a *Assistant) *PlayerInfoManager {
	p := &PlayerInfoManager{
		a: a,
	}

	a.Client().OnPlayerDataMsg(p.onPlayerDataMsg)
	a.Client().OnPrivilegeCardDataMsg(p.onPrivilegeCardDataMsg)

	return p
}

func (p *PlayerInfoManager) onPlayerDataMsg(ctx client.Context, msg *pb.PlayerDataMsg) error {
	p.playerInfo = msg

	p.a.SetLogger(logutil.DefaultLogger.With(
		zap.String("nickname", p.Nickname()),
		zap.String("username", p.a.Config().Username),
	))

	return nil
}

func (p *PlayerInfoManager) ServerId() int64 {
	return p.playerInfo.GetServerId()
}

func (p *PlayerInfoManager) UnionId() int64 {
	return p.playerInfo.GetUnionId()
}

func (p *PlayerInfoManager) Nickname() string {
	return p.playerInfo.GetNickName()
}

func (p *PlayerInfoManager) PlayerId() int64 {
	return p.playerInfo.GetPlayerId()
}

func (p *PlayerInfoManager) IsMonthCardVip() bool {
	return p.isMonthCardVip
}

func (p *PlayerInfoManager) IsYearCardVip() bool {
	return p.isYearCardVip
}

func (p *PlayerInfoManager) IsVip() bool {
	return p.isMonthCardVip || p.isYearCardVip
}

// privilegeCardReceiveReward
func (p *PlayerInfoManager) privilegeCardReceiveReward(ctx client.Context, typ int32) error {
	return p.a.Client().SendPrivilegeCardReceiveRewardReq(ctx, &pb.PrivilegeCardReceiveRewardReq{
		Type: &typ,
	})
}

// onPrivilegeCardDataMsg
func (p *PlayerInfoManager) onPrivilegeCardDataMsg(ctx client.Context, msg *pb.PrivilegeCardDataMsg) error {
	p.isYearCardVip = msg.GetMonthlyCardEndTime() != 0
	p.isMonthCardVip = time.Now().Before(time.UnixMilli(msg.GetMonthlyCardEndTime()))
	if p.isMonthCardVip {
		p.a.L().Info("检测到月卡", zap.String("endTime", time.UnixMilli(msg.GetMonthlyCardEndTime()).Format("2006-01-02 15:04:05")))
	}
	if p.isYearCardVip {
		p.a.L().Info("检测到年卡")
	}

	if !IsToday(msg.GetGetMonthlyCardRewardTime()) {
		p.privilegeCardReceiveReward(ctx, 1)
	}

	if !IsToday(msg.GetGetYearCardRewardTime()) {
		p.privilegeCardReceiveReward(ctx, 2)
	}

	return nil
}
