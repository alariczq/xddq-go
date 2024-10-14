package assist

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"xddq/pkg/game/client"
	"xddq/pkg/game/db"
	"xddq/pkg/game/protocol/pb"
	"xddq/pkg/utils"
)

type RoleManager struct {
	a *Assistant

	mu sync.RWMutex

	bigType       atomic.Int32
	littleType    atomic.Int32
	separationIdx atomic.Int32
	realmsId      atomic.Int32
	fightValue    atomic.Int64

	defaultSeparationIdx *int32

	dreamLv atomic.Int32
}

func NewRoleManager(a *Assistant) *RoleManager {
	am := &RoleManager{
		a: a,
	}

	if idx := a.Config().Role.SeparationIdx; idx != nil && (*idx == 0 || *idx == 1 || *idx == 2) {
		am.defaultSeparationIdx = idx
	}

	a.addRunner(am)

	a.Client().OnGetSeparationDataMsgListReq(am.onSeparationData)
	a.Client().OnDreamDataMsg(am.onDreamDataMsg)
	a.Client().OnPlayerAttributeDataMsg(am.onPlayerAttributeDataMsg)
	a.Client().OnPlayerSoulLiQuidDataMsg(am.onPlayerSoulLiQuidDataMsg)
	a.Client().OnPlayerMagicDataMsg(am.onPlayerMagicDataMsg)
	a.Client().OnSpiritPlayerDataMsg(am.onSpiritPlayerDataMsg)

	return am
}

func (rm *RoleManager) Run(ctx context.Context) error {
	if rm.defaultSeparationIdx != nil {
		rm.switchSeparation(ctx, *rm.defaultSeparationIdx)
	}
	return nil
}

func (rm *RoleManager) Stop() error {
	return nil
}

func (rm *RoleManager) onPlayerSoulLiQuidDataMsg(ctx client.Context, msg *pb.PlayerSoulLiQuidDataMsg) error {
	return nil
}

func (rm *RoleManager) onSeparationData(ctx client.Context, msg *pb.GetSeparationDataMsgListResp) error {
	return nil
}

func (rm *RoleManager) onDreamDataMsg(ctx client.Context, msg *pb.DreamDataMsg) error {
	rm.dreamLv.Store(msg.GetDreamLv())

	if time.Now().After(time.UnixMilli(msg.GetFreeSpeedUpCdEndTime())) {
		ctx.Client().SendDreamLvUpSpeedUpReq(ctx, &pb.DreamLvUpSpeedUpReq{
			SpeedUpType: utils.Ptr(int32(1)),
			UseTimes:    utils.Ptr(int32(1)),
			IsUseADTime: utils.Ptr(false),
		})
	}

	return nil
}

func (rm *RoleManager) onPlayerAttributeDataMsg(ctx client.Context, msg *pb.PlayerAttributeDataMsg) error {
	rm.fightValue.Store(msg.GetFightValue())
	rm.realmsId.Store(msg.GetRealmsId())
	rm.separationIdx.Store(msg.GetUseSeparationIdx())

	if realms := db.GetRealms(int(msg.GetRealmsId())); realms != nil {
		rm.bigType.Store(int32(realms.BigType))
		rm.littleType.Store(int32(realms.LittleType))
	}

	return nil
}

func (rm *RoleManager) SeparationIdx() int32 {
	return rm.separationIdx.Load()
}

func (rm *RoleManager) FightValue() int64 {
	return rm.fightValue.Load()
}

func (rm *RoleManager) BigType() int32 {
	return rm.bigType.Load()
}

func (rm *RoleManager) LittleType() int32 {
	return rm.littleType.Load()
}

func (rm *RoleManager) switchSeparation(ctx context.Context, idx int32) error {
	return rm.a.Client().SendSwitchSeparationReq(ctx, &pb.SwitchSeparationReq{
		SeparationIdx: &idx,
	})
}

func (rm *RoleManager) SwitchMagicSwitchPresetsReq(ctx client.Context, idx int32) error {
	return rm.a.Client().SendMagicSwitchPresetsReq(ctx, &pb.MagicSwitchPresetsReq{
		PresetsIndex: &idx,
	})
}

func (rm *RoleManager) SwitchBattleTeamReq(ctx client.Context, idx int32) error {
	return rm.a.Client().SendSwitchBattleTeamReq(ctx, &pb.SwitchBattleTeamReq{
		TeamNo: &idx,
	})
}

func (rm *RoleManager) SwitchSeparation(ctx context.Context, idx int32, fn func() error) error {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	recoverIdx := rm.SeparationIdx()
	defer func() {
		if rm.defaultSeparationIdx != nil {
			recoverIdx = *rm.defaultSeparationIdx
		}

		rm.switchSeparation(ctx, recoverIdx)
	}()

	if err := rm.switchSeparation(ctx, idx); err != nil {
		return err
	}

	return fn()
}

// onPlayerMagicDataMsg
func (rm *RoleManager) onPlayerMagicDataMsg(ctx client.Context, msg *pb.PlayerMagicDataMsg) error {
	return nil
}

// onSpiritPlayerDataMsg
func (rm *RoleManager) onSpiritPlayerDataMsg(ctx client.Context, msg *pb.SpiritPlayerDataMsg) error {
	return nil
}

// onMagicTreasurePlayerDataMsg
func (rm *RoleManager) onMagicTreasurePlayerDataMsg(ctx client.Context, msg *pb.MagicTreasurePlayerDataMsg) error {
	return nil
}

// onPlayerPetDataSync
func (rm *RoleManager) onPlayerPetDataSync(ctx client.Context, msg *pb.PlayerPetDataSync) error {
	return nil
}
