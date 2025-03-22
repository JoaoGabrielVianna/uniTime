package users

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/entity"
	"github.com/joaogabrielvianna/repository"
)

func CreateUser(c *gin.Context) {
	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		logger.ErrorLog("Erro ao fazer o bind dos dados do usuário: ")
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	user.CreatedAt = time.Now()

	createUser, err := repository.CreateUser(user)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Erro ao criar o usuário: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	logger.SuccessLog("Usuário criado com sucesso")

	c.JSON(http.StatusOK, gin.H{
		"message": "Usuário criado com sucesso",
		"course":  createUser,
	})
}
