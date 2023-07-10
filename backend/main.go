package main

import (
	connectorapi "prevue/pkg/connectorApi"
	statusapi "prevue/pkg/statusApi"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	connectorapi.Routes(router)
	statusapi.Routes(router)
	router.Run("0.0.0.0:8080")
}
