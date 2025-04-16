package users

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joaogabrielvianna/entity"
	"github.com/joaogabrielvianna/repository"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		logger.ErrorLog("Erro ao fazer o bind dos dados do usuário: ")
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Gerando o UUID para o ID do curso (se não estiver sendo gerado automaticamente)
	user.ID = uuid.New() // Gerando o UUID automaticamente
	user.CreatedAt = time.Now()

	// Criptografando a senha antes de salvar no banco de dados
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Erro ao criptografar a senha: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Erro interno ao criptografar a senha"})
		return
	}
	logger.InfoLog(fmt.Sprintf("Senha original recebida: %s", user.Password))
	// Atribuindo a senha criptografada ao usuário
	user.Password = string(hashedPassword)

	logger.InfoLog(fmt.Sprintf("Senha criptografada antes de salvar: %s", user.Password))

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
