package assist

import (
	"context"
	"errors"
	"time"

	"xddq/pkg/game/client"
	"xddq/pkg/game/db"
	"xddq/pkg/game/protocol/pb"
	"xddq/pkg/utils"

	"github.com/samber/lo"
	"go.uber.org/zap"
)

type UnionManager struct {
	a *Assistant

	unionInfo      *pb.MyUnionData
	bountyjobId    int
	daliytaskjobId int
	groupType      int32
}

func NewUnionManager(a *Assistant) *UnionManager {
	u := &UnionManager{
		a: a,
	}

	a.addRunner(u)

	// 妖盟数据
	a.Client().OnMyUnionData(u.onMyUnionData)

	// 砍价
	a.Client().OnCutPriceDataMsg(u.onCutPriceDataMsg)

	// 妖盟悬赏
	a.Client().OnUnionBountyEnterMapReq(u.onEnterUnionBountyMapResp)
	a.Client().OnUnionBountyOpenBountyEventReq(u.onUnionBountyOpenBountyEventResp)
	a.Client().OnUnionBountyOpenMonsterReq(u.onUnionBountyOpenMonsterResp)

	// 妖盟寻宝
	a.Client().OnUnionTreasureEnterReq(u.onUnionTreasureEnterResp)

	// 妖盟日常任务
	a.Client().OnReqUnionDailyTask(u.onRspUnionDailyTask)

	// 妖盟对战
	a.Client().OnSynPlayerInfo(u.onSynPlayerInfo)

	// 妖盟散财
	a.Client().OnUnionBlessingGiftSyncList(u.onUnionBlessingGiftSyncList)
	a.Client().OnUnionBlessingRewardReqMsg(u.onUnionBlessingRewardResp)

	return u
}

func (u *UnionManager) Run(ctx context.Context) error {
	if u.a.Config().Union.Bounty {
		u.enterUnionBountyMap(ctx)
		u.bountyjobId = u.a.addJob("unionBounty", "@every 5m", u.enterUnionBountyMap)
	}

	if u.a.Config().Union.UnionTask {
		u.check(ctx)
		u.daliytaskjobId = u.a.addJob("unionTask", "@every 5m", u.check)
	}

	return nil
}

func (u *UnionManager) Stop() error {
	u.a.removeJob(u.bountyjobId)
	u.a.removeJob(u.daliytaskjobId)
	return nil
}

func (u *UnionManager) Members() []*pb.UnionMemberData {
	if u.unionInfo == nil {
		return nil
	}

	return u.unionInfo.MemberList
}

// check 妖盟日常任务
func (u *UnionManager) check(ctx context.Context) error {

	u.a.Client().SendReqUnionDailyTask(ctx, &pb.ReqUnionDailyTask{})

	return nil
}

// onMyUnionData 妖盟数据
func (u *UnionManager) onMyUnionData(ctx client.Context, msg *pb.MyUnionData) error {
	u.a.L().Info("同步妖盟数据")

	u.unionInfo = msg
	return nil
}

// checkAndBuy 检查是否可以购买
func (u *UnionManager) checkAndBuy(ctx client.Context, msg *pb.CutPriceDataMsg) error {
	// 状态为 2 表示已经购买
	if msg.GetStatus() == 2 {
		return nil
	}

	// 2888 是初始价格,
	price := 2888 - msg.GetBargainPrice()
	if price > 0 {
		return nil
	}

	// 距离消失时间小于 5 分钟, 并且价格小于 0 则进行购买
	if time.Until(time.Unix(msg.GetDisappearTime(), 0)) < 5*time.Minute ||
		u.unionInfo != nil && int(msg.GetBargainTimes()) >= len(u.unionInfo.GetMemberList()) {
		u.a.L().Info("购买砍价", zap.Any("price", 2888-msg.GetBargainPrice()))

		if err := ctx.Send(ctx, &pb.CutPriceBuyReqMsg{
			BussinessId: utils.Ptr(msg.GetBussinessId()),
		}); err != nil {
			return err
		}
	}

	return nil
}

// onCutPriceDataMsg 砍价
func (u *UnionManager) onCutPriceDataMsg(ctx client.Context, msg *pb.CutPriceDataMsg) error {
	if !u.a.Config().Union.CutPrice {
		return nil
	}

	if msg.GetStatus() == 0 {
		if err := ctx.Send(ctx, &pb.CutPriceBargainReqMsg{
			BussinessId: utils.Ptr(msg.GetBussinessId()),
		}); err != nil {
			return err
		}
	}

	return u.checkAndBuy(ctx, msg)
}

// onRspUnionDailyTask 妖盟日常任务
func (u *UnionManager) onRspUnionDailyTask(ctx client.Context, msg *pb.RspUnionDailyTask) error {
	u.a.Storage().Cache().Condition(ctx, "union_daliy_task_done", func() error {

		progress := msg.GetProgress()
		tasks := []lo.Tuple2[int, int]{
			lo.T2(150, 0),
			lo.T2(250, 1),
			lo.T2(500, 2),
			lo.T2(750, 3),
			lo.T2(1000, 4),
		}

		done := 0

		for _, task := range tasks {
			if p, i := task.Unpack(); progress >= int32(p) {
				u.a.Client().SendReqUnionGetDailyTask(ctx, &pb.ReqUnionGetDailyTask{
					ActIndex: utils.Ptr(int32(i)),
				})
				done++
			}
		}

		if done >= len(tasks) {
			u.a.removeJob(u.daliytaskjobId)
			return nil
		}

		return errors.New("任务未完成")
	})
	return nil
}

// onSynPlayerInfo 妖盟对战
func (u *UnionManager) onSynPlayerInfo(ctx client.Context, msg *pb.SynPlayerInfo) error {
	u.a.L().Info("OnSynPlayerInfo", zap.Any("msg", msg))
	return nil
}

func (u *UnionManager) detach() {
	u.a.removeJob(u.bountyjobId)
}

func (u *UnionManager) enterUnionBountyMap(ctx context.Context) error {
	if hours := time.Now().Hour(); hours > 0 && hours < 8 {
		return nil
	}

	u.a.Client().SendUnionBountyEnterMapReq(ctx, &pb.UnionBountyEnterMapReq{})

	return nil
}

func (u *UnionManager) onEnterUnionBountyMapResp(ctx client.Context, msg *pb.UnionBountyEnterMapResp) error {
	u.groupType = msg.GetGroupType()

	playerData := msg.GetPlayerData()

	u.handleRob(ctx, msg)

	u.handleMonster(ctx, msg)

	// 领取悬赏奖励
	if msg.GetSettlementRedPoint() {
		u.a.Client().SendUnionBountyGetRewardEscortReq(ctx, &pb.UnionBountyGetRewardEscortReq{})
	}

	// 排行榜膜拜
	if msg.GetWorshipRedPoint() {
		u.a.Client().SendUnionBountyWorshipReq(ctx, &pb.UnionBountyWorshipReq{})
	}

	// 打开悬赏界面
	if playerData.GetBountyTimes() < 2 && msg.GetMyCart() == nil {
		u.a.Client().SendUnionBountyOpenBountyEventReq(ctx, &pb.UnionBountyOpenBountyEventReq{})
	}

	if playerData.GetBountyTimes() >= 2 && // 击杀次数
		playerData.GetRobTimes() >= 2 && // 抢夺次数
		playerData.GetHelpTimes() >= 1 && // 帮助次数
		msg.GetMyCart() == nil {
		u.detach()
	}

	return nil
}

func (u *UnionManager) handleRob(ctx client.Context, msg *pb.UnionBountyEnterMapResp) {
	if msg.GetPlayerData().GetRobTimes() >= 2 {
		return
	}

	robTimes := 0

	for _, cart := range msg.GetCartList() {
		if cart.GetAttackedByMe() ||
			cart.GetConfigId() < 1010 ||
			float64(cart.GetPlayerData().GetFightValue())*1.2 > float64(u.a.Role().FightValue()) {
			continue
		}

		u.a.Client().SendUnionBountyPlunderReq(ctx, &pb.UnionBountyPlunderReq{
			TargetId:  utils.Ptr(cart.GetPlayerData().GetPlayerId()),
			StartTime: utils.Ptr(cart.GetStartTime()),
			Energy:    utils.Ptr(int64(cart.GetEnergy())),
		})

		robTimes++
	}

	if robTimes == 0 && msg.GetPlayerData().GetRefreshTimes() < 3 {
		u.a.Client().SendUnionBountyRefreshMapReq(ctx, &pb.UnionBountyRefreshMapReq{})
	}
}

func (u *UnionManager) handleMonster(ctx client.Context, msg *pb.UnionBountyEnterMapResp) {
	playerData := msg.GetPlayerData()
	for _, monster := range msg.GetMonsterList() {
		if monster.GetPlayerId() == ctx.PlayerId() {
			u.a.Client().SendUnionBountyOpenMonsterReq(ctx, &pb.UnionBountyOpenMonsterReq{
				PlayerId: monster.PlayerId,
			})
		}

		if playerData.GetHelpTimes() < 1 {
			u.a.Client().SendUnionBountyDealAskHelpReq(ctx, &pb.UnionBountyDealAskHelpReq{
				PlayerId: monster.PlayerId,
			})
		}
	}
}

func (u *UnionManager) onUnionBountyOpenBountyEventResp(ctx client.Context, msg *pb.UnionBountyOpenBountyEventResp) error {
	if msg.GetRet() != 0 {
		u.a.L().Error("OpenBountyEventResp", zap.Error(db.GetGameError(msg.GetRet())))
		return nil
	}

	if msg.BountyInfo == nil {
		u.a.L().Info("Bounty End")
		return nil
	}

	u.a.Client().SendUnionBountyDealBountyReq(ctx, &pb.UnionBountyDealBountyReq{
		CurConfigId: msg.BountyInfo.CurConfigId,
		Type:        utils.Ptr(u.groupType),
	})

	return nil
}

func (u *UnionManager) onUnionBountyOpenMonsterResp(ctx client.Context, msg *pb.UnionBountyOpenMonsterResp) error {
	if msg.GetRet() != 0 {
		return db.GameError(msg)
	}

	if len(msg.GetBattleMainList()) >= 3 {
		ctx.Send(ctx, &pb.UnionBountyAttackMonsterReq{
			TeamPlayerList: lo.Map(msg.GetBattleMainList(), func(item *pb.BattleMain, _ int) *pb.UnionBountyMonsterTeamPlayerMsg {
				return &pb.UnionBountyMonsterTeamPlayerMsg{
					PlayerId: item.PlayerId,
					Pos:      item.Pos,
				}
			}),
			BossId: msg.Event,
		})
	}

	return nil
}

func (u *UnionManager) onUnionBountyAttackMonsterResp(ctx client.Context, msg *pb.UnionBountyAttackMonsterResp) error {
	if msg.GetRet() != 0 {
		return db.GameError(msg)
	}

	return u.enterUnionBountyMap(ctx)
}

// onUnionTreasureEnterResp 妖盟寻宝
func (u *UnionManager) onUnionTreasureEnterResp(ctx client.Context, msg *pb.UnionTreasureEnterResp) error {
	return nil
}

// onUnionBlessingGiftSyncList 同步妖盟散财数据
func (u *UnionManager) onUnionBlessingGiftSyncList(ctx client.Context, msg *pb.UnionBlessingGiftSyncList) error {
	for _, data := range msg.GetData() {
		u.a.L().Info("收到妖盟散财数据", zap.String("unionName", data.GetUnionName()), zap.String("id", data.GetId()))
		ctx.Client().SendUnionBlessingRewardReqMsg(ctx, &pb.UnionBlessingRewardReqMsg{
			Id: data.Id,
		})
	}
	return nil
}

func (u *UnionManager) onUnionBlessingRewardResp(_ client.Context, msg *pb.UnionBlessingRewardRespMsg) error {
	u.a.L().Info("收到妖盟散财奖励", zap.String("id", msg.GetId()), zap.Any("reward", db.ParseRewards(msg.GetReward())))
	return nil
}
