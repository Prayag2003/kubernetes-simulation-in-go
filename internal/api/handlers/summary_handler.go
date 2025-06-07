package handlers

import (
	"net/http"

	"github.com/Prayag2003/kubernetes-simulation/internal/analytics"
	"github.com/gin-gonic/gin"
)

func RegisterSummaryRoutes(r *gin.Engine) {
	r.GET("/api/summary", func(c *gin.Context) {
		c.JSON(http.StatusOK, analytics.Summary())
	})
}
