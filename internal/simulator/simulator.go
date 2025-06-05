package simulator

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Prayag2003/kubernetes-simulation/internal/kubeapi"
)

func StartWorkloadSim(kube *kubeapi.KubeAPI) {
	go func() {
		podCount := 0
		for {
			// Wait 1–3 seconds between pod creations
			delay := time.Duration(rand.Intn(3)+1) * time.Second
			time.Sleep(delay)

			name := fmt.Sprintf("workload-pod-%d", podCount)
			kube.CreatePod(name)
			podCount++

			// Randomly delete a pod occasionally
			if len(kube.Pods) > 3 && rand.Float64() < 0.3 {
				for id := range kube.Pods {
					kube.DeletePod(id)
					break
				}
			}
		}
	}()
}
