package main

import (
	"ahsen3d/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	if err := db.ConnectDB(); err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	// Create Gin router
	router := gin.Default()

	// Homepage route
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Server running with Gin!")
	})

	// Listen on port 3000
	log.Println("Gin server started on port 3000...")
	if err := router.Run(":3000"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
