package routes

import (
	"github.com/muksho/OSOORM/tree/development/backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/projects", controllers.CreateProject)

	return router
}
