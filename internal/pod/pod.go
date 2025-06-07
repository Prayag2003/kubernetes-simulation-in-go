package pod

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/Prayag2003/kubernetes-simulation/internal/analytics"
	"github.com/Prayag2003/kubernetes-simulation/internal/etcdstore"
	"github.com/Prayag2003/kubernetes-simulation/internal/models"
)

func StartPod(pod *models.Pod, stopChan chan struct{}) {
	analytics.Log("Pod", "info", "PodStart", fmt.Sprintf("Pod %s starting ...", pod.ID))

	pod.Status = models.Running
	pod.StartTime = time.Now()

	_ = etcdstore.GetStore().Set("/pods/"+pod.ID, pod)

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

			analytics.Log("Pod", "info", "PodMetrics", log)

			_ = etcdstore.GetStore().Set("/pods/"+pod.ID, pod)

		case <-stopChan:
			ticker.Stop()
			pod.Status = models.Succeeded
			analytics.Log("Pod", "warn", "PodStop", fmt.Sprintf("Pod %s stopped.", pod.ID))

			_ = etcdstore.GetStore().Set("/pods/"+pod.ID, pod)
			return
		}
	}
}
