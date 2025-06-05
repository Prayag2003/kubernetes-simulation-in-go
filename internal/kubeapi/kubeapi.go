package kubeapi

import (
	"sync"

	"github.com/Prayag2003/kubernetes-simulation/internal/models"
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
