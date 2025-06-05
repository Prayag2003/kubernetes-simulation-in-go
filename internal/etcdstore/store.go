package etcdstore

import (
	"encoding/json"
	"fmt"
	"sync"
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
	})
	return store
}

func (e *EtcdStore) Set(key string, value any) error {
	e.Lock()
	defer e.Unlock()

	bytes, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("Marshall failed: %w", err)
	}

	e.data[key] = bytes
	return nil
}

func (e *EtcdStore) Get(key string, out any) error {
	e.Lock()
	defer e.Unlock()

	val, exists := e.data[key]
	if !exists {
		return fmt.Errorf("Key not found: %s", key)
	}

	if err := json.Unmarshal(val, out); err != nil {
		return fmt.Errorf("unmarshall failed: %w", err)
	}

	return nil
}
