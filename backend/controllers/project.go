package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muksho/OSOORM/backend/config"
	"github.com/muksho/OSOORM/backend/models"
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
