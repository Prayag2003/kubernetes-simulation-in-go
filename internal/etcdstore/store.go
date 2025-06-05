package etcdstore

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/Prayag2003/kubernetes-simulation/utils"
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
		utils.LogInfo("EtcdStore", "Initialized in-memory etcd-style store.")
	})
	return store
}

func (e *EtcdStore) Set(key string, value any) error {
	e.Lock()
	defer e.Unlock()

	bytes, err := json.Marshal(value)
	if err != nil {
		utils.LogError("EtcdStore", fmt.Sprintf("Marshal failed for key=%s: %v", key, err))
		return fmt.Errorf("marshal failed: %w", err)
	}

	e.data[key] = bytes
	utils.LogSuccess("EtcdStore", fmt.Sprintf("Set key=%s", key))
	return nil
}

func (e *EtcdStore) Get(key string, out any) error {
	e.RLock()
	defer e.RUnlock()

	val, exists := e.data[key]
	if !exists {
		utils.LogWarn("EtcdStore", fmt.Sprintf("Get failed: key not found (%s)", key))
		return fmt.Errorf("key not found: %s", key)
	}

	if err := json.Unmarshal(val, out); err != nil {
		utils.LogError("EtcdStore", fmt.Sprintf("Unmarshal failed for key=%s: %v", key, err))
		return fmt.Errorf("unmarshal failed: %w", err)
	}

	utils.LogInfo("EtcdStore", fmt.Sprintf("Get success: key=%s", key))
	return nil
}

func (e *EtcdStore) Delete(key string) {
	e.Lock()
	defer e.Unlock()

	delete(e.data, key)
	utils.LogWarn("EtcdStore", fmt.Sprintf("Deleted key=%s", key))
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
	utils.LogInfo("EtcdStore", fmt.Sprintf("Listed %d keys with prefix=%s", count, prefix))
	return results
}
