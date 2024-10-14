package assist

import (
	"context"
	"xddq/pkg/game/client"
	"xddq/pkg/game/protocol/pb"
	"xddq/pkg/utils"
)

type AskWayManager struct {
	a              *Assistant
	activityId     int32
	competitionId  int32
	challengeTimes int32
	challengeInfo  []*pb.AskWayPlayerChallengePlayerInfo
}

func NewAskWayManager(a *Assistant) *AskWayManager {
	am := &AskWayManager{
		a: a,
	}

	a.addRunner(am)

	a.Client().OnAskWayEnterReq(am.onAskWayEnterRsp)

	return am
}

func (am *AskWayManager) Run(ctx context.Context) error {
	am.activityId = am.a.activity.GetActivityIdByName("问道盛会")

	return nil
}

func (am *AskWayManager) Stop() error {
	return nil
}

// enter
func (am *AskWayManager) enter() error {
	return am.a.Client().SendAskWayEnterReq(am.a.Context(), &pb.AskWayEnterReq{
		ActivityId: utils.Ptr(am.activityId),
	})
}

// onAskWayEnterRsp
func (am *AskWayManager) onAskWayEnterRsp(ctx client.Context, msg *pb.AskWayEnterRsp) error {
	playerData := msg.GetPlayerData()
	am.competitionId = playerData.GetCompetitionArea()
	am.challengeTimes = playerData.GetChallengeTimes()
	am.challengeInfo = playerData.GetChallengeInfo().GetPlayerList()

	return nil
}
