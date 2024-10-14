package client

import (
	"context"

	"xddq/pkg/game/protocol/pb"

	"google.golang.org/protobuf/proto"
)


// 登陆消息
func (c *Client) SendReqLoginMsg(ctx context.Context, req *pb.ReqLoginMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20001, data)
}

// 心跳
func (c *Client) SendReqPingMsg(ctx context.Context, req *pb.ReqPingMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20003, data)
}

// 预购买
func (c *Client) SendPreChargeReq(ctx context.Context, req *pb.PreChargeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20004, data)
}

// 创建角色请求
func (c *Client) SendCreateRoleReq(ctx context.Context, req *pb.CreateRoleReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20006, data)
}

// 改名
func (c *Client) SendChangeNickNameReq(ctx context.Context, req *pb.ChangeNickNameReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20008, data)
}

// 兑换码
func (c *Client) SendUseExchangeKeyReq(ctx context.Context, req *pb.UseExchangeKeyReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20009, data)
}

// 观看视频广告任务
func (c *Client) SendWatchAdReq(ctx context.Context, req *pb.WatchAdReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20019, data)
}

// 熊口逃生领奖
func (c *Client) SendCutRopeGetRewardReq(ctx context.Context, req *pb.CutRopeGetRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20020, data)
}

// 获取隐私注册协议
func (c *Client) SendGetPolicyInfoReq(ctx context.Context, req *pb.GetPolicyInfoReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20031, data)
}

// 获取绑定状态
func (c *Client) SendGetBindStatusReq(ctx context.Context, req *pb.GetBindStatusReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20033, data)
}

// 生成绑定码
func (c *Client) SendGetBindCodeReq(ctx context.Context, req *pb.GetBindCodeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20034, data)
}

// 解绑
func (c *Client) SendUntieBindCodeReq(ctx context.Context, req *pb.UntieBindCodeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20035, data)
}

// 获取绑定信息
func (c *Client) SendGetBindCodeInfoReq(ctx context.Context, req *pb.GetBindCodeInfoReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20036, data)
}

// 注销账号
func (c *Client) SendDeletePlayerReq(ctx context.Context, req *pb.DeletePlayerReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20037, data)
}

// 获取系统分组信息
func (c *Client) SendGetSystemGroupInfoReq(ctx context.Context, req *pb.GetSystemGroupInfoReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20040, data)
}

// 请求领取系统解锁奖励
func (c *Client) SendGetSysRewardReq(ctx context.Context, req *pb.GetSysRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20103, data)
}

// 领取特权卡奖励
func (c *Client) SendPrivilegeCardReceiveRewardReq(ctx context.Context, req *pb.PrivilegeCardReceiveRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20105, data)
}

// 切换形象请求
func (c *Client) SendChangePlayerCharaReq(ctx context.Context, req *pb.ChangePlayerCharaReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20107, data)
}

// 获取玩家详细数据
func (c *Client) SendGetPlayerDetailDataReq(ctx context.Context, req *pb.GetPlayerDetailDataReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20109, data)
}

// 请求聊天
func (c *Client) SendWorldChatReqMsg(ctx context.Context, req *pb.WorldChatReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20110, data)
}

// 请求活动消息列表
func (c *Client) SendActivityChatReqMsg(ctx context.Context, req *pb.ActivityChatReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20116, data)
}

// 获取仙树后面飘的玩家信息
func (c *Client) SendGetTopPlayerInfoReq(ctx context.Context, req *pb.GetTopPlayerInfoReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20118, data)
}

// 穿戴称号
func (c *Client) SendUseTitleReq(ctx context.Context, req *pb.UseTitleReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20120, data)
}

// 穿戴称号
func (c *Client) SendUseMedalReq(ctx context.Context, req *pb.UseMedalReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20123, data)
}

// 举报玩家
func (c *Client) SendWorldChatBlockReqMsg(ctx context.Context, req *pb.WorldChatBlockReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20130, data)
}

// 系统聊天列表，请求下发
func (c *Client) SendSystemChatReqMsg(ctx context.Context, req *pb.SystemChatReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20132, data)
}

// 聊天举报
func (c *Client) SendReportMessageReqMsg(ctx context.Context, req *pb.ReportMessageReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20133, data)
}

// 获取玩家(某个系统)锁定的详细数据-请求
func (c *Client) SendGetLockedPlayerDetailDataReq(ctx context.Context, req *pb.GetLockedPlayerDetailDataReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20134, data)
}

// 请求飞行员列表数据-请求
func (c *Client) SendGetPlayerPilotShowDataMsgListReq(ctx context.Context, req *pb.GetPlayerPilotShowDataMsgListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20135, data)
}

// 通用搜索玩家
func (c *Client) SendCommonSearchPlayerReq(ctx context.Context, req *pb.CommonSearchPlayerReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20136, data)
}

// 获取红包信息
func (c *Client) SendRedPacketOpenReq(ctx context.Context, req *pb.RedPacketOpenReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20141, data)
}

// 领取红包奖励
func (c *Client) SendRedPacketDrawReq(ctx context.Context, req *pb.RedPacketDrawReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20142, data)
}

// 皮肤羁绊任务领取
func (c *Client) SendCollectPlayerCharaReq(ctx context.Context, req *pb.CollectPlayerCharaReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20150, data)
}

// 皮肤等级升级
func (c *Client) SendUpgradePlayerCharaReq(ctx context.Context, req *pb.UpgradePlayerCharaReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20151, data)
}

// 处理装备
func (c *Client) SendReqDealEquipmentMsg(ctx context.Context, req *pb.ReqDealEquipmentMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20202, data)
}

// 做梦
func (c *Client) SendReqDreamMsg(ctx context.Context, req *pb.ReqDreamMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20203, data)
}

// 境界升级
func (c *Client) SendReqRealmsLeveUpMsg(ctx context.Context, req *pb.ReqRealmsLeveUpMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20204, data)
}

// 梦境升级
func (c *Client) SendDreamLvUpReq(ctx context.Context, req *pb.DreamLvUpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20205, data)
}

// 梦境升级加速
func (c *Client) SendDreamLvUpSpeedUpReq(ctx context.Context, req *pb.DreamLvUpSpeedUpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20206, data)
}

// 获取渡劫成功率
func (c *Client) SendGetTribulationSuccessProReq(ctx context.Context, req *pb.GetTribulationSuccessProReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20208, data)
}

// 获取未处理装备数据
func (c *Client) SendGetUnDealEquipmentMsgReq(ctx context.Context, req *pb.GetUnDealEquipmentMsgReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20209, data)
}

// 广告领奖
func (c *Client) SendGetAdRewardReq(ctx context.Context, req *pb.GetAdRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20211, data)
}

// 斩道也调用这个接口
func (c *Client) SendSoaringReq(ctx context.Context, req *pb.SoaringReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20212, data)
}

// 设置分身名字
func (c *Client) SendSetSeparationNameReq(ctx context.Context, req *pb.SetSeparationNameReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20213, data)
}

// 切换分身
func (c *Client) SendSwitchSeparationReq(ctx context.Context, req *pb.SwitchSeparationReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20214, data)
}

// 获取分身列表数据
func (c *Client) SendGetSeparationDataMsgListReq(ctx context.Context, req *pb.GetSeparationDataMsgListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20215, data)
}

// 设置订阅消息
func (c *Client) SendSetMessageSubscribeDataReqMsg(ctx context.Context, req *pb.SetMessageSubscribeDataReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20251, data)
}

// 微信头像授权
func (c *Client) SendAuthorizePlayerHeadReq(ctx context.Context, req *pb.AuthorizePlayerHeadReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20252, data)
}

// 使用道具
func (c *Client) SendUsePropReq(ctx context.Context, req *pb.UsePropReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20302, data)
}

// 合成道具
func (c *Client) SendCompoundPropReq(ctx context.Context, req *pb.CompoundPropReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20303, data)
}

// 挑战请求
func (c *Client) SendRankBattleChallengeReq(ctx context.Context, req *pb.RankBattleChallengeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20412, data)
}

// 妖榜复仇请求
func (c *Client) SendRankBattleRevengeReq(ctx context.Context, req *pb.RankBattleRevengeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20414, data)
}

// 获取跨服斗法参与区服
func (c *Client) SendRankBattleServerListReq(ctx context.Context, req *pb.RankBattleServerListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20418, data)
}

// 领取任务奖励响应
func (c *Client) SendTaskGetRewardReqMsg(ctx context.Context, req *pb.TaskGetRewardReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20503, data)
}

// 一键删除邮件
func (c *Client) SendMailDeleteAllReq(ctx context.Context, req *pb.MailDeleteAllReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20553, data)
}

// 删除邮件
func (c *Client) SendMailDeleteReq(ctx context.Context, req *pb.MailDeleteReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20554, data)
}

// 一键领取邮件
func (c *Client) SendMailGetAllRewardReq(ctx context.Context, req *pb.MailGetAllRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20555, data)
}

// 领取邮件
func (c *Client) SendMailGetRewardReq(ctx context.Context, req *pb.MailGetRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20556, data)
}

// 读取邮件
func (c *Client) SendMailReadReq(ctx context.Context, req *pb.MailReadReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20557, data)
}

// 购买商品
func (c *Client) SendBuyGoodsReq(ctx context.Context, req *pb.BuyGoodsReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20601, data)
}

// 商店列表
func (c *Client) SendMallRandomGoodsReqMsg(ctx context.Context, req *pb.MallRandomGoodsReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20603, data)
}

// 通宝购买请求
func (c *Client) SendMallChanceUseReqMsg(ctx context.Context, req *pb.MallChanceUseReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20604, data)
}

// 炼化
func (c *Client) SendRandomTalentReq(ctx context.Context, req *pb.RandomTalentReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20622, data)
}

// 处理神通
func (c *Client) SendDealTalentReq(ctx context.Context, req *pb.DealTalentReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20623, data)
}

// 阅读妖书
func (c *Client) SendReadBookReq(ctx context.Context, req *pb.ReadBookReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20624, data)
}

// 获取未处理神通数据
func (c *Client) SendGetUnDealTalentMsgReq(ctx context.Context, req *pb.GetUnDealTalentMsgReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20625, data)
}

// 赠礼
func (c *Client) SendReqDestinyGift(ctx context.Context, req *pb.ReqDestinyGift) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20652, data)
}

// 游历
func (c *Client) SendReqDestinyTravel(ctx context.Context, req *pb.ReqDestinyTravel) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20653, data)
}

// 切磋
func (c *Client) SendReqDestinyChallenge(ctx context.Context, req *pb.ReqDestinyChallenge) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20654, data)
}

// 解锁道具仙友
func (c *Client) SendReqUnlockDestinyByItem(ctx context.Context, req *pb.ReqUnlockDestinyByItem) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20655, data)
}

// 皮肤解锁
func (c *Client) SendReqUnlockDestinySkinByItem(ctx context.Context, req *pb.ReqUnlockDestinySkinByItem) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20657, data)
}

// 强化皮肤
func (c *Client) SendReqEnhanceDestinySkin(ctx context.Context, req *pb.ReqEnhanceDestinySkin) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20658, data)
}

// 穿戴皮肤
func (c *Client) SendReqWearDestinySkin(ctx context.Context, req *pb.ReqWearDestinySkin) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20659, data)
}

// 切换挚友联动形象
func (c *Client) SendDestinySwitchLinkageSkinReq(ctx context.Context, req *pb.DestinySwitchLinkageSkinReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20660, data)
}

// 解锁筋斗云请求
func (c *Client) SendUnlockCloudReq(ctx context.Context, req *pb.UnlockCloudReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20712, data)
}

// 筋斗云升级请求
func (c *Client) SendCloudLvUpReq(ctx context.Context, req *pb.CloudLvUpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20713, data)
}

// 筋斗云升阶请求
func (c *Client) SendCloudStageUpReq(ctx context.Context, req *pb.CloudStageUpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20714, data)
}

// 筋斗云穿戴请求
func (c *Client) SendCloudEquipReq(ctx context.Context, req *pb.CloudEquipReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20715, data)
}

// 筋斗云皮肤穿戴请求
func (c *Client) SendCloudEquipSkinReq(ctx context.Context, req *pb.CloudEquipSkinReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20716, data)
}

// 筋斗云皮肤升级 0-&gt;1 为激活
func (c *Client) SendCloudSkinLvUpReq(ctx context.Context, req *pb.CloudSkinLvUpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20717, data)
}

// 请求宝地挑战
func (c *Client) SendWildBossChallengeReq(ctx context.Context, req *pb.WildBossChallengeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20732, data)
}

// 请求宝地扫荡
func (c *Client) SendWildBossRepeatReq(ctx context.Context, req *pb.WildBossRepeatReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20733, data)
}

// 捕捉灵兽请求
func (c *Client) SendCatchPetReq(ctx context.Context, req *pb.CatchPetReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20741, data)
}

// 刷新灵兽池子请求
func (c *Client) SendRefreshPetPoolReq(ctx context.Context, req *pb.RefreshPetPoolReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20742, data)
}

// 开启灵兽背包格子请求
func (c *Client) SendAddPetBagCountReq(ctx context.Context, req *pb.AddPetBagCountReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20743, data)
}

// 灵兽放生请求
func (c *Client) SendReleasePetReq(ctx context.Context, req *pb.ReleasePetReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20744, data)
}

// 灵兽升级请求
func (c *Client) SendPetLevelUpReq(ctx context.Context, req *pb.PetLevelUpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20745, data)
}

// 灵兽吞噬请求
func (c *Client) SendPetGobbleUpReq(ctx context.Context, req *pb.PetGobbleUpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20746, data)
}

// 携带灵兽请求
func (c *Client) SendEquipPetReq(ctx context.Context, req *pb.EquipPetReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20747, data)
}

// 携带灵兽请求
func (c *Client) SendPetResetReq(ctx context.Context, req *pb.PetResetReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20748, data)
}

// 灵兽心愿选择
func (c *Client) SendChooseWishPetReq(ctx context.Context, req *pb.ChooseWishPetReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20749, data)
}

// 灵兽协助
func (c *Client) SendPetAssistantReq(ctx context.Context, req *pb.PetAssistantReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20750, data)
}

// 灵兽塑魂
func (c *Client) SendPetSoulShapeReq(ctx context.Context, req *pb.PetSoulShapeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20751, data)
}

// 灵兽塑魂技能选择
func (c *Client) SendSelectSoulShapeSkillReq(ctx context.Context, req *pb.SelectSoulShapeSkillReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20752, data)
}

// 灵兽洗练界面
func (c *Client) SendPetSkillRefreshViewReq(ctx context.Context, req *pb.PetSkillRefreshViewReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20753, data)
}

// 灵兽洗练
func (c *Client) SendPetSkillRefreshReq(ctx context.Context, req *pb.PetSkillRefreshReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20754, data)
}

// 灵兽洗练结果处理
func (c *Client) SendPetSkillRefreshResultReq(ctx context.Context, req *pb.PetSkillRefreshResultReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20755, data)
}

// 灵兽加锁
func (c *Client) SendPetLockStateReq(ctx context.Context, req *pb.PetLockStateReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20756, data)
}

// 切换灵兽联动形象
func (c *Client) SendPetSwitchLinkageSkinReq(ctx context.Context, req *pb.PetSwitchLinkageSkinReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20757, data)
}

// 灵兽新版上锁（需要密码解锁）
func (c *Client) SendPetLockReq(ctx context.Context, req *pb.PetLockReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20758, data)
}

// 选择加成请求
func (c *Client) SendSelectBuffReq(ctx context.Context, req *pb.SelectBuffReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20764, data)
}

// 获取选择选项
func (c *Client) SendGetSelectOptionReq(ctx context.Context, req *pb.GetSelectOptionReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20766, data)
}

// 保存选择选项
func (c *Client) SendSaveSelectOptionReq(ctx context.Context, req *pb.SaveSelectOptionReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20767, data)
}

// 获取排行列表
func (c *Client) SendGetRankListReq(ctx context.Context, req *pb.GetRankListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20801, data)
}

// 首充领取奖励
func (c *Client) SendGetMallRewardReq(ctx context.Context, req *pb.GetMallRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20802, data)
}

// 查看联盟积分
func (c *Client) SendGetUnionMemberScoreReq(ctx context.Context, req *pb.GetUnionMemberScoreReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20803, data)
}

// 精怪抽取请求
func (c *Client) SendSpiritDrawReq(ctx context.Context, req *pb.SpiritDrawReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20822, data)
}

// 精怪解锁请求
func (c *Client) SendSpiritUnlockReq(ctx context.Context, req *pb.SpiritUnlockReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20823, data)
}

// 精怪升级请求
func (c *Client) SendSpiritLvUpReq(ctx context.Context, req *pb.SpiritLvUpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20824, data)
}

// 切换出战队伍请求
func (c *Client) SendSwitchBattleTeamReq(ctx context.Context, req *pb.SwitchBattleTeamReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20825, data)
}

// 精怪上下阵请求
func (c *Client) SendSpiritBattleReq(ctx context.Context, req *pb.SpiritBattleReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20826, data)
}

// 精怪共鸣等级提升
func (c *Client) SendSpiritCombineLvUpReq(ctx context.Context, req *pb.SpiritCombineLvUpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20827, data)
}

// 切换精怪联动形象
func (c *Client) SendSpiritSwitchLinkageSkinReq(ctx context.Context, req *pb.SpiritSwitchLinkageSkinReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20828, data)
}

// 精怪等级解锁展示
func (c *Client) SendSpiritLevelUnlockShowReq(ctx context.Context, req *pb.SpiritLevelUnlockShowReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 20829, data)
}

// 请求活动通用数据
func (c *Client) SendReqGetActivityDetail(ctx context.Context, req *pb.ReqGetActivityDetail) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21003, data)
}

// 领取活动任务条件奖励
func (c *Client) SendReqGetActivityConditionReward21004(ctx context.Context, req *pb.ReqGetActivityConditionReward) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21004, data)
}

// 购买活动商品
func (c *Client) SendReqBuyActivityGoods(ctx context.Context, req *pb.ReqBuyActivityGoods) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21005, data)
}

// 领取活动排行奖励
func (c *Client) SendReqGetActivityRankReward(ctx context.Context, req *pb.ReqGetActivityRankReward) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21006, data)
}

// 领取基金活动任务条件奖励
func (c *Client) SendReqGetFundsConditionReward(ctx context.Context, req *pb.ReqGetFundsConditionReward) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21010, data)
}

// 运势抽奖请求
func (c *Client) SendLuckyDrawReq(ctx context.Context, req *pb.LuckyDrawReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21013, data)
}

// 自选礼包选择道具
func (c *Client) SendSelectItemsReq(ctx context.Context, req *pb.SelectItemsReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21014, data)
}

// 活动分享任务完成请求
func (c *Client) SendReqShareTaskDone21016(ctx context.Context, req *pb.ReqShareTaskDone) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21016, data)
}

// 请求榜单状态
func (c *Client) SendReqGetActivityRankState(ctx context.Context, req *pb.ReqGetActivityRankState) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21017, data)
}

// 通过数组一次性领取活动任务条件奖励
func (c *Client) SendReqGetActivityConditionRewardByArr(ctx context.Context, req *pb.ReqGetActivityConditionRewardByArr) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21022, data)
}

// 妖市觅宝抽奖
func (c *Client) SendSeekTreasureDrawReq(ctx context.Context, req *pb.SeekTreasureDrawReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21023, data)
}

// 妖市觅宝界面
func (c *Client) SendSeekTreasureViewReq(ctx context.Context, req *pb.SeekTreasureViewReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21024, data)
}

// 选择稀有奖励
func (c *Client) SendSeekTreasureChooseRareRewardReq(ctx context.Context, req *pb.SeekTreasureChooseRareRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21025, data)
}

// 选择稀有奖励
func (c *Client) SendSeekTreasureSelectRewardReqMsg(ctx context.Context, req *pb.SeekTreasureSelectRewardReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21028, data)
}

// 论坛分享
func (c *Client) SendReqShareTaskDone21030(ctx context.Context, req *pb.ReqShareTaskDone) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21030, data)
}

// 游戏圈分享
func (c *Client) SendReqShareTaskDone21031(ctx context.Context, req *pb.ReqShareTaskDone) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21031, data)
}

// 自选商品选中协议
func (c *Client) SendSelectActivityGoodsReq(ctx context.Context, req *pb.SelectActivityGoodsReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21032, data)
}

// 领取活动战令任务奖励
func (c *Client) SendReqGetActivityConditionReward21033(ctx context.Context, req *pb.ReqGetActivityConditionReward) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21033, data)
}

// 活动任务自选商品
func (c *Client) SendSelectActivityConditionGoodsReq(ctx context.Context, req *pb.SelectActivityConditionGoodsReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21034, data)
}

// 请求竞猜数据
func (c *Client) SendReqGuessInfoLoadMsg(ctx context.Context, req *pb.ReqGuessInfoLoadMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21040, data)
}

// 竞猜押注
func (c *Client) SendReqGuessMsg(ctx context.Context, req *pb.ReqGuessMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21041, data)
}

// 领取竞猜奖励
func (c *Client) SendReqGuessRewardMsg(ctx context.Context, req *pb.ReqGuessRewardMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21042, data)
}

// 查看竞猜奖励名单
func (c *Client) SendReqGuessResultDetailMsg(ctx context.Context, req *pb.ReqGuessResultDetailMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21043, data)
}

// 领取跨服获得内置的 妖盟充值活动的奖励
func (c *Client) SendReceiveCrossUnionRechargeRewardReq(ctx context.Context, req *pb.ReceiveCrossUnionRechargeRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21045, data)
}

// 查看自己所属妖盟的参与名单
func (c *Client) SendActivityGetJoinMemberListReq(ctx context.Context, req *pb.ActivityGetJoinMemberListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21046, data)
}

// 从PC打开
func (c *Client) SendReqShareTaskDone21047(ctx context.Context, req *pb.ReqShareTaskDone) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21047, data)
}

// 进入福地
func (c *Client) SendHomelandEnterReq(ctx context.Context, req *pb.HomelandEnterReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21052, data)
}

// 福地管理界面
func (c *Client) SendHomelandManageReq(ctx context.Context, req *pb.HomelandManageReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21053, data)
}

// 福地升级
func (c *Client) SendHomelandLevelUpReq(ctx context.Context, req *pb.HomelandLevelUpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21054, data)
}

// 福地刷新资源
func (c *Client) SendHomelandRefreshReq(ctx context.Context, req *pb.HomelandRefreshReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21055, data)
}

// 福地购买鼠宝
func (c *Client) SendHomelandBuyWorkerReq(ctx context.Context, req *pb.HomelandBuyWorkerReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21056, data)
}

// 福地日志
func (c *Client) SendHomelandGetRecordReq(ctx context.Context, req *pb.HomelandGetRecordReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21057, data)
}

// 福地探寻
func (c *Client) SendHomelandExploreReq(ctx context.Context, req *pb.HomelandExploreReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21058, data)
}

// 福地探寻刷新
func (c *Client) SendHomelandRefreshExploreReq(ctx context.Context, req *pb.HomelandRefreshExploreReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21059, data)
}

// 福地派遣鼠宝
func (c *Client) SendHomelandDispatchWorkerReq(ctx context.Context, req *pb.HomelandDispatchWorkerReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21060, data)
}

// 福地派遣鼠宝预览
func (c *Client) SendHomelandDispatchPreviewReq(ctx context.Context, req *pb.HomelandDispatchPreviewReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21061, data)
}

// 离开福地
func (c *Client) SendHomelandLeaveReq(ctx context.Context, req *pb.HomelandLeaveReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21063, data)
}

// 福地获取奖励
func (c *Client) SendHomelandGetRewardReq(ctx context.Context, req *pb.HomelandGetRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21065, data)
}

// 福地聚宝盆领奖
func (c *Client) SendCornucopiaGetRewardReq(ctx context.Context, req *pb.CornucopiaGetRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21067, data)
}

// 福地鼠鼠自动采集
func (c *Client) SendHomelandAutoDispatchWorkerReq(ctx context.Context, req *pb.HomelandAutoDispatchWorkerReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21068, data)
}

// 福地鼠宝总管试用
func (c *Client) SendHomelandUseFreeMouseManagerReq(ctx context.Context, req *pb.HomelandUseFreeMouseManagerReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21069, data)
}

// 鼠鼠总管保存设置
func (c *Client) SendHomelandSaveSettingsReq(ctx context.Context, req *pb.HomelandSaveSettingsReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21070, data)
}

// 桃源一键召回小妖
func (c *Client) SendHomelandDispatchAllWorkerReq(ctx context.Context, req *pb.HomelandDispatchAllWorkerReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21071, data)
}

// 请求报名
func (c *Client) SendBallGVGApplicationReqMsg(ctx context.Context, req *pb.BallGVGApplicationReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21101, data)
}

// 请求报名期战力排名
func (c *Client) SendBallGVGAbilityRankNumReqMsg(ctx context.Context, req *pb.BallGVGAbilityRankNumReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21102, data)
}

// 进入pvp活动请求协议
func (c *Client) SendBallGVGEnterActivityReqMsg(ctx context.Context, req *pb.BallGVGEnterActivityReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21103, data)
}

// 进入宝地(房间)请求协议
func (c *Client) SendBallGVGEnterPlaceReqMsg(ctx context.Context, req *pb.BallGVGEnterPlaceReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21104, data)
}

// 离开当前房间请求协议
func (c *Client) SendBallGVGLeavePlaceReqMsg(ctx context.Context, req *pb.BallGVGLeavePlaceReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21105, data)
}

// 移动协议
func (c *Client) SendBallGVGMoveReqMsg(ctx context.Context, req *pb.BallGVGMoveReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21106, data)
}

// 攻击协议
func (c *Client) SendBallGVGAttackReqMsg(ctx context.Context, req *pb.BallGVGAttackReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21107, data)
}

// 任命宗主权限
func (c *Client) SendBallGVGAppointMajorUserReqMsg(ctx context.Context, req *pb.BallGVGAppointMajorUserReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21108, data)
}

// 标记地点(重复标记视为取消标记)
func (c *Client) SendBallGVGMarkPlaceReqMsg(ctx context.Context, req *pb.BallGVGMarkPlaceReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21109, data)
}

// 获取所有房间数据
func (c *Client) SendBallGVGGetPlaceInfoReqMsg(ctx context.Context, req *pb.BallGVGGetPlaceInfoReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21111, data)
}

// 请求宗门玩家列表
func (c *Client) SendBallGVGGetUserRankReqMsg(ctx context.Context, req *pb.BallGVGGetUserRankReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21113, data)
}

// 请求报名期妖盟战力排行
func (c *Client) SendBallGVGAbilityRankReqMsg(ctx context.Context, req *pb.BallGVGAbilityRankReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21126, data)
}

// 获取参与妖盟的名字
func (c *Client) SendBallGVGUnionNameListReqMsg(ctx context.Context, req *pb.BallGVGUnionNameListReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21128, data)
}

// 获取上古遗迹妖盟成员信息
func (c *Client) SendBallGVGGetUnionUserListReqMsg(ctx context.Context, req *pb.BallGVGGetUnionUserListReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21129, data)
}

// 上古遗迹红点
func (c *Client) SendBallGVGRedDotReqMsg(ctx context.Context, req *pb.BallGVGRedDotReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21131, data)
}

// 仙缘进度领奖
func (c *Client) SendXyFundGetRewardReq(ctx context.Context, req *pb.XyFundGetRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21421, data)
}

// 获取商会同名姓名列表
func (c *Client) SendReqUnionSameNameListMsg(ctx context.Context, req *pb.ReqUnionSameNameListMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21501, data)
}

// 获取商会充值列表
func (c *Client) SendUnionRechargeUserReqMsg(ctx context.Context, req *pb.UnionRechargeUserReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 21511, data)
}

// 创建妖盟
func (c *Client) SendReqUnionCreate(ctx context.Context, req *pb.ReqUnionCreate) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22101, data)
}

// 进入妖盟
func (c *Client) SendReqUnionEnter(ctx context.Context, req *pb.ReqUnionEnter) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22102, data)
}

// 获取妖盟列表
func (c *Client) SendReqUnionList(ctx context.Context, req *pb.ReqUnionList) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22103, data)
}

// 加入妖盟
func (c *Client) SendReqUnionJoin(ctx context.Context, req *pb.ReqUnionJoin) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22104, data)
}

// 随机加入妖盟
func (c *Client) SendReqUnionRandomJoin(ctx context.Context, req *pb.ReqUnionRandomJoin) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22105, data)
}

// 妖盟详情
func (c *Client) SendReqUnionDetail22106(ctx context.Context, req *pb.ReqUnionDetail) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22106, data)
}

// 职位变更
func (c *Client) SendReqUnionPosition(ctx context.Context, req *pb.ReqUnionPosition) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22107, data)
}

// 获取申请加入妖盟的玩家列表
func (c *Client) SendReqUnionApplyPlayerList(ctx context.Context, req *pb.ReqUnionApplyPlayerList) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22108, data)
}

// 设置随机加入状态
func (c *Client) SendReqUnionRandomState(ctx context.Context, req *pb.ReqUnionRandomState) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22109, data)
}

// 设置入盟限制等级
func (c *Client) SendReqUnionLimitRealmsId(ctx context.Context, req *pb.ReqUnionLimitRealmsId) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22110, data)
}

// 清理所有申请玩家
func (c *Client) SendReqUnionClearApply(ctx context.Context, req *pb.ReqUnionClearApply) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22111, data)
}

// 通过新玩家的申请
func (c *Client) SendReqUnionPassApply(ctx context.Context, req *pb.ReqUnionPassApply) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22112, data)
}

// 解散妖盟
func (c *Client) SendReqUnionRemove(ctx context.Context, req *pb.ReqUnionRemove) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22113, data)
}

// 退出妖盟
func (c *Client) SendReqUnionExit(ctx context.Context, req *pb.ReqUnionExit) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22114, data)
}

// 获取妖盟动态
func (c *Client) SendReqUnionTrends(ctx context.Context, req *pb.ReqUnionTrends) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22115, data)
}

// 妖盟信息修改
func (c *Client) SendReqUnionModify(ctx context.Context, req *pb.ReqUnionModify) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22116, data)
}

// 获取每日任务列表
func (c *Client) SendReqUnionDailyTask(ctx context.Context, req *pb.ReqUnionDailyTask) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22117, data)
}

// 领取每日任务奖励
func (c *Client) SendReqUnionGetDailyTask(ctx context.Context, req *pb.ReqUnionGetDailyTask) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22118, data)
}

// 每日任务捐献
func (c *Client) SendReqUnionDailyDonate(ctx context.Context, req *pb.ReqUnionDailyDonate) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22120, data)
}

// 获取每日任务成员完成进度
func (c *Client) SendReqUnionDailyProgress(ctx context.Context, req *pb.ReqUnionDailyProgress) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22121, data)
}

// 逐出妖盟
func (c *Client) SendReqUnionEviction(ctx context.Context, req *pb.ReqUnionEviction) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22122, data)
}

// 加入的妖盟列表(我没有妖盟时的列表)
func (c *Client) SendReqUnionJoinList(ctx context.Context, req *pb.ReqUnionJoinList) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22123, data)
}

// 获取妖盟协助列表
func (c *Client) SendGetUnionHelpDataListReq(ctx context.Context, req *pb.GetUnionHelpDataListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22125, data)
}

// 请求妖盟协助
func (c *Client) SendRequestUnionHelpReq(ctx context.Context, req *pb.RequestUnionHelpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22126, data)
}

// 执行妖盟协助
func (c *Client) SendUnionHelpReq(ctx context.Context, req *pb.UnionHelpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22127, data)
}

// 获取妖盟协助状态
func (c *Client) SendGetUnionHelpStateReq(ctx context.Context, req *pb.GetUnionHelpStateReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22128, data)
}

// 查看妖盟的红点
func (c *Client) SendReqUnionWatchRedDot(ctx context.Context, req *pb.ReqUnionWatchRedDot) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22129, data)
}

// 砍价
func (c *Client) SendCutPriceBargainReqMsg(ctx context.Context, req *pb.CutPriceBargainReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22166, data)
}

// 砍价购买
func (c *Client) SendCutPriceBuyReqMsg(ctx context.Context, req *pb.CutPriceBuyReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22167, data)
}

// 绑定邀请码
func (c *Client) SendReqBindInviteCode(ctx context.Context, req *pb.ReqBindInviteCode) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22401, data)
}

// 获取绑定他人的邀请码
func (c *Client) SendReqGetBindInviteCode(ctx context.Context, req *pb.ReqGetBindInviteCode) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 22402, data)
}

// 请求挑战者列表
func (c *Client) SendReqHeroRankFightPlayerList(ctx context.Context, req *pb.ReqHeroRankFightPlayerList) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 23702, data)
}

// 请求挑战玩家
func (c *Client) SendReqHeroRankFight(ctx context.Context, req *pb.ReqHeroRankFight) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 23703, data)
}

// 请求扫荡玩家
func (c *Client) SendReqHeroRankClear(ctx context.Context, req *pb.ReqHeroRankClear) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 23704, data)
}

// 请求领取成就奖励
func (c *Client) SendReqHeroRankGetAchieve(ctx context.Context, req *pb.ReqHeroRankGetAchieve) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 23706, data)
}

// 请求购买体力
func (c *Client) SendReqHeroRankBuyEnergy(ctx context.Context, req *pb.ReqHeroRankBuyEnergy) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 23707, data)
}

// 抽奖
func (c *Client) SendReqPetDreamLandDraw(ctx context.Context, req *pb.ReqPetDreamLandDraw) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24101, data)
}

// 区域解锁
func (c *Client) SendReqPetDreamLandAdventurePlaceUnlock(ctx context.Context, req *pb.ReqPetDreamLandAdventurePlaceUnlock) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24102, data)
}

// 坑位解锁
func (c *Client) SendReqPetDreamLandAdventureSlotUnlock(ctx context.Context, req *pb.ReqPetDreamLandAdventureSlotUnlock) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24103, data)
}

// 灵兽派遣
func (c *Client) SendReqPetDreamLandAdventureDispatch(ctx context.Context, req *pb.ReqPetDreamLandAdventureDispatch) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24104, data)
}

// 领取收益
func (c *Client) SendReqPetDreamLandAdventureGetAward(ctx context.Context, req *pb.ReqPetDreamLandAdventureGetAward) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24105, data)
}

// 刷新邪祟
func (c *Client) SendRefreshEvilThingReq(ctx context.Context, req *pb.RefreshEvilThingReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24201, data)
}

// 镇压邪祟
func (c *Client) SendBattleEviThingReq(ctx context.Context, req *pb.BattleEviThingReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24202, data)
}

// 装备神通
func (c *Client) SendMagicEquipReq(ctx context.Context, req *pb.MagicEquipReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24401, data)
}

// 重置神通
func (c *Client) SendMagicResetReq(ctx context.Context, req *pb.MagicResetReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24402, data)
}

// 进阶神通
func (c *Client) SendMagicStageUpReq(ctx context.Context, req *pb.MagicStageUpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24403, data)
}

// 升级神通
func (c *Client) SendMagicLvUpReq(ctx context.Context, req *pb.MagicLvUpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24404, data)
}

// 装备印记
func (c *Client) SendMagicEquipMarkReq(ctx context.Context, req *pb.MagicEquipMarkReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24405, data)
}

// 卸下印记
func (c *Client) SendMagicUnsnatchMarkReq(ctx context.Context, req *pb.MagicUnsnatchMarkReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24406, data)
}

// 融合神通印记
func (c *Client) SendMagicFusionMarkReq(ctx context.Context, req *pb.MagicFusionMarkReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24407, data)
}

// 衍化
func (c *Client) SendMagicDerivationReq(ctx context.Context, req *pb.MagicDerivationReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24408, data)
}

// 升级组合
func (c *Client) SendMagicCombineLvUpReq(ctx context.Context, req *pb.MagicCombineLvUpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24409, data)
}

// 切换预设
func (c *Client) SendMagicSwitchPresetsReq(ctx context.Context, req *pb.MagicSwitchPresetsReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24410, data)
}

// 一键突破神通
func (c *Client) SendMagicStageUpAllReq(ctx context.Context, req *pb.MagicStageUpAllReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24411, data)
}

// 获取预设神通列表
func (c *Client) SendGetPresetMagicInfoReq(ctx context.Context, req *pb.GetPresetMagicInfoReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24412, data)
}

// 拆卸装备的印记
func (c *Client) SendUnEquipStoneReq(ctx context.Context, req *pb.UnEquipStoneReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24413, data)
}

// 添加储存数据
func (c *Client) SendSaveToServiceReq(ctx context.Context, req *pb.SaveToServiceReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24501, data)
}

// 仙宫点赞同步
func (c *Client) SendPalaceWorshipReq(ctx context.Context, req *pb.PalaceWorshipReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24802, data)
}

// 仙宫外部数据请求
func (c *Client) SendEnterPalaceReq(ctx context.Context, req *pb.EnterPalaceReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24803, data)
}

// 仙宫内部数据请求
func (c *Client) SendEnterPalaceInnerReq(ctx context.Context, req *pb.EnterPalaceInnerReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24804, data)
}

// 仙名录数据请求
func (c *Client) SendPalaceInnerBookReq(ctx context.Context, req *pb.PalaceInnerBookReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24805, data)
}

// 执行仙宫送福
func (c *Client) SendSendGiftReq(ctx context.Context, req *pb.SendGiftReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24806, data)
}

// 仙宫送福领奖
func (c *Client) SendGetSendGiftRewardReq(ctx context.Context, req *pb.GetSendGiftRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24807, data)
}

// 获得当前已开启的天宫称号id
func (c *Client) SendGetTitleIdListReq(ctx context.Context, req *pb.GetTitleIdListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24810, data)
}

// 仙宫获得成就奖励
func (c *Client) SendPalaceAchievementGetRewardReq(ctx context.Context, req *pb.PalaceAchievementGetRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 24857, data)
}

// 蛮荒 进入活动页
func (c *Client) SendManHuangEnterPanelReqMsg(ctx context.Context, req *pb.ManHuangEnterPanelReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25001, data)
}

// 蛮荒 创建队伍
func (c *Client) SendManHuangCreateTeamReqMsg(ctx context.Context, req *pb.ManHuangCreateTeamReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25002, data)
}

// 蛮荒 获取队伍列表
func (c *Client) SendManHuangGetTeamListReqMsg(ctx context.Context, req *pb.ManHuangGetTeamListReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25003, data)
}

// 蛮荒 查看队伍数据
func (c *Client) SendManHuangGetTeamInfoReqMsg(ctx context.Context, req *pb.ManHuangGetTeamInfoReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25004, data)
}

// 蛮荒 加入队伍
func (c *Client) SendManHuangJoinTeamReqMsg(ctx context.Context, req *pb.ManHuangJoinTeamReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25005, data)
}

// 蛮荒 取消申请队伍
func (c *Client) SendManHuangCancelTeamApplyReqMsg(ctx context.Context, req *pb.ManHuangCancelTeamApplyReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25006, data)
}

// 蛮荒 解散/退出队伍
func (c *Client) SendManHuangQuitTeamReqMsg(ctx context.Context, req *pb.ManHuangQuitTeamReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25007, data)
}

// 蛮荒 退出活动
func (c *Client) SendManHuangLeavePanelReqMsg(ctx context.Context, req *pb.ManHuangLeavePanelReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25008, data)
}

// 蛮荒 玩家申请入队同意
func (c *Client) SendManHuangApplyJoinTeamAgreeReqMsg(ctx context.Context, req *pb.ManHuangApplyJoinTeamAgreeReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25009, data)
}

// 蛮荒 玩家申请入队一键拒绝
func (c *Client) SendManHuangApplyJoinTeamRefusedReqMsg(ctx context.Context, req *pb.ManHuangApplyJoinTeamRefusedReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25010, data)
}

// 蛮荒 踢出队伍
func (c *Client) SendManHuangKickOutTeamReqMsg(ctx context.Context, req *pb.ManHuangKickOutTeamReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25011, data)
}

// 蛮荒 进入区域
func (c *Client) SendManHuangEnterRegionReqMsg(ctx context.Context, req *pb.ManHuangEnterRegionReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25012, data)
}

// 蛮荒 探索
func (c *Client) SendManHuangExploreReqMsg(ctx context.Context, req *pb.ManHuangExploreReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25013, data)
}

// 蛮荒 执行探索事件 
func (c *Client) SendManHuangEventHandleReqMsg(ctx context.Context, req *pb.ManHuangEventHandleReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25014, data)
}

// 蛮荒 action探索事件(标记，绕过) 
func (c *Client) SendManHuangEventActionReqMsg(ctx context.Context, req *pb.ManHuangEventActionReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25015, data)
}

// 蛮荒 日志-队友互助 
func (c *Client) SendManHuangLogHelpReqMsg(ctx context.Context, req *pb.ManHuangLogHelpReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25016, data)
}

// 蛮荒 日志-战场日志 
func (c *Client) SendManHuangLogBattlegroundReqMsg(ctx context.Context, req *pb.ManHuangLogBattlegroundReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25017, data)
}

// 蛮荒 日志-个人日志 
func (c *Client) SendManHuangLogPersonReqMsg(ctx context.Context, req *pb.ManHuangLogPersonReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25018, data)
}

// 蛮荒 排行-个人 
func (c *Client) SendManHuangRankPersonReqMsg(ctx context.Context, req *pb.ManHuangRankPersonReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25019, data)
}

// 蛮荒 排行-队伍 
func (c *Client) SendManHuangRankTeamReqMsg(ctx context.Context, req *pb.ManHuangRankTeamReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25020, data)
}

// 蛮荒 恢复体力
func (c *Client) SendManHuangRecoverEnergyReqMsg(ctx context.Context, req *pb.ManHuangRecoverEnergyReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25025, data)
}

// 蛮荒 执行协助
func (c *Client) SendManHuangHelpActionAttReqMsg(ctx context.Context, req *pb.ManHuangHelpActionAttReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25026, data)
}

// 蛮荒 开宝箱
func (c *Client) SendManHuangOpenBoxReqMsg(ctx context.Context, req *pb.ManHuangOpenBoxReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25027, data)
}

// 蛮荒 领取协助奖励
func (c *Client) SendManHuangHelpActionRewardReqMsg(ctx context.Context, req *pb.ManHuangHelpActionRewardReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25028, data)
}

// 蛮荒 发送同步玩家信息协议
func (c *Client) SendManHuangGetUserDataReqMsg(ctx context.Context, req *pb.ManHuangGetUserDataReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25030, data)
}

// 精炼
func (c *Client) SendReqEquipmentRefine(ctx context.Context, req *pb.ReqEquipmentRefine) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25501, data)
}

// 进阶
func (c *Client) SendReqEquipmentAdvance(ctx context.Context, req *pb.ReqEquipmentAdvance) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25502, data)
}

// 技能解锁
func (c *Client) SendReqEquipmentActivate(ctx context.Context, req *pb.ReqEquipmentActivate) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25503, data)
}

// 进入秘境
func (c *Client) SendSecretTowerEnterReq(ctx context.Context, req *pb.SecretTowerEnterReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25601, data)
}

// 挑战怪物
func (c *Client) SendSecretTowerFightReq(ctx context.Context, req *pb.SecretTowerFightReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25602, data)
}

// 领取阶段奖励
func (c *Client) SendSecretTowerGetStageRewardReq(ctx context.Context, req *pb.SecretTowerGetStageRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25603, data)
}

// 查看怪物属性
func (c *Client) SendSecretTowerViewMonsterAttrReq(ctx context.Context, req *pb.SecretTowerViewMonsterAttrReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25604, data)
}

// 查看秘境成就
func (c *Client) SendSecretTowerAchievementReq(ctx context.Context, req *pb.SecretTowerAchievementReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25606, data)
}

// 领取成就奖励
func (c *Client) SendSecretTowerGetAchiRewardReq(ctx context.Context, req *pb.SecretTowerGetAchiRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25607, data)
}

// 进入妖邪讨伐
func (c *Client) SendUnionBossInfoReqMsg(ctx context.Context, req *pb.UnionBossInfoReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25801, data)
}

// 布阵
func (c *Client) SendUnionBossBuffReqMsg(ctx context.Context, req *pb.UnionBossBuffReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25802, data)
}

// 查看讨伐奖励
func (c *Client) SendUnionBossRewardReqMsg(ctx context.Context, req *pb.UnionBossRewardReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25803, data)
}

// 领取讨伐奖励
func (c *Client) SendUnionBossRewardReceiveReqMsg(ctx context.Context, req *pb.UnionBossRewardReceiveReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25804, data)
}

// 挑战
func (c *Client) SendUnionBossBattleReqMsg(ctx context.Context, req *pb.UnionBossBattleReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25805, data)
}

// 领取讨伐成就奖励
func (c *Client) SendUnionBossAchieveRewardReqMsg(ctx context.Context, req *pb.UnionBossAchieveRewardReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25806, data)
}

// 请求已布阵玩家信息
func (c *Client) SendUnionBossAddBuffPlayerReqMsg(ctx context.Context, req *pb.UnionBossAddBuffPlayerReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 25809, data)
}

// 法宝抽取
func (c *Client) SendMagicTreasureDrawReq(ctx context.Context, req *pb.MagicTreasureDrawReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 26302, data)
}

// 法宝升级
func (c *Client) SendMagicTreasureLvUpReq(ctx context.Context, req *pb.MagicTreasureLvUpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 26303, data)
}

// 法宝激活
func (c *Client) SendMagicTreasureActiveReq(ctx context.Context, req *pb.MagicTreasureActiveReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 26304, data)
}

// 法宝升星
func (c *Client) SendMagicTreasureStarUpReq(ctx context.Context, req *pb.MagicTreasureStarUpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 26305, data)
}

// 法宝形象切换
func (c *Client) SendMagicTreasureSwitchLinkageSkinReq(ctx context.Context, req *pb.MagicTreasureSwitchLinkageSkinReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 26307, data)
}

// 小游戏领奖
func (c *Client) SendReceiveMiniGamesRewardReq(ctx context.Context, req *pb.ReceiveMiniGamesRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 200021, data)
}

// 秘境挑战
func (c *Client) SendStageMapChallengeReq(ctx context.Context, req *pb.StageMapChallengeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 200022, data)
}

// 禁地试炼领奖
func (c *Client) SendForbiddenTrailsGetRewardReq(ctx context.Context, req *pb.ForbiddenTrailsGetRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 200023, data)
}

// 西天救援进入
func (c *Client) SendEnterWestJourneyReq(ctx context.Context, req *pb.EnterWestJourneyReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 200024, data)
}

// 西天救援领奖
func (c *Client) SendGetWestJourneyRewardReq(ctx context.Context, req *pb.GetWestJourneyRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 200026, data)
}

// 查看本分组妖盟详情
func (c *Client) SendReqUnionDetail201604(ctx context.Context, req *pb.ReqUnionDetail) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 201604, data)
}

// 获取分组评级排行
func (c *Client) SendUnionPlayerGradeRankReqMsg(ctx context.Context, req *pb.UnionPlayerGradeRankReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 201630, data)
}

// 评级人数分配设置
func (c *Client) SendUnionMemberGradeSettingReq(ctx context.Context, req *pb.UnionMemberGradeSettingReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 201650, data)
}

// 获取妖盟跨服信息
func (c *Client) SendGetCrossUnionGroupServersReq(ctx context.Context, req *pb.GetCrossUnionGroupServersReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 201651, data)
}

// 发布招募信息
func (c *Client) SendUnionPublishRecruitReqMsg(ctx context.Context, req *pb.UnionPublishRecruitReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 201654, data)
}

// 移除黑名单
func (c *Client) SendRemoveBlackReqMsg(ctx context.Context, req *pb.RemoveBlackReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 203205, data)
}

// 获取黑名单列表
func (c *Client) SendGetBlackPlayerListReqMsg(ctx context.Context, req *pb.GetBlackPlayerListReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 203206, data)
}

// 检测消息是否有效
func (c *Client) SendWorldMessageCheckEffectiveReq(ctx context.Context, req *pb.WorldMessageCheckEffectiveReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 203208, data)
}

// 进入妖盟乱斗
func (c *Client) SendUnionBattleEnterReq(ctx context.Context, req *pb.UnionBattleEnterReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 205301, data)
}

// 查看自己所属联盟的参与名单
func (c *Client) SendUnionBattleGetJoinMemberListReq(ctx context.Context, req *pb.UnionBattleGetJoinMemberListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 205302, data)
}

// 匹配
func (c *Client) SendUnionBattleMatchReq(ctx context.Context, req *pb.UnionBattleMatchReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 205303, data)
}

// 挑战
func (c *Client) SendUnionBattleChallengeReq(ctx context.Context, req *pb.UnionBattleChallengeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 205304, data)
}

// 查看战报信息
func (c *Client) SendUnionBattleGetReportReq(ctx context.Context, req *pb.UnionBattleGetReportReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 205306, data)
}

// 购买buff
func (c *Client) SendUnionBattleBuyOpenBuffReq(ctx context.Context, req *pb.UnionBattleBuyOpenBuffReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 205307, data)
}

// 选择buff
func (c *Client) SendUnionBattleSelectBuffReq(ctx context.Context, req *pb.UnionBattleSelectBuffReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 205308, data)
}

// 获取本次挑战的回放列表
func (c *Client) SendUnionBattleGetPlayBackListReq(ctx context.Context, req *pb.UnionBattleGetPlayBackListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 205309, data)
}

// 领取联盟成就奖励
func (c *Client) SendUnionBattleReceiveUnionAchieveRewardReq(ctx context.Context, req *pb.UnionBattleReceiveUnionAchieveRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 205310, data)
}

// 查看联盟积分
func (c *Client) SendUnionBattleGetUnionMemberScoreReq(ctx context.Context, req *pb.UnionBattleGetUnionMemberScoreReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 205312, data)
}

// 红点
func (c *Client) SendUnionBattleRedDotReqMsg(ctx context.Context, req *pb.UnionBattleRedDotReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 205313, data)
}

// 领奖期 冠军弹窗
func (c *Client) SendAskDingGetChampionInfoReq(ctx context.Context, req *pb.AskDingGetChampionInfoReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 205935, data)
}

// 查询据点战基础信息
func (c *Client) SendUnionAreaWarReqMsg(ctx context.Context, req *pb.UnionAreaWarReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206001, data)
}

// 请求查看分组
func (c *Client) SendUnionAreaWarGroupReqMsg(ctx context.Context, req *pb.UnionAreaWarGroupReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206003, data)
}

// 请求防守人列表
func (c *Client) SendUnionAreaWarDefendersReqMsg(ctx context.Context, req *pb.UnionAreaWarDefendersReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206004, data)
}

// 刷新防守人列表
func (c *Client) SendUnionAreaWarDefendersUpdateReqMsg(ctx context.Context, req *pb.UnionAreaWarDefendersUpdateReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206005, data)
}

// 获取据点列表
func (c *Client) SendUnionAreWarListReqMsg(ctx context.Context, req *pb.UnionAreWarListReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206006, data)
}

// 刷新据点列表
func (c *Client) SendUnionAreWarListUpdateReqMsg(ctx context.Context, req *pb.UnionAreWarListUpdateReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206007, data)
}

// 请求战斗场景攻击
func (c *Client) SendUnionAreaWarAttackReqMsg(ctx context.Context, req *pb.UnionAreaWarAttackReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206009, data)
}

// 建设(捐献)
func (c *Client) SendUnionAreaWarDevelopReqMsg(ctx context.Context, req *pb.UnionAreaWarDevelopReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206011, data)
}

// 加载宝库信息
func (c *Client) SendUnionAreaWarTreasuryReqMsg(ctx context.Context, req *pb.UnionAreaWarTreasuryReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206013, data)
}

// 宝库抽奖
func (c *Client) SendUnionAreaWarTreasuryDrawReqMsg(ctx context.Context, req *pb.UnionAreaWarTreasuryDrawReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206014, data)
}

// 请求召唤神龙
func (c *Client) SendUnionAreaWarSummonDragonReqMsg(ctx context.Context, req *pb.UnionAreaWarSummonDragonReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206015, data)
}

// 查询据点战膜拜
func (c *Client) SendUnionAreaWarWorshipReq(ctx context.Context, req *pb.UnionAreaWarWorshipReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206016, data)
}

// 请求战斗场景数据
func (c *Client) SendUnionAreaWarFightSceneReqMsg(ctx context.Context, req *pb.UnionAreaWarFightSceneReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206019, data)
}

// 请求捐献建设数据
func (c *Client) SendUnionAreaWarConstructReqMsg(ctx context.Context, req *pb.UnionAreaWarConstructReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206020, data)
}

// 请求战报数据
func (c *Client) SendUnionAreaWarReportReqMsg(ctx context.Context, req *pb.UnionAreaWarReportReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206021, data)
}

// 请求战报伤害详情数据
func (c *Client) SendUnionAreaWarReportDetailReqMsg(ctx context.Context, req *pb.UnionAreaWarReportDetailReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206022, data)
}

// 请求联盟伤害数据
func (c *Client) SendUnionAreaWarUnionDamageReqMsg(ctx context.Context, req *pb.UnionAreaWarUnionDamageReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206023, data)
}

// 请求竞猜界面数据
func (c *Client) SendUnionAreaWarBetDataReqMsg(ctx context.Context, req *pb.UnionAreaWarBetDataReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206024, data)
}

// 请求竞猜选择
func (c *Client) SendUnionAreaWarBetSelectReqMsg(ctx context.Context, req *pb.UnionAreaWarBetSelectReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206025, data)
}

// 请求竞猜奖励
func (c *Client) SendUnionAreaWarBetRewardReqMsg(ctx context.Context, req *pb.UnionAreaWarBetRewardReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206026, data)
}

// 查看自己所属联盟的参与名单
func (c *Client) SendUnionAreaWarGetJoinMemberListReq(ctx context.Context, req *pb.UnionAreaWarGetJoinMemberListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206027, data)
}

// 请求查看联盟分段位伤害信息
func (c *Client) SendUnionAreaWarUnionGroupDamageReqMsg(ctx context.Context, req *pb.UnionAreaWarUnionGroupDamageReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206029, data)
}

// 请求红点信息
func (c *Client) SendUnionAreaWarRedDotReqMsg(ctx context.Context, req *pb.UnionAreaWarRedDotReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206030, data)
}

// 请求查看竞猜结果的获奖玩家列表
func (c *Client) SendUnionAreaWarGuessPlayersReqMsg(ctx context.Context, req *pb.UnionAreaWarGuessPlayersReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206031, data)
}

// 查看所有段位的分组数量
func (c *Client) SendUnionAreaWarDanGroupCountReq(ctx context.Context, req *pb.UnionAreaWarDanGroupCountReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206032, data)
}

// 进入活动
func (c *Client) SendAskWayEnterReq(ctx context.Context, req *pb.AskWayEnterReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206101, data)
}

// 斗法匹配
func (c *Client) SendAskWayMatchReq(ctx context.Context, req *pb.AskWayMatchReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206102, data)
}

// 斗法战斗
func (c *Client) SendAskWayBattleReq(ctx context.Context, req *pb.AskWayBattleReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206103, data)
}

// 获取战报
func (c *Client) SendAskWayGetReportReq(ctx context.Context, req *pb.AskWayGetReportReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206104, data)
}

// 斗法战斗回放
func (c *Client) SendAskWayBattleReplyReq(ctx context.Context, req *pb.AskWayBattleReplyReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206105, data)
}

// 领取段位奖励
func (c *Client) SendAskWayReceiveTierRewardReq(ctx context.Context, req *pb.AskWayReceiveTierRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206106, data)
}

// 领取斗法榜单积分奖励
func (c *Client) SendAskWayReceiveScoreRewardReq(ctx context.Context, req *pb.AskWayReceiveScoreRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206107, data)
}

// 购买挑战令道具
func (c *Client) SendAskWayBuyFightTicketReq(ctx context.Context, req *pb.AskWayBuyFightTicketReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206108, data)
}

// 查看玩家数据
func (c *Client) SendAskWayGetPlayerDetailReq(ctx context.Context, req *pb.AskWayGetPlayerDetailReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206109, data)
}

// 领取竞猜币
func (c *Client) SendAskWayGetGuessCoinReq(ctx context.Context, req *pb.AskWayGetGuessCoinReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206110, data)
}

// 登天竞猜
func (c *Client) SendAskWayToSkyGuessReq(ctx context.Context, req *pb.AskWayToSkyGuessReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206111, data)
}

// 登天战斗回放
func (c *Client) SendAskWayToSkyBattleReplyReq(ctx context.Context, req *pb.AskWayToSkyBattleReplyReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206112, data)
}

// 朝拜
func (c *Client) SendAskWayWorshipReq(ctx context.Context, req *pb.AskWayWorshipReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206113, data)
}

// 领取登天榜单奖励
func (c *Client) SendAskWayReceiveToSkyRankRewardReq(ctx context.Context, req *pb.AskWayReceiveToSkyRankRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206114, data)
}

// 往届妖神榜
func (c *Client) SendAskWayToSkyRosterReq(ctx context.Context, req *pb.AskWayToSkyRosterReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206115, data)
}

// 获取登天期组别数据
func (c *Client) SendAskWayToSkyGetBattleResultReq(ctx context.Context, req *pb.AskWayToSkyGetBattleResultReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206116, data)
}

// 获取竞猜信息
func (c *Client) SendAskWayGetGuessInfoReq(ctx context.Context, req *pb.AskWayGetGuessInfoReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206117, data)
}

// 获取登天战报数据
func (c *Client) SendAskWayToSkyGetReportReq(ctx context.Context, req *pb.AskWayToSkyGetReportReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206118, data)
}

// 获取海报数据
func (c *Client) SendAskWayGetGetCurStateInfoReq(ctx context.Context, req *pb.AskWayGetGetCurStateInfoReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206132, data)
}

// 妖盟夺位战主界面
func (c *Client) SendUnionFightMainReq(ctx context.Context, req *pb.UnionFightMainReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206701, data)
}

// 妖盟夺位战报名
func (c *Client) SendUnionFightApplyReq(ctx context.Context, req *pb.UnionFightApplyReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206702, data)
}

// 妖盟夺位战进入战斗界面
func (c *Client) SendUnionFightEnterReq(ctx context.Context, req *pb.UnionFightEnterReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206703, data)
}

// 妖盟夺位战请战
func (c *Client) SendUnionFightRequestReq(ctx context.Context, req *pb.UnionFightRequestReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206704, data)
}

// 妖盟夺位战上阵
func (c *Client) SendUnionFightPositionReq(ctx context.Context, req *pb.UnionFightPositionReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206705, data)
}

// 妖盟夺位战下阵
func (c *Client) SendUnionFightUnPositionReq(ctx context.Context, req *pb.UnionFightUnPositionReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206706, data)
}

// 妖盟夺位战小组排行
func (c *Client) SendUnionFightGroupRankReq(ctx context.Context, req *pb.UnionFightGroupRankReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206707, data)
}

// 妖盟夺位战离开
func (c *Client) SendUnionFightLeaveReq(ctx context.Context, req *pb.UnionFightLeaveReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206708, data)
}

// 妖盟夺位战至尊榜
func (c *Client) SendUnionFightSupremacyListReq(ctx context.Context, req *pb.UnionFightSupremacyListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206709, data)
}

// 妖盟夺位战膜拜
func (c *Client) SendUnionFightWorshipReq(ctx context.Context, req *pb.UnionFightWorshipReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206710, data)
}

// 妖盟夺位战回放
func (c *Client) SendUnionFightFightPlaybackReq(ctx context.Context, req *pb.UnionFightFightPlaybackReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206711, data)
}

// 妖盟夺位战对战记录
func (c *Client) SendUnionFightRecordReq(ctx context.Context, req *pb.UnionFightRecordReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206712, data)
}

// 妖盟夺位战历史最高
func (c *Client) SendUnionFightGetHistoryDataReq(ctx context.Context, req *pb.UnionFightGetHistoryDataReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206713, data)
}

// 妖盟夺位战获取锁定属性
func (c *Client) SendUnionFightGetLockedDetailReq(ctx context.Context, req *pb.UnionFightGetLockedDetailReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206714, data)
}

// 妖盟夺位战设置锁定状态
func (c *Client) SendUnionFightChangeLockStatusReq(ctx context.Context, req *pb.UnionFightChangeLockStatusReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206715, data)
}

// 妖盟夺位战今日详情
func (c *Client) SendUnionFightGetTodayResultReq(ctx context.Context, req *pb.UnionFightGetTodayResultReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206716, data)
}

// 获取妖盟排行列表
func (c *Client) SendUnionFightGetUnionRankListReq(ctx context.Context, req *pb.UnionFightGetUnionRankListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206717, data)
}

// 基础信息
func (c *Client) SendGodIslandBaseInfoReqMsg(ctx context.Context, req *pb.GodIslandBaseInfoReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206801, data)
}

// 查看分组
func (c *Client) SendGodIslandGroupReqMsg(ctx context.Context, req *pb.GodIslandGroupReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206802, data)
}

// 请求妖盟建筑战报
func (c *Client) SendGodIslandReportReqMsg(ctx context.Context, req *pb.GodIslandReportReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206803, data)
}

// 请求妖盟建筑战报详情数据
func (c *Client) SendGodIslandReportDetailReqMsg(ctx context.Context, req *pb.GodIslandReportDetailReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206804, data)
}

// 请求联盟战功数据
func (c *Client) SendGodIslandUnionDamageReqMsg(ctx context.Context, req *pb.GodIslandUnionDamageReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206805, data)
}

// 查看自己所属联盟的参与名单
func (c *Client) SendGodIslandGetJoinMemberListReq(ctx context.Context, req *pb.GodIslandGetJoinMemberListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206806, data)
}

// 妖盟分段位伤害信息
func (c *Client) SendGodIslandUnionGroupDamageReqMsg(ctx context.Context, req *pb.GodIslandUnionGroupDamageReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206807, data)
}

// 红点信息
func (c *Client) SendGodIslandRedDotReqMsg(ctx context.Context, req *pb.GodIslandRedDotReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206808, data)
}

// 点赞
func (c *Client) SendGodIslandWorshipReq(ctx context.Context, req *pb.GodIslandWorshipReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206809, data)
}

// 请求个人战报
func (c *Client) SendGodIslandPlayerReportReqMsg(ctx context.Context, req *pb.GodIslandPlayerReportReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206810, data)
}

// 妖力转灵力
func (c *Client) SendGodIslandUpdatePowerReq(ctx context.Context, req *pb.GodIslandUpdatePowerReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206811, data)
}

// 查看妖盟战功列表
func (c *Client) SendGodIslandUnionBattleScoreListReq(ctx context.Context, req *pb.GodIslandUnionBattleScoreListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206812, data)
}

// 灵液领取记录
func (c *Client) SendGodIslandLiquidReceiveRecordReq(ctx context.Context, req *pb.GodIslandLiquidReceiveRecordReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206813, data)
}

// 妖盟成员战功统计
func (c *Client) SendGodIslandUnionMemberScoreListReq(ctx context.Context, req *pb.GodIslandUnionMemberScoreListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206814, data)
}

// 活动心跳
func (c *Client) SendGodIslandHeartBeatReqMsg(ctx context.Context, req *pb.GodIslandHeartBeatReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206819, data)
}

// 游戏基础信息
func (c *Client) SendGodIslandGameInfoReqMsg(ctx context.Context, req *pb.GodIslandGameInfoReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206820, data)
}

// 游戏大地图城市信息(在滑动地图时候将视野内的城市id发送)
func (c *Client) SendGodIslandGameCityInfoReqMsg(ctx context.Context, req *pb.GodIslandGameCityInfoReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206823, data)
}

// npc事件选择
func (c *Client) SendGodIslandGameEventReqMsg(ctx context.Context, req *pb.GodIslandGameEventReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206824, data)
}

// 灵力球信息
func (c *Client) SendGodIslandGameSpiritualBallInfoReqMsg(ctx context.Context, req *pb.GodIslandGameSpiritualBallInfoReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206826, data)
}

// 使用灵力球
func (c *Client) SendGodIslandGameUseSpiritualBallReqMsg(ctx context.Context, req *pb.GodIslandGameUseSpiritualBallReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206827, data)
}

// 目标城市具体信息
func (c *Client) SendGodIslandGameTargetCityInfoReqMsg(ctx context.Context, req *pb.GodIslandGameTargetCityInfoReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206828, data)
}

// 目标城市队列信息
func (c *Client) SendGodIslandGameTargetCityLineInfoReqMsg(ctx context.Context, req *pb.GodIslandGameTargetCityLineInfoReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206829, data)
}

// 进攻
func (c *Client) SendGodIslandGameAttackReqMsg(ctx context.Context, req *pb.GodIslandGameAttackReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206830, data)
}

// 自动操作 托管
func (c *Client) SendGodIslandGameAutoAttackReqMsg(ctx context.Context, req *pb.GodIslandGameAutoAttackReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206831, data)
}

// 行军移动
func (c *Client) SendGodIslandGameMoveReqMsg(ctx context.Context, req *pb.GodIslandGameMoveReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206832, data)
}

// 小地图信息
func (c *Client) SendGodIslandGameMiniMapInfoReqMsg(ctx context.Context, req *pb.GodIslandGameMiniMapInfoReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206834, data)
}

// 设置联盟类型
func (c *Client) SendGodIslandGameSetUnionTypeReqMsg(ctx context.Context, req *pb.GodIslandGameSetUnionTypeReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206835, data)
}

// 联盟种草信息
func (c *Client) SendGodIslandGamePlantInfoReqMsg(ctx context.Context, req *pb.GodIslandGamePlantInfoReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206836, data)
}

// 使用仙果 补充灵力球的灵力
func (c *Client) SendGodIslandUseFruitReq(ctx context.Context, req *pb.GodIslandUseFruitReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206837, data)
}

// 联盟种草领取果实
func (c *Client) SendGodIslandGamePlantReceiveFruitsReqMsg(ctx context.Context, req *pb.GodIslandGamePlantReceiveFruitsReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206838, data)
}

// 联盟种草浇水
func (c *Client) SendGodIslandGamePlantWaterReqMsg(ctx context.Context, req *pb.GodIslandGamePlantWaterReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206839, data)
}

// 灵晶数据加载
func (c *Client) SendGodIslandCrystalInfoReqMsg(ctx context.Context, req *pb.GodIslandCrystalInfoReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206840, data)
}

// 领取灵晶
func (c *Client) SendGodIslandCrystalReceiveReqMsg(ctx context.Context, req *pb.GodIslandCrystalReceiveReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206841, data)
}

// 联盟指挥信息
func (c *Client) SendGodIslandGameStrategyInfoReqMsg(ctx context.Context, req *pb.GodIslandGameStrategyInfoReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206842, data)
}

// 联盟指挥设置
func (c *Client) SendGodIslandGameCommanderSetReqMsg(ctx context.Context, req *pb.GodIslandGameCommanderSetReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206843, data)
}

// 联盟指挥设置目标
func (c *Client) SendGodIslandGameSetUnionTargetReqMsg(ctx context.Context, req *pb.GodIslandGameSetUnionTargetReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206844, data)
}

// 联盟指挥请求支援
func (c *Client) SendGodIslandGameForHelpReqMsg(ctx context.Context, req *pb.GodIslandGameForHelpReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206845, data)
}

// 加速行军
func (c *Client) SendGodIslandGameAcclerateMoveMsg(ctx context.Context, req *pb.GodIslandGameAcclerateMoveMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206849, data)
}

// 联盟发送灵液
func (c *Client) SendGodIslandGameSendLiquidReqMsg(ctx context.Context, req *pb.GodIslandGameSendLiquidReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206851, data)
}

// 联盟灵液发送记录
func (c *Client) SendGodIslandGameSendLiquidRecordReqMsg(ctx context.Context, req *pb.GodIslandGameSendLiquidRecordReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206852, data)
}

// 请求城市界面buff信息
func (c *Client) SendGodIslandGameCityBuffInfoReqMsg(ctx context.Context, req *pb.GodIslandGameCityBuffInfoReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206855, data)
}

// 阵亡之后 补兵完成之后从场外回大本营
func (c *Client) SendGodIslandBackHomeReq(ctx context.Context, req *pb.GodIslandBackHomeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206856, data)
}

// 修改盟友
func (c *Client) SendGodIslandSetFriendUnionReq(ctx context.Context, req *pb.GodIslandSetFriendUnionReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206858, data)
}

// 星宿试炼战斗
func (c *Client) SendStarTrialChallengeReq(ctx context.Context, req *pb.StarTrialChallengeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206902, data)
}

// 获取日志信息
func (c *Client) SendStarTrialRecordReq(ctx context.Context, req *pb.StarTrialRecordReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206903, data)
}

// 战斗回放
func (c *Client) SendStarTrialBattleReplyReq(ctx context.Context, req *pb.StarTrialBattleReplyReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206904, data)
}

// 进入星空图鉴信息
func (c *Client) SendEnterStarTrialCodexMsgReq(ctx context.Context, req *pb.EnterStarTrialCodexMsgReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206905, data)
}

// 获取详细图鉴击败信息
func (c *Client) SendPlayerStarTrialCodexMsgReq(ctx context.Context, req *pb.PlayerStarTrialCodexMsgReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206906, data)
}

// 获取Boss属性信息
func (c *Client) SendGetBossDetailDataReq(ctx context.Context, req *pb.GetBossDetailDataReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206907, data)
}

// 每日领奖
func (c *Client) SendGetDailyFightRewardReq(ctx context.Context, req *pb.GetDailyFightRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206908, data)
}

// 获取分组信息
func (c *Client) SendGetStarTrialGroupInfoReq(ctx context.Context, req *pb.GetStarTrialGroupInfoReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206909, data)
}

// 查看玩家详情
func (c *Client) SendStarTrialGetPlayerDetailReq(ctx context.Context, req *pb.StarTrialGetPlayerDetailReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206910, data)
}

// 进入星宿试炼
func (c *Client) SendStarTrialEnterMainPanelReq(ctx context.Context, req *pb.StarTrialEnterMainPanelReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 206911, data)
}

// 进入聚灵阵
func (c *Client) SendGatherEnergyEnterNewReq(ctx context.Context, req *pb.GatherEnergyEnterNewReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 207001, data)
}

// 开启灵阵界面
func (c *Client) SendGatherEnergyOpenViewReq(ctx context.Context, req *pb.GatherEnergyOpenViewReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 207003, data)
}

// 开启某个灵阵
func (c *Client) SendGatherEnergyOpenReq(ctx context.Context, req *pb.GatherEnergyOpenReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 207004, data)
}

// 挑战某人
func (c *Client) SendGatherEnergyFightReq(ctx context.Context, req *pb.GatherEnergyFightReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 207007, data)
}

// 聚灵阵奖励领取
func (c *Client) SendGatherEnergyRewardReq(ctx context.Context, req *pb.GatherEnergyRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 207011, data)
}

// 聚灵阵排行榜点赞
func (c *Client) SendGatherEnergyLikeReq(ctx context.Context, req *pb.GatherEnergyLikeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 207012, data)
}

// 灵气转换
func (c *Client) SendGatherEnergyTransformReq(ctx context.Context, req *pb.GatherEnergyTransformReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 207015, data)
}

// 离开聚灵阵
func (c *Client) SendGatherEnergyLeaveReq(ctx context.Context, req *pb.GatherEnergyLeaveReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 207017, data)
}

// 聚灵阵公告
func (c *Client) SendGatherEnergyNoticeReq(ctx context.Context, req *pb.GatherEnergyNoticeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 207018, data)
}

// 聚灵阵获取广告奖励
func (c *Client) SendGatherEnergyGetADAwardReq(ctx context.Context, req *pb.GatherEnergyGetADAwardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 207019, data)
}

// 打开聚灵阵一级列表界面（参与人数）
func (c *Client) SendGatherEnergyFirstListViewReq(ctx context.Context, req *pb.GatherEnergyFirstListViewReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 207020, data)
}

// 打开聚灵阵二级列表界面（参与名字）
func (c *Client) SendGatherEnergySecondListViewReq(ctx context.Context, req *pb.GatherEnergySecondListViewReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 207021, data)
}

// 打开某个聚灵阵内部
func (c *Client) SendGatherEnergyInsideViewNewReq(ctx context.Context, req *pb.GatherEnergyInsideViewNewReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 207022, data)
}

// 打开战报界面
func (c *Client) SendGatherEnergyFightReportNewReq(ctx context.Context, req *pb.GatherEnergyFightReportNewReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 207023, data)
}

// 查询某人参与的阵法
func (c *Client) SendGatherEnergySearchNewReq(ctx context.Context, req *pb.GatherEnergySearchNewReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 207024, data)
}

// 加入聚灵阵
func (c *Client) SendGatherEnergyAttendNewReq(ctx context.Context, req *pb.GatherEnergyAttendNewReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 207025, data)
}

// 领取sdk奖励
func (c *Client) SendReceiveSdkRewardReq(ctx context.Context, req *pb.ReceiveSdkRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 207202, data)
}

// 领取qq卡券奖励
func (c *Client) SendQQCardGetRewardReq(ctx context.Context, req *pb.QQCardGetRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 207203, data)
}

// 支付宝启动参数
func (c *Client) SendAlipayStartParamReq(ctx context.Context, req *pb.AlipayStartParamReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 207204, data)
}

// Sdk七日登录奖励领取
func (c *Client) SendReceiveSdkDailyRewardReq(ctx context.Context, req *pb.ReceiveSdkDailyRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 207205, data)
}

// 美团启动参数
func (c *Client) SendMeiTuanStartParamReq(ctx context.Context, req *pb.MeiTuanStartParamReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 207206, data)
}

// 进入征战诸天
func (c *Client) SendSkyWarEnterReq(ctx context.Context, req *pb.SkyWarEnterReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 208401, data)
}

// 刷新对手玩家
func (c *Client) SendSkyWarRefreshEnemyReq(ctx context.Context, req *pb.SkyWarRefreshEnemyReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 208402, data)
}

// 斗法
func (c *Client) SendSkyWarFightReq(ctx context.Context, req *pb.SkyWarFightReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 208403, data)
}

// 获取排行榜数据
func (c *Client) SendSkyWarRankReq(ctx context.Context, req *pb.SkyWarRankReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 208404, data)
}

// 获取日志
func (c *Client) SendSkyWarLogReq(ctx context.Context, req *pb.SkyWarLogReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 208405, data)
}

// 战斗回放
func (c *Client) SendSkyWarLogPlaybackReq(ctx context.Context, req *pb.SkyWarLogPlaybackReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 208406, data)
}

// 获取布阵详情
func (c *Client) SendSkyWarFormationReq(ctx context.Context, req *pb.SkyWarFormationReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 208407, data)
}

// 设置顺位
func (c *Client) SendSkyWarSetOrderReq(ctx context.Context, req *pb.SkyWarSetOrderReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 208408, data)
}

// 诸天榜
func (c *Client) SendSkyWarSkyRankReq(ctx context.Context, req *pb.SkyWarSkyRankReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 208409, data)
}

// 膜拜
func (c *Client) SendSkyWarWorshipReq(ctx context.Context, req *pb.SkyWarWorshipReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 208410, data)
}

// 购买次数(斗法、刷新)
func (c *Client) SendSkyWarBuyTimesReq(ctx context.Context, req *pb.SkyWarBuyTimesReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 208411, data)
}

// 妖盟对决-同步妖盟对决数据
func (c *Client) SendUnionDuelSyncMsgReq(ctx context.Context, req *pb.UnionDuelSyncMsgReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 208501, data)
}

// 妖盟对决-开启主界面
func (c *Client) SendUnionDuelOpenViewReq(ctx context.Context, req *pb.UnionDuelOpenViewReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 208502, data)
}

// 妖盟对决-妖盟对决同步数据
func (c *Client) SendUnionDuelSyncDataReq(ctx context.Context, req *pb.UnionDuelSyncDataReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 208511, data)
}

// 评分完成
func (c *Client) SendMarkFinishReq(ctx context.Context, req *pb.MarkFinishReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 208602, data)
}

// 法则感悟
func (c *Client) SendWorldRulePerceptionReq(ctx context.Context, req *pb.WorldRulePerceptionReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209001, data)
}

// 法则感悟替换
func (c *Client) SendWorldRulePerceptionReplaceReq(ctx context.Context, req *pb.WorldRulePerceptionReplaceReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209002, data)
}

// 获取法则感悟记录
func (c *Client) SendWorldRuleGetPerceptionLogReq(ctx context.Context, req *pb.WorldRuleGetPerceptionLogReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209003, data)
}

// 切换法则方案
func (c *Client) SendWorldRuleSwitchProgrammeReq(ctx context.Context, req *pb.WorldRuleSwitchProgrammeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209004, data)
}

// 请求法则试炼挑战
func (c *Client) SendRuleTrialChallengeReq(ctx context.Context, req *pb.RuleTrialChallengeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209101, data)
}

// 请求法则试炼一键扫荡
func (c *Client) SendRuleTrialRepeatAllReq(ctx context.Context, req *pb.RuleTrialRepeatAllReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209103, data)
}

// 请求法则试炼Boss数据
func (c *Client) SendRuleTrialMonsterAttrReq(ctx context.Context, req *pb.RuleTrialMonsterAttrReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209106, data)
}

// 缩时香礼包
func (c *Client) SendADTimeTriggerReq(ctx context.Context, req *pb.ADTimeTriggerReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209201, data)
}

// 使用缩时香
func (c *Client) SendADTimeUseReq(ctx context.Context, req *pb.ADTimeUseReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209202, data)
}

// 九幽争霸数据登录同步
func (c *Client) SendHolyLandBattleApplyDataSyncReq(ctx context.Context, req *pb.HolyLandBattleApplyDataSyncReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209570, data)
}

// 补签接口
func (c *Client) SendReqSignInFundRepair(ctx context.Context, req *pb.ReqSignInFundRepair) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209601, data)
}

// 进入山海界面
func (c *Client) SendEnterMountainSeaReq(ctx context.Context, req *pb.EnterMountainSeaReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209701, data)
}

// 组队内部请求
func (c *Client) SendEnterMountainSeaTeamReq(ctx context.Context, req *pb.EnterMountainSeaTeamReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209702, data)
}

// 组队页面请求
func (c *Client) SendMountainSeaTeamStartReq(ctx context.Context, req *pb.MountainSeaTeamStartReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209703, data)
}

// 创建队伍
func (c *Client) SendMountainSeaCreateTeamReq(ctx context.Context, req *pb.MountainSeaCreateTeamReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209704, data)
}

// 获取队伍列表
func (c *Client) SendMountainSeaGetTeamListReq(ctx context.Context, req *pb.MountainSeaGetTeamListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209705, data)
}

// 查找队伍
func (c *Client) SendMountainSeaGetTeamInfoReq(ctx context.Context, req *pb.MountainSeaGetTeamInfoReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209706, data)
}

// 加入队伍
func (c *Client) SendMountainSeaJoinTeamReq(ctx context.Context, req *pb.MountainSeaJoinTeamReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209707, data)
}

// 取消加入队伍
func (c *Client) SendMountainSeaCancelTeamApplyReq(ctx context.Context, req *pb.MountainSeaCancelTeamApplyReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209708, data)
}

// 同意申请
func (c *Client) SendMountainSeaApplyJoinTeamAgreeReq(ctx context.Context, req *pb.MountainSeaApplyJoinTeamAgreeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209709, data)
}

// 一键拒绝
func (c *Client) SendMountainSeaApplyJoinTeamRefusedReq(ctx context.Context, req *pb.MountainSeaApplyJoinTeamRefusedReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209710, data)
}

// 退出队伍
func (c *Client) SendMountainSeaQuitTeamReq(ctx context.Context, req *pb.MountainSeaQuitTeamReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209711, data)
}

// 踢出队伍
func (c *Client) SendMountainSeaKickOutTeamReq(ctx context.Context, req *pb.MountainSeaKickOutTeamReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209712, data)
}

// 转让队长
func (c *Client) SendMountainSeaChangeLeaderReq(ctx context.Context, req *pb.MountainSeaChangeLeaderReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209713, data)
}

// 开始准备
func (c *Client) SendMountainSeaTeamPrepareReq(ctx context.Context, req *pb.MountainSeaTeamPrepareReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209714, data)
}

// 匹配
func (c *Client) SendMountainSeaMatchMemberReq(ctx context.Context, req *pb.MountainSeaMatchMemberReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209715, data)
}

// 开始战斗
func (c *Client) SendMountainSeaStartBattleReq(ctx context.Context, req *pb.MountainSeaStartBattleReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209716, data)
}

// 请求玩家数据
func (c *Client) SendMountainSeaGetPlayerInfoReq(ctx context.Context, req *pb.MountainSeaGetPlayerInfoReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209717, data)
}

// 膜拜
func (c *Client) SendMountainSeaWorshipReq(ctx context.Context, req *pb.MountainSeaWorshipReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209718, data)
}

// 战斗回放
func (c *Client) SendMountainSeaGetBattleVideoReq(ctx context.Context, req *pb.MountainSeaGetBattleVideoReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209719, data)
}

// 获取boss信息
func (c *Client) SendMountainSeaGetBossInfoReq(ctx context.Context, req *pb.MountainSeaGetBossInfoReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209720, data)
}

// 请求挑战次数
func (c *Client) SendMountainSeaChallengeTimeReq(ctx context.Context, req *pb.MountainSeaChallengeTimeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209721, data)
}

// 开始匹配
func (c *Client) SendMountainSeaStartMatchReq(ctx context.Context, req *pb.MountainSeaStartMatchReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209722, data)
}

// 邀请
func (c *Client) SendMountainSeaInviteReq(ctx context.Context, req *pb.MountainSeaInviteReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209723, data)
}

// 离开系统
func (c *Client) SendLeaveMountainSeaReq(ctx context.Context, req *pb.LeaveMountainSeaReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209724, data)
}

// 获取boss妖力
func (c *Client) SendMountainSeaGetBossPowerReq(ctx context.Context, req *pb.MountainSeaGetBossPowerReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209725, data)
}

// 进入布阵界面
func (c *Client) SendMountainSeaEnterBattleReq(ctx context.Context, req *pb.MountainSeaEnterBattleReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209750, data)
}

// 进入切换分身界面
func (c *Client) SendMountainSeaEnterSwitchSeparationReq(ctx context.Context, req *pb.MountainSeaEnterSwitchSeparationReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209752, data)
}

// 切换上阵的分身
func (c *Client) SendMountainSeaSwitchSeparationReq(ctx context.Context, req *pb.MountainSeaSwitchSeparationReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209753, data)
}

// 查看分身详细数据
func (c *Client) SendMountainSeaSeparationDetailReq(ctx context.Context, req *pb.MountainSeaSeparationDetailReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209754, data)
}

// 改变参战对象位置
func (c *Client) SendMountainSeaChangePosReq(ctx context.Context, req *pb.MountainSeaChangePosReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209755, data)
}

// 改变装配的队伍技能
func (c *Client) SendMountainSeaChangeTeamSkillReq(ctx context.Context, req *pb.MountainSeaChangeTeamSkillReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209756, data)
}

// 进行战斗
func (c *Client) SendMountainSeaDoBattleReq(ctx context.Context, req *pb.MountainSeaDoBattleReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209757, data)
}

// 邀请列表
func (c *Client) SendMountainSeaInviteListReq(ctx context.Context, req *pb.MountainSeaInviteListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209783, data)
}

// 一键拒绝邀请队伍
func (c *Client) SendMountainSeaInviteRefuseReq(ctx context.Context, req *pb.MountainSeaInviteRefuseReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209785, data)
}

// 设置接受指定邀请开关
func (c *Client) SendMountainSeaSetAppointReq(ctx context.Context, req *pb.MountainSeaSetAppointReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209786, data)
}

// 加入指定邀请队伍
func (c *Client) SendMountainSeaApplyJoinTeamReq(ctx context.Context, req *pb.MountainSeaApplyJoinTeamReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209787, data)
}

// 疯狂聚宝盆签到
func (c *Client) SendReqGetConditionReward(ctx context.Context, req *pb.ReqGetConditionReward) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 209901, data)
}

// 仙友回归-首次分享领奖
func (c *Client) SendRegressionShareReq(ctx context.Context, req *pb.RegressionShareReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 210001, data)
}

// 仙友回归-抽奖
func (c *Client) SendRegressionLotteryReq(ctx context.Context, req *pb.RegressionLotteryReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 210002, data)
}

// 仙友回归-绑定回归玩家数据
func (c *Client) SendGetRegressionPlayerDataReq(ctx context.Context, req *pb.GetRegressionPlayerDataReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 210003, data)
}

// 仙友回归-领取绑定玩家的任务奖励
func (c *Client) SendGetRegressionReceiveRewardReq(ctx context.Context, req *pb.GetRegressionReceiveRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 210004, data)
}

// 仙友回归-保存自选奖励
func (c *Client) SendRegressionSaveSelectItemReq(ctx context.Context, req *pb.RegressionSaveSelectItemReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 210005, data)
}

// 妖盟赐福
func (c *Client) SendUnionBlessingSendGiftReq(ctx context.Context, req *pb.UnionBlessingSendGiftReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 210601, data)
}

// 领取赐福奖励
func (c *Client) SendUnionBlessingRewardReqMsg(ctx context.Context, req *pb.UnionBlessingRewardReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 210602, data)
}

// 获取签到奖励
func (c *Client) SendGoodFortuneGetRewardReq(ctx context.Context, req *pb.GoodFortuneGetRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 210701, data)
}

// 通关请求
func (c *Client) SendWestTravelPassGameReq(ctx context.Context, req *pb.WestTravelPassGameReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 210901, data)
}

// 使用宴会道具
func (c *Client) SendFestivalCelebrationsUseBanquetItemReq(ctx context.Context, req *pb.FestivalCelebrationsUseBanquetItemReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211001, data)
}

// 领取宴会奖励
func (c *Client) SendFestivalCelebrationsDrawBanquetRewardReq(ctx context.Context, req *pb.FestivalCelebrationsDrawBanquetRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211002, data)
}

// 获取积分详情
func (c *Client) SendFestivalCelebrationsGetBanquetScoreDetailReq(ctx context.Context, req *pb.FestivalCelebrationsGetBanquetScoreDetailReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211003, data)
}

// 使用福缘道具
func (c *Client) SendFestivalCelebrationsUseLuckyFateItemReq(ctx context.Context, req *pb.FestivalCelebrationsUseLuckyFateItemReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211004, data)
}

// 获取签到奖励
func (c *Client) SendFestivalCelebrationsGetSignRewardReq(ctx context.Context, req *pb.FestivalCelebrationsGetSignRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211005, data)
}

// 请求合成
func (c *Client) SendFestivalCelebrationsCollectSynthesisReq(ctx context.Context, req *pb.FestivalCelebrationsCollectSynthesisReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211006, data)
}

// 请求填充
func (c *Client) SendFestivalCelebrationsCollectFillReq(ctx context.Context, req *pb.FestivalCelebrationsCollectFillReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211007, data)
}

// 领取大奖
func (c *Client) SendFestivalCelebrationsCollectDrawBigRewardReq(ctx context.Context, req *pb.FestivalCelebrationsCollectDrawBigRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211008, data)
}

// 获取领取名单
func (c *Client) SendFestivalCelebrationsCollectGetClaimListReq(ctx context.Context, req *pb.FestivalCelebrationsCollectGetClaimListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211009, data)
}

// 搜索玩家
func (c *Client) SendFestivalCelebrationsCollectSearchPlayerReq(ctx context.Context, req *pb.FestivalCelebrationsCollectSearchPlayerReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211010, data)
}

// 赠送道具
func (c *Client) SendFestivalCelebrationsCollectGiveReq(ctx context.Context, req *pb.FestivalCelebrationsCollectGiveReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211011, data)
}

// 索要道具
func (c *Client) SendFestivalCelebrationsCollectAskForReq(ctx context.Context, req *pb.FestivalCelebrationsCollectAskForReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211012, data)
}

// 领取彩蛋奖励
func (c *Client) SendFestivalCelebrationsEasterEggGetRewardReq(ctx context.Context, req *pb.FestivalCelebrationsEasterEggGetRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211013, data)
}

// 领取赠送五福
func (c *Client) SendFestivalCelebrationsCollectGetGivenReq(ctx context.Context, req *pb.FestivalCelebrationsCollectGetGivenReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211014, data)
}

// 领取福源奖励
func (c *Client) SendFestivalCelebrationsFuYuanGetRewardReq(ctx context.Context, req *pb.FestivalCelebrationsFuYuanGetRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211015, data)
}

// 获取绑定妖盟成员列表
func (c *Client) SendFestivalCelebrationsGetBindUnionMemberDataListReq(ctx context.Context, req *pb.FestivalCelebrationsGetBindUnionMemberDataListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211016, data)
}

// 获取周年回忆条目
func (c *Client) SendFestivalCelebrationsGetAnnualMemoryReq(ctx context.Context, req *pb.FestivalCelebrationsGetAnnualMemoryReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211021, data)
}

// 膜拜
func (c *Client) SendFestivalCelebrationsWorshipReq(ctx context.Context, req *pb.FestivalCelebrationsWorshipReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211022, data)
}

// 指定邀请
func (c *Client) SendDoubleDemonsAppointInviteReq(ctx context.Context, req *pb.DoubleDemonsAppointInviteReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211106, data)
}

// 同意邀请/结为伴侣
func (c *Client) SendDoubleDemonsAgreeReq(ctx context.Context, req *pb.DoubleDemonsAgreeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211112, data)
}

// 一键拒绝
func (c *Client) SendDoubleDemonsRefuseReq(ctx context.Context, req *pb.DoubleDemonsRefuseReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211113, data)
}

// 领取赠送奖励
func (c *Client) SendDoubleDemonsReceivePresentReq(ctx context.Context, req *pb.DoubleDemonsReceivePresentReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211114, data)
}

// 新春红包 进入系统
func (c *Client) SendNewYearRedBagEnterReqMsg(ctx context.Context, req *pb.NewYearRedBagEnterReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211501, data)
}

// 新春红包 退出系统
func (c *Client) SendNewYearRedBagExitReqMsg(ctx context.Context, req *pb.NewYearRedBagExitReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211502, data)
}

// 新春红包 打开红包
func (c *Client) SendNewYearRedBagOpenReqMsg(ctx context.Context, req *pb.NewYearRedBagOpenReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211503, data)
}

// 穿戴内丹
func (c *Client) SendPetKerneCarryReq(ctx context.Context, req *pb.PetKerneCarryReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211702, data)
}

// 内丹升级
func (c *Client) SendPetKerneUpLevelReq(ctx context.Context, req *pb.PetKerneUpLevelReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211703, data)
}

// 内丹升星
func (c *Client) SendPetKerneUpStarReq(ctx context.Context, req *pb.PetKerneUpStarReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211704, data)
}

// 内丹激活
func (c *Client) SendPetKerneActiveReq(ctx context.Context, req *pb.PetKerneActiveReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211705, data)
}

// 抽取灵兽内丹
func (c *Client) SendPetKernelDrawReq(ctx context.Context, req *pb.PetKernelDrawReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211706, data)
}

// 内丹共鸣升级或激活
func (c *Client) SendPetKerneCombineUpLevelReq(ctx context.Context, req *pb.PetKerneCombineUpLevelReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211707, data)
}

// 进入门徒系统
func (c *Client) SendEnterPupilSystemReq(ctx context.Context, req *pb.EnterPupilSystemReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211801, data)
}

// 门徒培养
func (c *Client) SendPupilTrainReq(ctx context.Context, req *pb.PupilTrainReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211802, data)
}

// 恢复培养次数
func (c *Client) SendPupilSiteTrainTimesRecoverReq(ctx context.Context, req *pb.PupilSiteTrainTimesRecoverReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211803, data)
}

// 招收门徒
func (c *Client) SendPupilRecruitReq(ctx context.Context, req *pb.PupilRecruitReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211805, data)
}

// 门徒出师
func (c *Client) SendPupilGraduateReq(ctx context.Context, req *pb.PupilGraduateReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211806, data)
}

// 获取出师门徒列表
func (c *Client) SendPupilGetGraduatedListReq(ctx context.Context, req *pb.PupilGetGraduatedListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211807, data)
}

// 获取上阵列表
func (c *Client) SendPupilGetFightListReq(ctx context.Context, req *pb.PupilGetFightListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211808, data)
}

// 门徒上阵
func (c *Client) SendPupilFightOnReq(ctx context.Context, req *pb.PupilFightOnReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211809, data)
}

// 获取门派统计信息
func (c *Client) SendPupilGetSectInfoReq(ctx context.Context, req *pb.PupilGetSectInfoReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211810, data)
}

// 搜索玩家
func (c *Client) SendSearchPlayerReq(ctx context.Context, req *pb.SearchPlayerReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211811, data)
}

// 领取广告奖励
func (c *Client) SendPupilGetAdRewardReq(ctx context.Context, req *pb.PupilGetAdRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211814, data)
}

// 锁定弟子
func (c *Client) SendPupilLockReq(ctx context.Context, req *pb.PupilLockReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211815, data)
}

// 精准查找
func (c *Client) SendPupilExactSearchReq(ctx context.Context, req *pb.PupilExactSearchReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211816, data)
}

// 获得出师未结义的门徒列表
func (c *Client) SendPupilGraduatedUnMarryListReq(ctx context.Context, req *pb.PupilGraduatedUnMarryListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211855, data)
}

// 结义记录列表
func (c *Client) SendMarriageRecordListReq(ctx context.Context, req *pb.MarriageRecordListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211856, data)
}

// 获取全服结义请求列表(包括妖盟)
func (c *Client) SendMarriageGetServerApplyReq(ctx context.Context, req *pb.MarriageGetServerApplyReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211857, data)
}

// 指定结义推荐玩家请求
func (c *Client) SendMarriageRecommendPlayerReq(ctx context.Context, req *pb.MarriageRecommendPlayerReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211858, data)
}

// 获取其他玩家发起的结义请求（别的玩家指定你结义的列表）
func (c *Client) SendMarriageGetAppointApplyReq(ctx context.Context, req *pb.MarriageGetAppointApplyReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211859, data)
}

// 处理结义请求
func (c *Client) SendMarriageApplyDealReq(ctx context.Context, req *pb.MarriageApplyDealReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211860, data)
}

// 发布结义请求
func (c *Client) SendMarriageApplyPublishReq(ctx context.Context, req *pb.MarriageApplyPublishReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211861, data)
}

// 取消结义请求
func (c *Client) SendMarriageApplyCancelReq(ctx context.Context, req *pb.MarriageApplyCancelReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211862, data)
}

// 刷新全服结义列表请求
func (c *Client) SendMarriageRefreshServerApplyReq(ctx context.Context, req *pb.MarriageRefreshServerApplyReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211863, data)
}

// 设置是否接受指向结义请求
func (c *Client) SendMarriageSetAppointReq(ctx context.Context, req *pb.MarriageSetAppointReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 211864, data)
}

// 洪荒灵兽灵根升级请求
func (c *Client) SendPetRootUpReq(ctx context.Context, req *pb.PetRootUpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 212201, data)
}

// 洪荒灵兽觉醒请求
func (c *Client) SendPetAwakeReq(ctx context.Context, req *pb.PetAwakeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 212202, data)
}

// 洪荒灵兽缘分升级请求
func (c *Client) SendPetFateUpReq(ctx context.Context, req *pb.PetFateUpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 212203, data)
}

// 切换灵兽皮肤
func (c *Client) SendPetSwitchSkinReq(ctx context.Context, req *pb.PetSwitchSkinReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 212204, data)
}

// 激活、升级灵兽皮肤
func (c *Client) SendPetUpSkinLevelReq(ctx context.Context, req *pb.PetUpSkinLevelReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 212205, data)
}

// 激活、升级灵兽羁绊
func (c *Client) SendPetUpSkinCombineReq(ctx context.Context, req *pb.PetUpSkinCombineReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 212206, data)
}

// 获取微信排行榜白名单
func (c *Client) SendGetWechatRankWhiteListReq(ctx context.Context, req *pb.GetWechatRankWhiteListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 212601, data)
}

// 筋斗云注灵请求
func (c *Client) SendCloudRefineReq(ctx context.Context, req *pb.CloudRefineReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 212901, data)
}

// 筋斗云注灵渡劫请求
func (c *Client) SendCloudRefineStarLvUpReq(ctx context.Context, req *pb.CloudRefineStarLvUpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 212902, data)
}

// 进入游戏
func (c *Client) SendComposeBallEnterGameReq(ctx context.Context, req *pb.ComposeBallEnterGameReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213001, data)
}

// 请求合成
func (c *Client) SendComposeBallComposeReq(ctx context.Context, req *pb.ComposeBallComposeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213002, data)
}

// 请求使用道具
func (c *Client) SendComposeBallUseItemReq(ctx context.Context, req *pb.ComposeBallUseItemReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213003, data)
}

// 请求增加体力
func (c *Client) SendComposeBallHpItemReq(ctx context.Context, req *pb.ComposeBallHpItemReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213004, data)
}

// 请求结算
func (c *Client) SendComposeBallSettleReq(ctx context.Context, req *pb.ComposeBallSettleReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213005, data)
}

// 进入活动
func (c *Client) SendMonopolyEnterActivityReq(ctx context.Context, req *pb.MonopolyEnterActivityReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213101, data)
}

// 进入地图
func (c *Client) SendMonopolyEnterMapReq(ctx context.Context, req *pb.MonopolyEnterMapReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213102, data)
}

// 摇色子
func (c *Client) SendMonopolyRollDiceReq(ctx context.Context, req *pb.MonopolyRollDiceReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213103, data)
}

// 获取协助列表
func (c *Client) SendMonopolyAssistListReq(ctx context.Context, req *pb.MonopolyAssistListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213104, data)
}

// 使用道具 补充体力
func (c *Client) SendMonopolyReplenishStrengthReq(ctx context.Context, req *pb.MonopolyReplenishStrengthReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213105, data)
}

// 协助队友出陷阱
func (c *Client) SendMonopolyRescueTrapReq(ctx context.Context, req *pb.MonopolyRescueTrapReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213107, data)
}

// 获取可掠夺的妖盟列表
func (c *Client) SendMonopolyRobListReq(ctx context.Context, req *pb.MonopolyRobListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213108, data)
}

// 掠夺其他妖盟
func (c *Client) SendMonopolyRobReq(ctx context.Context, req *pb.MonopolyRobReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213110, data)
}

// 获取妖盟日志列表
func (c *Client) SendMonopolyUnionLogListReq(ctx context.Context, req *pb.MonopolyUnionLogListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213111, data)
}

// 获取玩家日志详情
func (c *Client) SendMonopolyPlayerLogDetailReq(ctx context.Context, req *pb.MonopolyPlayerLogDetailReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213112, data)
}

// 协助队友攻击怪物
func (c *Client) SendMonopolyAssistAttackMonsterReq(ctx context.Context, req *pb.MonopolyAssistAttackMonsterReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213113, data)
}

// 领取队友帮忙打死怪后的奖励
func (c *Client) SendMonopolyReceiveMonsterRewardReq(ctx context.Context, req *pb.MonopolyReceiveMonsterRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213114, data)
}

// 红点详情
func (c *Client) SendMonopolyRedPointReqMsg(ctx context.Context, req *pb.MonopolyRedPointReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213115, data)
}

// 遥控骰子
func (c *Client) SendMonopolyRemoteRollDiceReq(ctx context.Context, req *pb.MonopolyRemoteRollDiceReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213116, data)
}

// 进入其他妖盟地图，请求详细数据
func (c *Client) SendMonopolyEnterRobMapReq(ctx context.Context, req *pb.MonopolyEnterRobMapReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213117, data)
}

// 自动操作解锁
func (c *Client) SendMonopolyAutoUnlockReq(ctx context.Context, req *pb.MonopolyAutoUnlockReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213118, data)
}

// 获取天赐福源列表
func (c *Client) SendMonopolyBlessingListReq(ctx context.Context, req *pb.MonopolyBlessingListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213119, data)
}

// 领取天赐祥瑞奖励
func (c *Client) SendMonopolyReceiveBlessingReq(ctx context.Context, req *pb.MonopolyReceiveBlessingReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213120, data)
}

// 神行符移动
func (c *Client) SendMonopolyQuickMoveReq(ctx context.Context, req *pb.MonopolyQuickMoveReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213121, data)
}

// 暗格移动
func (c *Client) SendMonopolyDarkGridMoveReq(ctx context.Context, req *pb.MonopolyDarkGridMoveReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213122, data)
}

// 处理事件同一接口
func (c *Client) SendMonopolyEventHandleReq(ctx context.Context, req *pb.MonopolyEventHandleReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213123, data)
}

// 查看怪物详情
func (c *Client) SendMonopolyMonsterAttrReq(ctx context.Context, req *pb.MonopolyMonsterAttrReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213124, data)
}

// 获取个人排名
func (c *Client) SendMonopolyGetPlayerRankReq(ctx context.Context, req *pb.MonopolyGetPlayerRankReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213125, data)
}

// 获取仇人列表
func (c *Client) SendMonopolyGetEnemyListReq(ctx context.Context, req *pb.MonopolyGetEnemyListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213126, data)
}

// 积分详情
func (c *Client) SendMonopolyScoreDetailReq(ctx context.Context, req *pb.MonopolyScoreDetailReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213127, data)
}

// 建筑积分详情
func (c *Client) SendMonopolyBuildingScoreDetailReq(ctx context.Context, req *pb.MonopolyBuildingScoreDetailReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213128, data)
}

// 分组内积分详情
func (c *Client) SendMonopolyGetGroupDetailReq(ctx context.Context, req *pb.MonopolyGetGroupDetailReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213141, data)
}

// 膜拜
func (c *Client) SendMonopolyWorshipReq(ctx context.Context, req *pb.MonopolyWorshipReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213142, data)
}

// 竞猜正确玩家数据
func (c *Client) SendMonopolyGuessPlayersReq(ctx context.Context, req *pb.MonopolyGuessPlayersReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213145, data)
}

// 请求竞猜界面数据
func (c *Client) SendMonopolyGetGuessDataReq(ctx context.Context, req *pb.MonopolyGetGuessDataReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213146, data)
}

// 请求竞猜选择
func (c *Client) SendMonopolyGuessSelectReq(ctx context.Context, req *pb.MonopolyGuessSelectReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213147, data)
}

// 请求竞猜奖励
func (c *Client) SendMonopolyGuessRewardReq(ctx context.Context, req *pb.MonopolyGuessRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213148, data)
}

// 镇魔登录同步
func (c *Client) SendTownDemonApplyDataSyncReq(ctx context.Context, req *pb.TownDemonApplyDataSyncReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213250, data)
}

// 宴会道具使用
func (c *Client) SendUsePeachBanquetItemReq(ctx context.Context, req *pb.UsePeachBanquetItemReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213520, data)
}

// 妖盟悬赏-打开护送地图
func (c *Client) SendUnionBountyEnterMapReq(ctx context.Context, req *pb.UnionBountyEnterMapReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213602, data)
}

// 妖盟悬赏-离开护送地图
func (c *Client) SendUnionBountyExitMapReq(ctx context.Context, req *pb.UnionBountyExitMapReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213603, data)
}

// 妖盟悬赏-掠夺某个镖车
func (c *Client) SendUnionBountyPlunderReq(ctx context.Context, req *pb.UnionBountyPlunderReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213604, data)
}

// 妖盟悬赏-处理悬赏事件
func (c *Client) SendUnionBountyDealBountyReq(ctx context.Context, req *pb.UnionBountyDealBountyReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213605, data)
}

// 妖盟悬赏-领取护送镖车奖励
func (c *Client) SendUnionBountyGetRewardEscortReq(ctx context.Context, req *pb.UnionBountyGetRewardEscortReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213606, data)
}

// 妖盟悬赏-查找某个玩家的镖车
func (c *Client) SendUnionBountyCheckEscortReq(ctx context.Context, req *pb.UnionBountyCheckEscortReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213608, data)
}

// 妖盟悬赏-打开战报
func (c *Client) SendUnionBountyGetReportReq(ctx context.Context, req *pb.UnionBountyGetReportReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213609, data)
}

// 妖盟悬赏-战报中打开某个玩家的镖车
func (c *Client) SendUnionBountyReportCheckEscortReq(ctx context.Context, req *pb.UnionBountyReportCheckEscortReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213611, data)
}

// 妖盟悬赏-排行榜膜拜
func (c *Client) SendUnionBountyWorshipReq(ctx context.Context, req *pb.UnionBountyWorshipReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213612, data)
}

// 妖盟悬赏-刷新护送地图
func (c *Client) SendUnionBountyRefreshMapReq(ctx context.Context, req *pb.UnionBountyRefreshMapReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213613, data)
}

// 妖盟悬赏-打开妖盟战报
func (c *Client) SendUnionBountyOpenBountyEventReq(ctx context.Context, req *pb.UnionBountyOpenBountyEventReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213614, data)
}

// 妖盟悬赏-反击
func (c *Client) SendUnionBountyRetaliateReq(ctx context.Context, req *pb.UnionBountyRetaliateReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213615, data)
}

// 妖盟悬赏-查看联盟积分
func (c *Client) SendUnionBountyGetMemberScoreReq(ctx context.Context, req *pb.UnionBountyGetMemberScoreReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213616, data)
}

// 妖盟悬赏-是否有正在押镖
func (c *Client) SendUnionBountyHaveEscortReq(ctx context.Context, req *pb.UnionBountyHaveEscortReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213617, data)
}

// 妖盟悬赏-查看镖车详情
func (c *Client) SendUnionBountyOpenEscortCartDetailReq(ctx context.Context, req *pb.UnionBountyOpenEscortCartDetailReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213618, data)
}

// 妖盟悬赏-打开妖兽界面
func (c *Client) SendUnionBountyOpenMonsterReq(ctx context.Context, req *pb.UnionBountyOpenMonsterReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213619, data)
}

// 妖盟悬赏-打妖兽
func (c *Client) SendUnionBountyAttackMonsterReq(ctx context.Context, req *pb.UnionBountyAttackMonsterReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213620, data)
}

// 妖盟悬赏-邀请
func (c *Client) SendUnionBountyAskHelpReq(ctx context.Context, req *pb.UnionBountyAskHelpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213621, data)
}

// 妖盟悬赏-打开邀请界面
func (c *Client) SendUnionBountyOpenAskHelpReq(ctx context.Context, req *pb.UnionBountyOpenAskHelpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213622, data)
}

// 妖盟悬赏-处理邀请（同意加入、拒绝、一键拒绝）
func (c *Client) SendUnionBountyDealAskHelpReq(ctx context.Context, req *pb.UnionBountyDealAskHelpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213623, data)
}

// 妖盟悬赏-队伍管理（加入队伍、离开队伍、踢出队伍）
func (c *Client) SendUnionBountyManageAskHelpReq(ctx context.Context, req *pb.UnionBountyManageAskHelpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213624, data)
}

// 妖盟悬赏-查看当前妖兽的属性
func (c *Client) SendUnionBountyGetMonsterAttributeReq(ctx context.Context, req *pb.UnionBountyGetMonsterAttributeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213626, data)
}

// 妖盟悬赏-变更参战玩家位置
func (c *Client) SendUnionBountyMonsterChangePosReq(ctx context.Context, req *pb.UnionBountyMonsterChangePosReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213627, data)
}

// 获取当前存在的妖兽
func (c *Client) SendUnionBountyGetExistMonsterReq(ctx context.Context, req *pb.UnionBountyGetExistMonsterReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213628, data)
}

// 进入宗门大比
func (c *Client) SendEnterPupilRankActivityReq(ctx context.Context, req *pb.EnterPupilRankActivityReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213801, data)
}

// 获取涨幅记录
func (c *Client) SendGetIncreaseRecordReq(ctx context.Context, req *pb.GetIncreaseRecordReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213802, data)
}

// 查看 排行榜上玩家的 涨幅记录
func (c *Client) SendPupilRankDetailReq(ctx context.Context, req *pb.PupilRankDetailReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 213803, data)
}

// 获取世交列表
func (c *Client) SendFriendGetListReq(ctx context.Context, req *pb.FriendGetListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 214101, data)
}

// 赠送
func (c *Client) SendFriendSendReq(ctx context.Context, req *pb.FriendSendReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 214102, data)
}

// 领取
func (c *Client) SendFriendReceiveReq(ctx context.Context, req *pb.FriendReceiveReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 214103, data)
}

// 一键赠送、领取
func (c *Client) SendFriendReceiveAllReq(ctx context.Context, req *pb.FriendReceiveAllReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 214104, data)
}

// 删除世交
func (c *Client) SendFriendDeleteReq(ctx context.Context, req *pb.FriendDeleteReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 214105, data)
}

// 好友消息列表
func (c *Client) SendFriendChatReqMsg(ctx context.Context, req *pb.FriendChatReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 214120, data)
}

// 飞升仙界-南天门协助
func (c *Client) SendFairyLandSouthDoorHelpReq(ctx context.Context, req *pb.FairyLandSouthDoorHelpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 214208, data)
}

// 小世界升级
func (c *Client) SendUniverseLevelUpReq(ctx context.Context, req *pb.UniverseLevelUpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 214303, data)
}

// 小世界轮盘抽奖
func (c *Client) SendUniverseDrawReq(ctx context.Context, req *pb.UniverseDrawReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 214304, data)
}

// 小世界技能激活
func (c *Client) SendUniverseSkillUnlockReq(ctx context.Context, req *pb.UniverseSkillUnlockReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 214305, data)
}

// 小世界技能上阵
func (c *Client) SendEquipUniverseSkillReq(ctx context.Context, req *pb.EquipUniverseSkillReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 214306, data)
}

// 小世界技能共鸣激活、升级
func (c *Client) SendUniverseSkillCombineLvUpReq(ctx context.Context, req *pb.UniverseSkillCombineLvUpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 214307, data)
}

// 洞察天机抽取
func (c *Client) SendUniverseSkillDrawReq(ctx context.Context, req *pb.UniverseSkillDrawReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 214308, data)
}

// 小世界技能升级
func (c *Client) SendUniverseSkillLvUpReq(ctx context.Context, req *pb.UniverseSkillLvUpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 214309, data)
}

// 天地轮盘二次抽奖
func (c *Client) SendUniverseDrawTwiceReq(ctx context.Context, req *pb.UniverseDrawTwiceReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 214310, data)
}

// 获取玩家结算信息
func (c *Client) SendGodDemonBattleGetPlayerSettlementReq(ctx context.Context, req *pb.GodDemonBattleGetPlayerSettlementReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 214527, data)
}

// 获取阵营结算信息
func (c *Client) SendGodDemonBattleGetCampSettlementReq(ctx context.Context, req *pb.GodDemonBattleGetCampSettlementReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 214528, data)
}

// 画中仙境-检测仙泽
func (c *Client) SendCheckFairyPoolReq(ctx context.Context, req *pb.CheckFairyPoolReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 214698, data)
}

// 进入活动
func (c *Client) SendSkyTradeEnterActivityReq(ctx context.Context, req *pb.SkyTradeEnterActivityReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215000, data)
}

// 进入地图
func (c *Client) SendSkyTradeEnterReq(ctx context.Context, req *pb.SkyTradeEnterReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215001, data)
}

// 查看分组信息
func (c *Client) SendSkyTradeGroupInfoReq(ctx context.Context, req *pb.SkyTradeGroupInfoReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215002, data)
}

// 加速移动
func (c *Client) SendSkyTradeAddSpeedReq(ctx context.Context, req *pb.SkyTradeAddSpeedReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215003, data)
}

// 请求移动
func (c *Client) SendSkyTradeGotoPortReq(ctx context.Context, req *pb.SkyTradeGotoPortReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215004, data)
}

// 购买货物
func (c *Client) SendSkyTradeDealReq(ctx context.Context, req *pb.SkyTradeDealReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215005, data)
}

// 挑战列表
func (c *Client) SendSkyTradeChallengeListReq(ctx context.Context, req *pb.SkyTradeChallengeListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215006, data)
}

// 请求挑战
func (c *Client) SendSkyTradeChallengeReq(ctx context.Context, req *pb.SkyTradeChallengeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215007, data)
}

// 妖盟成员跑商威望信息
func (c *Client) SendSkyTradeUnionFameReq(ctx context.Context, req *pb.SkyTradeUnionFameReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215008, data)
}

// 领取到港事件奖励
func (c *Client) SendSkyTradeGetRewardReq(ctx context.Context, req *pb.SkyTradeGetRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215009, data)
}

// 请求战报
func (c *Client) SendSkyTradeReportReq(ctx context.Context, req *pb.SkyTradeReportReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215010, data)
}

// 获取日志
func (c *Client) SendSkyTradePortLogReq(ctx context.Context, req *pb.SkyTradePortLogReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215011, data)
}

// 天地灵石信息
func (c *Client) SendSkyTradeSparInfoReq(ctx context.Context, req *pb.SkyTradeSparInfoReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215012, data)
}

// 领取天地灵石等级奖励
func (c *Client) SendSkyTradeGetSparPowerReq(ctx context.Context, req *pb.SkyTradeGetSparPowerReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215013, data)
}

// 领取天地灵石盟友奖励
func (c *Client) SendSkyTradeGetWelfareReq(ctx context.Context, req *pb.SkyTradeGetWelfareReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215014, data)
}

// 发送灵液
func (c *Client) SendSkyTradeSendWelfareReq(ctx context.Context, req *pb.SkyTradeSendWelfareReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215015, data)
}

// 灵液记录
func (c *Client) SendSkyTradeWelfareRecordReq(ctx context.Context, req *pb.SkyTradeWelfareRecordReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215016, data)
}

// 挑战界面心跳
func (c *Client) SendSkyTradeChallengeHeartBeatReq(ctx context.Context, req *pb.SkyTradeChallengeHeartBeatReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215017, data)
}

// 使用道具补充挑战次数
func (c *Client) SendSkyTradeUseRobItemReq(ctx context.Context, req *pb.SkyTradeUseRobItemReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215018, data)
}

// 请求详细战报信息
func (c *Client) SendSkyTradeReportDetailReq(ctx context.Context, req *pb.SkyTradeReportDetailReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215019, data)
}

// 竞猜结果
func (c *Client) SendSkyTradeGuessPlayersReq(ctx context.Context, req *pb.SkyTradeGuessPlayersReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215020, data)
}

// 请求竞猜界面数据
func (c *Client) SendSkyTradeGuessDataReq(ctx context.Context, req *pb.SkyTradeGuessDataReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215021, data)
}

// 设置竞猜选择
func (c *Client) SendSkyTradeGuessSelectReq(ctx context.Context, req *pb.SkyTradeGuessSelectReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215022, data)
}

// 领取竞猜奖励
func (c *Client) SendSkyTradeGuessRewardReq(ctx context.Context, req *pb.SkyTradeGuessRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215023, data)
}

// 请求上报所在地点稀缺货物
func (c *Client) SendSkyTradeReportGoodsReq(ctx context.Context, req *pb.SkyTradeReportGoodsReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215024, data)
}

// 请求刷新货物库存
func (c *Client) SendSkyTradeResetStockReq(ctx context.Context, req *pb.SkyTradeResetStockReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215025, data)
}

// 红点详情
func (c *Client) SendSkyTradeRedPointReq(ctx context.Context, req *pb.SkyTradeRedPointReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215026, data)
}

// 妖盟分段为伤害信息
func (c *Client) SendSkyTradeUnionGroupDamageReq(ctx context.Context, req *pb.SkyTradeUnionGroupDamageReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215027, data)
}

// 支付宝广告礼包-触发
func (c *Client) SendADGiftTriggerReq(ctx context.Context, req *pb.ADGiftTriggerReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215201, data)
}

// 支付宝广告礼包-领取奖励
func (c *Client) SendADGiftGetRewardReq(ctx context.Context, req *pb.ADGiftGetRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215202, data)
}

// 家园-进入
func (c *Client) SendYardEnterReq(ctx context.Context, req *pb.YardEnterReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215801, data)
}

// 家园-战斗解锁区域
func (c *Client) SendYardBattleReq(ctx context.Context, req *pb.YardBattleReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215802, data)
}

// 家园-建筑摆放
func (c *Client) SendYardPosActionReq(ctx context.Context, req *pb.YardPosActionReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215803, data)
}

// 家园升级
func (c *Client) SendYardLevelUpReq(ctx context.Context, req *pb.YardLevelUpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215804, data)
}

// 家园-清理杂物
func (c *Client) SendYardClearReq(ctx context.Context, req *pb.YardClearReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215805, data)
}

// 家园请求
func (c *Client) SendYardInviteReq(ctx context.Context, req *pb.YardInviteReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215806, data)
}

// 进入兑换商店
func (c *Client) SendYardEnterShopReq(ctx context.Context, req *pb.YardEnterShopReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215807, data)
}

// 刷新兑换商店
func (c *Client) SendYardRefreshShopReq(ctx context.Context, req *pb.YardRefreshShopReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215808, data)
}

// 羁绊共鸣升级
func (c *Client) SendYardSkillCombineLvUpReq(ctx context.Context, req *pb.YardSkillCombineLvUpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215821, data)
}

// 抽取建筑
func (c *Client) SendYardDrawReq(ctx context.Context, req *pb.YardDrawReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215822, data)
}

// 建筑升级
func (c *Client) SendBuildUpLevelReq(ctx context.Context, req *pb.BuildUpLevelReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215823, data)
}

// 装饰建筑解锁
func (c *Client) SendBuildUnlockReq(ctx context.Context, req *pb.BuildUnlockReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215824, data)
}

// 建筑生产
func (c *Client) SendYardBuildMakeReq(ctx context.Context, req *pb.YardBuildMakeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215825, data)
}

// 桃子协助
func (c *Client) SendYardBuildHelpReq(ctx context.Context, req *pb.YardBuildHelpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215826, data)
}

// 建筑领奖
func (c *Client) SendYardBuildGainRewardReq(ctx context.Context, req *pb.YardBuildGainRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215827, data)
}

// 建筑升级加速
func (c *Client) SendYardBuildSpeedUpReq(ctx context.Context, req *pb.YardBuildSpeedUpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215828, data)
}

// 妖盟协助列表请求
func (c *Client) SendGetYardUnionHelpDataListReq(ctx context.Context, req *pb.GetYardUnionHelpDataListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215829, data)
}

// 请求妖盟协助
func (c *Client) SendRequestYardUnionHelpReq(ctx context.Context, req *pb.RequestYardUnionHelpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215830, data)
}

// 协助他人
func (c *Client) SendYardUnionHelpReq(ctx context.Context, req *pb.YardUnionHelpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215831, data)
}

// 家园装饰建筑一键升级
func (c *Client) SendYardOneKeyLevelUpReq(ctx context.Context, req *pb.YardOneKeyLevelUpReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215833, data)
}

// 桃树协助列表
func (c *Client) SendYardPeachRecordListReq(ctx context.Context, req *pb.YardPeachRecordListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215834, data)
}

// 产物商人信息
func (c *Client) SendYardMerchantInfoReq(ctx context.Context, req *pb.YardMerchantInfoReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215835, data)
}

// 产物出售
func (c *Client) SendYardMerchantExchangeReq(ctx context.Context, req *pb.YardMerchantExchangeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215836, data)
}

// 留言板开关
func (c *Client) SendYardChatSwitchReq(ctx context.Context, req *pb.YardChatSwitchReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215837, data)
}

// 留言操作
func (c *Client) SendYardChatActionReq(ctx context.Context, req *pb.YardChatActionReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215839, data)
}

// 留言
func (c *Client) SendYardChatRecordReq(ctx context.Context, req *pb.YardChatRecordReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215840, data)
}

// 留言板信息
func (c *Client) SendYardChatRecordListReq(ctx context.Context, req *pb.YardChatRecordListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215841, data)
}

// 点赞
func (c *Client) SendYardGiveLikeReq(ctx context.Context, req *pb.YardGiveLikeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215842, data)
}

// 家园举报
func (c *Client) SendYardReportReq(ctx context.Context, req *pb.YardReportReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215852, data)
}

// 申请、撤销
func (c *Client) SendShuraBattleApplyReq(ctx context.Context, req *pb.ShuraBattleApplyReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 215910, data)
}

// 进入三界征途界面
func (c *Client) SendEnterPlanesTrialReq(ctx context.Context, req *pb.EnterPlanesTrialReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216001, data)
}

// 组队内部请求
func (c *Client) SendEnterPlanesTrialTeamReq(ctx context.Context, req *pb.EnterPlanesTrialTeamReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216002, data)
}

// 组队页面请求
func (c *Client) SendPlanesTrialTeamStartReq(ctx context.Context, req *pb.PlanesTrialTeamStartReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216003, data)
}

// 创建队伍
func (c *Client) SendPlanesTrialCreateTeamReq(ctx context.Context, req *pb.PlanesTrialCreateTeamReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216004, data)
}

// 获取队伍列表
func (c *Client) SendPlanesTrialGetTeamListReq(ctx context.Context, req *pb.PlanesTrialGetTeamListReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216005, data)
}

// 查找队伍
func (c *Client) SendPlanesTrialGetTeamInfoReq(ctx context.Context, req *pb.PlanesTrialGetTeamInfoReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216006, data)
}

// 三界征途-申请加入队伍
func (c *Client) SendPlanesTrialJoinTeamReq(ctx context.Context, req *pb.PlanesTrialJoinTeamReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216007, data)
}

// 取消加入队伍
func (c *Client) SendPlanesTrialCancelTeamApplyReq(ctx context.Context, req *pb.PlanesTrialCancelTeamApplyReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216008, data)
}

// 同意申请
func (c *Client) SendPlanesTrialApplyJoinTeamAgreeReq(ctx context.Context, req *pb.PlanesTrialApplyJoinTeamAgreeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216009, data)
}

// 一键拒绝
func (c *Client) SendPlanesTrialApplyJoinTeamRefusedReq(ctx context.Context, req *pb.PlanesTrialApplyJoinTeamRefusedReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216010, data)
}

// 退出队伍
func (c *Client) SendPlanesTrialQuitTeamReq(ctx context.Context, req *pb.PlanesTrialQuitTeamReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216011, data)
}

// 踢出队伍
func (c *Client) SendPlanesTrialKickOutTeamReq(ctx context.Context, req *pb.PlanesTrialKickOutTeamReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216012, data)
}

// 转让队长
func (c *Client) SendPlanesTrialChangeLeaderReq(ctx context.Context, req *pb.PlanesTrialChangeLeaderReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216013, data)
}

// 开始准备
func (c *Client) SendPlanesTrialTeamPrepareReq(ctx context.Context, req *pb.PlanesTrialTeamPrepareReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216014, data)
}

// 匹配
func (c *Client) SendPlanesTrialMatchMemberReq(ctx context.Context, req *pb.PlanesTrialMatchMemberReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216015, data)
}

// 开始战斗
func (c *Client) SendPlanesTrialStartBattleReq(ctx context.Context, req *pb.PlanesTrialStartBattleReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216016, data)
}

// 请求玩家数据
func (c *Client) SendPlanesTrialGetPlayerInfoReq(ctx context.Context, req *pb.PlanesTrialGetPlayerInfoReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216017, data)
}

// 战斗回放
func (c *Client) SendPlanesTrialGetBattleVideoReq(ctx context.Context, req *pb.PlanesTrialGetBattleVideoReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216019, data)
}

// 获取boss信息
func (c *Client) SendPlanesTrialGetBossInfoReq(ctx context.Context, req *pb.PlanesTrialGetBossInfoReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216020, data)
}

// 三界征途-获取挑战次数
func (c *Client) SendPlanesTrialChallengeTimeReq(ctx context.Context, req *pb.PlanesTrialChallengeTimeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216021, data)
}

// 开始匹配
func (c *Client) SendPlanesTrialStartMatchReq(ctx context.Context, req *pb.PlanesTrialStartMatchReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216022, data)
}

// 邀请
func (c *Client) SendPlanesTrialInviteReq(ctx context.Context, req *pb.PlanesTrialInviteReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216023, data)
}

// 离开系统
func (c *Client) SendLeavePlanesTrialReq(ctx context.Context, req *pb.LeavePlanesTrialReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216024, data)
}

// 获取boss妖力
func (c *Client) SendPlanesTrialGetBossPowerReq(ctx context.Context, req *pb.PlanesTrialGetBossPowerReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216025, data)
}

// 队长跳过战斗
func (c *Client) SendPlanesTrialSkipBattleReq(ctx context.Context, req *pb.PlanesTrialSkipBattleReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216026, data)
}

// 队长开始选buff
func (c *Client) SendPlanesTrialStartSelectBuffReq(ctx context.Context, req *pb.PlanesTrialStartSelectBuffReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216027, data)
}

// 选buff
func (c *Client) SendPlanesTrialSelectBuffReq(ctx context.Context, req *pb.PlanesTrialSelectBuffReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216028, data)
}

// 保存选择选项
func (c *Client) SendPlanesTrialSetBuffPreferenceReq(ctx context.Context, req *pb.PlanesTrialSetBuffPreferenceReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216029, data)
}

// 查看自己选的buff
func (c *Client) SendPlanesTrialGetSelectedBuffReq(ctx context.Context, req *pb.PlanesTrialGetSelectedBuffReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216030, data)
}

// 领取成就奖励
func (c *Client) SendPlanesTrialGetAchievementRewardReq(ctx context.Context, req *pb.PlanesTrialGetAchievementRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216031, data)
}

// 查看预设buff
func (c *Client) SendPlanesTrialGetBuffPreferenceReq(ctx context.Context, req *pb.PlanesTrialGetBuffPreferenceReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216033, data)
}

// 获取分身数据
func (c *Client) SendPlanesTrialGetGodBodyDataReq(ctx context.Context, req *pb.PlanesTrialGetGodBodyDataReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216034, data)
}

// 获取大奖信息
func (c *Client) SendPlanesTrialGetGrandPrizeInfoReq(ctx context.Context, req *pb.PlanesTrialGetGrandPrizeInfoReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216035, data)
}

// 领取大奖
func (c *Client) SendPlanesTrialReceiveGrandPrizeReq(ctx context.Context, req *pb.PlanesTrialReceiveGrandPrizeReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216036, data)
}

// 进入布阵
func (c *Client) SendPlanesTrialEnterBattleReq(ctx context.Context, req *pb.PlanesTrialEnterBattleReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216053, data)
}

// 进入切换分身界面
func (c *Client) SendPlanesTrialEnterSwitchSeparationReq(ctx context.Context, req *pb.PlanesTrialEnterSwitchSeparationReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216054, data)
}

// 切换上阵的分身
func (c *Client) SendPlanesTrialSwitchSeparationReq(ctx context.Context, req *pb.PlanesTrialSwitchSeparationReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216055, data)
}

// 查看上阵分身详细数据
func (c *Client) SendPlanesTrialSeparationDetailReq(ctx context.Context, req *pb.PlanesTrialSeparationDetailReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216056, data)
}

// 变更参战玩家位置
func (c *Client) SendPlanesTrialChangePosReq(ctx context.Context, req *pb.PlanesTrialChangePosReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216057, data)
}

// 进行战斗
func (c *Client) SendPlanesTrialDoBattleReq(ctx context.Context, req *pb.PlanesTrialDoBattleReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216058, data)
}

// 更新绑定数据
func (c *Client) SendPlanesTrialUpdateLockDataReq(ctx context.Context, req *pb.PlanesTrialUpdateLockDataReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216072, data)
}

// 获取被压制属性
func (c *Client) SendPlayerRestrainInfoMsgReq(ctx context.Context, req *pb.PlayerRestrainInfoMsgReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216088, data)
}

// 获取秘境人数请求
func (c *Client) SendGetPlanesTrialTrialMemberCountReq(ctx context.Context, req *pb.GetPlanesTrialTrialMemberCountReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216089, data)
}

// 获取排行榜
func (c *Client) SendPlanesTrialRankGetReq(ctx context.Context, req *pb.PlanesTrialRankGetReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216090, data)
}

// 心跳包
func (c *Client) SendPlanesTrialHeartbeatReq(ctx context.Context, req *pb.PlanesTrialHeartbeatReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216091, data)
}

// 获取自选奖励信息
func (c *Client) SendPlanesTrialGetSelectRewardDetailReq(ctx context.Context, req *pb.PlanesTrialGetSelectRewardDetailReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216092, data)
}

// 选择奖励
func (c *Client) SendPlanesTrialSelectRewardReq(ctx context.Context, req *pb.PlanesTrialSelectRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216093, data)
}

// 查看关卡录像数据
func (c *Client) SendPlanesTrialVideoInfoReq(ctx context.Context, req *pb.PlanesTrialVideoInfoReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216094, data)
}

// 播放关卡录像
func (c *Client) SendPlanesTrialPlayVideoReq(ctx context.Context, req *pb.PlanesTrialPlayVideoReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216095, data)
}

// 播放关卡录像
func (c *Client) SendPlanesTrialCloseSettleScreenReq(ctx context.Context, req *pb.PlanesTrialCloseSettleScreenReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216096, data)
}

// 妖盟寻宝地图互动
func (c *Client) SendUnionTreasureDrawChipReq(ctx context.Context, req *pb.UnionTreasureDrawChipReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216201, data)
}

// 妖盟寻宝索要
func (c *Client) SendUnionTreasureAskForReq(ctx context.Context, req *pb.UnionTreasureAskForReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216202, data)
}

// 妖盟寻宝碎片赠送
func (c *Client) SendUnionTreasureSendChipReq(ctx context.Context, req *pb.UnionTreasureSendChipReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216203, data)
}

// 妖盟寻宝遗迹互动
func (c *Client) SendUnionTreasureBattleRelicReq(ctx context.Context, req *pb.UnionTreasureBattleRelicReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216205, data)
}

// 妖盟寻宝遗迹挖宝
func (c *Client) SendUnionTreasureHuntTreasureReq(ctx context.Context, req *pb.UnionTreasureHuntTreasureReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216206, data)
}

// 妖盟寻宝进入界面
func (c *Client) SendUnionTreasureEnterReq(ctx context.Context, req *pb.UnionTreasureEnterReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216207, data)
}

// 妖盟寻宝获得碎片
func (c *Client) SendUnionTreasureGetGiveChipReq(ctx context.Context, req *pb.UnionTreasureGetGiveChipReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216209, data)
}

// 妖盟寻宝遗迹数据请求
func (c *Client) SendUnionTreasureAskForRelicReq(ctx context.Context, req *pb.UnionTreasureAskForRelicReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216213, data)
}

// 妖盟寻宝分享协助
func (c *Client) SendUnionTreasureShareRelicReq(ctx context.Context, req *pb.UnionTreasureShareRelicReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216216, data)
}

// 妖盟寻宝遗迹激活碎片
func (c *Client) SendUnionTreasureFillChipReq(ctx context.Context, req *pb.UnionTreasureFillChipReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216218, data)
}

// 妖盟寻宝遗迹消息过期
func (c *Client) SendCheckMessageValidReq(ctx context.Context, req *pb.CheckMessageValidReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216222, data)
}

// 妖盟寻宝战斗记录
func (c *Client) SendUnionTreasureGetBattleRecordReq(ctx context.Context, req *pb.UnionTreasureGetBattleRecordReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216223, data)
}

// 妖盟寻宝赠送请求
func (c *Client) SendUnionTreasureGetCollectGiveMsgReq(ctx context.Context, req *pb.UnionTreasureGetCollectGiveMsgReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216224, data)
}

// 妖盟寻宝退出请求
func (c *Client) SendUnionTreasureExitReq(ctx context.Context, req *pb.UnionTreasureExitReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216225, data)
}

// 美团订阅成功
func (c *Client) SendMeiTuanSubscribeSucceedReq(ctx context.Context, req *pb.MeiTuanSubscribeSucceedReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216701, data)
}

// 美团订阅领取
func (c *Client) SendMeiTuanSubscribeRewardReq(ctx context.Context, req *pb.MeiTuanSubscribeRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216702, data)
}

// 美团二楼登录
func (c *Client) SendMeiTuanDailyLoginReq(ctx context.Context, req *pb.MeiTuanDailyLoginReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216801, data)
}

// 美团登录领取
func (c *Client) SendMeiTuanDailyRewardReq(ctx context.Context, req *pb.MeiTuanDailyRewardReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 216802, data)
}

// 仙玉宝府进入请求
func (c *Client) SendYueBaoEnterReq(ctx context.Context, req *pb.YueBaoEnterReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 217001, data)
}

// 仙玉宝府存储
func (c *Client) SendYueBaoDepositReq(ctx context.Context, req *pb.YueBaoDepositReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 217002, data)
}

// 仙玉宝府互动请求
func (c *Client) SendYueBaoInteractReq(ctx context.Context, req *pb.YueBaoInteractReq) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 217003, data)
}

// 同步直播状态
func (c *Client) SendLiveShowNotifyReqMsg(ctx context.Context, req *pb.LiveShowNotifyReqMsg) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, 217701, data)
}


// 登陆消息
func (c *Client) OnReqLoginMsg(f func(ctx Context, msg *pb.RspLoginMsg)error) {
	c.RegisterHandler(1, func(ctx Context) error {
		msg := &pb.RspLoginMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 登录下发完成
func (c *Client) OnLoginOverMsg(f func(ctx Context, msg *pb.LoginOverMsg)error) {
	c.RegisterHandler(2, func(ctx Context) error {
		msg := &pb.LoginOverMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 心跳
func (c *Client) OnReqPingMsg(f func(ctx Context, msg *pb.RspPingMsg)error) {
	c.RegisterHandler(3, func(ctx Context) error {
		msg := &pb.RspPingMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 预购买
func (c *Client) OnPreChargeReq(f func(ctx Context, msg *pb.PreChargeResp)error) {
	c.RegisterHandler(4, func(ctx Context) error {
		msg := &pb.PreChargeResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 通知充值到账
func (c *Client) OnNotifyRechargeSuccess(f func(ctx Context, msg *pb.NotifyRechargeSuccess)error) {
	c.RegisterHandler(5, func(ctx Context) error {
		msg := &pb.NotifyRechargeSuccess{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 创建角色请求
func (c *Client) OnCreateRoleReq(f func(ctx Context, msg *pb.CreateRoleResp)error) {
	c.RegisterHandler(6, func(ctx Context) error {
		msg := &pb.CreateRoleResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 随机昵称回复
func (c *Client) OnRandomNickNameResp(f func(ctx Context, msg *pb.RandomNickNameResp)error) {
	c.RegisterHandler(7, func(ctx Context) error {
		msg := &pb.RandomNickNameResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 改名
func (c *Client) OnChangeNickNameReq(f func(ctx Context, msg *pb.ChangeNickNameResp)error) {
	c.RegisterHandler(8, func(ctx Context) error {
		msg := &pb.ChangeNickNameResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 兑换码
func (c *Client) OnUseExchangeKeyReq(f func(ctx Context, msg *pb.UseExchangeKeyResp)error) {
	c.RegisterHandler(9, func(ctx Context) error {
		msg := &pb.UseExchangeKeyResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 服务端配置
func (c *Client) OnGetServerConfigVersionResp(f func(ctx Context, msg *pb.GetServerConfigVersionResp)error) {
	c.RegisterHandler(10, func(ctx Context) error {
		msg := &pb.GetServerConfigVersionResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 其他人登录你的账户
func (c *Client) OnOtherLoginMsg(f func(ctx Context, msg *pb.OtherLoginMsg)error) {
	c.RegisterHandler(11, func(ctx Context) error {
		msg := &pb.OtherLoginMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 系统红点同步
func (c *Client) OnSystemRedPointSync(f func(ctx Context, msg *pb.SystemRedPointSync)error) {
	c.RegisterHandler(12, func(ctx Context) error {
		msg := &pb.SystemRedPointSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 活动红点同步
func (c *Client) OnActivityRedPointSync(f func(ctx Context, msg *pb.ActivityRedPointSync)error) {
	c.RegisterHandler(13, func(ctx Context) error {
		msg := &pb.ActivityRedPointSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 充值商品次数同步
func (c *Client) OnRechargeMallTimesSync(f func(ctx Context, msg *pb.RechargeMallTimesSync)error) {
	c.RegisterHandler(14, func(ctx Context) error {
		msg := &pb.RechargeMallTimesSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步后台配置
func (c *Client) OnRspConfigMsg(f func(ctx Context, msg *pb.RspConfigMsg)error) {
	c.RegisterHandler(16, func(ctx Context) error {
		msg := &pb.RspConfigMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 观看视频广告任务
func (c *Client) OnWatchAdReq(f func(ctx Context, msg *pb.WatchAdResp)error) {
	c.RegisterHandler(19, func(ctx Context) error {
		msg := &pb.WatchAdResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 熊口逃生领奖
func (c *Client) OnCutRopeGetRewardReq(f func(ctx Context, msg *pb.CutRopeGetRewardRsp)error) {
	c.RegisterHandler(20, func(ctx Context) error {
		msg := &pb.CutRopeGetRewardRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 小游戏领奖
func (c *Client) OnReceiveMiniGamesRewardReq(f func(ctx Context, msg *pb.ReceiveMiniGamesRewardResp)error) {
	c.RegisterHandler(21, func(ctx Context) error {
		msg := &pb.ReceiveMiniGamesRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 秘境挑战
func (c *Client) OnStageMapChallengeReq(f func(ctx Context, msg *pb.StageMapChallengeRsp)error) {
	c.RegisterHandler(22, func(ctx Context) error {
		msg := &pb.StageMapChallengeRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 禁地试炼领奖
func (c *Client) OnForbiddenTrailsGetRewardReq(f func(ctx Context, msg *pb.ForbiddenTrailsGetRewardResp)error) {
	c.RegisterHandler(23, func(ctx Context) error {
		msg := &pb.ForbiddenTrailsGetRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 西天救援进入
func (c *Client) OnEnterWestJourneyReq(f func(ctx Context, msg *pb.EnterWestJourneyResp)error) {
	c.RegisterHandler(24, func(ctx Context) error {
		msg := &pb.EnterWestJourneyResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 西天救援关卡数据同步
func (c *Client) OnMiniGamesWestJourneyData(f func(ctx Context, msg *pb.MiniGamesWestJourneyData)error) {
	c.RegisterHandler(25, func(ctx Context) error {
		msg := &pb.MiniGamesWestJourneyData{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 西天救援领奖
func (c *Client) OnGetWestJourneyRewardReq(f func(ctx Context, msg *pb.GetWestJourneyRewardResp)error) {
	c.RegisterHandler(26, func(ctx Context) error {
		msg := &pb.GetWestJourneyRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取隐私注册协议
func (c *Client) OnGetPolicyInfoReq(f func(ctx Context, msg *pb.GetPolicyInfoRsp)error) {
	c.RegisterHandler(31, func(ctx Context) error {
		msg := &pb.GetPolicyInfoRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步协议 回归玩家信息
func (c *Client) OnSyncRecallPlayerMsg(f func(ctx Context, msg *pb.SyncRecallPlayerMsg)error) {
	c.RegisterHandler(32, func(ctx Context) error {
		msg := &pb.SyncRecallPlayerMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取绑定状态
func (c *Client) OnGetBindStatusReq(f func(ctx Context, msg *pb.GetBindStatusResp)error) {
	c.RegisterHandler(33, func(ctx Context) error {
		msg := &pb.GetBindStatusResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 生成绑定码
func (c *Client) OnGetBindCodeReq(f func(ctx Context, msg *pb.GetBindCodeResp)error) {
	c.RegisterHandler(34, func(ctx Context) error {
		msg := &pb.GetBindCodeResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 解绑
func (c *Client) OnUntieBindCodeReq(f func(ctx Context, msg *pb.UntieBindCodeResp)error) {
	c.RegisterHandler(35, func(ctx Context) error {
		msg := &pb.UntieBindCodeResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取绑定信息
func (c *Client) OnGetBindCodeInfoReq(f func(ctx Context, msg *pb.GetBindCodeInfoResp)error) {
	c.RegisterHandler(36, func(ctx Context) error {
		msg := &pb.GetBindCodeInfoResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 注销账号
func (c *Client) OnDeletePlayerReq(f func(ctx Context, msg *pb.DeletePlayerResp)error) {
	c.RegisterHandler(37, func(ctx Context) error {
		msg := &pb.DeletePlayerResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取系统分组信息
func (c *Client) OnGetSystemGroupInfoReq(f func(ctx Context, msg *pb.GetSystemGroupInfoResp)error) {
	c.RegisterHandler(40, func(ctx Context) error {
		msg := &pb.GetSystemGroupInfoResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 用户信息同步
func (c *Client) OnPlayerDataMsg(f func(ctx Context, msg *pb.PlayerDataMsg)error) {
	c.RegisterHandler(101, func(ctx Context) error {
		msg := &pb.PlayerDataMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 系统解锁同步
func (c *Client) OnSystemUnlockSync(f func(ctx Context, msg *pb.SystemUnlockSync)error) {
	c.RegisterHandler(102, func(ctx Context) error {
		msg := &pb.SystemUnlockSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求领取系统解锁奖励
func (c *Client) OnGetSysRewardReq(f func(ctx Context, msg *pb.GetSysRewardResp)error) {
	c.RegisterHandler(103, func(ctx Context) error {
		msg := &pb.GetSysRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步特权卡数据
func (c *Client) OnPrivilegeCardDataMsg(f func(ctx Context, msg *pb.PrivilegeCardDataMsg)error) {
	c.RegisterHandler(104, func(ctx Context) error {
		msg := &pb.PrivilegeCardDataMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取特权卡奖励
func (c *Client) OnPrivilegeCardReceiveRewardReq(f func(ctx Context, msg *pb.PrivilegeCardReceiveRewardRsp)error) {
	c.RegisterHandler(105, func(ctx Context) error {
		msg := &pb.PrivilegeCardReceiveRewardRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 玩家形象同步
func (c *Client) OnPlayerCharaDataMsg(f func(ctx Context, msg *pb.PlayerCharaDataMsg)error) {
	c.RegisterHandler(106, func(ctx Context) error {
		msg := &pb.PlayerCharaDataMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 切换形象请求
func (c *Client) OnChangePlayerCharaReq(f func(ctx Context, msg *pb.ChangePlayerCharaResp)error) {
	c.RegisterHandler(107, func(ctx Context) error {
		msg := &pb.ChangePlayerCharaResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取形象回复
func (c *Client) OnGetPlayerCharaListResp(f func(ctx Context, msg *pb.GetPlayerCharaListResp)error) {
	c.RegisterHandler(108, func(ctx Context) error {
		msg := &pb.GetPlayerCharaListResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取玩家详细数据
func (c *Client) OnGetPlayerDetailDataReq(f func(ctx Context, msg *pb.GetPlayerDetailDataResp)error) {
	c.RegisterHandler(109, func(ctx Context) error {
		msg := &pb.GetPlayerDetailDataResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求聊天
func (c *Client) OnWorldChatReqMsg(f func(ctx Context, msg *pb.WorldChatRespMsg)error) {
	c.RegisterHandler(110, func(ctx Context) error {
		msg := &pb.WorldChatRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 世界消息同步
func (c *Client) OnWorldMessageListMsg20111(f func(ctx Context, msg *pb.WorldMessageListMsg)error) {
	c.RegisterHandler(111, func(ctx Context) error {
		msg := &pb.WorldMessageListMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 世界消息列表，登录时下发
func (c *Client) OnWorldMessageListMsg20112(f func(ctx Context, msg *pb.WorldMessageListMsg)error) {
	c.RegisterHandler(112, func(ctx Context) error {
		msg := &pb.WorldMessageListMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟消息同步
func (c *Client) OnWorldMessageListMsg20113(f func(ctx Context, msg *pb.WorldMessageListMsg)error) {
	c.RegisterHandler(113, func(ctx Context) error {
		msg := &pb.WorldMessageListMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟消息列表，登录时下发
func (c *Client) OnWorldMessageListMsg20114(f func(ctx Context, msg *pb.WorldMessageListMsg)error) {
	c.RegisterHandler(114, func(ctx Context) error {
		msg := &pb.WorldMessageListMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 活动消息同步
func (c *Client) OnWorldMessageListMsg20115(f func(ctx Context, msg *pb.WorldMessageListMsg)error) {
	c.RegisterHandler(115, func(ctx Context) error {
		msg := &pb.WorldMessageListMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求活动消息列表
func (c *Client) OnActivityChatReqMsg(f func(ctx Context, msg *pb.WorldMessageListMsg)error) {
	c.RegisterHandler(116, func(ctx Context) error {
		msg := &pb.WorldMessageListMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 跑马灯广播通知
func (c *Client) OnHorseRaceLampMsgAdd(f func(ctx Context, msg *pb.HorseRaceLampMsgAdd)error) {
	c.RegisterHandler(117, func(ctx Context) error {
		msg := &pb.HorseRaceLampMsgAdd{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取仙树后面飘的玩家信息
func (c *Client) OnGetTopPlayerInfoReq(f func(ctx Context, msg *pb.GetTopPlayerInfoResp)error) {
	c.RegisterHandler(118, func(ctx Context) error {
		msg := &pb.GetTopPlayerInfoResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 玩家称号同步
func (c *Client) OnTitleSyncMsg(f func(ctx Context, msg *pb.TitleSyncMsg)error) {
	c.RegisterHandler(119, func(ctx Context) error {
		msg := &pb.TitleSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 穿戴称号
func (c *Client) OnUseTitleReq(f func(ctx Context, msg *pb.UseTitleRsp)error) {
	c.RegisterHandler(120, func(ctx Context) error {
		msg := &pb.UseTitleRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 道具屏蔽列表同步
func (c *Client) OnGoodsShieldSync(f func(ctx Context, msg *pb.GoodsShieldSync)error) {
	c.RegisterHandler(121, func(ctx Context) error {
		msg := &pb.GoodsShieldSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 玩家勋章同步
func (c *Client) OnMedalSyncMsg(f func(ctx Context, msg *pb.MedalSyncMsg)error) {
	c.RegisterHandler(122, func(ctx Context) error {
		msg := &pb.MedalSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 穿戴称号
func (c *Client) OnUseMedalReq(f func(ctx Context, msg *pb.UseMedalRsp)error) {
	c.RegisterHandler(123, func(ctx Context) error {
		msg := &pb.UseMedalRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 活动海报屏蔽列表同步
func (c *Client) OnActivityPosterShieldSync(f func(ctx Context, msg *pb.ActivityPosterShieldSync)error) {
	c.RegisterHandler(124, func(ctx Context) error {
		msg := &pb.ActivityPosterShieldSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步渠道屏蔽商城付费项列表
func (c *Client) OnMallRechargeIdShieldSync(f func(ctx Context, msg *pb.MallRechargeIdShieldSync)error) {
	c.RegisterHandler(125, func(ctx Context) error {
		msg := &pb.MallRechargeIdShieldSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 举报玩家
func (c *Client) OnWorldChatBlockReqMsg(f func(ctx Context, msg *pb.WorldChatBlockRespMsg)error) {
	c.RegisterHandler(130, func(ctx Context) error {
		msg := &pb.WorldChatBlockRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 系统聊天同步
func (c *Client) OnWorldMessageListMsg20131(f func(ctx Context, msg *pb.WorldMessageListMsg)error) {
	c.RegisterHandler(131, func(ctx Context) error {
		msg := &pb.WorldMessageListMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 系统聊天列表，请求下发
func (c *Client) OnSystemChatReqMsg(f func(ctx Context, msg *pb.WorldMessageListMsg)error) {
	c.RegisterHandler(132, func(ctx Context) error {
		msg := &pb.WorldMessageListMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 聊天举报
func (c *Client) OnReportMessageReqMsg(f func(ctx Context, msg *pb.ReportMessageRespMsg)error) {
	c.RegisterHandler(133, func(ctx Context) error {
		msg := &pb.ReportMessageRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取玩家(某个系统)锁定的详细数据-请求
func (c *Client) OnGetLockedPlayerDetailDataReq(f func(ctx Context, msg *pb.GetLockedPlayerDetailDataResp)error) {
	c.RegisterHandler(134, func(ctx Context) error {
		msg := &pb.GetLockedPlayerDetailDataResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求飞行员列表数据-请求
func (c *Client) OnGetPlayerPilotShowDataMsgListReq(f func(ctx Context, msg *pb.GetPlayerPilotShowDataMsgListResp)error) {
	c.RegisterHandler(135, func(ctx Context) error {
		msg := &pb.GetPlayerPilotShowDataMsgListResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 通用搜索玩家
func (c *Client) OnCommonSearchPlayerReq(f func(ctx Context, msg *pb.CommonSearchPlayerResp)error) {
	c.RegisterHandler(136, func(ctx Context) error {
		msg := &pb.CommonSearchPlayerResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 红包状态同步
func (c *Client) OnRedPacketStateMsgSync(f func(ctx Context, msg *pb.RedPacketStateMsgSync)error) {
	c.RegisterHandler(140, func(ctx Context) error {
		msg := &pb.RedPacketStateMsgSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取红包信息
func (c *Client) OnRedPacketOpenReq(f func(ctx Context, msg *pb.RedPacketOpenResp)error) {
	c.RegisterHandler(141, func(ctx Context) error {
		msg := &pb.RedPacketOpenResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取红包奖励
func (c *Client) OnRedPacketDrawReq(f func(ctx Context, msg *pb.RedPacketDrawResp)error) {
	c.RegisterHandler(142, func(ctx Context) error {
		msg := &pb.RedPacketDrawResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 皮肤羁绊任务领取
func (c *Client) OnCollectPlayerCharaReq(f func(ctx Context, msg *pb.CollectPlayerCharaResp)error) {
	c.RegisterHandler(150, func(ctx Context) error {
		msg := &pb.CollectPlayerCharaResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 皮肤等级升级
func (c *Client) OnUpgradePlayerCharaReq(f func(ctx Context, msg *pb.UpgradePlayerCharaResp)error) {
	c.RegisterHandler(151, func(ctx Context) error {
		msg := &pb.UpgradePlayerCharaResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 玩家属性信息同步
func (c *Client) OnPlayerAttributeDataMsg(f func(ctx Context, msg *pb.PlayerAttributeDataMsg)error) {
	c.RegisterHandler(201, func(ctx Context) error {
		msg := &pb.PlayerAttributeDataMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 处理装备
func (c *Client) OnReqDealEquipmentMsg(f func(ctx Context, msg *pb.RspDealEquipmentMsg)error) {
	c.RegisterHandler(202, func(ctx Context) error {
		msg := &pb.RspDealEquipmentMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 做梦
func (c *Client) OnReqDreamMsg(f func(ctx Context, msg *pb.RspDreamMsg)error) {
	c.RegisterHandler(203, func(ctx Context) error {
		msg := &pb.RspDreamMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 境界升级
func (c *Client) OnReqRealmsLeveUpMsg(f func(ctx Context, msg *pb.RspRealmsLeveUpMsg)error) {
	c.RegisterHandler(204, func(ctx Context) error {
		msg := &pb.RspRealmsLeveUpMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 梦境升级
func (c *Client) OnDreamLvUpReq(f func(ctx Context, msg *pb.DreamLvUpResp)error) {
	c.RegisterHandler(205, func(ctx Context) error {
		msg := &pb.DreamLvUpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 梦境升级加速
func (c *Client) OnDreamLvUpSpeedUpReq(f func(ctx Context, msg *pb.DreamLvUpSpeedUpResp)error) {
	c.RegisterHandler(206, func(ctx Context) error {
		msg := &pb.DreamLvUpSpeedUpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 做梦玩家数据
func (c *Client) OnDreamDataMsg(f func(ctx Context, msg *pb.DreamDataMsg)error) {
	c.RegisterHandler(207, func(ctx Context) error {
		msg := &pb.DreamDataMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取渡劫成功率
func (c *Client) OnGetTribulationSuccessProReq(f func(ctx Context, msg *pb.GetTribulationSuccessProResp)error) {
	c.RegisterHandler(208, func(ctx Context) error {
		msg := &pb.GetTribulationSuccessProResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取未处理装备数据
func (c *Client) OnGetUnDealEquipmentMsgReq(f func(ctx Context, msg *pb.GetUnDealEquipmentMsgResp)error) {
	c.RegisterHandler(209, func(ctx Context) error {
		msg := &pb.GetUnDealEquipmentMsgResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 数据同步
func (c *Client) OnPlayerAdRewardDataMsg(f func(ctx Context, msg *pb.PlayerAdRewardDataMsg)error) {
	c.RegisterHandler(210, func(ctx Context) error {
		msg := &pb.PlayerAdRewardDataMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 广告领奖
func (c *Client) OnGetAdRewardReq(f func(ctx Context, msg *pb.GetAdRewardResp)error) {
	c.RegisterHandler(211, func(ctx Context) error {
		msg := &pb.GetAdRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 斩道也调用这个接口
func (c *Client) OnSoaringReq(f func(ctx Context, msg *pb.SoaringResp)error) {
	c.RegisterHandler(212, func(ctx Context) error {
		msg := &pb.SoaringResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 设置分身名字
func (c *Client) OnSetSeparationNameReq(f func(ctx Context, msg *pb.SetSeparationNameResp)error) {
	c.RegisterHandler(213, func(ctx Context) error {
		msg := &pb.SetSeparationNameResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 切换分身
func (c *Client) OnSwitchSeparationReq(f func(ctx Context, msg *pb.SwitchSeparationResp)error) {
	c.RegisterHandler(214, func(ctx Context) error {
		msg := &pb.SwitchSeparationResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取分身列表数据
func (c *Client) OnGetSeparationDataMsgListReq(f func(ctx Context, msg *pb.GetSeparationDataMsgListResp)error) {
	c.RegisterHandler(215, func(ctx Context) error {
		msg := &pb.GetSeparationDataMsgListResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 微信消息订阅登陆下发
func (c *Client) OnMessageSubscribeInfo(f func(ctx Context, msg *pb.MessageSubscribeInfo)error) {
	c.RegisterHandler(250, func(ctx Context) error {
		msg := &pb.MessageSubscribeInfo{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 设置订阅消息
func (c *Client) OnSetMessageSubscribeDataReqMsg(f func(ctx Context, msg *pb.SetMessageSubscribeDataRspMsg)error) {
	c.RegisterHandler(251, func(ctx Context) error {
		msg := &pb.SetMessageSubscribeDataRspMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 微信头像授权
func (c *Client) OnAuthorizePlayerHeadReq(f func(ctx Context, msg *pb.AuthorizePlayerHeadResp)error) {
	c.RegisterHandler(252, func(ctx Context) error {
		msg := &pb.AuthorizePlayerHeadResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步背包数据
func (c *Client) OnSyncBagMsg(f func(ctx Context, msg *pb.SyncBagMsg)error) {
	c.RegisterHandler(301, func(ctx Context) error {
		msg := &pb.SyncBagMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 使用道具
func (c *Client) OnUsePropReq(f func(ctx Context, msg *pb.UsePropResp)error) {
	c.RegisterHandler(302, func(ctx Context) error {
		msg := &pb.UsePropResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 合成道具
func (c *Client) OnCompoundPropReq(f func(ctx Context, msg *pb.CompoundPropResp)error) {
	c.RegisterHandler(303, func(ctx Context) error {
		msg := &pb.CompoundPropResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 测试战斗
func (c *Client) OnBattleRecordMsg(f func(ctx Context, msg *pb.BattleRecordMsg)error) {
	c.RegisterHandler(401, func(ctx Context) error {
		msg := &pb.BattleRecordMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 关卡挑战
func (c *Client) OnChallengeRspMsg(f func(ctx Context, msg *pb.ChallengeRspMsg)error) {
	c.RegisterHandler(402, func(ctx Context) error {
		msg := &pb.ChallengeRspMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 关卡数据同步
func (c *Client) OnPlayerStageData(f func(ctx Context, msg *pb.PlayerStageData)error) {
	c.RegisterHandler(403, func(ctx Context) error {
		msg := &pb.PlayerStageData{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看关卡怪物信息
func (c *Client) OnViewMonsterAttrResp20404(f func(ctx Context, msg *pb.ViewMonsterAttrResp)error) {
	c.RegisterHandler(404, func(ctx Context) error {
		msg := &pb.ViewMonsterAttrResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 测试组队战斗
func (c *Client) OnTeamBattleRecordMsg(f func(ctx Context, msg *pb.TeamBattleRecordMsg)error) {
	c.RegisterHandler(405, func(ctx Context) error {
		msg := &pb.TeamBattleRecordMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取对战列表
func (c *Client) OnGetBattleListResp(f func(ctx Context, msg *pb.GetBattleListResp)error) {
	c.RegisterHandler(410, func(ctx Context) error {
		msg := &pb.GetBattleListResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 刷新对战列表
func (c *Client) OnRefreshBattleListResp(f func(ctx Context, msg *pb.RefreshBattleListResp)error) {
	c.RegisterHandler(411, func(ctx Context) error {
		msg := &pb.RefreshBattleListResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 挑战请求
func (c *Client) OnRankBattleChallengeReq(f func(ctx Context, msg *pb.RankBattleChallengeResp)error) {
	c.RegisterHandler(412, func(ctx Context) error {
		msg := &pb.RankBattleChallengeResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取妖榜日志
func (c *Client) OnGetRankBattleLogResp(f func(ctx Context, msg *pb.GetRankBattleLogResp)error) {
	c.RegisterHandler(413, func(ctx Context) error {
		msg := &pb.GetRankBattleLogResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖榜复仇请求
func (c *Client) OnRankBattleRevengeReq(f func(ctx Context, msg *pb.RankBattleRevengeResp)error) {
	c.RegisterHandler(414, func(ctx Context) error {
		msg := &pb.RankBattleRevengeResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 斗法登录同步
func (c *Client) OnRankBattleSync(f func(ctx Context, msg *pb.RankBattleSync)error) {
	c.RegisterHandler(415, func(ctx Context) error {
		msg := &pb.RankBattleSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取跨服斗法参与区服
func (c *Client) OnRankBattleServerListReq(f func(ctx Context, msg *pb.RankBattleServerListResp)error) {
	c.RegisterHandler(418, func(ctx Context) error {
		msg := &pb.RankBattleServerListResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 玩家登录任务数据下发
func (c *Client) OnTaskDataListMsg20501(f func(ctx Context, msg *pb.TaskDataListMsg)error) {
	c.RegisterHandler(501, func(ctx Context) error {
		msg := &pb.TaskDataListMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 玩家任务数据同步
func (c *Client) OnTaskDataListMsg20502(f func(ctx Context, msg *pb.TaskDataListMsg)error) {
	c.RegisterHandler(502, func(ctx Context) error {
		msg := &pb.TaskDataListMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取任务奖励响应
func (c *Client) OnTaskGetRewardReqMsg(f func(ctx Context, msg *pb.TaskGetRewardRespMsg)error) {
	c.RegisterHandler(503, func(ctx Context) error {
		msg := &pb.TaskGetRewardRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 邮件列表数据同步
func (c *Client) OnMailListMsg20551(f func(ctx Context, msg *pb.MailListMsg)error) {
	c.RegisterHandler(551, func(ctx Context) error {
		msg := &pb.MailListMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 邮件数据同步
func (c *Client) OnMailListMsg20552(f func(ctx Context, msg *pb.MailListMsg)error) {
	c.RegisterHandler(552, func(ctx Context) error {
		msg := &pb.MailListMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 一键删除邮件
func (c *Client) OnMailDeleteAllReq(f func(ctx Context, msg *pb.MailDeleteAllResp)error) {
	c.RegisterHandler(553, func(ctx Context) error {
		msg := &pb.MailDeleteAllResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 删除邮件
func (c *Client) OnMailDeleteReq(f func(ctx Context, msg *pb.MailDeleteResp)error) {
	c.RegisterHandler(554, func(ctx Context) error {
		msg := &pb.MailDeleteResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 一键领取邮件
func (c *Client) OnMailGetAllRewardReq(f func(ctx Context, msg *pb.MailGetAllRewardResp)error) {
	c.RegisterHandler(555, func(ctx Context) error {
		msg := &pb.MailGetAllRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取邮件
func (c *Client) OnMailGetRewardReq(f func(ctx Context, msg *pb.MailGetRewardResp)error) {
	c.RegisterHandler(556, func(ctx Context) error {
		msg := &pb.MailGetRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 读取邮件
func (c *Client) OnMailReadReq(f func(ctx Context, msg *pb.MailReadResp)error) {
	c.RegisterHandler(557, func(ctx Context) error {
		msg := &pb.MailReadResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 购买商品
func (c *Client) OnBuyGoodsReq(f func(ctx Context, msg *pb.BuyGoodsResp)error) {
	c.RegisterHandler(601, func(ctx Context) error {
		msg := &pb.BuyGoodsResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 商品购买数量集合
func (c *Client) OnMallBuyCountListMsg(f func(ctx Context, msg *pb.MallBuyCountListMsg)error) {
	c.RegisterHandler(602, func(ctx Context) error {
		msg := &pb.MallBuyCountListMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 商店列表
func (c *Client) OnMallRandomGoodsReqMsg(f func(ctx Context, msg *pb.MallRandomGoodsRespMsg)error) {
	c.RegisterHandler(603, func(ctx Context) error {
		msg := &pb.MallRandomGoodsRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 通宝购买请求
func (c *Client) OnMallChanceUseReqMsg(f func(ctx Context, msg *pb.MallChanceUseRespMsg)error) {
	c.RegisterHandler(604, func(ctx Context) error {
		msg := &pb.MallChanceUseRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步神通数据
func (c *Client) OnTalentPlayerDataMsg(f func(ctx Context, msg *pb.TalentPlayerDataMsg)error) {
	c.RegisterHandler(621, func(ctx Context) error {
		msg := &pb.TalentPlayerDataMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 炼化
func (c *Client) OnRandomTalentReq(f func(ctx Context, msg *pb.RandomTalentResp)error) {
	c.RegisterHandler(622, func(ctx Context) error {
		msg := &pb.RandomTalentResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 处理神通
func (c *Client) OnDealTalentReq(f func(ctx Context, msg *pb.DealTalentResp)error) {
	c.RegisterHandler(623, func(ctx Context) error {
		msg := &pb.DealTalentResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 阅读妖书
func (c *Client) OnReadBookReq(f func(ctx Context, msg *pb.ReadBookResp)error) {
	c.RegisterHandler(624, func(ctx Context) error {
		msg := &pb.ReadBookResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取未处理神通数据
func (c *Client) OnGetUnDealTalentMsgReq(f func(ctx Context, msg *pb.GetUnDealTalentMsgResp)error) {
	c.RegisterHandler(625, func(ctx Context) error {
		msg := &pb.GetUnDealTalentMsgResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 数据同步
func (c *Client) OnDestinyData(f func(ctx Context, msg *pb.DestinyData)error) {
	c.RegisterHandler(651, func(ctx Context) error {
		msg := &pb.DestinyData{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 赠礼
func (c *Client) OnReqDestinyGift(f func(ctx Context, msg *pb.RspDestinyGift)error) {
	c.RegisterHandler(652, func(ctx Context) error {
		msg := &pb.RspDestinyGift{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 游历
func (c *Client) OnReqDestinyTravel(f func(ctx Context, msg *pb.RspDestinyTravel)error) {
	c.RegisterHandler(653, func(ctx Context) error {
		msg := &pb.RspDestinyTravel{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 切磋
func (c *Client) OnReqDestinyChallenge(f func(ctx Context, msg *pb.RspDestinyChallenge)error) {
	c.RegisterHandler(654, func(ctx Context) error {
		msg := &pb.RspDestinyChallenge{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 解锁道具仙友
func (c *Client) OnReqUnlockDestinyByItem(f func(ctx Context, msg *pb.RspUnlockDestinyByItem)error) {
	c.RegisterHandler(655, func(ctx Context) error {
		msg := &pb.RspUnlockDestinyByItem{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 皮肤解锁
func (c *Client) OnReqUnlockDestinySkinByItem(f func(ctx Context, msg *pb.RspUnlockDestinySkinByItem)error) {
	c.RegisterHandler(657, func(ctx Context) error {
		msg := &pb.RspUnlockDestinySkinByItem{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 强化皮肤
func (c *Client) OnReqEnhanceDestinySkin(f func(ctx Context, msg *pb.RspEnhanceDestinySkin)error) {
	c.RegisterHandler(658, func(ctx Context) error {
		msg := &pb.RspEnhanceDestinySkin{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 穿戴皮肤
func (c *Client) OnReqWearDestinySkin(f func(ctx Context, msg *pb.RspWearDestinySkin)error) {
	c.RegisterHandler(659, func(ctx Context) error {
		msg := &pb.RspWearDestinySkin{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 切换挚友联动形象
func (c *Client) OnDestinySwitchLinkageSkinReq(f func(ctx Context, msg *pb.DestinySwitchLinkageSkinResp)error) {
	c.RegisterHandler(660, func(ctx Context) error {
		msg := &pb.DestinySwitchLinkageSkinResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 筋斗云数据同步
func (c *Client) OnPlayerCloudDataMsg(f func(ctx Context, msg *pb.PlayerCloudDataMsg)error) {
	c.RegisterHandler(711, func(ctx Context) error {
		msg := &pb.PlayerCloudDataMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 解锁筋斗云请求
func (c *Client) OnUnlockCloudReq(f func(ctx Context, msg *pb.UnlockCloudResp)error) {
	c.RegisterHandler(712, func(ctx Context) error {
		msg := &pb.UnlockCloudResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 筋斗云升级请求
func (c *Client) OnCloudLvUpReq(f func(ctx Context, msg *pb.CloudLvUpResp)error) {
	c.RegisterHandler(713, func(ctx Context) error {
		msg := &pb.CloudLvUpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 筋斗云升阶请求
func (c *Client) OnCloudStageUpReq(f func(ctx Context, msg *pb.CloudStageUpResp)error) {
	c.RegisterHandler(714, func(ctx Context) error {
		msg := &pb.CloudStageUpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 筋斗云穿戴请求
func (c *Client) OnCloudEquipReq(f func(ctx Context, msg *pb.CloudEquipResp)error) {
	c.RegisterHandler(715, func(ctx Context) error {
		msg := &pb.CloudEquipResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 筋斗云皮肤穿戴请求
func (c *Client) OnCloudEquipSkinReq(f func(ctx Context, msg *pb.CloudEquipSkinResp)error) {
	c.RegisterHandler(716, func(ctx Context) error {
		msg := &pb.CloudEquipSkinResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 筋斗云皮肤升级 0-&gt;1 为激活
func (c *Client) OnCloudSkinLvUpReq(f func(ctx Context, msg *pb.CloudSkinLvUpResp)error) {
	c.RegisterHandler(717, func(ctx Context) error {
		msg := &pb.CloudSkinLvUpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步宝地
func (c *Client) OnWildBossDataSync(f func(ctx Context, msg *pb.WildBossDataSync)error) {
	c.RegisterHandler(731, func(ctx Context) error {
		msg := &pb.WildBossDataSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求宝地挑战
func (c *Client) OnWildBossChallengeReq(f func(ctx Context, msg *pb.WildBossChallengeResp)error) {
	c.RegisterHandler(732, func(ctx Context) error {
		msg := &pb.WildBossChallengeResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求宝地扫荡
func (c *Client) OnWildBossRepeatReq(f func(ctx Context, msg *pb.WildBossRepeatResp)error) {
	c.RegisterHandler(733, func(ctx Context) error {
		msg := &pb.WildBossRepeatResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步玩家灵兽数据
func (c *Client) OnPlayerPetDataSync(f func(ctx Context, msg *pb.PlayerPetDataSync)error) {
	c.RegisterHandler(740, func(ctx Context) error {
		msg := &pb.PlayerPetDataSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 捕捉灵兽请求
func (c *Client) OnCatchPetReq(f func(ctx Context, msg *pb.CatchPetResp)error) {
	c.RegisterHandler(741, func(ctx Context) error {
		msg := &pb.CatchPetResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 刷新灵兽池子请求
func (c *Client) OnRefreshPetPoolReq(f func(ctx Context, msg *pb.RefreshPetPoolResp)error) {
	c.RegisterHandler(742, func(ctx Context) error {
		msg := &pb.RefreshPetPoolResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 开启灵兽背包格子请求
func (c *Client) OnAddPetBagCountReq(f func(ctx Context, msg *pb.AddPetBagCountResp)error) {
	c.RegisterHandler(743, func(ctx Context) error {
		msg := &pb.AddPetBagCountResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 灵兽放生请求
func (c *Client) OnReleasePetReq(f func(ctx Context, msg *pb.ReleasePetResp)error) {
	c.RegisterHandler(744, func(ctx Context) error {
		msg := &pb.ReleasePetResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 灵兽升级请求
func (c *Client) OnPetLevelUpReq(f func(ctx Context, msg *pb.PetLevelUpResp)error) {
	c.RegisterHandler(745, func(ctx Context) error {
		msg := &pb.PetLevelUpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 灵兽吞噬请求
func (c *Client) OnPetGobbleUpReq(f func(ctx Context, msg *pb.PetGobbleUpResp)error) {
	c.RegisterHandler(746, func(ctx Context) error {
		msg := &pb.PetGobbleUpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 携带灵兽请求
func (c *Client) OnEquipPetReq(f func(ctx Context, msg *pb.EquipPetResp)error) {
	c.RegisterHandler(747, func(ctx Context) error {
		msg := &pb.EquipPetResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 携带灵兽请求
func (c *Client) OnPetResetReq(f func(ctx Context, msg *pb.PetResetResp)error) {
	c.RegisterHandler(748, func(ctx Context) error {
		msg := &pb.PetResetResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 灵兽心愿选择
func (c *Client) OnChooseWishPetReq(f func(ctx Context, msg *pb.ChooseWishPetResp)error) {
	c.RegisterHandler(749, func(ctx Context) error {
		msg := &pb.ChooseWishPetResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 灵兽协助
func (c *Client) OnPetAssistantReq(f func(ctx Context, msg *pb.PetAssistantResp)error) {
	c.RegisterHandler(750, func(ctx Context) error {
		msg := &pb.PetAssistantResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 灵兽塑魂
func (c *Client) OnPetSoulShapeReq(f func(ctx Context, msg *pb.PetSoulShapeResp)error) {
	c.RegisterHandler(751, func(ctx Context) error {
		msg := &pb.PetSoulShapeResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 灵兽塑魂技能选择
func (c *Client) OnSelectSoulShapeSkillReq(f func(ctx Context, msg *pb.SelectSoulShapeSkillResp)error) {
	c.RegisterHandler(752, func(ctx Context) error {
		msg := &pb.SelectSoulShapeSkillResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 灵兽洗练界面
func (c *Client) OnPetSkillRefreshViewReq(f func(ctx Context, msg *pb.PetSkillRefreshViewResp)error) {
	c.RegisterHandler(753, func(ctx Context) error {
		msg := &pb.PetSkillRefreshViewResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 灵兽洗练
func (c *Client) OnPetSkillRefreshReq(f func(ctx Context, msg *pb.PetSkillRefreshResp)error) {
	c.RegisterHandler(754, func(ctx Context) error {
		msg := &pb.PetSkillRefreshResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 灵兽洗练结果处理
func (c *Client) OnPetSkillRefreshResultReq(f func(ctx Context, msg *pb.PetSkillRefreshResultResp)error) {
	c.RegisterHandler(755, func(ctx Context) error {
		msg := &pb.PetSkillRefreshResultResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 灵兽加锁
func (c *Client) OnPetLockStateReq(f func(ctx Context, msg *pb.PetLockStateResp)error) {
	c.RegisterHandler(756, func(ctx Context) error {
		msg := &pb.PetLockStateResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 切换灵兽联动形象
func (c *Client) OnPetSwitchLinkageSkinReq(f func(ctx Context, msg *pb.PetSwitchLinkageSkinResp)error) {
	c.RegisterHandler(757, func(ctx Context) error {
		msg := &pb.PetSwitchLinkageSkinResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 灵兽新版上锁（需要密码解锁）
func (c *Client) OnPetLockReq(f func(ctx Context, msg *pb.PetLockResp)error) {
	c.RegisterHandler(758, func(ctx Context) error {
		msg := &pb.PetLockResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步数据
func (c *Client) OnTowerDataMsg(f func(ctx Context, msg *pb.TowerDataMsg)error) {
	c.RegisterHandler(761, func(ctx Context) error {
		msg := &pb.TowerDataMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 挑战
func (c *Client) OnTowerChallengeResp(f func(ctx Context, msg *pb.TowerChallengeResp)error) {
	c.RegisterHandler(762, func(ctx Context) error {
		msg := &pb.TowerChallengeResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 快速挑战
func (c *Client) OnQuickChallengeResp(f func(ctx Context, msg *pb.QuickChallengeResp)error) {
	c.RegisterHandler(763, func(ctx Context) error {
		msg := &pb.QuickChallengeResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 选择加成请求
func (c *Client) OnSelectBuffReq(f func(ctx Context, msg *pb.SelectBuffResp)error) {
	c.RegisterHandler(764, func(ctx Context) error {
		msg := &pb.SelectBuffResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看怪物属性
func (c *Client) OnViewMonsterAttrResp20765(f func(ctx Context, msg *pb.ViewMonsterAttrResp)error) {
	c.RegisterHandler(765, func(ctx Context) error {
		msg := &pb.ViewMonsterAttrResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取选择选项
func (c *Client) OnGetSelectOptionReq(f func(ctx Context, msg *pb.GetSelectOptionResp)error) {
	c.RegisterHandler(766, func(ctx Context) error {
		msg := &pb.GetSelectOptionResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 保存选择选项
func (c *Client) OnSaveSelectOptionReq(f func(ctx Context, msg *pb.SaveSelectOptionResp)error) {
	c.RegisterHandler(767, func(ctx Context) error {
		msg := &pb.SaveSelectOptionResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取排行列表
func (c *Client) OnGetRankListReq(f func(ctx Context, msg *pb.GetRankListResp)error) {
	c.RegisterHandler(801, func(ctx Context) error {
		msg := &pb.GetRankListResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 首充领取奖励
func (c *Client) OnGetMallRewardReq(f func(ctx Context, msg *pb.GetMallRewardResp)error) {
	c.RegisterHandler(802, func(ctx Context) error {
		msg := &pb.GetMallRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看联盟积分
func (c *Client) OnGetUnionMemberScoreReq(f func(ctx Context, msg *pb.GetUnionMemberScoreRsp)error) {
	c.RegisterHandler(803, func(ctx Context) error {
		msg := &pb.GetUnionMemberScoreRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步玩家精怪数据
func (c *Client) OnSpiritPlayerDataMsg(f func(ctx Context, msg *pb.SpiritPlayerDataMsg)error) {
	c.RegisterHandler(821, func(ctx Context) error {
		msg := &pb.SpiritPlayerDataMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 精怪抽取请求
func (c *Client) OnSpiritDrawReq(f func(ctx Context, msg *pb.SpiritDrawResp)error) {
	c.RegisterHandler(822, func(ctx Context) error {
		msg := &pb.SpiritDrawResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 精怪解锁请求
func (c *Client) OnSpiritUnlockReq(f func(ctx Context, msg *pb.SpiritUnlockResp)error) {
	c.RegisterHandler(823, func(ctx Context) error {
		msg := &pb.SpiritUnlockResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 精怪升级请求
func (c *Client) OnSpiritLvUpReq(f func(ctx Context, msg *pb.SpiritLvUpResp)error) {
	c.RegisterHandler(824, func(ctx Context) error {
		msg := &pb.SpiritLvUpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 切换出战队伍请求
func (c *Client) OnSwitchBattleTeamReq(f func(ctx Context, msg *pb.SwitchBattleTeamResp)error) {
	c.RegisterHandler(825, func(ctx Context) error {
		msg := &pb.SwitchBattleTeamResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 精怪上下阵请求
func (c *Client) OnSpiritBattleReq(f func(ctx Context, msg *pb.SpiritBattleResp)error) {
	c.RegisterHandler(826, func(ctx Context) error {
		msg := &pb.SpiritBattleResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 精怪共鸣等级提升
func (c *Client) OnSpiritCombineLvUpReq(f func(ctx Context, msg *pb.SpiritCombineLvUpResp)error) {
	c.RegisterHandler(827, func(ctx Context) error {
		msg := &pb.SpiritCombineLvUpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 切换精怪联动形象
func (c *Client) OnSpiritSwitchLinkageSkinReq(f func(ctx Context, msg *pb.SpiritSwitchLinkageSkinResp)error) {
	c.RegisterHandler(828, func(ctx Context) error {
		msg := &pb.SpiritSwitchLinkageSkinResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 精怪等级解锁展示
func (c *Client) OnSpiritLevelUnlockShowReq(f func(ctx Context, msg *pb.SpiritLevelUnlockShowResp)error) {
	c.RegisterHandler(829, func(ctx Context) error {
		msg := &pb.SpiritLevelUnlockShowResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 活动通用数据同步
func (c *Client) OnPushActivityList(f func(ctx Context, msg *pb.PushActivityList)error) {
	c.RegisterHandler(1001, func(ctx Context) error {
		msg := &pb.PushActivityList{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步活动详细配置
func (c *Client) OnActivityCommonDataListSync(f func(ctx Context, msg *pb.ActivityCommonDataListSync)error) {
	c.RegisterHandler(1002, func(ctx Context) error {
		msg := &pb.ActivityCommonDataListSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求活动通用数据
func (c *Client) OnReqGetActivityDetail(f func(ctx Context, msg *pb.RspGetActivityDetail)error) {
	c.RegisterHandler(1003, func(ctx Context) error {
		msg := &pb.RspGetActivityDetail{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取活动任务条件奖励
func (c *Client) OnReqGetActivityConditionReward21004(f func(ctx Context, msg *pb.RspGetActivityConditionReward)error) {
	c.RegisterHandler(1004, func(ctx Context) error {
		msg := &pb.RspGetActivityConditionReward{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 购买活动商品
func (c *Client) OnReqBuyActivityGoods(f func(ctx Context, msg *pb.RspBuyActivityGoods)error) {
	c.RegisterHandler(1005, func(ctx Context) error {
		msg := &pb.RspBuyActivityGoods{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取活动排行奖励
func (c *Client) OnReqGetActivityRankReward(f func(ctx Context, msg *pb.RespGetActivityRankReward)error) {
	c.RegisterHandler(1006, func(ctx Context) error {
		msg := &pb.RespGetActivityRankReward{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步活动任务条件数据 增量
func (c *Client) OnActivityConditionDataListSync(f func(ctx Context, msg *pb.ActivityConditionDataListSync)error) {
	c.RegisterHandler(1007, func(ctx Context) error {
		msg := &pb.ActivityConditionDataListSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步活动商品购买次数列表 增量
func (c *Client) OnActivityMallBuyCountDataListSync(f func(ctx Context, msg *pb.ActivityMallBuyCountDataListSync)error) {
	c.RegisterHandler(1008, func(ctx Context) error {
		msg := &pb.ActivityMallBuyCountDataListSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步活动用户数据
func (c *Client) OnActivityPlayerDataSync(f func(ctx Context, msg *pb.ActivityPlayerDataSync)error) {
	c.RegisterHandler(1009, func(ctx Context) error {
		msg := &pb.ActivityPlayerDataSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取基金活动任务条件奖励
func (c *Client) OnReqGetFundsConditionReward(f func(ctx Context, msg *pb.RspGetFundsConditionReward)error) {
	c.RegisterHandler(1010, func(ctx Context) error {
		msg := &pb.RspGetFundsConditionReward{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 运势抽奖请求
func (c *Client) OnLuckyDrawReq(f func(ctx Context, msg *pb.LuckyDrawResp)error) {
	c.RegisterHandler(1013, func(ctx Context) error {
		msg := &pb.LuckyDrawResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 自选礼包选择道具
func (c *Client) OnSelectItemsReq(f func(ctx Context, msg *pb.SelectItemsResp)error) {
	c.RegisterHandler(1014, func(ctx Context) error {
		msg := &pb.SelectItemsResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步活动积分数据
func (c *Client) OnActivityScoreDataMsgSync(f func(ctx Context, msg *pb.ActivityScoreDataMsgSync)error) {
	c.RegisterHandler(1015, func(ctx Context) error {
		msg := &pb.ActivityScoreDataMsgSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 活动分享任务完成请求
func (c *Client) OnReqShareTaskDone21016(f func(ctx Context, msg *pb.RespShareTaskDone)error) {
	c.RegisterHandler(1016, func(ctx Context) error {
		msg := &pb.RespShareTaskDone{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求榜单状态
func (c *Client) OnReqGetActivityRankState(f func(ctx Context, msg *pb.RespGetActivityRankState)error) {
	c.RegisterHandler(1017, func(ctx Context) error {
		msg := &pb.RespGetActivityRankState{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 通过数组一次性领取活动任务条件奖励
func (c *Client) OnReqGetActivityConditionRewardByArr(f func(ctx Context, msg *pb.RspGetActivityConditionReward)error) {
	c.RegisterHandler(1022, func(ctx Context) error {
		msg := &pb.RspGetActivityConditionReward{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖市觅宝抽奖
func (c *Client) OnSeekTreasureDrawReq(f func(ctx Context, msg *pb.SeekTreasureDrawResp)error) {
	c.RegisterHandler(1023, func(ctx Context) error {
		msg := &pb.SeekTreasureDrawResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖市觅宝界面
func (c *Client) OnSeekTreasureViewReq(f func(ctx Context, msg *pb.SeekTreasureViewResp)error) {
	c.RegisterHandler(1024, func(ctx Context) error {
		msg := &pb.SeekTreasureViewResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 选择稀有奖励
func (c *Client) OnSeekTreasureChooseRareRewardReq(f func(ctx Context, msg *pb.SeekTreasureChooseRareRewardResp)error) {
	c.RegisterHandler(1025, func(ctx Context) error {
		msg := &pb.SeekTreasureChooseRareRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步妖市觅宝抽奖公告
func (c *Client) OnSeekTreasureNoticeSyncInfo(f func(ctx Context, msg *pb.SeekTreasureNoticeSyncInfo)error) {
	c.RegisterHandler(1026, func(ctx Context) error {
		msg := &pb.SeekTreasureNoticeSyncInfo{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 活动商买自选商品列表同步
func (c *Client) OnActivityMallSelectDataListSync(f func(ctx Context, msg *pb.ActivityMallSelectDataListSync)error) {
	c.RegisterHandler(1027, func(ctx Context) error {
		msg := &pb.ActivityMallSelectDataListSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 选择稀有奖励
func (c *Client) OnSeekTreasureSelectRewardReqMsg(f func(ctx Context, msg *pb.SeekTreasureSelectRewardRespMsg)error) {
	c.RegisterHandler(1028, func(ctx Context) error {
		msg := &pb.SeekTreasureSelectRewardRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 论坛分享
func (c *Client) OnReqShareTaskDone21030(f func(ctx Context, msg *pb.RespShareTaskDone)error) {
	c.RegisterHandler(1030, func(ctx Context) error {
		msg := &pb.RespShareTaskDone{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 游戏圈分享
func (c *Client) OnReqShareTaskDone21031(f func(ctx Context, msg *pb.RespShareTaskDone)error) {
	c.RegisterHandler(1031, func(ctx Context) error {
		msg := &pb.RespShareTaskDone{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 自选商品选中协议
func (c *Client) OnSelectActivityGoodsReq(f func(ctx Context, msg *pb.SelectActivityGoodsResp)error) {
	c.RegisterHandler(1032, func(ctx Context) error {
		msg := &pb.SelectActivityGoodsResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取活动战令任务奖励
func (c *Client) OnReqGetActivityConditionReward21033(f func(ctx Context, msg *pb.RspGetActivityConditionReward)error) {
	c.RegisterHandler(1033, func(ctx Context) error {
		msg := &pb.RspGetActivityConditionReward{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 活动任务自选商品
func (c *Client) OnSelectActivityConditionGoodsReq(f func(ctx Context, msg *pb.SelectActivityConditionGoodsResp)error) {
	c.RegisterHandler(1034, func(ctx Context) error {
		msg := &pb.SelectActivityConditionGoodsResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步活动任务奖励自选记录
func (c *Client) OnActivityConditionSelectDataListSync(f func(ctx Context, msg *pb.ActivityConditionSelectDataListSync)error) {
	c.RegisterHandler(1035, func(ctx Context) error {
		msg := &pb.ActivityConditionSelectDataListSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求竞猜数据
func (c *Client) OnReqGuessInfoLoadMsg(f func(ctx Context, msg *pb.RespGuessInfoLoadMsg)error) {
	c.RegisterHandler(1040, func(ctx Context) error {
		msg := &pb.RespGuessInfoLoadMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 竞猜押注
func (c *Client) OnReqGuessMsg(f func(ctx Context, msg *pb.RespGuessMsg)error) {
	c.RegisterHandler(1041, func(ctx Context) error {
		msg := &pb.RespGuessMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取竞猜奖励
func (c *Client) OnReqGuessRewardMsg(f func(ctx Context, msg *pb.RespGuessRewardMsg)error) {
	c.RegisterHandler(1042, func(ctx Context) error {
		msg := &pb.RespGuessRewardMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看竞猜奖励名单
func (c *Client) OnReqGuessResultDetailMsg(f func(ctx Context, msg *pb.RespGuessResultDetailMsg)error) {
	c.RegisterHandler(1043, func(ctx Context) error {
		msg := &pb.RespGuessResultDetailMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 活动屏蔽同步
func (c *Client) OnActivityShieldSync(f func(ctx Context, msg *pb.ActivityShieldSync)error) {
	c.RegisterHandler(1044, func(ctx Context) error {
		msg := &pb.ActivityShieldSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取跨服获得内置的 妖盟充值活动的奖励
func (c *Client) OnReceiveCrossUnionRechargeRewardReq(f func(ctx Context, msg *pb.ReceiveCrossUnionRechargeRewardResp)error) {
	c.RegisterHandler(1045, func(ctx Context) error {
		msg := &pb.ReceiveCrossUnionRechargeRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看自己所属妖盟的参与名单
func (c *Client) OnActivityGetJoinMemberListReq(f func(ctx Context, msg *pb.ActivityGetJoinMemberListRsp)error) {
	c.RegisterHandler(1046, func(ctx Context) error {
		msg := &pb.ActivityGetJoinMemberListRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 从PC打开
func (c *Client) OnReqShareTaskDone21047(f func(ctx Context, msg *pb.RespShareTaskDone)error) {
	c.RegisterHandler(1047, func(ctx Context) error {
		msg := &pb.RespShareTaskDone{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步信息
func (c *Client) OnSyncHomelandMsg(f func(ctx Context, msg *pb.SyncHomelandMsg)error) {
	c.RegisterHandler(1051, func(ctx Context) error {
		msg := &pb.SyncHomelandMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入福地
func (c *Client) OnHomelandEnterReq(f func(ctx Context, msg *pb.HomelandEnterResp)error) {
	c.RegisterHandler(1052, func(ctx Context) error {
		msg := &pb.HomelandEnterResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 福地管理界面
func (c *Client) OnHomelandManageReq(f func(ctx Context, msg *pb.HomelandManageResp)error) {
	c.RegisterHandler(1053, func(ctx Context) error {
		msg := &pb.HomelandManageResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 福地升级
func (c *Client) OnHomelandLevelUpReq(f func(ctx Context, msg *pb.HomelandLevelUpResp)error) {
	c.RegisterHandler(1054, func(ctx Context) error {
		msg := &pb.HomelandLevelUpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 福地刷新资源
func (c *Client) OnHomelandRefreshReq(f func(ctx Context, msg *pb.HomelandRefreshResp)error) {
	c.RegisterHandler(1055, func(ctx Context) error {
		msg := &pb.HomelandRefreshResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 福地购买鼠宝
func (c *Client) OnHomelandBuyWorkerReq(f func(ctx Context, msg *pb.HomelandBuyWorkerResp)error) {
	c.RegisterHandler(1056, func(ctx Context) error {
		msg := &pb.HomelandBuyWorkerResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 福地日志
func (c *Client) OnHomelandGetRecordReq(f func(ctx Context, msg *pb.HomelandGetRecordResp)error) {
	c.RegisterHandler(1057, func(ctx Context) error {
		msg := &pb.HomelandGetRecordResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 福地探寻
func (c *Client) OnHomelandExploreReq(f func(ctx Context, msg *pb.HomelandExploreResp)error) {
	c.RegisterHandler(1058, func(ctx Context) error {
		msg := &pb.HomelandExploreResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 福地探寻刷新
func (c *Client) OnHomelandRefreshExploreReq(f func(ctx Context, msg *pb.HomelandRefreshExploreResp)error) {
	c.RegisterHandler(1059, func(ctx Context) error {
		msg := &pb.HomelandRefreshExploreResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 福地派遣鼠宝
func (c *Client) OnHomelandDispatchWorkerReq(f func(ctx Context, msg *pb.HomelandDispatchWorkerResp)error) {
	c.RegisterHandler(1060, func(ctx Context) error {
		msg := &pb.HomelandDispatchWorkerResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 福地派遣鼠宝预览
func (c *Client) OnHomelandDispatchPreviewReq(f func(ctx Context, msg *pb.HomelandDispatchPreviewResp)error) {
	c.RegisterHandler(1061, func(ctx Context) error {
		msg := &pb.HomelandDispatchPreviewResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步有奖励消息
func (c *Client) OnSyncHomelandHasRewardMsg(f func(ctx Context, msg *pb.SyncHomelandHasRewardMsg)error) {
	c.RegisterHandler(1062, func(ctx Context) error {
		msg := &pb.SyncHomelandHasRewardMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 离开福地
func (c *Client) OnHomelandLeaveReq(f func(ctx Context, msg *pb.HomelandLeaveResp)error) {
	c.RegisterHandler(1063, func(ctx Context) error {
		msg := &pb.HomelandLeaveResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步福地数据
func (c *Client) OnSyncPlayerHomelandChangeMsg(f func(ctx Context, msg *pb.SyncPlayerHomelandChangeMsg)error) {
	c.RegisterHandler(1064, func(ctx Context) error {
		msg := &pb.SyncPlayerHomelandChangeMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 福地获取奖励
func (c *Client) OnHomelandGetRewardReq(f func(ctx Context, msg *pb.HomelandGetRewardResp)error) {
	c.RegisterHandler(1065, func(ctx Context) error {
		msg := &pb.HomelandGetRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 福地聚宝盆数据同步
func (c *Client) OnSyncPlayerCornucopiaChangeMsg(f func(ctx Context, msg *pb.SyncPlayerCornucopiaChangeMsg)error) {
	c.RegisterHandler(1066, func(ctx Context) error {
		msg := &pb.SyncPlayerCornucopiaChangeMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 福地聚宝盆领奖
func (c *Client) OnCornucopiaGetRewardReq(f func(ctx Context, msg *pb.CornucopiaGetRewardRsp)error) {
	c.RegisterHandler(1067, func(ctx Context) error {
		msg := &pb.CornucopiaGetRewardRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 福地鼠鼠自动采集
func (c *Client) OnHomelandAutoDispatchWorkerReq(f func(ctx Context, msg *pb.HomelandAutoDispatchWorkerRsp)error) {
	c.RegisterHandler(1068, func(ctx Context) error {
		msg := &pb.HomelandAutoDispatchWorkerRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 福地鼠宝总管试用
func (c *Client) OnHomelandUseFreeMouseManagerReq(f func(ctx Context, msg *pb.HomelandUseFreeMouseManagerRsp)error) {
	c.RegisterHandler(1069, func(ctx Context) error {
		msg := &pb.HomelandUseFreeMouseManagerRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 鼠鼠总管保存设置
func (c *Client) OnHomelandSaveSettingsReq(f func(ctx Context, msg *pb.HomelandSaveSettingsRsp)error) {
	c.RegisterHandler(1070, func(ctx Context) error {
		msg := &pb.HomelandSaveSettingsRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 桃源一键召回小妖
func (c *Client) OnHomelandDispatchAllWorkerReq(f func(ctx Context, msg *pb.HomelandDispatchAllWorkerResp)error) {
	c.RegisterHandler(1071, func(ctx Context) error {
		msg := &pb.HomelandDispatchAllWorkerResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 福地通知玩家离开(只有偷其他玩家福地时才会通知)
func (c *Client) OnSyncLeaveHomeLand(f func(ctx Context, msg *pb.SyncLeaveHomeLand)error) {
	c.RegisterHandler(1072, func(ctx Context) error {
		msg := &pb.SyncLeaveHomeLand{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求报名
func (c *Client) OnBallGVGApplicationReqMsg(f func(ctx Context, msg *pb.BallGVGApplicationRespMsg)error) {
	c.RegisterHandler(1101, func(ctx Context) error {
		msg := &pb.BallGVGApplicationRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求报名期战力排名
func (c *Client) OnBallGVGAbilityRankNumReqMsg(f func(ctx Context, msg *pb.BallGVGAbilityRankNumRespMsg)error) {
	c.RegisterHandler(1102, func(ctx Context) error {
		msg := &pb.BallGVGAbilityRankNumRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入pvp活动请求协议
func (c *Client) OnBallGVGEnterActivityReqMsg(f func(ctx Context, msg *pb.BallGVGEnterActivityRespMsg)error) {
	c.RegisterHandler(1103, func(ctx Context) error {
		msg := &pb.BallGVGEnterActivityRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入宝地(房间)请求协议
func (c *Client) OnBallGVGEnterPlaceReqMsg(f func(ctx Context, msg *pb.BallGVGEnterPlaceRespMsg)error) {
	c.RegisterHandler(1104, func(ctx Context) error {
		msg := &pb.BallGVGEnterPlaceRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 离开当前房间请求协议
func (c *Client) OnBallGVGLeavePlaceReqMsg(f func(ctx Context, msg *pb.BallGVGLeavePlaceRespMsg)error) {
	c.RegisterHandler(1105, func(ctx Context) error {
		msg := &pb.BallGVGLeavePlaceRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 移动协议
func (c *Client) OnBallGVGMoveReqMsg(f func(ctx Context, msg *pb.BallGVGMoveRespMsg)error) {
	c.RegisterHandler(1106, func(ctx Context) error {
		msg := &pb.BallGVGMoveRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 攻击协议
func (c *Client) OnBallGVGAttackReqMsg(f func(ctx Context, msg *pb.BallGVGAttackRespMsg)error) {
	c.RegisterHandler(1107, func(ctx Context) error {
		msg := &pb.BallGVGAttackRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 任命宗主权限
func (c *Client) OnBallGVGAppointMajorUserReqMsg(f func(ctx Context, msg *pb.BallGVGAppointMajorUserRespMsg)error) {
	c.RegisterHandler(1108, func(ctx Context) error {
		msg := &pb.BallGVGAppointMajorUserRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 标记地点(重复标记视为取消标记)
func (c *Client) OnBallGVGMarkPlaceReqMsg(f func(ctx Context, msg *pb.BallGVGMarkPlaceRespMsg)error) {
	c.RegisterHandler(1109, func(ctx Context) error {
		msg := &pb.BallGVGMarkPlaceRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取所有房间数据
func (c *Client) OnBallGVGGetPlaceInfoReqMsg(f func(ctx Context, msg *pb.BallGVGGetPlaceInfoRespMsg)error) {
	c.RegisterHandler(1111, func(ctx Context) error {
		msg := &pb.BallGVGGetPlaceInfoRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求宗门玩家列表
func (c *Client) OnBallGVGGetUserRankReqMsg(f func(ctx Context, msg *pb.BallGVGGetUserRankRespMsg)error) {
	c.RegisterHandler(1113, func(ctx Context) error {
		msg := &pb.BallGVGGetUserRankRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 被打同步（有人被打即时同步）
func (c *Client) OnBallGVGAttackedSyncMsg(f func(ctx Context, msg *pb.BallGVGAttackedSyncMsg)error) {
	c.RegisterHandler(1116, func(ctx Context) error {
		msg := &pb.BallGVGAttackedSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 积分同步
func (c *Client) OnBallGVGCampScoreSyncMsg(f func(ctx Context, msg *pb.BallGVGCampScoreSyncMsg)error) {
	c.RegisterHandler(1117, func(ctx Context) error {
		msg := &pb.BallGVGCampScoreSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 宝地信息同步（位置同步）
func (c *Client) OnBallGVGPlaceSyncMsg(f func(ctx Context, msg *pb.BallGVGPlaceSyncMsg)error) {
	c.RegisterHandler(1118, func(ctx Context) error {
		msg := &pb.BallGVGPlaceSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 公告推送
func (c *Client) OnBallGVGNoticeSyncMsg(f func(ctx Context, msg *pb.BallGVGNoticeSyncMsg)error) {
	c.RegisterHandler(1119, func(ctx Context) error {
		msg := &pb.BallGVGNoticeSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步宗门信息（任命、标记）
func (c *Client) OnBallGVGCampSyncMsg(f func(ctx Context, msg *pb.BallGVGCampSyncMsg)error) {
	c.RegisterHandler(1120, func(ctx Context) error {
		msg := &pb.BallGVGCampSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 被打同步最新用户信息(通知被打人)
func (c *Client) OnBallGVGAttackedUserSyncMsg(f func(ctx Context, msg *pb.BallGVGAttackedUserSyncMsg)error) {
	c.RegisterHandler(1121, func(ctx Context) error {
		msg := &pb.BallGVGAttackedUserSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 房间抢占信息同步(人数变动，进房间，出房间  成功占领)
func (c *Client) OnBallGVGPlaceSeizeSyncMsg(f func(ctx Context, msg *pb.BallGVGPlaceSeizeSyncMsg)error) {
	c.RegisterHandler(1122, func(ctx Context) error {
		msg := &pb.BallGVGPlaceSeizeSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进房间同步信息（有人进房间及时同步）
func (c *Client) OnBallGVGEnterPlaceSync(f func(ctx Context, msg *pb.BallGVGEnterPlaceSync)error) {
	c.RegisterHandler(1123, func(ctx Context) error {
		msg := &pb.BallGVGEnterPlaceSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 出房间同步信息（有人出房间及时同步）
func (c *Client) OnBallGVGLeavePlaceSync(f func(ctx Context, msg *pb.BallGVGLeavePlaceSync)error) {
	c.RegisterHandler(1124, func(ctx Context) error {
		msg := &pb.BallGVGLeavePlaceSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求报名期妖盟战力排行
func (c *Client) OnBallGVGAbilityRankReqMsg(f func(ctx Context, msg *pb.BallGVGAbilityRankRespMsg)error) {
	c.RegisterHandler(1126, func(ctx Context) error {
		msg := &pb.BallGVGAbilityRankRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 战斗结算同步
func (c *Client) OnBallGVGEndDataSyncMsg(f func(ctx Context, msg *pb.BallGVGEndDataSyncMsg)error) {
	c.RegisterHandler(1127, func(ctx Context) error {
		msg := &pb.BallGVGEndDataSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取参与妖盟的名字
func (c *Client) OnBallGVGUnionNameListReqMsg(f func(ctx Context, msg *pb.BallGVGUnionNameListRespMsg)error) {
	c.RegisterHandler(1128, func(ctx Context) error {
		msg := &pb.BallGVGUnionNameListRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取上古遗迹妖盟成员信息
func (c *Client) OnBallGVGGetUnionUserListReqMsg(f func(ctx Context, msg *pb.BallGVGGetUnionUserListRespMsg)error) {
	c.RegisterHandler(1129, func(ctx Context) error {
		msg := &pb.BallGVGGetUnionUserListRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 上古遗迹红点
func (c *Client) OnBallGVGRedDotReqMsg(f func(ctx Context, msg *pb.BallGVGRedDotRespMsg)error) {
	c.RegisterHandler(1131, func(ctx Context) error {
		msg := &pb.BallGVGRedDotRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步表情包数据
func (c *Client) OnPlayerEmoticonDataSync(f func(ctx Context, msg *pb.PlayerEmoticonDataSync)error) {
	c.RegisterHandler(1201, func(ctx Context) error {
		msg := &pb.PlayerEmoticonDataSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 异兽入侵挑战
func (c *Client) OnInvadeChallengeResp(f func(ctx Context, msg *pb.InvadeChallengeResp)error) {
	c.RegisterHandler(1401, func(ctx Context) error {
		msg := &pb.InvadeChallengeResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 异兽入侵数据同步
func (c *Client) OnInvadeDataMsg(f func(ctx Context, msg *pb.InvadeDataMsg)error) {
	c.RegisterHandler(1402, func(ctx Context) error {
		msg := &pb.InvadeDataMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 仙缘进度领奖
func (c *Client) OnXyFundGetRewardReq(f func(ctx Context, msg *pb.XyFundGetRewardResp)error) {
	c.RegisterHandler(1421, func(ctx Context) error {
		msg := &pb.XyFundGetRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取商会同名姓名列表
func (c *Client) OnReqUnionSameNameListMsg(f func(ctx Context, msg *pb.RspUnionSameNameListMsg)error) {
	c.RegisterHandler(1501, func(ctx Context) error {
		msg := &pb.RspUnionSameNameListMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取商会充值列表
func (c *Client) OnUnionRechargeUserReqMsg(f func(ctx Context, msg *pb.UnionRechargeUserRspMsg)error) {
	c.RegisterHandler(1511, func(ctx Context) error {
		msg := &pb.UnionRechargeUserRspMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 商会充值活动数据同步
func (c *Client) OnUnionRechargeDataSync(f func(ctx Context, msg *pb.UnionRechargeDataSync)error) {
	c.RegisterHandler(1512, func(ctx Context) error {
		msg := &pb.UnionRechargeDataSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看本分组妖盟详情
func (c *Client) OnReqUnionDetail201604(f func(ctx Context, msg *pb.RspUnionDetail)error) {
	c.RegisterHandler(1604, func(ctx Context) error {
		msg := &pb.RspUnionDetail{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取分组评级排行
func (c *Client) OnUnionPlayerGradeRankReqMsg(f func(ctx Context, msg *pb.UnionPlayerGradeRankRespMsg)error) {
	c.RegisterHandler(1630, func(ctx Context) error {
		msg := &pb.UnionPlayerGradeRankRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 评级人数分配设置
func (c *Client) OnUnionMemberGradeSettingReq(f func(ctx Context, msg *pb.UnionMemberGradeSettingResp)error) {
	c.RegisterHandler(1650, func(ctx Context) error {
		msg := &pb.UnionMemberGradeSettingResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取妖盟跨服信息
func (c *Client) OnGetCrossUnionGroupServersReq(f func(ctx Context, msg *pb.GetCrossUnionGroupServersResp)error) {
	c.RegisterHandler(1651, func(ctx Context) error {
		msg := &pb.GetCrossUnionGroupServersResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 发布招募信息
func (c *Client) OnUnionPublishRecruitReqMsg(f func(ctx Context, msg *pb.UnionPublishRecruitRespMsg)error) {
	c.RegisterHandler(1654, func(ctx Context) error {
		msg := &pb.UnionPublishRecruitRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步招募次数
func (c *Client) OnNotifyUnionRecruitCountMsg(f func(ctx Context, msg *pb.NotifyUnionRecruitCountMsg)error) {
	c.RegisterHandler(1655, func(ctx Context) error {
		msg := &pb.NotifyUnionRecruitCountMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 登录或者评级变化时 同步玩家的评级数据
func (c *Client) OnNotifyPlayerGradeMsg(f func(ctx Context, msg *pb.NotifyPlayerGradeMsg)error) {
	c.RegisterHandler(1656, func(ctx Context) error {
		msg := &pb.NotifyPlayerGradeMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 创建妖盟
func (c *Client) OnReqUnionCreate(f func(ctx Context, msg *pb.RspUnionCreate)error) {
	c.RegisterHandler(2101, func(ctx Context) error {
		msg := &pb.RspUnionCreate{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入妖盟
func (c *Client) OnReqUnionEnter(f func(ctx Context, msg *pb.RspUnionEnter)error) {
	c.RegisterHandler(2102, func(ctx Context) error {
		msg := &pb.RspUnionEnter{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取妖盟列表
func (c *Client) OnReqUnionList(f func(ctx Context, msg *pb.RspUnionList)error) {
	c.RegisterHandler(2103, func(ctx Context) error {
		msg := &pb.RspUnionList{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 加入妖盟
func (c *Client) OnReqUnionJoin(f func(ctx Context, msg *pb.RspUnionJoin)error) {
	c.RegisterHandler(2104, func(ctx Context) error {
		msg := &pb.RspUnionJoin{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 随机加入妖盟
func (c *Client) OnReqUnionRandomJoin(f func(ctx Context, msg *pb.RspUnionRandomJoin)error) {
	c.RegisterHandler(2105, func(ctx Context) error {
		msg := &pb.RspUnionRandomJoin{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟详情
func (c *Client) OnReqUnionDetail22106(f func(ctx Context, msg *pb.RspUnionDetail)error) {
	c.RegisterHandler(2106, func(ctx Context) error {
		msg := &pb.RspUnionDetail{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 职位变更
func (c *Client) OnReqUnionPosition(f func(ctx Context, msg *pb.RspUnionPosition)error) {
	c.RegisterHandler(2107, func(ctx Context) error {
		msg := &pb.RspUnionPosition{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取申请加入妖盟的玩家列表
func (c *Client) OnReqUnionApplyPlayerList(f func(ctx Context, msg *pb.RspUnionApplyPlayerList)error) {
	c.RegisterHandler(2108, func(ctx Context) error {
		msg := &pb.RspUnionApplyPlayerList{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 设置随机加入状态
func (c *Client) OnReqUnionRandomState(f func(ctx Context, msg *pb.RspUnionRandomState)error) {
	c.RegisterHandler(2109, func(ctx Context) error {
		msg := &pb.RspUnionRandomState{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 设置入盟限制等级
func (c *Client) OnReqUnionLimitRealmsId(f func(ctx Context, msg *pb.RspUnionLimitRealmsId)error) {
	c.RegisterHandler(2110, func(ctx Context) error {
		msg := &pb.RspUnionLimitRealmsId{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 清理所有申请玩家
func (c *Client) OnReqUnionClearApply(f func(ctx Context, msg *pb.RspUnionClearApply)error) {
	c.RegisterHandler(2111, func(ctx Context) error {
		msg := &pb.RspUnionClearApply{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 通过新玩家的申请
func (c *Client) OnReqUnionPassApply(f func(ctx Context, msg *pb.RspUnionPassApply)error) {
	c.RegisterHandler(2112, func(ctx Context) error {
		msg := &pb.RspUnionPassApply{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 解散妖盟
func (c *Client) OnReqUnionRemove(f func(ctx Context, msg *pb.RspUnionRemove)error) {
	c.RegisterHandler(2113, func(ctx Context) error {
		msg := &pb.RspUnionRemove{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 退出妖盟
func (c *Client) OnReqUnionExit(f func(ctx Context, msg *pb.RspUnionExit)error) {
	c.RegisterHandler(2114, func(ctx Context) error {
		msg := &pb.RspUnionExit{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取妖盟动态
func (c *Client) OnReqUnionTrends(f func(ctx Context, msg *pb.RspUnionTrends)error) {
	c.RegisterHandler(2115, func(ctx Context) error {
		msg := &pb.RspUnionTrends{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟信息修改
func (c *Client) OnReqUnionModify(f func(ctx Context, msg *pb.RspUnionModify)error) {
	c.RegisterHandler(2116, func(ctx Context) error {
		msg := &pb.RspUnionModify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取每日任务列表
func (c *Client) OnReqUnionDailyTask(f func(ctx Context, msg *pb.RspUnionDailyTask)error) {
	c.RegisterHandler(2117, func(ctx Context) error {
		msg := &pb.RspUnionDailyTask{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取每日任务奖励
func (c *Client) OnReqUnionGetDailyTask(f func(ctx Context, msg *pb.RspUnionGetDailyTask)error) {
	c.RegisterHandler(2118, func(ctx Context) error {
		msg := &pb.RspUnionGetDailyTask{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 推送被踢出妖盟广播
func (c *Client) OnPushUnionEvictionMsg(f func(ctx Context, msg *pb.PushUnionEvictionMsg)error) {
	c.RegisterHandler(2119, func(ctx Context) error {
		msg := &pb.PushUnionEvictionMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 每日任务捐献
func (c *Client) OnReqUnionDailyDonate(f func(ctx Context, msg *pb.RspUnionDailyDonate)error) {
	c.RegisterHandler(2120, func(ctx Context) error {
		msg := &pb.RspUnionDailyDonate{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取每日任务成员完成进度
func (c *Client) OnReqUnionDailyProgress(f func(ctx Context, msg *pb.RspUnionDailyProgress)error) {
	c.RegisterHandler(2121, func(ctx Context) error {
		msg := &pb.RspUnionDailyProgress{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 逐出妖盟
func (c *Client) OnReqUnionEviction(f func(ctx Context, msg *pb.RspUnionEviction)error) {
	c.RegisterHandler(2122, func(ctx Context) error {
		msg := &pb.RspUnionEviction{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 加入的妖盟列表(我没有妖盟时的列表)
func (c *Client) OnReqUnionJoinList(f func(ctx Context, msg *pb.RspUnionJoinList)error) {
	c.RegisterHandler(2123, func(ctx Context) error {
		msg := &pb.RspUnionJoinList{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 推送我的妖盟数据更新
func (c *Client) OnMyUnionData(f func(ctx Context, msg *pb.MyUnionData)error) {
	c.RegisterHandler(2124, func(ctx Context) error {
		msg := &pb.MyUnionData{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取妖盟协助列表
func (c *Client) OnGetUnionHelpDataListReq(f func(ctx Context, msg *pb.GetUnionHelpDataListResp)error) {
	c.RegisterHandler(2125, func(ctx Context) error {
		msg := &pb.GetUnionHelpDataListResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求妖盟协助
func (c *Client) OnRequestUnionHelpReq(f func(ctx Context, msg *pb.RequestUnionHelpResp)error) {
	c.RegisterHandler(2126, func(ctx Context) error {
		msg := &pb.RequestUnionHelpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 执行妖盟协助
func (c *Client) OnUnionHelpReq(f func(ctx Context, msg *pb.UnionHelpResp)error) {
	c.RegisterHandler(2127, func(ctx Context) error {
		msg := &pb.UnionHelpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取妖盟协助状态
func (c *Client) OnGetUnionHelpStateReq(f func(ctx Context, msg *pb.GetUnionHelpStateResp)error) {
	c.RegisterHandler(2128, func(ctx Context) error {
		msg := &pb.GetUnionHelpStateResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看妖盟的红点
func (c *Client) OnReqUnionWatchRedDot(f func(ctx Context, msg *pb.RspUnionWatchRedDot)error) {
	c.RegisterHandler(2129, func(ctx Context) error {
		msg := &pb.RspUnionWatchRedDot{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 砍价数据同步
func (c *Client) OnCutPriceDataMsg(f func(ctx Context, msg *pb.CutPriceDataMsg)error) {
	c.RegisterHandler(2165, func(ctx Context) error {
		msg := &pb.CutPriceDataMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 砍价
func (c *Client) OnCutPriceBargainReqMsg(f func(ctx Context, msg *pb.CutPriceBargainRespMsg)error) {
	c.RegisterHandler(2166, func(ctx Context) error {
		msg := &pb.CutPriceBargainRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 砍价购买
func (c *Client) OnCutPriceBuyReqMsg(f func(ctx Context, msg *pb.CutPriceBuyRespMsg)error) {
	c.RegisterHandler(2167, func(ctx Context) error {
		msg := &pb.CutPriceBuyRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟时间同步
func (c *Client) OnUnionTimeMsg(f func(ctx Context, msg *pb.UnionTimeMsg)error) {
	c.RegisterHandler(2168, func(ctx Context) error {
		msg := &pb.UnionTimeMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 绑定邀请码
func (c *Client) OnReqBindInviteCode(f func(ctx Context, msg *pb.RspBindInviteCode)error) {
	c.RegisterHandler(2401, func(ctx Context) error {
		msg := &pb.RspBindInviteCode{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取绑定他人的邀请码
func (c *Client) OnReqGetBindInviteCode(f func(ctx Context, msg *pb.RspGetBindInviteCode)error) {
	c.RegisterHandler(2402, func(ctx Context) error {
		msg := &pb.RspGetBindInviteCode{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 移除黑名单
func (c *Client) OnRemoveBlackReqMsg(f func(ctx Context, msg *pb.RemoveBlackRespMsg)error) {
	c.RegisterHandler(3205, func(ctx Context) error {
		msg := &pb.RemoveBlackRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取黑名单列表
func (c *Client) OnGetBlackPlayerListReqMsg(f func(ctx Context, msg *pb.GetBlackPlayerListRespMsg)error) {
	c.RegisterHandler(3206, func(ctx Context) error {
		msg := &pb.GetBlackPlayerListRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 黑名单列表同步
func (c *Client) OnSyncBlackPlayerIdListMsg(f func(ctx Context, msg *pb.SyncBlackPlayerIdListMsg)error) {
	c.RegisterHandler(3207, func(ctx Context) error {
		msg := &pb.SyncBlackPlayerIdListMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 检测消息是否有效
func (c *Client) OnWorldMessageCheckEffectiveReq(f func(ctx Context, msg *pb.WorldMessageCheckEffectiveRsp)error) {
	c.RegisterHandler(3208, func(ctx Context) error {
		msg := &pb.WorldMessageCheckEffectiveRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入群英榜
func (c *Client) OnRspHeroRankEnter(f func(ctx Context, msg *pb.RspHeroRankEnter)error) {
	c.RegisterHandler(3700, func(ctx Context) error {
		msg := &pb.RspHeroRankEnter{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步群英榜玩家信息
func (c *Client) OnSynHeroRankPlayerInfo(f func(ctx Context, msg *pb.SynHeroRankPlayerInfo)error) {
	c.RegisterHandler(3701, func(ctx Context) error {
		msg := &pb.SynHeroRankPlayerInfo{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求挑战者列表
func (c *Client) OnReqHeroRankFightPlayerList(f func(ctx Context, msg *pb.RspHeroRankFightPlayerList)error) {
	c.RegisterHandler(3702, func(ctx Context) error {
		msg := &pb.RspHeroRankFightPlayerList{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求挑战玩家
func (c *Client) OnReqHeroRankFight(f func(ctx Context, msg *pb.RspHeroRankFight)error) {
	c.RegisterHandler(3703, func(ctx Context) error {
		msg := &pb.RspHeroRankFight{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求扫荡玩家
func (c *Client) OnReqHeroRankClear(f func(ctx Context, msg *pb.RspHeroRankClear)error) {
	c.RegisterHandler(3704, func(ctx Context) error {
		msg := &pb.RspHeroRankClear{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求日志数据
func (c *Client) OnRspHeroRankRecord(f func(ctx Context, msg *pb.RspHeroRankRecord)error) {
	c.RegisterHandler(3705, func(ctx Context) error {
		msg := &pb.RspHeroRankRecord{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求领取成就奖励
func (c *Client) OnReqHeroRankGetAchieve(f func(ctx Context, msg *pb.RspHeroRankGetAchieve)error) {
	c.RegisterHandler(3706, func(ctx Context) error {
		msg := &pb.RspHeroRankGetAchieve{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求购买体力
func (c *Client) OnReqHeroRankBuyEnergy(f func(ctx Context, msg *pb.RspHeroRankBuyEnergy)error) {
	c.RegisterHandler(3707, func(ctx Context) error {
		msg := &pb.RspHeroRankBuyEnergy{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 抽奖
func (c *Client) OnReqPetDreamLandDraw(f func(ctx Context, msg *pb.RspPetDreamLandDraw)error) {
	c.RegisterHandler(4101, func(ctx Context) error {
		msg := &pb.RspPetDreamLandDraw{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 区域解锁
func (c *Client) OnReqPetDreamLandAdventurePlaceUnlock(f func(ctx Context, msg *pb.RspPetDreamLandAdventurePlaceUnlock)error) {
	c.RegisterHandler(4102, func(ctx Context) error {
		msg := &pb.RspPetDreamLandAdventurePlaceUnlock{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 坑位解锁
func (c *Client) OnReqPetDreamLandAdventureSlotUnlock(f func(ctx Context, msg *pb.RspPetDreamLandAdventurePlaceSlot)error) {
	c.RegisterHandler(4103, func(ctx Context) error {
		msg := &pb.RspPetDreamLandAdventurePlaceSlot{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 灵兽派遣
func (c *Client) OnReqPetDreamLandAdventureDispatch(f func(ctx Context, msg *pb.RspPetDreamLandAdventureDispatch)error) {
	c.RegisterHandler(4104, func(ctx Context) error {
		msg := &pb.RspPetDreamLandAdventureDispatch{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取收益
func (c *Client) OnReqPetDreamLandAdventureGetAward(f func(ctx Context, msg *pb.RspPetDreamLandAdventureGetAward)error) {
	c.RegisterHandler(4105, func(ctx Context) error {
		msg := &pb.RspPetDreamLandAdventureGetAward{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 刷新邪祟
func (c *Client) OnRefreshEvilThingReq(f func(ctx Context, msg *pb.RefreshEvilThingResp)error) {
	c.RegisterHandler(4201, func(ctx Context) error {
		msg := &pb.RefreshEvilThingResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 镇压邪祟
func (c *Client) OnBattleEviThingReq(f func(ctx Context, msg *pb.BattleEvilThingResp)error) {
	c.RegisterHandler(4202, func(ctx Context) error {
		msg := &pb.BattleEvilThingResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 神通数据同步
func (c *Client) OnPlayerMagicDataMsg(f func(ctx Context, msg *pb.PlayerMagicDataMsg)error) {
	c.RegisterHandler(4400, func(ctx Context) error {
		msg := &pb.PlayerMagicDataMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 装备神通
func (c *Client) OnMagicEquipReq(f func(ctx Context, msg *pb.MagicEquipResp)error) {
	c.RegisterHandler(4401, func(ctx Context) error {
		msg := &pb.MagicEquipResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 重置神通
func (c *Client) OnMagicResetReq(f func(ctx Context, msg *pb.MagicResetResp)error) {
	c.RegisterHandler(4402, func(ctx Context) error {
		msg := &pb.MagicResetResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进阶神通
func (c *Client) OnMagicStageUpReq(f func(ctx Context, msg *pb.MagicStageUpResp)error) {
	c.RegisterHandler(4403, func(ctx Context) error {
		msg := &pb.MagicStageUpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 升级神通
func (c *Client) OnMagicLvUpReq(f func(ctx Context, msg *pb.MagicLvUpResp)error) {
	c.RegisterHandler(4404, func(ctx Context) error {
		msg := &pb.MagicLvUpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 装备印记
func (c *Client) OnMagicEquipMarkReq(f func(ctx Context, msg *pb.MagicEquipMarkResp)error) {
	c.RegisterHandler(4405, func(ctx Context) error {
		msg := &pb.MagicEquipMarkResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 卸下印记
func (c *Client) OnMagicUnsnatchMarkReq(f func(ctx Context, msg *pb.MagicUnsnatchMarkResp)error) {
	c.RegisterHandler(4406, func(ctx Context) error {
		msg := &pb.MagicUnsnatchMarkResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 融合神通印记
func (c *Client) OnMagicFusionMarkReq(f func(ctx Context, msg *pb.MagicFusionMarkResp)error) {
	c.RegisterHandler(4407, func(ctx Context) error {
		msg := &pb.MagicFusionMarkResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 衍化
func (c *Client) OnMagicDerivationReq(f func(ctx Context, msg *pb.MagicDerivationResp)error) {
	c.RegisterHandler(4408, func(ctx Context) error {
		msg := &pb.MagicDerivationResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 升级组合
func (c *Client) OnMagicCombineLvUpReq(f func(ctx Context, msg *pb.MagicCombineLvUpResp)error) {
	c.RegisterHandler(4409, func(ctx Context) error {
		msg := &pb.MagicCombineLvUpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 切换预设
func (c *Client) OnMagicSwitchPresetsReq(f func(ctx Context, msg *pb.MagicSwitchPresetsResp)error) {
	c.RegisterHandler(4410, func(ctx Context) error {
		msg := &pb.MagicSwitchPresetsResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 一键突破神通
func (c *Client) OnMagicStageUpAllReq(f func(ctx Context, msg *pb.MagicStageUpResp)error) {
	c.RegisterHandler(4411, func(ctx Context) error {
		msg := &pb.MagicStageUpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取预设神通列表
func (c *Client) OnGetPresetMagicInfoReq(f func(ctx Context, msg *pb.GetPresetMagicInfoResp)error) {
	c.RegisterHandler(4412, func(ctx Context) error {
		msg := &pb.GetPresetMagicInfoResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 拆卸装备的印记
func (c *Client) OnUnEquipStoneReq(f func(ctx Context, msg *pb.UnEquipStoneResp)error) {
	c.RegisterHandler(4413, func(ctx Context) error {
		msg := &pb.UnEquipStoneResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 添加储存数据
func (c *Client) OnSaveToServiceReq(f func(ctx Context, msg *pb.SaveToServiceRsp)error) {
	c.RegisterHandler(4501, func(ctx Context) error {
		msg := &pb.SaveToServiceRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 全量同步储存数据
func (c *Client) OnAllClientDataSync(f func(ctx Context, msg *pb.AllClientDataSync)error) {
	c.RegisterHandler(4502, func(ctx Context) error {
		msg := &pb.AllClientDataSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步灵兽蛋数据
func (c *Client) OnHoldPetEggDataSync(f func(ctx Context, msg *pb.HoldPetEggDataSync)error) {
	c.RegisterHandler(4600, func(ctx Context) error {
		msg := &pb.HoldPetEggDataSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步太初秘册数据
func (c *Client) OnHoldNestedBoxDataSync(f func(ctx Context, msg *pb.HoldNestedBoxDataSync)error) {
	c.RegisterHandler(4601, func(ctx Context) error {
		msg := &pb.HoldNestedBoxDataSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 仙宫数据同步
func (c *Client) OnPalaceSyncMsg(f func(ctx Context, msg *pb.PalaceSyncMsg)error) {
	c.RegisterHandler(4801, func(ctx Context) error {
		msg := &pb.PalaceSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 仙宫点赞同步
func (c *Client) OnPalaceWorshipReq(f func(ctx Context, msg *pb.PalaceWorshipRsp)error) {
	c.RegisterHandler(4802, func(ctx Context) error {
		msg := &pb.PalaceWorshipRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 仙宫外部数据请求
func (c *Client) OnEnterPalaceReq(f func(ctx Context, msg *pb.EnterPalaceRsp)error) {
	c.RegisterHandler(4803, func(ctx Context) error {
		msg := &pb.EnterPalaceRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 仙宫内部数据请求
func (c *Client) OnEnterPalaceInnerReq(f func(ctx Context, msg *pb.EnterPalaceInnerRsp)error) {
	c.RegisterHandler(4804, func(ctx Context) error {
		msg := &pb.EnterPalaceInnerRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 仙名录数据请求
func (c *Client) OnPalaceInnerBookReq(f func(ctx Context, msg *pb.PalaceInnerBookRsp)error) {
	c.RegisterHandler(4805, func(ctx Context) error {
		msg := &pb.PalaceInnerBookRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 执行仙宫送福
func (c *Client) OnSendGiftReq(f func(ctx Context, msg *pb.SendGiftRsp)error) {
	c.RegisterHandler(4806, func(ctx Context) error {
		msg := &pb.SendGiftRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 仙宫送福领奖
func (c *Client) OnGetSendGiftRewardReq(f func(ctx Context, msg *pb.GetSendGiftRewardRsp)error) {
	c.RegisterHandler(4807, func(ctx Context) error {
		msg := &pb.GetSendGiftRewardRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 仙宫送福数据同步
func (c *Client) OnSendGiftSyncMsg(f func(ctx Context, msg *pb.SendGiftSyncMsg)error) {
	c.RegisterHandler(4808, func(ctx Context) error {
		msg := &pb.SendGiftSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 仙宫神迹同步
func (c *Client) OnPalaceMiracleDataMsg(f func(ctx Context, msg *pb.PalaceMiracleDataMsg)error) {
	c.RegisterHandler(4809, func(ctx Context) error {
		msg := &pb.PalaceMiracleDataMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获得当前已开启的天宫称号id
func (c *Client) OnGetTitleIdListReq(f func(ctx Context, msg *pb.GetTitleIdListRsp)error) {
	c.RegisterHandler(4810, func(ctx Context) error {
		msg := &pb.GetTitleIdListRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 仙宫获得成就奖励
func (c *Client) OnPalaceAchievementGetRewardReq(f func(ctx Context, msg *pb.PalaceAchievementGetRewardRsp)error) {
	c.RegisterHandler(4857, func(ctx Context) error {
		msg := &pb.PalaceAchievementGetRewardRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 进入活动页
func (c *Client) OnManHuangEnterPanelReqMsg(f func(ctx Context, msg *pb.ManHuangEnterPanelRespMsg)error) {
	c.RegisterHandler(5001, func(ctx Context) error {
		msg := &pb.ManHuangEnterPanelRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 创建队伍
func (c *Client) OnManHuangCreateTeamReqMsg(f func(ctx Context, msg *pb.ManHuangCreateTeamRespMsg)error) {
	c.RegisterHandler(5002, func(ctx Context) error {
		msg := &pb.ManHuangCreateTeamRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 获取队伍列表
func (c *Client) OnManHuangGetTeamListReqMsg(f func(ctx Context, msg *pb.ManHuangGetTeamListRespMsg)error) {
	c.RegisterHandler(5003, func(ctx Context) error {
		msg := &pb.ManHuangGetTeamListRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 查看队伍数据
func (c *Client) OnManHuangGetTeamInfoReqMsg(f func(ctx Context, msg *pb.ManHuangGetTeamInfoRespMsg)error) {
	c.RegisterHandler(5004, func(ctx Context) error {
		msg := &pb.ManHuangGetTeamInfoRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 加入队伍
func (c *Client) OnManHuangJoinTeamReqMsg(f func(ctx Context, msg *pb.ManHuangJoinTeamRespMsg)error) {
	c.RegisterHandler(5005, func(ctx Context) error {
		msg := &pb.ManHuangJoinTeamRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 取消申请队伍
func (c *Client) OnManHuangCancelTeamApplyReqMsg(f func(ctx Context, msg *pb.ManHuangCancelTeamApplyRespMsg)error) {
	c.RegisterHandler(5006, func(ctx Context) error {
		msg := &pb.ManHuangCancelTeamApplyRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 解散/退出队伍
func (c *Client) OnManHuangQuitTeamReqMsg(f func(ctx Context, msg *pb.ManHuangQuitTeamRespMsg)error) {
	c.RegisterHandler(5007, func(ctx Context) error {
		msg := &pb.ManHuangQuitTeamRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 退出活动
func (c *Client) OnManHuangLeavePanelReqMsg(f func(ctx Context, msg *pb.ManHuangLeavePanelRespMsg)error) {
	c.RegisterHandler(5008, func(ctx Context) error {
		msg := &pb.ManHuangLeavePanelRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 玩家申请入队同意
func (c *Client) OnManHuangApplyJoinTeamAgreeReqMsg(f func(ctx Context, msg *pb.ManHuangApplyJoinTeamAgreeRespMsg)error) {
	c.RegisterHandler(5009, func(ctx Context) error {
		msg := &pb.ManHuangApplyJoinTeamAgreeRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 玩家申请入队一键拒绝
func (c *Client) OnManHuangApplyJoinTeamRefusedReqMsg(f func(ctx Context, msg *pb.ManHuangApplyJoinTeamRefusedRespMsg)error) {
	c.RegisterHandler(5010, func(ctx Context) error {
		msg := &pb.ManHuangApplyJoinTeamRefusedRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 踢出队伍
func (c *Client) OnManHuangKickOutTeamReqMsg(f func(ctx Context, msg *pb.ManHuangKickOutTeamRespMsg)error) {
	c.RegisterHandler(5011, func(ctx Context) error {
		msg := &pb.ManHuangKickOutTeamRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 进入区域
func (c *Client) OnManHuangEnterRegionReqMsg(f func(ctx Context, msg *pb.ManHuangEnterRegionRespMsg)error) {
	c.RegisterHandler(5012, func(ctx Context) error {
		msg := &pb.ManHuangEnterRegionRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 探索
func (c *Client) OnManHuangExploreReqMsg(f func(ctx Context, msg *pb.ManHuangExploreRespMsg)error) {
	c.RegisterHandler(5013, func(ctx Context) error {
		msg := &pb.ManHuangExploreRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 执行探索事件 
func (c *Client) OnManHuangEventHandleReqMsg(f func(ctx Context, msg *pb.ManHuangEventHandleRespMsg)error) {
	c.RegisterHandler(5014, func(ctx Context) error {
		msg := &pb.ManHuangEventHandleRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 action探索事件(标记，绕过) 
func (c *Client) OnManHuangEventActionReqMsg(f func(ctx Context, msg *pb.ManHuangEventActionRespMsg)error) {
	c.RegisterHandler(5015, func(ctx Context) error {
		msg := &pb.ManHuangEventActionRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 日志-队友互助 
func (c *Client) OnManHuangLogHelpReqMsg(f func(ctx Context, msg *pb.ManHuangLogHelpRespMsg)error) {
	c.RegisterHandler(5016, func(ctx Context) error {
		msg := &pb.ManHuangLogHelpRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 日志-战场日志 
func (c *Client) OnManHuangLogBattlegroundReqMsg(f func(ctx Context, msg *pb.ManHuangLogBattlegroundRespMsg)error) {
	c.RegisterHandler(5017, func(ctx Context) error {
		msg := &pb.ManHuangLogBattlegroundRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 日志-个人日志 
func (c *Client) OnManHuangLogPersonReqMsg(f func(ctx Context, msg *pb.ManHuangLogPersonRespMsg)error) {
	c.RegisterHandler(5018, func(ctx Context) error {
		msg := &pb.ManHuangLogPersonRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 排行-个人 
func (c *Client) OnManHuangRankPersonReqMsg(f func(ctx Context, msg *pb.ManHuangRankPersonRespMsg)error) {
	c.RegisterHandler(5019, func(ctx Context) error {
		msg := &pb.ManHuangRankPersonRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 排行-队伍 
func (c *Client) OnManHuangRankTeamReqMsg(f func(ctx Context, msg *pb.ManHuangRankTeamRespMsg)error) {
	c.RegisterHandler(5020, func(ctx Context) error {
		msg := &pb.ManHuangRankTeamRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 广播-仇人信息
func (c *Client) OnManHuangEnemyNotify(f func(ctx Context, msg *pb.ManHuangEnemyNotify)error) {
	c.RegisterHandler(5022, func(ctx Context) error {
		msg := &pb.ManHuangEnemyNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 广播-队长收到的广播
func (c *Client) OnManHuangTeamLeaderNotify(f func(ctx Context, msg *pb.ManHuangTeamLeaderNotify)error) {
	c.RegisterHandler(5023, func(ctx Context) error {
		msg := &pb.ManHuangTeamLeaderNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 广播-队员接收到的广播
func (c *Client) OnManHuangTeamMemberNotify(f func(ctx Context, msg *pb.ManHuangTeamMemberNotify)error) {
	c.RegisterHandler(5024, func(ctx Context) error {
		msg := &pb.ManHuangTeamMemberNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 恢复体力
func (c *Client) OnManHuangRecoverEnergyReqMsg(f func(ctx Context, msg *pb.ManHuangRecoverEnergyRespMsg)error) {
	c.RegisterHandler(5025, func(ctx Context) error {
		msg := &pb.ManHuangRecoverEnergyRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 执行协助
func (c *Client) OnManHuangHelpActionAttReqMsg(f func(ctx Context, msg *pb.ManHuangHelpActionAttRespMsg)error) {
	c.RegisterHandler(5026, func(ctx Context) error {
		msg := &pb.ManHuangHelpActionAttRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 开宝箱
func (c *Client) OnManHuangOpenBoxReqMsg(f func(ctx Context, msg *pb.ManHuangOpenBoxRespMsg)error) {
	c.RegisterHandler(5027, func(ctx Context) error {
		msg := &pb.ManHuangOpenBoxRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 领取协助奖励
func (c *Client) OnManHuangHelpActionRewardReqMsg(f func(ctx Context, msg *pb.ManHuangHelpActionRewardRespMsg)error) {
	c.RegisterHandler(5028, func(ctx Context) error {
		msg := &pb.ManHuangHelpActionRewardRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 事件广播
func (c *Client) OnManHuangEventInfoNotify(f func(ctx Context, msg *pb.ManHuangEventInfoNotify)error) {
	c.RegisterHandler(5029, func(ctx Context) error {
		msg := &pb.ManHuangEventInfoNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 发送同步玩家信息协议
func (c *Client) OnManHuangGetUserDataReqMsg(f func(ctx Context, msg *pb.ManHuangGetUserDataRespMsg)error) {
	c.RegisterHandler(5030, func(ctx Context) error {
		msg := &pb.ManHuangGetUserDataRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蛮荒 击杀广播
func (c *Client) OnManHuangMarqueeNotify(f func(ctx Context, msg *pb.ManHuangMarqueeNotify)error) {
	c.RegisterHandler(5031, func(ctx Context) error {
		msg := &pb.ManHuangMarqueeNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入妖盟乱斗
func (c *Client) OnUnionBattleEnterReq(f func(ctx Context, msg *pb.UnionBattleEnterRsp)error) {
	c.RegisterHandler(5301, func(ctx Context) error {
		msg := &pb.UnionBattleEnterRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看自己所属联盟的参与名单
func (c *Client) OnUnionBattleGetJoinMemberListReq(f func(ctx Context, msg *pb.UnionBattleGetJoinMemberListRsp)error) {
	c.RegisterHandler(5302, func(ctx Context) error {
		msg := &pb.UnionBattleGetJoinMemberListRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 匹配
func (c *Client) OnUnionBattleMatchReq(f func(ctx Context, msg *pb.UnionBattleMatchRsp)error) {
	c.RegisterHandler(5303, func(ctx Context) error {
		msg := &pb.UnionBattleMatchRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 挑战
func (c *Client) OnUnionBattleChallengeReq(f func(ctx Context, msg *pb.UnionBattleChallengeRsp)error) {
	c.RegisterHandler(5304, func(ctx Context) error {
		msg := &pb.UnionBattleChallengeRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看战报信息
func (c *Client) OnUnionBattleGetReportReq(f func(ctx Context, msg *pb.UnionBattleGetReportRsp)error) {
	c.RegisterHandler(5306, func(ctx Context) error {
		msg := &pb.UnionBattleGetReportRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 购买buff
func (c *Client) OnUnionBattleBuyOpenBuffReq(f func(ctx Context, msg *pb.UnionBattleBuyOpenBuffRsp)error) {
	c.RegisterHandler(5307, func(ctx Context) error {
		msg := &pb.UnionBattleBuyOpenBuffRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 选择buff
func (c *Client) OnUnionBattleSelectBuffReq(f func(ctx Context, msg *pb.UnionBattleSelectBuffRsp)error) {
	c.RegisterHandler(5308, func(ctx Context) error {
		msg := &pb.UnionBattleSelectBuffRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取本次挑战的回放列表
func (c *Client) OnUnionBattleGetPlayBackListReq(f func(ctx Context, msg *pb.UnionBattleGetPlayBackListRsp)error) {
	c.RegisterHandler(5309, func(ctx Context) error {
		msg := &pb.UnionBattleGetPlayBackListRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取联盟成就奖励
func (c *Client) OnUnionBattleReceiveUnionAchieveRewardReq(f func(ctx Context, msg *pb.UnionBattleReceiveUnionAchieveRewardRsp)error) {
	c.RegisterHandler(5310, func(ctx Context) error {
		msg := &pb.UnionBattleReceiveUnionAchieveRewardRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步玩家用户信息
func (c *Client) OnSynPlayerInfo(f func(ctx Context, msg *pb.SynPlayerInfo)error) {
	c.RegisterHandler(5311, func(ctx Context) error {
		msg := &pb.SynPlayerInfo{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看联盟积分
func (c *Client) OnUnionBattleGetUnionMemberScoreReq(f func(ctx Context, msg *pb.UnionBattleGetUnionMemberScoreRsp)error) {
	c.RegisterHandler(5312, func(ctx Context) error {
		msg := &pb.UnionBattleGetUnionMemberScoreRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 红点
func (c *Client) OnUnionBattleRedDotReqMsg(f func(ctx Context, msg *pb.UnionBattleRedDotRespMsg)error) {
	c.RegisterHandler(5313, func(ctx Context) error {
		msg := &pb.UnionBattleRedDotRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 精炼
func (c *Client) OnReqEquipmentRefine(f func(ctx Context, msg *pb.RespEquipmentRefine)error) {
	c.RegisterHandler(5501, func(ctx Context) error {
		msg := &pb.RespEquipmentRefine{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进阶
func (c *Client) OnReqEquipmentAdvance(f func(ctx Context, msg *pb.RespEquipmentAdvance)error) {
	c.RegisterHandler(5502, func(ctx Context) error {
		msg := &pb.RespEquipmentAdvance{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 技能解锁
func (c *Client) OnReqEquipmentActivate(f func(ctx Context, msg *pb.RespEquipmentActivate)error) {
	c.RegisterHandler(5503, func(ctx Context) error {
		msg := &pb.RespEquipmentActivate{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 数据同步
func (c *Client) OnEquipmentAdvanceDataMsg(f func(ctx Context, msg *pb.EquipmentAdvanceDataMsg)error) {
	c.RegisterHandler(5504, func(ctx Context) error {
		msg := &pb.EquipmentAdvanceDataMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入秘境
func (c *Client) OnSecretTowerEnterReq(f func(ctx Context, msg *pb.SecretTowerEnterRsp)error) {
	c.RegisterHandler(5601, func(ctx Context) error {
		msg := &pb.SecretTowerEnterRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 挑战怪物
func (c *Client) OnSecretTowerFightReq(f func(ctx Context, msg *pb.SecretTowerFightResp)error) {
	c.RegisterHandler(5602, func(ctx Context) error {
		msg := &pb.SecretTowerFightResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取阶段奖励
func (c *Client) OnSecretTowerGetStageRewardReq(f func(ctx Context, msg *pb.SecretTowerGetStageRewardRsp)error) {
	c.RegisterHandler(5603, func(ctx Context) error {
		msg := &pb.SecretTowerGetStageRewardRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看怪物属性
func (c *Client) OnSecretTowerViewMonsterAttrReq(f func(ctx Context, msg *pb.SecretTowerViewMonsterAttrResp)error) {
	c.RegisterHandler(5604, func(ctx Context) error {
		msg := &pb.SecretTowerViewMonsterAttrResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 秘境数据同步
func (c *Client) OnSynSecretTowerInfo(f func(ctx Context, msg *pb.SynSecretTowerInfo)error) {
	c.RegisterHandler(5605, func(ctx Context) error {
		msg := &pb.SynSecretTowerInfo{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看秘境成就
func (c *Client) OnSecretTowerAchievementReq(f func(ctx Context, msg *pb.SecretTowerAchievementRsp)error) {
	c.RegisterHandler(5606, func(ctx Context) error {
		msg := &pb.SecretTowerAchievementRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取成就奖励
func (c *Client) OnSecretTowerGetAchiRewardReq(f func(ctx Context, msg *pb.SecretTowerGetAchiRewardRsp)error) {
	c.RegisterHandler(5607, func(ctx Context) error {
		msg := &pb.SecretTowerGetAchiRewardRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入妖邪讨伐
func (c *Client) OnUnionBossInfoReqMsg(f func(ctx Context, msg *pb.UnionBossInfoRespMsg)error) {
	c.RegisterHandler(5801, func(ctx Context) error {
		msg := &pb.UnionBossInfoRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 布阵
func (c *Client) OnUnionBossBuffReqMsg(f func(ctx Context, msg *pb.UnionBossBuffRespMsg)error) {
	c.RegisterHandler(5802, func(ctx Context) error {
		msg := &pb.UnionBossBuffRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看讨伐奖励
func (c *Client) OnUnionBossRewardReqMsg(f func(ctx Context, msg *pb.UnionBossRewardRespMsg)error) {
	c.RegisterHandler(5803, func(ctx Context) error {
		msg := &pb.UnionBossRewardRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取讨伐奖励
func (c *Client) OnUnionBossRewardReceiveReqMsg(f func(ctx Context, msg *pb.UnionBossRewardReceiveRespMsg)error) {
	c.RegisterHandler(5804, func(ctx Context) error {
		msg := &pb.UnionBossRewardReceiveRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 挑战
func (c *Client) OnUnionBossBattleReqMsg(f func(ctx Context, msg *pb.UnionBossBattleRespMsg)error) {
	c.RegisterHandler(5805, func(ctx Context) error {
		msg := &pb.UnionBossBattleRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取讨伐成就奖励
func (c *Client) OnUnionBossAchieveRewardReqMsg(f func(ctx Context, msg *pb.UnionBossAchieveRewardRespMsg)error) {
	c.RegisterHandler(5806, func(ctx Context) error {
		msg := &pb.UnionBossAchieveRewardRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步boss信息
func (c *Client) OnUnionBossMsg(f func(ctx Context, msg *pb.UnionBossMsg)error) {
	c.RegisterHandler(5807, func(ctx Context) error {
		msg := &pb.UnionBossMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步阵法信息
func (c *Client) OnUnionBossBuff(f func(ctx Context, msg *pb.UnionBossBuff)error) {
	c.RegisterHandler(5808, func(ctx Context) error {
		msg := &pb.UnionBossBuff{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求已布阵玩家信息
func (c *Client) OnUnionBossAddBuffPlayerReqMsg(f func(ctx Context, msg *pb.UnionBossAddBuffPlayerRespMsg)error) {
	c.RegisterHandler(5809, func(ctx Context) error {
		msg := &pb.UnionBossAddBuffPlayerRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领奖期 冠军弹窗
func (c *Client) OnAskDingGetChampionInfoReq(f func(ctx Context, msg *pb.AskDingGetChampionInfoResp)error) {
	c.RegisterHandler(5935, func(ctx Context) error {
		msg := &pb.AskDingGetChampionInfoResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 玩家参赛后 同步玩家参赛资格
func (c *Client) OnAskDingSyncPlayerJoinMsg(f func(ctx Context, msg *pb.AskDingSyncPlayerJoinMsg)error) {
	c.RegisterHandler(5938, func(ctx Context) error {
		msg := &pb.AskDingSyncPlayerJoinMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查询据点战基础信息
func (c *Client) OnUnionAreaWarReqMsg(f func(ctx Context, msg *pb.UnionAreaWarRespMsg)error) {
	c.RegisterHandler(6001, func(ctx Context) error {
		msg := &pb.UnionAreaWarRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求查看分组
func (c *Client) OnUnionAreaWarGroupReqMsg(f func(ctx Context, msg *pb.UnionAreaWarGroupRespMsg)error) {
	c.RegisterHandler(6003, func(ctx Context) error {
		msg := &pb.UnionAreaWarGroupRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求防守人列表
func (c *Client) OnUnionAreaWarDefendersReqMsg(f func(ctx Context, msg *pb.UnionAreaWarDefendersRespMsg)error) {
	c.RegisterHandler(6004, func(ctx Context) error {
		msg := &pb.UnionAreaWarDefendersRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 刷新防守人列表
func (c *Client) OnUnionAreaWarDefendersUpdateReqMsg(f func(ctx Context, msg *pb.UnionAreaWarDefendersUpdateRespMsg)error) {
	c.RegisterHandler(6005, func(ctx Context) error {
		msg := &pb.UnionAreaWarDefendersUpdateRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取据点列表
func (c *Client) OnUnionAreWarListReqMsg(f func(ctx Context, msg *pb.UnionAreaWarListRespMsg)error) {
	c.RegisterHandler(6006, func(ctx Context) error {
		msg := &pb.UnionAreaWarListRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 刷新据点列表
func (c *Client) OnUnionAreWarListUpdateReqMsg(f func(ctx Context, msg *pb.UnionAreaWarListUpdateRespMsg)error) {
	c.RegisterHandler(6007, func(ctx Context) error {
		msg := &pb.UnionAreaWarListUpdateRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求战斗场景攻击
func (c *Client) OnUnionAreaWarAttackReqMsg(f func(ctx Context, msg *pb.UnionAreaWarAttackRespMsg)error) {
	c.RegisterHandler(6009, func(ctx Context) error {
		msg := &pb.UnionAreaWarAttackRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 建设(捐献)
func (c *Client) OnUnionAreaWarDevelopReqMsg(f func(ctx Context, msg *pb.UnionAreaWarDevelopRespMsg)error) {
	c.RegisterHandler(6011, func(ctx Context) error {
		msg := &pb.UnionAreaWarDevelopRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 加载宝库信息
func (c *Client) OnUnionAreaWarTreasuryReqMsg(f func(ctx Context, msg *pb.UnionAreaWarTreasuryRespMsg)error) {
	c.RegisterHandler(6013, func(ctx Context) error {
		msg := &pb.UnionAreaWarTreasuryRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 宝库抽奖
func (c *Client) OnUnionAreaWarTreasuryDrawReqMsg(f func(ctx Context, msg *pb.UnionAreaWarTreasuryDrawRespMsg)error) {
	c.RegisterHandler(6014, func(ctx Context) error {
		msg := &pb.UnionAreaWarTreasuryDrawRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求召唤神龙
func (c *Client) OnUnionAreaWarSummonDragonReqMsg(f func(ctx Context, msg *pb.UnionAreaWarSummonDragonRespMsg)error) {
	c.RegisterHandler(6015, func(ctx Context) error {
		msg := &pb.UnionAreaWarSummonDragonRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查询据点战膜拜
func (c *Client) OnUnionAreaWarWorshipReq(f func(ctx Context, msg *pb.UnionAreaWarWorshipRsp)error) {
	c.RegisterHandler(6016, func(ctx Context) error {
		msg := &pb.UnionAreaWarWorshipRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求战斗场景数据
func (c *Client) OnUnionAreaWarFightSceneReqMsg(f func(ctx Context, msg *pb.UnionAreaWarFightSceneRespMsg)error) {
	c.RegisterHandler(6019, func(ctx Context) error {
		msg := &pb.UnionAreaWarFightSceneRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求捐献建设数据
func (c *Client) OnUnionAreaWarConstructReqMsg(f func(ctx Context, msg *pb.UnionAreaWarConstructRespMsg)error) {
	c.RegisterHandler(6020, func(ctx Context) error {
		msg := &pb.UnionAreaWarConstructRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求战报数据
func (c *Client) OnUnionAreaWarReportReqMsg(f func(ctx Context, msg *pb.UnionAreaWarReportRespMsg)error) {
	c.RegisterHandler(6021, func(ctx Context) error {
		msg := &pb.UnionAreaWarReportRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求战报伤害详情数据
func (c *Client) OnUnionAreaWarReportDetailReqMsg(f func(ctx Context, msg *pb.UnionAreaWarReportDetailRespMsg)error) {
	c.RegisterHandler(6022, func(ctx Context) error {
		msg := &pb.UnionAreaWarReportDetailRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求联盟伤害数据
func (c *Client) OnUnionAreaWarUnionDamageReqMsg(f func(ctx Context, msg *pb.UnionAreaWarUnionDamageRespMsg)error) {
	c.RegisterHandler(6023, func(ctx Context) error {
		msg := &pb.UnionAreaWarUnionDamageRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求竞猜界面数据
func (c *Client) OnUnionAreaWarBetDataReqMsg(f func(ctx Context, msg *pb.UnionAreaWarBetDataRespMsg)error) {
	c.RegisterHandler(6024, func(ctx Context) error {
		msg := &pb.UnionAreaWarBetDataRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求竞猜选择
func (c *Client) OnUnionAreaWarBetSelectReqMsg(f func(ctx Context, msg *pb.UnionAreaWarBetSelectRespMsg)error) {
	c.RegisterHandler(6025, func(ctx Context) error {
		msg := &pb.UnionAreaWarBetSelectRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求竞猜奖励
func (c *Client) OnUnionAreaWarBetRewardReqMsg(f func(ctx Context, msg *pb.UnionAreaWarBetRewardRespMsg)error) {
	c.RegisterHandler(6026, func(ctx Context) error {
		msg := &pb.UnionAreaWarBetRewardRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看自己所属联盟的参与名单
func (c *Client) OnUnionAreaWarGetJoinMemberListReq(f func(ctx Context, msg *pb.UnionAreaWarGetJoinMemberListRsp)error) {
	c.RegisterHandler(6027, func(ctx Context) error {
		msg := &pb.UnionAreaWarGetJoinMemberListRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 神龙战场攻击推送
func (c *Client) OnSyncUnionAreaWarDragonAttackMsg(f func(ctx Context, msg *pb.SyncUnionAreaWarDragonAttackMsg)error) {
	c.RegisterHandler(6028, func(ctx Context) error {
		msg := &pb.SyncUnionAreaWarDragonAttackMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求查看联盟分段位伤害信息
func (c *Client) OnUnionAreaWarUnionGroupDamageReqMsg(f func(ctx Context, msg *pb.UnionAreaWarUnionGroupDamageRespMsg)error) {
	c.RegisterHandler(6029, func(ctx Context) error {
		msg := &pb.UnionAreaWarUnionGroupDamageRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求红点信息
func (c *Client) OnUnionAreaWarRedDotReqMsg(f func(ctx Context, msg *pb.UnionAreaWarRedDotRespMsg)error) {
	c.RegisterHandler(6030, func(ctx Context) error {
		msg := &pb.UnionAreaWarRedDotRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求查看竞猜结果的获奖玩家列表
func (c *Client) OnUnionAreaWarGuessPlayersReqMsg(f func(ctx Context, msg *pb.UnionAreaWarGuessPlayersRespMsg)error) {
	c.RegisterHandler(6031, func(ctx Context) error {
		msg := &pb.UnionAreaWarGuessPlayersRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看所有段位的分组数量
func (c *Client) OnUnionAreaWarDanGroupCountReq(f func(ctx Context, msg *pb.UnionAreaWarDanGroupCountResp)error) {
	c.RegisterHandler(6032, func(ctx Context) error {
		msg := &pb.UnionAreaWarDanGroupCountResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入活动
func (c *Client) OnAskWayEnterReq(f func(ctx Context, msg *pb.AskWayEnterRsp)error) {
	c.RegisterHandler(6101, func(ctx Context) error {
		msg := &pb.AskWayEnterRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 斗法匹配
func (c *Client) OnAskWayMatchReq(f func(ctx Context, msg *pb.AskWayMatchRsp)error) {
	c.RegisterHandler(6102, func(ctx Context) error {
		msg := &pb.AskWayMatchRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 斗法战斗
func (c *Client) OnAskWayBattleReq(f func(ctx Context, msg *pb.AskWayBattleRsp)error) {
	c.RegisterHandler(6103, func(ctx Context) error {
		msg := &pb.AskWayBattleRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取战报
func (c *Client) OnAskWayGetReportReq(f func(ctx Context, msg *pb.AskWayGetReportRsp)error) {
	c.RegisterHandler(6104, func(ctx Context) error {
		msg := &pb.AskWayGetReportRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 斗法战斗回放
func (c *Client) OnAskWayBattleReplyReq(f func(ctx Context, msg *pb.AskWayBattleReplyRsp)error) {
	c.RegisterHandler(6105, func(ctx Context) error {
		msg := &pb.AskWayBattleReplyRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取段位奖励
func (c *Client) OnAskWayReceiveTierRewardReq(f func(ctx Context, msg *pb.AskWayReceiveTierRewardRsp)error) {
	c.RegisterHandler(6106, func(ctx Context) error {
		msg := &pb.AskWayReceiveTierRewardRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取斗法榜单积分奖励
func (c *Client) OnAskWayReceiveScoreRewardReq(f func(ctx Context, msg *pb.AskWayReceiveScoreRewardRsp)error) {
	c.RegisterHandler(6107, func(ctx Context) error {
		msg := &pb.AskWayReceiveScoreRewardRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 购买挑战令道具
func (c *Client) OnAskWayBuyFightTicketReq(f func(ctx Context, msg *pb.AskWayBuyFightTicketRsp)error) {
	c.RegisterHandler(6108, func(ctx Context) error {
		msg := &pb.AskWayBuyFightTicketRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看玩家数据
func (c *Client) OnAskWayGetPlayerDetailReq(f func(ctx Context, msg *pb.AskWayGetPlayerDetailRsp)error) {
	c.RegisterHandler(6109, func(ctx Context) error {
		msg := &pb.AskWayGetPlayerDetailRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取竞猜币
func (c *Client) OnAskWayGetGuessCoinReq(f func(ctx Context, msg *pb.AskWayGetGuessCoinRsp)error) {
	c.RegisterHandler(6110, func(ctx Context) error {
		msg := &pb.AskWayGetGuessCoinRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 登天竞猜
func (c *Client) OnAskWayToSkyGuessReq(f func(ctx Context, msg *pb.AskWayToSkyGuessRsp)error) {
	c.RegisterHandler(6111, func(ctx Context) error {
		msg := &pb.AskWayToSkyGuessRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 登天战斗回放
func (c *Client) OnAskWayToSkyBattleReplyReq(f func(ctx Context, msg *pb.AskWayToSkyBattleReplyRsp)error) {
	c.RegisterHandler(6112, func(ctx Context) error {
		msg := &pb.AskWayToSkyBattleReplyRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 朝拜
func (c *Client) OnAskWayWorshipReq(f func(ctx Context, msg *pb.AskWayWorshipRsp)error) {
	c.RegisterHandler(6113, func(ctx Context) error {
		msg := &pb.AskWayWorshipRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取登天榜单奖励
func (c *Client) OnAskWayReceiveToSkyRankRewardReq(f func(ctx Context, msg *pb.AskWayReceiveToSkyRankRewardRsp)error) {
	c.RegisterHandler(6114, func(ctx Context) error {
		msg := &pb.AskWayReceiveToSkyRankRewardRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 往届妖神榜
func (c *Client) OnAskWayToSkyRosterReq(f func(ctx Context, msg *pb.AskWayToSkyRosterRsp)error) {
	c.RegisterHandler(6115, func(ctx Context) error {
		msg := &pb.AskWayToSkyRosterRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取登天期组别数据
func (c *Client) OnAskWayToSkyGetBattleResultReq(f func(ctx Context, msg *pb.AskWayToSkyGetBattleResultRsp)error) {
	c.RegisterHandler(6116, func(ctx Context) error {
		msg := &pb.AskWayToSkyGetBattleResultRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取竞猜信息
func (c *Client) OnAskWayGetGuessInfoReq(f func(ctx Context, msg *pb.AskWayGetGuessInfoRsp)error) {
	c.RegisterHandler(6117, func(ctx Context) error {
		msg := &pb.AskWayGetGuessInfoRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取登天战报数据
func (c *Client) OnAskWayToSkyGetReportReq(f func(ctx Context, msg *pb.AskWayToSkyGetReportRsp)error) {
	c.RegisterHandler(6118, func(ctx Context) error {
		msg := &pb.AskWayToSkyGetReportRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取海报数据
func (c *Client) OnAskWayGetGetCurStateInfoReq(f func(ctx Context, msg *pb.AskWayGetGetCurStateInfoRsp)error) {
	c.RegisterHandler(6132, func(ctx Context) error {
		msg := &pb.AskWayGetGetCurStateInfoRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 玩家法宝数据
func (c *Client) OnMagicTreasurePlayerDataMsg(f func(ctx Context, msg *pb.MagicTreasurePlayerDataMsg)error) {
	c.RegisterHandler(6301, func(ctx Context) error {
		msg := &pb.MagicTreasurePlayerDataMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 法宝抽取
func (c *Client) OnMagicTreasureDrawReq(f func(ctx Context, msg *pb.MagicTreasureDrawRsp)error) {
	c.RegisterHandler(6302, func(ctx Context) error {
		msg := &pb.MagicTreasureDrawRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 法宝升级
func (c *Client) OnMagicTreasureLvUpReq(f func(ctx Context, msg *pb.MagicTreasureLvUpRsp)error) {
	c.RegisterHandler(6303, func(ctx Context) error {
		msg := &pb.MagicTreasureLvUpRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 法宝激活
func (c *Client) OnMagicTreasureActiveReq(f func(ctx Context, msg *pb.MagicTreasureActiveRsp)error) {
	c.RegisterHandler(6304, func(ctx Context) error {
		msg := &pb.MagicTreasureActiveRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 法宝升星
func (c *Client) OnMagicTreasureStarUpReq(f func(ctx Context, msg *pb.MagicTreasureStarUpRsp)error) {
	c.RegisterHandler(6305, func(ctx Context) error {
		msg := &pb.MagicTreasureStarUpRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步法宝列表数据
func (c *Client) OnSyncMagicTreasureDataMsg(f func(ctx Context, msg *pb.SyncMagicTreasureDataMsg)error) {
	c.RegisterHandler(6306, func(ctx Context) error {
		msg := &pb.SyncMagicTreasureDataMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 法宝形象切换
func (c *Client) OnMagicTreasureSwitchLinkageSkinReq(f func(ctx Context, msg *pb.MagicTreasureSwitchLinkageSkinResp)error) {
	c.RegisterHandler(6307, func(ctx Context) error {
		msg := &pb.MagicTreasureSwitchLinkageSkinResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟夺位战主界面
func (c *Client) OnUnionFightMainReq(f func(ctx Context, msg *pb.UnionFightMainRsp)error) {
	c.RegisterHandler(6701, func(ctx Context) error {
		msg := &pb.UnionFightMainRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟夺位战报名
func (c *Client) OnUnionFightApplyReq(f func(ctx Context, msg *pb.UnionFightApplyRsp)error) {
	c.RegisterHandler(6702, func(ctx Context) error {
		msg := &pb.UnionFightApplyRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟夺位战进入战斗界面
func (c *Client) OnUnionFightEnterReq(f func(ctx Context, msg *pb.UnionFightEnterRsp)error) {
	c.RegisterHandler(6703, func(ctx Context) error {
		msg := &pb.UnionFightEnterRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟夺位战请战
func (c *Client) OnUnionFightRequestReq(f func(ctx Context, msg *pb.UnionFightRequestRsp)error) {
	c.RegisterHandler(6704, func(ctx Context) error {
		msg := &pb.UnionFightRequestRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟夺位战上阵
func (c *Client) OnUnionFightPositionReq(f func(ctx Context, msg *pb.UnionFightPositionRsp)error) {
	c.RegisterHandler(6705, func(ctx Context) error {
		msg := &pb.UnionFightPositionRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟夺位战下阵
func (c *Client) OnUnionFightUnPositionReq(f func(ctx Context, msg *pb.UnionFightUnPositionRsp)error) {
	c.RegisterHandler(6706, func(ctx Context) error {
		msg := &pb.UnionFightUnPositionRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟夺位战小组排行
func (c *Client) OnUnionFightGroupRankReq(f func(ctx Context, msg *pb.UnionFightGroupRankRsp)error) {
	c.RegisterHandler(6707, func(ctx Context) error {
		msg := &pb.UnionFightGroupRankRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟夺位战至尊榜
func (c *Client) OnUnionFightSupremacyListReq(f func(ctx Context, msg *pb.UnionFightSupremacyListRsp)error) {
	c.RegisterHandler(6709, func(ctx Context) error {
		msg := &pb.UnionFightSupremacyListRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟夺位战膜拜
func (c *Client) OnUnionFightWorshipReq(f func(ctx Context, msg *pb.UnionFightWorshipRsp)error) {
	c.RegisterHandler(6710, func(ctx Context) error {
		msg := &pb.UnionFightWorshipRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟夺位战回放
func (c *Client) OnUnionFightFightPlaybackReq(f func(ctx Context, msg *pb.UnionFightFightPlaybackRsp)error) {
	c.RegisterHandler(6711, func(ctx Context) error {
		msg := &pb.UnionFightFightPlaybackRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟夺位战对战记录
func (c *Client) OnUnionFightRecordReq(f func(ctx Context, msg *pb.UnionFightRecordRsp)error) {
	c.RegisterHandler(6712, func(ctx Context) error {
		msg := &pb.UnionFightRecordRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟夺位战历史最高
func (c *Client) OnUnionFightGetHistoryDataReq(f func(ctx Context, msg *pb.UnionFightGetHistoryDataRsp)error) {
	c.RegisterHandler(6713, func(ctx Context) error {
		msg := &pb.UnionFightGetHistoryDataRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟夺位战获取锁定属性
func (c *Client) OnUnionFightGetLockedDetailReq(f func(ctx Context, msg *pb.UnionFightGetLockedDetailRsp)error) {
	c.RegisterHandler(6714, func(ctx Context) error {
		msg := &pb.UnionFightGetLockedDetailRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟夺位战设置锁定状态
func (c *Client) OnUnionFightChangeLockStatusReq(f func(ctx Context, msg *pb.UnionFightChangeLockStatusRsp)error) {
	c.RegisterHandler(6715, func(ctx Context) error {
		msg := &pb.UnionFightChangeLockStatusRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟夺位战今日详情
func (c *Client) OnUnionFightGetTodayResultReq(f func(ctx Context, msg *pb.UnionFightGetTodayResultRsp)error) {
	c.RegisterHandler(6716, func(ctx Context) error {
		msg := &pb.UnionFightGetTodayResultRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取妖盟排行列表
func (c *Client) OnUnionFightGetUnionRankListReq(f func(ctx Context, msg *pb.RspUnionList)error) {
	c.RegisterHandler(6717, func(ctx Context) error {
		msg := &pb.RspUnionList{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟夺位战备战数据同步
func (c *Client) OnUnionFightPrepareDataSync(f func(ctx Context, msg *pb.UnionFightPrepareDataSync)error) {
	c.RegisterHandler(6721, func(ctx Context) error {
		msg := &pb.UnionFightPrepareDataSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟夺位战登录数据
func (c *Client) OnUnionFightApplyDataSync(f func(ctx Context, msg *pb.UnionFightApplyDataSync)error) {
	c.RegisterHandler(6730, func(ctx Context) error {
		msg := &pb.UnionFightApplyDataSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 基础信息
func (c *Client) OnGodIslandBaseInfoReqMsg(f func(ctx Context, msg *pb.GodIslandBaseInfoRespMsg)error) {
	c.RegisterHandler(6801, func(ctx Context) error {
		msg := &pb.GodIslandBaseInfoRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看分组
func (c *Client) OnGodIslandGroupReqMsg(f func(ctx Context, msg *pb.GodIslandGroupRespMsg)error) {
	c.RegisterHandler(6802, func(ctx Context) error {
		msg := &pb.GodIslandGroupRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求妖盟建筑战报
func (c *Client) OnGodIslandReportReqMsg(f func(ctx Context, msg *pb.GodIslandReportRespMsg)error) {
	c.RegisterHandler(6803, func(ctx Context) error {
		msg := &pb.GodIslandReportRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求妖盟建筑战报详情数据
func (c *Client) OnGodIslandReportDetailReqMsg(f func(ctx Context, msg *pb.GodIslandReportDetailRespMsg)error) {
	c.RegisterHandler(6804, func(ctx Context) error {
		msg := &pb.GodIslandReportDetailRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求联盟战功数据
func (c *Client) OnGodIslandUnionDamageReqMsg(f func(ctx Context, msg *pb.GodIslandUnionDamageRespMsg)error) {
	c.RegisterHandler(6805, func(ctx Context) error {
		msg := &pb.GodIslandUnionDamageRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看自己所属联盟的参与名单
func (c *Client) OnGodIslandGetJoinMemberListReq(f func(ctx Context, msg *pb.GodIslandGetJoinMemberListRsp)error) {
	c.RegisterHandler(6806, func(ctx Context) error {
		msg := &pb.GodIslandGetJoinMemberListRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟分段位伤害信息
func (c *Client) OnGodIslandUnionGroupDamageReqMsg(f func(ctx Context, msg *pb.GodIslandUnionGroupDamageRespMsg)error) {
	c.RegisterHandler(6807, func(ctx Context) error {
		msg := &pb.GodIslandUnionGroupDamageRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 红点信息
func (c *Client) OnGodIslandRedDotReqMsg(f func(ctx Context, msg *pb.GodIslandRedDotRespMsg)error) {
	c.RegisterHandler(6808, func(ctx Context) error {
		msg := &pb.GodIslandRedDotRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 点赞
func (c *Client) OnGodIslandWorshipReq(f func(ctx Context, msg *pb.GodIslandWorshipResp)error) {
	c.RegisterHandler(6809, func(ctx Context) error {
		msg := &pb.GodIslandWorshipResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求个人战报
func (c *Client) OnGodIslandPlayerReportReqMsg(f func(ctx Context, msg *pb.GodIslandPlayerReportRespMsg)error) {
	c.RegisterHandler(6810, func(ctx Context) error {
		msg := &pb.GodIslandPlayerReportRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖力转灵力
func (c *Client) OnGodIslandUpdatePowerReq(f func(ctx Context, msg *pb.GodIslandUpdatePowerResp)error) {
	c.RegisterHandler(6811, func(ctx Context) error {
		msg := &pb.GodIslandUpdatePowerResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看妖盟战功列表
func (c *Client) OnGodIslandUnionBattleScoreListReq(f func(ctx Context, msg *pb.GodIslandUnionBattleScoreListResp)error) {
	c.RegisterHandler(6812, func(ctx Context) error {
		msg := &pb.GodIslandUnionBattleScoreListResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 灵液领取记录
func (c *Client) OnGodIslandLiquidReceiveRecordReq(f func(ctx Context, msg *pb.GodIslandLiquidReceiveRecordResp)error) {
	c.RegisterHandler(6813, func(ctx Context) error {
		msg := &pb.GodIslandLiquidReceiveRecordResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟成员战功统计
func (c *Client) OnGodIslandUnionMemberScoreListReq(f func(ctx Context, msg *pb.GodIslandUnionMemberScoreListResp)error) {
	c.RegisterHandler(6814, func(ctx Context) error {
		msg := &pb.GodIslandUnionMemberScoreListResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 活动心跳
func (c *Client) OnGodIslandHeartBeatReqMsg(f func(ctx Context, msg *pb.GodIslandHeartBeatRespMsg)error) {
	c.RegisterHandler(6819, func(ctx Context) error {
		msg := &pb.GodIslandHeartBeatRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 游戏基础信息
func (c *Client) OnGodIslandGameInfoReqMsg(f func(ctx Context, msg *pb.GodIslandGameInfoRespMsg)error) {
	c.RegisterHandler(6820, func(ctx Context) error {
		msg := &pb.GodIslandGameInfoRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 游戏路线同步信息
func (c *Client) OnGodIslandGameRouteInfoSyncMsg(f func(ctx Context, msg *pb.GodIslandGameRouteInfoSyncMsg)error) {
	c.RegisterHandler(6822, func(ctx Context) error {
		msg := &pb.GodIslandGameRouteInfoSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 游戏大地图城市信息(在滑动地图时候将视野内的城市id发送)
func (c *Client) OnGodIslandGameCityInfoReqMsg(f func(ctx Context, msg *pb.GodIslandGameCityInfoRespMsg)error) {
	c.RegisterHandler(6823, func(ctx Context) error {
		msg := &pb.GodIslandGameCityInfoRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// npc事件选择
func (c *Client) OnGodIslandGameEventReqMsg(f func(ctx Context, msg *pb.GodIslandGameEventRespMsg)error) {
	c.RegisterHandler(6824, func(ctx Context) error {
		msg := &pb.GodIslandGameEventRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 城市buff同步
func (c *Client) OnGodIslandGameCityBuffSync(f func(ctx Context, msg *pb.GodIslandGameCityBuffSync)error) {
	c.RegisterHandler(6825, func(ctx Context) error {
		msg := &pb.GodIslandGameCityBuffSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 灵力球信息
func (c *Client) OnGodIslandGameSpiritualBallInfoReqMsg(f func(ctx Context, msg *pb.GodIslandGameSpiritualBallInfoRespMsg)error) {
	c.RegisterHandler(6826, func(ctx Context) error {
		msg := &pb.GodIslandGameSpiritualBallInfoRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 使用灵力球
func (c *Client) OnGodIslandGameUseSpiritualBallReqMsg(f func(ctx Context, msg *pb.GodIslandGameUseSpiritualBallRespMsg)error) {
	c.RegisterHandler(6827, func(ctx Context) error {
		msg := &pb.GodIslandGameUseSpiritualBallRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 目标城市具体信息
func (c *Client) OnGodIslandGameTargetCityInfoReqMsg(f func(ctx Context, msg *pb.GodIslandGameTargetCityInfoRespMsg)error) {
	c.RegisterHandler(6828, func(ctx Context) error {
		msg := &pb.GodIslandGameTargetCityInfoRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 目标城市队列信息
func (c *Client) OnGodIslandGameTargetCityLineInfoReqMsg(f func(ctx Context, msg *pb.GodIslandGameTargetCityLineInfoRespMsg)error) {
	c.RegisterHandler(6829, func(ctx Context) error {
		msg := &pb.GodIslandGameTargetCityLineInfoRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进攻
func (c *Client) OnGodIslandGameAttackReqMsg(f func(ctx Context, msg *pb.GodIslandGameAttackRespMsg)error) {
	c.RegisterHandler(6830, func(ctx Context) error {
		msg := &pb.GodIslandGameAttackRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 自动操作 托管
func (c *Client) OnGodIslandGameAutoAttackReqMsg(f func(ctx Context, msg *pb.GodIslandGameAutoAttackRespMsg)error) {
	c.RegisterHandler(6831, func(ctx Context) error {
		msg := &pb.GodIslandGameAutoAttackRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 行军移动
func (c *Client) OnGodIslandGameMoveReqMsg(f func(ctx Context, msg *pb.GodIslandGameMoveRespMsg)error) {
	c.RegisterHandler(6832, func(ctx Context) error {
		msg := &pb.GodIslandGameMoveRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// npc事件同步
func (c *Client) OnGodIslandGameEventSyncMsg(f func(ctx Context, msg *pb.GodIslandGameEventSyncMsg)error) {
	c.RegisterHandler(6833, func(ctx Context) error {
		msg := &pb.GodIslandGameEventSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 小地图信息
func (c *Client) OnGodIslandGameMiniMapInfoReqMsg(f func(ctx Context, msg *pb.GodIslandGameMiniMapInfoRespMsg)error) {
	c.RegisterHandler(6834, func(ctx Context) error {
		msg := &pb.GodIslandGameMiniMapInfoRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 设置联盟类型
func (c *Client) OnGodIslandGameSetUnionTypeReqMsg(f func(ctx Context, msg *pb.GodIslandGameSetUnionTypeRespMsg)error) {
	c.RegisterHandler(6835, func(ctx Context) error {
		msg := &pb.GodIslandGameSetUnionTypeRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 联盟种草信息
func (c *Client) OnGodIslandGamePlantInfoReqMsg(f func(ctx Context, msg *pb.GodIslandGamePlantInfoRespMsg)error) {
	c.RegisterHandler(6836, func(ctx Context) error {
		msg := &pb.GodIslandGamePlantInfoRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 使用仙果 补充灵力球的灵力
func (c *Client) OnGodIslandUseFruitReq(f func(ctx Context, msg *pb.GodIslandUseFruitResp)error) {
	c.RegisterHandler(6837, func(ctx Context) error {
		msg := &pb.GodIslandUseFruitResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 联盟种草领取果实
func (c *Client) OnGodIslandGamePlantReceiveFruitsReqMsg(f func(ctx Context, msg *pb.GodIslandGamePlantReceiveFruitsRespMsg)error) {
	c.RegisterHandler(6838, func(ctx Context) error {
		msg := &pb.GodIslandGamePlantReceiveFruitsRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 联盟种草浇水
func (c *Client) OnGodIslandGamePlantWaterReqMsg(f func(ctx Context, msg *pb.GodIslandGamePlantWaterRespMsg)error) {
	c.RegisterHandler(6839, func(ctx Context) error {
		msg := &pb.GodIslandGamePlantWaterRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 灵晶数据加载
func (c *Client) OnGodIslandCrystalInfoReqMsg(f func(ctx Context, msg *pb.GodIslandCrystalInfoRespMsg)error) {
	c.RegisterHandler(6840, func(ctx Context) error {
		msg := &pb.GodIslandCrystalInfoRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取灵晶
func (c *Client) OnGodIslandCrystalReceiveReqMsg(f func(ctx Context, msg *pb.GodIslandCrystalReceiveRespMsg)error) {
	c.RegisterHandler(6841, func(ctx Context) error {
		msg := &pb.GodIslandCrystalReceiveRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 联盟指挥信息
func (c *Client) OnGodIslandGameStrategyInfoReqMsg(f func(ctx Context, msg *pb.GodIslandGameStrategyInfoRespMsg)error) {
	c.RegisterHandler(6842, func(ctx Context) error {
		msg := &pb.GodIslandGameStrategyInfoRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 联盟指挥设置
func (c *Client) OnGodIslandGameCommanderSetReqMsg(f func(ctx Context, msg *pb.GodIslandGameCommanderSetRespMsg)error) {
	c.RegisterHandler(6843, func(ctx Context) error {
		msg := &pb.GodIslandGameCommanderSetRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 联盟指挥设置目标
func (c *Client) OnGodIslandGameSetUnionTargetReqMsg(f func(ctx Context, msg *pb.GodIslandGameSetUnionTargetRespMsg)error) {
	c.RegisterHandler(6844, func(ctx Context) error {
		msg := &pb.GodIslandGameSetUnionTargetRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 联盟指挥请求支援
func (c *Client) OnGodIslandGameForHelpReqMsg(f func(ctx Context, msg *pb.GodIslandGameForHelpRespMsg)error) {
	c.RegisterHandler(6845, func(ctx Context) error {
		msg := &pb.GodIslandGameForHelpRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 城市变化同步
func (c *Client) OnGodIslandGameCityChangeSyncMsg(f func(ctx Context, msg *pb.GodIslandGameCityChangeSyncMsg)error) {
	c.RegisterHandler(6846, func(ctx Context) error {
		msg := &pb.GodIslandGameCityChangeSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 连胜消息同步
func (c *Client) OnGodIslandGamePlayerWinSyncMsg(f func(ctx Context, msg *pb.GodIslandGamePlayerWinSyncMsg)error) {
	c.RegisterHandler(6847, func(ctx Context) error {
		msg := &pb.GodIslandGamePlayerWinSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 队列消息同步(大地图，城市队列，具体队列信息都通过这条去同步刷新)
func (c *Client) OnGodIslandGameLineInfoSyncMsg(f func(ctx Context, msg *pb.GodIslandGameLineInfoSyncMsg)error) {
	c.RegisterHandler(6848, func(ctx Context) error {
		msg := &pb.GodIslandGameLineInfoSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 加速行军
func (c *Client) OnGodIslandGameAcclerateMoveMsg(f func(ctx Context, msg *pb.GodIslandGameAcclerateMoveRespMsg)error) {
	c.RegisterHandler(6849, func(ctx Context) error {
		msg := &pb.GodIslandGameAcclerateMoveRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 我的游戏信息同步
func (c *Client) OnGodIslandGameMyGameInfoSyncMsg(f func(ctx Context, msg *pb.GodIslandGameMyGameInfoSyncMsg)error) {
	c.RegisterHandler(6850, func(ctx Context) error {
		msg := &pb.GodIslandGameMyGameInfoSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 联盟发送灵液
func (c *Client) OnGodIslandGameSendLiquidReqMsg(f func(ctx Context, msg *pb.GodIslandGameSendLiquidRespMsg)error) {
	c.RegisterHandler(6851, func(ctx Context) error {
		msg := &pb.GodIslandGameSendLiquidRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 联盟灵液发送记录
func (c *Client) OnGodIslandGameSendLiquidRecordReqMsg(f func(ctx Context, msg *pb.GodIslandGameSendLiquidRecordRespMsg)error) {
	c.RegisterHandler(6852, func(ctx Context) error {
		msg := &pb.GodIslandGameSendLiquidRecordRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 跑马灯消息同步
func (c *Client) OnGodIslandGameHorseLampSyncMsg(f func(ctx Context, msg *pb.GodIslandGameHorseLampSyncMsg)error) {
	c.RegisterHandler(6853, func(ctx Context) error {
		msg := &pb.GodIslandGameHorseLampSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 玩家的所向披靡同步
func (c *Client) OnGodIslandSuppressBuffSyncMsg(f func(ctx Context, msg *pb.GodIslandSuppressBuffSyncMsg)error) {
	c.RegisterHandler(6854, func(ctx Context) error {
		msg := &pb.GodIslandSuppressBuffSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求城市界面buff信息
func (c *Client) OnGodIslandGameCityBuffInfoReqMsg(f func(ctx Context, msg *pb.GodIslandGameCityBuffInfoRespMsg)error) {
	c.RegisterHandler(6855, func(ctx Context) error {
		msg := &pb.GodIslandGameCityBuffInfoRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 阵亡之后 补兵完成之后从场外回大本营
func (c *Client) OnGodIslandBackHomeReq(f func(ctx Context, msg *pb.GodIslandBackHomeResp)error) {
	c.RegisterHandler(6856, func(ctx Context) error {
		msg := &pb.GodIslandBackHomeResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 场内有妖盟被淘汰 同步下来
func (c *Client) OnGodIslandLeftOutSyncInfo(f func(ctx Context, msg *pb.GodIslandLeftOutSyncInfo)error) {
	c.RegisterHandler(6857, func(ctx Context) error {
		msg := &pb.GodIslandLeftOutSyncInfo{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 修改盟友
func (c *Client) OnGodIslandSetFriendUnionReq(f func(ctx Context, msg *pb.GodIslandSetFriendUnionResp)error) {
	c.RegisterHandler(6858, func(ctx Context) error {
		msg := &pb.GodIslandSetFriendUnionResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 盟友同步
func (c *Client) OnGodIslandSyncFriendUnionMsg(f func(ctx Context, msg *pb.GodIslandSyncFriendUnionMsg)error) {
	c.RegisterHandler(6859, func(ctx Context) error {
		msg := &pb.GodIslandSyncFriendUnionMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 指挥权变化同步
func (c *Client) OnGodIslandCommanderChangeSyncMsg(f func(ctx Context, msg *pb.GodIslandCommanderChangeSyncMsg)error) {
	c.RegisterHandler(6860, func(ctx Context) error {
		msg := &pb.GodIslandCommanderChangeSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 被击杀后同步
func (c *Client) OnGodIslandBeenKillSyncMsg(f func(ctx Context, msg *pb.GodIslandBeenKillSyncMsg)error) {
	c.RegisterHandler(6861, func(ctx Context) error {
		msg := &pb.GodIslandBeenKillSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 城市战斗状态
func (c *Client) OnGodIslandGameCityBattleSyncMsg(f func(ctx Context, msg *pb.GodIslandGameCityBattleSyncMsg)error) {
	c.RegisterHandler(6862, func(ctx Context) error {
		msg := &pb.GodIslandGameCityBattleSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 托管操作 结束后同步给客户端
func (c *Client) OnGodIslandAutoBattleStopSyncMsg(f func(ctx Context, msg *pb.GodIslandAutoBattleStopSyncMsg)error) {
	c.RegisterHandler(6863, func(ctx Context) error {
		msg := &pb.GodIslandAutoBattleStopSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 修罗城解锁同步
func (c *Client) OnGodIslandGhostCityUnlockSyncMsg(f func(ctx Context, msg *pb.GodIslandGhostCityUnlockSyncMsg)error) {
	c.RegisterHandler(6864, func(ctx Context) error {
		msg := &pb.GodIslandGhostCityUnlockSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 星宿试炼数据同步
func (c *Client) OnStarTrialDataMsg(f func(ctx Context, msg *pb.StarTrialDataMsg)error) {
	c.RegisterHandler(6901, func(ctx Context) error {
		msg := &pb.StarTrialDataMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 星宿试炼战斗
func (c *Client) OnStarTrialChallengeReq(f func(ctx Context, msg *pb.StarTrialChallengeResp)error) {
	c.RegisterHandler(6902, func(ctx Context) error {
		msg := &pb.StarTrialChallengeResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取日志信息
func (c *Client) OnStarTrialRecordReq(f func(ctx Context, msg *pb.StarTrialRecordResp)error) {
	c.RegisterHandler(6903, func(ctx Context) error {
		msg := &pb.StarTrialRecordResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 战斗回放
func (c *Client) OnStarTrialBattleReplyReq(f func(ctx Context, msg *pb.StarTrialBattleReplyResp)error) {
	c.RegisterHandler(6904, func(ctx Context) error {
		msg := &pb.StarTrialBattleReplyResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入星空图鉴信息
func (c *Client) OnEnterStarTrialCodexMsgReq(f func(ctx Context, msg *pb.EnterStarTrialCodexMsgResp)error) {
	c.RegisterHandler(6905, func(ctx Context) error {
		msg := &pb.EnterStarTrialCodexMsgResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取详细图鉴击败信息
func (c *Client) OnPlayerStarTrialCodexMsgReq(f func(ctx Context, msg *pb.RspPlayerStarTrialCodexMsg)error) {
	c.RegisterHandler(6906, func(ctx Context) error {
		msg := &pb.RspPlayerStarTrialCodexMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取Boss属性信息
func (c *Client) OnGetBossDetailDataReq(f func(ctx Context, msg *pb.GetBossDetailDataResp)error) {
	c.RegisterHandler(6907, func(ctx Context) error {
		msg := &pb.GetBossDetailDataResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 每日领奖
func (c *Client) OnGetDailyFightRewardReq(f func(ctx Context, msg *pb.GetDailyFightRewardResp)error) {
	c.RegisterHandler(6908, func(ctx Context) error {
		msg := &pb.GetDailyFightRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取分组信息
func (c *Client) OnGetStarTrialGroupInfoReq(f func(ctx Context, msg *pb.GetStarTrialGroupInfoResp)error) {
	c.RegisterHandler(6909, func(ctx Context) error {
		msg := &pb.GetStarTrialGroupInfoResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看玩家详情
func (c *Client) OnStarTrialGetPlayerDetailReq(f func(ctx Context, msg *pb.StarTrialGetPlayerDetailResp)error) {
	c.RegisterHandler(6910, func(ctx Context) error {
		msg := &pb.StarTrialGetPlayerDetailResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入星宿试炼
func (c *Client) OnStarTrialEnterMainPanelReq(f func(ctx Context, msg *pb.StarTrialEnterMainPanelResp)error) {
	c.RegisterHandler(6911, func(ctx Context) error {
		msg := &pb.StarTrialEnterMainPanelResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入聚灵阵
func (c *Client) OnGatherEnergyEnterNewReq(f func(ctx Context, msg *pb.GatherEnergyEnterNewResp)error) {
	c.RegisterHandler(7001, func(ctx Context) error {
		msg := &pb.GatherEnergyEnterNewResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步聚灵阵数据
func (c *Client) OnSyncGatherEnergyMsg(f func(ctx Context, msg *pb.SyncGatherEnergyMsg)error) {
	c.RegisterHandler(7002, func(ctx Context) error {
		msg := &pb.SyncGatherEnergyMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 开启灵阵界面
func (c *Client) OnGatherEnergyOpenViewReq(f func(ctx Context, msg *pb.GatherEnergyOpenViewResp)error) {
	c.RegisterHandler(7003, func(ctx Context) error {
		msg := &pb.GatherEnergyOpenViewResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 开启某个灵阵
func (c *Client) OnGatherEnergyOpenReq(f func(ctx Context, msg *pb.GatherEnergyOpenResp)error) {
	c.RegisterHandler(7004, func(ctx Context) error {
		msg := &pb.GatherEnergyOpenResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 挑战某人
func (c *Client) OnGatherEnergyFightReq(f func(ctx Context, msg *pb.GatherEnergyFightResp)error) {
	c.RegisterHandler(7007, func(ctx Context) error {
		msg := &pb.GatherEnergyFightResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 聚灵阵奖励推送
func (c *Client) OnGatherEnergyRewardShowResp(f func(ctx Context, msg *pb.GatherEnergyRewardShowResp)error) {
	c.RegisterHandler(7010, func(ctx Context) error {
		msg := &pb.GatherEnergyRewardShowResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 聚灵阵奖励领取
func (c *Client) OnGatherEnergyRewardReq(f func(ctx Context, msg *pb.GatherEnergyRewardResp)error) {
	c.RegisterHandler(7011, func(ctx Context) error {
		msg := &pb.GatherEnergyRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 聚灵阵排行榜点赞
func (c *Client) OnGatherEnergyLikeReq(f func(ctx Context, msg *pb.GatherEnergyLikeResp)error) {
	c.RegisterHandler(7012, func(ctx Context) error {
		msg := &pb.GatherEnergyLikeResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 灵气转换
func (c *Client) OnGatherEnergyTransformReq(f func(ctx Context, msg *pb.GatherEnergyTransformResp)error) {
	c.RegisterHandler(7015, func(ctx Context) error {
		msg := &pb.GatherEnergyTransformResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 离开聚灵阵
func (c *Client) OnGatherEnergyLeaveReq(f func(ctx Context, msg *pb.GatherEnergyLeaveResp)error) {
	c.RegisterHandler(7017, func(ctx Context) error {
		msg := &pb.GatherEnergyLeaveResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 聚灵阵公告
func (c *Client) OnGatherEnergyNoticeReq(f func(ctx Context, msg *pb.GatherEnergyNoticeResp)error) {
	c.RegisterHandler(7018, func(ctx Context) error {
		msg := &pb.GatherEnergyNoticeResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 聚灵阵获取广告奖励
func (c *Client) OnGatherEnergyGetADAwardReq(f func(ctx Context, msg *pb.GatherEnergyGetADAwardResp)error) {
	c.RegisterHandler(7019, func(ctx Context) error {
		msg := &pb.GatherEnergyGetADAwardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 打开聚灵阵一级列表界面（参与人数）
func (c *Client) OnGatherEnergyFirstListViewReq(f func(ctx Context, msg *pb.GatherEnergyFirstListViewResp)error) {
	c.RegisterHandler(7020, func(ctx Context) error {
		msg := &pb.GatherEnergyFirstListViewResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 打开聚灵阵二级列表界面（参与名字）
func (c *Client) OnGatherEnergySecondListViewReq(f func(ctx Context, msg *pb.GatherEnergySecondListViewResp)error) {
	c.RegisterHandler(7021, func(ctx Context) error {
		msg := &pb.GatherEnergySecondListViewResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 打开某个聚灵阵内部
func (c *Client) OnGatherEnergyInsideViewNewReq(f func(ctx Context, msg *pb.GatherEnergyInsideViewNewResp)error) {
	c.RegisterHandler(7022, func(ctx Context) error {
		msg := &pb.GatherEnergyInsideViewNewResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 打开战报界面
func (c *Client) OnGatherEnergyFightReportNewReq(f func(ctx Context, msg *pb.GatherEnergyFightReportNewResp)error) {
	c.RegisterHandler(7023, func(ctx Context) error {
		msg := &pb.GatherEnergyFightReportNewResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查询某人参与的阵法
func (c *Client) OnGatherEnergySearchNewReq(f func(ctx Context, msg *pb.GatherEnergySearchNewResp)error) {
	c.RegisterHandler(7024, func(ctx Context) error {
		msg := &pb.GatherEnergySearchNewResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 加入聚灵阵
func (c *Client) OnGatherEnergyAttendNewReq(f func(ctx Context, msg *pb.GatherEnergyAttendNewResp)error) {
	c.RegisterHandler(7025, func(ctx Context) error {
		msg := &pb.GatherEnergyAttendNewResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步sdk奖励数据
func (c *Client) OnReceiveSdkRewardSyn(f func(ctx Context, msg *pb.ReceiveSdkRewardSyn)error) {
	c.RegisterHandler(7201, func(ctx Context) error {
		msg := &pb.ReceiveSdkRewardSyn{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取sdk奖励
func (c *Client) OnReceiveSdkRewardReq(f func(ctx Context, msg *pb.ReceiveSdkRewardRsp)error) {
	c.RegisterHandler(7202, func(ctx Context) error {
		msg := &pb.ReceiveSdkRewardRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取qq卡券奖励
func (c *Client) OnQQCardGetRewardReq(f func(ctx Context, msg *pb.QQCardGetRewardRsp)error) {
	c.RegisterHandler(7203, func(ctx Context) error {
		msg := &pb.QQCardGetRewardRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 支付宝启动参数
func (c *Client) OnAlipayStartParamReq(f func(ctx Context, msg *pb.AlipayStartParamRsp)error) {
	c.RegisterHandler(7204, func(ctx Context) error {
		msg := &pb.AlipayStartParamRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// Sdk七日登录奖励领取
func (c *Client) OnReceiveSdkDailyRewardReq(f func(ctx Context, msg *pb.ReceiveSdkDailyRewardRsp)error) {
	c.RegisterHandler(7205, func(ctx Context) error {
		msg := &pb.ReceiveSdkDailyRewardRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 美团启动参数
func (c *Client) OnMeiTuanStartParamReq(f func(ctx Context, msg *pb.MeiTuanStartParamRsp)error) {
	c.RegisterHandler(7206, func(ctx Context) error {
		msg := &pb.MeiTuanStartParamRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入征战诸天
func (c *Client) OnSkyWarEnterReq(f func(ctx Context, msg *pb.SkyWarEnterRsp)error) {
	c.RegisterHandler(8401, func(ctx Context) error {
		msg := &pb.SkyWarEnterRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 刷新对手玩家
func (c *Client) OnSkyWarRefreshEnemyReq(f func(ctx Context, msg *pb.SkyWarRefreshEnemyRsp)error) {
	c.RegisterHandler(8402, func(ctx Context) error {
		msg := &pb.SkyWarRefreshEnemyRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 斗法
func (c *Client) OnSkyWarFightReq(f func(ctx Context, msg *pb.SkyWarFightRsp)error) {
	c.RegisterHandler(8403, func(ctx Context) error {
		msg := &pb.SkyWarFightRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取排行榜数据
func (c *Client) OnSkyWarRankReq(f func(ctx Context, msg *pb.SkyWarRankRsp)error) {
	c.RegisterHandler(8404, func(ctx Context) error {
		msg := &pb.SkyWarRankRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取日志
func (c *Client) OnSkyWarLogReq(f func(ctx Context, msg *pb.SkyWarLogRsp)error) {
	c.RegisterHandler(8405, func(ctx Context) error {
		msg := &pb.SkyWarLogRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 战斗回放
func (c *Client) OnSkyWarLogPlaybackReq(f func(ctx Context, msg *pb.SkyWarLogPlaybackRsp)error) {
	c.RegisterHandler(8406, func(ctx Context) error {
		msg := &pb.SkyWarLogPlaybackRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取布阵详情
func (c *Client) OnSkyWarFormationReq(f func(ctx Context, msg *pb.SkyWarFormationRsp)error) {
	c.RegisterHandler(8407, func(ctx Context) error {
		msg := &pb.SkyWarFormationRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 设置顺位
func (c *Client) OnSkyWarSetOrderReq(f func(ctx Context, msg *pb.SkyWarSetOrderRsp)error) {
	c.RegisterHandler(8408, func(ctx Context) error {
		msg := &pb.SkyWarSetOrderRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 诸天榜
func (c *Client) OnSkyWarSkyRankReq(f func(ctx Context, msg *pb.SkyWarSkyRankRsp)error) {
	c.RegisterHandler(8409, func(ctx Context) error {
		msg := &pb.SkyWarSkyRankRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 膜拜
func (c *Client) OnSkyWarWorshipReq(f func(ctx Context, msg *pb.SkyWarWorshipRsp)error) {
	c.RegisterHandler(8410, func(ctx Context) error {
		msg := &pb.SkyWarWorshipRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 购买次数(斗法、刷新)
func (c *Client) OnSkyWarBuyTimesReq(f func(ctx Context, msg *pb.SkyWarBuyTimesRsp)error) {
	c.RegisterHandler(8411, func(ctx Context) error {
		msg := &pb.SkyWarBuyTimesRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 数据同步
func (c *Client) OnSkyWarDataSync(f func(ctx Context, msg *pb.SkyWarDataSync)error) {
	c.RegisterHandler(8413, func(ctx Context) error {
		msg := &pb.SkyWarDataSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 征战诸天数据登录同步，阶段变化同步1
func (c *Client) OnSkyWarDataLoginSync(f func(ctx Context, msg *pb.SkyWarDataLoginSync)error) {
	c.RegisterHandler(8414, func(ctx Context) error {
		msg := &pb.SkyWarDataLoginSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟对决-同步妖盟对决数据
func (c *Client) OnUnionDuelSyncMsgReq(f func(ctx Context, msg *pb.SyncUnionDuelMsg)error) {
	c.RegisterHandler(8501, func(ctx Context) error {
		msg := &pb.SyncUnionDuelMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟对决-开启主界面
func (c *Client) OnUnionDuelOpenViewReq(f func(ctx Context, msg *pb.UnionDuelOpenViewResp)error) {
	c.RegisterHandler(8502, func(ctx Context) error {
		msg := &pb.UnionDuelOpenViewResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟对决-妖盟对决同步数据
func (c *Client) OnUnionDuelSyncDataReq(f func(ctx Context, msg *pb.UnionDuelSyncDataResp)error) {
	c.RegisterHandler(8511, func(ctx Context) error {
		msg := &pb.UnionDuelSyncDataResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步评分状态
func (c *Client) OnSyncMarkStateMsg(f func(ctx Context, msg *pb.SyncMarkStateMsg)error) {
	c.RegisterHandler(8601, func(ctx Context) error {
		msg := &pb.SyncMarkStateMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 评分完成
func (c *Client) OnMarkFinishReq(f func(ctx Context, msg *pb.MarkFinishResp)error) {
	c.RegisterHandler(8602, func(ctx Context) error {
		msg := &pb.MarkFinishResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 法则感悟
func (c *Client) OnWorldRulePerceptionReq(f func(ctx Context, msg *pb.WorldRulePerceptionResp)error) {
	c.RegisterHandler(9001, func(ctx Context) error {
		msg := &pb.WorldRulePerceptionResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 法则感悟替换
func (c *Client) OnWorldRulePerceptionReplaceReq(f func(ctx Context, msg *pb.WorldRulePerceptionReplaceResp)error) {
	c.RegisterHandler(9002, func(ctx Context) error {
		msg := &pb.WorldRulePerceptionReplaceResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取法则感悟记录
func (c *Client) OnWorldRuleGetPerceptionLogReq(f func(ctx Context, msg *pb.WorldRuleGetPerceptionLogResp)error) {
	c.RegisterHandler(9003, func(ctx Context) error {
		msg := &pb.WorldRuleGetPerceptionLogResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 切换法则方案
func (c *Client) OnWorldRuleSwitchProgrammeReq(f func(ctx Context, msg *pb.WorldRuleSwitchProgrammeResp)error) {
	c.RegisterHandler(9004, func(ctx Context) error {
		msg := &pb.WorldRuleSwitchProgrammeResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 天地法则玩家数据同步
func (c *Client) OnWorldRulePlayerDataMsg(f func(ctx Context, msg *pb.WorldRulePlayerDataMsg)error) {
	c.RegisterHandler(9005, func(ctx Context) error {
		msg := &pb.WorldRulePlayerDataMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求法则试炼挑战
func (c *Client) OnRuleTrialChallengeReq(f func(ctx Context, msg *pb.RuleTrialChallengeResp)error) {
	c.RegisterHandler(9101, func(ctx Context) error {
		msg := &pb.RuleTrialChallengeResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求法则试炼一键扫荡
func (c *Client) OnRuleTrialRepeatAllReq(f func(ctx Context, msg *pb.RuleTrialRepeatAllResp)error) {
	c.RegisterHandler(9103, func(ctx Context) error {
		msg := &pb.RuleTrialRepeatAllResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步法则试炼Boss妖力
func (c *Client) OnRuleTrialBossConfigSyncMsg(f func(ctx Context, msg *pb.RuleTrialBossConfigSyncMsg)error) {
	c.RegisterHandler(9104, func(ctx Context) error {
		msg := &pb.RuleTrialBossConfigSyncMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步法则试炼
func (c *Client) OnRuleTrialDataSync(f func(ctx Context, msg *pb.RuleTrialDataSync)error) {
	c.RegisterHandler(9105, func(ctx Context) error {
		msg := &pb.RuleTrialDataSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求法则试炼Boss数据
func (c *Client) OnRuleTrialMonsterAttrReq(f func(ctx Context, msg *pb.RuleTrialMonsterAttrResp)error) {
	c.RegisterHandler(9106, func(ctx Context) error {
		msg := &pb.RuleTrialMonsterAttrResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 九幽争霸数据登录同步
func (c *Client) OnHolyLandBattleApplyDataSyncReq(f func(ctx Context, msg *pb.HolyLandBattleApplyDataSync)error) {
	c.RegisterHandler(9570, func(ctx Context) error {
		msg := &pb.HolyLandBattleApplyDataSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 九幽争霸赛事时间戳数据登录前同步
func (c *Client) OnHolyLandBattleTimeStampsDataSync(f func(ctx Context, msg *pb.HolyLandBattleTimeStampsDataSync)error) {
	c.RegisterHandler(9575, func(ctx Context) error {
		msg := &pb.HolyLandBattleTimeStampsDataSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 补签接口
func (c *Client) OnReqSignInFundRepair(f func(ctx Context, msg *pb.RspSignInFundRepair)error) {
	c.RegisterHandler(9601, func(ctx Context) error {
		msg := &pb.RspSignInFundRepair{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入山海界面
func (c *Client) OnEnterMountainSeaReq(f func(ctx Context, msg *pb.EnterMountainSeaRsp)error) {
	c.RegisterHandler(9701, func(ctx Context) error {
		msg := &pb.EnterMountainSeaRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 组队内部请求
func (c *Client) OnEnterMountainSeaTeamReq(f func(ctx Context, msg *pb.EnterMountainSeaTeamRsp)error) {
	c.RegisterHandler(9702, func(ctx Context) error {
		msg := &pb.EnterMountainSeaTeamRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 组队页面请求
func (c *Client) OnMountainSeaTeamStartReq(f func(ctx Context, msg *pb.MountainSeaTeamStartRsp)error) {
	c.RegisterHandler(9703, func(ctx Context) error {
		msg := &pb.MountainSeaTeamStartRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 创建队伍
func (c *Client) OnMountainSeaCreateTeamReq(f func(ctx Context, msg *pb.MountainSeaCreateTeamRsp)error) {
	c.RegisterHandler(9704, func(ctx Context) error {
		msg := &pb.MountainSeaCreateTeamRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取队伍列表
func (c *Client) OnMountainSeaGetTeamListReq(f func(ctx Context, msg *pb.MountainSeaGetTeamListRsp)error) {
	c.RegisterHandler(9705, func(ctx Context) error {
		msg := &pb.MountainSeaGetTeamListRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查找队伍
func (c *Client) OnMountainSeaGetTeamInfoReq(f func(ctx Context, msg *pb.MountainSeaGetTeamInfoRsp)error) {
	c.RegisterHandler(9706, func(ctx Context) error {
		msg := &pb.MountainSeaGetTeamInfoRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 加入队伍
func (c *Client) OnMountainSeaJoinTeamReq(f func(ctx Context, msg *pb.MountainSeaJoinTeamRsp)error) {
	c.RegisterHandler(9707, func(ctx Context) error {
		msg := &pb.MountainSeaJoinTeamRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 取消加入队伍
func (c *Client) OnMountainSeaCancelTeamApplyReq(f func(ctx Context, msg *pb.MountainSeaCancelTeamApplyRsp)error) {
	c.RegisterHandler(9708, func(ctx Context) error {
		msg := &pb.MountainSeaCancelTeamApplyRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同意申请
func (c *Client) OnMountainSeaApplyJoinTeamAgreeReq(f func(ctx Context, msg *pb.MountainSeaApplyJoinTeamAgreeRsp)error) {
	c.RegisterHandler(9709, func(ctx Context) error {
		msg := &pb.MountainSeaApplyJoinTeamAgreeRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 一键拒绝
func (c *Client) OnMountainSeaApplyJoinTeamRefusedReq(f func(ctx Context, msg *pb.MountainSeaApplyJoinTeamRefusedRsp)error) {
	c.RegisterHandler(9710, func(ctx Context) error {
		msg := &pb.MountainSeaApplyJoinTeamRefusedRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 退出队伍
func (c *Client) OnMountainSeaQuitTeamReq(f func(ctx Context, msg *pb.MountainSeaQuitTeamRsp)error) {
	c.RegisterHandler(9711, func(ctx Context) error {
		msg := &pb.MountainSeaQuitTeamRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 踢出队伍
func (c *Client) OnMountainSeaKickOutTeamReq(f func(ctx Context, msg *pb.MountainSeaKickOutTeamRsp)error) {
	c.RegisterHandler(9712, func(ctx Context) error {
		msg := &pb.MountainSeaKickOutTeamRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 转让队长
func (c *Client) OnMountainSeaChangeLeaderReq(f func(ctx Context, msg *pb.MountainSeaChangeLeaderRsp)error) {
	c.RegisterHandler(9713, func(ctx Context) error {
		msg := &pb.MountainSeaChangeLeaderRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 开始准备
func (c *Client) OnMountainSeaTeamPrepareReq(f func(ctx Context, msg *pb.MountainSeaTeamPrepareRsp)error) {
	c.RegisterHandler(9714, func(ctx Context) error {
		msg := &pb.MountainSeaTeamPrepareRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 匹配
func (c *Client) OnMountainSeaMatchMemberReq(f func(ctx Context, msg *pb.MountainSeaMatchMemberRsp)error) {
	c.RegisterHandler(9715, func(ctx Context) error {
		msg := &pb.MountainSeaMatchMemberRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 开始战斗
func (c *Client) OnMountainSeaStartBattleReq(f func(ctx Context, msg *pb.MountainSeaStartBattleRsp)error) {
	c.RegisterHandler(9716, func(ctx Context) error {
		msg := &pb.MountainSeaStartBattleRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求玩家数据
func (c *Client) OnMountainSeaGetPlayerInfoReq(f func(ctx Context, msg *pb.MountainSeaGetPlayerInfoRsp)error) {
	c.RegisterHandler(9717, func(ctx Context) error {
		msg := &pb.MountainSeaGetPlayerInfoRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 膜拜
func (c *Client) OnMountainSeaWorshipReq(f func(ctx Context, msg *pb.MountainSeaWorshipRsp)error) {
	c.RegisterHandler(9718, func(ctx Context) error {
		msg := &pb.MountainSeaWorshipRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 战斗回放
func (c *Client) OnMountainSeaGetBattleVideoReq(f func(ctx Context, msg *pb.MountainSeaGetBattleVideoRsp)error) {
	c.RegisterHandler(9719, func(ctx Context) error {
		msg := &pb.MountainSeaGetBattleVideoRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取boss信息
func (c *Client) OnMountainSeaGetBossInfoReq(f func(ctx Context, msg *pb.MountainSeaGetBossInfoRsp)error) {
	c.RegisterHandler(9720, func(ctx Context) error {
		msg := &pb.MountainSeaGetBossInfoRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求挑战次数
func (c *Client) OnMountainSeaChallengeTimeReq(f func(ctx Context, msg *pb.MountainSeaChallengeTimeRsp)error) {
	c.RegisterHandler(9721, func(ctx Context) error {
		msg := &pb.MountainSeaChallengeTimeRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 开始匹配
func (c *Client) OnMountainSeaStartMatchReq(f func(ctx Context, msg *pb.MountainSeaStartMatchRsp)error) {
	c.RegisterHandler(9722, func(ctx Context) error {
		msg := &pb.MountainSeaStartMatchRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 邀请
func (c *Client) OnMountainSeaInviteReq(f func(ctx Context, msg *pb.MountainSeaInviteRsp)error) {
	c.RegisterHandler(9723, func(ctx Context) error {
		msg := &pb.MountainSeaInviteRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 离开系统
func (c *Client) OnLeaveMountainSeaReq(f func(ctx Context, msg *pb.LeaveMountainSeaRsp)error) {
	c.RegisterHandler(9724, func(ctx Context) error {
		msg := &pb.LeaveMountainSeaRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取boss妖力
func (c *Client) OnMountainSeaGetBossPowerReq(f func(ctx Context, msg *pb.MountainSeaGetBossPowerRsp)error) {
	c.RegisterHandler(9725, func(ctx Context) error {
		msg := &pb.MountainSeaGetBossPowerRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 队长接受到广播
func (c *Client) OnMountainSeaTeamLeaderNotify(f func(ctx Context, msg *pb.MountainSeaTeamLeaderNotify)error) {
	c.RegisterHandler(9731, func(ctx Context) error {
		msg := &pb.MountainSeaTeamLeaderNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 队员接收到的广播
func (c *Client) OnMountainSeaTeamMemberNotify(f func(ctx Context, msg *pb.MountainSeaTeamMemberNotify)error) {
	c.RegisterHandler(9732, func(ctx Context) error {
		msg := &pb.MountainSeaTeamMemberNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入布阵界面
func (c *Client) OnMountainSeaEnterBattleReq(f func(ctx Context, msg *pb.MountainSeaEnterBattleResp)error) {
	c.RegisterHandler(9750, func(ctx Context) error {
		msg := &pb.MountainSeaEnterBattleResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入切换分身界面
func (c *Client) OnMountainSeaEnterSwitchSeparationReq(f func(ctx Context, msg *pb.MountainSeaEnterSwitchSeparationRsp)error) {
	c.RegisterHandler(9752, func(ctx Context) error {
		msg := &pb.MountainSeaEnterSwitchSeparationRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 切换上阵的分身
func (c *Client) OnMountainSeaSwitchSeparationReq(f func(ctx Context, msg *pb.MountainSeaSwitchSeparationRsp)error) {
	c.RegisterHandler(9753, func(ctx Context) error {
		msg := &pb.MountainSeaSwitchSeparationRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看分身详细数据
func (c *Client) OnMountainSeaSeparationDetailReq(f func(ctx Context, msg *pb.MountainSeaSeparationDetailRsp)error) {
	c.RegisterHandler(9754, func(ctx Context) error {
		msg := &pb.MountainSeaSeparationDetailRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 改变参战对象位置
func (c *Client) OnMountainSeaChangePosReq(f func(ctx Context, msg *pb.MountainSeaChangePosRsp)error) {
	c.RegisterHandler(9755, func(ctx Context) error {
		msg := &pb.MountainSeaChangePosRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 改变装配的队伍技能
func (c *Client) OnMountainSeaChangeTeamSkillReq(f func(ctx Context, msg *pb.MountainSeaChangeTeamSkillRsp)error) {
	c.RegisterHandler(9756, func(ctx Context) error {
		msg := &pb.MountainSeaChangeTeamSkillRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进行战斗
func (c *Client) OnMountainSeaDoBattleReq(f func(ctx Context, msg *pb.MountainSeaDoBattleRsp)error) {
	c.RegisterHandler(9757, func(ctx Context) error {
		msg := &pb.MountainSeaDoBattleRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 广播进入布阵界面
func (c *Client) OnMountainSeaEnterBattleNotify(f func(ctx Context, msg *pb.MountainSeaEnterBattleNotify)error) {
	c.RegisterHandler(9770, func(ctx Context) error {
		msg := &pb.MountainSeaEnterBattleNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 切换上阵分身广播
func (c *Client) OnMountainSeaSwitchSeparationNotify(f func(ctx Context, msg *pb.MountainSeaSwitchSeparationNotify)error) {
	c.RegisterHandler(9771, func(ctx Context) error {
		msg := &pb.MountainSeaSwitchSeparationNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 装配的队伍技能变更广播
func (c *Client) OnMountainSeaChangeTeamSkillNotify(f func(ctx Context, msg *pb.MountainSeaChangeTeamSkillNotify)error) {
	c.RegisterHandler(9772, func(ctx Context) error {
		msg := &pb.MountainSeaChangeTeamSkillNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 广播战斗数据
func (c *Client) OnMountainSeaDoBattleNotify(f func(ctx Context, msg *pb.MountainSeaDoBattleNotify)error) {
	c.RegisterHandler(9773, func(ctx Context) error {
		msg := &pb.MountainSeaDoBattleNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 参战对象位置变更广播
func (c *Client) OnMountainSeaChangePosNotify(f func(ctx Context, msg *pb.MountainSeaChangePosNotify)error) {
	c.RegisterHandler(9774, func(ctx Context) error {
		msg := &pb.MountainSeaChangePosNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 邀请列表
func (c *Client) OnMountainSeaInviteListReq(f func(ctx Context, msg *pb.MountainSeaInviteListResp)error) {
	c.RegisterHandler(9783, func(ctx Context) error {
		msg := &pb.MountainSeaInviteListResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求红点
func (c *Client) OnMountainSeaRedPointResp(f func(ctx Context, msg *pb.MountainSeaRedPointResp)error) {
	c.RegisterHandler(9784, func(ctx Context) error {
		msg := &pb.MountainSeaRedPointResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 一键拒绝邀请队伍
func (c *Client) OnMountainSeaInviteRefuseReq(f func(ctx Context, msg *pb.MountainSeaInviteRefuseResp)error) {
	c.RegisterHandler(9785, func(ctx Context) error {
		msg := &pb.MountainSeaInviteRefuseResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 设置接受指定邀请开关
func (c *Client) OnMountainSeaSetAppointReq(f func(ctx Context, msg *pb.MountainSeaSetAppointResp)error) {
	c.RegisterHandler(9786, func(ctx Context) error {
		msg := &pb.MountainSeaSetAppointResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 加入指定邀请队伍
func (c *Client) OnMountainSeaApplyJoinTeamReq(f func(ctx Context, msg *pb.MountainSeaApplyJoinTeamRsp)error) {
	c.RegisterHandler(9787, func(ctx Context) error {
		msg := &pb.MountainSeaApplyJoinTeamRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 琼浆玉液数据同步
func (c *Client) OnPlayerSoulLiQuidDataMsg(f func(ctx Context, msg *pb.PlayerSoulLiQuidDataMsg)error) {
	c.RegisterHandler(9801, func(ctx Context) error {
		msg := &pb.PlayerSoulLiQuidDataMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 疯狂聚宝盆签到
func (c *Client) OnReqGetConditionReward(f func(ctx Context, msg *pb.RspGetConditionReward)error) {
	c.RegisterHandler(9901, func(ctx Context) error {
		msg := &pb.RspGetConditionReward{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 天道馈赠数据同步
func (c *Client) OnSkyPresentDataSync(f func(ctx Context, msg *pb.SkyPresentDataSync)error) {
	c.RegisterHandler(9910, func(ctx Context) error {
		msg := &pb.SkyPresentDataSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 仙友回归-首次分享领奖
func (c *Client) OnRegressionShareReq(f func(ctx Context, msg *pb.RegressionShareResp)error) {
	c.RegisterHandler(10001, func(ctx Context) error {
		msg := &pb.RegressionShareResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 仙友回归-抽奖
func (c *Client) OnRegressionLotteryReq(f func(ctx Context, msg *pb.RegressionLotteryResp)error) {
	c.RegisterHandler(10002, func(ctx Context) error {
		msg := &pb.RegressionLotteryResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 仙友回归-绑定回归玩家数据
func (c *Client) OnGetRegressionPlayerDataReq(f func(ctx Context, msg *pb.GetRegressionPlayerDataResp)error) {
	c.RegisterHandler(10003, func(ctx Context) error {
		msg := &pb.GetRegressionPlayerDataResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 仙友回归-领取绑定玩家的任务奖励
func (c *Client) OnGetRegressionReceiveRewardReq(f func(ctx Context, msg *pb.GetRegressionReceiveRewardResp)error) {
	c.RegisterHandler(10004, func(ctx Context) error {
		msg := &pb.GetRegressionReceiveRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 仙友回归-保存自选奖励
func (c *Client) OnRegressionSaveSelectItemReq(f func(ctx Context, msg *pb.RegressionSaveSelectItemResp)error) {
	c.RegisterHandler(10005, func(ctx Context) error {
		msg := &pb.RegressionSaveSelectItemResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟赐福
func (c *Client) OnUnionBlessingSendGiftReq(f func(ctx Context, msg *pb.UnionBlessingSendGiftResp)error) {
	c.RegisterHandler(10601, func(ctx Context) error {
		msg := &pb.UnionBlessingSendGiftResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取赐福奖励
func (c *Client) OnUnionBlessingRewardReqMsg(f func(ctx Context, msg *pb.UnionBlessingRewardRespMsg)error) {
	c.RegisterHandler(10602, func(ctx Context) error {
		msg := &pb.UnionBlessingRewardRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步妖盟散财数据
func (c *Client) OnUnionBlessingGiftSyncList(f func(ctx Context, msg *pb.UnionBlessingGiftSyncList)error) {
	c.RegisterHandler(10603, func(ctx Context) error {
		msg := &pb.UnionBlessingGiftSyncList{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步自己妖盟的散财的旗帜和时间
func (c *Client) OnUnionBlessingCountSyncList(f func(ctx Context, msg *pb.UnionBlessingCountSyncList)error) {
	c.RegisterHandler(10604, func(ctx Context) error {
		msg := &pb.UnionBlessingCountSyncList{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取签到奖励
func (c *Client) OnGoodFortuneGetRewardReq(f func(ctx Context, msg *pb.GoodFortuneGetRewardRsp)error) {
	c.RegisterHandler(10701, func(ctx Context) error {
		msg := &pb.GoodFortuneGetRewardRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 通关请求
func (c *Client) OnWestTravelPassGameReq(f func(ctx Context, msg *pb.WestTravelPassGameResp)error) {
	c.RegisterHandler(10901, func(ctx Context) error {
		msg := &pb.WestTravelPassGameResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 庆典玩家数据同步
func (c *Client) OnFestivalCelebrationsPlayerInfoMsg(f func(ctx Context, msg *pb.FestivalCelebrationsPlayerInfoMsg)error) {
	c.RegisterHandler(11000, func(ctx Context) error {
		msg := &pb.FestivalCelebrationsPlayerInfoMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 使用宴会道具
func (c *Client) OnFestivalCelebrationsUseBanquetItemReq(f func(ctx Context, msg *pb.FestivalCelebrationsUseBanquetItemRsp)error) {
	c.RegisterHandler(11001, func(ctx Context) error {
		msg := &pb.FestivalCelebrationsUseBanquetItemRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取宴会奖励
func (c *Client) OnFestivalCelebrationsDrawBanquetRewardReq(f func(ctx Context, msg *pb.FestivalCelebrationsDrawBanquetRewardRsp)error) {
	c.RegisterHandler(11002, func(ctx Context) error {
		msg := &pb.FestivalCelebrationsDrawBanquetRewardRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取积分详情
func (c *Client) OnFestivalCelebrationsGetBanquetScoreDetailReq(f func(ctx Context, msg *pb.FestivalCelebrationsGetBanquetScoreDetailRsp)error) {
	c.RegisterHandler(11003, func(ctx Context) error {
		msg := &pb.FestivalCelebrationsGetBanquetScoreDetailRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 使用福缘道具
func (c *Client) OnFestivalCelebrationsUseLuckyFateItemReq(f func(ctx Context, msg *pb.FestivalCelebrationsUseLuckyFateItemRsp)error) {
	c.RegisterHandler(11004, func(ctx Context) error {
		msg := &pb.FestivalCelebrationsUseLuckyFateItemRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取签到奖励
func (c *Client) OnFestivalCelebrationsGetSignRewardReq(f func(ctx Context, msg *pb.FestivalCelebrationsGetSignRewardRsp)error) {
	c.RegisterHandler(11005, func(ctx Context) error {
		msg := &pb.FestivalCelebrationsGetSignRewardRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求合成
func (c *Client) OnFestivalCelebrationsCollectSynthesisReq(f func(ctx Context, msg *pb.FestivalCelebrationsCollectSynthesisRsp)error) {
	c.RegisterHandler(11006, func(ctx Context) error {
		msg := &pb.FestivalCelebrationsCollectSynthesisRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求填充
func (c *Client) OnFestivalCelebrationsCollectFillReq(f func(ctx Context, msg *pb.FestivalCelebrationsCollectFillRsp)error) {
	c.RegisterHandler(11007, func(ctx Context) error {
		msg := &pb.FestivalCelebrationsCollectFillRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取大奖
func (c *Client) OnFestivalCelebrationsCollectDrawBigRewardReq(f func(ctx Context, msg *pb.FestivalCelebrationsCollectDrawBigRewardRsp)error) {
	c.RegisterHandler(11008, func(ctx Context) error {
		msg := &pb.FestivalCelebrationsCollectDrawBigRewardRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取领取名单
func (c *Client) OnFestivalCelebrationsCollectGetClaimListReq(f func(ctx Context, msg *pb.FestivalCelebrationsCollectGetClaimListRsp)error) {
	c.RegisterHandler(11009, func(ctx Context) error {
		msg := &pb.FestivalCelebrationsCollectGetClaimListRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 搜索玩家
func (c *Client) OnFestivalCelebrationsCollectSearchPlayerReq(f func(ctx Context, msg *pb.FestivalCelebrationsCollectSearchPlayerRsp)error) {
	c.RegisterHandler(11010, func(ctx Context) error {
		msg := &pb.FestivalCelebrationsCollectSearchPlayerRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 赠送道具
func (c *Client) OnFestivalCelebrationsCollectGiveReq(f func(ctx Context, msg *pb.FestivalCelebrationsCollectGiveRsp)error) {
	c.RegisterHandler(11011, func(ctx Context) error {
		msg := &pb.FestivalCelebrationsCollectGiveRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 索要道具
func (c *Client) OnFestivalCelebrationsCollectAskForReq(f func(ctx Context, msg *pb.FestivalCelebrationsCollectAskForRsp)error) {
	c.RegisterHandler(11012, func(ctx Context) error {
		msg := &pb.FestivalCelebrationsCollectAskForRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取彩蛋奖励
func (c *Client) OnFestivalCelebrationsEasterEggGetRewardReq(f func(ctx Context, msg *pb.FestivalCelebrationsEasterEggGetRewardRsp)error) {
	c.RegisterHandler(11013, func(ctx Context) error {
		msg := &pb.FestivalCelebrationsEasterEggGetRewardRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取赠送五福
func (c *Client) OnFestivalCelebrationsCollectGetGivenReq(f func(ctx Context, msg *pb.FestivalCelebrationsCollectGetGivenResp)error) {
	c.RegisterHandler(11014, func(ctx Context) error {
		msg := &pb.FestivalCelebrationsCollectGetGivenResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取福源奖励
func (c *Client) OnFestivalCelebrationsFuYuanGetRewardReq(f func(ctx Context, msg *pb.FestivalCelebrationsFuYuanGetRewardResp)error) {
	c.RegisterHandler(11015, func(ctx Context) error {
		msg := &pb.FestivalCelebrationsFuYuanGetRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取绑定妖盟成员列表
func (c *Client) OnFestivalCelebrationsGetBindUnionMemberDataListReq(f func(ctx Context, msg *pb.FestivalCelebrationsGetBindUnionMemberDataListRsp)error) {
	c.RegisterHandler(11016, func(ctx Context) error {
		msg := &pb.FestivalCelebrationsGetBindUnionMemberDataListRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取周年回忆条目
func (c *Client) OnFestivalCelebrationsGetAnnualMemoryReq(f func(ctx Context, msg *pb.FestivalCelebrationsGetAnnualMemoryRsp)error) {
	c.RegisterHandler(11021, func(ctx Context) error {
		msg := &pb.FestivalCelebrationsGetAnnualMemoryRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 膜拜
func (c *Client) OnFestivalCelebrationsWorshipReq(f func(ctx Context, msg *pb.FestivalCelebrationsWorshipResp)error) {
	c.RegisterHandler(11022, func(ctx Context) error {
		msg := &pb.FestivalCelebrationsWorshipResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 指定邀请
func (c *Client) OnDoubleDemonsAppointInviteReq(f func(ctx Context, msg *pb.DoubleDemonsAppointInviteResp)error) {
	c.RegisterHandler(11106, func(ctx Context) error {
		msg := &pb.DoubleDemonsAppointInviteResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 双妖成行登录同步数据
func (c *Client) OnDoubleDemonsDataLoginSync(f func(ctx Context, msg *pb.DoubleDemonsDataLoginSync)error) {
	c.RegisterHandler(11111, func(ctx Context) error {
		msg := &pb.DoubleDemonsDataLoginSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同意邀请/结为伴侣
func (c *Client) OnDoubleDemonsAgreeReq(f func(ctx Context, msg *pb.DoubleDemonsAgreeResp)error) {
	c.RegisterHandler(11112, func(ctx Context) error {
		msg := &pb.DoubleDemonsAgreeResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 一键拒绝
func (c *Client) OnDoubleDemonsRefuseReq(f func(ctx Context, msg *pb.DoubleDemonsRefuseResp)error) {
	c.RegisterHandler(11113, func(ctx Context) error {
		msg := &pb.DoubleDemonsRefuseResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取赠送奖励
func (c *Client) OnDoubleDemonsReceivePresentReq(f func(ctx Context, msg *pb.DoubleDemonsReceivePresentResp)error) {
	c.RegisterHandler(11114, func(ctx Context) error {
		msg := &pb.DoubleDemonsReceivePresentResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 新春红包 登录数据下发
func (c *Client) OnNewYearRedBagPlayerData(f func(ctx Context, msg *pb.NewYearRedBagPlayerData)error) {
	c.RegisterHandler(11500, func(ctx Context) error {
		msg := &pb.NewYearRedBagPlayerData{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 新春红包 进入系统
func (c *Client) OnNewYearRedBagEnterReqMsg(f func(ctx Context, msg *pb.NewYearRedBagEnterRespMsg)error) {
	c.RegisterHandler(11501, func(ctx Context) error {
		msg := &pb.NewYearRedBagEnterRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 新春红包 退出系统
func (c *Client) OnNewYearRedBagExitReqMsg(f func(ctx Context, msg *pb.NewYearRedBagExitRespMsg)error) {
	c.RegisterHandler(11502, func(ctx Context) error {
		msg := &pb.NewYearRedBagExitRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 新春红包 打开红包
func (c *Client) OnNewYearRedBagOpenReqMsg(f func(ctx Context, msg *pb.NewYearRedBagOpenRespMsg)error) {
	c.RegisterHandler(11503, func(ctx Context) error {
		msg := &pb.NewYearRedBagOpenRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 新春红包 同步公告
func (c *Client) OnNewYearSyncNoticeRespMsg(f func(ctx Context, msg *pb.NewYearSyncNoticeRespMsg)error) {
	c.RegisterHandler(11504, func(ctx Context) error {
		msg := &pb.NewYearSyncNoticeRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 穿戴内丹
func (c *Client) OnPetKerneCarryReq(f func(ctx Context, msg *pb.PetKerneCarryResp)error) {
	c.RegisterHandler(11702, func(ctx Context) error {
		msg := &pb.PetKerneCarryResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 内丹升级
func (c *Client) OnPetKerneUpLevelReq(f func(ctx Context, msg *pb.PetKerneUpLevelResp)error) {
	c.RegisterHandler(11703, func(ctx Context) error {
		msg := &pb.PetKerneUpLevelResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 内丹升星
func (c *Client) OnPetKerneUpStarReq(f func(ctx Context, msg *pb.PetKerneUpStarResp)error) {
	c.RegisterHandler(11704, func(ctx Context) error {
		msg := &pb.PetKerneUpStarResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 内丹激活
func (c *Client) OnPetKerneActiveReq(f func(ctx Context, msg *pb.PetKerneActiveResp)error) {
	c.RegisterHandler(11705, func(ctx Context) error {
		msg := &pb.PetKerneActiveResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 抽取灵兽内丹
func (c *Client) OnPetKernelDrawReq(f func(ctx Context, msg *pb.PetKernelDrawResp)error) {
	c.RegisterHandler(11706, func(ctx Context) error {
		msg := &pb.PetKernelDrawResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 内丹共鸣升级或激活
func (c *Client) OnPetKerneCombineUpLevelReq(f func(ctx Context, msg *pb.PetKerneCombineUpLevelResp)error) {
	c.RegisterHandler(11707, func(ctx Context) error {
		msg := &pb.PetKerneCombineUpLevelResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 内丹同步(商店是否开启 pieceShopOpen只有这个字段有用)
func (c *Client) OnPetKernelPlayerDataMsg(f func(ctx Context, msg *pb.PetKernelPlayerDataMsg)error) {
	c.RegisterHandler(11708, func(ctx Context) error {
		msg := &pb.PetKernelPlayerDataMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入门徒系统
func (c *Client) OnEnterPupilSystemReq(f func(ctx Context, msg *pb.EnterPupilSystemResp)error) {
	c.RegisterHandler(11801, func(ctx Context) error {
		msg := &pb.EnterPupilSystemResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 门徒培养
func (c *Client) OnPupilTrainReq(f func(ctx Context, msg *pb.PupilTrainResp)error) {
	c.RegisterHandler(11802, func(ctx Context) error {
		msg := &pb.PupilTrainResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 恢复培养次数
func (c *Client) OnPupilSiteTrainTimesRecoverReq(f func(ctx Context, msg *pb.PupilSiteTrainTimesRecoverResp)error) {
	c.RegisterHandler(11803, func(ctx Context) error {
		msg := &pb.PupilSiteTrainTimesRecoverResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 招收门徒
func (c *Client) OnPupilRecruitReq(f func(ctx Context, msg *pb.PupilRecruitResp)error) {
	c.RegisterHandler(11805, func(ctx Context) error {
		msg := &pb.PupilRecruitResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 门徒出师
func (c *Client) OnPupilGraduateReq(f func(ctx Context, msg *pb.PupilGraduateResp)error) {
	c.RegisterHandler(11806, func(ctx Context) error {
		msg := &pb.PupilGraduateResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取出师门徒列表
func (c *Client) OnPupilGetGraduatedListReq(f func(ctx Context, msg *pb.PupilGetGraduatedListResp)error) {
	c.RegisterHandler(11807, func(ctx Context) error {
		msg := &pb.PupilGetGraduatedListResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取上阵列表
func (c *Client) OnPupilGetFightListReq(f func(ctx Context, msg *pb.PupilGetFightListResp)error) {
	c.RegisterHandler(11808, func(ctx Context) error {
		msg := &pb.PupilGetFightListResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 门徒上阵
func (c *Client) OnPupilFightOnReq(f func(ctx Context, msg *pb.PupilFightOnResp)error) {
	c.RegisterHandler(11809, func(ctx Context) error {
		msg := &pb.PupilFightOnResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取门派统计信息
func (c *Client) OnPupilGetSectInfoReq(f func(ctx Context, msg *pb.PupilGetSectInfoResp)error) {
	c.RegisterHandler(11810, func(ctx Context) error {
		msg := &pb.PupilGetSectInfoResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 搜索玩家
func (c *Client) OnSearchPlayerReq(f func(ctx Context, msg *pb.SearchPlayerResp)error) {
	c.RegisterHandler(11811, func(ctx Context) error {
		msg := &pb.SearchPlayerResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 门徒系统登录同步
func (c *Client) OnPupilSystemLoginSync(f func(ctx Context, msg *pb.PupilSystemLoginSync)error) {
	c.RegisterHandler(11813, func(ctx Context) error {
		msg := &pb.PupilSystemLoginSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取广告奖励
func (c *Client) OnPupilGetAdRewardReq(f func(ctx Context, msg *pb.PupilGetAdRewardResp)error) {
	c.RegisterHandler(11814, func(ctx Context) error {
		msg := &pb.PupilGetAdRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 锁定弟子
func (c *Client) OnPupilLockReq(f func(ctx Context, msg *pb.PupilLockResp)error) {
	c.RegisterHandler(11815, func(ctx Context) error {
		msg := &pb.PupilLockResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 精准查找
func (c *Client) OnPupilExactSearchReq(f func(ctx Context, msg *pb.PupilExactSearchResp)error) {
	c.RegisterHandler(11816, func(ctx Context) error {
		msg := &pb.PupilExactSearchResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 结义用户数据同步
func (c *Client) OnMarriageUserDataMsgSync(f func(ctx Context, msg *pb.MarriageUserDataMsgSync)error) {
	c.RegisterHandler(11850, func(ctx Context) error {
		msg := &pb.MarriageUserDataMsgSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步其他玩家发起的结义请求（别的玩家指定你结义的列表）
func (c *Client) OnMarriageGetAppointApplySync(f func(ctx Context, msg *pb.MarriageGetAppointApplySync)error) {
	c.RegisterHandler(11851, func(ctx Context) error {
		msg := &pb.MarriageGetAppointApplySync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步出师未结义的门徒列表
func (c *Client) OnPupilGraduatedUnMarryListSync(f func(ctx Context, msg *pb.PupilGraduatedUnMarryListSync)error) {
	c.RegisterHandler(11852, func(ctx Context) error {
		msg := &pb.PupilGraduatedUnMarryListSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 结义记录数据同步
func (c *Client) OnMarriageRecordTempSync(f func(ctx Context, msg *pb.MarriageRecordTempSync)error) {
	c.RegisterHandler(11853, func(ctx Context) error {
		msg := &pb.MarriageRecordTempSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 拒绝结义通知同步
func (c *Client) OnMarriageRefuseNotifyMsgSync(f func(ctx Context, msg *pb.MarriageRefuseNotifyMsgSync)error) {
	c.RegisterHandler(11854, func(ctx Context) error {
		msg := &pb.MarriageRefuseNotifyMsgSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获得出师未结义的门徒列表
func (c *Client) OnPupilGraduatedUnMarryListReq(f func(ctx Context, msg *pb.PupilGraduatedUnMarryListResp)error) {
	c.RegisterHandler(11855, func(ctx Context) error {
		msg := &pb.PupilGraduatedUnMarryListResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 结义记录列表
func (c *Client) OnMarriageRecordListReq(f func(ctx Context, msg *pb.MarriageRecordListResp)error) {
	c.RegisterHandler(11856, func(ctx Context) error {
		msg := &pb.MarriageRecordListResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取全服结义请求列表(包括妖盟)
func (c *Client) OnMarriageGetServerApplyReq(f func(ctx Context, msg *pb.MarriageGetServerApplyResp)error) {
	c.RegisterHandler(11857, func(ctx Context) error {
		msg := &pb.MarriageGetServerApplyResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 指定结义推荐玩家请求
func (c *Client) OnMarriageRecommendPlayerReq(f func(ctx Context, msg *pb.MarriageRecommendPlayerResp)error) {
	c.RegisterHandler(11858, func(ctx Context) error {
		msg := &pb.MarriageRecommendPlayerResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取其他玩家发起的结义请求（别的玩家指定你结义的列表）
func (c *Client) OnMarriageGetAppointApplyReq(f func(ctx Context, msg *pb.MarriageGetAppointApplyResp)error) {
	c.RegisterHandler(11859, func(ctx Context) error {
		msg := &pb.MarriageGetAppointApplyResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 处理结义请求
func (c *Client) OnMarriageApplyDealReq(f func(ctx Context, msg *pb.MarriageApplyDealResp)error) {
	c.RegisterHandler(11860, func(ctx Context) error {
		msg := &pb.MarriageApplyDealResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 发布结义请求
func (c *Client) OnMarriageApplyPublishReq(f func(ctx Context, msg *pb.MarriageApplyPublishResp)error) {
	c.RegisterHandler(11861, func(ctx Context) error {
		msg := &pb.MarriageApplyPublishResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 取消结义请求
func (c *Client) OnMarriageApplyCancelReq(f func(ctx Context, msg *pb.MarriageApplyCancelResp)error) {
	c.RegisterHandler(11862, func(ctx Context) error {
		msg := &pb.MarriageApplyCancelResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 刷新全服结义列表请求
func (c *Client) OnMarriageRefreshServerApplyReq(f func(ctx Context, msg *pb.MarriageRefreshServerApplyResp)error) {
	c.RegisterHandler(11863, func(ctx Context) error {
		msg := &pb.MarriageRefreshServerApplyResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 设置是否接受指向结义请求
func (c *Client) OnMarriageSetAppointReq(f func(ctx Context, msg *pb.MarriageSetAppointResp)error) {
	c.RegisterHandler(11864, func(ctx Context) error {
		msg := &pb.MarriageSetAppointResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步其他玩家发起的结义请求时间
func (c *Client) OnMarriageGetAppointApplyTimeSync(f func(ctx Context, msg *pb.MarriageGetAppointApplyTimeSync)error) {
	c.RegisterHandler(11865, func(ctx Context) error {
		msg := &pb.MarriageGetAppointApplyTimeSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 洪荒灵兽灵根升级请求
func (c *Client) OnPetRootUpReq(f func(ctx Context, msg *pb.PetRootUpResp)error) {
	c.RegisterHandler(12201, func(ctx Context) error {
		msg := &pb.PetRootUpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 洪荒灵兽觉醒请求
func (c *Client) OnPetAwakeReq(f func(ctx Context, msg *pb.PetAwakeResp)error) {
	c.RegisterHandler(12202, func(ctx Context) error {
		msg := &pb.PetAwakeResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 洪荒灵兽缘分升级请求
func (c *Client) OnPetFateUpReq(f func(ctx Context, msg *pb.PetFateUpResp)error) {
	c.RegisterHandler(12203, func(ctx Context) error {
		msg := &pb.PetFateUpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 切换灵兽皮肤
func (c *Client) OnPetSwitchSkinReq(f func(ctx Context, msg *pb.PetSwitchSkinResp)error) {
	c.RegisterHandler(12204, func(ctx Context) error {
		msg := &pb.PetSwitchSkinResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 激活、升级灵兽皮肤
func (c *Client) OnPetUpSkinLevelReq(f func(ctx Context, msg *pb.PetUpSkinLevelResp)error) {
	c.RegisterHandler(12205, func(ctx Context) error {
		msg := &pb.PetUpSkinLevelResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 激活、升级灵兽羁绊
func (c *Client) OnPetUpSkinCombineReq(f func(ctx Context, msg *pb.PetUpSkinCombineResp)error) {
	c.RegisterHandler(12206, func(ctx Context) error {
		msg := &pb.PetUpSkinCombineResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取微信排行榜白名单
func (c *Client) OnGetWechatRankWhiteListReq(f func(ctx Context, msg *pb.GetWechatRankWhiteListRsp)error) {
	c.RegisterHandler(12601, func(ctx Context) error {
		msg := &pb.GetWechatRankWhiteListRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 筋斗云注灵数据同步
func (c *Client) OnCloudRefinePlayerDataMsg(f func(ctx Context, msg *pb.CloudRefinePlayerDataMsg)error) {
	c.RegisterHandler(12900, func(ctx Context) error {
		msg := &pb.CloudRefinePlayerDataMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 筋斗云注灵请求
func (c *Client) OnCloudRefineReq(f func(ctx Context, msg *pb.CloudRefineResp)error) {
	c.RegisterHandler(12901, func(ctx Context) error {
		msg := &pb.CloudRefineResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 筋斗云注灵渡劫请求
func (c *Client) OnCloudRefineStarLvUpReq(f func(ctx Context, msg *pb.CloudRefineStarLvUpResp)error) {
	c.RegisterHandler(12902, func(ctx Context) error {
		msg := &pb.CloudRefineStarLvUpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入游戏
func (c *Client) OnComposeBallEnterGameReq(f func(ctx Context, msg *pb.ComposeBallEnterGameResp)error) {
	c.RegisterHandler(13001, func(ctx Context) error {
		msg := &pb.ComposeBallEnterGameResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求合成
func (c *Client) OnComposeBallComposeReq(f func(ctx Context, msg *pb.ComposeBallComposeResp)error) {
	c.RegisterHandler(13002, func(ctx Context) error {
		msg := &pb.ComposeBallComposeResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求使用道具
func (c *Client) OnComposeBallUseItemReq(f func(ctx Context, msg *pb.ComposeBallUseItemResp)error) {
	c.RegisterHandler(13003, func(ctx Context) error {
		msg := &pb.ComposeBallUseItemResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求增加体力
func (c *Client) OnComposeBallHpItemReq(f func(ctx Context, msg *pb.ComposeBallHpItemResp)error) {
	c.RegisterHandler(13004, func(ctx Context) error {
		msg := &pb.ComposeBallHpItemResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求结算
func (c *Client) OnComposeBallSettleReq(f func(ctx Context, msg *pb.ComposeBallSettleResp)error) {
	c.RegisterHandler(13005, func(ctx Context) error {
		msg := &pb.ComposeBallSettleResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入活动
func (c *Client) OnMonopolyEnterActivityReq(f func(ctx Context, msg *pb.MonopolyEnterActivityResp)error) {
	c.RegisterHandler(13101, func(ctx Context) error {
		msg := &pb.MonopolyEnterActivityResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入地图
func (c *Client) OnMonopolyEnterMapReq(f func(ctx Context, msg *pb.MonopolyEnterMapResp)error) {
	c.RegisterHandler(13102, func(ctx Context) error {
		msg := &pb.MonopolyEnterMapResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 摇色子
func (c *Client) OnMonopolyRollDiceReq(f func(ctx Context, msg *pb.MonopolyRollDiceResp)error) {
	c.RegisterHandler(13103, func(ctx Context) error {
		msg := &pb.MonopolyRollDiceResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取协助列表
func (c *Client) OnMonopolyAssistListReq(f func(ctx Context, msg *pb.MonopolyAssistListResp)error) {
	c.RegisterHandler(13104, func(ctx Context) error {
		msg := &pb.MonopolyAssistListResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 使用道具 补充体力
func (c *Client) OnMonopolyReplenishStrengthReq(f func(ctx Context, msg *pb.MonopolyReplenishStrengthResp)error) {
	c.RegisterHandler(13105, func(ctx Context) error {
		msg := &pb.MonopolyReplenishStrengthResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 协助队友出陷阱
func (c *Client) OnMonopolyRescueTrapReq(f func(ctx Context, msg *pb.MonopolyRescueTrapResp)error) {
	c.RegisterHandler(13107, func(ctx Context) error {
		msg := &pb.MonopolyRescueTrapResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取可掠夺的妖盟列表
func (c *Client) OnMonopolyRobListReq(f func(ctx Context, msg *pb.MonopolyRobListResp)error) {
	c.RegisterHandler(13108, func(ctx Context) error {
		msg := &pb.MonopolyRobListResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 掠夺其他妖盟
func (c *Client) OnMonopolyRobReq(f func(ctx Context, msg *pb.MonopolyRobResp)error) {
	c.RegisterHandler(13110, func(ctx Context) error {
		msg := &pb.MonopolyRobResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取妖盟日志列表
func (c *Client) OnMonopolyUnionLogListReq(f func(ctx Context, msg *pb.MonopolyUnionLogListResp)error) {
	c.RegisterHandler(13111, func(ctx Context) error {
		msg := &pb.MonopolyUnionLogListResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取玩家日志详情
func (c *Client) OnMonopolyPlayerLogDetailReq(f func(ctx Context, msg *pb.MonopolyPlayerLogDetailResp)error) {
	c.RegisterHandler(13112, func(ctx Context) error {
		msg := &pb.MonopolyPlayerLogDetailResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 协助队友攻击怪物
func (c *Client) OnMonopolyAssistAttackMonsterReq(f func(ctx Context, msg *pb.MonopolyAssistAttackMonsterResp)error) {
	c.RegisterHandler(13113, func(ctx Context) error {
		msg := &pb.MonopolyAssistAttackMonsterResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取队友帮忙打死怪后的奖励
func (c *Client) OnMonopolyReceiveMonsterRewardReq(f func(ctx Context, msg *pb.MonopolyReceiveMonsterRewardResp)error) {
	c.RegisterHandler(13114, func(ctx Context) error {
		msg := &pb.MonopolyReceiveMonsterRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 红点详情
func (c *Client) OnMonopolyRedPointReqMsg(f func(ctx Context, msg *pb.MonopolyRedPointRespMsg)error) {
	c.RegisterHandler(13115, func(ctx Context) error {
		msg := &pb.MonopolyRedPointRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 遥控骰子
func (c *Client) OnMonopolyRemoteRollDiceReq(f func(ctx Context, msg *pb.MonopolyRemoteRollDiceResp)error) {
	c.RegisterHandler(13116, func(ctx Context) error {
		msg := &pb.MonopolyRemoteRollDiceResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入其他妖盟地图，请求详细数据
func (c *Client) OnMonopolyEnterRobMapReq(f func(ctx Context, msg *pb.MonopolyEnterRobMapResp)error) {
	c.RegisterHandler(13117, func(ctx Context) error {
		msg := &pb.MonopolyEnterRobMapResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 自动操作解锁
func (c *Client) OnMonopolyAutoUnlockReq(f func(ctx Context, msg *pb.MonopolyAutoUnlockResp)error) {
	c.RegisterHandler(13118, func(ctx Context) error {
		msg := &pb.MonopolyAutoUnlockResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取天赐福源列表
func (c *Client) OnMonopolyBlessingListReq(f func(ctx Context, msg *pb.MonopolyBlessingListResp)error) {
	c.RegisterHandler(13119, func(ctx Context) error {
		msg := &pb.MonopolyBlessingListResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取天赐祥瑞奖励
func (c *Client) OnMonopolyReceiveBlessingReq(f func(ctx Context, msg *pb.MonopolyReceiveBlessingResp)error) {
	c.RegisterHandler(13120, func(ctx Context) error {
		msg := &pb.MonopolyReceiveBlessingResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 神行符移动
func (c *Client) OnMonopolyQuickMoveReq(f func(ctx Context, msg *pb.MonopolyQuickMoveResp)error) {
	c.RegisterHandler(13121, func(ctx Context) error {
		msg := &pb.MonopolyQuickMoveResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 暗格移动
func (c *Client) OnMonopolyDarkGridMoveReq(f func(ctx Context, msg *pb.MonopolyDarkGridMoveResp)error) {
	c.RegisterHandler(13122, func(ctx Context) error {
		msg := &pb.MonopolyDarkGridMoveResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 处理事件同一接口
func (c *Client) OnMonopolyEventHandleReq(f func(ctx Context, msg *pb.MonopolyEventHandleResp)error) {
	c.RegisterHandler(13123, func(ctx Context) error {
		msg := &pb.MonopolyEventHandleResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看怪物详情
func (c *Client) OnMonopolyMonsterAttrReq(f func(ctx Context, msg *pb.MonopolyMonsterAttrResp)error) {
	c.RegisterHandler(13124, func(ctx Context) error {
		msg := &pb.MonopolyMonsterAttrResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取个人排名
func (c *Client) OnMonopolyGetPlayerRankReq(f func(ctx Context, msg *pb.MonopolyGetPlayerRankResp)error) {
	c.RegisterHandler(13125, func(ctx Context) error {
		msg := &pb.MonopolyGetPlayerRankResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取仇人列表
func (c *Client) OnMonopolyGetEnemyListReq(f func(ctx Context, msg *pb.MonopolyGetEnemyListResp)error) {
	c.RegisterHandler(13126, func(ctx Context) error {
		msg := &pb.MonopolyGetEnemyListResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 积分详情
func (c *Client) OnMonopolyScoreDetailReq(f func(ctx Context, msg *pb.MonopolyScoreDetailResp)error) {
	c.RegisterHandler(13127, func(ctx Context) error {
		msg := &pb.MonopolyScoreDetailResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 建筑积分详情
func (c *Client) OnMonopolyBuildingScoreDetailReq(f func(ctx Context, msg *pb.MonopolyBuildingScoreDetailResp)error) {
	c.RegisterHandler(13128, func(ctx Context) error {
		msg := &pb.MonopolyBuildingScoreDetailResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 弹幕推送
func (c *Client) OnMonopolyBarrageNotify(f func(ctx Context, msg *pb.MonopolyBarrageNotify)error) {
	c.RegisterHandler(13130, func(ctx Context) error {
		msg := &pb.MonopolyBarrageNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同妖盟移动同步
func (c *Client) OnMonopolyMoveNotify(f func(ctx Context, msg *pb.MonopolyMoveNotify)error) {
	c.RegisterHandler(13131, func(ctx Context) error {
		msg := &pb.MonopolyMoveNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 如果队友主动或者被协助 挣脱了陷阱 同步
func (c *Client) OnMonopolyNotifyPlayerEndTrap(f func(ctx Context, msg *pb.MonopolyNotifyPlayerEndTrap)error) {
	c.RegisterHandler(13132, func(ctx Context) error {
		msg := &pb.MonopolyNotifyPlayerEndTrap{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 触发建筑升级的推送
func (c *Client) OnMonopolyNotifyBuildingUpgrade(f func(ctx Context, msg *pb.MonopolyNotifyBuildingUpgrade)error) {
	c.RegisterHandler(13133, func(ctx Context) error {
		msg := &pb.MonopolyNotifyBuildingUpgrade{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 天赐福源推送
func (c *Client) OnMonopolySendBlessingNotify(f func(ctx Context, msg *pb.MonopolySendBlessingNotify)error) {
	c.RegisterHandler(13135, func(ctx Context) error {
		msg := &pb.MonopolySendBlessingNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟积分变动
func (c *Client) OnMonopolyScoreNotify(f func(ctx Context, msg *pb.MonopolyScoreNotify)error) {
	c.RegisterHandler(13136, func(ctx Context) error {
		msg := &pb.MonopolyScoreNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步排名信息
func (c *Client) OnMonopolyRankNotify(f func(ctx Context, msg *pb.MonopolyRankNotify)error) {
	c.RegisterHandler(13137, func(ctx Context) error {
		msg := &pb.MonopolyRankNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 天赐福源推送-我的
func (c *Client) OnMonopolyMyBlessingNotify(f func(ctx Context, msg *pb.MonopolyMyBlessingNotify)error) {
	c.RegisterHandler(13139, func(ctx Context) error {
		msg := &pb.MonopolyMyBlessingNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 协助红点推送
func (c *Client) OnMonopolyAssistRedNotify(f func(ctx Context, msg *pb.MonopolyAssistRedNotify)error) {
	c.RegisterHandler(13140, func(ctx Context) error {
		msg := &pb.MonopolyAssistRedNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 分组内积分详情
func (c *Client) OnMonopolyGetGroupDetailReq(f func(ctx Context, msg *pb.MonopolyGetGroupDetailResp)error) {
	c.RegisterHandler(13141, func(ctx Context) error {
		msg := &pb.MonopolyGetGroupDetailResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 膜拜
func (c *Client) OnMonopolyWorshipReq(f func(ctx Context, msg *pb.MonopolyWorshipResp)error) {
	c.RegisterHandler(13142, func(ctx Context) error {
		msg := &pb.MonopolyWorshipResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 协助击杀后积分推送
func (c *Client) OnMonopolyPlayerScoreChangeNotify(f func(ctx Context, msg *pb.MonopolyPlayerScoreChangeNotify)error) {
	c.RegisterHandler(13144, func(ctx Context) error {
		msg := &pb.MonopolyPlayerScoreChangeNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 竞猜正确玩家数据
func (c *Client) OnMonopolyGuessPlayersReq(f func(ctx Context, msg *pb.MonopolyGuessPlayersResp)error) {
	c.RegisterHandler(13145, func(ctx Context) error {
		msg := &pb.MonopolyGuessPlayersResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求竞猜界面数据
func (c *Client) OnMonopolyGetGuessDataReq(f func(ctx Context, msg *pb.MonopolyGetGuessDataResp)error) {
	c.RegisterHandler(13146, func(ctx Context) error {
		msg := &pb.MonopolyGetGuessDataResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求竞猜选择
func (c *Client) OnMonopolyGuessSelectReq(f func(ctx Context, msg *pb.MonopolyGuessSelectResp)error) {
	c.RegisterHandler(13147, func(ctx Context) error {
		msg := &pb.MonopolyGuessSelectResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求竞猜奖励
func (c *Client) OnMonopolyGuessRewardReq(f func(ctx Context, msg *pb.MonopolyGuessRewardResp)error) {
	c.RegisterHandler(13148, func(ctx Context) error {
		msg := &pb.MonopolyGuessRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 镇魔登录同步
func (c *Client) OnTownDemonApplyDataSyncReq(f func(ctx Context, msg *pb.TownDemonApplyDataSync)error) {
	c.RegisterHandler(13250, func(ctx Context) error {
		msg := &pb.TownDemonApplyDataSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 镇魔登录时间戳同步
func (c *Client) OnTownDemonTimeStampsDataSync(f func(ctx Context, msg *pb.TownDemonTimeStampsDataSync)error) {
	c.RegisterHandler(13251, func(ctx Context) error {
		msg := &pb.TownDemonTimeStampsDataSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 宴会开启同步(类仙宫送福)
func (c *Client) OnPeachBanquetOpenSync(f func(ctx Context, msg *pb.PeachBanquetOpenSync)error) {
	c.RegisterHandler(13510, func(ctx Context) error {
		msg := &pb.PeachBanquetOpenSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蟠桃宴登录同步
func (c *Client) OnPeachBanquetLoginSync(f func(ctx Context, msg *pb.PeachBanquetLoginSync)error) {
	c.RegisterHandler(13515, func(ctx Context) error {
		msg := &pb.PeachBanquetLoginSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 蟠桃宴道具同步
func (c *Client) OnPeachBanquetItemSync(f func(ctx Context, msg *pb.PeachBanquetItemSync)error) {
	c.RegisterHandler(13519, func(ctx Context) error {
		msg := &pb.PeachBanquetItemSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 宴会道具使用
func (c *Client) OnUsePeachBanquetItemReq(f func(ctx Context, msg *pb.UsePeachBanquetItemResp)error) {
	c.RegisterHandler(13520, func(ctx Context) error {
		msg := &pb.UsePeachBanquetItemResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟悬赏-打开护送地图
func (c *Client) OnUnionBountyEnterMapReq(f func(ctx Context, msg *pb.UnionBountyEnterMapResp)error) {
	c.RegisterHandler(13602, func(ctx Context) error {
		msg := &pb.UnionBountyEnterMapResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟悬赏-离开护送地图
func (c *Client) OnUnionBountyExitMapReq(f func(ctx Context, msg *pb.UnionBountyExitMapResp)error) {
	c.RegisterHandler(13603, func(ctx Context) error {
		msg := &pb.UnionBountyExitMapResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟悬赏-掠夺某个镖车
func (c *Client) OnUnionBountyPlunderReq(f func(ctx Context, msg *pb.UnionBountyPlunderResp)error) {
	c.RegisterHandler(13604, func(ctx Context) error {
		msg := &pb.UnionBountyPlunderResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟悬赏-处理悬赏事件
func (c *Client) OnUnionBountyDealBountyReq(f func(ctx Context, msg *pb.UnionBountyDealBountyResp)error) {
	c.RegisterHandler(13605, func(ctx Context) error {
		msg := &pb.UnionBountyDealBountyResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟悬赏-领取护送镖车奖励
func (c *Client) OnUnionBountyGetRewardEscortReq(f func(ctx Context, msg *pb.UnionBountyGetRewardEscortResp)error) {
	c.RegisterHandler(13606, func(ctx Context) error {
		msg := &pb.UnionBountyGetRewardEscortResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟悬赏-查找某个玩家的镖车
func (c *Client) OnUnionBountyCheckEscortReq(f func(ctx Context, msg *pb.UnionBountyCheckEscortResp)error) {
	c.RegisterHandler(13608, func(ctx Context) error {
		msg := &pb.UnionBountyCheckEscortResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟悬赏-打开战报
func (c *Client) OnUnionBountyGetReportReq(f func(ctx Context, msg *pb.UnionBountyGetReportResp)error) {
	c.RegisterHandler(13609, func(ctx Context) error {
		msg := &pb.UnionBountyGetReportResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟悬赏-战报中打开某个玩家的镖车
func (c *Client) OnUnionBountyReportCheckEscortReq(f func(ctx Context, msg *pb.UnionBountyReportCheckEscortResp)error) {
	c.RegisterHandler(13611, func(ctx Context) error {
		msg := &pb.UnionBountyReportCheckEscortResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟悬赏-排行榜膜拜
func (c *Client) OnUnionBountyWorshipReq(f func(ctx Context, msg *pb.UnionBountyWorshipResp)error) {
	c.RegisterHandler(13612, func(ctx Context) error {
		msg := &pb.UnionBountyWorshipResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟悬赏-刷新护送地图
func (c *Client) OnUnionBountyRefreshMapReq(f func(ctx Context, msg *pb.UnionBountyRefreshMapResp)error) {
	c.RegisterHandler(13613, func(ctx Context) error {
		msg := &pb.UnionBountyRefreshMapResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟悬赏-打开妖盟战报
func (c *Client) OnUnionBountyOpenBountyEventReq(f func(ctx Context, msg *pb.UnionBountyOpenBountyEventResp)error) {
	c.RegisterHandler(13614, func(ctx Context) error {
		msg := &pb.UnionBountyOpenBountyEventResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟悬赏-反击
func (c *Client) OnUnionBountyRetaliateReq(f func(ctx Context, msg *pb.UnionBountyRetaliateResp)error) {
	c.RegisterHandler(13615, func(ctx Context) error {
		msg := &pb.UnionBountyRetaliateResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟悬赏-查看联盟积分
func (c *Client) OnUnionBountyGetMemberScoreReq(f func(ctx Context, msg *pb.UnionBountyGetMemberScoreResp)error) {
	c.RegisterHandler(13616, func(ctx Context) error {
		msg := &pb.UnionBountyGetMemberScoreResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟悬赏-是否有正在押镖
func (c *Client) OnUnionBountyHaveEscortReq(f func(ctx Context, msg *pb.UnionBountyHaveEscortResp)error) {
	c.RegisterHandler(13617, func(ctx Context) error {
		msg := &pb.UnionBountyHaveEscortResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟悬赏-查看镖车详情
func (c *Client) OnUnionBountyOpenEscortCartDetailReq(f func(ctx Context, msg *pb.UnionBountyOpenEscortCartDetailRsp)error) {
	c.RegisterHandler(13618, func(ctx Context) error {
		msg := &pb.UnionBountyOpenEscortCartDetailRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟悬赏-打开妖兽界面
func (c *Client) OnUnionBountyOpenMonsterReq(f func(ctx Context, msg *pb.UnionBountyOpenMonsterResp)error) {
	c.RegisterHandler(13619, func(ctx Context) error {
		msg := &pb.UnionBountyOpenMonsterResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟悬赏-打妖兽
func (c *Client) OnUnionBountyAttackMonsterReq(f func(ctx Context, msg *pb.UnionBountyAttackMonsterResp)error) {
	c.RegisterHandler(13620, func(ctx Context) error {
		msg := &pb.UnionBountyAttackMonsterResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟悬赏-邀请
func (c *Client) OnUnionBountyAskHelpReq(f func(ctx Context, msg *pb.UnionBountyAskHelpResp)error) {
	c.RegisterHandler(13621, func(ctx Context) error {
		msg := &pb.UnionBountyAskHelpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟悬赏-打开邀请界面
func (c *Client) OnUnionBountyOpenAskHelpReq(f func(ctx Context, msg *pb.UnionBountyOpenAskHelpResp)error) {
	c.RegisterHandler(13622, func(ctx Context) error {
		msg := &pb.UnionBountyOpenAskHelpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟悬赏-处理邀请（同意加入、拒绝、一键拒绝）
func (c *Client) OnUnionBountyDealAskHelpReq(f func(ctx Context, msg *pb.UnionBountyDealAskHelpResp)error) {
	c.RegisterHandler(13623, func(ctx Context) error {
		msg := &pb.UnionBountyDealAskHelpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟悬赏-队伍管理（加入队伍、离开队伍、踢出队伍）
func (c *Client) OnUnionBountyManageAskHelpReq(f func(ctx Context, msg *pb.UnionBountyManageAskHelpResp)error) {
	c.RegisterHandler(13624, func(ctx Context) error {
		msg := &pb.UnionBountyManageAskHelpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟悬赏-同步妖兽房间信息
func (c *Client) OnUnionBountySynMonsterInfoMsg(f func(ctx Context, msg *pb.UnionBountySynMonsterInfoMsg)error) {
	c.RegisterHandler(13625, func(ctx Context) error {
		msg := &pb.UnionBountySynMonsterInfoMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟悬赏-查看当前妖兽的属性
func (c *Client) OnUnionBountyGetMonsterAttributeReq(f func(ctx Context, msg *pb.UnionBountyGetMonsterAttributeResp)error) {
	c.RegisterHandler(13626, func(ctx Context) error {
		msg := &pb.UnionBountyGetMonsterAttributeResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟悬赏-变更参战玩家位置
func (c *Client) OnUnionBountyMonsterChangePosReq(f func(ctx Context, msg *pb.UnionBountyMonsterChangePosRsp)error) {
	c.RegisterHandler(13627, func(ctx Context) error {
		msg := &pb.UnionBountyMonsterChangePosRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取当前存在的妖兽
func (c *Client) OnUnionBountyGetExistMonsterReq(f func(ctx Context, msg *pb.UnionBountyGetExistMonsterResp)error) {
	c.RegisterHandler(13628, func(ctx Context) error {
		msg := &pb.UnionBountyGetExistMonsterResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入宗门大比
func (c *Client) OnEnterPupilRankActivityReq(f func(ctx Context, msg *pb.EnterPupilRankActivityResp)error) {
	c.RegisterHandler(13801, func(ctx Context) error {
		msg := &pb.EnterPupilRankActivityResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取涨幅记录
func (c *Client) OnGetIncreaseRecordReq(f func(ctx Context, msg *pb.GetIncreaseRecordResp)error) {
	c.RegisterHandler(13802, func(ctx Context) error {
		msg := &pb.GetIncreaseRecordResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看 排行榜上玩家的 涨幅记录
func (c *Client) OnPupilRankDetailReq(f func(ctx Context, msg *pb.PupilRankDetailResp)error) {
	c.RegisterHandler(13803, func(ctx Context) error {
		msg := &pb.PupilRankDetailResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取世交列表
func (c *Client) OnFriendGetListReq(f func(ctx Context, msg *pb.FriendGetListResp)error) {
	c.RegisterHandler(14101, func(ctx Context) error {
		msg := &pb.FriendGetListResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 赠送
func (c *Client) OnFriendSendReq(f func(ctx Context, msg *pb.FriendSendResp)error) {
	c.RegisterHandler(14102, func(ctx Context) error {
		msg := &pb.FriendSendResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取
func (c *Client) OnFriendReceiveReq(f func(ctx Context, msg *pb.FriendReceiveResp)error) {
	c.RegisterHandler(14103, func(ctx Context) error {
		msg := &pb.FriendReceiveResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 一键赠送、领取
func (c *Client) OnFriendReceiveAllReq(f func(ctx Context, msg *pb.FriendReceiveAllResp)error) {
	c.RegisterHandler(14104, func(ctx Context) error {
		msg := &pb.FriendReceiveAllResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 删除世交
func (c *Client) OnFriendDeleteReq(f func(ctx Context, msg *pb.FriendDeleteResp)error) {
	c.RegisterHandler(14105, func(ctx Context) error {
		msg := &pb.FriendDeleteResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 好友消息列表
func (c *Client) OnFriendChatReqMsg(f func(ctx Context, msg *pb.WorldMessageListMsg)error) {
	c.RegisterHandler(14120, func(ctx Context) error {
		msg := &pb.WorldMessageListMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步世交信息列表
func (c *Client) OnSyncFriendList(f func(ctx Context, msg *pb.SyncFriendList)error) {
	c.RegisterHandler(14130, func(ctx Context) error {
		msg := &pb.SyncFriendList{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 好友消息同步
func (c *Client) OnWorldMessageListMsg214131(f func(ctx Context, msg *pb.WorldMessageListMsg)error) {
	c.RegisterHandler(14131, func(ctx Context) error {
		msg := &pb.WorldMessageListMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步世交信息
func (c *Client) OnSyncFriendInfo(f func(ctx Context, msg *pb.SyncFriendInfo)error) {
	c.RegisterHandler(14132, func(ctx Context) error {
		msg := &pb.SyncFriendInfo{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 飞升仙界-同步玩家数据
func (c *Client) OnPlayerFairyLandDataMsg(f func(ctx Context, msg *pb.PlayerFairyLandDataMsg)error) {
	c.RegisterHandler(14202, func(ctx Context) error {
		msg := &pb.PlayerFairyLandDataMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 飞升仙界-南天门协助
func (c *Client) OnFairyLandSouthDoorHelpReq(f func(ctx Context, msg *pb.FairyLandSouthDoorHelpResp)error) {
	c.RegisterHandler(14208, func(ctx Context) error {
		msg := &pb.FairyLandSouthDoorHelpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 飞升仙界-区服列表同步
func (c *Client) OnFairyLandServerSync(f func(ctx Context, msg *pb.FairyLandServerSync)error) {
	c.RegisterHandler(14211, func(ctx Context) error {
		msg := &pb.FairyLandServerSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步小世界信息
func (c *Client) OnUniverseDataMsgSync(f func(ctx Context, msg *pb.UniverseDataMsgSync)error) {
	c.RegisterHandler(14302, func(ctx Context) error {
		msg := &pb.UniverseDataMsgSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 小世界升级
func (c *Client) OnUniverseLevelUpReq(f func(ctx Context, msg *pb.UniverseLevelUpResp)error) {
	c.RegisterHandler(14303, func(ctx Context) error {
		msg := &pb.UniverseLevelUpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 小世界轮盘抽奖
func (c *Client) OnUniverseDrawReq(f func(ctx Context, msg *pb.UniverseDrawResp)error) {
	c.RegisterHandler(14304, func(ctx Context) error {
		msg := &pb.UniverseDrawResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 小世界技能激活
func (c *Client) OnUniverseSkillUnlockReq(f func(ctx Context, msg *pb.UniverseSkillUnlockResp)error) {
	c.RegisterHandler(14305, func(ctx Context) error {
		msg := &pb.UniverseSkillUnlockResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 小世界技能上阵
func (c *Client) OnEquipUniverseSkillReq(f func(ctx Context, msg *pb.EquipUniverseSkillResp)error) {
	c.RegisterHandler(14306, func(ctx Context) error {
		msg := &pb.EquipUniverseSkillResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 小世界技能共鸣激活、升级
func (c *Client) OnUniverseSkillCombineLvUpReq(f func(ctx Context, msg *pb.UniverseSkillCombineLvUpResp)error) {
	c.RegisterHandler(14307, func(ctx Context) error {
		msg := &pb.UniverseSkillCombineLvUpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 洞察天机抽取
func (c *Client) OnUniverseSkillDrawReq(f func(ctx Context, msg *pb.UniverseSkillDrawResp)error) {
	c.RegisterHandler(14308, func(ctx Context) error {
		msg := &pb.UniverseSkillDrawResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 小世界技能升级
func (c *Client) OnUniverseSkillLvUpReq(f func(ctx Context, msg *pb.UniverseSkillLvUpResp)error) {
	c.RegisterHandler(14309, func(ctx Context) error {
		msg := &pb.UniverseSkillLvUpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 天地轮盘二次抽奖
func (c *Client) OnUniverseDrawTwiceReq(f func(ctx Context, msg *pb.UniverseDrawTwiceResp)error) {
	c.RegisterHandler(14310, func(ctx Context) error {
		msg := &pb.UniverseDrawTwiceResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取玩家结算信息
func (c *Client) OnGodDemonBattleGetPlayerSettlementReq(f func(ctx Context, msg *pb.GodDemonBattleGetPlayerSettlementResp)error) {
	c.RegisterHandler(14527, func(ctx Context) error {
		msg := &pb.GodDemonBattleGetPlayerSettlementResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取阵营结算信息
func (c *Client) OnGodDemonBattleGetCampSettlementReq(f func(ctx Context, msg *pb.GodDemonBattleGetCampSettlementResp)error) {
	c.RegisterHandler(14528, func(ctx Context) error {
		msg := &pb.GodDemonBattleGetCampSettlementResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 画中仙境-检测仙泽
func (c *Client) OnCheckFairyPoolReq(f func(ctx Context, msg *pb.CheckFairyPoolResp)error) {
	c.RegisterHandler(14698, func(ctx Context) error {
		msg := &pb.CheckFairyPoolResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 画中仙境被抢夺同步
func (c *Client) OnBeingSnatchedSync(f func(ctx Context, msg *pb.BeingSnatchedSync)error) {
	c.RegisterHandler(14699, func(ctx Context) error {
		msg := &pb.BeingSnatchedSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入活动
func (c *Client) OnSkyTradeEnterActivityReq(f func(ctx Context, msg *pb.SkyTradeEnterActivityResp)error) {
	c.RegisterHandler(15000, func(ctx Context) error {
		msg := &pb.SkyTradeEnterActivityResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入地图
func (c *Client) OnSkyTradeEnterReq(f func(ctx Context, msg *pb.SkyTradeEnterResp)error) {
	c.RegisterHandler(15001, func(ctx Context) error {
		msg := &pb.SkyTradeEnterResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看分组信息
func (c *Client) OnSkyTradeGroupInfoReq(f func(ctx Context, msg *pb.SkyTradeGroupInfoResp)error) {
	c.RegisterHandler(15002, func(ctx Context) error {
		msg := &pb.SkyTradeGroupInfoResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 加速移动
func (c *Client) OnSkyTradeAddSpeedReq(f func(ctx Context, msg *pb.SkyTradeAddSpeedResp)error) {
	c.RegisterHandler(15003, func(ctx Context) error {
		msg := &pb.SkyTradeAddSpeedResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求移动
func (c *Client) OnSkyTradeGotoPortReq(f func(ctx Context, msg *pb.SkyTradeGotoPortResp)error) {
	c.RegisterHandler(15004, func(ctx Context) error {
		msg := &pb.SkyTradeGotoPortResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 购买货物
func (c *Client) OnSkyTradeDealReq(f func(ctx Context, msg *pb.SkyTradeDealResp)error) {
	c.RegisterHandler(15005, func(ctx Context) error {
		msg := &pb.SkyTradeDealResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 挑战列表
func (c *Client) OnSkyTradeChallengeListReq(f func(ctx Context, msg *pb.SkyTradeChallengeListResp)error) {
	c.RegisterHandler(15006, func(ctx Context) error {
		msg := &pb.SkyTradeChallengeListResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求挑战
func (c *Client) OnSkyTradeChallengeReq(f func(ctx Context, msg *pb.SkyTradeChallengeResp)error) {
	c.RegisterHandler(15007, func(ctx Context) error {
		msg := &pb.SkyTradeChallengeResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟成员跑商威望信息
func (c *Client) OnSkyTradeUnionFameReq(f func(ctx Context, msg *pb.SkyTradeUnionFameResp)error) {
	c.RegisterHandler(15008, func(ctx Context) error {
		msg := &pb.SkyTradeUnionFameResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取到港事件奖励
func (c *Client) OnSkyTradeGetRewardReq(f func(ctx Context, msg *pb.SkyTradeGetRewardResp)error) {
	c.RegisterHandler(15009, func(ctx Context) error {
		msg := &pb.SkyTradeGetRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求战报
func (c *Client) OnSkyTradeReportReq(f func(ctx Context, msg *pb.SkyTradeReportResp)error) {
	c.RegisterHandler(15010, func(ctx Context) error {
		msg := &pb.SkyTradeReportResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取日志
func (c *Client) OnSkyTradePortLogReq(f func(ctx Context, msg *pb.SkyTradeLogResp)error) {
	c.RegisterHandler(15011, func(ctx Context) error {
		msg := &pb.SkyTradeLogResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 天地灵石信息
func (c *Client) OnSkyTradeSparInfoReq(f func(ctx Context, msg *pb.SkyTradeSparInfoResp)error) {
	c.RegisterHandler(15012, func(ctx Context) error {
		msg := &pb.SkyTradeSparInfoResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取天地灵石等级奖励
func (c *Client) OnSkyTradeGetSparPowerReq(f func(ctx Context, msg *pb.SkyTradeGetSparPowerResp)error) {
	c.RegisterHandler(15013, func(ctx Context) error {
		msg := &pb.SkyTradeGetSparPowerResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取天地灵石盟友奖励
func (c *Client) OnSkyTradeGetWelfareReq(f func(ctx Context, msg *pb.SkyTradeGetWelfareResp)error) {
	c.RegisterHandler(15014, func(ctx Context) error {
		msg := &pb.SkyTradeGetWelfareResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 发送灵液
func (c *Client) OnSkyTradeSendWelfareReq(f func(ctx Context, msg *pb.SkyTradeSendWelfareResp)error) {
	c.RegisterHandler(15015, func(ctx Context) error {
		msg := &pb.SkyTradeSendWelfareResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 灵液记录
func (c *Client) OnSkyTradeWelfareRecordReq(f func(ctx Context, msg *pb.SkyTradeWelfareRecordResp)error) {
	c.RegisterHandler(15016, func(ctx Context) error {
		msg := &pb.SkyTradeWelfareRecordResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 挑战界面心跳
func (c *Client) OnSkyTradeChallengeHeartBeatReq(f func(ctx Context, msg *pb.SkyTradeChallengeHeartBeatResp)error) {
	c.RegisterHandler(15017, func(ctx Context) error {
		msg := &pb.SkyTradeChallengeHeartBeatResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 使用道具补充挑战次数
func (c *Client) OnSkyTradeUseRobItemReq(f func(ctx Context, msg *pb.SkyTradeUseRobItemResp)error) {
	c.RegisterHandler(15018, func(ctx Context) error {
		msg := &pb.SkyTradeUseRobItemResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求详细战报信息
func (c *Client) OnSkyTradeReportDetailReq(f func(ctx Context, msg *pb.SkyTradeReportDetailResp)error) {
	c.RegisterHandler(15019, func(ctx Context) error {
		msg := &pb.SkyTradeReportDetailResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 竞猜结果
func (c *Client) OnSkyTradeGuessPlayersReq(f func(ctx Context, msg *pb.SkyTradeGuessPlayersResp)error) {
	c.RegisterHandler(15020, func(ctx Context) error {
		msg := &pb.SkyTradeGuessPlayersResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求竞猜界面数据
func (c *Client) OnSkyTradeGuessDataReq(f func(ctx Context, msg *pb.SkyTradeGuessDataResp)error) {
	c.RegisterHandler(15021, func(ctx Context) error {
		msg := &pb.SkyTradeGuessDataResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 设置竞猜选择
func (c *Client) OnSkyTradeGuessSelectReq(f func(ctx Context, msg *pb.SkyTradeGuessSelectResp)error) {
	c.RegisterHandler(15022, func(ctx Context) error {
		msg := &pb.SkyTradeGuessSelectResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取竞猜奖励
func (c *Client) OnSkyTradeGuessRewardReq(f func(ctx Context, msg *pb.SkyTradeGuessRewardResp)error) {
	c.RegisterHandler(15023, func(ctx Context) error {
		msg := &pb.SkyTradeGuessRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求上报所在地点稀缺货物
func (c *Client) OnSkyTradeReportGoodsReq(f func(ctx Context, msg *pb.SkyTradeReportGoodsResp)error) {
	c.RegisterHandler(15024, func(ctx Context) error {
		msg := &pb.SkyTradeReportGoodsResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求刷新货物库存
func (c *Client) OnSkyTradeResetStockReq(f func(ctx Context, msg *pb.SkyTradeResetStockResp)error) {
	c.RegisterHandler(15025, func(ctx Context) error {
		msg := &pb.SkyTradeResetStockResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 红点详情
func (c *Client) OnSkyTradeRedPointReq(f func(ctx Context, msg *pb.SkyTradeRedPointResp)error) {
	c.RegisterHandler(15026, func(ctx Context) error {
		msg := &pb.SkyTradeRedPointResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟分段为伤害信息
func (c *Client) OnSkyTradeUnionGroupDamageReq(f func(ctx Context, msg *pb.SkyTradeUnionGroupDamageResp)error) {
	c.RegisterHandler(15027, func(ctx Context) error {
		msg := &pb.SkyTradeUnionGroupDamageResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟分数积分变更
func (c *Client) OnSkyTradeUnionDataSync(f func(ctx Context, msg *pb.SkyTradeUnionDataSync)error) {
	c.RegisterHandler(15050, func(ctx Context) error {
		msg := &pb.SkyTradeUnionDataSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 货物刷新
func (c *Client) OnSkyTradeResetSync(f func(ctx Context, msg *pb.SkyTradeResetSync)error) {
	c.RegisterHandler(15051, func(ctx Context) error {
		msg := &pb.SkyTradeResetSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步地图盟友的航行信息
func (c *Client) OnSkyTradeAirshipInfoSync(f func(ctx Context, msg *pb.SkyTradeAirshipInfoSync)error) {
	c.RegisterHandler(15052, func(ctx Context) error {
		msg := &pb.SkyTradeAirshipInfoSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步挑战信息
func (c *Client) OnSkyTradeChallengeInfoSync(f func(ctx Context, msg *pb.SkyTradeChallengeInfoSync)error) {
	c.RegisterHandler(15053, func(ctx Context) error {
		msg := &pb.SkyTradeChallengeInfoSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步紧缺货物信息
func (c *Client) OnSkyTradeRareGoodsSync(f func(ctx Context, msg *pb.SkyTradeRareGoodsSync)error) {
	c.RegisterHandler(15054, func(ctx Context) error {
		msg := &pb.SkyTradeRareGoodsSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步威望排名
func (c *Client) OnSkyTradeFameRankSync(f func(ctx Context, msg *pb.SkyTradeFameRankSync)error) {
	c.RegisterHandler(15055, func(ctx Context) error {
		msg := &pb.SkyTradeFameRankSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 支付宝广告礼包-触发
func (c *Client) OnADGiftTriggerReq(f func(ctx Context, msg *pb.ADGiftTriggerResp)error) {
	c.RegisterHandler(15201, func(ctx Context) error {
		msg := &pb.ADGiftTriggerResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 支付宝广告礼包-领取奖励
func (c *Client) OnADGiftGetRewardReq(f func(ctx Context, msg *pb.ADGiftGetRewardResp)error) {
	c.RegisterHandler(15202, func(ctx Context) error {
		msg := &pb.ADGiftGetRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 家园-进入
func (c *Client) OnYardEnterReq(f func(ctx Context, msg *pb.YardEnterResp)error) {
	c.RegisterHandler(15801, func(ctx Context) error {
		msg := &pb.YardEnterResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 家园-战斗解锁区域
func (c *Client) OnYardBattleReq(f func(ctx Context, msg *pb.YardBattleResp)error) {
	c.RegisterHandler(15802, func(ctx Context) error {
		msg := &pb.YardBattleResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 家园-建筑摆放
func (c *Client) OnYardPosActionReq(f func(ctx Context, msg *pb.YardPosActionResp)error) {
	c.RegisterHandler(15803, func(ctx Context) error {
		msg := &pb.YardPosActionResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 家园升级
func (c *Client) OnYardLevelUpReq(f func(ctx Context, msg *pb.YardLevelUpResp)error) {
	c.RegisterHandler(15804, func(ctx Context) error {
		msg := &pb.YardLevelUpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 家园-清理杂物
func (c *Client) OnYardClearReq(f func(ctx Context, msg *pb.YardClearResp)error) {
	c.RegisterHandler(15805, func(ctx Context) error {
		msg := &pb.YardClearResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 家园请求
func (c *Client) OnYardInviteReq(f func(ctx Context, msg *pb.YardInviteResp)error) {
	c.RegisterHandler(15806, func(ctx Context) error {
		msg := &pb.YardInviteResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入兑换商店
func (c *Client) OnYardEnterShopReq(f func(ctx Context, msg *pb.YardEnterShopResp)error) {
	c.RegisterHandler(15807, func(ctx Context) error {
		msg := &pb.YardEnterShopResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 刷新兑换商店
func (c *Client) OnYardRefreshShopReq(f func(ctx Context, msg *pb.YardRefreshShopResp)error) {
	c.RegisterHandler(15808, func(ctx Context) error {
		msg := &pb.YardRefreshShopResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 羁绊共鸣升级
func (c *Client) OnYardSkillCombineLvUpReq(f func(ctx Context, msg *pb.YardSkillCombineLvUpResp)error) {
	c.RegisterHandler(15821, func(ctx Context) error {
		msg := &pb.YardSkillCombineLvUpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 抽取建筑
func (c *Client) OnYardDrawReq(f func(ctx Context, msg *pb.YardDrawResp)error) {
	c.RegisterHandler(15822, func(ctx Context) error {
		msg := &pb.YardDrawResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 建筑升级
func (c *Client) OnBuildUpLevelReq(f func(ctx Context, msg *pb.BuildUpLevelResp)error) {
	c.RegisterHandler(15823, func(ctx Context) error {
		msg := &pb.BuildUpLevelResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 装饰建筑解锁
func (c *Client) OnBuildUnlockReq(f func(ctx Context, msg *pb.BuildUnlockResp)error) {
	c.RegisterHandler(15824, func(ctx Context) error {
		msg := &pb.BuildUnlockResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 建筑生产
func (c *Client) OnYardBuildMakeReq(f func(ctx Context, msg *pb.YardBuildMakeResp)error) {
	c.RegisterHandler(15825, func(ctx Context) error {
		msg := &pb.YardBuildMakeResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 桃子协助
func (c *Client) OnYardBuildHelpReq(f func(ctx Context, msg *pb.YardBuildhHelpResp)error) {
	c.RegisterHandler(15826, func(ctx Context) error {
		msg := &pb.YardBuildhHelpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 建筑领奖
func (c *Client) OnYardBuildGainRewardReq(f func(ctx Context, msg *pb.YardBuildGainRewardResp)error) {
	c.RegisterHandler(15827, func(ctx Context) error {
		msg := &pb.YardBuildGainRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 建筑升级加速
func (c *Client) OnYardBuildSpeedUpReq(f func(ctx Context, msg *pb.YardBuildhSpeedUpResp)error) {
	c.RegisterHandler(15828, func(ctx Context) error {
		msg := &pb.YardBuildhSpeedUpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟协助列表请求
func (c *Client) OnGetYardUnionHelpDataListReq(f func(ctx Context, msg *pb.GetYardUnionHelpDataListResp)error) {
	c.RegisterHandler(15829, func(ctx Context) error {
		msg := &pb.GetYardUnionHelpDataListResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求妖盟协助
func (c *Client) OnRequestYardUnionHelpReq(f func(ctx Context, msg *pb.RequestYardUnionHelpResp)error) {
	c.RegisterHandler(15830, func(ctx Context) error {
		msg := &pb.RequestYardUnionHelpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 协助他人
func (c *Client) OnYardUnionHelpReq(f func(ctx Context, msg *pb.YardUnionHelpResp)error) {
	c.RegisterHandler(15831, func(ctx Context) error {
		msg := &pb.YardUnionHelpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 家园装饰建筑一键升级
func (c *Client) OnYardOneKeyLevelUpReq(f func(ctx Context, msg *pb.YardOneKeyLevelUpResp)error) {
	c.RegisterHandler(15833, func(ctx Context) error {
		msg := &pb.YardOneKeyLevelUpResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 桃树协助列表
func (c *Client) OnYardPeachRecordListReq(f func(ctx Context, msg *pb.YardPeachRecordListResp)error) {
	c.RegisterHandler(15834, func(ctx Context) error {
		msg := &pb.YardPeachRecordListResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 产物商人信息
func (c *Client) OnYardMerchantInfoReq(f func(ctx Context, msg *pb.YardMerchantInfoResp)error) {
	c.RegisterHandler(15835, func(ctx Context) error {
		msg := &pb.YardMerchantInfoResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 产物出售
func (c *Client) OnYardMerchantExchangeReq(f func(ctx Context, msg *pb.YardMerchantExchangeResp)error) {
	c.RegisterHandler(15836, func(ctx Context) error {
		msg := &pb.YardMerchantExchangeResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 留言板开关
func (c *Client) OnYardChatSwitchReq(f func(ctx Context, msg *pb.YardChatSwitchResp)error) {
	c.RegisterHandler(15837, func(ctx Context) error {
		msg := &pb.YardChatSwitchResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 留言操作
func (c *Client) OnYardChatActionReq(f func(ctx Context, msg *pb.YardChatActionResp)error) {
	c.RegisterHandler(15839, func(ctx Context) error {
		msg := &pb.YardChatActionResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 留言
func (c *Client) OnYardChatRecordReq(f func(ctx Context, msg *pb.YardChatRecordResp)error) {
	c.RegisterHandler(15840, func(ctx Context) error {
		msg := &pb.YardChatRecordResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 留言板信息
func (c *Client) OnYardChatRecordListReq(f func(ctx Context, msg *pb.YardChatRecordListResp)error) {
	c.RegisterHandler(15841, func(ctx Context) error {
		msg := &pb.YardChatRecordListResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 点赞
func (c *Client) OnYardGiveLikeReq(f func(ctx Context, msg *pb.YardGiveLikeResp)error) {
	c.RegisterHandler(15842, func(ctx Context) error {
		msg := &pb.YardGiveLikeResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 家园登录同步
func (c *Client) OnYardLoginSync(f func(ctx Context, msg *pb.YardLoginSync)error) {
	c.RegisterHandler(15843, func(ctx Context) error {
		msg := &pb.YardLoginSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 装饰建筑等级同步
func (c *Client) OnYardBuildInfoMsg(f func(ctx Context, msg *pb.YardBuildInfoMsg)error) {
	c.RegisterHandler(15844, func(ctx Context) error {
		msg := &pb.YardBuildInfoMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 桃树协助信息同步
func (c *Client) OnYardPeachHelpDataSync(f func(ctx Context, msg *pb.YardPeachHelpDataSync)error) {
	c.RegisterHandler(15845, func(ctx Context) error {
		msg := &pb.YardPeachHelpDataSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 家园基础数据同步
func (c *Client) OnYardBaseMsgSync(f func(ctx Context, msg *pb.YardBaseMsgSync)error) {
	c.RegisterHandler(15847, func(ctx Context) error {
		msg := &pb.YardBaseMsgSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 家园建筑信息同步
func (c *Client) OnYardBuildInfoMsgSync(f func(ctx Context, msg *pb.YardBuildInfoMsgSync)error) {
	c.RegisterHandler(15848, func(ctx Context) error {
		msg := &pb.YardBuildInfoMsgSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 家园生产信息同步
func (c *Client) OnYardMakeMsgSync(f func(ctx Context, msg *pb.YardMakeMsgSync)error) {
	c.RegisterHandler(15849, func(ctx Context) error {
		msg := &pb.YardMakeMsgSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 家园每日数据同步
func (c *Client) OnYardRefreshDataSync(f func(ctx Context, msg *pb.YardRefreshDataSync)error) {
	c.RegisterHandler(15850, func(ctx Context) error {
		msg := &pb.YardRefreshDataSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 家园功能建筑升级同步
func (c *Client) OnYardBuildUpSync(f func(ctx Context, msg *pb.YardBuildUpSync)error) {
	c.RegisterHandler(15851, func(ctx Context) error {
		msg := &pb.YardBuildUpSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 家园举报
func (c *Client) OnYardReportReq(f func(ctx Context, msg *pb.YardReportResp)error) {
	c.RegisterHandler(15852, func(ctx Context) error {
		msg := &pb.YardReportResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 申请、撤销
func (c *Client) OnShuraBattleApplyReq(f func(ctx Context, msg *pb.ShuraBattleApplyResp)error) {
	c.RegisterHandler(15910, func(ctx Context) error {
		msg := &pb.ShuraBattleApplyResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入三界征途界面
func (c *Client) OnEnterPlanesTrialReq(f func(ctx Context, msg *pb.EnterPlanesTrialRsp)error) {
	c.RegisterHandler(16001, func(ctx Context) error {
		msg := &pb.EnterPlanesTrialRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 组队内部请求
func (c *Client) OnEnterPlanesTrialTeamReq(f func(ctx Context, msg *pb.EnterPlanesTrialTeamRsp)error) {
	c.RegisterHandler(16002, func(ctx Context) error {
		msg := &pb.EnterPlanesTrialTeamRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 组队页面请求
func (c *Client) OnPlanesTrialTeamStartReq(f func(ctx Context, msg *pb.PlanesTrialTeamStartRsp)error) {
	c.RegisterHandler(16003, func(ctx Context) error {
		msg := &pb.PlanesTrialTeamStartRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 创建队伍
func (c *Client) OnPlanesTrialCreateTeamReq(f func(ctx Context, msg *pb.PlanesTrialCreateTeamRsp)error) {
	c.RegisterHandler(16004, func(ctx Context) error {
		msg := &pb.PlanesTrialCreateTeamRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取队伍列表
func (c *Client) OnPlanesTrialGetTeamListReq(f func(ctx Context, msg *pb.PlanesTrialGetTeamListRsp)error) {
	c.RegisterHandler(16005, func(ctx Context) error {
		msg := &pb.PlanesTrialGetTeamListRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查找队伍
func (c *Client) OnPlanesTrialGetTeamInfoReq(f func(ctx Context, msg *pb.PlanesTrialGetTeamInfoRsp)error) {
	c.RegisterHandler(16006, func(ctx Context) error {
		msg := &pb.PlanesTrialGetTeamInfoRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 三界征途-申请加入队伍
func (c *Client) OnPlanesTrialJoinTeamReq(f func(ctx Context, msg *pb.PlanesTrialJoinTeamRsp)error) {
	c.RegisterHandler(16007, func(ctx Context) error {
		msg := &pb.PlanesTrialJoinTeamRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 取消加入队伍
func (c *Client) OnPlanesTrialCancelTeamApplyReq(f func(ctx Context, msg *pb.PlanesTrialCancelTeamApplyRsp)error) {
	c.RegisterHandler(16008, func(ctx Context) error {
		msg := &pb.PlanesTrialCancelTeamApplyRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同意申请
func (c *Client) OnPlanesTrialApplyJoinTeamAgreeReq(f func(ctx Context, msg *pb.PlanesTrialApplyJoinTeamAgreeRsp)error) {
	c.RegisterHandler(16009, func(ctx Context) error {
		msg := &pb.PlanesTrialApplyJoinTeamAgreeRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 一键拒绝
func (c *Client) OnPlanesTrialApplyJoinTeamRefusedReq(f func(ctx Context, msg *pb.PlanesTrialApplyJoinTeamRefusedRsp)error) {
	c.RegisterHandler(16010, func(ctx Context) error {
		msg := &pb.PlanesTrialApplyJoinTeamRefusedRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 退出队伍
func (c *Client) OnPlanesTrialQuitTeamReq(f func(ctx Context, msg *pb.PlanesTrialQuitTeamRsp)error) {
	c.RegisterHandler(16011, func(ctx Context) error {
		msg := &pb.PlanesTrialQuitTeamRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 踢出队伍
func (c *Client) OnPlanesTrialKickOutTeamReq(f func(ctx Context, msg *pb.PlanesTrialKickOutTeamRsp)error) {
	c.RegisterHandler(16012, func(ctx Context) error {
		msg := &pb.PlanesTrialKickOutTeamRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 转让队长
func (c *Client) OnPlanesTrialChangeLeaderReq(f func(ctx Context, msg *pb.PlanesTrialChangeLeaderRsp)error) {
	c.RegisterHandler(16013, func(ctx Context) error {
		msg := &pb.PlanesTrialChangeLeaderRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 开始准备
func (c *Client) OnPlanesTrialTeamPrepareReq(f func(ctx Context, msg *pb.PlanesTrialTeamPrepareRsp)error) {
	c.RegisterHandler(16014, func(ctx Context) error {
		msg := &pb.PlanesTrialTeamPrepareRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 匹配
func (c *Client) OnPlanesTrialMatchMemberReq(f func(ctx Context, msg *pb.PlanesTrialMatchMemberRsp)error) {
	c.RegisterHandler(16015, func(ctx Context) error {
		msg := &pb.PlanesTrialMatchMemberRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 开始战斗
func (c *Client) OnPlanesTrialStartBattleReq(f func(ctx Context, msg *pb.PlanesTrialStartBattleRsp)error) {
	c.RegisterHandler(16016, func(ctx Context) error {
		msg := &pb.PlanesTrialStartBattleRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 请求玩家数据
func (c *Client) OnPlanesTrialGetPlayerInfoReq(f func(ctx Context, msg *pb.PlanesTrialGetPlayerInfoRsp)error) {
	c.RegisterHandler(16017, func(ctx Context) error {
		msg := &pb.PlanesTrialGetPlayerInfoRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 战斗回放
func (c *Client) OnPlanesTrialGetBattleVideoReq(f func(ctx Context, msg *pb.PlanesTrialGetBattleVideoRsp)error) {
	c.RegisterHandler(16019, func(ctx Context) error {
		msg := &pb.PlanesTrialGetBattleVideoRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取boss信息
func (c *Client) OnPlanesTrialGetBossInfoReq(f func(ctx Context, msg *pb.PlanesTrialGetBossInfoRsp)error) {
	c.RegisterHandler(16020, func(ctx Context) error {
		msg := &pb.PlanesTrialGetBossInfoRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 三界征途-获取挑战次数
func (c *Client) OnPlanesTrialChallengeTimeReq(f func(ctx Context, msg *pb.PlanesTrialChallengeTimeRsp)error) {
	c.RegisterHandler(16021, func(ctx Context) error {
		msg := &pb.PlanesTrialChallengeTimeRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 开始匹配
func (c *Client) OnPlanesTrialStartMatchReq(f func(ctx Context, msg *pb.PlanesTrialStartMatchRsp)error) {
	c.RegisterHandler(16022, func(ctx Context) error {
		msg := &pb.PlanesTrialStartMatchRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 邀请
func (c *Client) OnPlanesTrialInviteReq(f func(ctx Context, msg *pb.PlanesTrialInviteRsp)error) {
	c.RegisterHandler(16023, func(ctx Context) error {
		msg := &pb.PlanesTrialInviteRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 离开系统
func (c *Client) OnLeavePlanesTrialReq(f func(ctx Context, msg *pb.LeavePlanesTrialRsp)error) {
	c.RegisterHandler(16024, func(ctx Context) error {
		msg := &pb.LeavePlanesTrialRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取boss妖力
func (c *Client) OnPlanesTrialGetBossPowerReq(f func(ctx Context, msg *pb.PlanesTrialGetBossPowerRsp)error) {
	c.RegisterHandler(16025, func(ctx Context) error {
		msg := &pb.PlanesTrialGetBossPowerRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 队长跳过战斗
func (c *Client) OnPlanesTrialSkipBattleReq(f func(ctx Context, msg *pb.PlanesTrialSkipBattleRsp)error) {
	c.RegisterHandler(16026, func(ctx Context) error {
		msg := &pb.PlanesTrialSkipBattleRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 队长开始选buff
func (c *Client) OnPlanesTrialStartSelectBuffReq(f func(ctx Context, msg *pb.PlanesTrialStartSelectBuffRsp)error) {
	c.RegisterHandler(16027, func(ctx Context) error {
		msg := &pb.PlanesTrialStartSelectBuffRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 选buff
func (c *Client) OnPlanesTrialSelectBuffReq(f func(ctx Context, msg *pb.PlanesTrialSelectBuffRsp)error) {
	c.RegisterHandler(16028, func(ctx Context) error {
		msg := &pb.PlanesTrialSelectBuffRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 保存选择选项
func (c *Client) OnPlanesTrialSetBuffPreferenceReq(f func(ctx Context, msg *pb.PlanesTrialSetBuffPreferenceRsp)error) {
	c.RegisterHandler(16029, func(ctx Context) error {
		msg := &pb.PlanesTrialSetBuffPreferenceRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看自己选的buff
func (c *Client) OnPlanesTrialGetSelectedBuffReq(f func(ctx Context, msg *pb.PlanesTrialGetSelectedBuffRsp)error) {
	c.RegisterHandler(16030, func(ctx Context) error {
		msg := &pb.PlanesTrialGetSelectedBuffRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取成就奖励
func (c *Client) OnPlanesTrialGetAchievementRewardReq(f func(ctx Context, msg *pb.PlanesTrialGetAchievementRewardRsp)error) {
	c.RegisterHandler(16031, func(ctx Context) error {
		msg := &pb.PlanesTrialGetAchievementRewardRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看预设buff
func (c *Client) OnPlanesTrialGetBuffPreferenceReq(f func(ctx Context, msg *pb.PlanesTrialGetBuffPreferenceRsp)error) {
	c.RegisterHandler(16033, func(ctx Context) error {
		msg := &pb.PlanesTrialGetBuffPreferenceRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取分身数据
func (c *Client) OnPlanesTrialGetGodBodyDataReq(f func(ctx Context, msg *pb.PlanesTrialGetGodBodyDataResp)error) {
	c.RegisterHandler(16034, func(ctx Context) error {
		msg := &pb.PlanesTrialGetGodBodyDataResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取大奖信息
func (c *Client) OnPlanesTrialGetGrandPrizeInfoReq(f func(ctx Context, msg *pb.PlanesTrialGetGrandPrizeInfoResp)error) {
	c.RegisterHandler(16035, func(ctx Context) error {
		msg := &pb.PlanesTrialGetGrandPrizeInfoResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 领取大奖
func (c *Client) OnPlanesTrialReceiveGrandPrizeReq(f func(ctx Context, msg *pb.PlanesTrialReceiveGrandPrizeResp)error) {
	c.RegisterHandler(16036, func(ctx Context) error {
		msg := &pb.PlanesTrialReceiveGrandPrizeResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 队长接受到广播
func (c *Client) OnPlanesTrialTeamLeaderNotify(f func(ctx Context, msg *pb.PlanesTrialTeamLeaderNotify)error) {
	c.RegisterHandler(16040, func(ctx Context) error {
		msg := &pb.PlanesTrialTeamLeaderNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 队员接收到的广播
func (c *Client) OnPlanesTrialTeamMemberNotify(f func(ctx Context, msg *pb.PlanesTrialTeamMemberNotify)error) {
	c.RegisterHandler(16041, func(ctx Context) error {
		msg := &pb.PlanesTrialTeamMemberNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入布阵
func (c *Client) OnPlanesTrialEnterBattleReq(f func(ctx Context, msg *pb.PlanesTrialEnterBattleResp)error) {
	c.RegisterHandler(16053, func(ctx Context) error {
		msg := &pb.PlanesTrialEnterBattleResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进入切换分身界面
func (c *Client) OnPlanesTrialEnterSwitchSeparationReq(f func(ctx Context, msg *pb.PlanesTrialEnterSwitchSeparationRsp)error) {
	c.RegisterHandler(16054, func(ctx Context) error {
		msg := &pb.PlanesTrialEnterSwitchSeparationRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 切换上阵的分身
func (c *Client) OnPlanesTrialSwitchSeparationReq(f func(ctx Context, msg *pb.PlanesTrialSwitchSeparationRsp)error) {
	c.RegisterHandler(16055, func(ctx Context) error {
		msg := &pb.PlanesTrialSwitchSeparationRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看上阵分身详细数据
func (c *Client) OnPlanesTrialSeparationDetailReq(f func(ctx Context, msg *pb.PlanesTrialSeparationDetailRsp)error) {
	c.RegisterHandler(16056, func(ctx Context) error {
		msg := &pb.PlanesTrialSeparationDetailRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 变更参战玩家位置
func (c *Client) OnPlanesTrialChangePosReq(f func(ctx Context, msg *pb.PlanesTrialChangePosRsp)error) {
	c.RegisterHandler(16057, func(ctx Context) error {
		msg := &pb.PlanesTrialChangePosRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 进行战斗
func (c *Client) OnPlanesTrialDoBattleReq(f func(ctx Context, msg *pb.PlanesTrialDoBattleRsp)error) {
	c.RegisterHandler(16058, func(ctx Context) error {
		msg := &pb.PlanesTrialDoBattleRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 通知进入战斗
func (c *Client) OnPlanesTrialDoBattleNotify(f func(ctx Context, msg *pb.PlanesTrialDoBattleNotify)error) {
	c.RegisterHandler(16065, func(ctx Context) error {
		msg := &pb.PlanesTrialDoBattleNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 广播进入布阵界面
func (c *Client) OnPlanesTrialEnterBattleNotify(f func(ctx Context, msg *pb.PlanesTrialEnterBattleNotify)error) {
	c.RegisterHandler(16070, func(ctx Context) error {
		msg := &pb.PlanesTrialEnterBattleNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 通知上阵的分身变更
func (c *Client) OnPlanesTrialSwitchSeparationNotify(f func(ctx Context, msg *pb.PlanesTrialSwitchSeparationNotify)error) {
	c.RegisterHandler(16071, func(ctx Context) error {
		msg := &pb.PlanesTrialSwitchSeparationNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 更新绑定数据
func (c *Client) OnPlanesTrialUpdateLockDataReq(f func(ctx Context, msg *pb.PlanesTrialUpdateLockDataResp)error) {
	c.RegisterHandler(16072, func(ctx Context) error {
		msg := &pb.PlanesTrialUpdateLockDataResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 通知参战玩家位置变更
func (c *Client) OnPlanesTrialChangePosNotify(f func(ctx Context, msg *pb.PlanesTrialChangePosNotify)error) {
	c.RegisterHandler(16074, func(ctx Context) error {
		msg := &pb.PlanesTrialChangePosNotify{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取被压制属性
func (c *Client) OnPlayerRestrainInfoMsgReq(f func(ctx Context, msg *pb.PlayerRestrainInfoMsgResp)error) {
	c.RegisterHandler(16088, func(ctx Context) error {
		msg := &pb.PlayerRestrainInfoMsgResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取秘境人数请求
func (c *Client) OnGetPlanesTrialTrialMemberCountReq(f func(ctx Context, msg *pb.GetPlanesTrialTrialMemberCountResp)error) {
	c.RegisterHandler(16089, func(ctx Context) error {
		msg := &pb.GetPlanesTrialTrialMemberCountResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取排行榜
func (c *Client) OnPlanesTrialRankGetReq(f func(ctx Context, msg *pb.PlanesTrialRankGetResp)error) {
	c.RegisterHandler(16090, func(ctx Context) error {
		msg := &pb.PlanesTrialRankGetResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 心跳包
func (c *Client) OnPlanesTrialHeartbeatReq(f func(ctx Context, msg *pb.PlanesTrialHeartbeatResp)error) {
	c.RegisterHandler(16091, func(ctx Context) error {
		msg := &pb.PlanesTrialHeartbeatResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 获取自选奖励信息
func (c *Client) OnPlanesTrialGetSelectRewardDetailReq(f func(ctx Context, msg *pb.PlanesTrialGetSelectRewardDetailResp)error) {
	c.RegisterHandler(16092, func(ctx Context) error {
		msg := &pb.PlanesTrialGetSelectRewardDetailResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 选择奖励
func (c *Client) OnPlanesTrialSelectRewardReq(f func(ctx Context, msg *pb.PlanesTrialSelectRewardResp)error) {
	c.RegisterHandler(16093, func(ctx Context) error {
		msg := &pb.PlanesTrialSelectRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 查看关卡录像数据
func (c *Client) OnPlanesTrialVideoInfoReq(f func(ctx Context, msg *pb.PlanesTrialVideoInfoRsp)error) {
	c.RegisterHandler(16094, func(ctx Context) error {
		msg := &pb.PlanesTrialVideoInfoRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 播放关卡录像
func (c *Client) OnPlanesTrialPlayVideoReq(f func(ctx Context, msg *pb.PlanesTrialPlayVideoRsp)error) {
	c.RegisterHandler(16095, func(ctx Context) error {
		msg := &pb.PlanesTrialPlayVideoRsp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 播放关卡录像
func (c *Client) OnPlanesTrialCloseSettleScreenReq(f func(ctx Context, msg *pb.PlanesTrialCloseSettleScreenResp)error) {
	c.RegisterHandler(16096, func(ctx Context) error {
		msg := &pb.PlanesTrialCloseSettleScreenResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟寻宝地图互动
func (c *Client) OnUnionTreasureDrawChipReq(f func(ctx Context, msg *pb.UnionTreasureDrawChipResp)error) {
	c.RegisterHandler(16201, func(ctx Context) error {
		msg := &pb.UnionTreasureDrawChipResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟寻宝索要
func (c *Client) OnUnionTreasureAskForReq(f func(ctx Context, msg *pb.UnionTreasureAskForResp)error) {
	c.RegisterHandler(16202, func(ctx Context) error {
		msg := &pb.UnionTreasureAskForResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟寻宝碎片赠送
func (c *Client) OnUnionTreasureSendChipReq(f func(ctx Context, msg *pb.UnionTreasureSendChipResp)error) {
	c.RegisterHandler(16203, func(ctx Context) error {
		msg := &pb.UnionTreasureSendChipResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟寻宝遗迹互动
func (c *Client) OnUnionTreasureBattleRelicReq(f func(ctx Context, msg *pb.UnionTreasureBattleRelicResp)error) {
	c.RegisterHandler(16205, func(ctx Context) error {
		msg := &pb.UnionTreasureBattleRelicResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟寻宝遗迹挖宝
func (c *Client) OnUnionTreasureHuntTreasureReq(f func(ctx Context, msg *pb.UnionTreasureHuntTreasureResp)error) {
	c.RegisterHandler(16206, func(ctx Context) error {
		msg := &pb.UnionTreasureHuntTreasureResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟寻宝进入界面
func (c *Client) OnUnionTreasureEnterReq(f func(ctx Context, msg *pb.UnionTreasureEnterResp)error) {
	c.RegisterHandler(16207, func(ctx Context) error {
		msg := &pb.UnionTreasureEnterResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟寻宝主要数据同步
func (c *Client) OnUnionTreasureCollectMsg(f func(ctx Context, msg *pb.UnionTreasureCollectMsg)error) {
	c.RegisterHandler(16208, func(ctx Context) error {
		msg := &pb.UnionTreasureCollectMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟寻宝获得碎片
func (c *Client) OnUnionTreasureGetGiveChipReq(f func(ctx Context, msg *pb.UnionTreasureGetGiveChipResp)error) {
	c.RegisterHandler(16209, func(ctx Context) error {
		msg := &pb.UnionTreasureGetGiveChipResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟寻宝遗迹数据请求
func (c *Client) OnUnionTreasureAskForRelicReq(f func(ctx Context, msg *pb.UnionTreasureAskForRelicResp)error) {
	c.RegisterHandler(16213, func(ctx Context) error {
		msg := &pb.UnionTreasureAskForRelicResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟寻宝分享协助
func (c *Client) OnUnionTreasureShareRelicReq(f func(ctx Context, msg *pb.UnionTreasureShareRelicResp)error) {
	c.RegisterHandler(16216, func(ctx Context) error {
		msg := &pb.UnionTreasureShareRelicResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟寻宝遗迹激活碎片
func (c *Client) OnUnionTreasureFillChipReq(f func(ctx Context, msg *pb.UnionTreasureFillChipResp)error) {
	c.RegisterHandler(16218, func(ctx Context) error {
		msg := &pb.UnionTreasureFillChipResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟寻宝时间同步
func (c *Client) OnUnionTreasuretLoginSync(f func(ctx Context, msg *pb.UnionTreasuretLoginSync)error) {
	c.RegisterHandler(16220, func(ctx Context) error {
		msg := &pb.UnionTreasuretLoginSync{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟寻宝遗迹消息过期
func (c *Client) OnCheckMessageValidReq(f func(ctx Context, msg *pb.CheckMessageValidResp)error) {
	c.RegisterHandler(16222, func(ctx Context) error {
		msg := &pb.CheckMessageValidResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟寻宝战斗记录
func (c *Client) OnUnionTreasureGetBattleRecordReq(f func(ctx Context, msg *pb.UnionTreasureGetBattleRecordResp)error) {
	c.RegisterHandler(16223, func(ctx Context) error {
		msg := &pb.UnionTreasureGetBattleRecordResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟寻宝赠送请求
func (c *Client) OnUnionTreasureGetCollectGiveMsgReq(f func(ctx Context, msg *pb.UnionTreasureGetCollectGiveMsgResp)error) {
	c.RegisterHandler(16224, func(ctx Context) error {
		msg := &pb.UnionTreasureGetCollectGiveMsgResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟寻宝退出请求
func (c *Client) OnUnionTreasureExitReq(f func(ctx Context, msg *pb.UnionTreasureExitResp)error) {
	c.RegisterHandler(16225, func(ctx Context) error {
		msg := &pb.UnionTreasureExitResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 妖盟寻宝升级弹窗
func (c *Client) OnUnionTreasureLvUpWindowsMsg(f func(ctx Context, msg *pb.UnionTreasureLvUpWindowsMsg)error) {
	c.RegisterHandler(16227, func(ctx Context) error {
		msg := &pb.UnionTreasureLvUpWindowsMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步美团订阅数据
func (c *Client) OnSyncMeiTuanSubscribeMsg(f func(ctx Context, msg *pb.SyncMeiTuanSubscribeMsg)error) {
	c.RegisterHandler(16700, func(ctx Context) error {
		msg := &pb.SyncMeiTuanSubscribeMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 美团订阅成功
func (c *Client) OnMeiTuanSubscribeSucceedReq(f func(ctx Context, msg *pb.MeiTuanSubscribeSucceedResp)error) {
	c.RegisterHandler(16701, func(ctx Context) error {
		msg := &pb.MeiTuanSubscribeSucceedResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 美团订阅领取
func (c *Client) OnMeiTuanSubscribeRewardReq(f func(ctx Context, msg *pb.MeiTuanSubscribeRewardResp)error) {
	c.RegisterHandler(16702, func(ctx Context) error {
		msg := &pb.MeiTuanSubscribeRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步美团日常信息
func (c *Client) OnSyncMeiTuanDailyMsg(f func(ctx Context, msg *pb.SyncMeiTuanDailyMsg)error) {
	c.RegisterHandler(16800, func(ctx Context) error {
		msg := &pb.SyncMeiTuanDailyMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 美团二楼登录
func (c *Client) OnMeiTuanDailyLoginReq(f func(ctx Context, msg *pb.MeiTuanDailyLoginResp)error) {
	c.RegisterHandler(16801, func(ctx Context) error {
		msg := &pb.MeiTuanDailyLoginResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 美团登录领取
func (c *Client) OnMeiTuanDailyRewardReq(f func(ctx Context, msg *pb.MeiTuanDailyRewardResp)error) {
	c.RegisterHandler(16802, func(ctx Context) error {
		msg := &pb.MeiTuanDailyRewardResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 仙玉宝府进入请求
func (c *Client) OnYueBaoEnterReq(f func(ctx Context, msg *pb.YueBaoEnterResp)error) {
	c.RegisterHandler(17001, func(ctx Context) error {
		msg := &pb.YueBaoEnterResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 仙玉宝府存储
func (c *Client) OnYueBaoDepositReq(f func(ctx Context, msg *pb.YueBaoDepositResp)error) {
	c.RegisterHandler(17002, func(ctx Context) error {
		msg := &pb.YueBaoDepositResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 仙玉宝府互动请求
func (c *Client) OnYueBaoInteractReq(f func(ctx Context, msg *pb.YueBaoInteractResp)error) {
	c.RegisterHandler(17003, func(ctx Context) error {
		msg := &pb.YueBaoInteractResp{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 仙玉宝府信息同步
func (c *Client) OnRechargeItemData(f func(ctx Context, msg *pb.RechargeItemData)error) {
	c.RegisterHandler(17004, func(ctx Context) error {
		msg := &pb.RechargeItemData{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 集卡登录同步
func (c *Client) OnMemoryCollectLoginSyncData(f func(ctx Context, msg *pb.MemoryCollectLoginSyncData)error) {
	c.RegisterHandler(17413, func(ctx Context) error {
		msg := &pb.MemoryCollectLoginSyncData{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

// 同步直播状态
func (c *Client) OnLiveShowNotifyReqMsg(f func(ctx Context, msg *pb.LiveShowNotifyRespMsg)error) {
	c.RegisterHandler(17701, func(ctx Context) error {
		msg := &pb.LiveShowNotifyRespMsg{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}
