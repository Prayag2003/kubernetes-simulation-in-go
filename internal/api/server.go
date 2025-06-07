package api

import (
	"net/http"

	"github.com/Prayag2003/kubernetes-simulation/internal/api/handlers"
	hpa "github.com/Prayag2003/kubernetes-simulation/internal/autoscaler"
	"github.com/Prayag2003/kubernetes-simulation/internal/kubeapi"
	"github.com/gin-gonic/gin"
)

func StartServer(kube *kubeapi.KubeAPI, hpaConfig hpa.HPAConfig) {
	router := gin.Default()

	hpa.StartHPA(kube, hpaConfig)

	handlers.RegisterPodRoutes(router, kube)
	handlers.RegisterNodeRoutes(router, kube)
	handlers.RegisterHPARoutes(router)
	handlers.RegisterLogRoutes(router)
	handlers.RegisterSummaryRoutes(router)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.Run(":8080")
}
