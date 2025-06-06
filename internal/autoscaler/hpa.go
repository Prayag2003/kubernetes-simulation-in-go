package hpa

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/Prayag2003/kubernetes-simulation/internal/kubeapi"
	"github.com/Prayag2003/kubernetes-simulation/utils"
)

type HPAConfig struct {
	TargetCPU int
	MinPods   int
	MaxPods   int
	Interval  time.Duration
}

func StartHPA(kube *kubeapi.KubeAPI, config HPAConfig) {
	go func() {
		for {
			time.Sleep(config.Interval)

			pods := kube.ListPods()
			numPods := len(pods)

			if numPods == 0 {
				utils.LogWarn("HPA", "No pods running, creating one...")
				kube.CreatePod("auto-pod")
				continue
			}

			// Simulate CPU load (for now random)
			totalCPU := 0
			for _, pod := range pods {
				cpu := rand.Intn(100)
				utils.LogInfo("HPA", "Pod "+pod.ID+" CPU usage: "+strconv.Itoa(cpu)+"%")
				utils.LogInfo("HPA", "Pod "+pod.ID+" CPU usage: "+strconv.Itoa(cpu)+"%")
				totalCPU += cpu
			}

			avgCPU := totalCPU / numPods
			utils.LogInfo("HPA", "Average CPU = "+strconv.Itoa(avgCPU)+"%")
			utils.LogInfo("HPA", "Average CPU = "+strconv.Itoa(avgCPU)+"%")

			if avgCPU > config.TargetCPU && numPods < config.MaxPods {
				utils.LogSuccess("HPA", "Scaling up... adding a pod")
				kube.CreatePod("auto-pod")
			} else if avgCPU < config.TargetCPU/2 && numPods > config.MinPods {
				utils.LogWarn("HPA", "Scaling down... removing a pod")
				kube.DeletePod(pods[0].ID)
			} else {
				utils.LogInfo("HPA", "No scaling needed")
			}
		}
	}()
}
