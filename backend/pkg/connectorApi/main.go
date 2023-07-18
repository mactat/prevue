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
	// router.POST(baseRoute+"/metrics", usersHandler)
	// router.POST(baseRoute+"/models", modelsHandler)
}

func helloHandler(c *gin.Context) {
	c.String(200, "Hi there!")
}

func sessionHandler(c *gin.Context) {
	var data types.SessionData
	c.BindJSON(&data)
	database := c.MustGet("database").(*sql.DB)
	modelId, err := db.SessionData(database, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert session data into the database"})
		log.Println("Failed to insert session data into the database")
	}
	log.Println(modelId)
	c.JSON(http.StatusOK, modelId)

	// err = db.InsertSessionProject(database, data)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert session data into the database"})
	// 	log.Println("Failed to insert session data into the database")
	// }
	// c.JSON(http.StatusOK, data)

}
