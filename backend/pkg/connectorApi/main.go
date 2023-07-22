package connectorapi

import (
	"database/sql"
	"log"
	"net/http"
	db "prevue/pkg/db"
	types "prevue/pkg/types"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var baseRoute = "/api/connector"

func Routes(router *gin.Engine) {
	router.GET(baseRoute, helloHandler)
	router.POST(baseRoute+"/session", sessionHandler)
	router.POST(baseRoute+"/metrics", metricsHandler)
}

func helloHandler(c *gin.Context) {
	c.String(200, "Hi there!")
}

func sessionHandler(c *gin.Context) {
	var dataSession types.SessionData

	c.BindJSON(&dataSession)
	database := c.MustGet("database").(*sql.DB)
	modelId, err := db.SessionData(database, dataSession)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert session data into the database"})
		log.Println("Failed to insert session data into the database")
	}
	log.Println(modelId)
	c.JSON(http.StatusOK, modelId)

}

func metricsHandler(c *gin.Context) {
	var dataMetrics types.SessionMetrics
	c.BindJSON(&dataMetrics)
	database := c.MustGet("database").(*sql.DB)

	err := db.MetricsData(database, dataMetrics)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert metrics data into the database"})
		log.Println("Failed to insert metrics data into the database")
	}
	c.JSON(http.StatusOK, err)

}
