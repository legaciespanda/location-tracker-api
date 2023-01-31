package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/legaciespanda/location-tracker-api/database"
	"github.com/legaciespanda/location-tracker-api/routes"
)

func main() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "2023"
	}

	r := gin.Default()

	// grouping the API
	r.Group("/ap1/v1")

	//enabling logging
	r.Use(gin.Logger())

	//registering the routes
	routes.ApiRoutes(r)

	//use authentication
	//r.Use(middleware.Authentication)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "IPFTracker REST API",
		})
	})

	database.ConnectDatabaseAndRunMigration()

	log.Fatal(r.Run(":" + port))

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("I just recovered from an error")
	// 	}
	// }()

	// panic("An error happened!")
}
