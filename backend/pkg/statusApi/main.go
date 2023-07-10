package statusapi

import (
	"github.com/gin-gonic/gin"
)

var baseRoute = "/api/status"

func Routes(router *gin.Engine) {
	router.GET(baseRoute+"/health", healthHandler)
	router.GET(baseRoute+"/readiness", readinessHandler)
}

func healthHandler(c *gin.Context) {
	c.String(200, "Good!")
}

func readinessHandler(c *gin.Context) {
	c.String(200, "Good!")
}
