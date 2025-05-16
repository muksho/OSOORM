package main

import (
	"github.com/muksho/OSOORM/config"
	"github.com/muksho/OSOORM/routes"
)

func main() {
	config.ConnectDatabase()

	router := routes.SetupRoutes()
	router.Run(":8080")
}
