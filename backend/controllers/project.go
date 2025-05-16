package controllers

import (
	"net/http"

	"github.com/muksho/OSOORM/tree/development/backend/config"
	"github.com/muksho/OSOORM/tree/development/backend/models"

	"github.com/gin-gonic/gin"
)

func CreateProject(c *gin.Context) {
	var project models.Project

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&project)
	c.JSON(http.StatusCreated, gin.H{"message": "Project created", "project": project})
}
