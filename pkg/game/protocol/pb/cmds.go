package pb

func (*ReqLoginMsg) MsgId() int32 {
	return 20001
}

func (*ReqLoginMsg) Comment() string {
	return "登陆消息"
}
		
func (*RspLoginMsg) MsgId() int32 {
	return 1
}

func (*RspLoginMsg) Comment() string {
	return "登陆消息"
}
		
func (*ReqPingMsg) MsgId() int32 {
	return 20003
}

func (*ReqPingMsg) Comment() string {
	return "心跳"
}
		
func (*RspPingMsg) MsgId() int32 {
	return 3
}

func (*RspPingMsg) Comment() string {
	return "心跳"
}
		
func (*PreChargeReq) MsgId() int32 {
	return 20004
}

func (*PreChargeReq) Comment() string {
	return "预购买"
}
		
func (*PreChargeResp) MsgId() int32 {
	return 4
}

func (*PreChargeResp) Comment() string {
	return "预购买"
}
		
func (*CreateRoleReq) MsgId() int32 {
	return 20006
}

func (*CreateRoleReq) Comment() string {
	return "创建角色请求"
}
		
func (*CreateRoleResp) MsgId() int32 {
	return 6
}

func (*CreateRoleResp) Comment() string {
	return "创建角色请求"
}
		
func (*ChangeNickNameReq) MsgId() int32 {
	return 20008
}

func (*ChangeNickNameReq) Comment() string {
	return "改名"
}
		
func (*ChangeNickNameResp) MsgId() int32 {
	return 8
}

func (*ChangeNickNameResp) Comment() string {
	return "改名"
}
		
func (*UseExchangeKeyReq) MsgId() int32 {
	return 20009
}

func (*UseExchangeKeyReq) Comment() string {
	return "兑换码"
}
		
func (*UseExchangeKeyResp) MsgId() int32 {
	return 9
}

func (*UseExchangeKeyResp) Comment() string {
	return "兑换码"
}
		
func (*WatchAdReq) MsgId() int32 {
	return 20019
}

func (*WatchAdReq) Comment() string {
	return "观看视频广告任务"
}
		
func (*WatchAdResp) MsgId() int32 {
	return 19
}

func (*WatchAdResp) Comment() string {
	return "观看视频广告任务"
}
		
func (*CutRopeGetRewardReq) MsgId() int32 {
	return 20020
}

func (*CutRopeGetRewardReq) Comment() string {
	return "熊口逃生领奖"
}
		
func (*CutRopeGetRewardRsp) MsgId() int32 {
	return 20
}

func (*CutRopeGetRewardRsp) Comment() string {
	return "熊口逃生领奖"
}
		
func (*GetPolicyInfoReq) MsgId() int32 {
	return 20031
}

func (*GetPolicyInfoReq) Comment() string {
	return "获取隐私注册协议"
}
		
func (*GetPolicyInfoRsp) MsgId() int32 {
	return 31
}

func (*GetPolicyInfoRsp) Comment() string {
	return "获取隐私注册协议"
}
		
func (*GetBindStatusReq) MsgId() int32 {
	return 20033
}

func (*GetBindStatusReq) Comment() string {
	return "获取绑定状态"
}
		
func (*GetBindStatusResp) MsgId() int32 {
	return 33
}

func (*GetBindStatusResp) Comment() string {
	return "获取绑定状态"
}
		
func (*GetBindCodeReq) MsgId() int32 {
	return 20034
}

func (*GetBindCodeReq) Comment() string {
	return "生成绑定码"
}
		
func (*GetBindCodeResp) MsgId() int32 {
	return 34
}

func (*GetBindCodeResp) Comment() string {
	return "生成绑定码"
}
		
func (*UntieBindCodeReq) MsgId() int32 {
	return 20035
}

func (*UntieBindCodeReq) Comment() string {
	return "解绑"
}
		
func (*UntieBindCodeResp) MsgId() int32 {
	return 35
}

func (*UntieBindCodeResp) Comment() string {
	return "解绑"
}
		
func (*GetBindCodeInfoReq) MsgId() int32 {
	return 20036
}

func (*GetBindCodeInfoReq) Comment() string {
	return "获取绑定信息"
}
		
func (*GetBindCodeInfoResp) MsgId() int32 {
	return 36
}

func (*GetBindCodeInfoResp) Comment() string {
	return "获取绑定信息"
}
		
func (*DeletePlayerReq) MsgId() int32 {
	return 20037
}

func (*DeletePlayerReq) Comment() string {
	return "注销账号"
}
		
func (*DeletePlayerResp) MsgId() int32 {
	return 37
}

func (*DeletePlayerResp) Comment() string {
	return "注销账号"
}
		
func (*GetSystemGroupInfoReq) MsgId() int32 {
	return 20040
}

func (*GetSystemGroupInfoReq) Comment() string {
	return "获取系统分组信息"
}
		
func (*GetSystemGroupInfoResp) MsgId() int32 {
	return 40
}

func (*GetSystemGroupInfoResp) Comment() string {
	return "获取系统分组信息"
}
		
func (*GetSysRewardReq) MsgId() int32 {
	return 20103
}

func (*GetSysRewardReq) Comment() string {
	return "请求领取系统解锁奖励"
}
		
func (*GetSysRewardResp) MsgId() int32 {
	return 103
}

func (*GetSysRewardResp) Comment() string {
	return "请求领取系统解锁奖励"
}
		
func (*PrivilegeCardReceiveRewardReq) MsgId() int32 {
	return 20105
}

func (*PrivilegeCardReceiveRewardReq) Comment() string {
	return "领取特权卡奖励"
}
		
func (*PrivilegeCardReceiveRewardRsp) MsgId() int32 {
	return 105
}

func (*PrivilegeCardReceiveRewardRsp) Comment() string {
	return "领取特权卡奖励"
}
		
func (*ChangePlayerCharaReq) MsgId() int32 {
	return 20107
}

func (*ChangePlayerCharaReq) Comment() string {
	return "切换形象请求"
}
		
func (*ChangePlayerCharaResp) MsgId() int32 {
	return 107
}

func (*ChangePlayerCharaResp) Comment() string {
	return "切换形象请求"
}
		
func (*WorldChatReqMsg) MsgId() int32 {
	return 20110
}

func (*WorldChatReqMsg) Comment() string {
	return "请求聊天"
}
		
func (*WorldChatRespMsg) MsgId() int32 {
	return 110
}

func (*WorldChatRespMsg) Comment() string {
	return "请求聊天"
}
		
func (*ActivityChatReqMsg) MsgId() int32 {
	return 20116
}

func (*ActivityChatReqMsg) Comment() string {
	return "请求活动消息列表"
}
		
func (*WorldMessageListMsg) MsgId() int32 {
	return 116
}

func (*WorldMessageListMsg) Comment() string {
	return "请求活动消息列表"
}
		
func (*GetTopPlayerInfoReq) MsgId() int32 {
	return 20118
}

func (*GetTopPlayerInfoReq) Comment() string {
	return "获取仙树后面飘的玩家信息"
}
		
func (*GetTopPlayerInfoResp) MsgId() int32 {
	return 118
}

func (*GetTopPlayerInfoResp) Comment() string {
	return "获取仙树后面飘的玩家信息"
}
		
func (*UseTitleReq) MsgId() int32 {
	return 20120
}

func (*UseTitleReq) Comment() string {
	return "穿戴称号"
}
		
func (*UseTitleRsp) MsgId() int32 {
	return 120
}

func (*UseTitleRsp) Comment() string {
	return "穿戴称号"
}
		
func (*UseMedalReq) MsgId() int32 {
	return 20123
}

func (*UseMedalReq) Comment() string {
	return "穿戴称号"
}
		
func (*UseMedalRsp) MsgId() int32 {
	return 123
}

func (*UseMedalRsp) Comment() string {
	return "穿戴称号"
}
		
func (*WorldChatBlockReqMsg) MsgId() int32 {
	return 20130
}

func (*WorldChatBlockReqMsg) Comment() string {
	return "举报玩家"
}
		
func (*WorldChatBlockRespMsg) MsgId() int32 {
	return 130
}

func (*WorldChatBlockRespMsg) Comment() string {
	return "举报玩家"
}
		
func (*SystemChatReqMsg) MsgId() int32 {
	return 20132
}

func (*SystemChatReqMsg) Comment() string {
	return "系统聊天列表，请求下发"
}
		
func (*ReportMessageReqMsg) MsgId() int32 {
	return 20133
}

func (*ReportMessageReqMsg) Comment() string {
	return "聊天举报"
}
		
func (*ReportMessageRespMsg) MsgId() int32 {
	return 133
}

func (*ReportMessageRespMsg) Comment() string {
	return "聊天举报"
}
		
func (*CommonSearchPlayerReq) MsgId() int32 {
	return 20136
}

func (*CommonSearchPlayerReq) Comment() string {
	return "通用搜索玩家"
}
		
func (*CommonSearchPlayerResp) MsgId() int32 {
	return 136
}

func (*CommonSearchPlayerResp) Comment() string {
	return "通用搜索玩家"
}
		
func (*RedPacketOpenReq) MsgId() int32 {
	return 20141
}

func (*RedPacketOpenReq) Comment() string {
	return "获取红包信息"
}
		
func (*RedPacketOpenResp) MsgId() int32 {
	return 141
}

func (*RedPacketOpenResp) Comment() string {
	return "获取红包信息"
}
		
func (*RedPacketDrawReq) MsgId() int32 {
	return 20142
}

func (*RedPacketDrawReq) Comment() string {
	return "领取红包奖励"
}
		
func (*RedPacketDrawResp) MsgId() int32 {
	return 142
}

func (*RedPacketDrawResp) Comment() string {
	return "领取红包奖励"
}
		
func (*CollectPlayerCharaReq) MsgId() int32 {
	return 20150
}

func (*CollectPlayerCharaReq) Comment() string {
	return "皮肤羁绊任务领取"
}
		
func (*CollectPlayerCharaResp) MsgId() int32 {
	return 150
}

func (*CollectPlayerCharaResp) Comment() string {
	return "皮肤羁绊任务领取"
}
		
func (*UpgradePlayerCharaReq) MsgId() int32 {
	return 20151
}

func (*UpgradePlayerCharaReq) Comment() string {
	return "皮肤等级升级"
}
		
func (*UpgradePlayerCharaResp) MsgId() int32 {
	return 151
}

func (*UpgradePlayerCharaResp) Comment() string {
	return "皮肤等级升级"
}
		
func (*ReqDealEquipmentMsg) MsgId() int32 {
	return 20202
}

func (*ReqDealEquipmentMsg) Comment() string {
	return "处理装备"
}
		
func (*RspDealEquipmentMsg) MsgId() int32 {
	return 202
}

func (*RspDealEquipmentMsg) Comment() string {
	return "处理装备"
}
		
func (*ReqDreamMsg) MsgId() int32 {
	return 20203
}

func (*ReqDreamMsg) Comment() string {
	return "做梦"
}
		
func (*RspDreamMsg) MsgId() int32 {
	return 203
}

func (*RspDreamMsg) Comment() string {
	return "做梦"
}
		
func (*ReqRealmsLeveUpMsg) MsgId() int32 {
	return 20204
}

func (*ReqRealmsLeveUpMsg) Comment() string {
	return "境界升级"
}
		
func (*RspRealmsLeveUpMsg) MsgId() int32 {
	return 204
}

func (*RspRealmsLeveUpMsg) Comment() string {
	return "境界升级"
}
		
func (*DreamLvUpReq) MsgId() int32 {
	return 20205
}

func (*DreamLvUpReq) Comment() string {
	return "梦境升级"
}
		
func (*DreamLvUpResp) MsgId() int32 {
	return 205
}

func (*DreamLvUpResp) Comment() string {
	return "梦境升级"
}
		
func (*DreamLvUpSpeedUpReq) MsgId() int32 {
	return 20206
}

func (*DreamLvUpSpeedUpReq) Comment() string {
	return "梦境升级加速"
}
		
func (*DreamLvUpSpeedUpResp) MsgId() int32 {
	return 206
}

func (*DreamLvUpSpeedUpResp) Comment() string {
	return "梦境升级加速"
}
		
func (*GetTribulationSuccessProReq) MsgId() int32 {
	return 20208
}

func (*GetTribulationSuccessProReq) Comment() string {
	return "获取渡劫成功率"
}
		
func (*GetTribulationSuccessProResp) MsgId() int32 {
	return 208
}

func (*GetTribulationSuccessProResp) Comment() string {
	return "获取渡劫成功率"
}
		
func (*GetUnDealEquipmentMsgReq) MsgId() int32 {
	return 20209
}

func (*GetUnDealEquipmentMsgReq) Comment() string {
	return "获取未处理装备数据"
}
		
func (*GetUnDealEquipmentMsgResp) MsgId() int32 {
	return 209
}

func (*GetUnDealEquipmentMsgResp) Comment() string {
	return "获取未处理装备数据"
}
		
func (*GetAdRewardReq) MsgId() int32 {
	return 20211
}

func (*GetAdRewardReq) Comment() string {
	return "广告领奖"
}
		
func (*GetAdRewardResp) MsgId() int32 {
	return 211
}

func (*GetAdRewardResp) Comment() string {
	return "广告领奖"
}
		
func (*SoaringReq) MsgId() int32 {
	return 20212
}

func (*SoaringReq) Comment() string {
	return "斩道也调用这个接口"
}
		
func (*SoaringResp) MsgId() int32 {
	return 212
}

func (*SoaringResp) Comment() string {
	return "斩道也调用这个接口"
}
		
func (*SetSeparationNameReq) MsgId() int32 {
	return 20213
}

func (*SetSeparationNameReq) Comment() string {
	return "设置分身名字"
}
		
func (*SetSeparationNameResp) MsgId() int32 {
	return 213
}

func (*SetSeparationNameResp) Comment() string {
	return "设置分身名字"
}
		
func (*SwitchSeparationReq) MsgId() int32 {
	return 20214
}

func (*SwitchSeparationReq) Comment() string {
	return "切换分身"
}
		
func (*SwitchSeparationResp) MsgId() int32 {
	return 214
}

func (*SwitchSeparationResp) Comment() string {
	return "切换分身"
}
		
func (*GetSeparationDataMsgListReq) MsgId() int32 {
	return 20215
}

func (*GetSeparationDataMsgListReq) Comment() string {
	return "获取分身列表数据"
}
		
func (*GetSeparationDataMsgListResp) MsgId() int32 {
	return 215
}

func (*GetSeparationDataMsgListResp) Comment() string {
	return "获取分身列表数据"
}
		
func (*SetMessageSubscribeDataReqMsg) MsgId() int32 {
	return 20251
}

func (*SetMessageSubscribeDataReqMsg) Comment() string {
	return "设置订阅消息"
}
		
func (*SetMessageSubscribeDataRspMsg) MsgId() int32 {
	return 251
}

func (*SetMessageSubscribeDataRspMsg) Comment() string {
	return "设置订阅消息"
}
		
func (*AuthorizePlayerHeadReq) MsgId() int32 {
	return 20252
}

func (*AuthorizePlayerHeadReq) Comment() string {
	return "微信头像授权"
}
		
func (*AuthorizePlayerHeadResp) MsgId() int32 {
	return 252
}

func (*AuthorizePlayerHeadResp) Comment() string {
	return "微信头像授权"
}
		
func (*UsePropReq) MsgId() int32 {
	return 20302
}

func (*UsePropReq) Comment() string {
	return "使用道具"
}
		
func (*UsePropResp) MsgId() int32 {
	return 302
}

func (*UsePropResp) Comment() string {
	return "使用道具"
}
		
func (*CompoundPropReq) MsgId() int32 {
	return 20303
}

func (*CompoundPropReq) Comment() string {
	return "合成道具"
}
		
func (*CompoundPropResp) MsgId() int32 {
	return 303
}

func (*CompoundPropResp) Comment() string {
	return "合成道具"
}
		
func (*RankBattleChallengeReq) MsgId() int32 {
	return 20412
}

func (*RankBattleChallengeReq) Comment() string {
	return "挑战请求"
}
		
func (*RankBattleChallengeResp) MsgId() int32 {
	return 412
}

func (*RankBattleChallengeResp) Comment() string {
	return "挑战请求"
}
		
func (*RankBattleRevengeReq) MsgId() int32 {
	return 20414
}

func (*RankBattleRevengeReq) Comment() string {
	return "妖榜复仇请求"
}
		
func (*RankBattleRevengeResp) MsgId() int32 {
	return 414
}

func (*RankBattleRevengeResp) Comment() string {
	return "妖榜复仇请求"
}
		
func (*RankBattleServerListReq) MsgId() int32 {
	return 20418
}

func (*RankBattleServerListReq) Comment() string {
	return "获取跨服斗法参与区服"
}
		
func (*RankBattleServerListResp) MsgId() int32 {
	return 418
}

func (*RankBattleServerListResp) Comment() string {
	return "获取跨服斗法参与区服"
}
		
func (*TaskGetRewardReqMsg) MsgId() int32 {
	return 20503
}

func (*TaskGetRewardReqMsg) Comment() string {
	return "领取任务奖励响应"
}
		
func (*TaskGetRewardRespMsg) MsgId() int32 {
	return 503
}

func (*TaskGetRewardRespMsg) Comment() string {
	return "领取任务奖励响应"
}
		
func (*MailDeleteAllReq) MsgId() int32 {
	return 20553
}

func (*MailDeleteAllReq) Comment() string {
	return "一键删除邮件"
}
		
func (*MailDeleteAllResp) MsgId() int32 {
	return 553
}

func (*MailDeleteAllResp) Comment() string {
	return "一键删除邮件"
}
		
func (*MailDeleteReq) MsgId() int32 {
	return 20554
}

func (*MailDeleteReq) Comment() string {
	return "删除邮件"
}
		
func (*MailDeleteResp) MsgId() int32 {
	return 554
}

func (*MailDeleteResp) Comment() string {
	return "删除邮件"
}
		
func (*MailGetAllRewardReq) MsgId() int32 {
	return 20555
}

func (*MailGetAllRewardReq) Comment() string {
	return "一键领取邮件"
}
		
func (*MailGetAllRewardResp) MsgId() int32 {
	return 555
}

func (*MailGetAllRewardResp) Comment() string {
	return "一键领取邮件"
}
		
func (*MailGetRewardReq) MsgId() int32 {
	return 20556
}

func (*MailGetRewardReq) Comment() string {
	return "领取邮件"
}
		
func (*MailGetRewardResp) MsgId() int32 {
	return 556
}

func (*MailGetRewardResp) Comment() string {
	return "领取邮件"
}
		
func (*MailReadReq) MsgId() int32 {
	return 20557
}

func (*MailReadReq) Comment() string {
	return "读取邮件"
}
		
func (*MailReadResp) MsgId() int32 {
	return 557
}

func (*MailReadResp) Comment() string {
	return "读取邮件"
}
		
func (*BuyGoodsReq) MsgId() int32 {
	return 20601
}

func (*BuyGoodsReq) Comment() string {
	return "购买商品"
}
		
func (*BuyGoodsResp) MsgId() int32 {
	return 601
}

func (*BuyGoodsResp) Comment() string {
	return "购买商品"
}
		
func (*MallRandomGoodsReqMsg) MsgId() int32 {
	return 20603
}

func (*MallRandomGoodsReqMsg) Comment() string {
	return "商店列表"
}
		
func (*MallRandomGoodsRespMsg) MsgId() int32 {
	return 603
}

func (*MallRandomGoodsRespMsg) Comment() string {
	return "商店列表"
}
		
func (*MallChanceUseReqMsg) MsgId() int32 {
	return 20604
}

func (*MallChanceUseReqMsg) Comment() string {
	return "通宝购买请求"
}
		
func (*MallChanceUseRespMsg) MsgId() int32 {
	return 604
}

func (*MallChanceUseRespMsg) Comment() string {
	return "通宝购买请求"
}
		
func (*RandomTalentReq) MsgId() int32 {
	return 20622
}

func (*RandomTalentReq) Comment() string {
	return "炼化"
}
		
func (*RandomTalentResp) MsgId() int32 {
	return 622
}

func (*RandomTalentResp) Comment() string {
	return "炼化"
}
		
func (*DealTalentReq) MsgId() int32 {
	return 20623
}

func (*DealTalentReq) Comment() string {
	return "处理神通"
}
		
func (*DealTalentResp) MsgId() int32 {
	return 623
}

func (*DealTalentResp) Comment() string {
	return "处理神通"
}
		
func (*ReadBookReq) MsgId() int32 {
	return 20624
}

func (*ReadBookReq) Comment() string {
	return "阅读妖书"
}
		
func (*ReadBookResp) MsgId() int32 {
	return 624
}

func (*ReadBookResp) Comment() string {
	return "阅读妖书"
}
		
func (*GetUnDealTalentMsgReq) MsgId() int32 {
	return 20625
}

func (*GetUnDealTalentMsgReq) Comment() string {
	return "获取未处理神通数据"
}
		
func (*GetUnDealTalentMsgResp) MsgId() int32 {
	return 625
}

func (*GetUnDealTalentMsgResp) Comment() string {
	return "获取未处理神通数据"
}
		
func (*ReqDestinyGift) MsgId() int32 {
	return 20652
}

func (*ReqDestinyGift) Comment() string {
	return "赠礼"
}
		
func (*RspDestinyGift) MsgId() int32 {
	return 652
}

func (*RspDestinyGift) Comment() string {
	return "赠礼"
}
		
func (*ReqDestinyTravel) MsgId() int32 {
	return 20653
}

func (*ReqDestinyTravel) Comment() string {
	return "游历"
}
		
func (*RspDestinyTravel) MsgId() int32 {
	return 653
}

func (*RspDestinyTravel) Comment() string {
	return "游历"
}
		
func (*ReqDestinyChallenge) MsgId() int32 {
	return 20654
}

func (*ReqDestinyChallenge) Comment() string {
	return "切磋"
}
		
func (*RspDestinyChallenge) MsgId() int32 {
	return 654
}

func (*RspDestinyChallenge) Comment() string {
	return "切磋"
}
		
func (*ReqUnlockDestinyByItem) MsgId() int32 {
	return 20655
}

func (*ReqUnlockDestinyByItem) Comment() string {
	return "解锁道具仙友"
}
		
func (*RspUnlockDestinyByItem) MsgId() int32 {
	return 655
}

func (*RspUnlockDestinyByItem) Comment() string {
	return "解锁道具仙友"
}
		
func (*ReqUnlockDestinySkinByItem) MsgId() int32 {
	return 20657
}

func (*ReqUnlockDestinySkinByItem) Comment() string {
	return "皮肤解锁"
}
		
func (*RspUnlockDestinySkinByItem) MsgId() int32 {
	return 657
}

func (*RspUnlockDestinySkinByItem) Comment() string {
	return "皮肤解锁"
}
		
func (*ReqEnhanceDestinySkin) MsgId() int32 {
	return 20658
}

func (*ReqEnhanceDestinySkin) Comment() string {
	return "强化皮肤"
}
		
func (*RspEnhanceDestinySkin) MsgId() int32 {
	return 658
}

func (*RspEnhanceDestinySkin) Comment() string {
	return "强化皮肤"
}
		
func (*ReqWearDestinySkin) MsgId() int32 {
	return 20659
}

func (*ReqWearDestinySkin) Comment() string {
	return "穿戴皮肤"
}
		
func (*RspWearDestinySkin) MsgId() int32 {
	return 659
}

func (*RspWearDestinySkin) Comment() string {
	return "穿戴皮肤"
}
		
func (*DestinySwitchLinkageSkinReq) MsgId() int32 {
	return 20660
}

func (*DestinySwitchLinkageSkinReq) Comment() string {
	return "切换挚友联动形象"
}
		
func (*DestinySwitchLinkageSkinResp) MsgId() int32 {
	return 660
}

func (*DestinySwitchLinkageSkinResp) Comment() string {
	return "切换挚友联动形象"
}
		
func (*UnlockCloudReq) MsgId() int32 {
	return 20712
}

func (*UnlockCloudReq) Comment() string {
	return "解锁筋斗云请求"
}
		
func (*UnlockCloudResp) MsgId() int32 {
	return 712
}

func (*UnlockCloudResp) Comment() string {
	return "解锁筋斗云请求"
}
		
func (*CloudLvUpReq) MsgId() int32 {
	return 20713
}

func (*CloudLvUpReq) Comment() string {
	return "筋斗云升级请求"
}
		
func (*CloudLvUpResp) MsgId() int32 {
	return 713
}

func (*CloudLvUpResp) Comment() string {
	return "筋斗云升级请求"
}
		
func (*CloudStageUpReq) MsgId() int32 {
	return 20714
}

func (*CloudStageUpReq) Comment() string {
	return "筋斗云升阶请求"
}
		
func (*CloudStageUpResp) MsgId() int32 {
	return 714
}

func (*CloudStageUpResp) Comment() string {
	return "筋斗云升阶请求"
}
		
func (*CloudEquipReq) MsgId() int32 {
	return 20715
}

func (*CloudEquipReq) Comment() string {
	return "筋斗云穿戴请求"
}
		
func (*CloudEquipResp) MsgId() int32 {
	return 715
}

func (*CloudEquipResp) Comment() string {
	return "筋斗云穿戴请求"
}
		
func (*CloudEquipSkinReq) MsgId() int32 {
	return 20716
}

func (*CloudEquipSkinReq) Comment() string {
	return "筋斗云皮肤穿戴请求"
}
		
func (*CloudEquipSkinResp) MsgId() int32 {
	return 716
}

func (*CloudEquipSkinResp) Comment() string {
	return "筋斗云皮肤穿戴请求"
}
		
func (*CloudSkinLvUpReq) MsgId() int32 {
	return 20717
}

func (*CloudSkinLvUpReq) Comment() string {
	return "筋斗云皮肤升级 0->1 为激活"
}
		
func (*CloudSkinLvUpResp) MsgId() int32 {
	return 717
}

func (*CloudSkinLvUpResp) Comment() string {
	return "筋斗云皮肤升级 0->1 为激活"
}
		
func (*WildBossChallengeReq) MsgId() int32 {
	return 20732
}

func (*WildBossChallengeReq) Comment() string {
	return "请求宝地挑战"
}
		
func (*WildBossChallengeResp) MsgId() int32 {
	return 732
}

func (*WildBossChallengeResp) Comment() string {
	return "请求宝地挑战"
}
		
func (*WildBossRepeatReq) MsgId() int32 {
	return 20733
}

func (*WildBossRepeatReq) Comment() string {
	return "请求宝地扫荡"
}
		
func (*WildBossRepeatResp) MsgId() int32 {
	return 733
}

func (*WildBossRepeatResp) Comment() string {
	return "请求宝地扫荡"
}
		
func (*CatchPetReq) MsgId() int32 {
	return 20741
}

func (*CatchPetReq) Comment() string {
	return "捕捉灵兽请求"
}
		
func (*CatchPetResp) MsgId() int32 {
	return 741
}

func (*CatchPetResp) Comment() string {
	return "捕捉灵兽请求"
}
		
func (*RefreshPetPoolReq) MsgId() int32 {
	return 20742
}

func (*RefreshPetPoolReq) Comment() string {
	return "刷新灵兽池子请求"
}
		
func (*RefreshPetPoolResp) MsgId() int32 {
	return 742
}

func (*RefreshPetPoolResp) Comment() string {
	return "刷新灵兽池子请求"
}
		
func (*AddPetBagCountReq) MsgId() int32 {
	return 20743
}

func (*AddPetBagCountReq) Comment() string {
	return "开启灵兽背包格子请求"
}
		
func (*AddPetBagCountResp) MsgId() int32 {
	return 743
}

func (*AddPetBagCountResp) Comment() string {
	return "开启灵兽背包格子请求"
}
		
func (*ReleasePetReq) MsgId() int32 {
	return 20744
}

func (*ReleasePetReq) Comment() string {
	return "灵兽放生请求"
}
		
func (*ReleasePetResp) MsgId() int32 {
	return 744
}

func (*ReleasePetResp) Comment() string {
	return "灵兽放生请求"
}
		
func (*PetLevelUpReq) MsgId() int32 {
	return 20745
}

func (*PetLevelUpReq) Comment() string {
	return "灵兽升级请求"
}
		
func (*PetLevelUpResp) MsgId() int32 {
	return 745
}

func (*PetLevelUpResp) Comment() string {
	return "灵兽升级请求"
}
		
func (*PetGobbleUpReq) MsgId() int32 {
	return 20746
}

func (*PetGobbleUpReq) Comment() string {
	return "灵兽吞噬请求"
}
		
func (*PetGobbleUpResp) MsgId() int32 {
	return 746
}

func (*PetGobbleUpResp) Comment() string {
	return "灵兽吞噬请求"
}
		
func (*EquipPetReq) MsgId() int32 {
	return 20747
}

func (*EquipPetReq) Comment() string {
	return "携带灵兽请求"
}
		
func (*EquipPetResp) MsgId() int32 {
	return 747
}

func (*EquipPetResp) Comment() string {
	return "携带灵兽请求"
}
		
func (*PetResetReq) MsgId() int32 {
	return 20748
}

func (*PetResetReq) Comment() string {
	return "携带灵兽请求"
}
		
func (*PetResetResp) MsgId() int32 {
	return 748
}

func (*PetResetResp) Comment() string {
	return "携带灵兽请求"
}
		
func (*ChooseWishPetReq) MsgId() int32 {
	return 20749
}

func (*ChooseWishPetReq) Comment() string {
	return "灵兽心愿选择"
}
		
func (*ChooseWishPetResp) MsgId() int32 {
	return 749
}

func (*ChooseWishPetResp) Comment() string {
	return "灵兽心愿选择"
}
		
func (*PetAssistantReq) MsgId() int32 {
	return 20750
}

func (*PetAssistantReq) Comment() string {
	return "灵兽协助"
}
		
func (*PetAssistantResp) MsgId() int32 {
	return 750
}

func (*PetAssistantResp) Comment() string {
	return "灵兽协助"
}
		
func (*PetSoulShapeReq) MsgId() int32 {
	return 20751
}

func (*PetSoulShapeReq) Comment() string {
	return "灵兽塑魂"
}
		
func (*PetSoulShapeResp) MsgId() int32 {
	return 751
}

func (*PetSoulShapeResp) Comment() string {
	return "灵兽塑魂"
}
		
func (*SelectSoulShapeSkillReq) MsgId() int32 {
	return 20752
}

func (*SelectSoulShapeSkillReq) Comment() string {
	return "灵兽塑魂技能选择"
}
		
func (*SelectSoulShapeSkillResp) MsgId() int32 {
	return 752
}

func (*SelectSoulShapeSkillResp) Comment() string {
	return "灵兽塑魂技能选择"
}
		
func (*PetSkillRefreshViewReq) MsgId() int32 {
	return 20753
}

func (*PetSkillRefreshViewReq) Comment() string {
	return "灵兽洗练界面"
}
		
func (*PetSkillRefreshViewResp) MsgId() int32 {
	return 753
}

func (*PetSkillRefreshViewResp) Comment() string {
	return "灵兽洗练界面"
}
		
func (*PetSkillRefreshReq) MsgId() int32 {
	return 20754
}

func (*PetSkillRefreshReq) Comment() string {
	return "灵兽洗练"
}
		
func (*PetSkillRefreshResp) MsgId() int32 {
	return 754
}

func (*PetSkillRefreshResp) Comment() string {
	return "灵兽洗练"
}
		
func (*PetSkillRefreshResultReq) MsgId() int32 {
	return 20755
}

func (*PetSkillRefreshResultReq) Comment() string {
	return "灵兽洗练结果处理"
}
		
func (*PetSkillRefreshResultResp) MsgId() int32 {
	return 755
}

func (*PetSkillRefreshResultResp) Comment() string {
	return "灵兽洗练结果处理"
}
		
func (*PetLockStateReq) MsgId() int32 {
	return 20756
}

func (*PetLockStateReq) Comment() string {
	return "灵兽加锁"
}
		
func (*PetLockStateResp) MsgId() int32 {
	return 756
}

func (*PetLockStateResp) Comment() string {
	return "灵兽加锁"
}
		
func (*PetSwitchLinkageSkinReq) MsgId() int32 {
	return 20757
}

func (*PetSwitchLinkageSkinReq) Comment() string {
	return "切换灵兽联动形象"
}
		
func (*PetSwitchLinkageSkinResp) MsgId() int32 {
	return 757
}

func (*PetSwitchLinkageSkinResp) Comment() string {
	return "切换灵兽联动形象"
}
		
func (*PetLockReq) MsgId() int32 {
	return 20758
}

func (*PetLockReq) Comment() string {
	return "灵兽新版上锁（需要密码解锁）"
}
		
func (*PetLockResp) MsgId() int32 {
	return 758
}

func (*PetLockResp) Comment() string {
	return "灵兽新版上锁（需要密码解锁）"
}
		
func (*SelectBuffReq) MsgId() int32 {
	return 20764
}

func (*SelectBuffReq) Comment() string {
	return "选择加成请求"
}
		
func (*SelectBuffResp) MsgId() int32 {
	return 764
}

func (*SelectBuffResp) Comment() string {
	return "选择加成请求"
}
		
func (*GetSelectOptionReq) MsgId() int32 {
	return 20766
}

func (*GetSelectOptionReq) Comment() string {
	return "获取选择选项"
}
		
func (*GetSelectOptionResp) MsgId() int32 {
	return 766
}

func (*GetSelectOptionResp) Comment() string {
	return "获取选择选项"
}
		
func (*SaveSelectOptionReq) MsgId() int32 {
	return 20767
}

func (*SaveSelectOptionReq) Comment() string {
	return "保存选择选项"
}
		
func (*SaveSelectOptionResp) MsgId() int32 {
	return 767
}

func (*SaveSelectOptionResp) Comment() string {
	return "保存选择选项"
}
		
func (*GetRankListReq) MsgId() int32 {
	return 20801
}

func (*GetRankListReq) Comment() string {
	return "获取排行列表"
}
		
func (*GetRankListResp) MsgId() int32 {
	return 801
}

func (*GetRankListResp) Comment() string {
	return "获取排行列表"
}
		
func (*GetMallRewardReq) MsgId() int32 {
	return 20802
}

func (*GetMallRewardReq) Comment() string {
	return "首充领取奖励"
}
		
func (*GetMallRewardResp) MsgId() int32 {
	return 802
}

func (*GetMallRewardResp) Comment() string {
	return "首充领取奖励"
}
		
func (*GetUnionMemberScoreReq) MsgId() int32 {
	return 20803
}

func (*GetUnionMemberScoreReq) Comment() string {
	return "查看联盟积分"
}
		
func (*GetUnionMemberScoreRsp) MsgId() int32 {
	return 803
}

func (*GetUnionMemberScoreRsp) Comment() string {
	return "查看联盟积分"
}
		
func (*SpiritDrawReq) MsgId() int32 {
	return 20822
}

func (*SpiritDrawReq) Comment() string {
	return "精怪抽取请求"
}
		
func (*SpiritDrawResp) MsgId() int32 {
	return 822
}

func (*SpiritDrawResp) Comment() string {
	return "精怪抽取请求"
}
		
func (*SpiritUnlockReq) MsgId() int32 {
	return 20823
}

func (*SpiritUnlockReq) Comment() string {
	return "精怪解锁请求"
}
		
func (*SpiritUnlockResp) MsgId() int32 {
	return 823
}

func (*SpiritUnlockResp) Comment() string {
	return "精怪解锁请求"
}
		
func (*SpiritLvUpReq) MsgId() int32 {
	return 20824
}

func (*SpiritLvUpReq) Comment() string {
	return "精怪升级请求"
}
		
func (*SpiritLvUpResp) MsgId() int32 {
	return 824
}

func (*SpiritLvUpResp) Comment() string {
	return "精怪升级请求"
}
		
func (*SwitchBattleTeamReq) MsgId() int32 {
	return 20825
}

func (*SwitchBattleTeamReq) Comment() string {
	return "切换出战队伍请求"
}
		
func (*SwitchBattleTeamResp) MsgId() int32 {
	return 825
}

func (*SwitchBattleTeamResp) Comment() string {
	return "切换出战队伍请求"
}
		
func (*SpiritBattleReq) MsgId() int32 {
	return 20826
}

func (*SpiritBattleReq) Comment() string {
	return "精怪上下阵请求"
}
		
func (*SpiritBattleResp) MsgId() int32 {
	return 826
}

func (*SpiritBattleResp) Comment() string {
	return "精怪上下阵请求"
}
		
func (*SpiritCombineLvUpReq) MsgId() int32 {
	return 20827
}

func (*SpiritCombineLvUpReq) Comment() string {
	return "精怪共鸣等级提升"
}
		
func (*SpiritCombineLvUpResp) MsgId() int32 {
	return 827
}

func (*SpiritCombineLvUpResp) Comment() string {
	return "精怪共鸣等级提升"
}
		
func (*SpiritSwitchLinkageSkinReq) MsgId() int32 {
	return 20828
}

func (*SpiritSwitchLinkageSkinReq) Comment() string {
	return "切换精怪联动形象"
}
		
func (*SpiritSwitchLinkageSkinResp) MsgId() int32 {
	return 828
}

func (*SpiritSwitchLinkageSkinResp) Comment() string {
	return "切换精怪联动形象"
}
		
func (*SpiritLevelUnlockShowReq) MsgId() int32 {
	return 20829
}

func (*SpiritLevelUnlockShowReq) Comment() string {
	return "精怪等级解锁展示"
}
		
func (*SpiritLevelUnlockShowResp) MsgId() int32 {
	return 829
}

func (*SpiritLevelUnlockShowResp) Comment() string {
	return "精怪等级解锁展示"
}
		
func (*ReqGetActivityDetail) MsgId() int32 {
	return 21003
}

func (*ReqGetActivityDetail) Comment() string {
	return "请求活动通用数据"
}
		
func (*RspGetActivityDetail) MsgId() int32 {
	return 1003
}

func (*RspGetActivityDetail) Comment() string {
	return "请求活动通用数据"
}
		
func (*ReqGetActivityConditionReward) MsgId() int32 {
	return 21004
}

func (*ReqGetActivityConditionReward) Comment() string {
	return "领取活动任务条件奖励"
}
		
func (*RspGetActivityConditionReward) MsgId() int32 {
	return 1004
}

func (*RspGetActivityConditionReward) Comment() string {
	return "领取活动任务条件奖励"
}
		
func (*ReqBuyActivityGoods) MsgId() int32 {
	return 21005
}

func (*ReqBuyActivityGoods) Comment() string {
	return "购买活动商品"
}
		
func (*RspBuyActivityGoods) MsgId() int32 {
	return 1005
}

func (*RspBuyActivityGoods) Comment() string {
	return "购买活动商品"
}
		
func (*ReqGetActivityRankReward) MsgId() int32 {
	return 21006
}

func (*ReqGetActivityRankReward) Comment() string {
	return "领取活动排行奖励"
}
		
func (*RespGetActivityRankReward) MsgId() int32 {
	return 1006
}

func (*RespGetActivityRankReward) Comment() string {
	return "领取活动排行奖励"
}
		
func (*ReqGetFundsConditionReward) MsgId() int32 {
	return 21010
}

func (*ReqGetFundsConditionReward) Comment() string {
	return "领取基金活动任务条件奖励"
}
		
func (*RspGetFundsConditionReward) MsgId() int32 {
	return 1010
}

func (*RspGetFundsConditionReward) Comment() string {
	return "领取基金活动任务条件奖励"
}
		
func (*LuckyDrawReq) MsgId() int32 {
	return 21013
}

func (*LuckyDrawReq) Comment() string {
	return "运势抽奖请求"
}
		
func (*LuckyDrawResp) MsgId() int32 {
	return 1013
}

func (*LuckyDrawResp) Comment() string {
	return "运势抽奖请求"
}
		
func (*SelectItemsReq) MsgId() int32 {
	return 21014
}

func (*SelectItemsReq) Comment() string {
	return "自选礼包选择道具"
}
		
func (*SelectItemsResp) MsgId() int32 {
	return 1014
}

func (*SelectItemsResp) Comment() string {
	return "自选礼包选择道具"
}
		
func (*ReqShareTaskDone) MsgId() int32 {
	return 21016
}

func (*ReqShareTaskDone) Comment() string {
	return "活动分享任务完成请求"
}
		
func (*RespShareTaskDone) MsgId() int32 {
	return 1016
}

func (*RespShareTaskDone) Comment() string {
	return "活动分享任务完成请求"
}
		
func (*ReqGetActivityRankState) MsgId() int32 {
	return 21017
}

func (*ReqGetActivityRankState) Comment() string {
	return "请求榜单状态"
}
		
func (*RespGetActivityRankState) MsgId() int32 {
	return 1017
}

func (*RespGetActivityRankState) Comment() string {
	return "请求榜单状态"
}
		
func (*ReqGetActivityConditionRewardByArr) MsgId() int32 {
	return 21022
}

func (*ReqGetActivityConditionRewardByArr) Comment() string {
	return "通过数组一次性领取活动任务条件奖励"
}
		
func (*SeekTreasureDrawReq) MsgId() int32 {
	return 21023
}

func (*SeekTreasureDrawReq) Comment() string {
	return "妖市觅宝抽奖"
}
		
func (*SeekTreasureDrawResp) MsgId() int32 {
	return 1023
}

func (*SeekTreasureDrawResp) Comment() string {
	return "妖市觅宝抽奖"
}
		
func (*SeekTreasureViewReq) MsgId() int32 {
	return 21024
}

func (*SeekTreasureViewReq) Comment() string {
	return "妖市觅宝界面"
}
		
func (*SeekTreasureViewResp) MsgId() int32 {
	return 1024
}

func (*SeekTreasureViewResp) Comment() string {
	return "妖市觅宝界面"
}
		
func (*SeekTreasureChooseRareRewardReq) MsgId() int32 {
	return 21025
}

func (*SeekTreasureChooseRareRewardReq) Comment() string {
	return "选择稀有奖励"
}
		
func (*SeekTreasureChooseRareRewardResp) MsgId() int32 {
	return 1025
}

func (*SeekTreasureChooseRareRewardResp) Comment() string {
	return "选择稀有奖励"
}
		
func (*SeekTreasureSelectRewardReqMsg) MsgId() int32 {
	return 21028
}

func (*SeekTreasureSelectRewardReqMsg) Comment() string {
	return "选择稀有奖励"
}
		
func (*SeekTreasureSelectRewardRespMsg) MsgId() int32 {
	return 1028
}

func (*SeekTreasureSelectRewardRespMsg) Comment() string {
	return "选择稀有奖励"
}
		
func (*SelectActivityGoodsReq) MsgId() int32 {
	return 21032
}

func (*SelectActivityGoodsReq) Comment() string {
	return "自选商品选中协议"
}
		
func (*SelectActivityGoodsResp) MsgId() int32 {
	return 1032
}

func (*SelectActivityGoodsResp) Comment() string {
	return "自选商品选中协议"
}
		
func (*SelectActivityConditionGoodsReq) MsgId() int32 {
	return 21034
}

func (*SelectActivityConditionGoodsReq) Comment() string {
	return "活动任务自选商品"
}
		
func (*SelectActivityConditionGoodsResp) MsgId() int32 {
	return 1034
}

func (*SelectActivityConditionGoodsResp) Comment() string {
	return "活动任务自选商品"
}
		
func (*ReqGuessInfoLoadMsg) MsgId() int32 {
	return 21040
}

func (*ReqGuessInfoLoadMsg) Comment() string {
	return "请求竞猜数据"
}
		
func (*RespGuessInfoLoadMsg) MsgId() int32 {
	return 1040
}

func (*RespGuessInfoLoadMsg) Comment() string {
	return "请求竞猜数据"
}
		
func (*ReqGuessMsg) MsgId() int32 {
	return 21041
}

func (*ReqGuessMsg) Comment() string {
	return "竞猜押注"
}
		
func (*RespGuessMsg) MsgId() int32 {
	return 1041
}

func (*RespGuessMsg) Comment() string {
	return "竞猜押注"
}
		
func (*ReqGuessRewardMsg) MsgId() int32 {
	return 21042
}

func (*ReqGuessRewardMsg) Comment() string {
	return "领取竞猜奖励"
}
		
func (*RespGuessRewardMsg) MsgId() int32 {
	return 1042
}

func (*RespGuessRewardMsg) Comment() string {
	return "领取竞猜奖励"
}
		
func (*ReqGuessResultDetailMsg) MsgId() int32 {
	return 21043
}

func (*ReqGuessResultDetailMsg) Comment() string {
	return "查看竞猜奖励名单"
}
		
func (*RespGuessResultDetailMsg) MsgId() int32 {
	return 1043
}

func (*RespGuessResultDetailMsg) Comment() string {
	return "查看竞猜奖励名单"
}
		
func (*ReceiveCrossUnionRechargeRewardReq) MsgId() int32 {
	return 21045
}

func (*ReceiveCrossUnionRechargeRewardReq) Comment() string {
	return "领取跨服获得内置的 妖盟充值活动的奖励"
}
		
func (*ReceiveCrossUnionRechargeRewardResp) MsgId() int32 {
	return 1045
}

func (*ReceiveCrossUnionRechargeRewardResp) Comment() string {
	return "领取跨服获得内置的 妖盟充值活动的奖励"
}
		
func (*ActivityGetJoinMemberListReq) MsgId() int32 {
	return 21046
}

func (*ActivityGetJoinMemberListReq) Comment() string {
	return "查看自己所属妖盟的参与名单"
}
		
func (*ActivityGetJoinMemberListRsp) MsgId() int32 {
	return 1046
}

func (*ActivityGetJoinMemberListRsp) Comment() string {
	return "查看自己所属妖盟的参与名单"
}
		
func (*HomelandEnterReq) MsgId() int32 {
	return 21052
}

func (*HomelandEnterReq) Comment() string {
	return "进入福地"
}
		
func (*HomelandEnterResp) MsgId() int32 {
	return 1052
}

func (*HomelandEnterResp) Comment() string {
	return "进入福地"
}
		
func (*HomelandManageReq) MsgId() int32 {
	return 21053
}

func (*HomelandManageReq) Comment() string {
	return "福地管理界面"
}
		
func (*HomelandManageResp) MsgId() int32 {
	return 1053
}

func (*HomelandManageResp) Comment() string {
	return "福地管理界面"
}
		
func (*HomelandLevelUpReq) MsgId() int32 {
	return 21054
}

func (*HomelandLevelUpReq) Comment() string {
	return "福地升级"
}
		
func (*HomelandLevelUpResp) MsgId() int32 {
	return 1054
}

func (*HomelandLevelUpResp) Comment() string {
	return "福地升级"
}
		
func (*HomelandRefreshReq) MsgId() int32 {
	return 21055
}

func (*HomelandRefreshReq) Comment() string {
	return "福地刷新资源"
}
		
func (*HomelandRefreshResp) MsgId() int32 {
	return 1055
}

func (*HomelandRefreshResp) Comment() string {
	return "福地刷新资源"
}
		
func (*HomelandBuyWorkerReq) MsgId() int32 {
	return 21056
}

func (*HomelandBuyWorkerReq) Comment() string {
	return "福地购买鼠宝"
}
		
func (*HomelandBuyWorkerResp) MsgId() int32 {
	return 1056
}

func (*HomelandBuyWorkerResp) Comment() string {
	return "福地购买鼠宝"
}
		
func (*HomelandGetRecordReq) MsgId() int32 {
	return 21057
}

func (*HomelandGetRecordReq) Comment() string {
	return "福地日志"
}
		
func (*HomelandGetRecordResp) MsgId() int32 {
	return 1057
}

func (*HomelandGetRecordResp) Comment() string {
	return "福地日志"
}
		
func (*HomelandExploreReq) MsgId() int32 {
	return 21058
}

func (*HomelandExploreReq) Comment() string {
	return "福地探寻"
}
		
func (*HomelandExploreResp) MsgId() int32 {
	return 1058
}

func (*HomelandExploreResp) Comment() string {
	return "福地探寻"
}
		
func (*HomelandRefreshExploreReq) MsgId() int32 {
	return 21059
}

func (*HomelandRefreshExploreReq) Comment() string {
	return "福地探寻刷新"
}
		
func (*HomelandRefreshExploreResp) MsgId() int32 {
	return 1059
}

func (*HomelandRefreshExploreResp) Comment() string {
	return "福地探寻刷新"
}
		
func (*HomelandDispatchWorkerReq) MsgId() int32 {
	return 21060
}

func (*HomelandDispatchWorkerReq) Comment() string {
	return "福地派遣鼠宝"
}
		
func (*HomelandDispatchWorkerResp) MsgId() int32 {
	return 1060
}

func (*HomelandDispatchWorkerResp) Comment() string {
	return "福地派遣鼠宝"
}
		
func (*HomelandDispatchPreviewReq) MsgId() int32 {
	return 21061
}

func (*HomelandDispatchPreviewReq) Comment() string {
	return "福地派遣鼠宝预览"
}
		
func (*HomelandDispatchPreviewResp) MsgId() int32 {
	return 1061
}

func (*HomelandDispatchPreviewResp) Comment() string {
	return "福地派遣鼠宝预览"
}
		
func (*HomelandLeaveReq) MsgId() int32 {
	return 21063
}

func (*HomelandLeaveReq) Comment() string {
	return "离开福地"
}
		
func (*HomelandLeaveResp) MsgId() int32 {
	return 1063
}

func (*HomelandLeaveResp) Comment() string {
	return "离开福地"
}
		
func (*HomelandGetRewardReq) MsgId() int32 {
	return 21065
}

func (*HomelandGetRewardReq) Comment() string {
	return "福地获取奖励"
}
		
func (*HomelandGetRewardResp) MsgId() int32 {
	return 1065
}

func (*HomelandGetRewardResp) Comment() string {
	return "福地获取奖励"
}
		
func (*CornucopiaGetRewardReq) MsgId() int32 {
	return 21067
}

func (*CornucopiaGetRewardReq) Comment() string {
	return "福地聚宝盆领奖"
}
		
func (*CornucopiaGetRewardRsp) MsgId() int32 {
	return 1067
}

func (*CornucopiaGetRewardRsp) Comment() string {
	return "福地聚宝盆领奖"
}
		
func (*HomelandAutoDispatchWorkerReq) MsgId() int32 {
	return 21068
}

func (*HomelandAutoDispatchWorkerReq) Comment() string {
	return "福地鼠鼠自动采集"
}
		
func (*HomelandAutoDispatchWorkerRsp) MsgId() int32 {
	return 1068
}

func (*HomelandAutoDispatchWorkerRsp) Comment() string {
	return "福地鼠鼠自动采集"
}
		
func (*HomelandUseFreeMouseManagerReq) MsgId() int32 {
	return 21069
}

func (*HomelandUseFreeMouseManagerReq) Comment() string {
	return "福地鼠宝总管试用"
}
		
func (*HomelandUseFreeMouseManagerRsp) MsgId() int32 {
	return 1069
}

func (*HomelandUseFreeMouseManagerRsp) Comment() string {
	return "福地鼠宝总管试用"
}
		
func (*HomelandSaveSettingsReq) MsgId() int32 {
	return 21070
}

func (*HomelandSaveSettingsReq) Comment() string {
	return "鼠鼠总管保存设置"
}
		
func (*HomelandSaveSettingsRsp) MsgId() int32 {
	return 1070
}

func (*HomelandSaveSettingsRsp) Comment() string {
	return "鼠鼠总管保存设置"
}
		
func (*HomelandDispatchAllWorkerReq) MsgId() int32 {
	return 21071
}

func (*HomelandDispatchAllWorkerReq) Comment() string {
	return "桃源一键召回小妖"
}
		
func (*HomelandDispatchAllWorkerResp) MsgId() int32 {
	return 1071
}

func (*HomelandDispatchAllWorkerResp) Comment() string {
	return "桃源一键召回小妖"
}
		
func (*BallGVGApplicationReqMsg) MsgId() int32 {
	return 21101
}

func (*BallGVGApplicationReqMsg) Comment() string {
	return "请求报名"
}
		
func (*BallGVGApplicationRespMsg) MsgId() int32 {
	return 1101
}

func (*BallGVGApplicationRespMsg) Comment() string {
	return "请求报名"
}
		
func (*BallGVGAbilityRankNumReqMsg) MsgId() int32 {
	return 21102
}

func (*BallGVGAbilityRankNumReqMsg) Comment() string {
	return "请求报名期战力排名"
}
		
func (*BallGVGAbilityRankNumRespMsg) MsgId() int32 {
	return 1102
}

func (*BallGVGAbilityRankNumRespMsg) Comment() string {
	return "请求报名期战力排名"
}
		
func (*BallGVGEnterActivityReqMsg) MsgId() int32 {
	return 21103
}

func (*BallGVGEnterActivityReqMsg) Comment() string {
	return "进入pvp活动请求协议"
}
		
func (*BallGVGEnterActivityRespMsg) MsgId() int32 {
	return 1103
}

func (*BallGVGEnterActivityRespMsg) Comment() string {
	return "进入pvp活动请求协议"
}
		
func (*BallGVGEnterPlaceReqMsg) MsgId() int32 {
	return 21104
}

func (*BallGVGEnterPlaceReqMsg) Comment() string {
	return "进入宝地(房间)请求协议"
}
		
func (*BallGVGEnterPlaceRespMsg) MsgId() int32 {
	return 1104
}

func (*BallGVGEnterPlaceRespMsg) Comment() string {
	return "进入宝地(房间)请求协议"
}
		
func (*BallGVGLeavePlaceReqMsg) MsgId() int32 {
	return 21105
}

func (*BallGVGLeavePlaceReqMsg) Comment() string {
	return "离开当前房间请求协议"
}
		
func (*BallGVGLeavePlaceRespMsg) MsgId() int32 {
	return 1105
}

func (*BallGVGLeavePlaceRespMsg) Comment() string {
	return "离开当前房间请求协议"
}
		
func (*BallGVGMoveReqMsg) MsgId() int32 {
	return 21106
}

func (*BallGVGMoveReqMsg) Comment() string {
	return "移动协议"
}
		
func (*BallGVGMoveRespMsg) MsgId() int32 {
	return 1106
}

func (*BallGVGMoveRespMsg) Comment() string {
	return "移动协议"
}
		
func (*BallGVGAttackReqMsg) MsgId() int32 {
	return 21107
}

func (*BallGVGAttackReqMsg) Comment() string {
	return "攻击协议"
}
		
func (*BallGVGAttackRespMsg) MsgId() int32 {
	return 1107
}

func (*BallGVGAttackRespMsg) Comment() string {
	return "攻击协议"
}
		
func (*BallGVGAppointMajorUserReqMsg) MsgId() int32 {
	return 21108
}

func (*BallGVGAppointMajorUserReqMsg) Comment() string {
	return "任命宗主权限"
}
		
func (*BallGVGAppointMajorUserRespMsg) MsgId() int32 {
	return 1108
}

func (*BallGVGAppointMajorUserRespMsg) Comment() string {
	return "任命宗主权限"
}
		
func (*BallGVGMarkPlaceReqMsg) MsgId() int32 {
	return 21109
}

func (*BallGVGMarkPlaceReqMsg) Comment() string {
	return "标记地点(重复标记视为取消标记)"
}
		
func (*BallGVGMarkPlaceRespMsg) MsgId() int32 {
	return 1109
}

func (*BallGVGMarkPlaceRespMsg) Comment() string {
	return "标记地点(重复标记视为取消标记)"
}
		
func (*BallGVGGetPlaceInfoReqMsg) MsgId() int32 {
	return 21111
}

func (*BallGVGGetPlaceInfoReqMsg) Comment() string {
	return "获取所有房间数据"
}
		
func (*BallGVGGetPlaceInfoRespMsg) MsgId() int32 {
	return 1111
}

func (*BallGVGGetPlaceInfoRespMsg) Comment() string {
	return "获取所有房间数据"
}
		
func (*BallGVGGetUserRankReqMsg) MsgId() int32 {
	return 21113
}

func (*BallGVGGetUserRankReqMsg) Comment() string {
	return "请求宗门玩家列表"
}
		
func (*BallGVGGetUserRankRespMsg) MsgId() int32 {
	return 1113
}

func (*BallGVGGetUserRankRespMsg) Comment() string {
	return "请求宗门玩家列表"
}
		
func (*BallGVGAbilityRankReqMsg) MsgId() int32 {
	return 21126
}

func (*BallGVGAbilityRankReqMsg) Comment() string {
	return "请求报名期妖盟战力排行"
}
		
func (*BallGVGAbilityRankRespMsg) MsgId() int32 {
	return 1126
}

func (*BallGVGAbilityRankRespMsg) Comment() string {
	return "请求报名期妖盟战力排行"
}
		
func (*BallGVGUnionNameListReqMsg) MsgId() int32 {
	return 21128
}

func (*BallGVGUnionNameListReqMsg) Comment() string {
	return "获取参与妖盟的名字"
}
		
func (*BallGVGUnionNameListRespMsg) MsgId() int32 {
	return 1128
}

func (*BallGVGUnionNameListRespMsg) Comment() string {
	return "获取参与妖盟的名字"
}
		
func (*BallGVGGetUnionUserListReqMsg) MsgId() int32 {
	return 21129
}

func (*BallGVGGetUnionUserListReqMsg) Comment() string {
	return "获取上古遗迹妖盟成员信息"
}
		
func (*BallGVGGetUnionUserListRespMsg) MsgId() int32 {
	return 1129
}

func (*BallGVGGetUnionUserListRespMsg) Comment() string {
	return "获取上古遗迹妖盟成员信息"
}
		
func (*BallGVGRedDotReqMsg) MsgId() int32 {
	return 21131
}

func (*BallGVGRedDotReqMsg) Comment() string {
	return "上古遗迹红点"
}
		
func (*BallGVGRedDotRespMsg) MsgId() int32 {
	return 1131
}

func (*BallGVGRedDotRespMsg) Comment() string {
	return "上古遗迹红点"
}
		
func (*XyFundGetRewardReq) MsgId() int32 {
	return 21421
}

func (*XyFundGetRewardReq) Comment() string {
	return "仙缘进度领奖"
}
		
func (*XyFundGetRewardResp) MsgId() int32 {
	return 1421
}

func (*XyFundGetRewardResp) Comment() string {
	return "仙缘进度领奖"
}
		
func (*ReqUnionSameNameListMsg) MsgId() int32 {
	return 21501
}

func (*ReqUnionSameNameListMsg) Comment() string {
	return "获取商会同名姓名列表"
}
		
func (*RspUnionSameNameListMsg) MsgId() int32 {
	return 1501
}

func (*RspUnionSameNameListMsg) Comment() string {
	return "获取商会同名姓名列表"
}
		
func (*UnionRechargeUserReqMsg) MsgId() int32 {
	return 21511
}

func (*UnionRechargeUserReqMsg) Comment() string {
	return "获取商会充值列表"
}
		
func (*UnionRechargeUserRspMsg) MsgId() int32 {
	return 1511
}

func (*UnionRechargeUserRspMsg) Comment() string {
	return "获取商会充值列表"
}
		
func (*ReqUnionCreate) MsgId() int32 {
	return 22101
}

func (*ReqUnionCreate) Comment() string {
	return "创建妖盟"
}
		
func (*RspUnionCreate) MsgId() int32 {
	return 2101
}

func (*RspUnionCreate) Comment() string {
	return "创建妖盟"
}
		
func (*ReqUnionEnter) MsgId() int32 {
	return 22102
}

func (*ReqUnionEnter) Comment() string {
	return "进入妖盟"
}
		
func (*RspUnionEnter) MsgId() int32 {
	return 2102
}

func (*RspUnionEnter) Comment() string {
	return "进入妖盟"
}
		
func (*ReqUnionList) MsgId() int32 {
	return 22103
}

func (*ReqUnionList) Comment() string {
	return "获取妖盟列表"
}
		
func (*RspUnionList) MsgId() int32 {
	return 2103
}

func (*RspUnionList) Comment() string {
	return "获取妖盟列表"
}
		
func (*ReqUnionJoin) MsgId() int32 {
	return 22104
}

func (*ReqUnionJoin) Comment() string {
	return "加入妖盟"
}
		
func (*RspUnionJoin) MsgId() int32 {
	return 2104
}

func (*RspUnionJoin) Comment() string {
	return "加入妖盟"
}
		
func (*ReqUnionRandomJoin) MsgId() int32 {
	return 22105
}

func (*ReqUnionRandomJoin) Comment() string {
	return "随机加入妖盟"
}
		
func (*RspUnionRandomJoin) MsgId() int32 {
	return 2105
}

func (*RspUnionRandomJoin) Comment() string {
	return "随机加入妖盟"
}
		
func (*ReqUnionDetail) MsgId() int32 {
	return 22106
}

func (*ReqUnionDetail) Comment() string {
	return "妖盟详情"
}
		
func (*RspUnionDetail) MsgId() int32 {
	return 2106
}

func (*RspUnionDetail) Comment() string {
	return "妖盟详情"
}
		
func (*ReqUnionPosition) MsgId() int32 {
	return 22107
}

func (*ReqUnionPosition) Comment() string {
	return "职位变更"
}
		
func (*RspUnionPosition) MsgId() int32 {
	return 2107
}

func (*RspUnionPosition) Comment() string {
	return "职位变更"
}
		
func (*ReqUnionApplyPlayerList) MsgId() int32 {
	return 22108
}

func (*ReqUnionApplyPlayerList) Comment() string {
	return "获取申请加入妖盟的玩家列表"
}
		
func (*RspUnionApplyPlayerList) MsgId() int32 {
	return 2108
}

func (*RspUnionApplyPlayerList) Comment() string {
	return "获取申请加入妖盟的玩家列表"
}
		
func (*ReqUnionRandomState) MsgId() int32 {
	return 22109
}

func (*ReqUnionRandomState) Comment() string {
	return "设置随机加入状态"
}
		
func (*RspUnionRandomState) MsgId() int32 {
	return 2109
}

func (*RspUnionRandomState) Comment() string {
	return "设置随机加入状态"
}
		
func (*ReqUnionLimitRealmsId) MsgId() int32 {
	return 22110
}

func (*ReqUnionLimitRealmsId) Comment() string {
	return "设置入盟限制等级"
}
		
func (*RspUnionLimitRealmsId) MsgId() int32 {
	return 2110
}

func (*RspUnionLimitRealmsId) Comment() string {
	return "设置入盟限制等级"
}
		
func (*ReqUnionClearApply) MsgId() int32 {
	return 22111
}

func (*ReqUnionClearApply) Comment() string {
	return "清理所有申请玩家"
}
		
func (*RspUnionClearApply) MsgId() int32 {
	return 2111
}

func (*RspUnionClearApply) Comment() string {
	return "清理所有申请玩家"
}
		
func (*ReqUnionPassApply) MsgId() int32 {
	return 22112
}

func (*ReqUnionPassApply) Comment() string {
	return "通过新玩家的申请"
}
		
func (*RspUnionPassApply) MsgId() int32 {
	return 2112
}

func (*RspUnionPassApply) Comment() string {
	return "通过新玩家的申请"
}
		
func (*ReqUnionRemove) MsgId() int32 {
	return 22113
}

func (*ReqUnionRemove) Comment() string {
	return "解散妖盟"
}
		
func (*RspUnionRemove) MsgId() int32 {
	return 2113
}

func (*RspUnionRemove) Comment() string {
	return "解散妖盟"
}
		
func (*ReqUnionExit) MsgId() int32 {
	return 22114
}

func (*ReqUnionExit) Comment() string {
	return "退出妖盟"
}
		
func (*RspUnionExit) MsgId() int32 {
	return 2114
}

func (*RspUnionExit) Comment() string {
	return "退出妖盟"
}
		
func (*ReqUnionTrends) MsgId() int32 {
	return 22115
}

func (*ReqUnionTrends) Comment() string {
	return "获取妖盟动态"
}
		
func (*RspUnionTrends) MsgId() int32 {
	return 2115
}

func (*RspUnionTrends) Comment() string {
	return "获取妖盟动态"
}
		
func (*ReqUnionModify) MsgId() int32 {
	return 22116
}

func (*ReqUnionModify) Comment() string {
	return "妖盟信息修改"
}
		
func (*RspUnionModify) MsgId() int32 {
	return 2116
}

func (*RspUnionModify) Comment() string {
	return "妖盟信息修改"
}
		
func (*ReqUnionDailyTask) MsgId() int32 {
	return 22117
}

func (*ReqUnionDailyTask) Comment() string {
	return "获取每日任务列表"
}
		
func (*RspUnionDailyTask) MsgId() int32 {
	return 2117
}

func (*RspUnionDailyTask) Comment() string {
	return "获取每日任务列表"
}
		
func (*ReqUnionGetDailyTask) MsgId() int32 {
	return 22118
}

func (*ReqUnionGetDailyTask) Comment() string {
	return "领取每日任务奖励"
}
		
func (*RspUnionGetDailyTask) MsgId() int32 {
	return 2118
}

func (*RspUnionGetDailyTask) Comment() string {
	return "领取每日任务奖励"
}
		
func (*ReqUnionDailyDonate) MsgId() int32 {
	return 22120
}

func (*ReqUnionDailyDonate) Comment() string {
	return "每日任务捐献"
}
		
func (*RspUnionDailyDonate) MsgId() int32 {
	return 2120
}

func (*RspUnionDailyDonate) Comment() string {
	return "每日任务捐献"
}
		
func (*ReqUnionDailyProgress) MsgId() int32 {
	return 22121
}

func (*ReqUnionDailyProgress) Comment() string {
	return "获取每日任务成员完成进度"
}
		
func (*RspUnionDailyProgress) MsgId() int32 {
	return 2121
}

func (*RspUnionDailyProgress) Comment() string {
	return "获取每日任务成员完成进度"
}
		
func (*ReqUnionEviction) MsgId() int32 {
	return 22122
}

func (*ReqUnionEviction) Comment() string {
	return "逐出妖盟"
}
		
func (*RspUnionEviction) MsgId() int32 {
	return 2122
}

func (*RspUnionEviction) Comment() string {
	return "逐出妖盟"
}
		
func (*ReqUnionJoinList) MsgId() int32 {
	return 22123
}

func (*ReqUnionJoinList) Comment() string {
	return "加入的妖盟列表(我没有妖盟时的列表)"
}
		
func (*RspUnionJoinList) MsgId() int32 {
	return 2123
}

func (*RspUnionJoinList) Comment() string {
	return "加入的妖盟列表(我没有妖盟时的列表)"
}
		
func (*GetUnionHelpDataListReq) MsgId() int32 {
	return 22125
}

func (*GetUnionHelpDataListReq) Comment() string {
	return "获取妖盟协助列表"
}
		
func (*GetUnionHelpDataListResp) MsgId() int32 {
	return 2125
}

func (*GetUnionHelpDataListResp) Comment() string {
	return "获取妖盟协助列表"
}
		
func (*RequestUnionHelpReq) MsgId() int32 {
	return 22126
}

func (*RequestUnionHelpReq) Comment() string {
	return "请求妖盟协助"
}
		
func (*RequestUnionHelpResp) MsgId() int32 {
	return 2126
}

func (*RequestUnionHelpResp) Comment() string {
	return "请求妖盟协助"
}
		
func (*UnionHelpReq) MsgId() int32 {
	return 22127
}

func (*UnionHelpReq) Comment() string {
	return "执行妖盟协助"
}
		
func (*UnionHelpResp) MsgId() int32 {
	return 2127
}

func (*UnionHelpResp) Comment() string {
	return "执行妖盟协助"
}
		
func (*GetUnionHelpStateReq) MsgId() int32 {
	return 22128
}

func (*GetUnionHelpStateReq) Comment() string {
	return "获取妖盟协助状态"
}
		
func (*GetUnionHelpStateResp) MsgId() int32 {
	return 2128
}

func (*GetUnionHelpStateResp) Comment() string {
	return "获取妖盟协助状态"
}
		
func (*ReqUnionWatchRedDot) MsgId() int32 {
	return 22129
}

func (*ReqUnionWatchRedDot) Comment() string {
	return "查看妖盟的红点"
}
		
func (*RspUnionWatchRedDot) MsgId() int32 {
	return 2129
}

func (*RspUnionWatchRedDot) Comment() string {
	return "查看妖盟的红点"
}
		
func (*CutPriceBargainReqMsg) MsgId() int32 {
	return 22166
}

func (*CutPriceBargainReqMsg) Comment() string {
	return "砍价"
}
		
func (*CutPriceBargainRespMsg) MsgId() int32 {
	return 2166
}

func (*CutPriceBargainRespMsg) Comment() string {
	return "砍价"
}
		
func (*CutPriceBuyReqMsg) MsgId() int32 {
	return 22167
}

func (*CutPriceBuyReqMsg) Comment() string {
	return "砍价购买"
}
		
func (*CutPriceBuyRespMsg) MsgId() int32 {
	return 2167
}

func (*CutPriceBuyRespMsg) Comment() string {
	return "砍价购买"
}
		
func (*ReqBindInviteCode) MsgId() int32 {
	return 22401
}

func (*ReqBindInviteCode) Comment() string {
	return "绑定邀请码"
}
		
func (*RspBindInviteCode) MsgId() int32 {
	return 2401
}

func (*RspBindInviteCode) Comment() string {
	return "绑定邀请码"
}
		
func (*ReqGetBindInviteCode) MsgId() int32 {
	return 22402
}

func (*ReqGetBindInviteCode) Comment() string {
	return "获取绑定他人的邀请码"
}
		
func (*RspGetBindInviteCode) MsgId() int32 {
	return 2402
}

func (*RspGetBindInviteCode) Comment() string {
	return "获取绑定他人的邀请码"
}
		
func (*ReqHeroRankFightPlayerList) MsgId() int32 {
	return 23702
}

func (*ReqHeroRankFightPlayerList) Comment() string {
	return "请求挑战者列表"
}
		
func (*RspHeroRankFightPlayerList) MsgId() int32 {
	return 3702
}

func (*RspHeroRankFightPlayerList) Comment() string {
	return "请求挑战者列表"
}
		
func (*ReqHeroRankFight) MsgId() int32 {
	return 23703
}

func (*ReqHeroRankFight) Comment() string {
	return "请求挑战玩家"
}
		
func (*RspHeroRankFight) MsgId() int32 {
	return 3703
}

func (*RspHeroRankFight) Comment() string {
	return "请求挑战玩家"
}
		
func (*ReqHeroRankClear) MsgId() int32 {
	return 23704
}

func (*ReqHeroRankClear) Comment() string {
	return "请求扫荡玩家"
}
		
func (*RspHeroRankClear) MsgId() int32 {
	return 3704
}

func (*RspHeroRankClear) Comment() string {
	return "请求扫荡玩家"
}
		
func (*ReqHeroRankGetAchieve) MsgId() int32 {
	return 23706
}

func (*ReqHeroRankGetAchieve) Comment() string {
	return "请求领取成就奖励"
}
		
func (*RspHeroRankGetAchieve) MsgId() int32 {
	return 3706
}

func (*RspHeroRankGetAchieve) Comment() string {
	return "请求领取成就奖励"
}
		
func (*ReqHeroRankBuyEnergy) MsgId() int32 {
	return 23707
}

func (*ReqHeroRankBuyEnergy) Comment() string {
	return "请求购买体力"
}
		
func (*RspHeroRankBuyEnergy) MsgId() int32 {
	return 3707
}

func (*RspHeroRankBuyEnergy) Comment() string {
	return "请求购买体力"
}
		
func (*ReqPetDreamLandDraw) MsgId() int32 {
	return 24101
}

func (*ReqPetDreamLandDraw) Comment() string {
	return "抽奖"
}
		
func (*RspPetDreamLandDraw) MsgId() int32 {
	return 4101
}

func (*RspPetDreamLandDraw) Comment() string {
	return "抽奖"
}
		
func (*ReqPetDreamLandAdventurePlaceUnlock) MsgId() int32 {
	return 24102
}

func (*ReqPetDreamLandAdventurePlaceUnlock) Comment() string {
	return "区域解锁"
}
		
func (*RspPetDreamLandAdventurePlaceUnlock) MsgId() int32 {
	return 4102
}

func (*RspPetDreamLandAdventurePlaceUnlock) Comment() string {
	return "区域解锁"
}
		
func (*ReqPetDreamLandAdventureSlotUnlock) MsgId() int32 {
	return 24103
}

func (*ReqPetDreamLandAdventureSlotUnlock) Comment() string {
	return "坑位解锁"
}
		
func (*RspPetDreamLandAdventurePlaceSlot) MsgId() int32 {
	return 4103
}

func (*RspPetDreamLandAdventurePlaceSlot) Comment() string {
	return "坑位解锁"
}
		
func (*ReqPetDreamLandAdventureDispatch) MsgId() int32 {
	return 24104
}

func (*ReqPetDreamLandAdventureDispatch) Comment() string {
	return "灵兽派遣"
}
		
func (*RspPetDreamLandAdventureDispatch) MsgId() int32 {
	return 4104
}

func (*RspPetDreamLandAdventureDispatch) Comment() string {
	return "灵兽派遣"
}
		
func (*ReqPetDreamLandAdventureGetAward) MsgId() int32 {
	return 24105
}

func (*ReqPetDreamLandAdventureGetAward) Comment() string {
	return "领取收益"
}
		
func (*RspPetDreamLandAdventureGetAward) MsgId() int32 {
	return 4105
}

func (*RspPetDreamLandAdventureGetAward) Comment() string {
	return "领取收益"
}
		
func (*RefreshEvilThingReq) MsgId() int32 {
	return 24201
}

func (*RefreshEvilThingReq) Comment() string {
	return "刷新邪祟"
}
		
func (*RefreshEvilThingResp) MsgId() int32 {
	return 4201
}

func (*RefreshEvilThingResp) Comment() string {
	return "刷新邪祟"
}
		
func (*BattleEviThingReq) MsgId() int32 {
	return 24202
}

func (*BattleEviThingReq) Comment() string {
	return "镇压邪祟"
}
		
func (*BattleEvilThingResp) MsgId() int32 {
	return 4202
}

func (*BattleEvilThingResp) Comment() string {
	return "镇压邪祟"
}
		
func (*MagicEquipReq) MsgId() int32 {
	return 24401
}

func (*MagicEquipReq) Comment() string {
	return "装备神通"
}
		
func (*MagicEquipResp) MsgId() int32 {
	return 4401
}

func (*MagicEquipResp) Comment() string {
	return "装备神通"
}
		
func (*MagicResetReq) MsgId() int32 {
	return 24402
}

func (*MagicResetReq) Comment() string {
	return "重置神通"
}
		
func (*MagicResetResp) MsgId() int32 {
	return 4402
}

func (*MagicResetResp) Comment() string {
	return "重置神通"
}
		
func (*MagicStageUpReq) MsgId() int32 {
	return 24403
}

func (*MagicStageUpReq) Comment() string {
	return "进阶神通"
}
		
func (*MagicStageUpResp) MsgId() int32 {
	return 4403
}

func (*MagicStageUpResp) Comment() string {
	return "进阶神通"
}
		
func (*MagicLvUpReq) MsgId() int32 {
	return 24404
}

func (*MagicLvUpReq) Comment() string {
	return "升级神通"
}
		
func (*MagicLvUpResp) MsgId() int32 {
	return 4404
}

func (*MagicLvUpResp) Comment() string {
	return "升级神通"
}
		
func (*MagicEquipMarkReq) MsgId() int32 {
	return 24405
}

func (*MagicEquipMarkReq) Comment() string {
	return "装备印记"
}
		
func (*MagicEquipMarkResp) MsgId() int32 {
	return 4405
}

func (*MagicEquipMarkResp) Comment() string {
	return "装备印记"
}
		
func (*MagicUnsnatchMarkReq) MsgId() int32 {
	return 24406
}

func (*MagicUnsnatchMarkReq) Comment() string {
	return "卸下印记"
}
		
func (*MagicUnsnatchMarkResp) MsgId() int32 {
	return 4406
}

func (*MagicUnsnatchMarkResp) Comment() string {
	return "卸下印记"
}
		
func (*MagicFusionMarkReq) MsgId() int32 {
	return 24407
}

func (*MagicFusionMarkReq) Comment() string {
	return "融合神通印记"
}
		
func (*MagicFusionMarkResp) MsgId() int32 {
	return 4407
}

func (*MagicFusionMarkResp) Comment() string {
	return "融合神通印记"
}
		
func (*MagicDerivationReq) MsgId() int32 {
	return 24408
}

func (*MagicDerivationReq) Comment() string {
	return "衍化"
}
		
func (*MagicDerivationResp) MsgId() int32 {
	return 4408
}

func (*MagicDerivationResp) Comment() string {
	return "衍化"
}
		
func (*MagicCombineLvUpReq) MsgId() int32 {
	return 24409
}

func (*MagicCombineLvUpReq) Comment() string {
	return "升级组合"
}
		
func (*MagicCombineLvUpResp) MsgId() int32 {
	return 4409
}

func (*MagicCombineLvUpResp) Comment() string {
	return "升级组合"
}
		
func (*MagicSwitchPresetsReq) MsgId() int32 {
	return 24410
}

func (*MagicSwitchPresetsReq) Comment() string {
	return "切换预设"
}
		
func (*MagicSwitchPresetsResp) MsgId() int32 {
	return 4410
}

func (*MagicSwitchPresetsResp) Comment() string {
	return "切换预设"
}
		
func (*MagicStageUpAllReq) MsgId() int32 {
	return 24411
}

func (*MagicStageUpAllReq) Comment() string {
	return "一键突破神通"
}
		
func (*GetPresetMagicInfoReq) MsgId() int32 {
	return 24412
}

func (*GetPresetMagicInfoReq) Comment() string {
	return "获取预设神通列表"
}
		
func (*GetPresetMagicInfoResp) MsgId() int32 {
	return 4412
}

func (*GetPresetMagicInfoResp) Comment() string {
	return "获取预设神通列表"
}
		
func (*UnEquipStoneReq) MsgId() int32 {
	return 24413
}

func (*UnEquipStoneReq) Comment() string {
	return "拆卸装备的印记"
}
		
func (*UnEquipStoneResp) MsgId() int32 {
	return 4413
}

func (*UnEquipStoneResp) Comment() string {
	return "拆卸装备的印记"
}
		
func (*SaveToServiceReq) MsgId() int32 {
	return 24501
}

func (*SaveToServiceReq) Comment() string {
	return "添加储存数据"
}
		
func (*SaveToServiceRsp) MsgId() int32 {
	return 4501
}

func (*SaveToServiceRsp) Comment() string {
	return "添加储存数据"
}
		
func (*PalaceWorshipReq) MsgId() int32 {
	return 24802
}

func (*PalaceWorshipReq) Comment() string {
	return "仙宫点赞同步"
}
		
func (*PalaceWorshipRsp) MsgId() int32 {
	return 4802
}

func (*PalaceWorshipRsp) Comment() string {
	return "仙宫点赞同步"
}
		
func (*EnterPalaceReq) MsgId() int32 {
	return 24803
}

func (*EnterPalaceReq) Comment() string {
	return "仙宫外部数据请求"
}
		
func (*EnterPalaceRsp) MsgId() int32 {
	return 4803
}

func (*EnterPalaceRsp) Comment() string {
	return "仙宫外部数据请求"
}
		
func (*EnterPalaceInnerReq) MsgId() int32 {
	return 24804
}

func (*EnterPalaceInnerReq) Comment() string {
	return "仙宫内部数据请求"
}
		
func (*EnterPalaceInnerRsp) MsgId() int32 {
	return 4804
}

func (*EnterPalaceInnerRsp) Comment() string {
	return "仙宫内部数据请求"
}
		
func (*PalaceInnerBookReq) MsgId() int32 {
	return 24805
}

func (*PalaceInnerBookReq) Comment() string {
	return "仙名录数据请求"
}
		
func (*PalaceInnerBookRsp) MsgId() int32 {
	return 4805
}

func (*PalaceInnerBookRsp) Comment() string {
	return "仙名录数据请求"
}
		
func (*SendGiftReq) MsgId() int32 {
	return 24806
}

func (*SendGiftReq) Comment() string {
	return "执行仙宫送福"
}
		
func (*SendGiftRsp) MsgId() int32 {
	return 4806
}

func (*SendGiftRsp) Comment() string {
	return "执行仙宫送福"
}
		
func (*GetSendGiftRewardReq) MsgId() int32 {
	return 24807
}

func (*GetSendGiftRewardReq) Comment() string {
	return "仙宫送福领奖"
}
		
func (*GetSendGiftRewardRsp) MsgId() int32 {
	return 4807
}

func (*GetSendGiftRewardRsp) Comment() string {
	return "仙宫送福领奖"
}
		
func (*GetTitleIdListReq) MsgId() int32 {
	return 24810
}

func (*GetTitleIdListReq) Comment() string {
	return "获得当前已开启的天宫称号id"
}
		
func (*GetTitleIdListRsp) MsgId() int32 {
	return 4810
}

func (*GetTitleIdListRsp) Comment() string {
	return "获得当前已开启的天宫称号id"
}
		
func (*PalaceAchievementGetRewardReq) MsgId() int32 {
	return 24857
}

func (*PalaceAchievementGetRewardReq) Comment() string {
	return "仙宫获得成就奖励"
}
		
func (*PalaceAchievementGetRewardRsp) MsgId() int32 {
	return 4857
}

func (*PalaceAchievementGetRewardRsp) Comment() string {
	return "仙宫获得成就奖励"
}
		
func (*ManHuangEnterPanelReqMsg) MsgId() int32 {
	return 25001
}

func (*ManHuangEnterPanelReqMsg) Comment() string {
	return "蛮荒 进入活动页"
}
		
func (*ManHuangEnterPanelRespMsg) MsgId() int32 {
	return 5001
}

func (*ManHuangEnterPanelRespMsg) Comment() string {
	return "蛮荒 进入活动页"
}
		
func (*ManHuangCreateTeamReqMsg) MsgId() int32 {
	return 25002
}

func (*ManHuangCreateTeamReqMsg) Comment() string {
	return "蛮荒 创建队伍"
}
		
func (*ManHuangCreateTeamRespMsg) MsgId() int32 {
	return 5002
}

func (*ManHuangCreateTeamRespMsg) Comment() string {
	return "蛮荒 创建队伍"
}
		
func (*ManHuangGetTeamListReqMsg) MsgId() int32 {
	return 25003
}

func (*ManHuangGetTeamListReqMsg) Comment() string {
	return "蛮荒 获取队伍列表"
}
		
func (*ManHuangGetTeamListRespMsg) MsgId() int32 {
	return 5003
}

func (*ManHuangGetTeamListRespMsg) Comment() string {
	return "蛮荒 获取队伍列表"
}
		
func (*ManHuangGetTeamInfoReqMsg) MsgId() int32 {
	return 25004
}

func (*ManHuangGetTeamInfoReqMsg) Comment() string {
	return "蛮荒 查看队伍数据"
}
		
func (*ManHuangGetTeamInfoRespMsg) MsgId() int32 {
	return 5004
}

func (*ManHuangGetTeamInfoRespMsg) Comment() string {
	return "蛮荒 查看队伍数据"
}
		
func (*ManHuangJoinTeamReqMsg) MsgId() int32 {
	return 25005
}

func (*ManHuangJoinTeamReqMsg) Comment() string {
	return "蛮荒 加入队伍"
}
		
func (*ManHuangJoinTeamRespMsg) MsgId() int32 {
	return 5005
}

func (*ManHuangJoinTeamRespMsg) Comment() string {
	return "蛮荒 加入队伍"
}
		
func (*ManHuangCancelTeamApplyReqMsg) MsgId() int32 {
	return 25006
}

func (*ManHuangCancelTeamApplyReqMsg) Comment() string {
	return "蛮荒 取消申请队伍"
}
		
func (*ManHuangCancelTeamApplyRespMsg) MsgId() int32 {
	return 5006
}

func (*ManHuangCancelTeamApplyRespMsg) Comment() string {
	return "蛮荒 取消申请队伍"
}
		
func (*ManHuangQuitTeamReqMsg) MsgId() int32 {
	return 25007
}

func (*ManHuangQuitTeamReqMsg) Comment() string {
	return "蛮荒 解散/退出队伍"
}
		
func (*ManHuangQuitTeamRespMsg) MsgId() int32 {
	return 5007
}

func (*ManHuangQuitTeamRespMsg) Comment() string {
	return "蛮荒 解散/退出队伍"
}
		
func (*ManHuangLeavePanelReqMsg) MsgId() int32 {
	return 25008
}

func (*ManHuangLeavePanelReqMsg) Comment() string {
	return "蛮荒 退出活动"
}
		
func (*ManHuangLeavePanelRespMsg) MsgId() int32 {
	return 5008
}

func (*ManHuangLeavePanelRespMsg) Comment() string {
	return "蛮荒 退出活动"
}
		
func (*ManHuangApplyJoinTeamAgreeReqMsg) MsgId() int32 {
	return 25009
}

func (*ManHuangApplyJoinTeamAgreeReqMsg) Comment() string {
	return "蛮荒 玩家申请入队同意"
}
		
func (*ManHuangApplyJoinTeamAgreeRespMsg) MsgId() int32 {
	return 5009
}

func (*ManHuangApplyJoinTeamAgreeRespMsg) Comment() string {
	return "蛮荒 玩家申请入队同意"
}
		
func (*ManHuangApplyJoinTeamRefusedReqMsg) MsgId() int32 {
	return 25010
}

func (*ManHuangApplyJoinTeamRefusedReqMsg) Comment() string {
	return "蛮荒 玩家申请入队一键拒绝"
}
		
func (*ManHuangApplyJoinTeamRefusedRespMsg) MsgId() int32 {
	return 5010
}

func (*ManHuangApplyJoinTeamRefusedRespMsg) Comment() string {
	return "蛮荒 玩家申请入队一键拒绝"
}
		
func (*ManHuangKickOutTeamReqMsg) MsgId() int32 {
	return 25011
}

func (*ManHuangKickOutTeamReqMsg) Comment() string {
	return "蛮荒 踢出队伍"
}
		
func (*ManHuangKickOutTeamRespMsg) MsgId() int32 {
	return 5011
}

func (*ManHuangKickOutTeamRespMsg) Comment() string {
	return "蛮荒 踢出队伍"
}
		
func (*ManHuangEnterRegionReqMsg) MsgId() int32 {
	return 25012
}

func (*ManHuangEnterRegionReqMsg) Comment() string {
	return "蛮荒 进入区域"
}
		
func (*ManHuangEnterRegionRespMsg) MsgId() int32 {
	return 5012
}

func (*ManHuangEnterRegionRespMsg) Comment() string {
	return "蛮荒 进入区域"
}
		
func (*ManHuangExploreReqMsg) MsgId() int32 {
	return 25013
}

func (*ManHuangExploreReqMsg) Comment() string {
	return "蛮荒 探索"
}
		
func (*ManHuangExploreRespMsg) MsgId() int32 {
	return 5013
}

func (*ManHuangExploreRespMsg) Comment() string {
	return "蛮荒 探索"
}
		
func (*ManHuangEventHandleReqMsg) MsgId() int32 {
	return 25014
}

func (*ManHuangEventHandleReqMsg) Comment() string {
	return "蛮荒 执行探索事件 "
}
		
func (*ManHuangEventHandleRespMsg) MsgId() int32 {
	return 5014
}

func (*ManHuangEventHandleRespMsg) Comment() string {
	return "蛮荒 执行探索事件 "
}
		
func (*ManHuangEventActionReqMsg) MsgId() int32 {
	return 25015
}

func (*ManHuangEventActionReqMsg) Comment() string {
	return "蛮荒 action探索事件(标记，绕过) "
}
		
func (*ManHuangEventActionRespMsg) MsgId() int32 {
	return 5015
}

func (*ManHuangEventActionRespMsg) Comment() string {
	return "蛮荒 action探索事件(标记，绕过) "
}
		
func (*ManHuangLogHelpReqMsg) MsgId() int32 {
	return 25016
}

func (*ManHuangLogHelpReqMsg) Comment() string {
	return "蛮荒 日志-队友互助 "
}
		
func (*ManHuangLogHelpRespMsg) MsgId() int32 {
	return 5016
}

func (*ManHuangLogHelpRespMsg) Comment() string {
	return "蛮荒 日志-队友互助 "
}
		
func (*ManHuangLogBattlegroundReqMsg) MsgId() int32 {
	return 25017
}

func (*ManHuangLogBattlegroundReqMsg) Comment() string {
	return "蛮荒 日志-战场日志 "
}
		
func (*ManHuangLogBattlegroundRespMsg) MsgId() int32 {
	return 5017
}

func (*ManHuangLogBattlegroundRespMsg) Comment() string {
	return "蛮荒 日志-战场日志 "
}
		
func (*ManHuangLogPersonReqMsg) MsgId() int32 {
	return 25018
}

func (*ManHuangLogPersonReqMsg) Comment() string {
	return "蛮荒 日志-个人日志 "
}
		
func (*ManHuangLogPersonRespMsg) MsgId() int32 {
	return 5018
}

func (*ManHuangLogPersonRespMsg) Comment() string {
	return "蛮荒 日志-个人日志 "
}
		
func (*ManHuangRankPersonReqMsg) MsgId() int32 {
	return 25019
}

func (*ManHuangRankPersonReqMsg) Comment() string {
	return "蛮荒 排行-个人 "
}
		
func (*ManHuangRankPersonRespMsg) MsgId() int32 {
	return 5019
}

func (*ManHuangRankPersonRespMsg) Comment() string {
	return "蛮荒 排行-个人 "
}
		
func (*ManHuangRankTeamReqMsg) MsgId() int32 {
	return 25020
}

func (*ManHuangRankTeamReqMsg) Comment() string {
	return "蛮荒 排行-队伍 "
}
		
func (*ManHuangRankTeamRespMsg) MsgId() int32 {
	return 5020
}

func (*ManHuangRankTeamRespMsg) Comment() string {
	return "蛮荒 排行-队伍 "
}
		
func (*ManHuangRecoverEnergyReqMsg) MsgId() int32 {
	return 25025
}

func (*ManHuangRecoverEnergyReqMsg) Comment() string {
	return "蛮荒 恢复体力"
}
		
func (*ManHuangRecoverEnergyRespMsg) MsgId() int32 {
	return 5025
}

func (*ManHuangRecoverEnergyRespMsg) Comment() string {
	return "蛮荒 恢复体力"
}
		
func (*ManHuangHelpActionAttReqMsg) MsgId() int32 {
	return 25026
}

func (*ManHuangHelpActionAttReqMsg) Comment() string {
	return "蛮荒 执行协助"
}
		
func (*ManHuangHelpActionAttRespMsg) MsgId() int32 {
	return 5026
}

func (*ManHuangHelpActionAttRespMsg) Comment() string {
	return "蛮荒 执行协助"
}
		
func (*ManHuangOpenBoxReqMsg) MsgId() int32 {
	return 25027
}

func (*ManHuangOpenBoxReqMsg) Comment() string {
	return "蛮荒 开宝箱"
}
		
func (*ManHuangOpenBoxRespMsg) MsgId() int32 {
	return 5027
}

func (*ManHuangOpenBoxRespMsg) Comment() string {
	return "蛮荒 开宝箱"
}
		
func (*ManHuangHelpActionRewardReqMsg) MsgId() int32 {
	return 25028
}

func (*ManHuangHelpActionRewardReqMsg) Comment() string {
	return "蛮荒 领取协助奖励"
}
		
func (*ManHuangHelpActionRewardRespMsg) MsgId() int32 {
	return 5028
}

func (*ManHuangHelpActionRewardRespMsg) Comment() string {
	return "蛮荒 领取协助奖励"
}
		
func (*ManHuangGetUserDataReqMsg) MsgId() int32 {
	return 25030
}

func (*ManHuangGetUserDataReqMsg) Comment() string {
	return "蛮荒 发送同步玩家信息协议"
}
		
func (*ManHuangGetUserDataRespMsg) MsgId() int32 {
	return 5030
}

func (*ManHuangGetUserDataRespMsg) Comment() string {
	return "蛮荒 发送同步玩家信息协议"
}
		
func (*ReqEquipmentRefine) MsgId() int32 {
	return 25501
}

func (*ReqEquipmentRefine) Comment() string {
	return "精炼"
}
		
func (*RespEquipmentRefine) MsgId() int32 {
	return 5501
}

func (*RespEquipmentRefine) Comment() string {
	return "精炼"
}
		
func (*ReqEquipmentAdvance) MsgId() int32 {
	return 25502
}

func (*ReqEquipmentAdvance) Comment() string {
	return "进阶"
}
		
func (*RespEquipmentAdvance) MsgId() int32 {
	return 5502
}

func (*RespEquipmentAdvance) Comment() string {
	return "进阶"
}
		
func (*ReqEquipmentActivate) MsgId() int32 {
	return 25503
}

func (*ReqEquipmentActivate) Comment() string {
	return "技能解锁"
}
		
func (*RespEquipmentActivate) MsgId() int32 {
	return 5503
}

func (*RespEquipmentActivate) Comment() string {
	return "技能解锁"
}
		
func (*SecretTowerEnterReq) MsgId() int32 {
	return 25601
}

func (*SecretTowerEnterReq) Comment() string {
	return "进入秘境"
}
		
func (*SecretTowerEnterRsp) MsgId() int32 {
	return 5601
}

func (*SecretTowerEnterRsp) Comment() string {
	return "进入秘境"
}
		
func (*SecretTowerFightReq) MsgId() int32 {
	return 25602
}

func (*SecretTowerFightReq) Comment() string {
	return "挑战怪物"
}
		
func (*SecretTowerFightResp) MsgId() int32 {
	return 5602
}

func (*SecretTowerFightResp) Comment() string {
	return "挑战怪物"
}
		
func (*SecretTowerGetStageRewardReq) MsgId() int32 {
	return 25603
}

func (*SecretTowerGetStageRewardReq) Comment() string {
	return "领取阶段奖励"
}
		
func (*SecretTowerGetStageRewardRsp) MsgId() int32 {
	return 5603
}

func (*SecretTowerGetStageRewardRsp) Comment() string {
	return "领取阶段奖励"
}
		
func (*SecretTowerViewMonsterAttrReq) MsgId() int32 {
	return 25604
}

func (*SecretTowerViewMonsterAttrReq) Comment() string {
	return "查看怪物属性"
}
		
func (*SecretTowerViewMonsterAttrResp) MsgId() int32 {
	return 5604
}

func (*SecretTowerViewMonsterAttrResp) Comment() string {
	return "查看怪物属性"
}
		
func (*SecretTowerAchievementReq) MsgId() int32 {
	return 25606
}

func (*SecretTowerAchievementReq) Comment() string {
	return "查看秘境成就"
}
		
func (*SecretTowerAchievementRsp) MsgId() int32 {
	return 5606
}

func (*SecretTowerAchievementRsp) Comment() string {
	return "查看秘境成就"
}
		
func (*SecretTowerGetAchiRewardReq) MsgId() int32 {
	return 25607
}

func (*SecretTowerGetAchiRewardReq) Comment() string {
	return "领取成就奖励"
}
		
func (*SecretTowerGetAchiRewardRsp) MsgId() int32 {
	return 5607
}

func (*SecretTowerGetAchiRewardRsp) Comment() string {
	return "领取成就奖励"
}
		
func (*UnionBossInfoReqMsg) MsgId() int32 {
	return 25801
}

func (*UnionBossInfoReqMsg) Comment() string {
	return "进入妖邪讨伐"
}
		
func (*UnionBossInfoRespMsg) MsgId() int32 {
	return 5801
}

func (*UnionBossInfoRespMsg) Comment() string {
	return "进入妖邪讨伐"
}
		
func (*UnionBossBuffReqMsg) MsgId() int32 {
	return 25802
}

func (*UnionBossBuffReqMsg) Comment() string {
	return "布阵"
}
		
func (*UnionBossBuffRespMsg) MsgId() int32 {
	return 5802
}

func (*UnionBossBuffRespMsg) Comment() string {
	return "布阵"
}
		
func (*UnionBossRewardReqMsg) MsgId() int32 {
	return 25803
}

func (*UnionBossRewardReqMsg) Comment() string {
	return "查看讨伐奖励"
}
		
func (*UnionBossRewardRespMsg) MsgId() int32 {
	return 5803
}

func (*UnionBossRewardRespMsg) Comment() string {
	return "查看讨伐奖励"
}
		
func (*UnionBossRewardReceiveReqMsg) MsgId() int32 {
	return 25804
}

func (*UnionBossRewardReceiveReqMsg) Comment() string {
	return "领取讨伐奖励"
}
		
func (*UnionBossRewardReceiveRespMsg) MsgId() int32 {
	return 5804
}

func (*UnionBossRewardReceiveRespMsg) Comment() string {
	return "领取讨伐奖励"
}
		
func (*UnionBossBattleReqMsg) MsgId() int32 {
	return 25805
}

func (*UnionBossBattleReqMsg) Comment() string {
	return "挑战"
}
		
func (*UnionBossBattleRespMsg) MsgId() int32 {
	return 5805
}

func (*UnionBossBattleRespMsg) Comment() string {
	return "挑战"
}
		
func (*UnionBossAchieveRewardReqMsg) MsgId() int32 {
	return 25806
}

func (*UnionBossAchieveRewardReqMsg) Comment() string {
	return "领取讨伐成就奖励"
}
		
func (*UnionBossAchieveRewardRespMsg) MsgId() int32 {
	return 5806
}

func (*UnionBossAchieveRewardRespMsg) Comment() string {
	return "领取讨伐成就奖励"
}
		
func (*UnionBossAddBuffPlayerReqMsg) MsgId() int32 {
	return 25809
}

func (*UnionBossAddBuffPlayerReqMsg) Comment() string {
	return "请求已布阵玩家信息"
}
		
func (*UnionBossAddBuffPlayerRespMsg) MsgId() int32 {
	return 5809
}

func (*UnionBossAddBuffPlayerRespMsg) Comment() string {
	return "请求已布阵玩家信息"
}
		
func (*MagicTreasureDrawReq) MsgId() int32 {
	return 26302
}

func (*MagicTreasureDrawReq) Comment() string {
	return "法宝抽取"
}
		
func (*MagicTreasureDrawRsp) MsgId() int32 {
	return 6302
}

func (*MagicTreasureDrawRsp) Comment() string {
	return "法宝抽取"
}
		
func (*MagicTreasureLvUpReq) MsgId() int32 {
	return 26303
}

func (*MagicTreasureLvUpReq) Comment() string {
	return "法宝升级"
}
		
func (*MagicTreasureLvUpRsp) MsgId() int32 {
	return 6303
}

func (*MagicTreasureLvUpRsp) Comment() string {
	return "法宝升级"
}
		
func (*MagicTreasureActiveReq) MsgId() int32 {
	return 26304
}

func (*MagicTreasureActiveReq) Comment() string {
	return "法宝激活"
}
		
func (*MagicTreasureActiveRsp) MsgId() int32 {
	return 6304
}

func (*MagicTreasureActiveRsp) Comment() string {
	return "法宝激活"
}
		
func (*MagicTreasureStarUpReq) MsgId() int32 {
	return 26305
}

func (*MagicTreasureStarUpReq) Comment() string {
	return "法宝升星"
}
		
func (*MagicTreasureStarUpRsp) MsgId() int32 {
	return 6305
}

func (*MagicTreasureStarUpRsp) Comment() string {
	return "法宝升星"
}
		
func (*MagicTreasureSwitchLinkageSkinReq) MsgId() int32 {
	return 26307
}

func (*MagicTreasureSwitchLinkageSkinReq) Comment() string {
	return "法宝形象切换"
}
		
func (*MagicTreasureSwitchLinkageSkinResp) MsgId() int32 {
	return 6307
}

func (*MagicTreasureSwitchLinkageSkinResp) Comment() string {
	return "法宝形象切换"
}
		
func (*ReceiveMiniGamesRewardReq) MsgId() int32 {
	return 200021
}

func (*ReceiveMiniGamesRewardReq) Comment() string {
	return "小游戏领奖"
}
		
func (*ReceiveMiniGamesRewardResp) MsgId() int32 {
	return 21
}

func (*ReceiveMiniGamesRewardResp) Comment() string {
	return "小游戏领奖"
}
		
func (*StageMapChallengeReq) MsgId() int32 {
	return 200022
}

func (*StageMapChallengeReq) Comment() string {
	return "秘境挑战"
}
		
func (*StageMapChallengeRsp) MsgId() int32 {
	return 22
}

func (*StageMapChallengeRsp) Comment() string {
	return "秘境挑战"
}
		
func (*ForbiddenTrailsGetRewardReq) MsgId() int32 {
	return 200023
}

func (*ForbiddenTrailsGetRewardReq) Comment() string {
	return "禁地试炼领奖"
}
		
func (*ForbiddenTrailsGetRewardResp) MsgId() int32 {
	return 23
}

func (*ForbiddenTrailsGetRewardResp) Comment() string {
	return "禁地试炼领奖"
}
		
func (*EnterWestJourneyReq) MsgId() int32 {
	return 200024
}

func (*EnterWestJourneyReq) Comment() string {
	return "西天救援进入"
}
		
func (*EnterWestJourneyResp) MsgId() int32 {
	return 24
}

func (*EnterWestJourneyResp) Comment() string {
	return "西天救援进入"
}
		
func (*GetWestJourneyRewardReq) MsgId() int32 {
	return 200026
}

func (*GetWestJourneyRewardReq) Comment() string {
	return "西天救援领奖"
}
		
func (*GetWestJourneyRewardResp) MsgId() int32 {
	return 26
}

func (*GetWestJourneyRewardResp) Comment() string {
	return "西天救援领奖"
}
		
func (*UnionPlayerGradeRankReqMsg) MsgId() int32 {
	return 201630
}

func (*UnionPlayerGradeRankReqMsg) Comment() string {
	return "获取分组评级排行"
}
		
func (*UnionPlayerGradeRankRespMsg) MsgId() int32 {
	return 1630
}

func (*UnionPlayerGradeRankRespMsg) Comment() string {
	return "获取分组评级排行"
}
		
func (*UnionMemberGradeSettingReq) MsgId() int32 {
	return 201650
}

func (*UnionMemberGradeSettingReq) Comment() string {
	return "评级人数分配设置"
}
		
func (*UnionMemberGradeSettingResp) MsgId() int32 {
	return 1650
}

func (*UnionMemberGradeSettingResp) Comment() string {
	return "评级人数分配设置"
}
		
func (*GetCrossUnionGroupServersReq) MsgId() int32 {
	return 201651
}

func (*GetCrossUnionGroupServersReq) Comment() string {
	return "获取妖盟跨服信息"
}
		
func (*GetCrossUnionGroupServersResp) MsgId() int32 {
	return 1651
}

func (*GetCrossUnionGroupServersResp) Comment() string {
	return "获取妖盟跨服信息"
}
		
func (*UnionPublishRecruitReqMsg) MsgId() int32 {
	return 201654
}

func (*UnionPublishRecruitReqMsg) Comment() string {
	return "发布招募信息"
}
		
func (*UnionPublishRecruitRespMsg) MsgId() int32 {
	return 1654
}

func (*UnionPublishRecruitRespMsg) Comment() string {
	return "发布招募信息"
}
		
func (*RemoveBlackReqMsg) MsgId() int32 {
	return 203205
}

func (*RemoveBlackReqMsg) Comment() string {
	return "移除黑名单"
}
		
func (*RemoveBlackRespMsg) MsgId() int32 {
	return 3205
}

func (*RemoveBlackRespMsg) Comment() string {
	return "移除黑名单"
}
		
func (*GetBlackPlayerListReqMsg) MsgId() int32 {
	return 203206
}

func (*GetBlackPlayerListReqMsg) Comment() string {
	return "获取黑名单列表"
}
		
func (*GetBlackPlayerListRespMsg) MsgId() int32 {
	return 3206
}

func (*GetBlackPlayerListRespMsg) Comment() string {
	return "获取黑名单列表"
}
		
func (*WorldMessageCheckEffectiveReq) MsgId() int32 {
	return 203208
}

func (*WorldMessageCheckEffectiveReq) Comment() string {
	return "检测消息是否有效"
}
		
func (*WorldMessageCheckEffectiveRsp) MsgId() int32 {
	return 3208
}

func (*WorldMessageCheckEffectiveRsp) Comment() string {
	return "检测消息是否有效"
}
		
func (*UnionBattleEnterReq) MsgId() int32 {
	return 205301
}

func (*UnionBattleEnterReq) Comment() string {
	return "进入妖盟乱斗"
}
		
func (*UnionBattleEnterRsp) MsgId() int32 {
	return 5301
}

func (*UnionBattleEnterRsp) Comment() string {
	return "进入妖盟乱斗"
}
		
func (*UnionBattleGetJoinMemberListReq) MsgId() int32 {
	return 205302
}

func (*UnionBattleGetJoinMemberListReq) Comment() string {
	return "查看自己所属联盟的参与名单"
}
		
func (*UnionBattleGetJoinMemberListRsp) MsgId() int32 {
	return 5302
}

func (*UnionBattleGetJoinMemberListRsp) Comment() string {
	return "查看自己所属联盟的参与名单"
}
		
func (*UnionBattleMatchReq) MsgId() int32 {
	return 205303
}

func (*UnionBattleMatchReq) Comment() string {
	return "匹配"
}
		
func (*UnionBattleMatchRsp) MsgId() int32 {
	return 5303
}

func (*UnionBattleMatchRsp) Comment() string {
	return "匹配"
}
		
func (*UnionBattleChallengeReq) MsgId() int32 {
	return 205304
}

func (*UnionBattleChallengeReq) Comment() string {
	return "挑战"
}
		
func (*UnionBattleChallengeRsp) MsgId() int32 {
	return 5304
}

func (*UnionBattleChallengeRsp) Comment() string {
	return "挑战"
}
		
func (*UnionBattleGetReportReq) MsgId() int32 {
	return 205306
}

func (*UnionBattleGetReportReq) Comment() string {
	return "查看战报信息"
}
		
func (*UnionBattleGetReportRsp) MsgId() int32 {
	return 5306
}

func (*UnionBattleGetReportRsp) Comment() string {
	return "查看战报信息"
}
		
func (*UnionBattleBuyOpenBuffReq) MsgId() int32 {
	return 205307
}

func (*UnionBattleBuyOpenBuffReq) Comment() string {
	return "购买buff"
}
		
func (*UnionBattleBuyOpenBuffRsp) MsgId() int32 {
	return 5307
}

func (*UnionBattleBuyOpenBuffRsp) Comment() string {
	return "购买buff"
}
		
func (*UnionBattleSelectBuffReq) MsgId() int32 {
	return 205308
}

func (*UnionBattleSelectBuffReq) Comment() string {
	return "选择buff"
}
		
func (*UnionBattleSelectBuffRsp) MsgId() int32 {
	return 5308
}

func (*UnionBattleSelectBuffRsp) Comment() string {
	return "选择buff"
}
		
func (*UnionBattleGetPlayBackListReq) MsgId() int32 {
	return 205309
}

func (*UnionBattleGetPlayBackListReq) Comment() string {
	return "获取本次挑战的回放列表"
}
		
func (*UnionBattleGetPlayBackListRsp) MsgId() int32 {
	return 5309
}

func (*UnionBattleGetPlayBackListRsp) Comment() string {
	return "获取本次挑战的回放列表"
}
		
func (*UnionBattleReceiveUnionAchieveRewardReq) MsgId() int32 {
	return 205310
}

func (*UnionBattleReceiveUnionAchieveRewardReq) Comment() string {
	return "领取联盟成就奖励"
}
		
func (*UnionBattleReceiveUnionAchieveRewardRsp) MsgId() int32 {
	return 5310
}

func (*UnionBattleReceiveUnionAchieveRewardRsp) Comment() string {
	return "领取联盟成就奖励"
}
		
func (*UnionBattleGetUnionMemberScoreReq) MsgId() int32 {
	return 205312
}

func (*UnionBattleGetUnionMemberScoreReq) Comment() string {
	return "查看联盟积分"
}
		
func (*UnionBattleGetUnionMemberScoreRsp) MsgId() int32 {
	return 5312
}

func (*UnionBattleGetUnionMemberScoreRsp) Comment() string {
	return "查看联盟积分"
}
		
func (*UnionBattleRedDotReqMsg) MsgId() int32 {
	return 205313
}

func (*UnionBattleRedDotReqMsg) Comment() string {
	return "红点"
}
		
func (*UnionBattleRedDotRespMsg) MsgId() int32 {
	return 5313
}

func (*UnionBattleRedDotRespMsg) Comment() string {
	return "红点"
}
		
func (*AskDingGetChampionInfoReq) MsgId() int32 {
	return 205935
}

func (*AskDingGetChampionInfoReq) Comment() string {
	return "领奖期 冠军弹窗"
}
		
func (*AskDingGetChampionInfoResp) MsgId() int32 {
	return 5935
}

func (*AskDingGetChampionInfoResp) Comment() string {
	return "领奖期 冠军弹窗"
}
		
func (*UnionAreaWarReqMsg) MsgId() int32 {
	return 206001
}

func (*UnionAreaWarReqMsg) Comment() string {
	return "查询据点战基础信息"
}
		
func (*UnionAreaWarRespMsg) MsgId() int32 {
	return 6001
}

func (*UnionAreaWarRespMsg) Comment() string {
	return "查询据点战基础信息"
}
		
func (*UnionAreaWarGroupReqMsg) MsgId() int32 {
	return 206003
}

func (*UnionAreaWarGroupReqMsg) Comment() string {
	return "请求查看分组"
}
		
func (*UnionAreaWarGroupRespMsg) MsgId() int32 {
	return 6003
}

func (*UnionAreaWarGroupRespMsg) Comment() string {
	return "请求查看分组"
}
		
func (*UnionAreaWarDefendersReqMsg) MsgId() int32 {
	return 206004
}

func (*UnionAreaWarDefendersReqMsg) Comment() string {
	return "请求防守人列表"
}
		
func (*UnionAreaWarDefendersRespMsg) MsgId() int32 {
	return 6004
}

func (*UnionAreaWarDefendersRespMsg) Comment() string {
	return "请求防守人列表"
}
		
func (*UnionAreaWarDefendersUpdateReqMsg) MsgId() int32 {
	return 206005
}

func (*UnionAreaWarDefendersUpdateReqMsg) Comment() string {
	return "刷新防守人列表"
}
		
func (*UnionAreaWarDefendersUpdateRespMsg) MsgId() int32 {
	return 6005
}

func (*UnionAreaWarDefendersUpdateRespMsg) Comment() string {
	return "刷新防守人列表"
}
		
func (*UnionAreWarListReqMsg) MsgId() int32 {
	return 206006
}

func (*UnionAreWarListReqMsg) Comment() string {
	return "获取据点列表"
}
		
func (*UnionAreaWarListRespMsg) MsgId() int32 {
	return 6006
}

func (*UnionAreaWarListRespMsg) Comment() string {
	return "获取据点列表"
}
		
func (*UnionAreWarListUpdateReqMsg) MsgId() int32 {
	return 206007
}

func (*UnionAreWarListUpdateReqMsg) Comment() string {
	return "刷新据点列表"
}
		
func (*UnionAreaWarListUpdateRespMsg) MsgId() int32 {
	return 6007
}

func (*UnionAreaWarListUpdateRespMsg) Comment() string {
	return "刷新据点列表"
}
		
func (*UnionAreaWarAttackReqMsg) MsgId() int32 {
	return 206009
}

func (*UnionAreaWarAttackReqMsg) Comment() string {
	return "请求战斗场景攻击"
}
		
func (*UnionAreaWarAttackRespMsg) MsgId() int32 {
	return 6009
}

func (*UnionAreaWarAttackRespMsg) Comment() string {
	return "请求战斗场景攻击"
}
		
func (*UnionAreaWarDevelopReqMsg) MsgId() int32 {
	return 206011
}

func (*UnionAreaWarDevelopReqMsg) Comment() string {
	return "建设(捐献)"
}
		
func (*UnionAreaWarDevelopRespMsg) MsgId() int32 {
	return 6011
}

func (*UnionAreaWarDevelopRespMsg) Comment() string {
	return "建设(捐献)"
}
		
func (*UnionAreaWarTreasuryReqMsg) MsgId() int32 {
	return 206013
}

func (*UnionAreaWarTreasuryReqMsg) Comment() string {
	return "加载宝库信息"
}
		
func (*UnionAreaWarTreasuryRespMsg) MsgId() int32 {
	return 6013
}

func (*UnionAreaWarTreasuryRespMsg) Comment() string {
	return "加载宝库信息"
}
		
func (*UnionAreaWarTreasuryDrawReqMsg) MsgId() int32 {
	return 206014
}

func (*UnionAreaWarTreasuryDrawReqMsg) Comment() string {
	return "宝库抽奖"
}
		
func (*UnionAreaWarTreasuryDrawRespMsg) MsgId() int32 {
	return 6014
}

func (*UnionAreaWarTreasuryDrawRespMsg) Comment() string {
	return "宝库抽奖"
}
		
func (*UnionAreaWarSummonDragonReqMsg) MsgId() int32 {
	return 206015
}

func (*UnionAreaWarSummonDragonReqMsg) Comment() string {
	return "请求召唤神龙"
}
		
func (*UnionAreaWarSummonDragonRespMsg) MsgId() int32 {
	return 6015
}

func (*UnionAreaWarSummonDragonRespMsg) Comment() string {
	return "请求召唤神龙"
}
		
func (*UnionAreaWarWorshipReq) MsgId() int32 {
	return 206016
}

func (*UnionAreaWarWorshipReq) Comment() string {
	return "查询据点战膜拜"
}
		
func (*UnionAreaWarWorshipRsp) MsgId() int32 {
	return 6016
}

func (*UnionAreaWarWorshipRsp) Comment() string {
	return "查询据点战膜拜"
}
		
func (*UnionAreaWarFightSceneReqMsg) MsgId() int32 {
	return 206019
}

func (*UnionAreaWarFightSceneReqMsg) Comment() string {
	return "请求战斗场景数据"
}
		
func (*UnionAreaWarFightSceneRespMsg) MsgId() int32 {
	return 6019
}

func (*UnionAreaWarFightSceneRespMsg) Comment() string {
	return "请求战斗场景数据"
}
		
func (*UnionAreaWarConstructReqMsg) MsgId() int32 {
	return 206020
}

func (*UnionAreaWarConstructReqMsg) Comment() string {
	return "请求捐献建设数据"
}
		
func (*UnionAreaWarConstructRespMsg) MsgId() int32 {
	return 6020
}

func (*UnionAreaWarConstructRespMsg) Comment() string {
	return "请求捐献建设数据"
}
		
func (*UnionAreaWarReportReqMsg) MsgId() int32 {
	return 206021
}

func (*UnionAreaWarReportReqMsg) Comment() string {
	return "请求战报数据"
}
		
func (*UnionAreaWarReportRespMsg) MsgId() int32 {
	return 6021
}

func (*UnionAreaWarReportRespMsg) Comment() string {
	return "请求战报数据"
}
		
func (*UnionAreaWarReportDetailReqMsg) MsgId() int32 {
	return 206022
}

func (*UnionAreaWarReportDetailReqMsg) Comment() string {
	return "请求战报伤害详情数据"
}
		
func (*UnionAreaWarReportDetailRespMsg) MsgId() int32 {
	return 6022
}

func (*UnionAreaWarReportDetailRespMsg) Comment() string {
	return "请求战报伤害详情数据"
}
		
func (*UnionAreaWarUnionDamageReqMsg) MsgId() int32 {
	return 206023
}

func (*UnionAreaWarUnionDamageReqMsg) Comment() string {
	return "请求联盟伤害数据"
}
		
func (*UnionAreaWarUnionDamageRespMsg) MsgId() int32 {
	return 6023
}

func (*UnionAreaWarUnionDamageRespMsg) Comment() string {
	return "请求联盟伤害数据"
}
		
func (*UnionAreaWarBetDataReqMsg) MsgId() int32 {
	return 206024
}

func (*UnionAreaWarBetDataReqMsg) Comment() string {
	return "请求竞猜界面数据"
}
		
func (*UnionAreaWarBetDataRespMsg) MsgId() int32 {
	return 6024
}

func (*UnionAreaWarBetDataRespMsg) Comment() string {
	return "请求竞猜界面数据"
}
		
func (*UnionAreaWarBetSelectReqMsg) MsgId() int32 {
	return 206025
}

func (*UnionAreaWarBetSelectReqMsg) Comment() string {
	return "请求竞猜选择"
}
		
func (*UnionAreaWarBetSelectRespMsg) MsgId() int32 {
	return 6025
}

func (*UnionAreaWarBetSelectRespMsg) Comment() string {
	return "请求竞猜选择"
}
		
func (*UnionAreaWarBetRewardReqMsg) MsgId() int32 {
	return 206026
}

func (*UnionAreaWarBetRewardReqMsg) Comment() string {
	return "请求竞猜奖励"
}
		
func (*UnionAreaWarBetRewardRespMsg) MsgId() int32 {
	return 6026
}

func (*UnionAreaWarBetRewardRespMsg) Comment() string {
	return "请求竞猜奖励"
}
		
func (*UnionAreaWarGetJoinMemberListReq) MsgId() int32 {
	return 206027
}

func (*UnionAreaWarGetJoinMemberListReq) Comment() string {
	return "查看自己所属联盟的参与名单"
}
		
func (*UnionAreaWarGetJoinMemberListRsp) MsgId() int32 {
	return 6027
}

func (*UnionAreaWarGetJoinMemberListRsp) Comment() string {
	return "查看自己所属联盟的参与名单"
}
		
func (*UnionAreaWarUnionGroupDamageReqMsg) MsgId() int32 {
	return 206029
}

func (*UnionAreaWarUnionGroupDamageReqMsg) Comment() string {
	return "请求查看联盟分段位伤害信息"
}
		
func (*UnionAreaWarUnionGroupDamageRespMsg) MsgId() int32 {
	return 6029
}

func (*UnionAreaWarUnionGroupDamageRespMsg) Comment() string {
	return "请求查看联盟分段位伤害信息"
}
		
func (*UnionAreaWarRedDotReqMsg) MsgId() int32 {
	return 206030
}

func (*UnionAreaWarRedDotReqMsg) Comment() string {
	return "请求红点信息"
}
		
func (*UnionAreaWarRedDotRespMsg) MsgId() int32 {
	return 6030
}

func (*UnionAreaWarRedDotRespMsg) Comment() string {
	return "请求红点信息"
}
		
func (*UnionAreaWarGuessPlayersReqMsg) MsgId() int32 {
	return 206031
}

func (*UnionAreaWarGuessPlayersReqMsg) Comment() string {
	return "请求查看竞猜结果的获奖玩家列表"
}
		
func (*UnionAreaWarGuessPlayersRespMsg) MsgId() int32 {
	return 6031
}

func (*UnionAreaWarGuessPlayersRespMsg) Comment() string {
	return "请求查看竞猜结果的获奖玩家列表"
}
		
func (*UnionAreaWarDanGroupCountReq) MsgId() int32 {
	return 206032
}

func (*UnionAreaWarDanGroupCountReq) Comment() string {
	return "查看所有段位的分组数量"
}
		
func (*UnionAreaWarDanGroupCountResp) MsgId() int32 {
	return 6032
}

func (*UnionAreaWarDanGroupCountResp) Comment() string {
	return "查看所有段位的分组数量"
}
		
func (*AskWayEnterReq) MsgId() int32 {
	return 206101
}

func (*AskWayEnterReq) Comment() string {
	return "进入活动"
}
		
func (*AskWayEnterRsp) MsgId() int32 {
	return 6101
}

func (*AskWayEnterRsp) Comment() string {
	return "进入活动"
}
		
func (*AskWayMatchReq) MsgId() int32 {
	return 206102
}

func (*AskWayMatchReq) Comment() string {
	return "斗法匹配"
}
		
func (*AskWayMatchRsp) MsgId() int32 {
	return 6102
}

func (*AskWayMatchRsp) Comment() string {
	return "斗法匹配"
}
		
func (*AskWayBattleReq) MsgId() int32 {
	return 206103
}

func (*AskWayBattleReq) Comment() string {
	return "斗法战斗"
}
		
func (*AskWayBattleRsp) MsgId() int32 {
	return 6103
}

func (*AskWayBattleRsp) Comment() string {
	return "斗法战斗"
}
		
func (*AskWayGetReportReq) MsgId() int32 {
	return 206104
}

func (*AskWayGetReportReq) Comment() string {
	return "获取战报"
}
		
func (*AskWayGetReportRsp) MsgId() int32 {
	return 6104
}

func (*AskWayGetReportRsp) Comment() string {
	return "获取战报"
}
		
func (*AskWayBattleReplyReq) MsgId() int32 {
	return 206105
}

func (*AskWayBattleReplyReq) Comment() string {
	return "斗法战斗回放"
}
		
func (*AskWayBattleReplyRsp) MsgId() int32 {
	return 6105
}

func (*AskWayBattleReplyRsp) Comment() string {
	return "斗法战斗回放"
}
		
func (*AskWayReceiveTierRewardReq) MsgId() int32 {
	return 206106
}

func (*AskWayReceiveTierRewardReq) Comment() string {
	return "领取段位奖励"
}
		
func (*AskWayReceiveTierRewardRsp) MsgId() int32 {
	return 6106
}

func (*AskWayReceiveTierRewardRsp) Comment() string {
	return "领取段位奖励"
}
		
func (*AskWayReceiveScoreRewardReq) MsgId() int32 {
	return 206107
}

func (*AskWayReceiveScoreRewardReq) Comment() string {
	return "领取斗法榜单积分奖励"
}
		
func (*AskWayReceiveScoreRewardRsp) MsgId() int32 {
	return 6107
}

func (*AskWayReceiveScoreRewardRsp) Comment() string {
	return "领取斗法榜单积分奖励"
}
		
func (*AskWayBuyFightTicketReq) MsgId() int32 {
	return 206108
}

func (*AskWayBuyFightTicketReq) Comment() string {
	return "购买挑战令道具"
}
		
func (*AskWayBuyFightTicketRsp) MsgId() int32 {
	return 6108
}

func (*AskWayBuyFightTicketRsp) Comment() string {
	return "购买挑战令道具"
}
		
func (*AskWayGetPlayerDetailReq) MsgId() int32 {
	return 206109
}

func (*AskWayGetPlayerDetailReq) Comment() string {
	return "查看玩家数据"
}
		
func (*AskWayGetPlayerDetailRsp) MsgId() int32 {
	return 6109
}

func (*AskWayGetPlayerDetailRsp) Comment() string {
	return "查看玩家数据"
}
		
func (*AskWayGetGuessCoinReq) MsgId() int32 {
	return 206110
}

func (*AskWayGetGuessCoinReq) Comment() string {
	return "领取竞猜币"
}
		
func (*AskWayGetGuessCoinRsp) MsgId() int32 {
	return 6110
}

func (*AskWayGetGuessCoinRsp) Comment() string {
	return "领取竞猜币"
}
		
func (*AskWayToSkyGuessReq) MsgId() int32 {
	return 206111
}

func (*AskWayToSkyGuessReq) Comment() string {
	return "登天竞猜"
}
		
func (*AskWayToSkyGuessRsp) MsgId() int32 {
	return 6111
}

func (*AskWayToSkyGuessRsp) Comment() string {
	return "登天竞猜"
}
		
func (*AskWayToSkyBattleReplyReq) MsgId() int32 {
	return 206112
}

func (*AskWayToSkyBattleReplyReq) Comment() string {
	return "登天战斗回放"
}
		
func (*AskWayToSkyBattleReplyRsp) MsgId() int32 {
	return 6112
}

func (*AskWayToSkyBattleReplyRsp) Comment() string {
	return "登天战斗回放"
}
		
func (*AskWayWorshipReq) MsgId() int32 {
	return 206113
}

func (*AskWayWorshipReq) Comment() string {
	return "朝拜"
}
		
func (*AskWayWorshipRsp) MsgId() int32 {
	return 6113
}

func (*AskWayWorshipRsp) Comment() string {
	return "朝拜"
}
		
func (*AskWayReceiveToSkyRankRewardReq) MsgId() int32 {
	return 206114
}

func (*AskWayReceiveToSkyRankRewardReq) Comment() string {
	return "领取登天榜单奖励"
}
		
func (*AskWayReceiveToSkyRankRewardRsp) MsgId() int32 {
	return 6114
}

func (*AskWayReceiveToSkyRankRewardRsp) Comment() string {
	return "领取登天榜单奖励"
}
		
func (*AskWayToSkyRosterReq) MsgId() int32 {
	return 206115
}

func (*AskWayToSkyRosterReq) Comment() string {
	return "往届妖神榜"
}
		
func (*AskWayToSkyRosterRsp) MsgId() int32 {
	return 6115
}

func (*AskWayToSkyRosterRsp) Comment() string {
	return "往届妖神榜"
}
		
func (*AskWayToSkyGetBattleResultReq) MsgId() int32 {
	return 206116
}

func (*AskWayToSkyGetBattleResultReq) Comment() string {
	return "获取登天期组别数据"
}
		
func (*AskWayToSkyGetBattleResultRsp) MsgId() int32 {
	return 6116
}

func (*AskWayToSkyGetBattleResultRsp) Comment() string {
	return "获取登天期组别数据"
}
		
func (*AskWayGetGuessInfoReq) MsgId() int32 {
	return 206117
}

func (*AskWayGetGuessInfoReq) Comment() string {
	return "获取竞猜信息"
}
		
func (*AskWayGetGuessInfoRsp) MsgId() int32 {
	return 6117
}

func (*AskWayGetGuessInfoRsp) Comment() string {
	return "获取竞猜信息"
}
		
func (*AskWayToSkyGetReportReq) MsgId() int32 {
	return 206118
}

func (*AskWayToSkyGetReportReq) Comment() string {
	return "获取登天战报数据"
}
		
func (*AskWayToSkyGetReportRsp) MsgId() int32 {
	return 6118
}

func (*AskWayToSkyGetReportRsp) Comment() string {
	return "获取登天战报数据"
}
		
func (*AskWayGetGetCurStateInfoReq) MsgId() int32 {
	return 206132
}

func (*AskWayGetGetCurStateInfoReq) Comment() string {
	return "获取海报数据"
}
		
func (*AskWayGetGetCurStateInfoRsp) MsgId() int32 {
	return 6132
}

func (*AskWayGetGetCurStateInfoRsp) Comment() string {
	return "获取海报数据"
}
		
func (*UnionFightMainReq) MsgId() int32 {
	return 206701
}

func (*UnionFightMainReq) Comment() string {
	return "妖盟夺位战主界面"
}
		
func (*UnionFightMainRsp) MsgId() int32 {
	return 6701
}

func (*UnionFightMainRsp) Comment() string {
	return "妖盟夺位战主界面"
}
		
func (*UnionFightApplyReq) MsgId() int32 {
	return 206702
}

func (*UnionFightApplyReq) Comment() string {
	return "妖盟夺位战报名"
}
		
func (*UnionFightApplyRsp) MsgId() int32 {
	return 6702
}

func (*UnionFightApplyRsp) Comment() string {
	return "妖盟夺位战报名"
}
		
func (*UnionFightEnterReq) MsgId() int32 {
	return 206703
}

func (*UnionFightEnterReq) Comment() string {
	return "妖盟夺位战进入战斗界面"
}
		
func (*UnionFightEnterRsp) MsgId() int32 {
	return 6703
}

func (*UnionFightEnterRsp) Comment() string {
	return "妖盟夺位战进入战斗界面"
}
		
func (*UnionFightRequestReq) MsgId() int32 {
	return 206704
}

func (*UnionFightRequestReq) Comment() string {
	return "妖盟夺位战请战"
}
		
func (*UnionFightRequestRsp) MsgId() int32 {
	return 6704
}

func (*UnionFightRequestRsp) Comment() string {
	return "妖盟夺位战请战"
}
		
func (*UnionFightPositionReq) MsgId() int32 {
	return 206705
}

func (*UnionFightPositionReq) Comment() string {
	return "妖盟夺位战上阵"
}
		
func (*UnionFightPositionRsp) MsgId() int32 {
	return 6705
}

func (*UnionFightPositionRsp) Comment() string {
	return "妖盟夺位战上阵"
}
		
func (*UnionFightUnPositionReq) MsgId() int32 {
	return 206706
}

func (*UnionFightUnPositionReq) Comment() string {
	return "妖盟夺位战下阵"
}
		
func (*UnionFightUnPositionRsp) MsgId() int32 {
	return 6706
}

func (*UnionFightUnPositionRsp) Comment() string {
	return "妖盟夺位战下阵"
}
		
func (*UnionFightGroupRankReq) MsgId() int32 {
	return 206707
}

func (*UnionFightGroupRankReq) Comment() string {
	return "妖盟夺位战小组排行"
}
		
func (*UnionFightGroupRankRsp) MsgId() int32 {
	return 6707
}

func (*UnionFightGroupRankRsp) Comment() string {
	return "妖盟夺位战小组排行"
}
		
func (*UnionFightLeaveReq) MsgId() int32 {
	return 206708
}

func (*UnionFightLeaveReq) Comment() string {
	return "妖盟夺位战离开"
}
		
func (*UnionFightSupremacyListReq) MsgId() int32 {
	return 206709
}

func (*UnionFightSupremacyListReq) Comment() string {
	return "妖盟夺位战至尊榜"
}
		
func (*UnionFightSupremacyListRsp) MsgId() int32 {
	return 6709
}

func (*UnionFightSupremacyListRsp) Comment() string {
	return "妖盟夺位战至尊榜"
}
		
func (*UnionFightWorshipReq) MsgId() int32 {
	return 206710
}

func (*UnionFightWorshipReq) Comment() string {
	return "妖盟夺位战膜拜"
}
		
func (*UnionFightWorshipRsp) MsgId() int32 {
	return 6710
}

func (*UnionFightWorshipRsp) Comment() string {
	return "妖盟夺位战膜拜"
}
		
func (*UnionFightFightPlaybackReq) MsgId() int32 {
	return 206711
}

func (*UnionFightFightPlaybackReq) Comment() string {
	return "妖盟夺位战回放"
}
		
func (*UnionFightFightPlaybackRsp) MsgId() int32 {
	return 6711
}

func (*UnionFightFightPlaybackRsp) Comment() string {
	return "妖盟夺位战回放"
}
		
func (*UnionFightRecordReq) MsgId() int32 {
	return 206712
}

func (*UnionFightRecordReq) Comment() string {
	return "妖盟夺位战对战记录"
}
		
func (*UnionFightRecordRsp) MsgId() int32 {
	return 6712
}

func (*UnionFightRecordRsp) Comment() string {
	return "妖盟夺位战对战记录"
}
		
func (*UnionFightGetHistoryDataReq) MsgId() int32 {
	return 206713
}

func (*UnionFightGetHistoryDataReq) Comment() string {
	return "妖盟夺位战历史最高"
}
		
func (*UnionFightGetHistoryDataRsp) MsgId() int32 {
	return 6713
}

func (*UnionFightGetHistoryDataRsp) Comment() string {
	return "妖盟夺位战历史最高"
}
		
func (*UnionFightGetLockedDetailReq) MsgId() int32 {
	return 206714
}

func (*UnionFightGetLockedDetailReq) Comment() string {
	return "妖盟夺位战获取锁定属性"
}
		
func (*UnionFightGetLockedDetailRsp) MsgId() int32 {
	return 6714
}

func (*UnionFightGetLockedDetailRsp) Comment() string {
	return "妖盟夺位战获取锁定属性"
}
		
func (*UnionFightChangeLockStatusReq) MsgId() int32 {
	return 206715
}

func (*UnionFightChangeLockStatusReq) Comment() string {
	return "妖盟夺位战设置锁定状态"
}
		
func (*UnionFightChangeLockStatusRsp) MsgId() int32 {
	return 6715
}

func (*UnionFightChangeLockStatusRsp) Comment() string {
	return "妖盟夺位战设置锁定状态"
}
		
func (*UnionFightGetTodayResultReq) MsgId() int32 {
	return 206716
}

func (*UnionFightGetTodayResultReq) Comment() string {
	return "妖盟夺位战今日详情"
}
		
func (*UnionFightGetTodayResultRsp) MsgId() int32 {
	return 6716
}

func (*UnionFightGetTodayResultRsp) Comment() string {
	return "妖盟夺位战今日详情"
}
		
func (*UnionFightGetUnionRankListReq) MsgId() int32 {
	return 206717
}

func (*UnionFightGetUnionRankListReq) Comment() string {
	return "获取妖盟排行列表"
}
		
func (*GodIslandBaseInfoReqMsg) MsgId() int32 {
	return 206801
}

func (*GodIslandBaseInfoReqMsg) Comment() string {
	return "基础信息"
}
		
func (*GodIslandBaseInfoRespMsg) MsgId() int32 {
	return 6801
}

func (*GodIslandBaseInfoRespMsg) Comment() string {
	return "基础信息"
}
		
func (*GodIslandGroupReqMsg) MsgId() int32 {
	return 206802
}

func (*GodIslandGroupReqMsg) Comment() string {
	return "查看分组"
}
		
func (*GodIslandGroupRespMsg) MsgId() int32 {
	return 6802
}

func (*GodIslandGroupRespMsg) Comment() string {
	return "查看分组"
}
		
func (*GodIslandReportReqMsg) MsgId() int32 {
	return 206803
}

func (*GodIslandReportReqMsg) Comment() string {
	return "请求妖盟建筑战报"
}
		
func (*GodIslandReportRespMsg) MsgId() int32 {
	return 6803
}

func (*GodIslandReportRespMsg) Comment() string {
	return "请求妖盟建筑战报"
}
		
func (*GodIslandReportDetailReqMsg) MsgId() int32 {
	return 206804
}

func (*GodIslandReportDetailReqMsg) Comment() string {
	return "请求妖盟建筑战报详情数据"
}
		
func (*GodIslandReportDetailRespMsg) MsgId() int32 {
	return 6804
}

func (*GodIslandReportDetailRespMsg) Comment() string {
	return "请求妖盟建筑战报详情数据"
}
		
func (*GodIslandUnionDamageReqMsg) MsgId() int32 {
	return 206805
}

func (*GodIslandUnionDamageReqMsg) Comment() string {
	return "请求联盟战功数据"
}
		
func (*GodIslandUnionDamageRespMsg) MsgId() int32 {
	return 6805
}

func (*GodIslandUnionDamageRespMsg) Comment() string {
	return "请求联盟战功数据"
}
		
func (*GodIslandGetJoinMemberListReq) MsgId() int32 {
	return 206806
}

func (*GodIslandGetJoinMemberListReq) Comment() string {
	return "查看自己所属联盟的参与名单"
}
		
func (*GodIslandGetJoinMemberListRsp) MsgId() int32 {
	return 6806
}

func (*GodIslandGetJoinMemberListRsp) Comment() string {
	return "查看自己所属联盟的参与名单"
}
		
func (*GodIslandUnionGroupDamageReqMsg) MsgId() int32 {
	return 206807
}

func (*GodIslandUnionGroupDamageReqMsg) Comment() string {
	return "妖盟分段位伤害信息"
}
		
func (*GodIslandUnionGroupDamageRespMsg) MsgId() int32 {
	return 6807
}

func (*GodIslandUnionGroupDamageRespMsg) Comment() string {
	return "妖盟分段位伤害信息"
}
		
func (*GodIslandRedDotReqMsg) MsgId() int32 {
	return 206808
}

func (*GodIslandRedDotReqMsg) Comment() string {
	return "红点信息"
}
		
func (*GodIslandRedDotRespMsg) MsgId() int32 {
	return 6808
}

func (*GodIslandRedDotRespMsg) Comment() string {
	return "红点信息"
}
		
func (*GodIslandWorshipReq) MsgId() int32 {
	return 206809
}

func (*GodIslandWorshipReq) Comment() string {
	return "点赞"
}
		
func (*GodIslandWorshipResp) MsgId() int32 {
	return 6809
}

func (*GodIslandWorshipResp) Comment() string {
	return "点赞"
}
		
func (*GodIslandPlayerReportReqMsg) MsgId() int32 {
	return 206810
}

func (*GodIslandPlayerReportReqMsg) Comment() string {
	return "请求个人战报"
}
		
func (*GodIslandPlayerReportRespMsg) MsgId() int32 {
	return 6810
}

func (*GodIslandPlayerReportRespMsg) Comment() string {
	return "请求个人战报"
}
		
func (*GodIslandUpdatePowerReq) MsgId() int32 {
	return 206811
}

func (*GodIslandUpdatePowerReq) Comment() string {
	return "妖力转灵力"
}
		
func (*GodIslandUpdatePowerResp) MsgId() int32 {
	return 6811
}

func (*GodIslandUpdatePowerResp) Comment() string {
	return "妖力转灵力"
}
		
func (*GodIslandUnionBattleScoreListReq) MsgId() int32 {
	return 206812
}

func (*GodIslandUnionBattleScoreListReq) Comment() string {
	return "查看妖盟战功列表"
}
		
func (*GodIslandUnionBattleScoreListResp) MsgId() int32 {
	return 6812
}

func (*GodIslandUnionBattleScoreListResp) Comment() string {
	return "查看妖盟战功列表"
}
		
func (*GodIslandLiquidReceiveRecordReq) MsgId() int32 {
	return 206813
}

func (*GodIslandLiquidReceiveRecordReq) Comment() string {
	return "灵液领取记录"
}
		
func (*GodIslandLiquidReceiveRecordResp) MsgId() int32 {
	return 6813
}

func (*GodIslandLiquidReceiveRecordResp) Comment() string {
	return "灵液领取记录"
}
		
func (*GodIslandUnionMemberScoreListReq) MsgId() int32 {
	return 206814
}

func (*GodIslandUnionMemberScoreListReq) Comment() string {
	return "妖盟成员战功统计"
}
		
func (*GodIslandUnionMemberScoreListResp) MsgId() int32 {
	return 6814
}

func (*GodIslandUnionMemberScoreListResp) Comment() string {
	return "妖盟成员战功统计"
}
		
func (*GodIslandHeartBeatReqMsg) MsgId() int32 {
	return 206819
}

func (*GodIslandHeartBeatReqMsg) Comment() string {
	return "活动心跳"
}
		
func (*GodIslandHeartBeatRespMsg) MsgId() int32 {
	return 6819
}

func (*GodIslandHeartBeatRespMsg) Comment() string {
	return "活动心跳"
}
		
func (*GodIslandGameInfoReqMsg) MsgId() int32 {
	return 206820
}

func (*GodIslandGameInfoReqMsg) Comment() string {
	return "游戏基础信息"
}
		
func (*GodIslandGameInfoRespMsg) MsgId() int32 {
	return 6820
}

func (*GodIslandGameInfoRespMsg) Comment() string {
	return "游戏基础信息"
}
		
func (*GodIslandGameCityInfoReqMsg) MsgId() int32 {
	return 206823
}

func (*GodIslandGameCityInfoReqMsg) Comment() string {
	return "游戏大地图城市信息(在滑动地图时候将视野内的城市id发送)"
}
		
func (*GodIslandGameCityInfoRespMsg) MsgId() int32 {
	return 6823
}

func (*GodIslandGameCityInfoRespMsg) Comment() string {
	return "游戏大地图城市信息(在滑动地图时候将视野内的城市id发送)"
}
		
func (*GodIslandGameEventReqMsg) MsgId() int32 {
	return 206824
}

func (*GodIslandGameEventReqMsg) Comment() string {
	return "npc事件选择"
}
		
func (*GodIslandGameEventRespMsg) MsgId() int32 {
	return 6824
}

func (*GodIslandGameEventRespMsg) Comment() string {
	return "npc事件选择"
}
		
func (*GodIslandGameSpiritualBallInfoReqMsg) MsgId() int32 {
	return 206826
}

func (*GodIslandGameSpiritualBallInfoReqMsg) Comment() string {
	return "灵力球信息"
}
		
func (*GodIslandGameSpiritualBallInfoRespMsg) MsgId() int32 {
	return 6826
}

func (*GodIslandGameSpiritualBallInfoRespMsg) Comment() string {
	return "灵力球信息"
}
		
func (*GodIslandGameUseSpiritualBallReqMsg) MsgId() int32 {
	return 206827
}

func (*GodIslandGameUseSpiritualBallReqMsg) Comment() string {
	return "使用灵力球"
}
		
func (*GodIslandGameUseSpiritualBallRespMsg) MsgId() int32 {
	return 6827
}

func (*GodIslandGameUseSpiritualBallRespMsg) Comment() string {
	return "使用灵力球"
}
		
func (*GodIslandGameTargetCityInfoReqMsg) MsgId() int32 {
	return 206828
}

func (*GodIslandGameTargetCityInfoReqMsg) Comment() string {
	return "目标城市具体信息"
}
		
func (*GodIslandGameTargetCityInfoRespMsg) MsgId() int32 {
	return 6828
}

func (*GodIslandGameTargetCityInfoRespMsg) Comment() string {
	return "目标城市具体信息"
}
		
func (*GodIslandGameTargetCityLineInfoReqMsg) MsgId() int32 {
	return 206829
}

func (*GodIslandGameTargetCityLineInfoReqMsg) Comment() string {
	return "目标城市队列信息"
}
		
func (*GodIslandGameTargetCityLineInfoRespMsg) MsgId() int32 {
	return 6829
}

func (*GodIslandGameTargetCityLineInfoRespMsg) Comment() string {
	return "目标城市队列信息"
}
		
func (*GodIslandGameAttackReqMsg) MsgId() int32 {
	return 206830
}

func (*GodIslandGameAttackReqMsg) Comment() string {
	return "进攻"
}
		
func (*GodIslandGameAttackRespMsg) MsgId() int32 {
	return 6830
}

func (*GodIslandGameAttackRespMsg) Comment() string {
	return "进攻"
}
		
func (*GodIslandGameAutoAttackReqMsg) MsgId() int32 {
	return 206831
}

func (*GodIslandGameAutoAttackReqMsg) Comment() string {
	return "自动操作 托管"
}
		
func (*GodIslandGameAutoAttackRespMsg) MsgId() int32 {
	return 6831
}

func (*GodIslandGameAutoAttackRespMsg) Comment() string {
	return "自动操作 托管"
}
		
func (*GodIslandGameMoveReqMsg) MsgId() int32 {
	return 206832
}

func (*GodIslandGameMoveReqMsg) Comment() string {
	return "行军移动"
}
		
func (*GodIslandGameMoveRespMsg) MsgId() int32 {
	return 6832
}

func (*GodIslandGameMoveRespMsg) Comment() string {
	return "行军移动"
}
		
func (*GodIslandGameMiniMapInfoReqMsg) MsgId() int32 {
	return 206834
}

func (*GodIslandGameMiniMapInfoReqMsg) Comment() string {
	return "小地图信息"
}
		
func (*GodIslandGameMiniMapInfoRespMsg) MsgId() int32 {
	return 6834
}

func (*GodIslandGameMiniMapInfoRespMsg) Comment() string {
	return "小地图信息"
}
		
func (*GodIslandGameSetUnionTypeReqMsg) MsgId() int32 {
	return 206835
}

func (*GodIslandGameSetUnionTypeReqMsg) Comment() string {
	return "设置联盟类型"
}
		
func (*GodIslandGameSetUnionTypeRespMsg) MsgId() int32 {
	return 6835
}

func (*GodIslandGameSetUnionTypeRespMsg) Comment() string {
	return "设置联盟类型"
}
		
func (*GodIslandGamePlantInfoReqMsg) MsgId() int32 {
	return 206836
}

func (*GodIslandGamePlantInfoReqMsg) Comment() string {
	return "联盟种草信息"
}
		
func (*GodIslandGamePlantInfoRespMsg) MsgId() int32 {
	return 6836
}

func (*GodIslandGamePlantInfoRespMsg) Comment() string {
	return "联盟种草信息"
}
		
func (*GodIslandUseFruitReq) MsgId() int32 {
	return 206837
}

func (*GodIslandUseFruitReq) Comment() string {
	return "使用仙果 补充灵力球的灵力"
}
		
func (*GodIslandUseFruitResp) MsgId() int32 {
	return 6837
}

func (*GodIslandUseFruitResp) Comment() string {
	return "使用仙果 补充灵力球的灵力"
}
		
func (*GodIslandGamePlantReceiveFruitsReqMsg) MsgId() int32 {
	return 206838
}

func (*GodIslandGamePlantReceiveFruitsReqMsg) Comment() string {
	return "联盟种草领取果实"
}
		
func (*GodIslandGamePlantReceiveFruitsRespMsg) MsgId() int32 {
	return 6838
}

func (*GodIslandGamePlantReceiveFruitsRespMsg) Comment() string {
	return "联盟种草领取果实"
}
		
func (*GodIslandGamePlantWaterReqMsg) MsgId() int32 {
	return 206839
}

func (*GodIslandGamePlantWaterReqMsg) Comment() string {
	return "联盟种草浇水"
}
		
func (*GodIslandGamePlantWaterRespMsg) MsgId() int32 {
	return 6839
}

func (*GodIslandGamePlantWaterRespMsg) Comment() string {
	return "联盟种草浇水"
}
		
func (*GodIslandCrystalInfoReqMsg) MsgId() int32 {
	return 206840
}

func (*GodIslandCrystalInfoReqMsg) Comment() string {
	return "灵晶数据加载"
}
		
func (*GodIslandCrystalInfoRespMsg) MsgId() int32 {
	return 6840
}

func (*GodIslandCrystalInfoRespMsg) Comment() string {
	return "灵晶数据加载"
}
		
func (*GodIslandCrystalReceiveReqMsg) MsgId() int32 {
	return 206841
}

func (*GodIslandCrystalReceiveReqMsg) Comment() string {
	return "领取灵晶"
}
		
func (*GodIslandCrystalReceiveRespMsg) MsgId() int32 {
	return 6841
}

func (*GodIslandCrystalReceiveRespMsg) Comment() string {
	return "领取灵晶"
}
		
func (*GodIslandGameStrategyInfoReqMsg) MsgId() int32 {
	return 206842
}

func (*GodIslandGameStrategyInfoReqMsg) Comment() string {
	return "联盟指挥信息"
}
		
func (*GodIslandGameStrategyInfoRespMsg) MsgId() int32 {
	return 6842
}

func (*GodIslandGameStrategyInfoRespMsg) Comment() string {
	return "联盟指挥信息"
}
		
func (*GodIslandGameCommanderSetReqMsg) MsgId() int32 {
	return 206843
}

func (*GodIslandGameCommanderSetReqMsg) Comment() string {
	return "联盟指挥设置"
}
		
func (*GodIslandGameCommanderSetRespMsg) MsgId() int32 {
	return 6843
}

func (*GodIslandGameCommanderSetRespMsg) Comment() string {
	return "联盟指挥设置"
}
		
func (*GodIslandGameSetUnionTargetReqMsg) MsgId() int32 {
	return 206844
}

func (*GodIslandGameSetUnionTargetReqMsg) Comment() string {
	return "联盟指挥设置目标"
}
		
func (*GodIslandGameSetUnionTargetRespMsg) MsgId() int32 {
	return 6844
}

func (*GodIslandGameSetUnionTargetRespMsg) Comment() string {
	return "联盟指挥设置目标"
}
		
func (*GodIslandGameForHelpReqMsg) MsgId() int32 {
	return 206845
}

func (*GodIslandGameForHelpReqMsg) Comment() string {
	return "联盟指挥请求支援"
}
		
func (*GodIslandGameForHelpRespMsg) MsgId() int32 {
	return 6845
}

func (*GodIslandGameForHelpRespMsg) Comment() string {
	return "联盟指挥请求支援"
}
		
func (*GodIslandGameAcclerateMoveMsg) MsgId() int32 {
	return 206849
}

func (*GodIslandGameAcclerateMoveMsg) Comment() string {
	return "加速行军"
}
		
func (*GodIslandGameAcclerateMoveRespMsg) MsgId() int32 {
	return 6849
}

func (*GodIslandGameAcclerateMoveRespMsg) Comment() string {
	return "加速行军"
}
		
func (*GodIslandGameSendLiquidReqMsg) MsgId() int32 {
	return 206851
}

func (*GodIslandGameSendLiquidReqMsg) Comment() string {
	return "联盟发送灵液"
}
		
func (*GodIslandGameSendLiquidRespMsg) MsgId() int32 {
	return 6851
}

func (*GodIslandGameSendLiquidRespMsg) Comment() string {
	return "联盟发送灵液"
}
		
func (*GodIslandGameSendLiquidRecordReqMsg) MsgId() int32 {
	return 206852
}

func (*GodIslandGameSendLiquidRecordReqMsg) Comment() string {
	return "联盟灵液发送记录"
}
		
func (*GodIslandGameSendLiquidRecordRespMsg) MsgId() int32 {
	return 6852
}

func (*GodIslandGameSendLiquidRecordRespMsg) Comment() string {
	return "联盟灵液发送记录"
}
		
func (*GodIslandGameCityBuffInfoReqMsg) MsgId() int32 {
	return 206855
}

func (*GodIslandGameCityBuffInfoReqMsg) Comment() string {
	return "请求城市界面buff信息"
}
		
func (*GodIslandGameCityBuffInfoRespMsg) MsgId() int32 {
	return 6855
}

func (*GodIslandGameCityBuffInfoRespMsg) Comment() string {
	return "请求城市界面buff信息"
}
		
func (*GodIslandBackHomeReq) MsgId() int32 {
	return 206856
}

func (*GodIslandBackHomeReq) Comment() string {
	return "阵亡之后 补兵完成之后从场外回大本营"
}
		
func (*GodIslandBackHomeResp) MsgId() int32 {
	return 6856
}

func (*GodIslandBackHomeResp) Comment() string {
	return "阵亡之后 补兵完成之后从场外回大本营"
}
		
func (*GodIslandSetFriendUnionReq) MsgId() int32 {
	return 206858
}

func (*GodIslandSetFriendUnionReq) Comment() string {
	return "修改盟友"
}
		
func (*GodIslandSetFriendUnionResp) MsgId() int32 {
	return 6858
}

func (*GodIslandSetFriendUnionResp) Comment() string {
	return "修改盟友"
}
		
func (*StarTrialChallengeReq) MsgId() int32 {
	return 206902
}

func (*StarTrialChallengeReq) Comment() string {
	return "星宿试炼战斗"
}
		
func (*StarTrialChallengeResp) MsgId() int32 {
	return 6902
}

func (*StarTrialChallengeResp) Comment() string {
	return "星宿试炼战斗"
}
		
func (*StarTrialRecordReq) MsgId() int32 {
	return 206903
}

func (*StarTrialRecordReq) Comment() string {
	return "获取日志信息"
}
		
func (*StarTrialRecordResp) MsgId() int32 {
	return 6903
}

func (*StarTrialRecordResp) Comment() string {
	return "获取日志信息"
}
		
func (*StarTrialBattleReplyReq) MsgId() int32 {
	return 206904
}

func (*StarTrialBattleReplyReq) Comment() string {
	return "战斗回放"
}
		
func (*StarTrialBattleReplyResp) MsgId() int32 {
	return 6904
}

func (*StarTrialBattleReplyResp) Comment() string {
	return "战斗回放"
}
		
func (*EnterStarTrialCodexMsgReq) MsgId() int32 {
	return 206905
}

func (*EnterStarTrialCodexMsgReq) Comment() string {
	return "进入星空图鉴信息"
}
		
func (*EnterStarTrialCodexMsgResp) MsgId() int32 {
	return 6905
}

func (*EnterStarTrialCodexMsgResp) Comment() string {
	return "进入星空图鉴信息"
}
		
func (*PlayerStarTrialCodexMsgReq) MsgId() int32 {
	return 206906
}

func (*PlayerStarTrialCodexMsgReq) Comment() string {
	return "获取详细图鉴击败信息"
}
		
func (*RspPlayerStarTrialCodexMsg) MsgId() int32 {
	return 6906
}

func (*RspPlayerStarTrialCodexMsg) Comment() string {
	return "获取详细图鉴击败信息"
}
		
func (*GetBossDetailDataReq) MsgId() int32 {
	return 206907
}

func (*GetBossDetailDataReq) Comment() string {
	return "获取Boss属性信息"
}
		
func (*GetBossDetailDataResp) MsgId() int32 {
	return 6907
}

func (*GetBossDetailDataResp) Comment() string {
	return "获取Boss属性信息"
}
		
func (*GetDailyFightRewardReq) MsgId() int32 {
	return 206908
}

func (*GetDailyFightRewardReq) Comment() string {
	return "每日领奖"
}
		
func (*GetDailyFightRewardResp) MsgId() int32 {
	return 6908
}

func (*GetDailyFightRewardResp) Comment() string {
	return "每日领奖"
}
		
func (*GetStarTrialGroupInfoReq) MsgId() int32 {
	return 206909
}

func (*GetStarTrialGroupInfoReq) Comment() string {
	return "获取分组信息"
}
		
func (*GetStarTrialGroupInfoResp) MsgId() int32 {
	return 6909
}

func (*GetStarTrialGroupInfoResp) Comment() string {
	return "获取分组信息"
}
		
func (*StarTrialGetPlayerDetailReq) MsgId() int32 {
	return 206910
}

func (*StarTrialGetPlayerDetailReq) Comment() string {
	return "查看玩家详情"
}
		
func (*StarTrialGetPlayerDetailResp) MsgId() int32 {
	return 6910
}

func (*StarTrialGetPlayerDetailResp) Comment() string {
	return "查看玩家详情"
}
		
func (*StarTrialEnterMainPanelReq) MsgId() int32 {
	return 206911
}

func (*StarTrialEnterMainPanelReq) Comment() string {
	return "进入星宿试炼"
}
		
func (*StarTrialEnterMainPanelResp) MsgId() int32 {
	return 6911
}

func (*StarTrialEnterMainPanelResp) Comment() string {
	return "进入星宿试炼"
}
		
func (*GatherEnergyEnterNewReq) MsgId() int32 {
	return 207001
}

func (*GatherEnergyEnterNewReq) Comment() string {
	return "进入聚灵阵"
}
		
func (*GatherEnergyEnterNewResp) MsgId() int32 {
	return 7001
}

func (*GatherEnergyEnterNewResp) Comment() string {
	return "进入聚灵阵"
}
		
func (*GatherEnergyOpenViewReq) MsgId() int32 {
	return 207003
}

func (*GatherEnergyOpenViewReq) Comment() string {
	return "开启灵阵界面"
}
		
func (*GatherEnergyOpenViewResp) MsgId() int32 {
	return 7003
}

func (*GatherEnergyOpenViewResp) Comment() string {
	return "开启灵阵界面"
}
		
func (*GatherEnergyOpenReq) MsgId() int32 {
	return 207004
}

func (*GatherEnergyOpenReq) Comment() string {
	return "开启某个灵阵"
}
		
func (*GatherEnergyOpenResp) MsgId() int32 {
	return 7004
}

func (*GatherEnergyOpenResp) Comment() string {
	return "开启某个灵阵"
}
		
func (*GatherEnergyFightReq) MsgId() int32 {
	return 207007
}

func (*GatherEnergyFightReq) Comment() string {
	return "挑战某人"
}
		
func (*GatherEnergyFightResp) MsgId() int32 {
	return 7007
}

func (*GatherEnergyFightResp) Comment() string {
	return "挑战某人"
}
		
func (*GatherEnergyRewardReq) MsgId() int32 {
	return 207011
}

func (*GatherEnergyRewardReq) Comment() string {
	return "聚灵阵奖励领取"
}
		
func (*GatherEnergyRewardResp) MsgId() int32 {
	return 7011
}

func (*GatherEnergyRewardResp) Comment() string {
	return "聚灵阵奖励领取"
}
		
func (*GatherEnergyLikeReq) MsgId() int32 {
	return 207012
}

func (*GatherEnergyLikeReq) Comment() string {
	return "聚灵阵排行榜点赞"
}
		
func (*GatherEnergyLikeResp) MsgId() int32 {
	return 7012
}

func (*GatherEnergyLikeResp) Comment() string {
	return "聚灵阵排行榜点赞"
}
		
func (*GatherEnergyTransformReq) MsgId() int32 {
	return 207015
}

func (*GatherEnergyTransformReq) Comment() string {
	return "灵气转换"
}
		
func (*GatherEnergyTransformResp) MsgId() int32 {
	return 7015
}

func (*GatherEnergyTransformResp) Comment() string {
	return "灵气转换"
}
		
func (*GatherEnergyLeaveReq) MsgId() int32 {
	return 207017
}

func (*GatherEnergyLeaveReq) Comment() string {
	return "离开聚灵阵"
}
		
func (*GatherEnergyLeaveResp) MsgId() int32 {
	return 7017
}

func (*GatherEnergyLeaveResp) Comment() string {
	return "离开聚灵阵"
}
		
func (*GatherEnergyNoticeReq) MsgId() int32 {
	return 207018
}

func (*GatherEnergyNoticeReq) Comment() string {
	return "聚灵阵公告"
}
		
func (*GatherEnergyNoticeResp) MsgId() int32 {
	return 7018
}

func (*GatherEnergyNoticeResp) Comment() string {
	return "聚灵阵公告"
}
		
func (*GatherEnergyGetADAwardReq) MsgId() int32 {
	return 207019
}

func (*GatherEnergyGetADAwardReq) Comment() string {
	return "聚灵阵获取广告奖励"
}
		
func (*GatherEnergyGetADAwardResp) MsgId() int32 {
	return 7019
}

func (*GatherEnergyGetADAwardResp) Comment() string {
	return "聚灵阵获取广告奖励"
}
		
func (*GatherEnergyFirstListViewReq) MsgId() int32 {
	return 207020
}

func (*GatherEnergyFirstListViewReq) Comment() string {
	return "打开聚灵阵一级列表界面（参与人数）"
}
		
func (*GatherEnergyFirstListViewResp) MsgId() int32 {
	return 7020
}

func (*GatherEnergyFirstListViewResp) Comment() string {
	return "打开聚灵阵一级列表界面（参与人数）"
}
		
func (*GatherEnergySecondListViewReq) MsgId() int32 {
	return 207021
}

func (*GatherEnergySecondListViewReq) Comment() string {
	return "打开聚灵阵二级列表界面（参与名字）"
}
		
func (*GatherEnergySecondListViewResp) MsgId() int32 {
	return 7021
}

func (*GatherEnergySecondListViewResp) Comment() string {
	return "打开聚灵阵二级列表界面（参与名字）"
}
		
func (*GatherEnergyInsideViewNewReq) MsgId() int32 {
	return 207022
}

func (*GatherEnergyInsideViewNewReq) Comment() string {
	return "打开某个聚灵阵内部"
}
		
func (*GatherEnergyInsideViewNewResp) MsgId() int32 {
	return 7022
}

func (*GatherEnergyInsideViewNewResp) Comment() string {
	return "打开某个聚灵阵内部"
}
		
func (*GatherEnergyFightReportNewReq) MsgId() int32 {
	return 207023
}

func (*GatherEnergyFightReportNewReq) Comment() string {
	return "打开战报界面"
}
		
func (*GatherEnergyFightReportNewResp) MsgId() int32 {
	return 7023
}

func (*GatherEnergyFightReportNewResp) Comment() string {
	return "打开战报界面"
}
		
func (*GatherEnergySearchNewReq) MsgId() int32 {
	return 207024
}

func (*GatherEnergySearchNewReq) Comment() string {
	return "查询某人参与的阵法"
}
		
func (*GatherEnergySearchNewResp) MsgId() int32 {
	return 7024
}

func (*GatherEnergySearchNewResp) Comment() string {
	return "查询某人参与的阵法"
}
		
func (*GatherEnergyAttendNewReq) MsgId() int32 {
	return 207025
}

func (*GatherEnergyAttendNewReq) Comment() string {
	return "加入聚灵阵"
}
		
func (*GatherEnergyAttendNewResp) MsgId() int32 {
	return 7025
}

func (*GatherEnergyAttendNewResp) Comment() string {
	return "加入聚灵阵"
}
		
func (*ReceiveSdkRewardReq) MsgId() int32 {
	return 207202
}

func (*ReceiveSdkRewardReq) Comment() string {
	return "领取sdk奖励"
}
		
func (*ReceiveSdkRewardRsp) MsgId() int32 {
	return 7202
}

func (*ReceiveSdkRewardRsp) Comment() string {
	return "领取sdk奖励"
}
		
func (*QQCardGetRewardReq) MsgId() int32 {
	return 207203
}

func (*QQCardGetRewardReq) Comment() string {
	return "领取qq卡券奖励"
}
		
func (*QQCardGetRewardRsp) MsgId() int32 {
	return 7203
}

func (*QQCardGetRewardRsp) Comment() string {
	return "领取qq卡券奖励"
}
		
func (*AlipayStartParamReq) MsgId() int32 {
	return 207204
}

func (*AlipayStartParamReq) Comment() string {
	return "支付宝启动参数"
}
		
func (*AlipayStartParamRsp) MsgId() int32 {
	return 7204
}

func (*AlipayStartParamRsp) Comment() string {
	return "支付宝启动参数"
}
		
func (*ReceiveSdkDailyRewardReq) MsgId() int32 {
	return 207205
}

func (*ReceiveSdkDailyRewardReq) Comment() string {
	return "Sdk七日登录奖励领取"
}
		
func (*ReceiveSdkDailyRewardRsp) MsgId() int32 {
	return 7205
}

func (*ReceiveSdkDailyRewardRsp) Comment() string {
	return "Sdk七日登录奖励领取"
}
		
func (*MeiTuanStartParamReq) MsgId() int32 {
	return 207206
}

func (*MeiTuanStartParamReq) Comment() string {
	return "美团启动参数"
}
		
func (*MeiTuanStartParamRsp) MsgId() int32 {
	return 7206
}

func (*MeiTuanStartParamRsp) Comment() string {
	return "美团启动参数"
}
		
func (*SkyWarEnterReq) MsgId() int32 {
	return 208401
}

func (*SkyWarEnterReq) Comment() string {
	return "进入征战诸天"
}
		
func (*SkyWarEnterRsp) MsgId() int32 {
	return 8401
}

func (*SkyWarEnterRsp) Comment() string {
	return "进入征战诸天"
}
		
func (*SkyWarRefreshEnemyReq) MsgId() int32 {
	return 208402
}

func (*SkyWarRefreshEnemyReq) Comment() string {
	return "刷新对手玩家"
}
		
func (*SkyWarRefreshEnemyRsp) MsgId() int32 {
	return 8402
}

func (*SkyWarRefreshEnemyRsp) Comment() string {
	return "刷新对手玩家"
}
		
func (*SkyWarFightReq) MsgId() int32 {
	return 208403
}

func (*SkyWarFightReq) Comment() string {
	return "斗法"
}
		
func (*SkyWarFightRsp) MsgId() int32 {
	return 8403
}

func (*SkyWarFightRsp) Comment() string {
	return "斗法"
}
		
func (*SkyWarRankReq) MsgId() int32 {
	return 208404
}

func (*SkyWarRankReq) Comment() string {
	return "获取排行榜数据"
}
		
func (*SkyWarRankRsp) MsgId() int32 {
	return 8404
}

func (*SkyWarRankRsp) Comment() string {
	return "获取排行榜数据"
}
		
func (*SkyWarLogReq) MsgId() int32 {
	return 208405
}

func (*SkyWarLogReq) Comment() string {
	return "获取日志"
}
		
func (*SkyWarLogRsp) MsgId() int32 {
	return 8405
}

func (*SkyWarLogRsp) Comment() string {
	return "获取日志"
}
		
func (*SkyWarLogPlaybackReq) MsgId() int32 {
	return 208406
}

func (*SkyWarLogPlaybackReq) Comment() string {
	return "战斗回放"
}
		
func (*SkyWarLogPlaybackRsp) MsgId() int32 {
	return 8406
}

func (*SkyWarLogPlaybackRsp) Comment() string {
	return "战斗回放"
}
		
func (*SkyWarFormationReq) MsgId() int32 {
	return 208407
}

func (*SkyWarFormationReq) Comment() string {
	return "获取布阵详情"
}
		
func (*SkyWarFormationRsp) MsgId() int32 {
	return 8407
}

func (*SkyWarFormationRsp) Comment() string {
	return "获取布阵详情"
}
		
func (*SkyWarSetOrderReq) MsgId() int32 {
	return 208408
}

func (*SkyWarSetOrderReq) Comment() string {
	return "设置顺位"
}
		
func (*SkyWarSetOrderRsp) MsgId() int32 {
	return 8408
}

func (*SkyWarSetOrderRsp) Comment() string {
	return "设置顺位"
}
		
func (*SkyWarSkyRankReq) MsgId() int32 {
	return 208409
}

func (*SkyWarSkyRankReq) Comment() string {
	return "诸天榜"
}
		
func (*SkyWarSkyRankRsp) MsgId() int32 {
	return 8409
}

func (*SkyWarSkyRankRsp) Comment() string {
	return "诸天榜"
}
		
func (*SkyWarWorshipReq) MsgId() int32 {
	return 208410
}

func (*SkyWarWorshipReq) Comment() string {
	return "膜拜"
}
		
func (*SkyWarWorshipRsp) MsgId() int32 {
	return 8410
}

func (*SkyWarWorshipRsp) Comment() string {
	return "膜拜"
}
		
func (*SkyWarBuyTimesReq) MsgId() int32 {
	return 208411
}

func (*SkyWarBuyTimesReq) Comment() string {
	return "购买次数(斗法、刷新)"
}
		
func (*SkyWarBuyTimesRsp) MsgId() int32 {
	return 8411
}

func (*SkyWarBuyTimesRsp) Comment() string {
	return "购买次数(斗法、刷新)"
}
		
func (*UnionDuelSyncMsgReq) MsgId() int32 {
	return 208501
}

func (*UnionDuelSyncMsgReq) Comment() string {
	return "妖盟对决-同步妖盟对决数据"
}
		
func (*SyncUnionDuelMsg) MsgId() int32 {
	return 8501
}

func (*SyncUnionDuelMsg) Comment() string {
	return "妖盟对决-同步妖盟对决数据"
}
		
func (*UnionDuelOpenViewReq) MsgId() int32 {
	return 208502
}

func (*UnionDuelOpenViewReq) Comment() string {
	return "妖盟对决-开启主界面"
}
		
func (*UnionDuelOpenViewResp) MsgId() int32 {
	return 8502
}

func (*UnionDuelOpenViewResp) Comment() string {
	return "妖盟对决-开启主界面"
}
		
func (*UnionDuelSyncDataReq) MsgId() int32 {
	return 208511
}

func (*UnionDuelSyncDataReq) Comment() string {
	return "妖盟对决-妖盟对决同步数据"
}
		
func (*UnionDuelSyncDataResp) MsgId() int32 {
	return 8511
}

func (*UnionDuelSyncDataResp) Comment() string {
	return "妖盟对决-妖盟对决同步数据"
}
		
func (*MarkFinishReq) MsgId() int32 {
	return 208602
}

func (*MarkFinishReq) Comment() string {
	return "评分完成"
}
		
func (*MarkFinishResp) MsgId() int32 {
	return 8602
}

func (*MarkFinishResp) Comment() string {
	return "评分完成"
}
		
func (*WorldRulePerceptionReq) MsgId() int32 {
	return 209001
}

func (*WorldRulePerceptionReq) Comment() string {
	return "法则感悟"
}
		
func (*WorldRulePerceptionResp) MsgId() int32 {
	return 9001
}

func (*WorldRulePerceptionResp) Comment() string {
	return "法则感悟"
}
		
func (*WorldRulePerceptionReplaceReq) MsgId() int32 {
	return 209002
}

func (*WorldRulePerceptionReplaceReq) Comment() string {
	return "法则感悟替换"
}
		
func (*WorldRulePerceptionReplaceResp) MsgId() int32 {
	return 9002
}

func (*WorldRulePerceptionReplaceResp) Comment() string {
	return "法则感悟替换"
}
		
func (*WorldRuleGetPerceptionLogReq) MsgId() int32 {
	return 209003
}

func (*WorldRuleGetPerceptionLogReq) Comment() string {
	return "获取法则感悟记录"
}
		
func (*WorldRuleGetPerceptionLogResp) MsgId() int32 {
	return 9003
}

func (*WorldRuleGetPerceptionLogResp) Comment() string {
	return "获取法则感悟记录"
}
		
func (*WorldRuleSwitchProgrammeReq) MsgId() int32 {
	return 209004
}

func (*WorldRuleSwitchProgrammeReq) Comment() string {
	return "切换法则方案"
}
		
func (*WorldRuleSwitchProgrammeResp) MsgId() int32 {
	return 9004
}

func (*WorldRuleSwitchProgrammeResp) Comment() string {
	return "切换法则方案"
}
		
func (*RuleTrialChallengeReq) MsgId() int32 {
	return 209101
}

func (*RuleTrialChallengeReq) Comment() string {
	return "请求法则试炼挑战"
}
		
func (*RuleTrialChallengeResp) MsgId() int32 {
	return 9101
}

func (*RuleTrialChallengeResp) Comment() string {
	return "请求法则试炼挑战"
}
		
func (*RuleTrialRepeatAllReq) MsgId() int32 {
	return 209103
}

func (*RuleTrialRepeatAllReq) Comment() string {
	return "请求法则试炼一键扫荡"
}
		
func (*RuleTrialRepeatAllResp) MsgId() int32 {
	return 9103
}

func (*RuleTrialRepeatAllResp) Comment() string {
	return "请求法则试炼一键扫荡"
}
		
func (*RuleTrialMonsterAttrReq) MsgId() int32 {
	return 209106
}

func (*RuleTrialMonsterAttrReq) Comment() string {
	return "请求法则试炼Boss数据"
}
		
func (*RuleTrialMonsterAttrResp) MsgId() int32 {
	return 9106
}

func (*RuleTrialMonsterAttrResp) Comment() string {
	return "请求法则试炼Boss数据"
}
		
func (*ADTimeTriggerReq) MsgId() int32 {
	return 209201
}

func (*ADTimeTriggerReq) Comment() string {
	return "缩时香礼包"
}
		
func (*ADTimeUseReq) MsgId() int32 {
	return 209202
}

func (*ADTimeUseReq) Comment() string {
	return "使用缩时香"
}
		
func (*HolyLandBattleApplyDataSyncReq) MsgId() int32 {
	return 209570
}

func (*HolyLandBattleApplyDataSyncReq) Comment() string {
	return "九幽争霸数据登录同步"
}
		
func (*HolyLandBattleApplyDataSync) MsgId() int32 {
	return 9570
}

func (*HolyLandBattleApplyDataSync) Comment() string {
	return "九幽争霸数据登录同步"
}
		
func (*ReqSignInFundRepair) MsgId() int32 {
	return 209601
}

func (*ReqSignInFundRepair) Comment() string {
	return "补签接口"
}
		
func (*RspSignInFundRepair) MsgId() int32 {
	return 9601
}

func (*RspSignInFundRepair) Comment() string {
	return "补签接口"
}
		
func (*EnterMountainSeaReq) MsgId() int32 {
	return 209701
}

func (*EnterMountainSeaReq) Comment() string {
	return "进入山海界面"
}
		
func (*EnterMountainSeaRsp) MsgId() int32 {
	return 9701
}

func (*EnterMountainSeaRsp) Comment() string {
	return "进入山海界面"
}
		
func (*EnterMountainSeaTeamReq) MsgId() int32 {
	return 209702
}

func (*EnterMountainSeaTeamReq) Comment() string {
	return "组队内部请求"
}
		
func (*EnterMountainSeaTeamRsp) MsgId() int32 {
	return 9702
}

func (*EnterMountainSeaTeamRsp) Comment() string {
	return "组队内部请求"
}
		
func (*MountainSeaTeamStartReq) MsgId() int32 {
	return 209703
}

func (*MountainSeaTeamStartReq) Comment() string {
	return "组队页面请求"
}
		
func (*MountainSeaTeamStartRsp) MsgId() int32 {
	return 9703
}

func (*MountainSeaTeamStartRsp) Comment() string {
	return "组队页面请求"
}
		
func (*MountainSeaCreateTeamReq) MsgId() int32 {
	return 209704
}

func (*MountainSeaCreateTeamReq) Comment() string {
	return "创建队伍"
}
		
func (*MountainSeaCreateTeamRsp) MsgId() int32 {
	return 9704
}

func (*MountainSeaCreateTeamRsp) Comment() string {
	return "创建队伍"
}
		
func (*MountainSeaGetTeamListReq) MsgId() int32 {
	return 209705
}

func (*MountainSeaGetTeamListReq) Comment() string {
	return "获取队伍列表"
}
		
func (*MountainSeaGetTeamListRsp) MsgId() int32 {
	return 9705
}

func (*MountainSeaGetTeamListRsp) Comment() string {
	return "获取队伍列表"
}
		
func (*MountainSeaGetTeamInfoReq) MsgId() int32 {
	return 209706
}

func (*MountainSeaGetTeamInfoReq) Comment() string {
	return "查找队伍"
}
		
func (*MountainSeaGetTeamInfoRsp) MsgId() int32 {
	return 9706
}

func (*MountainSeaGetTeamInfoRsp) Comment() string {
	return "查找队伍"
}
		
func (*MountainSeaJoinTeamReq) MsgId() int32 {
	return 209707
}

func (*MountainSeaJoinTeamReq) Comment() string {
	return "加入队伍"
}
		
func (*MountainSeaJoinTeamRsp) MsgId() int32 {
	return 9707
}

func (*MountainSeaJoinTeamRsp) Comment() string {
	return "加入队伍"
}
		
func (*MountainSeaCancelTeamApplyReq) MsgId() int32 {
	return 209708
}

func (*MountainSeaCancelTeamApplyReq) Comment() string {
	return "取消加入队伍"
}
		
func (*MountainSeaCancelTeamApplyRsp) MsgId() int32 {
	return 9708
}

func (*MountainSeaCancelTeamApplyRsp) Comment() string {
	return "取消加入队伍"
}
		
func (*MountainSeaApplyJoinTeamAgreeReq) MsgId() int32 {
	return 209709
}

func (*MountainSeaApplyJoinTeamAgreeReq) Comment() string {
	return "同意申请"
}
		
func (*MountainSeaApplyJoinTeamAgreeRsp) MsgId() int32 {
	return 9709
}

func (*MountainSeaApplyJoinTeamAgreeRsp) Comment() string {
	return "同意申请"
}
		
func (*MountainSeaApplyJoinTeamRefusedReq) MsgId() int32 {
	return 209710
}

func (*MountainSeaApplyJoinTeamRefusedReq) Comment() string {
	return "一键拒绝"
}
		
func (*MountainSeaApplyJoinTeamRefusedRsp) MsgId() int32 {
	return 9710
}

func (*MountainSeaApplyJoinTeamRefusedRsp) Comment() string {
	return "一键拒绝"
}
		
func (*MountainSeaQuitTeamReq) MsgId() int32 {
	return 209711
}

func (*MountainSeaQuitTeamReq) Comment() string {
	return "退出队伍"
}
		
func (*MountainSeaQuitTeamRsp) MsgId() int32 {
	return 9711
}

func (*MountainSeaQuitTeamRsp) Comment() string {
	return "退出队伍"
}
		
func (*MountainSeaKickOutTeamReq) MsgId() int32 {
	return 209712
}

func (*MountainSeaKickOutTeamReq) Comment() string {
	return "踢出队伍"
}
		
func (*MountainSeaKickOutTeamRsp) MsgId() int32 {
	return 9712
}

func (*MountainSeaKickOutTeamRsp) Comment() string {
	return "踢出队伍"
}
		
func (*MountainSeaChangeLeaderReq) MsgId() int32 {
	return 209713
}

func (*MountainSeaChangeLeaderReq) Comment() string {
	return "转让队长"
}
		
func (*MountainSeaChangeLeaderRsp) MsgId() int32 {
	return 9713
}

func (*MountainSeaChangeLeaderRsp) Comment() string {
	return "转让队长"
}
		
func (*MountainSeaTeamPrepareReq) MsgId() int32 {
	return 209714
}

func (*MountainSeaTeamPrepareReq) Comment() string {
	return "开始准备"
}
		
func (*MountainSeaTeamPrepareRsp) MsgId() int32 {
	return 9714
}

func (*MountainSeaTeamPrepareRsp) Comment() string {
	return "开始准备"
}
		
func (*MountainSeaMatchMemberReq) MsgId() int32 {
	return 209715
}

func (*MountainSeaMatchMemberReq) Comment() string {
	return "匹配"
}
		
func (*MountainSeaMatchMemberRsp) MsgId() int32 {
	return 9715
}

func (*MountainSeaMatchMemberRsp) Comment() string {
	return "匹配"
}
		
func (*MountainSeaStartBattleReq) MsgId() int32 {
	return 209716
}

func (*MountainSeaStartBattleReq) Comment() string {
	return "开始战斗"
}
		
func (*MountainSeaStartBattleRsp) MsgId() int32 {
	return 9716
}

func (*MountainSeaStartBattleRsp) Comment() string {
	return "开始战斗"
}
		
func (*MountainSeaGetPlayerInfoReq) MsgId() int32 {
	return 209717
}

func (*MountainSeaGetPlayerInfoReq) Comment() string {
	return "请求玩家数据"
}
		
func (*MountainSeaGetPlayerInfoRsp) MsgId() int32 {
	return 9717
}

func (*MountainSeaGetPlayerInfoRsp) Comment() string {
	return "请求玩家数据"
}
		
func (*MountainSeaWorshipReq) MsgId() int32 {
	return 209718
}

func (*MountainSeaWorshipReq) Comment() string {
	return "膜拜"
}
		
func (*MountainSeaWorshipRsp) MsgId() int32 {
	return 9718
}

func (*MountainSeaWorshipRsp) Comment() string {
	return "膜拜"
}
		
func (*MountainSeaGetBattleVideoReq) MsgId() int32 {
	return 209719
}

func (*MountainSeaGetBattleVideoReq) Comment() string {
	return "战斗回放"
}
		
func (*MountainSeaGetBattleVideoRsp) MsgId() int32 {
	return 9719
}

func (*MountainSeaGetBattleVideoRsp) Comment() string {
	return "战斗回放"
}
		
func (*MountainSeaGetBossInfoReq) MsgId() int32 {
	return 209720
}

func (*MountainSeaGetBossInfoReq) Comment() string {
	return "获取boss信息"
}
		
func (*MountainSeaGetBossInfoRsp) MsgId() int32 {
	return 9720
}

func (*MountainSeaGetBossInfoRsp) Comment() string {
	return "获取boss信息"
}
		
func (*MountainSeaChallengeTimeReq) MsgId() int32 {
	return 209721
}

func (*MountainSeaChallengeTimeReq) Comment() string {
	return "请求挑战次数"
}
		
func (*MountainSeaChallengeTimeRsp) MsgId() int32 {
	return 9721
}

func (*MountainSeaChallengeTimeRsp) Comment() string {
	return "请求挑战次数"
}
		
func (*MountainSeaStartMatchReq) MsgId() int32 {
	return 209722
}

func (*MountainSeaStartMatchReq) Comment() string {
	return "开始匹配"
}
		
func (*MountainSeaStartMatchRsp) MsgId() int32 {
	return 9722
}

func (*MountainSeaStartMatchRsp) Comment() string {
	return "开始匹配"
}
		
func (*MountainSeaInviteReq) MsgId() int32 {
	return 209723
}

func (*MountainSeaInviteReq) Comment() string {
	return "邀请"
}
		
func (*MountainSeaInviteRsp) MsgId() int32 {
	return 9723
}

func (*MountainSeaInviteRsp) Comment() string {
	return "邀请"
}
		
func (*LeaveMountainSeaReq) MsgId() int32 {
	return 209724
}

func (*LeaveMountainSeaReq) Comment() string {
	return "离开系统"
}
		
func (*LeaveMountainSeaRsp) MsgId() int32 {
	return 9724
}

func (*LeaveMountainSeaRsp) Comment() string {
	return "离开系统"
}
		
func (*MountainSeaGetBossPowerReq) MsgId() int32 {
	return 209725
}

func (*MountainSeaGetBossPowerReq) Comment() string {
	return "获取boss妖力"
}
		
func (*MountainSeaGetBossPowerRsp) MsgId() int32 {
	return 9725
}

func (*MountainSeaGetBossPowerRsp) Comment() string {
	return "获取boss妖力"
}
		
func (*MountainSeaEnterBattleReq) MsgId() int32 {
	return 209750
}

func (*MountainSeaEnterBattleReq) Comment() string {
	return "进入布阵界面"
}
		
func (*MountainSeaEnterBattleResp) MsgId() int32 {
	return 9750
}

func (*MountainSeaEnterBattleResp) Comment() string {
	return "进入布阵界面"
}
		
func (*MountainSeaEnterSwitchSeparationReq) MsgId() int32 {
	return 209752
}

func (*MountainSeaEnterSwitchSeparationReq) Comment() string {
	return "进入切换分身界面"
}
		
func (*MountainSeaEnterSwitchSeparationRsp) MsgId() int32 {
	return 9752
}

func (*MountainSeaEnterSwitchSeparationRsp) Comment() string {
	return "进入切换分身界面"
}
		
func (*MountainSeaSwitchSeparationReq) MsgId() int32 {
	return 209753
}

func (*MountainSeaSwitchSeparationReq) Comment() string {
	return "切换上阵的分身"
}
		
func (*MountainSeaSwitchSeparationRsp) MsgId() int32 {
	return 9753
}

func (*MountainSeaSwitchSeparationRsp) Comment() string {
	return "切换上阵的分身"
}
		
func (*MountainSeaSeparationDetailReq) MsgId() int32 {
	return 209754
}

func (*MountainSeaSeparationDetailReq) Comment() string {
	return "查看分身详细数据"
}
		
func (*MountainSeaSeparationDetailRsp) MsgId() int32 {
	return 9754
}

func (*MountainSeaSeparationDetailRsp) Comment() string {
	return "查看分身详细数据"
}
		
func (*MountainSeaChangePosReq) MsgId() int32 {
	return 209755
}

func (*MountainSeaChangePosReq) Comment() string {
	return "改变参战对象位置"
}
		
func (*MountainSeaChangePosRsp) MsgId() int32 {
	return 9755
}

func (*MountainSeaChangePosRsp) Comment() string {
	return "改变参战对象位置"
}
		
func (*MountainSeaChangeTeamSkillReq) MsgId() int32 {
	return 209756
}

func (*MountainSeaChangeTeamSkillReq) Comment() string {
	return "改变装配的队伍技能"
}
		
func (*MountainSeaChangeTeamSkillRsp) MsgId() int32 {
	return 9756
}

func (*MountainSeaChangeTeamSkillRsp) Comment() string {
	return "改变装配的队伍技能"
}
		
func (*MountainSeaDoBattleReq) MsgId() int32 {
	return 209757
}

func (*MountainSeaDoBattleReq) Comment() string {
	return "进行战斗"
}
		
func (*MountainSeaDoBattleRsp) MsgId() int32 {
	return 9757
}

func (*MountainSeaDoBattleRsp) Comment() string {
	return "进行战斗"
}
		
func (*MountainSeaInviteListReq) MsgId() int32 {
	return 209783
}

func (*MountainSeaInviteListReq) Comment() string {
	return "邀请列表"
}
		
func (*MountainSeaInviteListResp) MsgId() int32 {
	return 9783
}

func (*MountainSeaInviteListResp) Comment() string {
	return "邀请列表"
}
		
func (*MountainSeaInviteRefuseReq) MsgId() int32 {
	return 209785
}

func (*MountainSeaInviteRefuseReq) Comment() string {
	return "一键拒绝邀请队伍"
}
		
func (*MountainSeaInviteRefuseResp) MsgId() int32 {
	return 9785
}

func (*MountainSeaInviteRefuseResp) Comment() string {
	return "一键拒绝邀请队伍"
}
		
func (*MountainSeaSetAppointReq) MsgId() int32 {
	return 209786
}

func (*MountainSeaSetAppointReq) Comment() string {
	return "设置接受指定邀请开关"
}
		
func (*MountainSeaSetAppointResp) MsgId() int32 {
	return 9786
}

func (*MountainSeaSetAppointResp) Comment() string {
	return "设置接受指定邀请开关"
}
		
func (*MountainSeaApplyJoinTeamReq) MsgId() int32 {
	return 209787
}

func (*MountainSeaApplyJoinTeamReq) Comment() string {
	return "加入指定邀请队伍"
}
		
func (*MountainSeaApplyJoinTeamRsp) MsgId() int32 {
	return 9787
}

func (*MountainSeaApplyJoinTeamRsp) Comment() string {
	return "加入指定邀请队伍"
}
		
func (*ReqGetConditionReward) MsgId() int32 {
	return 209901
}

func (*ReqGetConditionReward) Comment() string {
	return "疯狂聚宝盆签到"
}
		
func (*RspGetConditionReward) MsgId() int32 {
	return 9901
}

func (*RspGetConditionReward) Comment() string {
	return "疯狂聚宝盆签到"
}
		
func (*RegressionShareReq) MsgId() int32 {
	return 210001
}

func (*RegressionShareReq) Comment() string {
	return "仙友回归-首次分享领奖"
}
		
func (*RegressionShareResp) MsgId() int32 {
	return 10001
}

func (*RegressionShareResp) Comment() string {
	return "仙友回归-首次分享领奖"
}
		
func (*RegressionLotteryReq) MsgId() int32 {
	return 210002
}

func (*RegressionLotteryReq) Comment() string {
	return "仙友回归-抽奖"
}
		
func (*RegressionLotteryResp) MsgId() int32 {
	return 10002
}

func (*RegressionLotteryResp) Comment() string {
	return "仙友回归-抽奖"
}
		
func (*GetRegressionPlayerDataReq) MsgId() int32 {
	return 210003
}

func (*GetRegressionPlayerDataReq) Comment() string {
	return "仙友回归-绑定回归玩家数据"
}
		
func (*GetRegressionPlayerDataResp) MsgId() int32 {
	return 10003
}

func (*GetRegressionPlayerDataResp) Comment() string {
	return "仙友回归-绑定回归玩家数据"
}
		
func (*GetRegressionReceiveRewardReq) MsgId() int32 {
	return 210004
}

func (*GetRegressionReceiveRewardReq) Comment() string {
	return "仙友回归-领取绑定玩家的任务奖励"
}
		
func (*GetRegressionReceiveRewardResp) MsgId() int32 {
	return 10004
}

func (*GetRegressionReceiveRewardResp) Comment() string {
	return "仙友回归-领取绑定玩家的任务奖励"
}
		
func (*RegressionSaveSelectItemReq) MsgId() int32 {
	return 210005
}

func (*RegressionSaveSelectItemReq) Comment() string {
	return "仙友回归-保存自选奖励"
}
		
func (*RegressionSaveSelectItemResp) MsgId() int32 {
	return 10005
}

func (*RegressionSaveSelectItemResp) Comment() string {
	return "仙友回归-保存自选奖励"
}
		
func (*UnionBlessingSendGiftReq) MsgId() int32 {
	return 210601
}

func (*UnionBlessingSendGiftReq) Comment() string {
	return "妖盟赐福"
}
		
func (*UnionBlessingSendGiftResp) MsgId() int32 {
	return 10601
}

func (*UnionBlessingSendGiftResp) Comment() string {
	return "妖盟赐福"
}
		
func (*UnionBlessingRewardReqMsg) MsgId() int32 {
	return 210602
}

func (*UnionBlessingRewardReqMsg) Comment() string {
	return "领取赐福奖励"
}
		
func (*UnionBlessingRewardRespMsg) MsgId() int32 {
	return 10602
}

func (*UnionBlessingRewardRespMsg) Comment() string {
	return "领取赐福奖励"
}
		
func (*GoodFortuneGetRewardReq) MsgId() int32 {
	return 210701
}

func (*GoodFortuneGetRewardReq) Comment() string {
	return "获取签到奖励"
}
		
func (*GoodFortuneGetRewardRsp) MsgId() int32 {
	return 10701
}

func (*GoodFortuneGetRewardRsp) Comment() string {
	return "获取签到奖励"
}
		
func (*WestTravelPassGameReq) MsgId() int32 {
	return 210901
}

func (*WestTravelPassGameReq) Comment() string {
	return "通关请求"
}
		
func (*WestTravelPassGameResp) MsgId() int32 {
	return 10901
}

func (*WestTravelPassGameResp) Comment() string {
	return "通关请求"
}
		
func (*FestivalCelebrationsUseBanquetItemReq) MsgId() int32 {
	return 211001
}

func (*FestivalCelebrationsUseBanquetItemReq) Comment() string {
	return "使用宴会道具"
}
		
func (*FestivalCelebrationsUseBanquetItemRsp) MsgId() int32 {
	return 11001
}

func (*FestivalCelebrationsUseBanquetItemRsp) Comment() string {
	return "使用宴会道具"
}
		
func (*FestivalCelebrationsDrawBanquetRewardReq) MsgId() int32 {
	return 211002
}

func (*FestivalCelebrationsDrawBanquetRewardReq) Comment() string {
	return "领取宴会奖励"
}
		
func (*FestivalCelebrationsDrawBanquetRewardRsp) MsgId() int32 {
	return 11002
}

func (*FestivalCelebrationsDrawBanquetRewardRsp) Comment() string {
	return "领取宴会奖励"
}
		
func (*FestivalCelebrationsGetBanquetScoreDetailReq) MsgId() int32 {
	return 211003
}

func (*FestivalCelebrationsGetBanquetScoreDetailReq) Comment() string {
	return "获取积分详情"
}
		
func (*FestivalCelebrationsGetBanquetScoreDetailRsp) MsgId() int32 {
	return 11003
}

func (*FestivalCelebrationsGetBanquetScoreDetailRsp) Comment() string {
	return "获取积分详情"
}
		
func (*FestivalCelebrationsUseLuckyFateItemReq) MsgId() int32 {
	return 211004
}

func (*FestivalCelebrationsUseLuckyFateItemReq) Comment() string {
	return "使用福缘道具"
}
		
func (*FestivalCelebrationsUseLuckyFateItemRsp) MsgId() int32 {
	return 11004
}

func (*FestivalCelebrationsUseLuckyFateItemRsp) Comment() string {
	return "使用福缘道具"
}
		
func (*FestivalCelebrationsGetSignRewardReq) MsgId() int32 {
	return 211005
}

func (*FestivalCelebrationsGetSignRewardReq) Comment() string {
	return "获取签到奖励"
}
		
func (*FestivalCelebrationsGetSignRewardRsp) MsgId() int32 {
	return 11005
}

func (*FestivalCelebrationsGetSignRewardRsp) Comment() string {
	return "获取签到奖励"
}
		
func (*FestivalCelebrationsCollectSynthesisReq) MsgId() int32 {
	return 211006
}

func (*FestivalCelebrationsCollectSynthesisReq) Comment() string {
	return "请求合成"
}
		
func (*FestivalCelebrationsCollectSynthesisRsp) MsgId() int32 {
	return 11006
}

func (*FestivalCelebrationsCollectSynthesisRsp) Comment() string {
	return "请求合成"
}
		
func (*FestivalCelebrationsCollectFillReq) MsgId() int32 {
	return 211007
}

func (*FestivalCelebrationsCollectFillReq) Comment() string {
	return "请求填充"
}
		
func (*FestivalCelebrationsCollectFillRsp) MsgId() int32 {
	return 11007
}

func (*FestivalCelebrationsCollectFillRsp) Comment() string {
	return "请求填充"
}
		
func (*FestivalCelebrationsCollectDrawBigRewardReq) MsgId() int32 {
	return 211008
}

func (*FestivalCelebrationsCollectDrawBigRewardReq) Comment() string {
	return "领取大奖"
}
		
func (*FestivalCelebrationsCollectDrawBigRewardRsp) MsgId() int32 {
	return 11008
}

func (*FestivalCelebrationsCollectDrawBigRewardRsp) Comment() string {
	return "领取大奖"
}
		
func (*FestivalCelebrationsCollectGetClaimListReq) MsgId() int32 {
	return 211009
}

func (*FestivalCelebrationsCollectGetClaimListReq) Comment() string {
	return "获取领取名单"
}
		
func (*FestivalCelebrationsCollectGetClaimListRsp) MsgId() int32 {
	return 11009
}

func (*FestivalCelebrationsCollectGetClaimListRsp) Comment() string {
	return "获取领取名单"
}
		
func (*FestivalCelebrationsCollectSearchPlayerReq) MsgId() int32 {
	return 211010
}

func (*FestivalCelebrationsCollectSearchPlayerReq) Comment() string {
	return "搜索玩家"
}
		
func (*FestivalCelebrationsCollectSearchPlayerRsp) MsgId() int32 {
	return 11010
}

func (*FestivalCelebrationsCollectSearchPlayerRsp) Comment() string {
	return "搜索玩家"
}
		
func (*FestivalCelebrationsCollectGiveReq) MsgId() int32 {
	return 211011
}

func (*FestivalCelebrationsCollectGiveReq) Comment() string {
	return "赠送道具"
}
		
func (*FestivalCelebrationsCollectGiveRsp) MsgId() int32 {
	return 11011
}

func (*FestivalCelebrationsCollectGiveRsp) Comment() string {
	return "赠送道具"
}
		
func (*FestivalCelebrationsCollectAskForReq) MsgId() int32 {
	return 211012
}

func (*FestivalCelebrationsCollectAskForReq) Comment() string {
	return "索要道具"
}
		
func (*FestivalCelebrationsCollectAskForRsp) MsgId() int32 {
	return 11012
}

func (*FestivalCelebrationsCollectAskForRsp) Comment() string {
	return "索要道具"
}
		
func (*FestivalCelebrationsEasterEggGetRewardReq) MsgId() int32 {
	return 211013
}

func (*FestivalCelebrationsEasterEggGetRewardReq) Comment() string {
	return "领取彩蛋奖励"
}
		
func (*FestivalCelebrationsEasterEggGetRewardRsp) MsgId() int32 {
	return 11013
}

func (*FestivalCelebrationsEasterEggGetRewardRsp) Comment() string {
	return "领取彩蛋奖励"
}
		
func (*FestivalCelebrationsCollectGetGivenReq) MsgId() int32 {
	return 211014
}

func (*FestivalCelebrationsCollectGetGivenReq) Comment() string {
	return "领取赠送五福"
}
		
func (*FestivalCelebrationsCollectGetGivenResp) MsgId() int32 {
	return 11014
}

func (*FestivalCelebrationsCollectGetGivenResp) Comment() string {
	return "领取赠送五福"
}
		
func (*FestivalCelebrationsFuYuanGetRewardReq) MsgId() int32 {
	return 211015
}

func (*FestivalCelebrationsFuYuanGetRewardReq) Comment() string {
	return "领取福源奖励"
}
		
func (*FestivalCelebrationsFuYuanGetRewardResp) MsgId() int32 {
	return 11015
}

func (*FestivalCelebrationsFuYuanGetRewardResp) Comment() string {
	return "领取福源奖励"
}
		
func (*FestivalCelebrationsGetBindUnionMemberDataListReq) MsgId() int32 {
	return 211016
}

func (*FestivalCelebrationsGetBindUnionMemberDataListReq) Comment() string {
	return "获取绑定妖盟成员列表"
}
		
func (*FestivalCelebrationsGetBindUnionMemberDataListRsp) MsgId() int32 {
	return 11016
}

func (*FestivalCelebrationsGetBindUnionMemberDataListRsp) Comment() string {
	return "获取绑定妖盟成员列表"
}
		
func (*FestivalCelebrationsGetAnnualMemoryReq) MsgId() int32 {
	return 211021
}

func (*FestivalCelebrationsGetAnnualMemoryReq) Comment() string {
	return "获取周年回忆条目"
}
		
func (*FestivalCelebrationsGetAnnualMemoryRsp) MsgId() int32 {
	return 11021
}

func (*FestivalCelebrationsGetAnnualMemoryRsp) Comment() string {
	return "获取周年回忆条目"
}
		
func (*FestivalCelebrationsWorshipReq) MsgId() int32 {
	return 211022
}

func (*FestivalCelebrationsWorshipReq) Comment() string {
	return "膜拜"
}
		
func (*FestivalCelebrationsWorshipResp) MsgId() int32 {
	return 11022
}

func (*FestivalCelebrationsWorshipResp) Comment() string {
	return "膜拜"
}
		
func (*DoubleDemonsAppointInviteReq) MsgId() int32 {
	return 211106
}

func (*DoubleDemonsAppointInviteReq) Comment() string {
	return "指定邀请"
}
		
func (*DoubleDemonsAppointInviteResp) MsgId() int32 {
	return 11106
}

func (*DoubleDemonsAppointInviteResp) Comment() string {
	return "指定邀请"
}
		
func (*DoubleDemonsAgreeReq) MsgId() int32 {
	return 211112
}

func (*DoubleDemonsAgreeReq) Comment() string {
	return "同意邀请/结为伴侣"
}
		
func (*DoubleDemonsAgreeResp) MsgId() int32 {
	return 11112
}

func (*DoubleDemonsAgreeResp) Comment() string {
	return "同意邀请/结为伴侣"
}
		
func (*DoubleDemonsRefuseReq) MsgId() int32 {
	return 211113
}

func (*DoubleDemonsRefuseReq) Comment() string {
	return "一键拒绝"
}
		
func (*DoubleDemonsRefuseResp) MsgId() int32 {
	return 11113
}

func (*DoubleDemonsRefuseResp) Comment() string {
	return "一键拒绝"
}
		
func (*DoubleDemonsReceivePresentReq) MsgId() int32 {
	return 211114
}

func (*DoubleDemonsReceivePresentReq) Comment() string {
	return "领取赠送奖励"
}
		
func (*DoubleDemonsReceivePresentResp) MsgId() int32 {
	return 11114
}

func (*DoubleDemonsReceivePresentResp) Comment() string {
	return "领取赠送奖励"
}
		
func (*NewYearRedBagEnterReqMsg) MsgId() int32 {
	return 211501
}

func (*NewYearRedBagEnterReqMsg) Comment() string {
	return "新春红包 进入系统"
}
		
func (*NewYearRedBagEnterRespMsg) MsgId() int32 {
	return 11501
}

func (*NewYearRedBagEnterRespMsg) Comment() string {
	return "新春红包 进入系统"
}
		
func (*NewYearRedBagExitReqMsg) MsgId() int32 {
	return 211502
}

func (*NewYearRedBagExitReqMsg) Comment() string {
	return "新春红包 退出系统"
}
		
func (*NewYearRedBagExitRespMsg) MsgId() int32 {
	return 11502
}

func (*NewYearRedBagExitRespMsg) Comment() string {
	return "新春红包 退出系统"
}
		
func (*NewYearRedBagOpenReqMsg) MsgId() int32 {
	return 211503
}

func (*NewYearRedBagOpenReqMsg) Comment() string {
	return "新春红包 打开红包"
}
		
func (*NewYearRedBagOpenRespMsg) MsgId() int32 {
	return 11503
}

func (*NewYearRedBagOpenRespMsg) Comment() string {
	return "新春红包 打开红包"
}
		
func (*PetKerneCarryReq) MsgId() int32 {
	return 211702
}

func (*PetKerneCarryReq) Comment() string {
	return "穿戴内丹"
}
		
func (*PetKerneCarryResp) MsgId() int32 {
	return 11702
}

func (*PetKerneCarryResp) Comment() string {
	return "穿戴内丹"
}
		
func (*PetKerneUpLevelReq) MsgId() int32 {
	return 211703
}

func (*PetKerneUpLevelReq) Comment() string {
	return "内丹升级"
}
		
func (*PetKerneUpLevelResp) MsgId() int32 {
	return 11703
}

func (*PetKerneUpLevelResp) Comment() string {
	return "内丹升级"
}
		
func (*PetKerneUpStarReq) MsgId() int32 {
	return 211704
}

func (*PetKerneUpStarReq) Comment() string {
	return "内丹升星"
}
		
func (*PetKerneUpStarResp) MsgId() int32 {
	return 11704
}

func (*PetKerneUpStarResp) Comment() string {
	return "内丹升星"
}
		
func (*PetKerneActiveReq) MsgId() int32 {
	return 211705
}

func (*PetKerneActiveReq) Comment() string {
	return "内丹激活"
}
		
func (*PetKerneActiveResp) MsgId() int32 {
	return 11705
}

func (*PetKerneActiveResp) Comment() string {
	return "内丹激活"
}
		
func (*PetKernelDrawReq) MsgId() int32 {
	return 211706
}

func (*PetKernelDrawReq) Comment() string {
	return "抽取灵兽内丹"
}
		
func (*PetKernelDrawResp) MsgId() int32 {
	return 11706
}

func (*PetKernelDrawResp) Comment() string {
	return "抽取灵兽内丹"
}
		
func (*PetKerneCombineUpLevelReq) MsgId() int32 {
	return 211707
}

func (*PetKerneCombineUpLevelReq) Comment() string {
	return "内丹共鸣升级或激活"
}
		
func (*PetKerneCombineUpLevelResp) MsgId() int32 {
	return 11707
}

func (*PetKerneCombineUpLevelResp) Comment() string {
	return "内丹共鸣升级或激活"
}
		
func (*EnterPupilSystemReq) MsgId() int32 {
	return 211801
}

func (*EnterPupilSystemReq) Comment() string {
	return "进入门徒系统"
}
		
func (*EnterPupilSystemResp) MsgId() int32 {
	return 11801
}

func (*EnterPupilSystemResp) Comment() string {
	return "进入门徒系统"
}
		
func (*PupilTrainReq) MsgId() int32 {
	return 211802
}

func (*PupilTrainReq) Comment() string {
	return "门徒培养"
}
		
func (*PupilTrainResp) MsgId() int32 {
	return 11802
}

func (*PupilTrainResp) Comment() string {
	return "门徒培养"
}
		
func (*PupilSiteTrainTimesRecoverReq) MsgId() int32 {
	return 211803
}

func (*PupilSiteTrainTimesRecoverReq) Comment() string {
	return "恢复培养次数"
}
		
func (*PupilSiteTrainTimesRecoverResp) MsgId() int32 {
	return 11803
}

func (*PupilSiteTrainTimesRecoverResp) Comment() string {
	return "恢复培养次数"
}
		
func (*PupilRecruitReq) MsgId() int32 {
	return 211805
}

func (*PupilRecruitReq) Comment() string {
	return "招收门徒"
}
		
func (*PupilRecruitResp) MsgId() int32 {
	return 11805
}

func (*PupilRecruitResp) Comment() string {
	return "招收门徒"
}
		
func (*PupilGraduateReq) MsgId() int32 {
	return 211806
}

func (*PupilGraduateReq) Comment() string {
	return "门徒出师"
}
		
func (*PupilGraduateResp) MsgId() int32 {
	return 11806
}

func (*PupilGraduateResp) Comment() string {
	return "门徒出师"
}
		
func (*PupilGetGraduatedListReq) MsgId() int32 {
	return 211807
}

func (*PupilGetGraduatedListReq) Comment() string {
	return "获取出师门徒列表"
}
		
func (*PupilGetGraduatedListResp) MsgId() int32 {
	return 11807
}

func (*PupilGetGraduatedListResp) Comment() string {
	return "获取出师门徒列表"
}
		
func (*PupilGetFightListReq) MsgId() int32 {
	return 211808
}

func (*PupilGetFightListReq) Comment() string {
	return "获取上阵列表"
}
		
func (*PupilGetFightListResp) MsgId() int32 {
	return 11808
}

func (*PupilGetFightListResp) Comment() string {
	return "获取上阵列表"
}
		
func (*PupilFightOnReq) MsgId() int32 {
	return 211809
}

func (*PupilFightOnReq) Comment() string {
	return "门徒上阵"
}
		
func (*PupilFightOnResp) MsgId() int32 {
	return 11809
}

func (*PupilFightOnResp) Comment() string {
	return "门徒上阵"
}
		
func (*PupilGetSectInfoReq) MsgId() int32 {
	return 211810
}

func (*PupilGetSectInfoReq) Comment() string {
	return "获取门派统计信息"
}
		
func (*PupilGetSectInfoResp) MsgId() int32 {
	return 11810
}

func (*PupilGetSectInfoResp) Comment() string {
	return "获取门派统计信息"
}
		
func (*SearchPlayerReq) MsgId() int32 {
	return 211811
}

func (*SearchPlayerReq) Comment() string {
	return "搜索玩家"
}
		
func (*SearchPlayerResp) MsgId() int32 {
	return 11811
}

func (*SearchPlayerResp) Comment() string {
	return "搜索玩家"
}
		
func (*PupilGetAdRewardReq) MsgId() int32 {
	return 211814
}

func (*PupilGetAdRewardReq) Comment() string {
	return "领取广告奖励"
}
		
func (*PupilGetAdRewardResp) MsgId() int32 {
	return 11814
}

func (*PupilGetAdRewardResp) Comment() string {
	return "领取广告奖励"
}
		
func (*PupilLockReq) MsgId() int32 {
	return 211815
}

func (*PupilLockReq) Comment() string {
	return "锁定弟子"
}
		
func (*PupilLockResp) MsgId() int32 {
	return 11815
}

func (*PupilLockResp) Comment() string {
	return "锁定弟子"
}
		
func (*PupilExactSearchReq) MsgId() int32 {
	return 211816
}

func (*PupilExactSearchReq) Comment() string {
	return "精准查找"
}
		
func (*PupilExactSearchResp) MsgId() int32 {
	return 11816
}

func (*PupilExactSearchResp) Comment() string {
	return "精准查找"
}
		
func (*PupilGraduatedUnMarryListReq) MsgId() int32 {
	return 211855
}

func (*PupilGraduatedUnMarryListReq) Comment() string {
	return "获得出师未结义的门徒列表"
}
		
func (*PupilGraduatedUnMarryListResp) MsgId() int32 {
	return 11855
}

func (*PupilGraduatedUnMarryListResp) Comment() string {
	return "获得出师未结义的门徒列表"
}
		
func (*MarriageRecordListReq) MsgId() int32 {
	return 211856
}

func (*MarriageRecordListReq) Comment() string {
	return "结义记录列表"
}
		
func (*MarriageRecordListResp) MsgId() int32 {
	return 11856
}

func (*MarriageRecordListResp) Comment() string {
	return "结义记录列表"
}
		
func (*MarriageGetServerApplyReq) MsgId() int32 {
	return 211857
}

func (*MarriageGetServerApplyReq) Comment() string {
	return "获取全服结义请求列表(包括妖盟)"
}
		
func (*MarriageGetServerApplyResp) MsgId() int32 {
	return 11857
}

func (*MarriageGetServerApplyResp) Comment() string {
	return "获取全服结义请求列表(包括妖盟)"
}
		
func (*MarriageRecommendPlayerReq) MsgId() int32 {
	return 211858
}

func (*MarriageRecommendPlayerReq) Comment() string {
	return "指定结义推荐玩家请求"
}
		
func (*MarriageRecommendPlayerResp) MsgId() int32 {
	return 11858
}

func (*MarriageRecommendPlayerResp) Comment() string {
	return "指定结义推荐玩家请求"
}
		
func (*MarriageGetAppointApplyReq) MsgId() int32 {
	return 211859
}

func (*MarriageGetAppointApplyReq) Comment() string {
	return "获取其他玩家发起的结义请求（别的玩家指定你结义的列表）"
}
		
func (*MarriageGetAppointApplyResp) MsgId() int32 {
	return 11859
}

func (*MarriageGetAppointApplyResp) Comment() string {
	return "获取其他玩家发起的结义请求（别的玩家指定你结义的列表）"
}
		
func (*MarriageApplyDealReq) MsgId() int32 {
	return 211860
}

func (*MarriageApplyDealReq) Comment() string {
	return "处理结义请求"
}
		
func (*MarriageApplyDealResp) MsgId() int32 {
	return 11860
}

func (*MarriageApplyDealResp) Comment() string {
	return "处理结义请求"
}
		
func (*MarriageApplyPublishReq) MsgId() int32 {
	return 211861
}

func (*MarriageApplyPublishReq) Comment() string {
	return "发布结义请求"
}
		
func (*MarriageApplyPublishResp) MsgId() int32 {
	return 11861
}

func (*MarriageApplyPublishResp) Comment() string {
	return "发布结义请求"
}
		
func (*MarriageApplyCancelReq) MsgId() int32 {
	return 211862
}

func (*MarriageApplyCancelReq) Comment() string {
	return "取消结义请求"
}
		
func (*MarriageApplyCancelResp) MsgId() int32 {
	return 11862
}

func (*MarriageApplyCancelResp) Comment() string {
	return "取消结义请求"
}
		
func (*MarriageRefreshServerApplyReq) MsgId() int32 {
	return 211863
}

func (*MarriageRefreshServerApplyReq) Comment() string {
	return "刷新全服结义列表请求"
}
		
func (*MarriageRefreshServerApplyResp) MsgId() int32 {
	return 11863
}

func (*MarriageRefreshServerApplyResp) Comment() string {
	return "刷新全服结义列表请求"
}
		
func (*MarriageSetAppointReq) MsgId() int32 {
	return 211864
}

func (*MarriageSetAppointReq) Comment() string {
	return "设置是否接受指向结义请求"
}
		
func (*MarriageSetAppointResp) MsgId() int32 {
	return 11864
}

func (*MarriageSetAppointResp) Comment() string {
	return "设置是否接受指向结义请求"
}
		
func (*PetRootUpReq) MsgId() int32 {
	return 212201
}

func (*PetRootUpReq) Comment() string {
	return "洪荒灵兽灵根升级请求"
}
		
func (*PetRootUpResp) MsgId() int32 {
	return 12201
}

func (*PetRootUpResp) Comment() string {
	return "洪荒灵兽灵根升级请求"
}
		
func (*PetAwakeReq) MsgId() int32 {
	return 212202
}

func (*PetAwakeReq) Comment() string {
	return "洪荒灵兽觉醒请求"
}
		
func (*PetAwakeResp) MsgId() int32 {
	return 12202
}

func (*PetAwakeResp) Comment() string {
	return "洪荒灵兽觉醒请求"
}
		
func (*PetFateUpReq) MsgId() int32 {
	return 212203
}

func (*PetFateUpReq) Comment() string {
	return "洪荒灵兽缘分升级请求"
}
		
func (*PetFateUpResp) MsgId() int32 {
	return 12203
}

func (*PetFateUpResp) Comment() string {
	return "洪荒灵兽缘分升级请求"
}
		
func (*PetSwitchSkinReq) MsgId() int32 {
	return 212204
}

func (*PetSwitchSkinReq) Comment() string {
	return "切换灵兽皮肤"
}
		
func (*PetSwitchSkinResp) MsgId() int32 {
	return 12204
}

func (*PetSwitchSkinResp) Comment() string {
	return "切换灵兽皮肤"
}
		
func (*PetUpSkinLevelReq) MsgId() int32 {
	return 212205
}

func (*PetUpSkinLevelReq) Comment() string {
	return "激活、升级灵兽皮肤"
}
		
func (*PetUpSkinLevelResp) MsgId() int32 {
	return 12205
}

func (*PetUpSkinLevelResp) Comment() string {
	return "激活、升级灵兽皮肤"
}
		
func (*PetUpSkinCombineReq) MsgId() int32 {
	return 212206
}

func (*PetUpSkinCombineReq) Comment() string {
	return "激活、升级灵兽羁绊"
}
		
func (*PetUpSkinCombineResp) MsgId() int32 {
	return 12206
}

func (*PetUpSkinCombineResp) Comment() string {
	return "激活、升级灵兽羁绊"
}
		
func (*GetWechatRankWhiteListReq) MsgId() int32 {
	return 212601
}

func (*GetWechatRankWhiteListReq) Comment() string {
	return "获取微信排行榜白名单"
}
		
func (*GetWechatRankWhiteListRsp) MsgId() int32 {
	return 12601
}

func (*GetWechatRankWhiteListRsp) Comment() string {
	return "获取微信排行榜白名单"
}
		
func (*CloudRefineReq) MsgId() int32 {
	return 212901
}

func (*CloudRefineReq) Comment() string {
	return "筋斗云注灵请求"
}
		
func (*CloudRefineResp) MsgId() int32 {
	return 12901
}

func (*CloudRefineResp) Comment() string {
	return "筋斗云注灵请求"
}
		
func (*CloudRefineStarLvUpReq) MsgId() int32 {
	return 212902
}

func (*CloudRefineStarLvUpReq) Comment() string {
	return "筋斗云注灵渡劫请求"
}
		
func (*CloudRefineStarLvUpResp) MsgId() int32 {
	return 12902
}

func (*CloudRefineStarLvUpResp) Comment() string {
	return "筋斗云注灵渡劫请求"
}
		
func (*ComposeBallEnterGameReq) MsgId() int32 {
	return 213001
}

func (*ComposeBallEnterGameReq) Comment() string {
	return "进入游戏"
}
		
func (*ComposeBallEnterGameResp) MsgId() int32 {
	return 13001
}

func (*ComposeBallEnterGameResp) Comment() string {
	return "进入游戏"
}
		
func (*ComposeBallComposeReq) MsgId() int32 {
	return 213002
}

func (*ComposeBallComposeReq) Comment() string {
	return "请求合成"
}
		
func (*ComposeBallComposeResp) MsgId() int32 {
	return 13002
}

func (*ComposeBallComposeResp) Comment() string {
	return "请求合成"
}
		
func (*ComposeBallUseItemReq) MsgId() int32 {
	return 213003
}

func (*ComposeBallUseItemReq) Comment() string {
	return "请求使用道具"
}
		
func (*ComposeBallUseItemResp) MsgId() int32 {
	return 13003
}

func (*ComposeBallUseItemResp) Comment() string {
	return "请求使用道具"
}
		
func (*ComposeBallHpItemReq) MsgId() int32 {
	return 213004
}

func (*ComposeBallHpItemReq) Comment() string {
	return "请求增加体力"
}
		
func (*ComposeBallHpItemResp) MsgId() int32 {
	return 13004
}

func (*ComposeBallHpItemResp) Comment() string {
	return "请求增加体力"
}
		
func (*ComposeBallSettleReq) MsgId() int32 {
	return 213005
}

func (*ComposeBallSettleReq) Comment() string {
	return "请求结算"
}
		
func (*ComposeBallSettleResp) MsgId() int32 {
	return 13005
}

func (*ComposeBallSettleResp) Comment() string {
	return "请求结算"
}
		
func (*MonopolyEnterActivityReq) MsgId() int32 {
	return 213101
}

func (*MonopolyEnterActivityReq) Comment() string {
	return "进入活动"
}
		
func (*MonopolyEnterActivityResp) MsgId() int32 {
	return 13101
}

func (*MonopolyEnterActivityResp) Comment() string {
	return "进入活动"
}
		
func (*MonopolyEnterMapReq) MsgId() int32 {
	return 213102
}

func (*MonopolyEnterMapReq) Comment() string {
	return "进入地图"
}
		
func (*MonopolyEnterMapResp) MsgId() int32 {
	return 13102
}

func (*MonopolyEnterMapResp) Comment() string {
	return "进入地图"
}
		
func (*MonopolyRollDiceReq) MsgId() int32 {
	return 213103
}

func (*MonopolyRollDiceReq) Comment() string {
	return "摇色子"
}
		
func (*MonopolyRollDiceResp) MsgId() int32 {
	return 13103
}

func (*MonopolyRollDiceResp) Comment() string {
	return "摇色子"
}
		
func (*MonopolyAssistListReq) MsgId() int32 {
	return 213104
}

func (*MonopolyAssistListReq) Comment() string {
	return "获取协助列表"
}
		
func (*MonopolyAssistListResp) MsgId() int32 {
	return 13104
}

func (*MonopolyAssistListResp) Comment() string {
	return "获取协助列表"
}
		
func (*MonopolyReplenishStrengthReq) MsgId() int32 {
	return 213105
}

func (*MonopolyReplenishStrengthReq) Comment() string {
	return "使用道具 补充体力"
}
		
func (*MonopolyReplenishStrengthResp) MsgId() int32 {
	return 13105
}

func (*MonopolyReplenishStrengthResp) Comment() string {
	return "使用道具 补充体力"
}
		
func (*MonopolyRescueTrapReq) MsgId() int32 {
	return 213107
}

func (*MonopolyRescueTrapReq) Comment() string {
	return "协助队友出陷阱"
}
		
func (*MonopolyRescueTrapResp) MsgId() int32 {
	return 13107
}

func (*MonopolyRescueTrapResp) Comment() string {
	return "协助队友出陷阱"
}
		
func (*MonopolyRobListReq) MsgId() int32 {
	return 213108
}

func (*MonopolyRobListReq) Comment() string {
	return "获取可掠夺的妖盟列表"
}
		
func (*MonopolyRobListResp) MsgId() int32 {
	return 13108
}

func (*MonopolyRobListResp) Comment() string {
	return "获取可掠夺的妖盟列表"
}
		
func (*MonopolyRobReq) MsgId() int32 {
	return 213110
}

func (*MonopolyRobReq) Comment() string {
	return "掠夺其他妖盟"
}
		
func (*MonopolyRobResp) MsgId() int32 {
	return 13110
}

func (*MonopolyRobResp) Comment() string {
	return "掠夺其他妖盟"
}
		
func (*MonopolyUnionLogListReq) MsgId() int32 {
	return 213111
}

func (*MonopolyUnionLogListReq) Comment() string {
	return "获取妖盟日志列表"
}
		
func (*MonopolyUnionLogListResp) MsgId() int32 {
	return 13111
}

func (*MonopolyUnionLogListResp) Comment() string {
	return "获取妖盟日志列表"
}
		
func (*MonopolyPlayerLogDetailReq) MsgId() int32 {
	return 213112
}

func (*MonopolyPlayerLogDetailReq) Comment() string {
	return "获取玩家日志详情"
}
		
func (*MonopolyPlayerLogDetailResp) MsgId() int32 {
	return 13112
}

func (*MonopolyPlayerLogDetailResp) Comment() string {
	return "获取玩家日志详情"
}
		
func (*MonopolyAssistAttackMonsterReq) MsgId() int32 {
	return 213113
}

func (*MonopolyAssistAttackMonsterReq) Comment() string {
	return "协助队友攻击怪物"
}
		
func (*MonopolyAssistAttackMonsterResp) MsgId() int32 {
	return 13113
}

func (*MonopolyAssistAttackMonsterResp) Comment() string {
	return "协助队友攻击怪物"
}
		
func (*MonopolyReceiveMonsterRewardReq) MsgId() int32 {
	return 213114
}

func (*MonopolyReceiveMonsterRewardReq) Comment() string {
	return "领取队友帮忙打死怪后的奖励"
}
		
func (*MonopolyReceiveMonsterRewardResp) MsgId() int32 {
	return 13114
}

func (*MonopolyReceiveMonsterRewardResp) Comment() string {
	return "领取队友帮忙打死怪后的奖励"
}
		
func (*MonopolyRedPointReqMsg) MsgId() int32 {
	return 213115
}

func (*MonopolyRedPointReqMsg) Comment() string {
	return "红点详情"
}
		
func (*MonopolyRedPointRespMsg) MsgId() int32 {
	return 13115
}

func (*MonopolyRedPointRespMsg) Comment() string {
	return "红点详情"
}
		
func (*MonopolyRemoteRollDiceReq) MsgId() int32 {
	return 213116
}

func (*MonopolyRemoteRollDiceReq) Comment() string {
	return "遥控骰子"
}
		
func (*MonopolyRemoteRollDiceResp) MsgId() int32 {
	return 13116
}

func (*MonopolyRemoteRollDiceResp) Comment() string {
	return "遥控骰子"
}
		
func (*MonopolyEnterRobMapReq) MsgId() int32 {
	return 213117
}

func (*MonopolyEnterRobMapReq) Comment() string {
	return "进入其他妖盟地图，请求详细数据"
}
		
func (*MonopolyEnterRobMapResp) MsgId() int32 {
	return 13117
}

func (*MonopolyEnterRobMapResp) Comment() string {
	return "进入其他妖盟地图，请求详细数据"
}
		
func (*MonopolyAutoUnlockReq) MsgId() int32 {
	return 213118
}

func (*MonopolyAutoUnlockReq) Comment() string {
	return "自动操作解锁"
}
		
func (*MonopolyAutoUnlockResp) MsgId() int32 {
	return 13118
}

func (*MonopolyAutoUnlockResp) Comment() string {
	return "自动操作解锁"
}
		
func (*MonopolyBlessingListReq) MsgId() int32 {
	return 213119
}

func (*MonopolyBlessingListReq) Comment() string {
	return "获取天赐福源列表"
}
		
func (*MonopolyBlessingListResp) MsgId() int32 {
	return 13119
}

func (*MonopolyBlessingListResp) Comment() string {
	return "获取天赐福源列表"
}
		
func (*MonopolyReceiveBlessingReq) MsgId() int32 {
	return 213120
}

func (*MonopolyReceiveBlessingReq) Comment() string {
	return "领取天赐祥瑞奖励"
}
		
func (*MonopolyReceiveBlessingResp) MsgId() int32 {
	return 13120
}

func (*MonopolyReceiveBlessingResp) Comment() string {
	return "领取天赐祥瑞奖励"
}
		
func (*MonopolyQuickMoveReq) MsgId() int32 {
	return 213121
}

func (*MonopolyQuickMoveReq) Comment() string {
	return "神行符移动"
}
		
func (*MonopolyQuickMoveResp) MsgId() int32 {
	return 13121
}

func (*MonopolyQuickMoveResp) Comment() string {
	return "神行符移动"
}
		
func (*MonopolyDarkGridMoveReq) MsgId() int32 {
	return 213122
}

func (*MonopolyDarkGridMoveReq) Comment() string {
	return "暗格移动"
}
		
func (*MonopolyDarkGridMoveResp) MsgId() int32 {
	return 13122
}

func (*MonopolyDarkGridMoveResp) Comment() string {
	return "暗格移动"
}
		
func (*MonopolyEventHandleReq) MsgId() int32 {
	return 213123
}

func (*MonopolyEventHandleReq) Comment() string {
	return "处理事件同一接口"
}
		
func (*MonopolyEventHandleResp) MsgId() int32 {
	return 13123
}

func (*MonopolyEventHandleResp) Comment() string {
	return "处理事件同一接口"
}
		
func (*MonopolyMonsterAttrReq) MsgId() int32 {
	return 213124
}

func (*MonopolyMonsterAttrReq) Comment() string {
	return "查看怪物详情"
}
		
func (*MonopolyMonsterAttrResp) MsgId() int32 {
	return 13124
}

func (*MonopolyMonsterAttrResp) Comment() string {
	return "查看怪物详情"
}
		
func (*MonopolyGetPlayerRankReq) MsgId() int32 {
	return 213125
}

func (*MonopolyGetPlayerRankReq) Comment() string {
	return "获取个人排名"
}
		
func (*MonopolyGetPlayerRankResp) MsgId() int32 {
	return 13125
}

func (*MonopolyGetPlayerRankResp) Comment() string {
	return "获取个人排名"
}
		
func (*MonopolyGetEnemyListReq) MsgId() int32 {
	return 213126
}

func (*MonopolyGetEnemyListReq) Comment() string {
	return "获取仇人列表"
}
		
func (*MonopolyGetEnemyListResp) MsgId() int32 {
	return 13126
}

func (*MonopolyGetEnemyListResp) Comment() string {
	return "获取仇人列表"
}
		
func (*MonopolyScoreDetailReq) MsgId() int32 {
	return 213127
}

func (*MonopolyScoreDetailReq) Comment() string {
	return "积分详情"
}
		
func (*MonopolyScoreDetailResp) MsgId() int32 {
	return 13127
}

func (*MonopolyScoreDetailResp) Comment() string {
	return "积分详情"
}
		
func (*MonopolyBuildingScoreDetailReq) MsgId() int32 {
	return 213128
}

func (*MonopolyBuildingScoreDetailReq) Comment() string {
	return "建筑积分详情"
}
		
func (*MonopolyBuildingScoreDetailResp) MsgId() int32 {
	return 13128
}

func (*MonopolyBuildingScoreDetailResp) Comment() string {
	return "建筑积分详情"
}
		
func (*MonopolyGetGroupDetailReq) MsgId() int32 {
	return 213141
}

func (*MonopolyGetGroupDetailReq) Comment() string {
	return "分组内积分详情"
}
		
func (*MonopolyGetGroupDetailResp) MsgId() int32 {
	return 13141
}

func (*MonopolyGetGroupDetailResp) Comment() string {
	return "分组内积分详情"
}
		
func (*MonopolyWorshipReq) MsgId() int32 {
	return 213142
}

func (*MonopolyWorshipReq) Comment() string {
	return "膜拜"
}
		
func (*MonopolyWorshipResp) MsgId() int32 {
	return 13142
}

func (*MonopolyWorshipResp) Comment() string {
	return "膜拜"
}
		
func (*MonopolyGuessPlayersReq) MsgId() int32 {
	return 213145
}

func (*MonopolyGuessPlayersReq) Comment() string {
	return "竞猜正确玩家数据"
}
		
func (*MonopolyGuessPlayersResp) MsgId() int32 {
	return 13145
}

func (*MonopolyGuessPlayersResp) Comment() string {
	return "竞猜正确玩家数据"
}
		
func (*MonopolyGetGuessDataReq) MsgId() int32 {
	return 213146
}

func (*MonopolyGetGuessDataReq) Comment() string {
	return "请求竞猜界面数据"
}
		
func (*MonopolyGetGuessDataResp) MsgId() int32 {
	return 13146
}

func (*MonopolyGetGuessDataResp) Comment() string {
	return "请求竞猜界面数据"
}
		
func (*MonopolyGuessSelectReq) MsgId() int32 {
	return 213147
}

func (*MonopolyGuessSelectReq) Comment() string {
	return "请求竞猜选择"
}
		
func (*MonopolyGuessSelectResp) MsgId() int32 {
	return 13147
}

func (*MonopolyGuessSelectResp) Comment() string {
	return "请求竞猜选择"
}
		
func (*MonopolyGuessRewardReq) MsgId() int32 {
	return 213148
}

func (*MonopolyGuessRewardReq) Comment() string {
	return "请求竞猜奖励"
}
		
func (*MonopolyGuessRewardResp) MsgId() int32 {
	return 13148
}

func (*MonopolyGuessRewardResp) Comment() string {
	return "请求竞猜奖励"
}
		
func (*TownDemonApplyDataSyncReq) MsgId() int32 {
	return 213250
}

func (*TownDemonApplyDataSyncReq) Comment() string {
	return "镇魔登录同步"
}
		
func (*TownDemonApplyDataSync) MsgId() int32 {
	return 13250
}

func (*TownDemonApplyDataSync) Comment() string {
	return "镇魔登录同步"
}
		
func (*UsePeachBanquetItemReq) MsgId() int32 {
	return 213520
}

func (*UsePeachBanquetItemReq) Comment() string {
	return "宴会道具使用"
}
		
func (*UsePeachBanquetItemResp) MsgId() int32 {
	return 13520
}

func (*UsePeachBanquetItemResp) Comment() string {
	return "宴会道具使用"
}
		
func (*UnionBountyEnterMapReq) MsgId() int32 {
	return 213602
}

func (*UnionBountyEnterMapReq) Comment() string {
	return "妖盟悬赏-打开护送地图"
}
		
func (*UnionBountyEnterMapResp) MsgId() int32 {
	return 13602
}

func (*UnionBountyEnterMapResp) Comment() string {
	return "妖盟悬赏-打开护送地图"
}
		
func (*UnionBountyExitMapReq) MsgId() int32 {
	return 213603
}

func (*UnionBountyExitMapReq) Comment() string {
	return "妖盟悬赏-离开护送地图"
}
		
func (*UnionBountyExitMapResp) MsgId() int32 {
	return 13603
}

func (*UnionBountyExitMapResp) Comment() string {
	return "妖盟悬赏-离开护送地图"
}
		
func (*UnionBountyPlunderReq) MsgId() int32 {
	return 213604
}

func (*UnionBountyPlunderReq) Comment() string {
	return "妖盟悬赏-掠夺某个镖车"
}
		
func (*UnionBountyPlunderResp) MsgId() int32 {
	return 13604
}

func (*UnionBountyPlunderResp) Comment() string {
	return "妖盟悬赏-掠夺某个镖车"
}
		
func (*UnionBountyDealBountyReq) MsgId() int32 {
	return 213605
}

func (*UnionBountyDealBountyReq) Comment() string {
	return "妖盟悬赏-处理悬赏事件"
}
		
func (*UnionBountyDealBountyResp) MsgId() int32 {
	return 13605
}

func (*UnionBountyDealBountyResp) Comment() string {
	return "妖盟悬赏-处理悬赏事件"
}
		
func (*UnionBountyGetRewardEscortReq) MsgId() int32 {
	return 213606
}

func (*UnionBountyGetRewardEscortReq) Comment() string {
	return "妖盟悬赏-领取护送镖车奖励"
}
		
func (*UnionBountyGetRewardEscortResp) MsgId() int32 {
	return 13606
}

func (*UnionBountyGetRewardEscortResp) Comment() string {
	return "妖盟悬赏-领取护送镖车奖励"
}
		
func (*UnionBountyCheckEscortReq) MsgId() int32 {
	return 213608
}

func (*UnionBountyCheckEscortReq) Comment() string {
	return "妖盟悬赏-查找某个玩家的镖车"
}
		
func (*UnionBountyCheckEscortResp) MsgId() int32 {
	return 13608
}

func (*UnionBountyCheckEscortResp) Comment() string {
	return "妖盟悬赏-查找某个玩家的镖车"
}
		
func (*UnionBountyGetReportReq) MsgId() int32 {
	return 213609
}

func (*UnionBountyGetReportReq) Comment() string {
	return "妖盟悬赏-打开战报"
}
		
func (*UnionBountyGetReportResp) MsgId() int32 {
	return 13609
}

func (*UnionBountyGetReportResp) Comment() string {
	return "妖盟悬赏-打开战报"
}
		
func (*UnionBountyReportCheckEscortReq) MsgId() int32 {
	return 213611
}

func (*UnionBountyReportCheckEscortReq) Comment() string {
	return "妖盟悬赏-战报中打开某个玩家的镖车"
}
		
func (*UnionBountyReportCheckEscortResp) MsgId() int32 {
	return 13611
}

func (*UnionBountyReportCheckEscortResp) Comment() string {
	return "妖盟悬赏-战报中打开某个玩家的镖车"
}
		
func (*UnionBountyWorshipReq) MsgId() int32 {
	return 213612
}

func (*UnionBountyWorshipReq) Comment() string {
	return "妖盟悬赏-排行榜膜拜"
}
		
func (*UnionBountyWorshipResp) MsgId() int32 {
	return 13612
}

func (*UnionBountyWorshipResp) Comment() string {
	return "妖盟悬赏-排行榜膜拜"
}
		
func (*UnionBountyRefreshMapReq) MsgId() int32 {
	return 213613
}

func (*UnionBountyRefreshMapReq) Comment() string {
	return "妖盟悬赏-刷新护送地图"
}
		
func (*UnionBountyRefreshMapResp) MsgId() int32 {
	return 13613
}

func (*UnionBountyRefreshMapResp) Comment() string {
	return "妖盟悬赏-刷新护送地图"
}
		
func (*UnionBountyOpenBountyEventReq) MsgId() int32 {
	return 213614
}

func (*UnionBountyOpenBountyEventReq) Comment() string {
	return "妖盟悬赏-打开妖盟战报"
}
		
func (*UnionBountyOpenBountyEventResp) MsgId() int32 {
	return 13614
}

func (*UnionBountyOpenBountyEventResp) Comment() string {
	return "妖盟悬赏-打开妖盟战报"
}
		
func (*UnionBountyRetaliateReq) MsgId() int32 {
	return 213615
}

func (*UnionBountyRetaliateReq) Comment() string {
	return "妖盟悬赏-反击"
}
		
func (*UnionBountyRetaliateResp) MsgId() int32 {
	return 13615
}

func (*UnionBountyRetaliateResp) Comment() string {
	return "妖盟悬赏-反击"
}
		
func (*UnionBountyGetMemberScoreReq) MsgId() int32 {
	return 213616
}

func (*UnionBountyGetMemberScoreReq) Comment() string {
	return "妖盟悬赏-查看联盟积分"
}
		
func (*UnionBountyGetMemberScoreResp) MsgId() int32 {
	return 13616
}

func (*UnionBountyGetMemberScoreResp) Comment() string {
	return "妖盟悬赏-查看联盟积分"
}
		
func (*UnionBountyHaveEscortReq) MsgId() int32 {
	return 213617
}

func (*UnionBountyHaveEscortReq) Comment() string {
	return "妖盟悬赏-是否有正在押镖"
}
		
func (*UnionBountyHaveEscortResp) MsgId() int32 {
	return 13617
}

func (*UnionBountyHaveEscortResp) Comment() string {
	return "妖盟悬赏-是否有正在押镖"
}
		
func (*UnionBountyOpenEscortCartDetailReq) MsgId() int32 {
	return 213618
}

func (*UnionBountyOpenEscortCartDetailReq) Comment() string {
	return "妖盟悬赏-查看镖车详情"
}
		
func (*UnionBountyOpenEscortCartDetailRsp) MsgId() int32 {
	return 13618
}

func (*UnionBountyOpenEscortCartDetailRsp) Comment() string {
	return "妖盟悬赏-查看镖车详情"
}
		
func (*UnionBountyOpenMonsterReq) MsgId() int32 {
	return 213619
}

func (*UnionBountyOpenMonsterReq) Comment() string {
	return "妖盟悬赏-打开妖兽界面"
}
		
func (*UnionBountyOpenMonsterResp) MsgId() int32 {
	return 13619
}

func (*UnionBountyOpenMonsterResp) Comment() string {
	return "妖盟悬赏-打开妖兽界面"
}
		
func (*UnionBountyAttackMonsterReq) MsgId() int32 {
	return 213620
}

func (*UnionBountyAttackMonsterReq) Comment() string {
	return "妖盟悬赏-打妖兽"
}
		
func (*UnionBountyAttackMonsterResp) MsgId() int32 {
	return 13620
}

func (*UnionBountyAttackMonsterResp) Comment() string {
	return "妖盟悬赏-打妖兽"
}
		
func (*UnionBountyAskHelpReq) MsgId() int32 {
	return 213621
}

func (*UnionBountyAskHelpReq) Comment() string {
	return "妖盟悬赏-邀请"
}
		
func (*UnionBountyAskHelpResp) MsgId() int32 {
	return 13621
}

func (*UnionBountyAskHelpResp) Comment() string {
	return "妖盟悬赏-邀请"
}
		
func (*UnionBountyOpenAskHelpReq) MsgId() int32 {
	return 213622
}

func (*UnionBountyOpenAskHelpReq) Comment() string {
	return "妖盟悬赏-打开邀请界面"
}
		
func (*UnionBountyOpenAskHelpResp) MsgId() int32 {
	return 13622
}

func (*UnionBountyOpenAskHelpResp) Comment() string {
	return "妖盟悬赏-打开邀请界面"
}
		
func (*UnionBountyDealAskHelpReq) MsgId() int32 {
	return 213623
}

func (*UnionBountyDealAskHelpReq) Comment() string {
	return "妖盟悬赏-处理邀请（同意加入、拒绝、一键拒绝）"
}
		
func (*UnionBountyDealAskHelpResp) MsgId() int32 {
	return 13623
}

func (*UnionBountyDealAskHelpResp) Comment() string {
	return "妖盟悬赏-处理邀请（同意加入、拒绝、一键拒绝）"
}
		
func (*UnionBountyManageAskHelpReq) MsgId() int32 {
	return 213624
}

func (*UnionBountyManageAskHelpReq) Comment() string {
	return "妖盟悬赏-队伍管理（加入队伍、离开队伍、踢出队伍）"
}
		
func (*UnionBountyManageAskHelpResp) MsgId() int32 {
	return 13624
}

func (*UnionBountyManageAskHelpResp) Comment() string {
	return "妖盟悬赏-队伍管理（加入队伍、离开队伍、踢出队伍）"
}
		
func (*UnionBountyGetMonsterAttributeReq) MsgId() int32 {
	return 213626
}

func (*UnionBountyGetMonsterAttributeReq) Comment() string {
	return "妖盟悬赏-查看当前妖兽的属性"
}
		
func (*UnionBountyGetMonsterAttributeResp) MsgId() int32 {
	return 13626
}

func (*UnionBountyGetMonsterAttributeResp) Comment() string {
	return "妖盟悬赏-查看当前妖兽的属性"
}
		
func (*UnionBountyMonsterChangePosReq) MsgId() int32 {
	return 213627
}

func (*UnionBountyMonsterChangePosReq) Comment() string {
	return "妖盟悬赏-变更参战玩家位置"
}
		
func (*UnionBountyMonsterChangePosRsp) MsgId() int32 {
	return 13627
}

func (*UnionBountyMonsterChangePosRsp) Comment() string {
	return "妖盟悬赏-变更参战玩家位置"
}
		
func (*UnionBountyGetExistMonsterReq) MsgId() int32 {
	return 213628
}

func (*UnionBountyGetExistMonsterReq) Comment() string {
	return "获取当前存在的妖兽"
}
		
func (*UnionBountyGetExistMonsterResp) MsgId() int32 {
	return 13628
}

func (*UnionBountyGetExistMonsterResp) Comment() string {
	return "获取当前存在的妖兽"
}
		
func (*EnterPupilRankActivityReq) MsgId() int32 {
	return 213801
}

func (*EnterPupilRankActivityReq) Comment() string {
	return "进入宗门大比"
}
		
func (*EnterPupilRankActivityResp) MsgId() int32 {
	return 13801
}

func (*EnterPupilRankActivityResp) Comment() string {
	return "进入宗门大比"
}
		
func (*GetIncreaseRecordReq) MsgId() int32 {
	return 213802
}

func (*GetIncreaseRecordReq) Comment() string {
	return "获取涨幅记录"
}
		
func (*GetIncreaseRecordResp) MsgId() int32 {
	return 13802
}

func (*GetIncreaseRecordResp) Comment() string {
	return "获取涨幅记录"
}
		
func (*PupilRankDetailReq) MsgId() int32 {
	return 213803
}

func (*PupilRankDetailReq) Comment() string {
	return "查看 排行榜上玩家的 涨幅记录"
}
		
func (*PupilRankDetailResp) MsgId() int32 {
	return 13803
}

func (*PupilRankDetailResp) Comment() string {
	return "查看 排行榜上玩家的 涨幅记录"
}
		
func (*FriendGetListReq) MsgId() int32 {
	return 214101
}

func (*FriendGetListReq) Comment() string {
	return "获取世交列表"
}
		
func (*FriendGetListResp) MsgId() int32 {
	return 14101
}

func (*FriendGetListResp) Comment() string {
	return "获取世交列表"
}
		
func (*FriendSendReq) MsgId() int32 {
	return 214102
}

func (*FriendSendReq) Comment() string {
	return "赠送"
}
		
func (*FriendSendResp) MsgId() int32 {
	return 14102
}

func (*FriendSendResp) Comment() string {
	return "赠送"
}
		
func (*FriendReceiveReq) MsgId() int32 {
	return 214103
}

func (*FriendReceiveReq) Comment() string {
	return "领取"
}
		
func (*FriendReceiveResp) MsgId() int32 {
	return 14103
}

func (*FriendReceiveResp) Comment() string {
	return "领取"
}
		
func (*FriendReceiveAllReq) MsgId() int32 {
	return 214104
}

func (*FriendReceiveAllReq) Comment() string {
	return "一键赠送、领取"
}
		
func (*FriendReceiveAllResp) MsgId() int32 {
	return 14104
}

func (*FriendReceiveAllResp) Comment() string {
	return "一键赠送、领取"
}
		
func (*FriendDeleteReq) MsgId() int32 {
	return 214105
}

func (*FriendDeleteReq) Comment() string {
	return "删除世交"
}
		
func (*FriendDeleteResp) MsgId() int32 {
	return 14105
}

func (*FriendDeleteResp) Comment() string {
	return "删除世交"
}
		
func (*FriendChatReqMsg) MsgId() int32 {
	return 214120
}

func (*FriendChatReqMsg) Comment() string {
	return "好友消息列表"
}
		
func (*FairyLandSouthDoorHelpReq) MsgId() int32 {
	return 214208
}

func (*FairyLandSouthDoorHelpReq) Comment() string {
	return "飞升仙界-南天门协助"
}
		
func (*FairyLandSouthDoorHelpResp) MsgId() int32 {
	return 14208
}

func (*FairyLandSouthDoorHelpResp) Comment() string {
	return "飞升仙界-南天门协助"
}
		
func (*UniverseLevelUpReq) MsgId() int32 {
	return 214303
}

func (*UniverseLevelUpReq) Comment() string {
	return "小世界升级"
}
		
func (*UniverseLevelUpResp) MsgId() int32 {
	return 14303
}

func (*UniverseLevelUpResp) Comment() string {
	return "小世界升级"
}
		
func (*UniverseDrawReq) MsgId() int32 {
	return 214304
}

func (*UniverseDrawReq) Comment() string {
	return "小世界轮盘抽奖"
}
		
func (*UniverseDrawResp) MsgId() int32 {
	return 14304
}

func (*UniverseDrawResp) Comment() string {
	return "小世界轮盘抽奖"
}
		
func (*UniverseSkillUnlockReq) MsgId() int32 {
	return 214305
}

func (*UniverseSkillUnlockReq) Comment() string {
	return "小世界技能激活"
}
		
func (*UniverseSkillUnlockResp) MsgId() int32 {
	return 14305
}

func (*UniverseSkillUnlockResp) Comment() string {
	return "小世界技能激活"
}
		
func (*EquipUniverseSkillReq) MsgId() int32 {
	return 214306
}

func (*EquipUniverseSkillReq) Comment() string {
	return "小世界技能上阵"
}
		
func (*EquipUniverseSkillResp) MsgId() int32 {
	return 14306
}

func (*EquipUniverseSkillResp) Comment() string {
	return "小世界技能上阵"
}
		
func (*UniverseSkillCombineLvUpReq) MsgId() int32 {
	return 214307
}

func (*UniverseSkillCombineLvUpReq) Comment() string {
	return "小世界技能共鸣激活、升级"
}
		
func (*UniverseSkillCombineLvUpResp) MsgId() int32 {
	return 14307
}

func (*UniverseSkillCombineLvUpResp) Comment() string {
	return "小世界技能共鸣激活、升级"
}
		
func (*UniverseSkillDrawReq) MsgId() int32 {
	return 214308
}

func (*UniverseSkillDrawReq) Comment() string {
	return "洞察天机抽取"
}
		
func (*UniverseSkillDrawResp) MsgId() int32 {
	return 14308
}

func (*UniverseSkillDrawResp) Comment() string {
	return "洞察天机抽取"
}
		
func (*UniverseSkillLvUpReq) MsgId() int32 {
	return 214309
}

func (*UniverseSkillLvUpReq) Comment() string {
	return "小世界技能升级"
}
		
func (*UniverseSkillLvUpResp) MsgId() int32 {
	return 14309
}

func (*UniverseSkillLvUpResp) Comment() string {
	return "小世界技能升级"
}
		
func (*UniverseDrawTwiceReq) MsgId() int32 {
	return 214310
}

func (*UniverseDrawTwiceReq) Comment() string {
	return "天地轮盘二次抽奖"
}
		
func (*UniverseDrawTwiceResp) MsgId() int32 {
	return 14310
}

func (*UniverseDrawTwiceResp) Comment() string {
	return "天地轮盘二次抽奖"
}
		
func (*GodDemonBattleGetPlayerSettlementReq) MsgId() int32 {
	return 214527
}

func (*GodDemonBattleGetPlayerSettlementReq) Comment() string {
	return "获取玩家结算信息"
}
		
func (*GodDemonBattleGetPlayerSettlementResp) MsgId() int32 {
	return 14527
}

func (*GodDemonBattleGetPlayerSettlementResp) Comment() string {
	return "获取玩家结算信息"
}
		
func (*GodDemonBattleGetCampSettlementReq) MsgId() int32 {
	return 214528
}

func (*GodDemonBattleGetCampSettlementReq) Comment() string {
	return "获取阵营结算信息"
}
		
func (*GodDemonBattleGetCampSettlementResp) MsgId() int32 {
	return 14528
}

func (*GodDemonBattleGetCampSettlementResp) Comment() string {
	return "获取阵营结算信息"
}
		
func (*CheckFairyPoolReq) MsgId() int32 {
	return 214698
}

func (*CheckFairyPoolReq) Comment() string {
	return "画中仙境-检测仙泽"
}
		
func (*CheckFairyPoolResp) MsgId() int32 {
	return 14698
}

func (*CheckFairyPoolResp) Comment() string {
	return "画中仙境-检测仙泽"
}
		
func (*SkyTradeEnterActivityReq) MsgId() int32 {
	return 215000
}

func (*SkyTradeEnterActivityReq) Comment() string {
	return "进入活动"
}
		
func (*SkyTradeEnterActivityResp) MsgId() int32 {
	return 15000
}

func (*SkyTradeEnterActivityResp) Comment() string {
	return "进入活动"
}
		
func (*SkyTradeEnterReq) MsgId() int32 {
	return 215001
}

func (*SkyTradeEnterReq) Comment() string {
	return "进入地图"
}
		
func (*SkyTradeEnterResp) MsgId() int32 {
	return 15001
}

func (*SkyTradeEnterResp) Comment() string {
	return "进入地图"
}
		
func (*SkyTradeGroupInfoReq) MsgId() int32 {
	return 215002
}

func (*SkyTradeGroupInfoReq) Comment() string {
	return "查看分组信息"
}
		
func (*SkyTradeGroupInfoResp) MsgId() int32 {
	return 15002
}

func (*SkyTradeGroupInfoResp) Comment() string {
	return "查看分组信息"
}
		
func (*SkyTradeAddSpeedReq) MsgId() int32 {
	return 215003
}

func (*SkyTradeAddSpeedReq) Comment() string {
	return "加速移动"
}
		
func (*SkyTradeAddSpeedResp) MsgId() int32 {
	return 15003
}

func (*SkyTradeAddSpeedResp) Comment() string {
	return "加速移动"
}
		
func (*SkyTradeGotoPortReq) MsgId() int32 {
	return 215004
}

func (*SkyTradeGotoPortReq) Comment() string {
	return "请求移动"
}
		
func (*SkyTradeGotoPortResp) MsgId() int32 {
	return 15004
}

func (*SkyTradeGotoPortResp) Comment() string {
	return "请求移动"
}
		
func (*SkyTradeDealReq) MsgId() int32 {
	return 215005
}

func (*SkyTradeDealReq) Comment() string {
	return "购买货物"
}
		
func (*SkyTradeDealResp) MsgId() int32 {
	return 15005
}

func (*SkyTradeDealResp) Comment() string {
	return "购买货物"
}
		
func (*SkyTradeChallengeListReq) MsgId() int32 {
	return 215006
}

func (*SkyTradeChallengeListReq) Comment() string {
	return "挑战列表"
}
		
func (*SkyTradeChallengeListResp) MsgId() int32 {
	return 15006
}

func (*SkyTradeChallengeListResp) Comment() string {
	return "挑战列表"
}
		
func (*SkyTradeChallengeReq) MsgId() int32 {
	return 215007
}

func (*SkyTradeChallengeReq) Comment() string {
	return "请求挑战"
}
		
func (*SkyTradeChallengeResp) MsgId() int32 {
	return 15007
}

func (*SkyTradeChallengeResp) Comment() string {
	return "请求挑战"
}
		
func (*SkyTradeUnionFameReq) MsgId() int32 {
	return 215008
}

func (*SkyTradeUnionFameReq) Comment() string {
	return "妖盟成员跑商威望信息"
}
		
func (*SkyTradeUnionFameResp) MsgId() int32 {
	return 15008
}

func (*SkyTradeUnionFameResp) Comment() string {
	return "妖盟成员跑商威望信息"
}
		
func (*SkyTradeGetRewardReq) MsgId() int32 {
	return 215009
}

func (*SkyTradeGetRewardReq) Comment() string {
	return "领取到港事件奖励"
}
		
func (*SkyTradeGetRewardResp) MsgId() int32 {
	return 15009
}

func (*SkyTradeGetRewardResp) Comment() string {
	return "领取到港事件奖励"
}
		
func (*SkyTradeReportReq) MsgId() int32 {
	return 215010
}

func (*SkyTradeReportReq) Comment() string {
	return "请求战报"
}
		
func (*SkyTradeReportResp) MsgId() int32 {
	return 15010
}

func (*SkyTradeReportResp) Comment() string {
	return "请求战报"
}
		
func (*SkyTradePortLogReq) MsgId() int32 {
	return 215011
}

func (*SkyTradePortLogReq) Comment() string {
	return "获取日志"
}
		
func (*SkyTradeLogResp) MsgId() int32 {
	return 15011
}

func (*SkyTradeLogResp) Comment() string {
	return "获取日志"
}
		
func (*SkyTradeSparInfoReq) MsgId() int32 {
	return 215012
}

func (*SkyTradeSparInfoReq) Comment() string {
	return "天地灵石信息"
}
		
func (*SkyTradeSparInfoResp) MsgId() int32 {
	return 15012
}

func (*SkyTradeSparInfoResp) Comment() string {
	return "天地灵石信息"
}
		
func (*SkyTradeGetSparPowerReq) MsgId() int32 {
	return 215013
}

func (*SkyTradeGetSparPowerReq) Comment() string {
	return "领取天地灵石等级奖励"
}
		
func (*SkyTradeGetSparPowerResp) MsgId() int32 {
	return 15013
}

func (*SkyTradeGetSparPowerResp) Comment() string {
	return "领取天地灵石等级奖励"
}
		
func (*SkyTradeGetWelfareReq) MsgId() int32 {
	return 215014
}

func (*SkyTradeGetWelfareReq) Comment() string {
	return "领取天地灵石盟友奖励"
}
		
func (*SkyTradeGetWelfareResp) MsgId() int32 {
	return 15014
}

func (*SkyTradeGetWelfareResp) Comment() string {
	return "领取天地灵石盟友奖励"
}
		
func (*SkyTradeSendWelfareReq) MsgId() int32 {
	return 215015
}

func (*SkyTradeSendWelfareReq) Comment() string {
	return "发送灵液"
}
		
func (*SkyTradeSendWelfareResp) MsgId() int32 {
	return 15015
}

func (*SkyTradeSendWelfareResp) Comment() string {
	return "发送灵液"
}
		
func (*SkyTradeWelfareRecordReq) MsgId() int32 {
	return 215016
}

func (*SkyTradeWelfareRecordReq) Comment() string {
	return "灵液记录"
}
		
func (*SkyTradeWelfareRecordResp) MsgId() int32 {
	return 15016
}

func (*SkyTradeWelfareRecordResp) Comment() string {
	return "灵液记录"
}
		
func (*SkyTradeChallengeHeartBeatReq) MsgId() int32 {
	return 215017
}

func (*SkyTradeChallengeHeartBeatReq) Comment() string {
	return "挑战界面心跳"
}
		
func (*SkyTradeChallengeHeartBeatResp) MsgId() int32 {
	return 15017
}

func (*SkyTradeChallengeHeartBeatResp) Comment() string {
	return "挑战界面心跳"
}
		
func (*SkyTradeUseRobItemReq) MsgId() int32 {
	return 215018
}

func (*SkyTradeUseRobItemReq) Comment() string {
	return "使用道具补充挑战次数"
}
		
func (*SkyTradeUseRobItemResp) MsgId() int32 {
	return 15018
}

func (*SkyTradeUseRobItemResp) Comment() string {
	return "使用道具补充挑战次数"
}
		
func (*SkyTradeReportDetailReq) MsgId() int32 {
	return 215019
}

func (*SkyTradeReportDetailReq) Comment() string {
	return "请求详细战报信息"
}
		
func (*SkyTradeReportDetailResp) MsgId() int32 {
	return 15019
}

func (*SkyTradeReportDetailResp) Comment() string {
	return "请求详细战报信息"
}
		
func (*SkyTradeGuessPlayersReq) MsgId() int32 {
	return 215020
}

func (*SkyTradeGuessPlayersReq) Comment() string {
	return "竞猜结果"
}
		
func (*SkyTradeGuessPlayersResp) MsgId() int32 {
	return 15020
}

func (*SkyTradeGuessPlayersResp) Comment() string {
	return "竞猜结果"
}
		
func (*SkyTradeGuessDataReq) MsgId() int32 {
	return 215021
}

func (*SkyTradeGuessDataReq) Comment() string {
	return "请求竞猜界面数据"
}
		
func (*SkyTradeGuessDataResp) MsgId() int32 {
	return 15021
}

func (*SkyTradeGuessDataResp) Comment() string {
	return "请求竞猜界面数据"
}
		
func (*SkyTradeGuessSelectReq) MsgId() int32 {
	return 215022
}

func (*SkyTradeGuessSelectReq) Comment() string {
	return "设置竞猜选择"
}
		
func (*SkyTradeGuessSelectResp) MsgId() int32 {
	return 15022
}

func (*SkyTradeGuessSelectResp) Comment() string {
	return "设置竞猜选择"
}
		
func (*SkyTradeGuessRewardReq) MsgId() int32 {
	return 215023
}

func (*SkyTradeGuessRewardReq) Comment() string {
	return "领取竞猜奖励"
}
		
func (*SkyTradeGuessRewardResp) MsgId() int32 {
	return 15023
}

func (*SkyTradeGuessRewardResp) Comment() string {
	return "领取竞猜奖励"
}
		
func (*SkyTradeReportGoodsReq) MsgId() int32 {
	return 215024
}

func (*SkyTradeReportGoodsReq) Comment() string {
	return "请求上报所在地点稀缺货物"
}
		
func (*SkyTradeReportGoodsResp) MsgId() int32 {
	return 15024
}

func (*SkyTradeReportGoodsResp) Comment() string {
	return "请求上报所在地点稀缺货物"
}
		
func (*SkyTradeResetStockReq) MsgId() int32 {
	return 215025
}

func (*SkyTradeResetStockReq) Comment() string {
	return "请求刷新货物库存"
}
		
func (*SkyTradeResetStockResp) MsgId() int32 {
	return 15025
}

func (*SkyTradeResetStockResp) Comment() string {
	return "请求刷新货物库存"
}
		
func (*SkyTradeRedPointReq) MsgId() int32 {
	return 215026
}

func (*SkyTradeRedPointReq) Comment() string {
	return "红点详情"
}
		
func (*SkyTradeRedPointResp) MsgId() int32 {
	return 15026
}

func (*SkyTradeRedPointResp) Comment() string {
	return "红点详情"
}
		
func (*SkyTradeUnionGroupDamageReq) MsgId() int32 {
	return 215027
}

func (*SkyTradeUnionGroupDamageReq) Comment() string {
	return "妖盟分段为伤害信息"
}
		
func (*SkyTradeUnionGroupDamageResp) MsgId() int32 {
	return 15027
}

func (*SkyTradeUnionGroupDamageResp) Comment() string {
	return "妖盟分段为伤害信息"
}
		
func (*ADGiftTriggerReq) MsgId() int32 {
	return 215201
}

func (*ADGiftTriggerReq) Comment() string {
	return "支付宝广告礼包-触发"
}
		
func (*ADGiftTriggerResp) MsgId() int32 {
	return 15201
}

func (*ADGiftTriggerResp) Comment() string {
	return "支付宝广告礼包-触发"
}
		
func (*ADGiftGetRewardReq) MsgId() int32 {
	return 215202
}

func (*ADGiftGetRewardReq) Comment() string {
	return "支付宝广告礼包-领取奖励"
}
		
func (*ADGiftGetRewardResp) MsgId() int32 {
	return 15202
}

func (*ADGiftGetRewardResp) Comment() string {
	return "支付宝广告礼包-领取奖励"
}
		
func (*YardEnterReq) MsgId() int32 {
	return 215801
}

func (*YardEnterReq) Comment() string {
	return "家园-进入"
}
		
func (*YardEnterResp) MsgId() int32 {
	return 15801
}

func (*YardEnterResp) Comment() string {
	return "家园-进入"
}
		
func (*YardBattleReq) MsgId() int32 {
	return 215802
}

func (*YardBattleReq) Comment() string {
	return "家园-战斗解锁区域"
}
		
func (*YardBattleResp) MsgId() int32 {
	return 15802
}

func (*YardBattleResp) Comment() string {
	return "家园-战斗解锁区域"
}
		
func (*YardPosActionReq) MsgId() int32 {
	return 215803
}

func (*YardPosActionReq) Comment() string {
	return "家园-建筑摆放"
}
		
func (*YardPosActionResp) MsgId() int32 {
	return 15803
}

func (*YardPosActionResp) Comment() string {
	return "家园-建筑摆放"
}
		
func (*YardLevelUpReq) MsgId() int32 {
	return 215804
}

func (*YardLevelUpReq) Comment() string {
	return "家园升级"
}
		
func (*YardLevelUpResp) MsgId() int32 {
	return 15804
}

func (*YardLevelUpResp) Comment() string {
	return "家园升级"
}
		
func (*YardClearReq) MsgId() int32 {
	return 215805
}

func (*YardClearReq) Comment() string {
	return "家园-清理杂物"
}
		
func (*YardClearResp) MsgId() int32 {
	return 15805
}

func (*YardClearResp) Comment() string {
	return "家园-清理杂物"
}
		
func (*YardInviteReq) MsgId() int32 {
	return 215806
}

func (*YardInviteReq) Comment() string {
	return "家园请求"
}
		
func (*YardInviteResp) MsgId() int32 {
	return 15806
}

func (*YardInviteResp) Comment() string {
	return "家园请求"
}
		
func (*YardEnterShopReq) MsgId() int32 {
	return 215807
}

func (*YardEnterShopReq) Comment() string {
	return "进入兑换商店"
}
		
func (*YardEnterShopResp) MsgId() int32 {
	return 15807
}

func (*YardEnterShopResp) Comment() string {
	return "进入兑换商店"
}
		
func (*YardRefreshShopReq) MsgId() int32 {
	return 215808
}

func (*YardRefreshShopReq) Comment() string {
	return "刷新兑换商店"
}
		
func (*YardRefreshShopResp) MsgId() int32 {
	return 15808
}

func (*YardRefreshShopResp) Comment() string {
	return "刷新兑换商店"
}
		
func (*YardSkillCombineLvUpReq) MsgId() int32 {
	return 215821
}

func (*YardSkillCombineLvUpReq) Comment() string {
	return "羁绊共鸣升级"
}
		
func (*YardSkillCombineLvUpResp) MsgId() int32 {
	return 15821
}

func (*YardSkillCombineLvUpResp) Comment() string {
	return "羁绊共鸣升级"
}
		
func (*YardDrawReq) MsgId() int32 {
	return 215822
}

func (*YardDrawReq) Comment() string {
	return "抽取建筑"
}
		
func (*YardDrawResp) MsgId() int32 {
	return 15822
}

func (*YardDrawResp) Comment() string {
	return "抽取建筑"
}
		
func (*BuildUpLevelReq) MsgId() int32 {
	return 215823
}

func (*BuildUpLevelReq) Comment() string {
	return "建筑升级"
}
		
func (*BuildUpLevelResp) MsgId() int32 {
	return 15823
}

func (*BuildUpLevelResp) Comment() string {
	return "建筑升级"
}
		
func (*BuildUnlockReq) MsgId() int32 {
	return 215824
}

func (*BuildUnlockReq) Comment() string {
	return "装饰建筑解锁"
}
		
func (*BuildUnlockResp) MsgId() int32 {
	return 15824
}

func (*BuildUnlockResp) Comment() string {
	return "装饰建筑解锁"
}
		
func (*YardBuildMakeReq) MsgId() int32 {
	return 215825
}

func (*YardBuildMakeReq) Comment() string {
	return "建筑生产"
}
		
func (*YardBuildMakeResp) MsgId() int32 {
	return 15825
}

func (*YardBuildMakeResp) Comment() string {
	return "建筑生产"
}
		
func (*YardBuildHelpReq) MsgId() int32 {
	return 215826
}

func (*YardBuildHelpReq) Comment() string {
	return "桃子协助"
}
		
func (*YardBuildhHelpResp) MsgId() int32 {
	return 15826
}

func (*YardBuildhHelpResp) Comment() string {
	return "桃子协助"
}
		
func (*YardBuildGainRewardReq) MsgId() int32 {
	return 215827
}

func (*YardBuildGainRewardReq) Comment() string {
	return "建筑领奖"
}
		
func (*YardBuildGainRewardResp) MsgId() int32 {
	return 15827
}

func (*YardBuildGainRewardResp) Comment() string {
	return "建筑领奖"
}
		
func (*YardBuildSpeedUpReq) MsgId() int32 {
	return 215828
}

func (*YardBuildSpeedUpReq) Comment() string {
	return "建筑升级加速"
}
		
func (*YardBuildhSpeedUpResp) MsgId() int32 {
	return 15828
}

func (*YardBuildhSpeedUpResp) Comment() string {
	return "建筑升级加速"
}
		
func (*GetYardUnionHelpDataListReq) MsgId() int32 {
	return 215829
}

func (*GetYardUnionHelpDataListReq) Comment() string {
	return "妖盟协助列表请求"
}
		
func (*GetYardUnionHelpDataListResp) MsgId() int32 {
	return 15829
}

func (*GetYardUnionHelpDataListResp) Comment() string {
	return "妖盟协助列表请求"
}
		
func (*RequestYardUnionHelpReq) MsgId() int32 {
	return 215830
}

func (*RequestYardUnionHelpReq) Comment() string {
	return "请求妖盟协助"
}
		
func (*RequestYardUnionHelpResp) MsgId() int32 {
	return 15830
}

func (*RequestYardUnionHelpResp) Comment() string {
	return "请求妖盟协助"
}
		
func (*YardUnionHelpReq) MsgId() int32 {
	return 215831
}

func (*YardUnionHelpReq) Comment() string {
	return "协助他人"
}
		
func (*YardUnionHelpResp) MsgId() int32 {
	return 15831
}

func (*YardUnionHelpResp) Comment() string {
	return "协助他人"
}
		
func (*YardOneKeyLevelUpReq) MsgId() int32 {
	return 215833
}

func (*YardOneKeyLevelUpReq) Comment() string {
	return "家园装饰建筑一键升级"
}
		
func (*YardOneKeyLevelUpResp) MsgId() int32 {
	return 15833
}

func (*YardOneKeyLevelUpResp) Comment() string {
	return "家园装饰建筑一键升级"
}
		
func (*YardPeachRecordListReq) MsgId() int32 {
	return 215834
}

func (*YardPeachRecordListReq) Comment() string {
	return "桃树协助列表"
}
		
func (*YardPeachRecordListResp) MsgId() int32 {
	return 15834
}

func (*YardPeachRecordListResp) Comment() string {
	return "桃树协助列表"
}
		
func (*YardMerchantInfoReq) MsgId() int32 {
	return 215835
}

func (*YardMerchantInfoReq) Comment() string {
	return "产物商人信息"
}
		
func (*YardMerchantInfoResp) MsgId() int32 {
	return 15835
}

func (*YardMerchantInfoResp) Comment() string {
	return "产物商人信息"
}
		
func (*YardMerchantExchangeReq) MsgId() int32 {
	return 215836
}

func (*YardMerchantExchangeReq) Comment() string {
	return "产物出售"
}
		
func (*YardMerchantExchangeResp) MsgId() int32 {
	return 15836
}

func (*YardMerchantExchangeResp) Comment() string {
	return "产物出售"
}
		
func (*YardChatSwitchReq) MsgId() int32 {
	return 215837
}

func (*YardChatSwitchReq) Comment() string {
	return "留言板开关"
}
		
func (*YardChatSwitchResp) MsgId() int32 {
	return 15837
}

func (*YardChatSwitchResp) Comment() string {
	return "留言板开关"
}
		
func (*YardChatActionReq) MsgId() int32 {
	return 215839
}

func (*YardChatActionReq) Comment() string {
	return "留言操作"
}
		
func (*YardChatActionResp) MsgId() int32 {
	return 15839
}

func (*YardChatActionResp) Comment() string {
	return "留言操作"
}
		
func (*YardChatRecordReq) MsgId() int32 {
	return 215840
}

func (*YardChatRecordReq) Comment() string {
	return "留言"
}
		
func (*YardChatRecordResp) MsgId() int32 {
	return 15840
}

func (*YardChatRecordResp) Comment() string {
	return "留言"
}
		
func (*YardChatRecordListReq) MsgId() int32 {
	return 215841
}

func (*YardChatRecordListReq) Comment() string {
	return "留言板信息"
}
		
func (*YardChatRecordListResp) MsgId() int32 {
	return 15841
}

func (*YardChatRecordListResp) Comment() string {
	return "留言板信息"
}
		
func (*YardGiveLikeReq) MsgId() int32 {
	return 215842
}

func (*YardGiveLikeReq) Comment() string {
	return "点赞"
}
		
func (*YardGiveLikeResp) MsgId() int32 {
	return 15842
}

func (*YardGiveLikeResp) Comment() string {
	return "点赞"
}
		
func (*YardReportReq) MsgId() int32 {
	return 215852
}

func (*YardReportReq) Comment() string {
	return "家园举报"
}
		
func (*YardReportResp) MsgId() int32 {
	return 15852
}

func (*YardReportResp) Comment() string {
	return "家园举报"
}
		
func (*ShuraBattleApplyReq) MsgId() int32 {
	return 215910
}

func (*ShuraBattleApplyReq) Comment() string {
	return "申请、撤销"
}
		
func (*ShuraBattleApplyResp) MsgId() int32 {
	return 15910
}

func (*ShuraBattleApplyResp) Comment() string {
	return "申请、撤销"
}
		
func (*EnterPlanesTrialReq) MsgId() int32 {
	return 216001
}

func (*EnterPlanesTrialReq) Comment() string {
	return "进入三界征途界面"
}
		
func (*EnterPlanesTrialRsp) MsgId() int32 {
	return 16001
}

func (*EnterPlanesTrialRsp) Comment() string {
	return "进入三界征途界面"
}
		
func (*EnterPlanesTrialTeamReq) MsgId() int32 {
	return 216002
}

func (*EnterPlanesTrialTeamReq) Comment() string {
	return "组队内部请求"
}
		
func (*EnterPlanesTrialTeamRsp) MsgId() int32 {
	return 16002
}

func (*EnterPlanesTrialTeamRsp) Comment() string {
	return "组队内部请求"
}
		
func (*PlanesTrialTeamStartReq) MsgId() int32 {
	return 216003
}

func (*PlanesTrialTeamStartReq) Comment() string {
	return "组队页面请求"
}
		
func (*PlanesTrialTeamStartRsp) MsgId() int32 {
	return 16003
}

func (*PlanesTrialTeamStartRsp) Comment() string {
	return "组队页面请求"
}
		
func (*PlanesTrialCreateTeamReq) MsgId() int32 {
	return 216004
}

func (*PlanesTrialCreateTeamReq) Comment() string {
	return "创建队伍"
}
		
func (*PlanesTrialCreateTeamRsp) MsgId() int32 {
	return 16004
}

func (*PlanesTrialCreateTeamRsp) Comment() string {
	return "创建队伍"
}
		
func (*PlanesTrialGetTeamListReq) MsgId() int32 {
	return 216005
}

func (*PlanesTrialGetTeamListReq) Comment() string {
	return "获取队伍列表"
}
		
func (*PlanesTrialGetTeamListRsp) MsgId() int32 {
	return 16005
}

func (*PlanesTrialGetTeamListRsp) Comment() string {
	return "获取队伍列表"
}
		
func (*PlanesTrialGetTeamInfoReq) MsgId() int32 {
	return 216006
}

func (*PlanesTrialGetTeamInfoReq) Comment() string {
	return "查找队伍"
}
		
func (*PlanesTrialGetTeamInfoRsp) MsgId() int32 {
	return 16006
}

func (*PlanesTrialGetTeamInfoRsp) Comment() string {
	return "查找队伍"
}
		
func (*PlanesTrialJoinTeamReq) MsgId() int32 {
	return 216007
}

func (*PlanesTrialJoinTeamReq) Comment() string {
	return "三界征途-申请加入队伍"
}
		
func (*PlanesTrialJoinTeamRsp) MsgId() int32 {
	return 16007
}

func (*PlanesTrialJoinTeamRsp) Comment() string {
	return "三界征途-申请加入队伍"
}
		
func (*PlanesTrialCancelTeamApplyReq) MsgId() int32 {
	return 216008
}

func (*PlanesTrialCancelTeamApplyReq) Comment() string {
	return "取消加入队伍"
}
		
func (*PlanesTrialCancelTeamApplyRsp) MsgId() int32 {
	return 16008
}

func (*PlanesTrialCancelTeamApplyRsp) Comment() string {
	return "取消加入队伍"
}
		
func (*PlanesTrialApplyJoinTeamAgreeReq) MsgId() int32 {
	return 216009
}

func (*PlanesTrialApplyJoinTeamAgreeReq) Comment() string {
	return "同意申请"
}
		
func (*PlanesTrialApplyJoinTeamAgreeRsp) MsgId() int32 {
	return 16009
}

func (*PlanesTrialApplyJoinTeamAgreeRsp) Comment() string {
	return "同意申请"
}
		
func (*PlanesTrialApplyJoinTeamRefusedReq) MsgId() int32 {
	return 216010
}

func (*PlanesTrialApplyJoinTeamRefusedReq) Comment() string {
	return "一键拒绝"
}
		
func (*PlanesTrialApplyJoinTeamRefusedRsp) MsgId() int32 {
	return 16010
}

func (*PlanesTrialApplyJoinTeamRefusedRsp) Comment() string {
	return "一键拒绝"
}
		
func (*PlanesTrialQuitTeamReq) MsgId() int32 {
	return 216011
}

func (*PlanesTrialQuitTeamReq) Comment() string {
	return "退出队伍"
}
		
func (*PlanesTrialQuitTeamRsp) MsgId() int32 {
	return 16011
}

func (*PlanesTrialQuitTeamRsp) Comment() string {
	return "退出队伍"
}
		
func (*PlanesTrialKickOutTeamReq) MsgId() int32 {
	return 216012
}

func (*PlanesTrialKickOutTeamReq) Comment() string {
	return "踢出队伍"
}
		
func (*PlanesTrialKickOutTeamRsp) MsgId() int32 {
	return 16012
}

func (*PlanesTrialKickOutTeamRsp) Comment() string {
	return "踢出队伍"
}
		
func (*PlanesTrialChangeLeaderReq) MsgId() int32 {
	return 216013
}

func (*PlanesTrialChangeLeaderReq) Comment() string {
	return "转让队长"
}
		
func (*PlanesTrialChangeLeaderRsp) MsgId() int32 {
	return 16013
}

func (*PlanesTrialChangeLeaderRsp) Comment() string {
	return "转让队长"
}
		
func (*PlanesTrialTeamPrepareReq) MsgId() int32 {
	return 216014
}

func (*PlanesTrialTeamPrepareReq) Comment() string {
	return "开始准备"
}
		
func (*PlanesTrialTeamPrepareRsp) MsgId() int32 {
	return 16014
}

func (*PlanesTrialTeamPrepareRsp) Comment() string {
	return "开始准备"
}
		
func (*PlanesTrialMatchMemberReq) MsgId() int32 {
	return 216015
}

func (*PlanesTrialMatchMemberReq) Comment() string {
	return "匹配"
}
		
func (*PlanesTrialMatchMemberRsp) MsgId() int32 {
	return 16015
}

func (*PlanesTrialMatchMemberRsp) Comment() string {
	return "匹配"
}
		
func (*PlanesTrialStartBattleReq) MsgId() int32 {
	return 216016
}

func (*PlanesTrialStartBattleReq) Comment() string {
	return "开始战斗"
}
		
func (*PlanesTrialStartBattleRsp) MsgId() int32 {
	return 16016
}

func (*PlanesTrialStartBattleRsp) Comment() string {
	return "开始战斗"
}
		
func (*PlanesTrialGetPlayerInfoReq) MsgId() int32 {
	return 216017
}

func (*PlanesTrialGetPlayerInfoReq) Comment() string {
	return "请求玩家数据"
}
		
func (*PlanesTrialGetPlayerInfoRsp) MsgId() int32 {
	return 16017
}

func (*PlanesTrialGetPlayerInfoRsp) Comment() string {
	return "请求玩家数据"
}
		
func (*PlanesTrialGetBattleVideoReq) MsgId() int32 {
	return 216019
}

func (*PlanesTrialGetBattleVideoReq) Comment() string {
	return "战斗回放"
}
		
func (*PlanesTrialGetBattleVideoRsp) MsgId() int32 {
	return 16019
}

func (*PlanesTrialGetBattleVideoRsp) Comment() string {
	return "战斗回放"
}
		
func (*PlanesTrialGetBossInfoReq) MsgId() int32 {
	return 216020
}

func (*PlanesTrialGetBossInfoReq) Comment() string {
	return "获取boss信息"
}
		
func (*PlanesTrialGetBossInfoRsp) MsgId() int32 {
	return 16020
}

func (*PlanesTrialGetBossInfoRsp) Comment() string {
	return "获取boss信息"
}
		
func (*PlanesTrialChallengeTimeReq) MsgId() int32 {
	return 216021
}

func (*PlanesTrialChallengeTimeReq) Comment() string {
	return "三界征途-获取挑战次数"
}
		
func (*PlanesTrialChallengeTimeRsp) MsgId() int32 {
	return 16021
}

func (*PlanesTrialChallengeTimeRsp) Comment() string {
	return "三界征途-获取挑战次数"
}
		
func (*PlanesTrialStartMatchReq) MsgId() int32 {
	return 216022
}

func (*PlanesTrialStartMatchReq) Comment() string {
	return "开始匹配"
}
		
func (*PlanesTrialStartMatchRsp) MsgId() int32 {
	return 16022
}

func (*PlanesTrialStartMatchRsp) Comment() string {
	return "开始匹配"
}
		
func (*PlanesTrialInviteReq) MsgId() int32 {
	return 216023
}

func (*PlanesTrialInviteReq) Comment() string {
	return "邀请"
}
		
func (*PlanesTrialInviteRsp) MsgId() int32 {
	return 16023
}

func (*PlanesTrialInviteRsp) Comment() string {
	return "邀请"
}
		
func (*LeavePlanesTrialReq) MsgId() int32 {
	return 216024
}

func (*LeavePlanesTrialReq) Comment() string {
	return "离开系统"
}
		
func (*LeavePlanesTrialRsp) MsgId() int32 {
	return 16024
}

func (*LeavePlanesTrialRsp) Comment() string {
	return "离开系统"
}
		
func (*PlanesTrialGetBossPowerReq) MsgId() int32 {
	return 216025
}

func (*PlanesTrialGetBossPowerReq) Comment() string {
	return "获取boss妖力"
}
		
func (*PlanesTrialGetBossPowerRsp) MsgId() int32 {
	return 16025
}

func (*PlanesTrialGetBossPowerRsp) Comment() string {
	return "获取boss妖力"
}
		
func (*PlanesTrialSkipBattleReq) MsgId() int32 {
	return 216026
}

func (*PlanesTrialSkipBattleReq) Comment() string {
	return "队长跳过战斗"
}
		
func (*PlanesTrialSkipBattleRsp) MsgId() int32 {
	return 16026
}

func (*PlanesTrialSkipBattleRsp) Comment() string {
	return "队长跳过战斗"
}
		
func (*PlanesTrialStartSelectBuffReq) MsgId() int32 {
	return 216027
}

func (*PlanesTrialStartSelectBuffReq) Comment() string {
	return "队长开始选buff"
}
		
func (*PlanesTrialStartSelectBuffRsp) MsgId() int32 {
	return 16027
}

func (*PlanesTrialStartSelectBuffRsp) Comment() string {
	return "队长开始选buff"
}
		
func (*PlanesTrialSelectBuffReq) MsgId() int32 {
	return 216028
}

func (*PlanesTrialSelectBuffReq) Comment() string {
	return "选buff"
}
		
func (*PlanesTrialSelectBuffRsp) MsgId() int32 {
	return 16028
}

func (*PlanesTrialSelectBuffRsp) Comment() string {
	return "选buff"
}
		
func (*PlanesTrialSetBuffPreferenceReq) MsgId() int32 {
	return 216029
}

func (*PlanesTrialSetBuffPreferenceReq) Comment() string {
	return "保存选择选项"
}
		
func (*PlanesTrialSetBuffPreferenceRsp) MsgId() int32 {
	return 16029
}

func (*PlanesTrialSetBuffPreferenceRsp) Comment() string {
	return "保存选择选项"
}
		
func (*PlanesTrialGetSelectedBuffReq) MsgId() int32 {
	return 216030
}

func (*PlanesTrialGetSelectedBuffReq) Comment() string {
	return "查看自己选的buff"
}
		
func (*PlanesTrialGetSelectedBuffRsp) MsgId() int32 {
	return 16030
}

func (*PlanesTrialGetSelectedBuffRsp) Comment() string {
	return "查看自己选的buff"
}
		
func (*PlanesTrialGetAchievementRewardReq) MsgId() int32 {
	return 216031
}

func (*PlanesTrialGetAchievementRewardReq) Comment() string {
	return "领取成就奖励"
}
		
func (*PlanesTrialGetAchievementRewardRsp) MsgId() int32 {
	return 16031
}

func (*PlanesTrialGetAchievementRewardRsp) Comment() string {
	return "领取成就奖励"
}
		
func (*PlanesTrialGetBuffPreferenceReq) MsgId() int32 {
	return 216033
}

func (*PlanesTrialGetBuffPreferenceReq) Comment() string {
	return "查看预设buff"
}
		
func (*PlanesTrialGetBuffPreferenceRsp) MsgId() int32 {
	return 16033
}

func (*PlanesTrialGetBuffPreferenceRsp) Comment() string {
	return "查看预设buff"
}
		
func (*PlanesTrialGetGodBodyDataReq) MsgId() int32 {
	return 216034
}

func (*PlanesTrialGetGodBodyDataReq) Comment() string {
	return "获取分身数据"
}
		
func (*PlanesTrialGetGodBodyDataResp) MsgId() int32 {
	return 16034
}

func (*PlanesTrialGetGodBodyDataResp) Comment() string {
	return "获取分身数据"
}
		
func (*PlanesTrialGetGrandPrizeInfoReq) MsgId() int32 {
	return 216035
}

func (*PlanesTrialGetGrandPrizeInfoReq) Comment() string {
	return "获取大奖信息"
}
		
func (*PlanesTrialGetGrandPrizeInfoResp) MsgId() int32 {
	return 16035
}

func (*PlanesTrialGetGrandPrizeInfoResp) Comment() string {
	return "获取大奖信息"
}
		
func (*PlanesTrialReceiveGrandPrizeReq) MsgId() int32 {
	return 216036
}

func (*PlanesTrialReceiveGrandPrizeReq) Comment() string {
	return "领取大奖"
}
		
func (*PlanesTrialReceiveGrandPrizeResp) MsgId() int32 {
	return 16036
}

func (*PlanesTrialReceiveGrandPrizeResp) Comment() string {
	return "领取大奖"
}
		
func (*PlanesTrialEnterBattleReq) MsgId() int32 {
	return 216053
}

func (*PlanesTrialEnterBattleReq) Comment() string {
	return "进入布阵"
}
		
func (*PlanesTrialEnterBattleResp) MsgId() int32 {
	return 16053
}

func (*PlanesTrialEnterBattleResp) Comment() string {
	return "进入布阵"
}
		
func (*PlanesTrialEnterSwitchSeparationReq) MsgId() int32 {
	return 216054
}

func (*PlanesTrialEnterSwitchSeparationReq) Comment() string {
	return "进入切换分身界面"
}
		
func (*PlanesTrialEnterSwitchSeparationRsp) MsgId() int32 {
	return 16054
}

func (*PlanesTrialEnterSwitchSeparationRsp) Comment() string {
	return "进入切换分身界面"
}
		
func (*PlanesTrialSwitchSeparationReq) MsgId() int32 {
	return 216055
}

func (*PlanesTrialSwitchSeparationReq) Comment() string {
	return "切换上阵的分身"
}
		
func (*PlanesTrialSwitchSeparationRsp) MsgId() int32 {
	return 16055
}

func (*PlanesTrialSwitchSeparationRsp) Comment() string {
	return "切换上阵的分身"
}
		
func (*PlanesTrialSeparationDetailReq) MsgId() int32 {
	return 216056
}

func (*PlanesTrialSeparationDetailReq) Comment() string {
	return "查看上阵分身详细数据"
}
		
func (*PlanesTrialSeparationDetailRsp) MsgId() int32 {
	return 16056
}

func (*PlanesTrialSeparationDetailRsp) Comment() string {
	return "查看上阵分身详细数据"
}
		
func (*PlanesTrialChangePosReq) MsgId() int32 {
	return 216057
}

func (*PlanesTrialChangePosReq) Comment() string {
	return "变更参战玩家位置"
}
		
func (*PlanesTrialChangePosRsp) MsgId() int32 {
	return 16057
}

func (*PlanesTrialChangePosRsp) Comment() string {
	return "变更参战玩家位置"
}
		
func (*PlanesTrialDoBattleReq) MsgId() int32 {
	return 216058
}

func (*PlanesTrialDoBattleReq) Comment() string {
	return "进行战斗"
}
		
func (*PlanesTrialDoBattleRsp) MsgId() int32 {
	return 16058
}

func (*PlanesTrialDoBattleRsp) Comment() string {
	return "进行战斗"
}
		
func (*PlanesTrialUpdateLockDataReq) MsgId() int32 {
	return 216072
}

func (*PlanesTrialUpdateLockDataReq) Comment() string {
	return "更新绑定数据"
}
		
func (*PlanesTrialUpdateLockDataResp) MsgId() int32 {
	return 16072
}

func (*PlanesTrialUpdateLockDataResp) Comment() string {
	return "更新绑定数据"
}
		
func (*PlayerRestrainInfoMsgReq) MsgId() int32 {
	return 216088
}

func (*PlayerRestrainInfoMsgReq) Comment() string {
	return "获取被压制属性"
}
		
func (*PlayerRestrainInfoMsgResp) MsgId() int32 {
	return 16088
}

func (*PlayerRestrainInfoMsgResp) Comment() string {
	return "获取被压制属性"
}
		
func (*PlanesTrialRankGetReq) MsgId() int32 {
	return 216090
}

func (*PlanesTrialRankGetReq) Comment() string {
	return "获取排行榜"
}
		
func (*PlanesTrialRankGetResp) MsgId() int32 {
	return 16090
}

func (*PlanesTrialRankGetResp) Comment() string {
	return "获取排行榜"
}
		
func (*PlanesTrialHeartbeatReq) MsgId() int32 {
	return 216091
}

func (*PlanesTrialHeartbeatReq) Comment() string {
	return "心跳包"
}
		
func (*PlanesTrialHeartbeatResp) MsgId() int32 {
	return 16091
}

func (*PlanesTrialHeartbeatResp) Comment() string {
	return "心跳包"
}
		
func (*PlanesTrialGetSelectRewardDetailReq) MsgId() int32 {
	return 216092
}

func (*PlanesTrialGetSelectRewardDetailReq) Comment() string {
	return "获取自选奖励信息"
}
		
func (*PlanesTrialGetSelectRewardDetailResp) MsgId() int32 {
	return 16092
}

func (*PlanesTrialGetSelectRewardDetailResp) Comment() string {
	return "获取自选奖励信息"
}
		
func (*PlanesTrialSelectRewardReq) MsgId() int32 {
	return 216093
}

func (*PlanesTrialSelectRewardReq) Comment() string {
	return "选择奖励"
}
		
func (*PlanesTrialSelectRewardResp) MsgId() int32 {
	return 16093
}

func (*PlanesTrialSelectRewardResp) Comment() string {
	return "选择奖励"
}
		
func (*PlanesTrialVideoInfoReq) MsgId() int32 {
	return 216094
}

func (*PlanesTrialVideoInfoReq) Comment() string {
	return "查看关卡录像数据"
}
		
func (*PlanesTrialVideoInfoRsp) MsgId() int32 {
	return 16094
}

func (*PlanesTrialVideoInfoRsp) Comment() string {
	return "查看关卡录像数据"
}
		
func (*PlanesTrialPlayVideoReq) MsgId() int32 {
	return 216095
}

func (*PlanesTrialPlayVideoReq) Comment() string {
	return "播放关卡录像"
}
		
func (*PlanesTrialPlayVideoRsp) MsgId() int32 {
	return 16095
}

func (*PlanesTrialPlayVideoRsp) Comment() string {
	return "播放关卡录像"
}
		
func (*PlanesTrialCloseSettleScreenReq) MsgId() int32 {
	return 216096
}

func (*PlanesTrialCloseSettleScreenReq) Comment() string {
	return "播放关卡录像"
}
		
func (*PlanesTrialCloseSettleScreenResp) MsgId() int32 {
	return 16096
}

func (*PlanesTrialCloseSettleScreenResp) Comment() string {
	return "播放关卡录像"
}
		
func (*UnionTreasureDrawChipReq) MsgId() int32 {
	return 216201
}

func (*UnionTreasureDrawChipReq) Comment() string {
	return "妖盟寻宝地图互动"
}
		
func (*UnionTreasureDrawChipResp) MsgId() int32 {
	return 16201
}

func (*UnionTreasureDrawChipResp) Comment() string {
	return "妖盟寻宝地图互动"
}
		
func (*UnionTreasureAskForReq) MsgId() int32 {
	return 216202
}

func (*UnionTreasureAskForReq) Comment() string {
	return "妖盟寻宝索要"
}
		
func (*UnionTreasureAskForResp) MsgId() int32 {
	return 16202
}

func (*UnionTreasureAskForResp) Comment() string {
	return "妖盟寻宝索要"
}
		
func (*UnionTreasureSendChipReq) MsgId() int32 {
	return 216203
}

func (*UnionTreasureSendChipReq) Comment() string {
	return "妖盟寻宝碎片赠送"
}
		
func (*UnionTreasureSendChipResp) MsgId() int32 {
	return 16203
}

func (*UnionTreasureSendChipResp) Comment() string {
	return "妖盟寻宝碎片赠送"
}
		
func (*UnionTreasureBattleRelicReq) MsgId() int32 {
	return 216205
}

func (*UnionTreasureBattleRelicReq) Comment() string {
	return "妖盟寻宝遗迹互动"
}
		
func (*UnionTreasureBattleRelicResp) MsgId() int32 {
	return 16205
}

func (*UnionTreasureBattleRelicResp) Comment() string {
	return "妖盟寻宝遗迹互动"
}
		
func (*UnionTreasureHuntTreasureReq) MsgId() int32 {
	return 216206
}

func (*UnionTreasureHuntTreasureReq) Comment() string {
	return "妖盟寻宝遗迹挖宝"
}
		
func (*UnionTreasureHuntTreasureResp) MsgId() int32 {
	return 16206
}

func (*UnionTreasureHuntTreasureResp) Comment() string {
	return "妖盟寻宝遗迹挖宝"
}
		
func (*UnionTreasureEnterReq) MsgId() int32 {
	return 216207
}

func (*UnionTreasureEnterReq) Comment() string {
	return "妖盟寻宝进入界面"
}
		
func (*UnionTreasureEnterResp) MsgId() int32 {
	return 16207
}

func (*UnionTreasureEnterResp) Comment() string {
	return "妖盟寻宝进入界面"
}
		
func (*UnionTreasureGetGiveChipReq) MsgId() int32 {
	return 216209
}

func (*UnionTreasureGetGiveChipReq) Comment() string {
	return "妖盟寻宝获得碎片"
}
		
func (*UnionTreasureGetGiveChipResp) MsgId() int32 {
	return 16209
}

func (*UnionTreasureGetGiveChipResp) Comment() string {
	return "妖盟寻宝获得碎片"
}
		
func (*UnionTreasureAskForRelicReq) MsgId() int32 {
	return 216213
}

func (*UnionTreasureAskForRelicReq) Comment() string {
	return "妖盟寻宝遗迹数据请求"
}
		
func (*UnionTreasureAskForRelicResp) MsgId() int32 {
	return 16213
}

func (*UnionTreasureAskForRelicResp) Comment() string {
	return "妖盟寻宝遗迹数据请求"
}
		
func (*UnionTreasureShareRelicReq) MsgId() int32 {
	return 216216
}

func (*UnionTreasureShareRelicReq) Comment() string {
	return "妖盟寻宝分享协助"
}
		
func (*UnionTreasureShareRelicResp) MsgId() int32 {
	return 16216
}

func (*UnionTreasureShareRelicResp) Comment() string {
	return "妖盟寻宝分享协助"
}
		
func (*UnionTreasureFillChipReq) MsgId() int32 {
	return 216218
}

func (*UnionTreasureFillChipReq) Comment() string {
	return "妖盟寻宝遗迹激活碎片"
}
		
func (*UnionTreasureFillChipResp) MsgId() int32 {
	return 16218
}

func (*UnionTreasureFillChipResp) Comment() string {
	return "妖盟寻宝遗迹激活碎片"
}
		
func (*UnionTreasureGetBattleRecordReq) MsgId() int32 {
	return 216223
}

func (*UnionTreasureGetBattleRecordReq) Comment() string {
	return "妖盟寻宝战斗记录"
}
		
func (*UnionTreasureGetBattleRecordResp) MsgId() int32 {
	return 16223
}

func (*UnionTreasureGetBattleRecordResp) Comment() string {
	return "妖盟寻宝战斗记录"
}
		
func (*UnionTreasureGetCollectGiveMsgReq) MsgId() int32 {
	return 216224
}

func (*UnionTreasureGetCollectGiveMsgReq) Comment() string {
	return "妖盟寻宝赠送请求"
}
		
func (*UnionTreasureGetCollectGiveMsgResp) MsgId() int32 {
	return 16224
}

func (*UnionTreasureGetCollectGiveMsgResp) Comment() string {
	return "妖盟寻宝赠送请求"
}
		
func (*UnionTreasureExitReq) MsgId() int32 {
	return 216225
}

func (*UnionTreasureExitReq) Comment() string {
	return "妖盟寻宝退出请求"
}
		
func (*UnionTreasureExitResp) MsgId() int32 {
	return 16225
}

func (*UnionTreasureExitResp) Comment() string {
	return "妖盟寻宝退出请求"
}
		
func (*MeiTuanSubscribeSucceedReq) MsgId() int32 {
	return 216701
}

func (*MeiTuanSubscribeSucceedReq) Comment() string {
	return "美团订阅成功"
}
		
func (*MeiTuanSubscribeSucceedResp) MsgId() int32 {
	return 16701
}

func (*MeiTuanSubscribeSucceedResp) Comment() string {
	return "美团订阅成功"
}
		
func (*MeiTuanSubscribeRewardReq) MsgId() int32 {
	return 216702
}

func (*MeiTuanSubscribeRewardReq) Comment() string {
	return "美团订阅领取"
}
		
func (*MeiTuanSubscribeRewardResp) MsgId() int32 {
	return 16702
}

func (*MeiTuanSubscribeRewardResp) Comment() string {
	return "美团订阅领取"
}
		
func (*YueBaoEnterReq) MsgId() int32 {
	return 217001
}

func (*YueBaoEnterReq) Comment() string {
	return "仙玉宝府进入请求"
}
		
func (*YueBaoEnterResp) MsgId() int32 {
	return 17001
}

func (*YueBaoEnterResp) Comment() string {
	return "仙玉宝府进入请求"
}
		
func (*YueBaoDepositReq) MsgId() int32 {
	return 217002
}

func (*YueBaoDepositReq) Comment() string {
	return "仙玉宝府存储"
}
		
func (*YueBaoDepositResp) MsgId() int32 {
	return 17002
}

func (*YueBaoDepositResp) Comment() string {
	return "仙玉宝府存储"
}
		
func (*YueBaoInteractReq) MsgId() int32 {
	return 217003
}

func (*YueBaoInteractReq) Comment() string {
	return "仙玉宝府互动请求"
}
		
func (*YueBaoInteractResp) MsgId() int32 {
	return 17003
}

func (*YueBaoInteractResp) Comment() string {
	return "仙玉宝府互动请求"
}
		
func (*LiveShowNotifyReqMsg) MsgId() int32 {
	return 217701
}

func (*LiveShowNotifyReqMsg) Comment() string {
	return "同步直播状态"
}
		
func (*LoginOverMsg) MsgId() int32 {
	return 2
}

func (*LoginOverMsg) Comment() string {
	return "登录下发完成"
}
		
func (*NotifyRechargeSuccess) MsgId() int32 {
	return 5
}

func (*NotifyRechargeSuccess) Comment() string {
	return "通知充值到账"
}
		
func (*RandomNickNameResp) MsgId() int32 {
	return 7
}

func (*RandomNickNameResp) Comment() string {
	return "随机昵称回复"
}
		
func (*GetServerConfigVersionResp) MsgId() int32 {
	return 10
}

func (*GetServerConfigVersionResp) Comment() string {
	return "服务端配置"
}
		
func (*OtherLoginMsg) MsgId() int32 {
	return 11
}

func (*OtherLoginMsg) Comment() string {
	return "其他人登录你的账户"
}
		
func (*SystemRedPointSync) MsgId() int32 {
	return 12
}

func (*SystemRedPointSync) Comment() string {
	return "系统红点同步"
}
		
func (*ActivityRedPointSync) MsgId() int32 {
	return 13
}

func (*ActivityRedPointSync) Comment() string {
	return "活动红点同步"
}
		
func (*RechargeMallTimesSync) MsgId() int32 {
	return 14
}

func (*RechargeMallTimesSync) Comment() string {
	return "充值商品次数同步"
}
		
func (*RspConfigMsg) MsgId() int32 {
	return 16
}

func (*RspConfigMsg) Comment() string {
	return "同步后台配置"
}
		
func (*MiniGamesWestJourneyData) MsgId() int32 {
	return 25
}

func (*MiniGamesWestJourneyData) Comment() string {
	return "西天救援关卡数据同步"
}
		
func (*SyncRecallPlayerMsg) MsgId() int32 {
	return 32
}

func (*SyncRecallPlayerMsg) Comment() string {
	return "同步协议 回归玩家信息"
}
		
func (*PlayerDataMsg) MsgId() int32 {
	return 101
}

func (*PlayerDataMsg) Comment() string {
	return "用户信息同步"
}
		
func (*SystemUnlockSync) MsgId() int32 {
	return 102
}

func (*SystemUnlockSync) Comment() string {
	return "系统解锁同步"
}
		
func (*PrivilegeCardDataMsg) MsgId() int32 {
	return 104
}

func (*PrivilegeCardDataMsg) Comment() string {
	return "同步特权卡数据"
}
		
func (*PlayerCharaDataMsg) MsgId() int32 {
	return 106
}

func (*PlayerCharaDataMsg) Comment() string {
	return "玩家形象同步"
}
		
func (*GetPlayerCharaListResp) MsgId() int32 {
	return 108
}

func (*GetPlayerCharaListResp) Comment() string {
	return "获取形象回复"
}
		
func (*HorseRaceLampMsgAdd) MsgId() int32 {
	return 117
}

func (*HorseRaceLampMsgAdd) Comment() string {
	return "跑马灯广播通知"
}
		
func (*TitleSyncMsg) MsgId() int32 {
	return 119
}

func (*TitleSyncMsg) Comment() string {
	return "玩家称号同步"
}
		
func (*GoodsShieldSync) MsgId() int32 {
	return 121
}

func (*GoodsShieldSync) Comment() string {
	return "道具屏蔽列表同步"
}
		
func (*MedalSyncMsg) MsgId() int32 {
	return 122
}

func (*MedalSyncMsg) Comment() string {
	return "玩家勋章同步"
}
		
func (*ActivityPosterShieldSync) MsgId() int32 {
	return 124
}

func (*ActivityPosterShieldSync) Comment() string {
	return "活动海报屏蔽列表同步"
}
		
func (*MallRechargeIdShieldSync) MsgId() int32 {
	return 125
}

func (*MallRechargeIdShieldSync) Comment() string {
	return "同步渠道屏蔽商城付费项列表"
}
		
func (*RedPacketStateMsgSync) MsgId() int32 {
	return 140
}

func (*RedPacketStateMsgSync) Comment() string {
	return "红包状态同步"
}
		
func (*PlayerAttributeDataMsg) MsgId() int32 {
	return 201
}

func (*PlayerAttributeDataMsg) Comment() string {
	return "玩家属性信息同步"
}
		
func (*DreamDataMsg) MsgId() int32 {
	return 207
}

func (*DreamDataMsg) Comment() string {
	return "做梦玩家数据"
}
		
func (*PlayerAdRewardDataMsg) MsgId() int32 {
	return 210
}

func (*PlayerAdRewardDataMsg) Comment() string {
	return "数据同步"
}
		
func (*MessageSubscribeInfo) MsgId() int32 {
	return 250
}

func (*MessageSubscribeInfo) Comment() string {
	return "微信消息订阅登陆下发"
}
		
func (*SyncBagMsg) MsgId() int32 {
	return 301
}

func (*SyncBagMsg) Comment() string {
	return "同步背包数据"
}
		
func (*BattleRecordMsg) MsgId() int32 {
	return 401
}

func (*BattleRecordMsg) Comment() string {
	return "测试战斗"
}
		
func (*ChallengeRspMsg) MsgId() int32 {
	return 402
}

func (*ChallengeRspMsg) Comment() string {
	return "关卡挑战"
}
		
func (*PlayerStageData) MsgId() int32 {
	return 403
}

func (*PlayerStageData) Comment() string {
	return "关卡数据同步"
}
		
func (*ViewMonsterAttrResp) MsgId() int32 {
	return 404
}

func (*ViewMonsterAttrResp) Comment() string {
	return "查看关卡怪物信息"
}
		
func (*TeamBattleRecordMsg) MsgId() int32 {
	return 405
}

func (*TeamBattleRecordMsg) Comment() string {
	return "测试组队战斗"
}
		
func (*GetBattleListResp) MsgId() int32 {
	return 410
}

func (*GetBattleListResp) Comment() string {
	return "获取对战列表"
}
		
func (*RefreshBattleListResp) MsgId() int32 {
	return 411
}

func (*RefreshBattleListResp) Comment() string {
	return "刷新对战列表"
}
		
func (*GetRankBattleLogResp) MsgId() int32 {
	return 413
}

func (*GetRankBattleLogResp) Comment() string {
	return "获取妖榜日志"
}
		
func (*RankBattleSync) MsgId() int32 {
	return 415
}

func (*RankBattleSync) Comment() string {
	return "斗法登录同步"
}
		
func (*TaskDataListMsg) MsgId() int32 {
	return 501
}

func (*TaskDataListMsg) Comment() string {
	return "玩家登录任务数据下发"
}
		
func (*MailListMsg) MsgId() int32 {
	return 551
}

func (*MailListMsg) Comment() string {
	return "邮件列表数据同步"
}
		
func (*MallBuyCountListMsg) MsgId() int32 {
	return 602
}

func (*MallBuyCountListMsg) Comment() string {
	return "商品购买数量集合"
}
		
func (*TalentPlayerDataMsg) MsgId() int32 {
	return 621
}

func (*TalentPlayerDataMsg) Comment() string {
	return "同步神通数据"
}
		
func (*DestinyData) MsgId() int32 {
	return 651
}

func (*DestinyData) Comment() string {
	return "数据同步"
}
		
func (*PlayerCloudDataMsg) MsgId() int32 {
	return 711
}

func (*PlayerCloudDataMsg) Comment() string {
	return "筋斗云数据同步"
}
		
func (*WildBossDataSync) MsgId() int32 {
	return 731
}

func (*WildBossDataSync) Comment() string {
	return "同步宝地"
}
		
func (*PlayerPetDataSync) MsgId() int32 {
	return 740
}

func (*PlayerPetDataSync) Comment() string {
	return "同步玩家灵兽数据"
}
		
func (*TowerDataMsg) MsgId() int32 {
	return 761
}

func (*TowerDataMsg) Comment() string {
	return "同步数据"
}
		
func (*TowerChallengeResp) MsgId() int32 {
	return 762
}

func (*TowerChallengeResp) Comment() string {
	return "挑战"
}
		
func (*QuickChallengeResp) MsgId() int32 {
	return 763
}

func (*QuickChallengeResp) Comment() string {
	return "快速挑战"
}
		
func (*SpiritPlayerDataMsg) MsgId() int32 {
	return 821
}

func (*SpiritPlayerDataMsg) Comment() string {
	return "同步玩家精怪数据"
}
		
func (*PushActivityList) MsgId() int32 {
	return 1001
}

func (*PushActivityList) Comment() string {
	return "活动通用数据同步"
}
		
func (*ActivityCommonDataListSync) MsgId() int32 {
	return 1002
}

func (*ActivityCommonDataListSync) Comment() string {
	return "同步活动详细配置"
}
		
func (*ActivityConditionDataListSync) MsgId() int32 {
	return 1007
}

func (*ActivityConditionDataListSync) Comment() string {
	return "同步活动任务条件数据 增量"
}
		
func (*ActivityMallBuyCountDataListSync) MsgId() int32 {
	return 1008
}

func (*ActivityMallBuyCountDataListSync) Comment() string {
	return "同步活动商品购买次数列表 增量"
}
		
func (*ActivityPlayerDataSync) MsgId() int32 {
	return 1009
}

func (*ActivityPlayerDataSync) Comment() string {
	return "同步活动用户数据"
}
		
func (*ActivityScoreDataMsgSync) MsgId() int32 {
	return 1015
}

func (*ActivityScoreDataMsgSync) Comment() string {
	return "同步活动积分数据"
}
		
func (*SeekTreasureNoticeSyncInfo) MsgId() int32 {
	return 1026
}

func (*SeekTreasureNoticeSyncInfo) Comment() string {
	return "同步妖市觅宝抽奖公告"
}
		
func (*ActivityMallSelectDataListSync) MsgId() int32 {
	return 1027
}

func (*ActivityMallSelectDataListSync) Comment() string {
	return "活动商买自选商品列表同步"
}
		
func (*ActivityConditionSelectDataListSync) MsgId() int32 {
	return 1035
}

func (*ActivityConditionSelectDataListSync) Comment() string {
	return "同步活动任务奖励自选记录"
}
		
func (*ActivityShieldSync) MsgId() int32 {
	return 1044
}

func (*ActivityShieldSync) Comment() string {
	return "活动屏蔽同步"
}
		
func (*SyncHomelandMsg) MsgId() int32 {
	return 1051
}

func (*SyncHomelandMsg) Comment() string {
	return "同步信息"
}
		
func (*SyncHomelandHasRewardMsg) MsgId() int32 {
	return 1062
}

func (*SyncHomelandHasRewardMsg) Comment() string {
	return "同步有奖励消息"
}
		
func (*SyncPlayerHomelandChangeMsg) MsgId() int32 {
	return 1064
}

func (*SyncPlayerHomelandChangeMsg) Comment() string {
	return "同步福地数据"
}
		
func (*SyncPlayerCornucopiaChangeMsg) MsgId() int32 {
	return 1066
}

func (*SyncPlayerCornucopiaChangeMsg) Comment() string {
	return "福地聚宝盆数据同步"
}
		
func (*SyncLeaveHomeLand) MsgId() int32 {
	return 1072
}

func (*SyncLeaveHomeLand) Comment() string {
	return "福地通知玩家离开(只有偷其他玩家福地时才会通知)"
}
		
func (*BallGVGAttackedSyncMsg) MsgId() int32 {
	return 1116
}

func (*BallGVGAttackedSyncMsg) Comment() string {
	return "被打同步（有人被打即时同步）"
}
		
func (*BallGVGCampScoreSyncMsg) MsgId() int32 {
	return 1117
}

func (*BallGVGCampScoreSyncMsg) Comment() string {
	return "积分同步"
}
		
func (*BallGVGPlaceSyncMsg) MsgId() int32 {
	return 1118
}

func (*BallGVGPlaceSyncMsg) Comment() string {
	return "宝地信息同步（位置同步）"
}
		
func (*BallGVGNoticeSyncMsg) MsgId() int32 {
	return 1119
}

func (*BallGVGNoticeSyncMsg) Comment() string {
	return "公告推送"
}
		
func (*BallGVGCampSyncMsg) MsgId() int32 {
	return 1120
}

func (*BallGVGCampSyncMsg) Comment() string {
	return "同步宗门信息（任命、标记）"
}
		
func (*BallGVGAttackedUserSyncMsg) MsgId() int32 {
	return 1121
}

func (*BallGVGAttackedUserSyncMsg) Comment() string {
	return "被打同步最新用户信息(通知被打人)"
}
		
func (*BallGVGPlaceSeizeSyncMsg) MsgId() int32 {
	return 1122
}

func (*BallGVGPlaceSeizeSyncMsg) Comment() string {
	return "房间抢占信息同步(人数变动，进房间，出房间  成功占领)"
}
		
func (*BallGVGEnterPlaceSync) MsgId() int32 {
	return 1123
}

func (*BallGVGEnterPlaceSync) Comment() string {
	return "进房间同步信息（有人进房间及时同步）"
}
		
func (*BallGVGLeavePlaceSync) MsgId() int32 {
	return 1124
}

func (*BallGVGLeavePlaceSync) Comment() string {
	return "出房间同步信息（有人出房间及时同步）"
}
		
func (*BallGVGEndDataSyncMsg) MsgId() int32 {
	return 1127
}

func (*BallGVGEndDataSyncMsg) Comment() string {
	return "战斗结算同步"
}
		
func (*PlayerEmoticonDataSync) MsgId() int32 {
	return 1201
}

func (*PlayerEmoticonDataSync) Comment() string {
	return "同步表情包数据"
}
		
func (*InvadeChallengeResp) MsgId() int32 {
	return 1401
}

func (*InvadeChallengeResp) Comment() string {
	return "异兽入侵挑战"
}
		
func (*InvadeDataMsg) MsgId() int32 {
	return 1402
}

func (*InvadeDataMsg) Comment() string {
	return "异兽入侵数据同步"
}
		
func (*UnionRechargeDataSync) MsgId() int32 {
	return 1512
}

func (*UnionRechargeDataSync) Comment() string {
	return "商会充值活动数据同步"
}
		
func (*NotifyUnionRecruitCountMsg) MsgId() int32 {
	return 1655
}

func (*NotifyUnionRecruitCountMsg) Comment() string {
	return "同步招募次数"
}
		
func (*NotifyPlayerGradeMsg) MsgId() int32 {
	return 1656
}

func (*NotifyPlayerGradeMsg) Comment() string {
	return "登录或者评级变化时 同步玩家的评级数据"
}
		
func (*PushUnionEvictionMsg) MsgId() int32 {
	return 2119
}

func (*PushUnionEvictionMsg) Comment() string {
	return "推送被踢出妖盟广播"
}
		
func (*MyUnionData) MsgId() int32 {
	return 2124
}

func (*MyUnionData) Comment() string {
	return "推送我的妖盟数据更新"
}
		
func (*CutPriceDataMsg) MsgId() int32 {
	return 2165
}

func (*CutPriceDataMsg) Comment() string {
	return "砍价数据同步"
}
		
func (*UnionTimeMsg) MsgId() int32 {
	return 2168
}

func (*UnionTimeMsg) Comment() string {
	return "妖盟时间同步"
}
		
func (*SyncBlackPlayerIdListMsg) MsgId() int32 {
	return 3207
}

func (*SyncBlackPlayerIdListMsg) Comment() string {
	return "黑名单列表同步"
}
		
func (*RspHeroRankEnter) MsgId() int32 {
	return 3700
}

func (*RspHeroRankEnter) Comment() string {
	return "进入群英榜"
}
		
func (*SynHeroRankPlayerInfo) MsgId() int32 {
	return 3701
}

func (*SynHeroRankPlayerInfo) Comment() string {
	return "同步群英榜玩家信息"
}
		
func (*RspHeroRankRecord) MsgId() int32 {
	return 3705
}

func (*RspHeroRankRecord) Comment() string {
	return "请求日志数据"
}
		
func (*PlayerMagicDataMsg) MsgId() int32 {
	return 4400
}

func (*PlayerMagicDataMsg) Comment() string {
	return "神通数据同步"
}
		
func (*AllClientDataSync) MsgId() int32 {
	return 4502
}

func (*AllClientDataSync) Comment() string {
	return "全量同步储存数据"
}
		
func (*HoldPetEggDataSync) MsgId() int32 {
	return 4600
}

func (*HoldPetEggDataSync) Comment() string {
	return "同步灵兽蛋数据"
}
		
func (*HoldNestedBoxDataSync) MsgId() int32 {
	return 4601
}

func (*HoldNestedBoxDataSync) Comment() string {
	return "同步太初秘册数据"
}
		
func (*PalaceSyncMsg) MsgId() int32 {
	return 4801
}

func (*PalaceSyncMsg) Comment() string {
	return "仙宫数据同步"
}
		
func (*SendGiftSyncMsg) MsgId() int32 {
	return 4808
}

func (*SendGiftSyncMsg) Comment() string {
	return "仙宫送福数据同步"
}
		
func (*PalaceMiracleDataMsg) MsgId() int32 {
	return 4809
}

func (*PalaceMiracleDataMsg) Comment() string {
	return "仙宫神迹同步"
}
		
func (*ManHuangEnemyNotify) MsgId() int32 {
	return 5022
}

func (*ManHuangEnemyNotify) Comment() string {
	return "蛮荒 广播-仇人信息"
}
		
func (*ManHuangTeamLeaderNotify) MsgId() int32 {
	return 5023
}

func (*ManHuangTeamLeaderNotify) Comment() string {
	return "蛮荒 广播-队长收到的广播"
}
		
func (*ManHuangTeamMemberNotify) MsgId() int32 {
	return 5024
}

func (*ManHuangTeamMemberNotify) Comment() string {
	return "蛮荒 广播-队员接收到的广播"
}
		
func (*ManHuangEventInfoNotify) MsgId() int32 {
	return 5029
}

func (*ManHuangEventInfoNotify) Comment() string {
	return "蛮荒 事件广播"
}
		
func (*ManHuangMarqueeNotify) MsgId() int32 {
	return 5031
}

func (*ManHuangMarqueeNotify) Comment() string {
	return "蛮荒 击杀广播"
}
		
func (*SynPlayerInfo) MsgId() int32 {
	return 5311
}

func (*SynPlayerInfo) Comment() string {
	return "同步玩家用户信息"
}
		
func (*EquipmentAdvanceDataMsg) MsgId() int32 {
	return 5504
}

func (*EquipmentAdvanceDataMsg) Comment() string {
	return "数据同步"
}
		
func (*SynSecretTowerInfo) MsgId() int32 {
	return 5605
}

func (*SynSecretTowerInfo) Comment() string {
	return "秘境数据同步"
}
		
func (*UnionBossMsg) MsgId() int32 {
	return 5807
}

func (*UnionBossMsg) Comment() string {
	return "同步boss信息"
}
		
func (*UnionBossBuff) MsgId() int32 {
	return 5808
}

func (*UnionBossBuff) Comment() string {
	return "同步阵法信息"
}
		
func (*AskDingSyncPlayerJoinMsg) MsgId() int32 {
	return 5938
}

func (*AskDingSyncPlayerJoinMsg) Comment() string {
	return "玩家参赛后 同步玩家参赛资格"
}
		
func (*SyncUnionAreaWarDragonAttackMsg) MsgId() int32 {
	return 6028
}

func (*SyncUnionAreaWarDragonAttackMsg) Comment() string {
	return "神龙战场攻击推送"
}
		
func (*MagicTreasurePlayerDataMsg) MsgId() int32 {
	return 6301
}

func (*MagicTreasurePlayerDataMsg) Comment() string {
	return "玩家法宝数据"
}
		
func (*SyncMagicTreasureDataMsg) MsgId() int32 {
	return 6306
}

func (*SyncMagicTreasureDataMsg) Comment() string {
	return "同步法宝列表数据"
}
		
func (*UnionFightPrepareDataSync) MsgId() int32 {
	return 6721
}

func (*UnionFightPrepareDataSync) Comment() string {
	return "妖盟夺位战备战数据同步"
}
		
func (*UnionFightApplyDataSync) MsgId() int32 {
	return 6730
}

func (*UnionFightApplyDataSync) Comment() string {
	return "妖盟夺位战登录数据"
}
		
func (*GodIslandGameRouteInfoSyncMsg) MsgId() int32 {
	return 6822
}

func (*GodIslandGameRouteInfoSyncMsg) Comment() string {
	return "游戏路线同步信息"
}
		
func (*GodIslandGameCityBuffSync) MsgId() int32 {
	return 6825
}

func (*GodIslandGameCityBuffSync) Comment() string {
	return "城市buff同步"
}
		
func (*GodIslandGameEventSyncMsg) MsgId() int32 {
	return 6833
}

func (*GodIslandGameEventSyncMsg) Comment() string {
	return "npc事件同步"
}
		
func (*GodIslandGameCityChangeSyncMsg) MsgId() int32 {
	return 6846
}

func (*GodIslandGameCityChangeSyncMsg) Comment() string {
	return "城市变化同步"
}
		
func (*GodIslandGamePlayerWinSyncMsg) MsgId() int32 {
	return 6847
}

func (*GodIslandGamePlayerWinSyncMsg) Comment() string {
	return "连胜消息同步"
}
		
func (*GodIslandGameLineInfoSyncMsg) MsgId() int32 {
	return 6848
}

func (*GodIslandGameLineInfoSyncMsg) Comment() string {
	return "队列消息同步(大地图，城市队列，具体队列信息都通过这条去同步刷新)"
}
		
func (*GodIslandGameMyGameInfoSyncMsg) MsgId() int32 {
	return 6850
}

func (*GodIslandGameMyGameInfoSyncMsg) Comment() string {
	return "我的游戏信息同步"
}
		
func (*GodIslandGameHorseLampSyncMsg) MsgId() int32 {
	return 6853
}

func (*GodIslandGameHorseLampSyncMsg) Comment() string {
	return "跑马灯消息同步"
}
		
func (*GodIslandSuppressBuffSyncMsg) MsgId() int32 {
	return 6854
}

func (*GodIslandSuppressBuffSyncMsg) Comment() string {
	return "玩家的所向披靡同步"
}
		
func (*GodIslandLeftOutSyncInfo) MsgId() int32 {
	return 6857
}

func (*GodIslandLeftOutSyncInfo) Comment() string {
	return "场内有妖盟被淘汰 同步下来"
}
		
func (*GodIslandSyncFriendUnionMsg) MsgId() int32 {
	return 6859
}

func (*GodIslandSyncFriendUnionMsg) Comment() string {
	return "盟友同步"
}
		
func (*GodIslandCommanderChangeSyncMsg) MsgId() int32 {
	return 6860
}

func (*GodIslandCommanderChangeSyncMsg) Comment() string {
	return "指挥权变化同步"
}
		
func (*GodIslandBeenKillSyncMsg) MsgId() int32 {
	return 6861
}

func (*GodIslandBeenKillSyncMsg) Comment() string {
	return "被击杀后同步"
}
		
func (*GodIslandGameCityBattleSyncMsg) MsgId() int32 {
	return 6862
}

func (*GodIslandGameCityBattleSyncMsg) Comment() string {
	return "城市战斗状态"
}
		
func (*GodIslandAutoBattleStopSyncMsg) MsgId() int32 {
	return 6863
}

func (*GodIslandAutoBattleStopSyncMsg) Comment() string {
	return "托管操作 结束后同步给客户端"
}
		
func (*GodIslandGhostCityUnlockSyncMsg) MsgId() int32 {
	return 6864
}

func (*GodIslandGhostCityUnlockSyncMsg) Comment() string {
	return "修罗城解锁同步"
}
		
func (*StarTrialDataMsg) MsgId() int32 {
	return 6901
}

func (*StarTrialDataMsg) Comment() string {
	return "星宿试炼数据同步"
}
		
func (*SyncGatherEnergyMsg) MsgId() int32 {
	return 7002
}

func (*SyncGatherEnergyMsg) Comment() string {
	return "同步聚灵阵数据"
}
		
func (*GatherEnergyRewardShowResp) MsgId() int32 {
	return 7010
}

func (*GatherEnergyRewardShowResp) Comment() string {
	return "聚灵阵奖励推送"
}
		
func (*ReceiveSdkRewardSyn) MsgId() int32 {
	return 7201
}

func (*ReceiveSdkRewardSyn) Comment() string {
	return "同步sdk奖励数据"
}
		
func (*SkyWarDataSync) MsgId() int32 {
	return 8413
}

func (*SkyWarDataSync) Comment() string {
	return "数据同步"
}
		
func (*SkyWarDataLoginSync) MsgId() int32 {
	return 8414
}

func (*SkyWarDataLoginSync) Comment() string {
	return "征战诸天数据登录同步，阶段变化同步1"
}
		
func (*SyncMarkStateMsg) MsgId() int32 {
	return 8601
}

func (*SyncMarkStateMsg) Comment() string {
	return "同步评分状态"
}
		
func (*WorldRulePlayerDataMsg) MsgId() int32 {
	return 9005
}

func (*WorldRulePlayerDataMsg) Comment() string {
	return "天地法则玩家数据同步"
}
		
func (*RuleTrialBossConfigSyncMsg) MsgId() int32 {
	return 9104
}

func (*RuleTrialBossConfigSyncMsg) Comment() string {
	return "同步法则试炼Boss妖力"
}
		
func (*RuleTrialDataSync) MsgId() int32 {
	return 9105
}

func (*RuleTrialDataSync) Comment() string {
	return "同步法则试炼"
}
		
func (*HolyLandBattleTimeStampsDataSync) MsgId() int32 {
	return 9575
}

func (*HolyLandBattleTimeStampsDataSync) Comment() string {
	return "九幽争霸赛事时间戳数据登录前同步"
}
		
func (*MountainSeaTeamLeaderNotify) MsgId() int32 {
	return 9731
}

func (*MountainSeaTeamLeaderNotify) Comment() string {
	return "队长接受到广播"
}
		
func (*MountainSeaTeamMemberNotify) MsgId() int32 {
	return 9732
}

func (*MountainSeaTeamMemberNotify) Comment() string {
	return "队员接收到的广播"
}
		
func (*MountainSeaEnterBattleNotify) MsgId() int32 {
	return 9770
}

func (*MountainSeaEnterBattleNotify) Comment() string {
	return "广播进入布阵界面"
}
		
func (*MountainSeaSwitchSeparationNotify) MsgId() int32 {
	return 9771
}

func (*MountainSeaSwitchSeparationNotify) Comment() string {
	return "切换上阵分身广播"
}
		
func (*MountainSeaChangeTeamSkillNotify) MsgId() int32 {
	return 9772
}

func (*MountainSeaChangeTeamSkillNotify) Comment() string {
	return "装配的队伍技能变更广播"
}
		
func (*MountainSeaDoBattleNotify) MsgId() int32 {
	return 9773
}

func (*MountainSeaDoBattleNotify) Comment() string {
	return "广播战斗数据"
}
		
func (*MountainSeaChangePosNotify) MsgId() int32 {
	return 9774
}

func (*MountainSeaChangePosNotify) Comment() string {
	return "参战对象位置变更广播"
}
		
func (*MountainSeaRedPointResp) MsgId() int32 {
	return 9784
}

func (*MountainSeaRedPointResp) Comment() string {
	return "请求红点"
}
		
func (*PlayerSoulLiQuidDataMsg) MsgId() int32 {
	return 9801
}

func (*PlayerSoulLiQuidDataMsg) Comment() string {
	return "琼浆玉液数据同步"
}
		
func (*SkyPresentDataSync) MsgId() int32 {
	return 9910
}

func (*SkyPresentDataSync) Comment() string {
	return "天道馈赠数据同步"
}
		
func (*UnionBlessingGiftSyncList) MsgId() int32 {
	return 10603
}

func (*UnionBlessingGiftSyncList) Comment() string {
	return "同步妖盟散财数据"
}
		
func (*UnionBlessingCountSyncList) MsgId() int32 {
	return 10604
}

func (*UnionBlessingCountSyncList) Comment() string {
	return "同步自己妖盟的散财的旗帜和时间"
}
		
func (*FestivalCelebrationsPlayerInfoMsg) MsgId() int32 {
	return 11000
}

func (*FestivalCelebrationsPlayerInfoMsg) Comment() string {
	return "庆典玩家数据同步"
}
		
func (*DoubleDemonsDataLoginSync) MsgId() int32 {
	return 11111
}

func (*DoubleDemonsDataLoginSync) Comment() string {
	return "双妖成行登录同步数据"
}
		
func (*NewYearRedBagPlayerData) MsgId() int32 {
	return 11500
}

func (*NewYearRedBagPlayerData) Comment() string {
	return "新春红包 登录数据下发"
}
		
func (*NewYearSyncNoticeRespMsg) MsgId() int32 {
	return 11504
}

func (*NewYearSyncNoticeRespMsg) Comment() string {
	return "新春红包 同步公告"
}
		
func (*PetKernelPlayerDataMsg) MsgId() int32 {
	return 11708
}

func (*PetKernelPlayerDataMsg) Comment() string {
	return "内丹同步(商店是否开启 pieceShopOpen只有这个字段有用)"
}
		
func (*PupilSystemLoginSync) MsgId() int32 {
	return 11813
}

func (*PupilSystemLoginSync) Comment() string {
	return "门徒系统登录同步"
}
		
func (*MarriageUserDataMsgSync) MsgId() int32 {
	return 11850
}

func (*MarriageUserDataMsgSync) Comment() string {
	return "结义用户数据同步"
}
		
func (*MarriageGetAppointApplySync) MsgId() int32 {
	return 11851
}

func (*MarriageGetAppointApplySync) Comment() string {
	return "同步其他玩家发起的结义请求（别的玩家指定你结义的列表）"
}
		
func (*PupilGraduatedUnMarryListSync) MsgId() int32 {
	return 11852
}

func (*PupilGraduatedUnMarryListSync) Comment() string {
	return "同步出师未结义的门徒列表"
}
		
func (*MarriageRecordTempSync) MsgId() int32 {
	return 11853
}

func (*MarriageRecordTempSync) Comment() string {
	return "结义记录数据同步"
}
		
func (*MarriageRefuseNotifyMsgSync) MsgId() int32 {
	return 11854
}

func (*MarriageRefuseNotifyMsgSync) Comment() string {
	return "拒绝结义通知同步"
}
		
func (*MarriageGetAppointApplyTimeSync) MsgId() int32 {
	return 11865
}

func (*MarriageGetAppointApplyTimeSync) Comment() string {
	return "同步其他玩家发起的结义请求时间"
}
		
func (*CloudRefinePlayerDataMsg) MsgId() int32 {
	return 12900
}

func (*CloudRefinePlayerDataMsg) Comment() string {
	return "筋斗云注灵数据同步"
}
		
func (*MonopolyBarrageNotify) MsgId() int32 {
	return 13130
}

func (*MonopolyBarrageNotify) Comment() string {
	return "弹幕推送"
}
		
func (*MonopolyMoveNotify) MsgId() int32 {
	return 13131
}

func (*MonopolyMoveNotify) Comment() string {
	return "同妖盟移动同步"
}
		
func (*MonopolyNotifyPlayerEndTrap) MsgId() int32 {
	return 13132
}

func (*MonopolyNotifyPlayerEndTrap) Comment() string {
	return "如果队友主动或者被协助 挣脱了陷阱 同步"
}
		
func (*MonopolyNotifyBuildingUpgrade) MsgId() int32 {
	return 13133
}

func (*MonopolyNotifyBuildingUpgrade) Comment() string {
	return "触发建筑升级的推送"
}
		
func (*MonopolySendBlessingNotify) MsgId() int32 {
	return 13135
}

func (*MonopolySendBlessingNotify) Comment() string {
	return "天赐福源推送"
}
		
func (*MonopolyScoreNotify) MsgId() int32 {
	return 13136
}

func (*MonopolyScoreNotify) Comment() string {
	return "妖盟积分变动"
}
		
func (*MonopolyRankNotify) MsgId() int32 {
	return 13137
}

func (*MonopolyRankNotify) Comment() string {
	return "同步排名信息"
}
		
func (*MonopolyMyBlessingNotify) MsgId() int32 {
	return 13139
}

func (*MonopolyMyBlessingNotify) Comment() string {
	return "天赐福源推送-我的"
}
		
func (*MonopolyAssistRedNotify) MsgId() int32 {
	return 13140
}

func (*MonopolyAssistRedNotify) Comment() string {
	return "协助红点推送"
}
		
func (*MonopolyPlayerScoreChangeNotify) MsgId() int32 {
	return 13144
}

func (*MonopolyPlayerScoreChangeNotify) Comment() string {
	return "协助击杀后积分推送"
}
		
func (*TownDemonTimeStampsDataSync) MsgId() int32 {
	return 13251
}

func (*TownDemonTimeStampsDataSync) Comment() string {
	return "镇魔登录时间戳同步"
}
		
func (*PeachBanquetOpenSync) MsgId() int32 {
	return 13510
}

func (*PeachBanquetOpenSync) Comment() string {
	return "宴会开启同步(类仙宫送福)"
}
		
func (*PeachBanquetLoginSync) MsgId() int32 {
	return 13515
}

func (*PeachBanquetLoginSync) Comment() string {
	return "蟠桃宴登录同步"
}
		
func (*PeachBanquetItemSync) MsgId() int32 {
	return 13519
}

func (*PeachBanquetItemSync) Comment() string {
	return "蟠桃宴道具同步"
}
		
func (*UnionBountySynMonsterInfoMsg) MsgId() int32 {
	return 13625
}

func (*UnionBountySynMonsterInfoMsg) Comment() string {
	return "妖盟悬赏-同步妖兽房间信息"
}
		
func (*SyncFriendList) MsgId() int32 {
	return 14130
}

func (*SyncFriendList) Comment() string {
	return "同步世交信息列表"
}
		
func (*SyncFriendInfo) MsgId() int32 {
	return 14132
}

func (*SyncFriendInfo) Comment() string {
	return "同步世交信息"
}
		
func (*PlayerFairyLandDataMsg) MsgId() int32 {
	return 14202
}

func (*PlayerFairyLandDataMsg) Comment() string {
	return "飞升仙界-同步玩家数据"
}
		
func (*FairyLandServerSync) MsgId() int32 {
	return 14211
}

func (*FairyLandServerSync) Comment() string {
	return "飞升仙界-区服列表同步"
}
		
func (*UniverseDataMsgSync) MsgId() int32 {
	return 14302
}

func (*UniverseDataMsgSync) Comment() string {
	return "同步小世界信息"
}
		
func (*BeingSnatchedSync) MsgId() int32 {
	return 14699
}

func (*BeingSnatchedSync) Comment() string {
	return "画中仙境被抢夺同步"
}
		
func (*SkyTradeUnionDataSync) MsgId() int32 {
	return 15050
}

func (*SkyTradeUnionDataSync) Comment() string {
	return "妖盟分数积分变更"
}
		
func (*SkyTradeResetSync) MsgId() int32 {
	return 15051
}

func (*SkyTradeResetSync) Comment() string {
	return "货物刷新"
}
		
func (*SkyTradeAirshipInfoSync) MsgId() int32 {
	return 15052
}

func (*SkyTradeAirshipInfoSync) Comment() string {
	return "同步地图盟友的航行信息"
}
		
func (*SkyTradeChallengeInfoSync) MsgId() int32 {
	return 15053
}

func (*SkyTradeChallengeInfoSync) Comment() string {
	return "同步挑战信息"
}
		
func (*SkyTradeRareGoodsSync) MsgId() int32 {
	return 15054
}

func (*SkyTradeRareGoodsSync) Comment() string {
	return "同步紧缺货物信息"
}
		
func (*SkyTradeFameRankSync) MsgId() int32 {
	return 15055
}

func (*SkyTradeFameRankSync) Comment() string {
	return "同步威望排名"
}
		
func (*YardLoginSync) MsgId() int32 {
	return 15843
}

func (*YardLoginSync) Comment() string {
	return "家园登录同步"
}
		
func (*YardBuildInfoMsg) MsgId() int32 {
	return 15844
}

func (*YardBuildInfoMsg) Comment() string {
	return "装饰建筑等级同步"
}
		
func (*YardPeachHelpDataSync) MsgId() int32 {
	return 15845
}

func (*YardPeachHelpDataSync) Comment() string {
	return "桃树协助信息同步"
}
		
func (*YardBaseMsgSync) MsgId() int32 {
	return 15847
}

func (*YardBaseMsgSync) Comment() string {
	return "家园基础数据同步"
}
		
func (*YardBuildInfoMsgSync) MsgId() int32 {
	return 15848
}

func (*YardBuildInfoMsgSync) Comment() string {
	return "家园建筑信息同步"
}
		
func (*YardMakeMsgSync) MsgId() int32 {
	return 15849
}

func (*YardMakeMsgSync) Comment() string {
	return "家园生产信息同步"
}
		
func (*YardRefreshDataSync) MsgId() int32 {
	return 15850
}

func (*YardRefreshDataSync) Comment() string {
	return "家园每日数据同步"
}
		
func (*YardBuildUpSync) MsgId() int32 {
	return 15851
}

func (*YardBuildUpSync) Comment() string {
	return "家园功能建筑升级同步"
}
		
func (*PlanesTrialTeamLeaderNotify) MsgId() int32 {
	return 16040
}

func (*PlanesTrialTeamLeaderNotify) Comment() string {
	return "队长接受到广播"
}
		
func (*PlanesTrialTeamMemberNotify) MsgId() int32 {
	return 16041
}

func (*PlanesTrialTeamMemberNotify) Comment() string {
	return "队员接收到的广播"
}
		
func (*PlanesTrialDoBattleNotify) MsgId() int32 {
	return 16065
}

func (*PlanesTrialDoBattleNotify) Comment() string {
	return "通知进入战斗"
}
		
func (*PlanesTrialEnterBattleNotify) MsgId() int32 {
	return 16070
}

func (*PlanesTrialEnterBattleNotify) Comment() string {
	return "广播进入布阵界面"
}
		
func (*PlanesTrialSwitchSeparationNotify) MsgId() int32 {
	return 16071
}

func (*PlanesTrialSwitchSeparationNotify) Comment() string {
	return "通知上阵的分身变更"
}
		
func (*PlanesTrialChangePosNotify) MsgId() int32 {
	return 16074
}

func (*PlanesTrialChangePosNotify) Comment() string {
	return "通知参战玩家位置变更"
}
		
func (*UnionTreasureCollectMsg) MsgId() int32 {
	return 16208
}

func (*UnionTreasureCollectMsg) Comment() string {
	return "妖盟寻宝主要数据同步"
}
		
func (*UnionTreasuretLoginSync) MsgId() int32 {
	return 16220
}

func (*UnionTreasuretLoginSync) Comment() string {
	return "妖盟寻宝时间同步"
}
		
func (*UnionTreasureLvUpWindowsMsg) MsgId() int32 {
	return 16227
}

func (*UnionTreasureLvUpWindowsMsg) Comment() string {
	return "妖盟寻宝升级弹窗"
}
		
func (*SyncMeiTuanSubscribeMsg) MsgId() int32 {
	return 16700
}

func (*SyncMeiTuanSubscribeMsg) Comment() string {
	return "同步美团订阅数据"
}
		
func (*RechargeItemData) MsgId() int32 {
	return 17004
}

func (*RechargeItemData) Comment() string {
	return "仙玉宝府信息同步"
}
		
func (*MemoryCollectLoginSyncData) MsgId() int32 {
	return 17413
}

func (*MemoryCollectLoginSyncData) Comment() string {
	return "集卡登录同步"
}
		
func (*LiveShowNotifyRespMsg) MsgId() int32 {
	return 17701
}

func (*LiveShowNotifyRespMsg) Comment() string {
	return "同步直播状态"
}
		