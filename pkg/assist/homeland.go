package assist

import (
	"context"
	"time"
	"xddq/pkg/game/client"
	"xddq/pkg/game/protocol/pb"
	"xddq/pkg/utils"

	"go.uber.org/zap"
)

type HomelandManager struct {
	a *Assistant

	totalWorkerNum int32
	freeWorkerNum  int32
	energy         int32

	lastRefresh time.Time
}

func NewHomelandManager(a *Assistant) *HomelandManager {
	h := &HomelandManager{
		a: a,
	}

	if !a.Config().Homeland.Enabled {
		return h
	}

	a.addRunner(h)

	a.Client().OnSyncHomelandMsg(h.onSyncHomeland)
	a.Client().OnSyncHomelandHasRewardMsg(h.onSyncHomelandHasRewardMsg)
	a.Client().OnHomelandManageReq(h.onManageHomelandResp)
	a.Client().OnHomelandEnterReq(h.onEnterHomelandResp)
	a.Client().OnHomelandExploreReq(h.onExploreResp)

	return h
}

func (h *HomelandManager) Run(ctx context.Context) error {
	h.a.addJob("homelandManage", "@every 1s", h.explore)

	return nil
}

func (h *HomelandManager) Stop() error {
	return nil
}

func (h *HomelandManager) onSyncHomeland(ctx client.Context, msg *pb.SyncHomelandMsg) error {
	h.a.L().Info("福地信息同步", zap.Any("msg", msg))
	h.totalWorkerNum = msg.GetTotalWorkerNum()
	h.freeWorkerNum = msg.GetFreeWorkerNum()
	h.energy = msg.GetEnergy()

	return nil
}

func (h *HomelandManager) scanHomeland(ctx context.Context) error {
	// h.manageHomeland()
	h.enter(h.a.PlayerInfo().PlayerId())

	return nil
}

// explore 探索
func (h *HomelandManager) explore(ctx context.Context) error {
	return h.a.Client().SendHomelandExploreReq(ctx, &pb.HomelandExploreReq{})
}

// manage homeland
func (h *HomelandManager) manage() error {
	return h.a.Client().SendHomelandManageReq(h.a.Context(), &pb.HomelandManageReq{})
}

// enter homeland 进入福地
func (h *HomelandManager) enter(playerId int64) error {
	return h.a.Client().SendHomelandEnterReq(h.a.Context(), &pb.HomelandEnterReq{PlayerId: &playerId})
}

// dispatch 派遣
func (h *HomelandManager) dispatch(playerId int64, pos int32, workerNum int32) error {
	return h.a.Client().SendHomelandDispatchWorkerReq(h.a.Context(), &pb.HomelandDispatchWorkerReq{
		PlayerId:  &playerId,
		Pos:       &pos,
		WorkerNum: &workerNum,
	})
}

// 撤回
func (h *HomelandManager) reset(playerId int64, pos int32) error {
	return h.dispatch(playerId, pos, 0)
}

// 刷新探索
func (h *HomelandManager) refreshExplore() error {
	return h.a.Client().SendHomelandRefreshExploreReq(h.a.Context(), &pb.HomelandRefreshExploreReq{})
}

// 刷新资源
func (h *HomelandManager) refresh() error {
	return h.a.Client().SendHomelandRefreshReq(h.a.Context(), &pb.HomelandRefreshReq{
		Type:        utils.Ptr(int32(1)),
		Position:    utils.Ptr(int32(-1)),
		ItemId:      utils.Ptr(int32(0)),
		IsUseADTime: utils.Ptr(false),
	})
}

// onExploreResp
func (h *HomelandManager) onExploreResp(ctx client.Context, msg *pb.HomelandExploreResp) error {
	// exploreData := msg.GetExploreData()

	return nil
}

func (h *HomelandManager) onManageHomelandResp(ctx client.Context, msg *pb.HomelandManageResp) error {
	return nil
}

func (h *HomelandManager) onRefresh(ctx client.Context, msg *pb.HomelandRefreshResp) error {
	return nil
}

// onEnterHomelandResp
func (h *HomelandManager) onEnterHomelandResp(ctx client.Context, msg *pb.HomelandEnterResp) error {
	h.a.L().Info("进入福地")
	// hld := msg.GetHomeland()
	return nil
}

// onSyncHomelandHasRewardMsg
func (h *HomelandManager) onSyncHomelandHasRewardMsg(ctx client.Context, msg *pb.SyncHomelandHasRewardMsg) error {
	h.a.L().Info("福地有奖励可领取同步")
	return nil
}
