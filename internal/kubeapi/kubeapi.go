package kubeapi

import (
	"fmt"
	"sync"

	"github.com/Prayag2003/kubernetes-simulation/internal/models"
	"github.com/Prayag2003/kubernetes-simulation/internal/pod"
	"github.com/Prayag2003/kubernetes-simulation/utils"
	"github.com/google/uuid"
)

type KubeAPI struct {
	mu        sync.Mutex
	Pods      map[string]*models.Pod   // ID ==> Pod
	StopChans map[string]chan struct{} // ID ==> stop signal
}

func NewKubeAPI() *KubeAPI {
	return &KubeAPI{
		Pods:      make(map[string]*models.Pod),
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

	stopChan := make(chan struct{})
	k.Pods[id] = p
	k.StopChans[id] = stopChan

	go pod.StartPod(p, stopChan)
	utils.LogSuccess("KubeAPI", fmt.Sprintf("Created Pod: ID=%s, Name=%s", id, name))

	return id
}

func (k *KubeAPI) DeletePod(id string) {
	k.mu.Lock()
	defer k.mu.Unlock()

	stopChan, exists := k.StopChans[id]
	if !exists {
		utils.LogError("KubeAPI", fmt.Sprintf("Pod ID=%s not found", id))
		return
	}

	close(stopChan)
	delete(k.Pods, id)
	delete(k.StopChans, id)

	utils.LogWarn("KubeAPI", fmt.Sprintf("Deleted Pod ID=%s", id))
}
