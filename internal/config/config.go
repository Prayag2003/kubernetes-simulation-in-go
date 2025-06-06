package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type HPAFileConfig struct {
	Enabled         bool `yaml:"enabled"`
	TargetCPU       int  `yaml:"targetCPU"`
	MinPods         int  `yaml:"minPods"`
	MaxPods         int  `yaml:"maxPods"`
	IntervalSeconds int  `yaml:"intervalSeconds"`
}

func LoadHPAConfigFromFile(path string) (*HPAFileConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg HPAFileConfig
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
