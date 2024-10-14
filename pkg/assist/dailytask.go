package assist

import (
	"context"
	"xddq/pkg/game/client"
	"xddq/pkg/game/protocol/pb"
	"xddq/pkg/utils"
)

type DailyTaskManager struct {
	a *Assistant
}

func NewDailyTaskManager(a *Assistant) *DailyTaskManager {
	dt := &DailyTaskManager{
		a: a,
	}
	if a.Config().DailyTask {

		a.addRunner(dt)

		a.Client().OnGoodFortuneGetRewardReq(dt.onGoodFortuneGetRewardRsp)
	}

	return dt
}

func (dt *DailyTaskManager) Run(ctx context.Context) error {
	dt.a.L().Debug("daily task manager running")

	dt.checkin(ctx)
	dt.shareTask(ctx)

	dt.friendOneKey(ctx, 1)
	dt.friendOneKey(ctx, 2)

	// dt.a.Client().SendReqBuyActivityGoods(ctx, &pb.ReqBuyActivityGoods{
	// 	ActivityId: utils.Ptr(int32(dt.a.Activity().GetActivityIdByName("宝华堂"))),
	// 	MallId:     utils.Ptr(int32(400000003)),
	// 	Count:      utils.Ptr(int64(1)),
	// })

	return nil
}

// shareTask 分享任务
func (dt *DailyTaskManager) shareTask(ctx context.Context) error {
	if shareTask := dt.a.Storage().Cache().GetBool("share_task"); shareTask {
		return nil
	}
	defer dt.a.Storage().Cache().Set(ctx, "share_task", true)

	dt.a.L().Info("start share task")

	dt.a.Client().SendReqShareTaskDone21016(ctx, &pb.ReqShareTaskDone{
		ActivityId:  utils.Ptr(int32(0)),
		ConditionId: utils.Ptr(int32(0)),
	})
	dt.a.Client().SendReqShareTaskDone21030(ctx, &pb.ReqShareTaskDone{
		ActivityId:  utils.Ptr(int32(0)),
		ConditionId: utils.Ptr(int32(0)),
	})
	dt.a.Client().SendReqShareTaskDone21031(ctx, &pb.ReqShareTaskDone{
		ActivityId:  utils.Ptr(int32(0)),
		ConditionId: utils.Ptr(int32(0)),
	})

	return nil
}

// checkin
func (dt *DailyTaskManager) checkin(ctx context.Context) {
	key := "checkin"
	if checkin := dt.a.Storage().Cache().GetBool(key); checkin {
		return
	}
	defer dt.a.Storage().Cache().Set(ctx, key, true)

	dt.a.L().Info("start checkin")

	aid := dt.a.Activity().GetActivityIdByName("福泽")
	for i := range 8 {
		dt.a.Client().SendGoodFortuneGetRewardReq(ctx, &pb.GoodFortuneGetRewardReq{
			ActivityId:  utils.Ptr(aid),
			ConditionId: utils.Ptr(int32(10000 + i)),
			Type:        utils.Ptr(int32(1)),
		})
	}

	aid = dt.a.Activity().GetActivityIdByName("疯狂聚宝盆")
	for i := range 7 {
		dt.a.Client().SendReqGetConditionReward(ctx, &pb.ReqGetConditionReward{
			ActivityId:  utils.Ptr(int32(aid)),
			ConditionId: utils.Ptr(int32(10001 + i)),
		})
	}
}

func (dt *DailyTaskManager) onGoodFortuneGetRewardRsp(ctx client.Context, msg *pb.GoodFortuneGetRewardRsp) error {
	return nil
}

func (dt *DailyTaskManager) Stop() error {
	dt.a.L().Debug("daily task manager stop")

	return nil
}

func (dt *DailyTaskManager) friendOneKey(ctx context.Context, typ int32) error {
	return dt.a.Client().SendFriendReceiveAllReq(ctx, &pb.FriendReceiveAllReq{
		Type: &typ,
	})
}
