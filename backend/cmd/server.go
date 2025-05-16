package main

import (
	"github.com/muksho/OSOORM/tree/development/backend/config"
	"github.com/muksho/OSOORM/tree/development/backend/routes"
)

func main() {
	config.ConnectDatabase()

	router := routes.SetupRoutes()
	router.Run(":8080")
}
