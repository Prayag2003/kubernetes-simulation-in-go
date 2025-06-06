package etcdstore

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/Prayag2003/kubernetes-simulation/internal/analytics"
)

type EtcdStore struct {
	sync.RWMutex
	data map[string][]byte
}

var store *EtcdStore
var once sync.Once

func GetStore() *EtcdStore {
	once.Do(func() {
		store = &EtcdStore{
			data: make(map[string][]byte),
		}
		analytics.Log("EtcdStore", "info", "Init", "Initialized in-memory etcd-style store.")
	})
	return store
}

func (e *EtcdStore) Set(key string, value any) error {
	e.Lock()
	defer e.Unlock()

	bytes, err := json.Marshal(value)
	if err != nil {
		analytics.Log("EtcdStore", "error", "MarshalFailed", fmt.Sprintf("key=%s: %v", key, err))
		return fmt.Errorf("marshal failed: %w", err)
	}

	e.data[key] = bytes
	analytics.Log("EtcdStore", "success", "SetKey", fmt.Sprintf("key=%s", key))
	return nil
}

func (e *EtcdStore) Get(key string, out any) error {
	e.RLock()
	defer e.RUnlock()

	val, exists := e.data[key]
	if !exists {
		analytics.Log("EtcdStore", "warn", "KeyNotFound", fmt.Sprintf("key=%s", key))
		return fmt.Errorf("key not found: %s", key)
	}

	if err := json.Unmarshal(val, out); err != nil {
		analytics.Log("EtcdStore", "error", "UnmarshalFailed", fmt.Sprintf("key=%s: %v", key, err))
		return fmt.Errorf("unmarshal failed: %w", err)
	}

	analytics.Log("EtcdStore", "info", "GetKey", fmt.Sprintf("key=%s", key))
	return nil
}

func (e *EtcdStore) Delete(key string) {
	e.Lock()
	defer e.Unlock()

	delete(e.data, key)
	analytics.Log("EtcdStore", "warn", "DeletedKey", fmt.Sprintf("key=%s", key))
}

func (e *EtcdStore) List(prefix string) map[string][]byte {
	e.RLock()
	defer e.RUnlock()

	results := make(map[string][]byte)
	count := 0
	for k, v := range e.data {
		if len(k) >= len(prefix) && k[:len(prefix)] == prefix {
			results[k] = v
			count++
		}
	}
	analytics.Log("EtcdStore", "info", "ListKeys", fmt.Sprintf("Listed %d keys with prefix=%s", count, prefix))
	return results
}
