package main

import (
	"flag"
	"time"

	hpa "github.com/Prayag2003/kubernetes-simulation/internal/autoscaler"
	"github.com/Prayag2003/kubernetes-simulation/internal/config"
	"github.com/Prayag2003/kubernetes-simulation/internal/etcdstore"
	"github.com/Prayag2003/kubernetes-simulation/internal/kubeapi"
	"github.com/Prayag2003/kubernetes-simulation/internal/models"
	"github.com/Prayag2003/kubernetes-simulation/internal/simulator"
	"github.com/Prayag2003/kubernetes-simulation/utils"
)

func main() {
	hpaConfigPath := flag.String("hpa-config", "", "Path to HPA config YAML file")
	flag.Parse()

	kube := kubeapi.NewKubeAPI()
	utils.LogInfo("Main", "Starting workload simulation...")
	simulator.StartWorkloadSim(kube)

	// Load and apply HPA config
	if *hpaConfigPath != "" {
		cfg, err := config.LoadHPAConfigFromFile(*hpaConfigPath)
		if err != nil {
			utils.LogError("Main", "Failed to load HPA config: "+err.Error())
		} else if cfg.Enabled {
			utils.LogInfo("Main", "HPA config loaded. Starting autoscaler...")
			hpa.StartHPA(kube, hpa.HPAConfig{
				TargetCPU: cfg.TargetCPU,
				MinPods:   cfg.MinPods,
				MaxPods:   cfg.MaxPods,
				Interval:  time.Duration(cfg.IntervalSeconds) * time.Second,
			})
		} else {
			utils.LogWarn("Main", "HPA config disabled.")
		}
	}

	// etcd test pod
	dummyPod := models.Pod{ID: "demo-pod-1", Name: "hello-world"}
	_ = etcdstore.GetStore().Set("/pods/demo-pod-1", dummyPod)

	time.Sleep(30 * time.Second)

	utils.LogWarn("Main", "Cleaning up all pods...")
	for _, pod := range kube.ListPods() {
		kube.DeletePod(pod.ID)
		etcdstore.GetStore().Delete("/pods/" + pod.ID)
	}

	utils.LogSuccess("Main", "Simulation completed.")
}
