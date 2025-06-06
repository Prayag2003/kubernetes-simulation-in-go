package kubeapi

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/Prayag2003/kubernetes-simulation/internal/etcdstore"
	"github.com/Prayag2003/kubernetes-simulation/internal/models"
	"github.com/Prayag2003/kubernetes-simulation/internal/pod"
	"github.com/Prayag2003/kubernetes-simulation/utils"
	"github.com/google/uuid"
)

type KubeAPI struct {
	mu        sync.Mutex
	StopChans map[string]chan struct{} // ID ==> stop signal
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

	// Save to etcd
	key := fmt.Sprintf("/pods/%s", id)
	if err := etcdstore.GetStore().Set(key, p); err != nil {
		utils.LogError("KubeAPI", fmt.Sprintf("Failed to persist pod %s: %v", id, err))
		return ""
	}

	// Start pod process
	stopChan := make(chan struct{})
	k.StopChans[id] = stopChan
	go pod.StartPod(p, stopChan)

	utils.LogSuccess("KubeAPI", fmt.Sprintf("Created Pod: ID=%s, Name=%s", id, name))
	return id
}

func (k *KubeAPI) DeletePod(id string) {
	k.mu.Lock()
	defer k.mu.Unlock()

	stopChan, exists := k.StopChans[id]
	if exists {
		close(stopChan)
		delete(k.StopChans, id)
	}

	key := fmt.Sprintf("/pods/%s", id)
	etcdstore.GetStore().Delete(key)
	utils.LogWarn("KubeAPI", fmt.Sprintf("Deleted Pod ID=%s", id))
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
