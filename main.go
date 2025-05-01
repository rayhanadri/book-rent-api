package main

import (
	"library-api/config" // Import the config package
	"library-api/route"  // Import the route package
)

func init() {
	// Load environment variables and connect to the database
	config.LoadEnv()
	// Initialize the database connection
	config.Connect()
}

func main() {
	// @title Library API
	// @version 1.0
	// @description This is a sample server for a library API.
	// @host localhost:8080
	// @BasePath /api/v1/
	route.ExecRouter()

}
