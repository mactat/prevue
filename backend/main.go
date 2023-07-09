package main

import (
	connectorapi "prevue/pkg/connectorApi"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	connectorapi.Routes(router)
	router.Run("localhost:8080")
}
