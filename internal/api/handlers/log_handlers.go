package handlers

import (
	"net/http"

	"github.com/Prayag2003/kubernetes-simulation/internal/analytics"
	"github.com/gin-gonic/gin"
)

func RegisterLogRoutes(r *gin.Engine) {
	r.GET("/api/logs", func(c *gin.Context) {
		c.JSON(http.StatusOK, analytics.GetLogs())
	})
}
