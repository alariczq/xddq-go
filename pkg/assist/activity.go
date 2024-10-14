package assist

import (
	"context"
	"fmt"
	"sync"
	"time"

	"xddq/pkg/game/client"
	"xddq/pkg/game/db"
	"xddq/pkg/game/protocol/pb"
	"xddq/pkg/utils"

	"go.uber.org/zap"
)

type ActivityConfig struct {
	*pb.ActivityMainConfig `json:"-"`
	Id                     int32     `json:"id"`
	Name                   string    `json:"name"`
	Type                   int32     `json:"type"`
	GroupType              int32     `json:"group_type"`
	BeginTime              time.Time `json:"begin_time"`
	EndTime                time.Time `json:"end_time"`
}

type ActivityManager struct {
	a *Assistant

	mu               sync.RWMutex
	systemActivityId map[int32]*ActivityConfig
	activitiesDetail map[int32]*pb.ActivityCommonData
}

func NewActivityManager(a *Assistant) *ActivityManager {
	am := &ActivityManager{
		a:                a,
		systemActivityId: make(map[int32]*ActivityConfig),
		activitiesDetail: make(map[int32]*pb.ActivityCommonData),
	}

	a.Client().OnReqGetActivityDetail(am.onRspGetActivityDetail)
	a.Client().OnReqGetActivityConditionReward21033(am.onRspGetActivityConditionReward)
	a.Client().OnPushActivityList(am.onPushActivityList)
	a.Client().OnActivityCommonDataListSync(am.onActivityCommonDataListSync)
	a.Client().OnActivityConditionDataListSync(am.onActivityConditionDataListSync)
	a.Client().OnReqGetActivityConditionReward21004(am.onReqGetActivityConditionReward21004)

	return am
}

func (am *ActivityManager) onActivityCommonDataListSync(ctx client.Context, msg *pb.ActivityCommonDataListSync) error {
	for _, activity := range msg.GetActivityDataList() {
		for _, cond := range activity.GetConditionDataList() {
			if !cond.GetIsGetReward() && cond.GetCompleteTime() > 0 {
				am.getActivityReward(ctx, activity.GetActivityId(), cond.GetConditionId())
			}
		}
		detail := activity.GetDetailConfig()
		if len(detail.GetCommonConfig().GetMainConfig().GetServerId()) > 1 {
			detail.CommonConfig.MainConfig.ServerId = detail.CommonConfig.MainConfig.ServerId[:1]
		}
		am.activitiesDetail[activity.GetActivityId()] = activity
	}

	return nil
}

func (am *ActivityManager) GetActivityIdByName(name string) int32 {
	am.mu.RLock()
	defer am.mu.RUnlock()

	for _, activity := range am.systemActivityId {
		if activity.Name == name {
			return activity.Id
		}
	}

	return 0
}

func (am *ActivityManager) ListActivity() []*ActivityConfig {
	am.mu.RLock()
	defer am.mu.RUnlock()

	list := make([]*ActivityConfig, 0, len(am.systemActivityId))
	for _, activity := range am.systemActivityId {
		list = append(list, activity)
	}
	return list
}

func (am *ActivityManager) ListActivityDetail() map[int32]*pb.ActivityCommonData {
	am.mu.RLock()
	defer am.mu.RUnlock()

	return am.activitiesDetail
}

func (am *ActivityManager) onPushActivityList(ctx client.Context, msg *pb.PushActivityList) error {
	am.a.L().Info("同步活动推送数据")

	am.mu.Lock()
	defer am.mu.Unlock()

	now := time.Now().UnixMilli()

	for _, activity := range msg.MainConfig {
		if activity.GetEndTime() < now {
			continue
		}

		if c, ok := am.systemActivityId[activity.GetType()]; !ok || activity.GetEndTime() > c.EndTime.UnixMilli() {
			am.systemActivityId[activity.GetType()] = &ActivityConfig{
				ActivityMainConfig: activity,
				Id:                 activity.GetActivityId(),
				Type:               activity.GetType(),
				GroupType:          activity.GetGroupType(),
				Name:               db.GetSystemInfo(int(activity.GetType())).GetName(),
				BeginTime:          time.UnixMilli(activity.GetBeginTime()),
				EndTime:            time.UnixMilli(activity.GetEndTime()),
			}
		}

		// am.getActivityDetail(ctx, activity.GetActivityId())
	}

	return nil
}

// ReqGetActivityDetail
func (am *ActivityManager) getActivityDetail(ctx context.Context, activityId int32) error {
	return am.a.Client().SendReqGetActivityDetail(ctx, &pb.ReqGetActivityDetail{
		ActivityId: &activityId,
	})
}

// onRspGetActivityDetail
func (am *ActivityManager) onRspGetActivityDetail(ctx client.Context, msg *pb.RspGetActivityDetail) error {
	if msg.GetRet() != 0 {
		return db.GameError(msg)
	}

	am.mu.Lock()
	defer am.mu.Unlock()

	activity := msg.GetActivity()
	detail := activity.GetDetailConfig()
	if len(detail.GetCommonConfig().GetMainConfig().GetServerId()) > 1 {
		detail.CommonConfig.MainConfig.ServerId = detail.CommonConfig.MainConfig.ServerId[:1]
	}

	for _, condition := range activity.GetConditionDataList() {
		if !condition.GetIsGetReward() && condition.GetCompleteTime() > 0 {
			am.getActivityReward(ctx, activity.GetActivityId(), condition.GetConditionId())
		}
	}

	am.activitiesDetail[activity.GetActivityId()] = activity

	am.buyMallGoods(ctx, activity)

	return nil
}

func (am *ActivityManager) getActivityReward(ctx context.Context, activityId int32, conditionId int32) error {
	return am.a.Client().SendReqGetActivityConditionReward21033(ctx, &pb.ReqGetActivityConditionReward{
		ActivityId:  &activityId,
		ConditionId: &conditionId,
	})
}

// onRspGetActivityConditionReward
func (am *ActivityManager) onRspGetActivityConditionReward(ctx client.Context, msg *pb.RspGetActivityConditionReward) error {
	if msg.GetRet() != 0 {
		return db.GameError(msg)
	}

	am.a.L().Info("领取活动条件奖励", zap.Stringer("rewards", db.ParseRewards(msg.GetRewards())))

	return nil
}

// onActivityConditionDataListSync
func (am *ActivityManager) onActivityConditionDataListSync(ctx client.Context, msg *pb.ActivityConditionDataListSync) error {
	// for _, activity := range msg.GetActivityConditionDataList() {
	// 	for _, condition := range activity.GetConditionDataList() {
	// 		if condition.GetIsGetReward() {
	// 			am.a.L().Info("有可领取物品", zap.Any("condition", condition))
	// 		}
	// 	}
	// }
	return nil
}

// onReqGetActivityConditionReward21004
func (am *ActivityManager) onReqGetActivityConditionReward21004(ctx client.Context, msg *pb.RspGetActivityConditionReward) error {
	return db.GameError(msg)
}

func (am *ActivityManager) buyMallGoods(ctx context.Context, activities ...*pb.ActivityCommonData) error {
	for _, activity := range activities {
		mallConfig := activity.GetDetailConfig().GetCommonConfig().GetMallConfig()
		for _, mall := range mallConfig {
			if price := mall.GetMallTempMsg().GetPrice(); price != "100000=0" {
				continue
			}

			id := mall.GetMallTempMsg().GetId()
			buyLimit := mall.GetMallTempMsg().GetBuyLimit()
			name := mall.GetMallTempMsg().GetName()
			if id == 9655276 || id == 9712892 || id == 9655196 {
				continue
			}

			for i := 0; i < int(buyLimit); i++ {
				am.a.reward.PushDelayFunc(func() {
					am.a.Storage().Cache().Condition(ctx, fmt.Sprintf("%d_%d", activity.GetActivityId(), id), func() error {
						am.a.L().Info(
							"购买商城物品",
							zap.Int32("activity_id", activity.GetActivityId()),
							zap.Int32("mall_id", id),
							zap.String("name", name),
							zap.Int("times", i+1),
						)

						return am.a.Client().SendReqBuyActivityGoods(ctx, &pb.ReqBuyActivityGoods{
							ActivityId: utils.Ptr(activity.GetActivityId()),
							Count:      utils.Ptr(int64(1)),
							MallId:     utils.Ptr(id),
						})
					})
				})
			}

		}
	}

	return nil
}
