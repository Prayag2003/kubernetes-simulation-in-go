package hpa

import (
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/Prayag2003/kubernetes-simulation/internal/analytics"
	"github.com/Prayag2003/kubernetes-simulation/internal/kubeapi"
)

type HPAConfig struct {
	TargetCPU int
	MinPods   int
	MaxPods   int
	Interval  time.Duration
}

var (
	currentConfig HPAConfig
	configMu      = sync.RWMutex{}
)

func StartHPA(kube *kubeapi.KubeAPI, config HPAConfig) {
	currentConfig = config
	go func() {
		for {
			time.Sleep(config.Interval)

			pods := kube.ListPods()
			numPods := len(pods)

			if numPods == 0 {
				analytics.Log("HPA", "warn", "NoPods", "No pods running, creating one...")
				kube.CreatePod("auto-pod")
				continue
			}

			totalCPU := 0
			for _, pod := range pods {
				cpu := rand.Intn(100)
				analytics.Log("HPA", "info", "PodCPU", "Pod "+pod.ID+" CPU usage: "+strconv.Itoa(cpu)+"%")
				totalCPU += cpu
			}

			avgCPU := totalCPU / numPods
			analytics.Log("HPA", "info", "AvgCPU", "Average CPU = "+strconv.Itoa(avgCPU)+"%")

			if avgCPU > config.TargetCPU && numPods < config.MaxPods {
				analytics.Log("HPA", "success", "ScaleUp", "Scaling up... adding a pod")
				kube.CreatePod("auto-pod")
			} else if avgCPU < config.TargetCPU/2 && numPods > config.MinPods {
				analytics.Log("HPA", "warn", "ScaleDown", "Scaling down... removing a pod")
				kube.DeletePod(pods[0].ID)
			} else {
				analytics.Log("HPA", "info", "Stable", "No scaling needed")
			}
		}
	}()
}

func GetHPAConfig() HPAConfig {
	configMu.RLock()
	defer configMu.RUnlock()
	return currentConfig
}

func UpdateHPAConfig(config HPAConfig) {
	configMu.Lock()
	currentConfig = config
	configMu.Unlock()
}
