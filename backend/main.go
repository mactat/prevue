package main

import (
	"os"
	connectorapi "prevue/pkg/connectorApi"
	db "prevue/pkg/db"
	statusapi "prevue/pkg/statusApi"

	"github.com/gin-gonic/gin"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
)

func main() {
	router := gin.Default()

	database := db.Connect(dbname, user, password, host, port)
	defer db.Close(database)

	db.CreateTables(database)

	// Provide db variable to controllers
	router.Use(func(c *gin.Context) {
		c.Set("database", database)
		c.Next()
	})
	connectorapi.Routes(router)
	statusapi.Routes(router)
	router.Run("0.0.0.0:8080")
}
