package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/repository"
)

// Function para deletar o user
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	_, err := repository.DeleteUser(id)
	fmt.Println("ERRO MEU------------------", err, "ID", id)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Erro ao deletar o usuário com ID %s: %v", id, err))
		c.JSON(http.StatusNotFound, gin.H{"message": "User não encontrado"})
		return
	}

	logger.SuccessLog(fmt.Sprintf("Usuário com ID %s deletado com sucesso", id))

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "User deletado com sucesso",
	})
}
