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
	route.ExecRouter()

}
