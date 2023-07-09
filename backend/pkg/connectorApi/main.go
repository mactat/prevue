package connectorapi

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var baseRoute = "/api/connector"

type MetricsData struct {
	Accuracy float64 `json:"accuracy"`
	MSE      float64 `json:"mse"`
}

type connectorData struct {
	Uid           string      `json:"uid"`
	ProjectName   string      `json:"projectName"`
	ConnectorName string      `json:"connectorName"`
	Metrics       MetricsData `json:"metricsData"`
}

func Routes(router *gin.Engine) {
	router.GET(baseRoute, helloHandler)
	router.POST(baseRoute+"/metrics", metricsHandler)
}

func helloHandler(c *gin.Context) {
	c.String(200, "Hi there!")
}

func metricsHandler(c *gin.Context) {
	var data connectorData
	c.BindJSON(&data)
	fmt.Println(data)
	c.JSON(200, data)
}
