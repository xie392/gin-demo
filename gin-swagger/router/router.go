package router

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the routing for the API
func SetupRouter(r *gin.Engine) {
	r.GET("/ping", ping)
}

// ping example
// @Summary ping example
// @Schemes
// @Description Do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} string "pong"
// @Router /ping [get]
func ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
