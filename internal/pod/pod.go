package pod

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/Prayag2003/kubernetes-simulation/internal/models"
)

func StartPod(pod *models.Pod, stopChan chan struct{}) {
	fmt.Printf("[Pod %s] starting ....\n", pod.ID)
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

			log := fmt.Sprintf("[Pod %s] CPU: %.2fm, MEM: %.2fMB", pod.ID, cpu, mem)
			pod.Logs = append(pod.Logs, log)

		case <-stopChan:
			ticker.Stop()
			pod.Status = models.Succeeded
			fmt.Printf("[Pod %s] stopped.\n", pod.ID)
			return
		}
	}
}
