package handlers

import (
	"net/http"

	"github.com/Prayag2003/kubernetes-simulation/internal/kubeapi"
	"github.com/gin-gonic/gin"
)

func RegisterNodeRoutes(r *gin.Engine, kube *kubeapi.KubeAPI) {
	r.GET("/api/nodes", func(c *gin.Context) {
		c.JSON(http.StatusOK, kube.ListPods())
	})
}
