package main

import (
	"ahsen3d/db"
	"ahsen3d/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	if err := db.ConnectDB(); err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	// Initialize router
	router := gin.Default()
	// Static files
	router.NoRoute(func(c *gin.Context) {
		file := c.Request.URL.Path
		if file == "/" {
			file = "/index.html"
		}
		c.File("./static" + file)
	}) // All routes
	routes.Routes(router)
	// Start server
	log.Println("Gin server started on port 3000...")
	if err := router.Run(":3000"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
