package scheduler

import (
	"fmt"
	"time"

	"github.com/Prayag2003/kubernetes-simulation/internal/analytics"
	"github.com/Prayag2003/kubernetes-simulation/internal/kubeapi"
	"github.com/Prayag2003/kubernetes-simulation/internal/models"
	node "github.com/Prayag2003/kubernetes-simulation/internal/nodes"
)

func StartScheduler(kube *kubeapi.KubeAPI) {
	go func() {
		for {
			time.Sleep(2 * time.Second)

			for _, pod := range kube.ListPods() {
				if pod.Status != models.Pending {
					continue
				}

				scheduled := false

				for _, n := range node.Nodes {
					if n.UsedCPU+pod.Resources.CPU <= n.CPU && n.UsedMem+pod.Resources.Memory <= n.Memory {
						n.UsedCPU += pod.Resources.CPU
						n.UsedMem += pod.Resources.Memory

						analytics.Log("Scheduler", "Info", fmt.Sprintf("Scheduled %s to %s", pod.ID, n.Name), "")
						scheduled = true
						break
					}
				}

				if !scheduled {
					analytics.Log("Scheduler", "Warn", fmt.Sprintf("No node has enough resources for %s", pod.ID), "")
				}
			}
		}
	}()
}
