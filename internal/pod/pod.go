package pod

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/Prayag2003/kubernetes-simulation/internal/analytics"
	"github.com/Prayag2003/kubernetes-simulation/internal/models"
)

func StartPod(pod *models.Pod, stopChan chan struct{}) {
	analytics.Log("Pod", "info", "PodStart", fmt.Sprintf("Pod %s starting ...", pod.ID))
	pod.Status = models.Running
	pod.StartTime = time.Now()

	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			cpu := rand.Float64()*100 + 100 // 100-200 millicores
			mem := rand.Float64()*50 + 50   // 50-100 MB
			pod.Resources = models.ResourceUsage{
				CPU:    cpu,
				Memory: mem,
			}

			analytics.Log("Pod", "info", "PodMetrics", fmt.Sprintf("Pod %s CPU: %.2fm, MEM: %.2fMB", pod.ID, cpu, mem))
			log := fmt.Sprintf("[Pod %s] CPU: %.2fm, MEM: %.2fMB", pod.ID, cpu, mem)
			pod.Logs = append(pod.Logs, log)

		case <-stopChan:
			ticker.Stop()
			pod.Status = models.Succeeded
			analytics.Log("Pod", "warn", "PodStop", fmt.Sprintf("Pod %s stopped.", pod.ID))
			return
		}
	}
}
