package handlers

import (
	"net/http"
	"posts-api/internal/config"
	"posts-api/internal/models"

	"github.com/gin-gonic/gin"
)

func RetornarTodos(c *gin.Context) {

	var usuarios []models.User
	config.DB.Find(&usuarios)

	c.JSON(http.StatusOK, gin.H{
		"usuarios": usuarios,
	})
}
