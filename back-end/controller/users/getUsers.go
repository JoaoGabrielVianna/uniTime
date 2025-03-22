package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/repository"
)

// Função para listar todos os usuários
func GetUsers(c *gin.Context) {
	users, err := repository.GetUsers()
	if err != nil {
		// Log de erro, com a mensagem do erro que ocorreu
		logger.ErrorLog(fmt.Sprintf("Erro ao listar os usuários: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Log de sucesso, indicando que a listagem de usuários foi bem-sucedida
	logger.SuccessLog("Usuários listados com sucesso")

	// Retorna a lista de usuários
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
