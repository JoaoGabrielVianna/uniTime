package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/repository"
)

// Função para listar todos os usuários
func GetUsers(c *gin.Context) {
	users, err := repository.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})

}
