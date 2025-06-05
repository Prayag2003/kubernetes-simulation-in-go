package kubeapi

import (
	"fmt"
	"sync"

	"github.com/Prayag2003/kubernetes-simulation/internal/models"
	"github.com/Prayag2003/kubernetes-simulation/internal/pod"
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
	fmt.Sprintf("[KubeAPI] has created Pod with ID = %s Name = %s\n", p.ID, p.Name)

	return id
}
