package assist

import (
	"xddq/pkg/game/client"
	"xddq/pkg/game/db"
	"xddq/pkg/game/protocol/pb"

	"go.uber.org/zap"
)

type HelperManager struct {
	a *Assistant
}

func NewHelperManager(a *Assistant) *HelperManager {
	hm := &HelperManager{
		a: a,
	}

	a.Client().OnUseExchangeKeyReq(hm.onUseExchangeKeyResp)

	return hm
}

// onUseExchangeKeyResp
func (hm *HelperManager) onUseExchangeKeyResp(ctx client.Context, msg *pb.UseExchangeKeyResp) error {
	if msg.GetRet() != 0 {
		hm.a.L().Error("use exchange key resp", zap.Error(db.GetGameError(msg.GetRet())))
		return nil
	}

	hm.a.L().Info("use exchange key resp", zap.Any("reward", db.ParseRewards(msg.GetReward())))

	return nil
}
