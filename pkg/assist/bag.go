package assist

import (
	"strconv"
	"sync"

	"xddq/pkg/game/client"
	"xddq/pkg/game/db"
	"xddq/pkg/game/protocol/pb"
)

type BagItem struct {
	Name   string
	PropId int32
	Num    int64
}

type BagManager struct {
	a     *Assistant
	mu    sync.RWMutex
	items map[int32]*BagItem
	malls map[int32]int32
}

func NewBagManager(a *Assistant) *BagManager {
	b := &BagManager{
		a:     a,
		items: make(map[int32]*BagItem),
		malls: make(map[int32]int32),
	}

	a.Client().OnSyncBagMsg(b.onSyncBag)
	a.Client().OnMallBuyCountListMsg(b.onMallBuyCountListMsg)

	return b
}

func (b *BagManager) onSyncBag(ctx client.Context, msg *pb.SyncBagMsg) error {
	b.a.L().Info("同步背包数据")

	b.mu.Lock()
	defer b.mu.Unlock()

	for _, item := range msg.GetBagData() {
		num, _ := strconv.ParseInt(item.GetNum(), 10, 64)
		b.items[item.GetPropId()] = &BagItem{
			Name:   db.GetItemName(item.GetPropId()),
			PropId: item.GetPropId(),
			Num:    num,
		}
	}

	return nil
}

func (b *BagManager) onMallBuyCountListMsg(ctx client.Context, msg *pb.MallBuyCountListMsg) error {
	b.a.L().Info("同步商城购买数量列表")

	b.mu.Lock()
	defer b.mu.Unlock()

	for _, item := range msg.GetMallBuyCountList() {
		b.malls[item.GetMallId()] = item.GetCount()
	}

	return nil
}

func (b *BagManager) GetItem(propId int32) *BagItem {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.items[propId]
}

func (b *BagManager) GetItemByName(name string) *BagItem {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, item := range b.items {
		if item.Name == name {
			return item
		}
	}

	return nil
}

func (b *BagManager) GetMallCount(mallId int32) int32 {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.malls[mallId]
}
