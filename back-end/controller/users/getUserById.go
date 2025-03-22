package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/repository"
)

// Função para obter o usuário por ID
func GetUserById(c *gin.Context) {
	id := c.Param("id")
	// id, err := strconv.Atoi(idStr) // O erro da conversão agora é tratado
	// if err != nil {
	// 	logger.ErrorLog(fmt.Sprintf("Erro ao converter ID para inteiro: %v", err))
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
	// 	return
	// }

	user, err := repository.GetUserById(id)
	if err != nil {
		// Log de erro, quando o usuário não é encontrado
		logger.ErrorLog(fmt.Sprintf("Erro ao buscar o usuário com ID %s: %v", id, err))
		c.JSON(http.StatusNotFound, gin.H{"message": "Usuário não encontrado"})
		return
	}

	// Log de sucesso, indicando que o usuário foi encontrado com sucesso
	logger.SuccessLog(fmt.Sprintf("Usuário com ID %s encontrado", id))

	c.JSON(http.StatusOK, gin.H{"user": user})
}
