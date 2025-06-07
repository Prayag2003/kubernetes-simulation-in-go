package handlers

import (
	"net/http"
	"time"

	hpa "github.com/Prayag2003/kubernetes-simulation/internal/autoscaler"
	"github.com/gin-gonic/gin"
)

type HPAConfigInput struct {
	TargetCPU int    `json:"target_cpu" binding:"required"`
	MinPods   int    `json:"min_pods" binding:"required"`
	MaxPods   int    `json:"max_pods" binding:"required"`
	Interval  string `json:"interval" binding:"required"`
}

func RegisterHPARoutes(r *gin.Engine) {
	r.GET("/api/hpa/config", func(c *gin.Context) {
		c.JSON(http.StatusOK, hpa.GetHPAConfig())
	})

	r.POST("/api/hpa/config", func(c *gin.Context) {
		var input HPAConfigInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		intervalDuration, err := time.ParseDuration(input.Interval)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval format"})
			return
		}

		newConfig := hpa.HPAConfig{
			TargetCPU: input.TargetCPU,
			MinPods:   input.MinPods,
			MaxPods:   input.MaxPods,
			Interval:  intervalDuration,
		}

		hpa.UpdateHPAConfig(newConfig)
		c.JSON(http.StatusOK, gin.H{"message": "HPA config updated", "config": newConfig})
	})
}
