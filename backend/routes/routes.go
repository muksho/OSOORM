package routes

import (
	"github.com/muksho/OSOORM/backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/projects", controllers.CreateProject)

	return router
}
