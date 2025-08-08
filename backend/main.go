package main

import (
	"ahsen3d/db"
	"log"
)

func main() {

	// Initialize the database connection
	if err := db.ConnectDB(); err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

}
