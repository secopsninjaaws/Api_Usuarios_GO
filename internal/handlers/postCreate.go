package handlers

import (
	"posts-api/internal/config"
	"posts-api/internal/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {

	var usuario struct {
		Nome      string
		Sobrenome string
	}

	c.BindJSON(&usuario)

	user := models.User{Nome: usuario.Nome, Sobrenome: usuario.Sobrenome}

	result := config.DB.Create(&user)
	if result.Error != nil {
		panic("Erro ao cadastrar usuário.")
	}
}
