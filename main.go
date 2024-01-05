// main.go
package main

import (
	"fmt"
	"log"

	// main.go
	"myapp/database"
	"myapp/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	fmt.Println("Testing.............")
	db, err := database.InitializeDB()
	if err != nil {
		log.Fatal("Could not initialize the database:", err)
	}

	// Initialize Gin
	router := gin.Default()

	// Set up routes
	routes.SetupRoutes(router, db)

	// Run the server
	fmt.Println("Starting the server on :8080...")
	if err := router.Run(":4000"); err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}
