package assist

import (
	"context"
	"math/rand/v2"
	"slices"

	"xddq/pkg/game/client"
	"xddq/pkg/game/db"
	"xddq/pkg/game/protocol/pb"
	"xddq/pkg/utils"

	"github.com/samber/lo"
	"go.uber.org/zap"
)

type UniverseManager struct {
	a               *Assistant
	stoneNum        int32
	freeDrawTimes   int32
	drawTwiceStatus int32

	selectedPos []int32
}

func NewUniverseManager(a *Assistant) *UniverseManager {
	u := &UniverseManager{
		a: a,
	}

	if a.Config().Universe {
		a.addRunner(u)
		a.Client().OnUniverseDataMsgSync(u.onUniverseDataMsgSync)
		a.Client().OnUniverseDrawTwiceReq(u.onUniverseDrawTwiceResp)
		a.Client().OnUniverseDrawReq(u.onUniverseDrawResp)
	}

	return u
}

func (u *UniverseManager) Run(ctx context.Context) error {
	u.a.addJob("universeCheck", "@every 1s", u.check)
	return nil
}

func (u *UniverseManager) Stop() error {
	return nil
}

func (u *UniverseManager) check(ctx context.Context) error {
	if u.stoneNum > 0 {
		u.draw(ctx, u.stoneNum)
	}

	if u.drawTwiceStatus > 1 {
		u.drawTwice(ctx)
	}

	return nil
}

func (u *UniverseManager) drawTwice(ctx context.Context) error {
	return u.a.Client().SendUniverseDrawTwiceReq(ctx, &pb.UniverseDrawTwiceReq{
		Pos: utils.Ptr(u.randPos()),
	})
}

func (u *UniverseManager) draw(ctx context.Context, num int32) error {
	return u.a.Client().SendUniverseDrawReq(ctx, &pb.UniverseDrawReq{
		Bei: &num,
	})
}

// levelUp
func (u *UniverseManager) levelUp(ctx client.Context) error {
	return u.a.Client().SendUniverseLevelUpReq(ctx, &pb.UniverseLevelUpReq{})
}

// onUniversLevelUpResp
func (u *UniverseManager) onUniversLevelUpResp(_ client.Context, msg *pb.UniverseLevelUpResp) error {
	u.a.L().Info("onUniversLevelUpResp", zap.Any("msg", msg))
	return nil
}

func (u *UniverseManager) randPos() (pos int32) {
	if u.drawTwiceStatus == 2 {
		return 1 + rand.Int32N(3)
	}

	pos = 1 + rand.Int32N(12)
	for slices.Contains(u.selectedPos, pos) {
		pos = 1 + rand.Int32N(12)
	}

	return
}

// onUniverseDataMsgSync
func (u *UniverseManager) onUniverseDataMsgSync(ctx client.Context, msg *pb.UniverseDataMsgSync) error {
	um := msg.UniverseDataMsg
	u.stoneNum = um.GetStoneNum()
	u.freeDrawTimes = um.GetFreeDrawTimes()
	u.drawTwiceStatus = um.GetDrawTwiceData().GetDrawStatus()
	u.selectedPos = lo.Map(um.GetDrawTwiceData().GetDrawMap(), func(item *pb.UniverseDrawMap, _ int) int32 {
		return item.GetPos()
	})

	return nil
}

func (u *UniverseManager) onUniverseDrawResp(ctx client.Context, msg *pb.UniverseDrawResp) error {
	return db.GameError(msg)
}

func (u *UniverseManager) onUniverseDrawTwiceResp(ctx client.Context, msg *pb.UniverseDrawTwiceResp) error {
	if msg.GetRet() != 0 {
		return db.GameError(msg)
	}
	if msg.GetRewards() != "" {
		u.selectedPos = nil
		u.a.L().Info("onUniverseDrawTwiceResp", zap.Stringer("rewards", db.ParseRewards(msg.GetRewards())))
		return nil
	}

	u.selectedPos = append(u.selectedPos, msg.GetDrawMap().GetPos())

	return nil
}
