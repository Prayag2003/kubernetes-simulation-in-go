package pod

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/Prayag2003/kubernetes-simulation/internal/models"
	"github.com/Prayag2003/kubernetes-simulation/utils"
)

func StartPod(pod *models.Pod, stopChan chan struct{}) {
	utils.LogInfo("Pod", fmt.Sprintf("Pod %s starting ...", pod.ID))
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

			utils.LogInfo("Pod", fmt.Sprintf("Pod %s CPU: %.2fm, MEM: %.2fMB", pod.ID, cpu, mem))
			log := fmt.Sprintf("[Pod %s] CPU: %.2fm, MEM: %.2fMB", pod.ID, cpu, mem)
			pod.Logs = append(pod.Logs, log)

		case <-stopChan:
			ticker.Stop()
			pod.Status = models.Succeeded
			utils.LogWarn("Pod", fmt.Sprintf("Pod %s stopped.", pod.ID))
			return
		}
	}
}
