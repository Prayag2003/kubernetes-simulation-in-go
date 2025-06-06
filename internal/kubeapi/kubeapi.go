package kubeapi

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/Prayag2003/kubernetes-simulation/internal/analytics"
	"github.com/Prayag2003/kubernetes-simulation/internal/etcdstore"
	"github.com/Prayag2003/kubernetes-simulation/internal/models"
	"github.com/Prayag2003/kubernetes-simulation/internal/pod"
	"github.com/google/uuid"
)

type KubeAPI struct {
	mu        sync.Mutex
	StopChans map[string]chan struct{}
}

func NewKubeAPI() *KubeAPI {
	return &KubeAPI{
		StopChans: make(map[string]chan struct{}),
	}
}

func (k *KubeAPI) CreatePod(name string) string {
	k.mu.Lock()
	defer k.mu.Unlock()

	id := uuid.New().String()
	p := &models.Pod{
		ID:     id,
		Name:   name,
		Status: models.Pending,
	}

	key := fmt.Sprintf("/pods/%s", id)
	if err := etcdstore.GetStore().Set(key, p); err != nil {
		analytics.Log("KubeAPI", "error", "PersistPodFailed", fmt.Sprintf("podID=%s: %v", id, err))
		return ""
	}

	stopChan := make(chan struct{})
	k.StopChans[id] = stopChan
	go pod.StartPod(p, stopChan)

	analytics.Log("KubeAPI", "success", "CreatedPod", fmt.Sprintf("ID=%s, Name=%s", id, name))
	return id
}

func (k *KubeAPI) DeletePod(id string) {
	k.mu.Lock()
	defer k.mu.Unlock()

	if stopChan, exists := k.StopChans[id]; exists {
		close(stopChan)
		delete(k.StopChans, id)
	}

	key := fmt.Sprintf("/pods/%s", id)
	etcdstore.GetStore().Delete(key)
	analytics.Log("KubeAPI", "warn", "DeletedPod", fmt.Sprintf("ID=%s", id))
}

func (k *KubeAPI) ListPods() []*models.Pod {
	raw := etcdstore.GetStore().List("/pods/")
	result := make([]*models.Pod, 0, len(raw))

	for _, val := range raw {
		var p models.Pod
		if err := json.Unmarshal(val, &p); err == nil {
			result = append(result, &p)
		}
	}
	return result
}
