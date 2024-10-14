package assist

import (
	"context"

	"xddq/pkg/game/client"
	"xddq/pkg/game/protocol/pb"
	"xddq/pkg/utils"

	"go.uber.org/zap"
)

type DestinyManager struct {
	a *Assistant
}

func NewDestinyManager(a *Assistant) *DestinyManager {
	d := &DestinyManager{
		a: a,
	}

	if a.Config().Destiny {
		a.addRunner(d)

		a.Client().OnDestinyData(d.onDestinyData)
	}

	return d
}

func (d *DestinyManager) Run(ctx context.Context) error {
	d.a.addJob("destiny", "@every 10m", d.destinyTravel)

	return nil
}

func (d *DestinyManager) Stop() error {
	return nil
}

func (d *DestinyManager) onDestinyData(ctx client.Context, msg *pb.DestinyData) error {
	if msg.GetPlayerDestinyDataMsg().GetPower() > 0 {
		d.a.L().Info("destiny travel", zap.Int32("power", msg.GetPlayerDestinyDataMsg().GetPower()))
		d.destinyTravel(ctx)
	}

	return nil
}

func (d *DestinyManager) destinyTravel(ctx context.Context) error {
	return d.a.Client().SendReqDestinyTravel(ctx, &pb.ReqDestinyTravel{
		IsOneKey: utils.Ptr(true),
	})
}
