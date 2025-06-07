package handlers

import (
	"net/http"

	"github.com/Prayag2003/kubernetes-simulation/internal/kubeapi"
	"github.com/gin-gonic/gin"
)

func RegisterPodRoutes(r *gin.Engine, kube *kubeapi.KubeAPI) {
	r.GET("/api/pods", func(c *gin.Context) {
		c.JSON(http.StatusOK, kube.ListPods())
	})

	r.GET("/api/pods/:id", func(c *gin.Context) {
		id := c.Param("id")
		pod := kube.GetPod(id)
		if pod == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Pod not found"})
			return
		}
		c.JSON(http.StatusOK, pod)
	})

	r.POST("/api/pods", func(c *gin.Context) {
		var request struct {
			Name string `json:"name"`
		}
		if err := c.BindJSON(&request); err != nil || request.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pod name"})
			return
		}
		id := kube.CreatePod(request.Name)
		c.JSON(http.StatusCreated, gin.H{"message": "Pod created", "id": id})
	})

	r.DELETE("/api/pods/:id", func(c *gin.Context) {
		id := c.Param("id")
		kube.DeletePod(id)
		c.JSON(http.StatusOK, gin.H{"message": "Pod deleted"})
	})
}
