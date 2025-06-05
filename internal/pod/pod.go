package pod

import (
	"fmt"
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
			log := fmt.Sprintf("[Pod %s] running at %v\n", pod.ID, time.Now())
			pod.Logs = append(pod.Logs, log)

		case <-stopChan:
			ticker.Stop()
			pod.Status = models.Succeeded
			fmt.Printf("[Pod %s] stopped.\n", pod.ID)
			return
		}
	}
}
