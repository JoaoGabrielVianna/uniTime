package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/entity"
	"github.com/joaogabrielvianna/repository"
)

// Function para atualizar um user
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var updatedUser *entity.User

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		logger.ErrorLog(fmt.Sprintf("Erro ao fazer o bind dos dados do usuário: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	updatedUser, err := repository.UpdateUser(id, *updatedUser)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Erro ao atualizar o usuário com ID %s: %v", id, err))
		c.JSON(http.StatusNotFound, gin.H{"message": "User não encontrado"})
		return
	}

	logger.SuccessLog(fmt.Sprintf("Usuário com ID %s atualizado com sucesso", id))

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"course": updatedUser,
	})
}
