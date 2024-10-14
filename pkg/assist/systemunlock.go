package assist

import (
	"math/big"

	"xddq/pkg/game/client"
	"xddq/pkg/game/protocol/pb"
)

type SystemUnlockManager struct {
	flags *big.Int
}

func NewSystemUnlockManager(a *Assistant) *SystemUnlockManager {
	s := &SystemUnlockManager{
		flags: big.NewInt(0),
	}
	a.Client().OnSystemUnlockSync(s.onSyncSystemUnlock)

	return s
}

func (s *SystemUnlockManager) onSyncSystemUnlock(ctx client.Context, msg *pb.SystemUnlockSync) error {
	s.flags.SetString(msg.GetUnlockInfo(), 10)
	return nil
}

func (s *SystemUnlockManager) IsSystemUnlocked(systemId int) bool {
	return s.flags.Bit(systemId) == 1
}
