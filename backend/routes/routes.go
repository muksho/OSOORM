package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muksho/OSOORM/backend/controllers"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/projects", controllers.CreateProject)

	return router
}
