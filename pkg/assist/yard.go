package assist

import (
	"context"
	"strconv"
	"time"

	"xddq/pkg/game/client"
	"xddq/pkg/game/db"
	"xddq/pkg/game/protocol/pb"
	"xddq/pkg/utils"

	"go.uber.org/zap"
	"golang.org/x/exp/maps"
)

type YardManager struct {
	a  *Assistant
	mc chan *pb.YardMerchantInfoResp

	memberMerchantInfos map[int32]*MerchantInfo

	danNum      int64
	grassNum    int64
	buildingMsg []*pb.YardBuildingMsg
}

func NewYardManager(a *Assistant) *YardManager {
	y := &YardManager{
		a:                   a,
		mc:                  make(chan *pb.YardMerchantInfoResp),
		memberMerchantInfos: make(map[int32]*MerchantInfo),
	}

	if a.Config().Yard {
		a.addRunner(y)

		a.Client().OnYardMerchantInfoReq(y.onYardMerchantInfo)
		a.Client().OnYardEnterReq(y.onYardEnterResp)
		a.Client().OnYardMakeMsgSync(y.onYardMakeMsgSync)
		a.Client().OnYardLoginSync(y.onYardLoginSync)
		a.Client().OnYardBuildGainRewardReq(y.onYardBuildGainRewardResp)
		a.Client().OnYardBuildMakeReq(y.onYardBuildMakeResp)
		a.Client().OnYardBuildInfoMsgSync(y.onYardBuildInfoMsgSync)
	}

	return y
}

type MerchantInfo struct {
	PlayerId int64  `json:"playerId"`
	NickName string `json:"nickname"`

	Name   string `json:"name"`
	CorpId int32  `json:"corpId"`
	Price  int64  `json:"price"`
}

func (y *YardManager) Run(ctx context.Context) error {
	if !y.a.Config().Yard {
		return nil
	}

	y.a.L().Info("yard manager running")

	y.check(ctx)
	y.a.addJob("yard", "@every 10m", y.check)

	if err := y.makeMerchantInfo(ctx); err != nil {
		return err
	}

	return nil
}

func (y *YardManager) Stop() error {
	y.a.L().Info("yard manager stop")

	return nil
}

func (y *YardManager) makeMerchantInfo(ctx context.Context) error {
	if err := y.a.Storage().Cache().Get("yard_merchant_infos", &y.memberMerchantInfos); err == nil {
		return nil
	}

	for _, member := range y.a.Union().Members() {
		p := member.PlayerData
		m, err := y.GetMerchantInfo(ctx, p.GetPlayerId())
		if err != nil {
			y.a.L().Error("get merchant info", zap.Error(err))
			continue
		}

		for _, item := range m.MerchantItemMsg {
			y.memberMerchantInfos[item.GetCropId()] = &MerchantInfo{
				PlayerId: p.GetPlayerId(),
				NickName: p.GetNickName(),
				Name:     db.GetItemName(item.GetCropId()),
				CorpId:   item.GetCropId(),
				Price:    item.GetPrice(),
			}
		}
	}

	y.a.Storage().Cache().Set(ctx, "yard_merchant_infos", y.memberMerchantInfos)

	return nil
}

// onYardLoginSync
func (y *YardManager) onYardLoginSync(ctx client.Context, msg *pb.YardLoginSync) error {
	y.grassNum, _ = strconv.ParseInt(msg.GetGrassNum(), 10, 64)
	return nil
}

func (y *YardManager) GetMerchantInfoByName(ctx context.Context, name string) *MerchantInfo {
	for _, m := range y.memberMerchantInfos {
		if m.Name == name {
			return m
		}
	}
	return nil
}

func (y *YardManager) ListMerchantInfo() []*MerchantInfo {
	return maps.Values(y.memberMerchantInfos)
}

func (y *YardManager) GetMerchantInfo(ctx context.Context, id int64) (msg *pb.YardMerchantInfoResp, err error) {
	y.a.Client().SendYardMerchantInfoReq(ctx, &pb.YardMerchantInfoReq{
		TargetPlayerId: &id,
	})

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case msg := <-y.mc:
		return msg, nil
	}
}

func (y *YardManager) onYardMerchantInfo(ctx client.Context, msg *pb.YardMerchantInfoResp) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case y.mc <- msg:
	default:
	}

	return nil
}

func (y *YardManager) check(ctx context.Context) error {
	return y.enter(ctx, y.a.PlayerInfo().PlayerId())
}

func (y *YardManager) enter(ctx context.Context, id int64) error {
	return y.a.Client().SendYardEnterReq(ctx, &pb.YardEnterReq{
		TargetPlayerId: &id,
	})
}

var actionName = map[int32]string{
	1001: "采药",
	1002: "收丹药",
	1003: "收桃子",
	1004: "收孕育",
}

func (y *YardManager) processBuilding(ctx context.Context, building *pb.YardBuildingMsg) {
	cell := building.GetYardCellMsg()
	detail := building.GetYardBuildDetailMsg()
	uniqueId := cell.GetUniqueId()
	buildId := cell.GetBuildId()

	if actionName[cell.GetBuildId()] != "" &&
		(time.Since(time.UnixMilli(detail.GetStartTime())) > 10*time.Minute || detail.GetStatus() == 0) {
		y.a.L().Info(
			"收菜",
			zap.String("action", actionName[buildId]),
			zap.Int64("uniqueId", uniqueId),
		)
		y.gainReward(ctx, uniqueId, buildId)
	}

	switch cell.GetBuildId() {
	case 1002:
		num := y.grassNum / 500
		if detail.GetStatus() == 0 && num > 0 {
			y.a.L().Info(
				"炼丹",
				zap.Int64("uniqueId", uniqueId),
				zap.Int32("buildId", buildId),
			)

			y.a.Client().SendYardBuildMakeReq(ctx, &pb.YardBuildMakeReq{
				UniqueId: &uniqueId,
				BuildId:  &buildId,
				Num:      utils.Ptr(int32(num)),
				IsCancel: utils.Ptr(false),
			})
		}

	case 1004:
		if detail.GetStatus() == 0 {
			y.a.L().Info("孕育", zap.Int64("uniqueId", cell.GetUniqueId()))

			y.a.Client().SendYardBuildMakeReq(ctx, &pb.YardBuildMakeReq{
				UniqueId:  &uniqueId,
				BuildId:   &buildId,
				ProductId: utils.Ptr(int32(400000 + detail.GetLevel())),
				Num:       utils.Ptr(int32(100)),
				IsCancel:  utils.Ptr(false),
			})
		}
	}
}

// onYardEnterResp
func (y *YardManager) onYardEnterResp(ctx client.Context, msg *pb.YardEnterResp) error {
	drawData := msg.GetDrawData()

	for i := drawData.GetAdCount(); i < 2; i++ {
		y.draw(ctx, true, 0)
	}

	for i := drawData.GetFreeDrawTimes(); i < 1; i++ {
		y.draw(ctx, false, 0)
	}

	areaInfo := msg.GetAreaInfo()
	for _, area := range areaInfo {
		for _, building := range area.GetBuildingMsg() {
			y.processBuilding(ctx, building)
		}
	}

	return nil
}

// gainReward
func (y *YardManager) gainReward(ctx context.Context, uniqueId int64, buildId int32) {
	y.a.Client().SendYardBuildGainRewardReq(ctx, &pb.YardBuildGainRewardReq{
		UniqueId: &uniqueId,
		BuildId:  &buildId,
	})
}

func (y *YardManager) onYardBuildMakeResp(ctx client.Context, msg *pb.YardBuildMakeResp) error {
	return db.GameError(msg)
}

func (y *YardManager) onYardBuildGainRewardResp(ctx client.Context, msg *pb.YardBuildGainRewardResp) error {
	if msg.GetRet() != 0 {
		return nil
	}

	y.a.L().Info("收菜成功", zap.Any("reward", db.ParseRewards(msg.GetRewards())))

	return nil
}

// draw
func (y *YardManager) draw(ctx context.Context, isAd bool, poolId int32) {
	y.a.Client().SendYardDrawReq(ctx, &pb.YardDrawReq{
		IsTen:       utils.Ptr(false),
		IsUseADTime: utils.Ptr(false),
		IsADType:    &isAd,
		PoolId:      &poolId,
	})
}

// onYardBuildInfoMsgSync
func (y *YardManager) onYardBuildInfoMsgSync(ctx client.Context, msg *pb.YardBuildInfoMsgSync) error {
	return nil
}

// onYardMakeMsgSync
func (y *YardManager) onYardMakeMsgSync(ctx client.Context, msg *pb.YardMakeMsgSync) error {
	y.grassNum, _ = strconv.ParseInt(msg.GetGrassNum(), 10, 64)
	y.danNum = msg.GetDanNum()

	return nil
}
