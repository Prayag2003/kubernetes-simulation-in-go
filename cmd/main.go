package main

import (
	"fmt"
	"time"

	"github.com/Prayag2003/kubernetes-simulation/internal/kubeapi"
)

func main() {
	kube := kubeapi.NewKubeAPI()

	var podIDs []string
	for i := 1; i <= 5; i++ {
		name := fmt.Sprintf("nginx-%d", i)
		id := kube.CreatePod(name)
		podIDs = append(podIDs, id)
	}

	time.Sleep(6 * time.Second)

	for _, id := range podIDs {
		kube.DeletePod(id)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("[Main] All pods cleaned up. Simulation successful.")
}
