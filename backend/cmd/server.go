package main

import (
	"github.com/muksho/OSOORM/backend/config"
	"github.com/muksho/OSOORM/backend/routes"
)

func main() {
	config.ConnectDatabase()

	router := routes.SetupRoutes()
	router.Run(":8080")
}
