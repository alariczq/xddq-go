package assist

import (
	"xddq/pkg/game/client"
	"xddq/pkg/game/protocol/pb"
)

func (wb *ChallengeManager) willBossMaxTimes() (times int32) {
	if times = 6; wb.a.PlayerInfo().IsMonthCardVip() {
		times += 2
	}
	return
}

// onWildBossDataSync 妖王
func (wb *ChallengeManager) onWildBossDataSync(ctx client.Context, msg *pb.WildBossDataSync) error {
	wb.willBossPassId = msg.GetData().GetPassId()
	wb.willBossTimes = msg.GetData().GetUseRepeatTimes()
	for i := wb.willBossTimes; i < wb.willBossMaxTimes(); i++ {
		wb.a.Client().SendWildBossRepeatReq(ctx, &pb.WildBossRepeatReq{})
	}

	return nil
}
