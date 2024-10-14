package assist

import (
	"context"
	"time"

	"xddq/pkg/game/client"
	"xddq/pkg/game/db"
	"xddq/pkg/game/protocol/pb"
	"xddq/pkg/utils"
)

type YuebaoManager struct {
	a       *Assistant
	endTime int64
}

func NewYuebaoManager(a *Assistant) *YuebaoManager {
	y := &YuebaoManager{
		a: a,
	}

	if a.Config().Yuebao {
		a.addRunner(y)

		a.Client().OnYueBaoEnterReq(y.onYueBaoEnterResp)
		a.Client().OnYueBaoInteractReq(y.onYueBaoInteractResp)
		a.Client().OnYueBaoDepositReq(y.onYueBaoDepositResp)
	}

	return y
}

func (y *YuebaoManager) Run(ctx context.Context) error {
	if !y.a.Config().Yuebao {
		return nil
	}

	y.enter(ctx)
	y.a.addJob("yuebao", "@every 1m", func(ctx context.Context) error {
		if time.Now().UnixMilli() < y.endTime {
			return nil
		}

		return y.enter(ctx)
	})

	return nil
}

func (y *YuebaoManager) activityId() int32 {
	return y.a.activity.GetActivityIdByName("呱仙阁")
}

func (y *YuebaoManager) enter(ctx context.Context) error {
	return y.a.Client().SendYueBaoEnterReq(ctx, &pb.YueBaoEnterReq{ActivityId: utils.Ptr(y.activityId())})
}

func (y *YuebaoManager) Stop() error {
	return nil
}

func (y *YuebaoManager) onYueBaoEnterResp(ctx client.Context, msg *pb.YueBaoEnterResp) error {
	if msg.GetRet() != 0 {
		return db.GetGameError(msg.GetRet())
	}

	if y.endTime = msg.GetPlayerData().GetEndTime(); y.endTime == 0 {
		return y.deposit(ctx)
	}

	if now := time.Now().UnixMilli(); now > y.endTime {
		ctx.Client().SendYueBaoInteractReq(ctx, &pb.YueBaoInteractReq{
			ActivityId: utils.Ptr(y.activityId()),
		})
	}

	return nil
}

// deposit 存钱
func (y *YuebaoManager) deposit(ctx client.Context) error {
	return ctx.Client().SendYueBaoDepositReq(ctx, &pb.YueBaoDepositReq{
		ActivityId: utils.Ptr(y.activityId()),
		Index:      utils.Ptr(int32(1)),
		DepositNum: utils.Ptr(int32(3000)),
	})
}

func (y *YuebaoManager) onYueBaoInteractResp(ctx client.Context, msg *pb.YueBaoInteractResp) error {
	if msg.GetRet() != 0 {
		return db.GetGameError(msg.GetRet())
	}

	return y.deposit(ctx)
}

func (y *YuebaoManager) onYueBaoDepositResp(ctx client.Context, msg *pb.YueBaoDepositResp) error {
	if msg.GetRet() != 0 {
		return db.GetGameError(msg.GetRet())
	}

	return nil
}
