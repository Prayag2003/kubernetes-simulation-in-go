package main

import (
	"time"

	"github.com/Prayag2003/kubernetes-simulation/internal/kubeapi"
	"github.com/Prayag2003/kubernetes-simulation/internal/simulator"
	"github.com/Prayag2003/kubernetes-simulation/utils"
)

func main() {
	kube := kubeapi.NewKubeAPI()
	utils.LogInfo("Main", "Starting workload simulation...")
	simulator.StartWorkloadSim(kube)

	time.Sleep(30 * time.Second)

	utils.LogWarn("Main", "Cleaning up all pods...")
	for id := range kube.Pods {
		kube.DeletePod(id)
	}

	utils.LogSuccess("Main", "Simulation completed.")
}
