package assist

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"time"
	"xddq/pkg/game/client"
	"xddq/pkg/game/protocol/pb"
	"xddq/pkg/utils"
)

type PlayerDataStorageManager struct {
	a *Assistant

	mu sync.RWMutex
	kv map[string]string
}

func NewPlayerDataStorageManager(a *Assistant) *PlayerDataStorageManager {
	ds := &PlayerDataStorageManager{
		a:  a,
		kv: make(map[string]string),
	}

	a.Client().OnAllClientDataSync(ds.onAllClientDataSync)

	return ds
}

// onAllClientDataSync
func (ds *PlayerDataStorageManager) onAllClientDataSync(ctx client.Context, msg *pb.AllClientDataSync) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	for _, data := range msg.DataList {
		ds.kv[data.GetKey()] = data.GetValue()
	}

	return nil
}

func (ds *PlayerDataStorageManager) Set(ctx context.Context, key string, val any) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	bytes, err := json.Marshal(val)
	if err != nil {
		return err
	}

	ds.kv[key] = string(bytes)

	return ds.save(ctx, key, string(bytes))
}

func (ds *PlayerDataStorageManager) Delete(ctx context.Context, key string) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	delete(ds.kv, key)

	return ds.save(ctx, key, "\x7f")
}

func (ds *PlayerDataStorageManager) Get(key string, val any) error {
	ds.mu.RLock()
	value, ok := ds.kv[key]
	ds.mu.RUnlock()

	if !ok {
		return errors.New("not found")
	}

	return json.Unmarshal([]byte(value), val)
}

func (ds *PlayerDataStorageManager) GetInt64(key string) int64 {
	var value int64
	if err := ds.Get(key, &value); err != nil {
		return 0
	}
	return value
}

func (ds *PlayerDataStorageManager) Cache() *CacheStorage {
	return &CacheStorage{ds: ds}
}

func (ds *PlayerDataStorageManager) GetBool(key string) bool {
	var value bool
	if err := ds.Get(key, &value); err != nil {
		return false
	}
	return value
}

func (ds *PlayerDataStorageManager) save(ctx context.Context, key string, val string) error {
	return ds.a.Client().SendSaveToServiceReq(ctx, &pb.SaveToServiceReq{DataList: []*pb.TempKeyValueData{
		{
			Key:   &key,
			Value: utils.Ptr(val),
		},
	}})
}

type CacheStorage struct {
	ds *PlayerDataStorageManager
}

type KV struct {
	Val       json.RawMessage `json:"v"`
	ExpiredAt int64           `json:"e"`
}

func (cs *CacheStorage) SetWithTimeout(ctx context.Context, key string, val any, timeout time.Duration) error {
	return cs.Set(ctx, key, val, time.Now().Add(timeout).UnixMilli())
}

func (cs *CacheStorage) Set(ctx context.Context, key string, val any, expiredAt ...int64) error {
	exp := TomorrowMilli()
	if len(expiredAt) > 0 {
		exp = expiredAt[0]
	}

	cs.ds.mu.Lock()
	defer cs.ds.mu.Unlock()

	data, err := json.Marshal(val)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(KV{
		Val:       data,
		ExpiredAt: exp,
	})
	if err != nil {
		return err
	}

	cs.ds.kv[key] = string(bytes)

	return cs.ds.save(ctx, key, string(bytes))
}

func (cs *CacheStorage) Get(key string, val any) error {
	cs.ds.mu.RLock()
	value, ok := cs.ds.kv[key]
	cs.ds.mu.RUnlock()

	if !ok {
		return errors.New("not found")
	}

	kv := KV{}
	if err := json.Unmarshal([]byte(value), &kv); err != nil {
		return err
	}

	if kv.ExpiredAt < time.Now().UnixMilli() {
		return errors.New("expired")
	}

	return json.Unmarshal(kv.Val, val)
}

func (cs *CacheStorage) GetInt64(key string) int64 {
	var value int64
	if err := cs.Get(key, &value); err != nil {
		return 0
	}
	return value
}

func (cs *CacheStorage) GetBool(key string) bool {
	var value bool
	if err := cs.Get(key, &value); err != nil {
		return false
	}
	return value
}

func (cs *CacheStorage) Condition(ctx context.Context, key string, fn func() error) error {
	if cs.GetBool(key) {
		return nil
	}

	if err := fn(); err != nil {
		return err
	}

	cs.Set(ctx, key, true)

	return nil
}
