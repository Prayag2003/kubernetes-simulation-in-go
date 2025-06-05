package models

import "time"

type PodStatus string

const (
	Pending   PodStatus = "Pending"
	Running   PodStatus = "Running"
	Failed    PodStatus = "Failed"
	Succeeded PodStatus = "Succeeded"
)

type Pod struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Status        PodStatus `json:"status"`
	StartTime     time.Time `json:"start_time"`
	Logs          []string
	ResourceUsage ResourceUsage
}

type ResourceUsage struct {
	CPU    float64
	Memory float64
}
