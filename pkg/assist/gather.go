package assist

import (
	"context"
	"slices"
	"sync/atomic"
	"time"
	"xddq/pkg/game/client"
	"xddq/pkg/game/db"
	"xddq/pkg/game/protocol/pb"
	"xddq/pkg/utils"

	"github.com/golang-module/carbon/v2"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type GatherManager struct {
	a *Assistant

	checkJobId int
	scanJobId  int

	attendNum int32
	openNum   int32

	scanning  atomic.Bool
	minIncome int64
}

func NewGatherManager(a *Assistant) *GatherManager {
	gm := &GatherManager{
		a: a,
	}

	conf := gm.a.Config().Gather
	if !conf.Enabled {
		gm.a.L().Debug("gather manager disabled")
		return gm
	}

	gm.minIncome = conf.MinIncome

	a.addRunner(gm)

	a.Client().OnGatherEnergyEnterNewReq(gm.onGatherEnergyEnterNewResp)
	a.Client().OnGatherEnergyFirstListViewReq(gm.onGatherEnergyFirstListViewResp)
	a.Client().OnGatherEnergyAttendNewReq(gm.onGatherEnergyAttendNewResp)
	a.Client().OnSyncGatherEnergyMsg(gm.onSyncGatherEnergyMsg)
	return gm
}

func (gm *GatherManager) Run(ctx context.Context) error {
	gm.a.L().Debug("gather manager running")

	gm.check(ctx)
	gm.checkJobId = gm.a.addJob("gather", "@every 1m", gm.check)

	gm.a.L().Info("聚灵阵最低收益!!!", zap.Int64("minIncome", gm.minIncome))

	return nil
}

func (gm *GatherManager) Stop() error {
	gm.a.L().Debug("gather manager stop")

	gm.stopCheck()
	gm.stopScan()
	return nil
}

// 检查状态
func (gm *GatherManager) check(ctx context.Context) error {
	if now := time.Now(); now.Hour() < 10 && now.Hour() > 21 {
		return nil
	}

	gm.enter(ctx)

	return nil
}

func (gm *GatherManager) scan(ctx context.Context) error {
	return gm.list(ctx)
}

func (gm *GatherManager) startScan() {
	if gm.scanning.Load() {
		return
	}
	gm.scanning.Store(true)

	// 扫描列表
	gm.scanJobId = gm.a.addJob("gatherScan", "@every 1s", gm.scan)
}

func (gm *GatherManager) stopScan() {
	if !gm.scanning.Load() {
		return
	}
	gm.scanning.Store(false)

	if gm.scanJobId == 0 {
		return
	}

	gm.a.removeJob(gm.scanJobId)
	gm.scanJobId = 0
}

func (gm *GatherManager) stopCheck() {
	gm.a.removeJob(gm.checkJobId)
	gm.checkJobId = 0
}

func (gm *GatherManager) list(ctx context.Context) error {
	return gm.a.Client().SendGatherEnergyFirstListViewReq(ctx, &pb.GatherEnergyFirstListViewReq{
		Offset:     utils.Ptr(int32(0)),
		FilterType: utils.Ptr(int32(1)),
	})
}

// onGatherEnergyFirstListViewResp
func (gm *GatherManager) onGatherEnergyFirstListViewResp(ctx client.Context, msg *pb.GatherEnergyFirstListViewResp) error {
	if msg.GetRet() != 0 {
		return db.GetGameError(msg.GetRet())
	}

	now := time.Now()
	today2130 := carbon.Now().SetHour(21).SetMinute(30).StdTime()
	today2125 := carbon.Now().SetHour(21).SetMinute(25).StdTime()

	list := lo.Filter(msg.List, func(item *pb.GatherEnergyFirstListMsg, _ int) bool {
		if item.GetListSize() >= 5 {
			return false
		}

		energy := item.GetEnergyBaseMsg()
		endTime := time.UnixMilli(energy.GetEndTime())

		if now.Before(today2130) {
			if endTime.After(today2125) {
				return false
			}

			if gm.minIncome > 0 {
				return energy.GetIncome() >= gm.minIncome
			}

			nowHour := now.Hour()
			switch {
			case nowHour == 10:
				return energy.GetIncome() >= 600
			case nowHour == 11 || nowHour == 12:
				return energy.GetIncome() >= 500
			default:
				return energy.GetIncome() >= 468
			}
		}

		if endTime.Sub(now) < 12*time.Hour {
			return false
		}

		if gm.minIncome > 0 {
			return energy.GetIncome() >= gm.minIncome
		}

		nowMin := now.Minute()
		switch {
		case nowMin < 45:
			return energy.GetIncome() >= 700
		case nowMin < 50:
			return energy.GetIncome() >= 600
		case nowMin < 55:
			return energy.GetIncome() >= 500
		case nowMin < 59:
			return energy.GetIncome() >= 468
		default:
			return energy.GetIncome() >= 400
		}
	})

	slices.SortFunc(list, func(a, b *pb.GatherEnergyFirstListMsg) int {
		age := a.GetEnergyBaseMsg()
		bge := b.GetEnergyBaseMsg()
		aMin := time.UnixMilli(age.GetEndTime()).Sub(now).Minutes()
		bMin := time.UnixMilli(bge.GetEndTime()).Sub(now).Minutes()

		return int(bMin*float64(bge.GetIncome())) - int(aMin*float64(age.GetIncome()))
	})

	for _, item := range list {
		if gm.attendNum > 0 {
			break
		}

		gm.a.L().Info("尝试进入聚灵阵", zap.String("nickname", item.OpenerMsg.GetNickName()))

		gm.attend(ctx, item.GetEnergyBaseMsg().GetId())

		break
	}

	return nil
}

func (gm *GatherManager) like(ctx context.Context) error {
	return gm.a.Client().SendGatherEnergyLikeReq(ctx, &pb.GatherEnergyLikeReq{})
}

func (gm *GatherManager) getReward(ctx context.Context, id int64) error {
	return gm.a.Client().SendGatherEnergyRewardReq(ctx, &pb.GatherEnergyRewardReq{Id: &id})
}

func (gm *GatherManager) getADReward(ctx context.Context) error {
	return gm.a.Client().SendGatherEnergyGetADAwardReq(ctx, &pb.GatherEnergyGetADAwardReq{IsUseADTime: utils.Ptr(false)})
}

func (gm *GatherManager) attend(ctx context.Context, id int64) error {
	return gm.a.Client().SendGatherEnergyAttendNewReq(ctx, &pb.GatherEnergyAttendNewReq{
		Id: utils.Ptr(id),
	})
}

func (gm *GatherManager) onGatherEnergyAttendNewResp(ctx client.Context, msg *pb.GatherEnergyAttendNewResp) error {
	if msg.GetRet() != 0 {
		return db.GetGameError(msg.GetRet())
	}

	detail := msg.GetDetailInfo()

	var owner string
	if teamList := detail.GetTeamList(); len(teamList) > 0 {
		owner = *teamList[0].GetPlayerInfo().NickName
	}

	ge := detail.GetEnergyBaseMsg()
	gm.a.L().Info(
		"进入聚灵阵",
		zap.String("owner", owner),
		zap.Int64("income", ge.GetIncome()),
		zap.Duration("duration", time.Until(time.UnixMilli(ge.GetEndTime()))),
	)

	gm.check(ctx)

	return nil
}

func (gm *GatherManager) enter(ctx context.Context) error {
	return gm.a.Client().SendGatherEnergyEnterNewReq(ctx, &pb.GatherEnergyEnterNewReq{
		TotalNum: utils.Ptr(int32(0)),
	})
}

// onGatherEnergyEnterNewResp
func (gm *GatherManager) onGatherEnergyEnterNewResp(ctx client.Context, msg *pb.GatherEnergyEnterNewResp) error {
	if msg.GetRet() != 0 {
		return db.GetGameError(msg.GetRet())
	}
	defer func() {
		gm.a.Client().SendGatherEnergyLeaveReq(ctx, &pb.GatherEnergyLeaveReq{})
	}()

	ge := msg.GetGatherEnergy()
	gm.openNum = ge.GetOpenNum()
	gm.attendNum = ge.GetAttendNum()

	for i := ge.GetGetTimes(); i < 3; i++ {
		gm.getADReward(ctx)
	}

	for _, item := range msg.GetReward() {
		if attacker := item.GetAttackerInfo(); attacker != nil {
			gm.a.L().Warn("聚灵阵被驱逐", zap.String("nickname", attacker.GetNickName()))
		}
		gm.getReward(ctx, item.GetId())
	}

	if !ge.GetHadLike() {
		gm.like(ctx)
	}

	now := time.Now()
	hour := now.Hour()

	if gm.attendNum == 0 && (hour >= 10 && hour < 22) {
		gm.startScan()
	}

	if (gm.attendNum > 0 || hour < 10 || hour >= 22) && gm.scanning.Load() {
		gm.a.L().Info("已加入聚灵阵，停止扫描")
		gm.stopScan()
	}

	return nil
}

// onSyncGatherEnergyMsg
func (gm *GatherManager) onSyncGatherEnergyMsg(ctx client.Context, msg *pb.SyncGatherEnergyMsg) error {
	return nil
}
