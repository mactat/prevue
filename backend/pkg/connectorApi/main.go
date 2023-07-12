package connectorapi

import (
	"database/sql"
	db "prevue/pkg/db"
	types "prevue/pkg/types"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var baseRoute = "/api/connector"

func Routes(router *gin.Engine) {
	router.GET(baseRoute, helloHandler)
	router.POST(baseRoute+"/metrics", metricsHandler)
}

func helloHandler(c *gin.Context) {
	c.String(200, "Hi there!")
}

func metricsHandler(c *gin.Context) {
	var data types.ConnectorData
	c.BindJSON(&data)
	database := c.MustGet("database").(*sql.DB)
	err := db.Insert(database, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into the database"})
		log.Println("Failed to insert data into the database")
	}
	c.JSON(http.StatusOK, data)
	
}
