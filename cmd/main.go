package main

import (
	"flag"
	"time"

	"github.com/Prayag2003/kubernetes-simulation/internal/analytics"
	"github.com/Prayag2003/kubernetes-simulation/internal/api"
	hpa "github.com/Prayag2003/kubernetes-simulation/internal/autoscaler"
	"github.com/Prayag2003/kubernetes-simulation/internal/config"
	"github.com/Prayag2003/kubernetes-simulation/internal/kubeapi"
)

func main() {
	hpaConfigPath := flag.String("hpa-config", "", "Path to optional HPA config YAML file")
	flag.Parse()

	kube := kubeapi.NewKubeAPI()
	var hpaConfig hpa.HPAConfig

	if *hpaConfigPath != "" {
		cfg, err := config.LoadHPAConfigFromFile(*hpaConfigPath)
		if err != nil {
			analytics.Log("Main", "error", "HPAConfigLoad", "Failed to load HPA config: "+err.Error())
		} else if cfg.Enabled {
			analytics.Log("Main", "info", "HPAEnabled", "Loaded HPA config from "+*hpaConfigPath)
			hpaConfig = hpa.HPAConfig{
				TargetCPU: cfg.TargetCPU,
				MinPods:   cfg.MinPods,
				MaxPods:   cfg.MaxPods,
				Interval:  time.Duration(cfg.IntervalSeconds) * time.Second,
			}
		} else {
			analytics.Log("Main", "warn", "HPADisabled", "HPA is disabled in the config.")
		}
	} else {
		analytics.Log("Main", "warn", "HPAMockConfig", "No HPA config file provided. Using mock config.")
		hpaConfig = hpa.HPAConfig{
			TargetCPU: 50,
			MinPods:   1,
			MaxPods:   5,
			Interval:  10 * time.Second,
		}
	}

	api.StartServer(kube, hpaConfig)
}
