package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/entity"
	"github.com/joaogabrielvianna/repository"
	"github.com/joaogabrielvianna/service"
)

// Função para criar um usuário
func CreateUser(c *gin.Context) {
	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Adicionando a data
	user.CreatedAt = time.Now()

	// Chamando o serviço para criar o usuário
	createUser := service.CreateUser(user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Usuário criado com sucesso",
		"user":    createUser,
	})
}

// Função para listar todos os usuários
func GetUsers(c *gin.Context) {
	users := service.GetUsers()
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// Function para obter o user por ID
func GetUserById(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido, deve ser um número.",
		})
		return
	}

	// Procurando o usuário pelo ID
	for _, user := range repository.Users {
		if user.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"user": user,
			})
			return
		}
	}

	// Caso não encontre o usuário
	c.JSON(http.StatusNotFound, gin.H{
		"message": "User não encontrado",
	})
}

// Function para atualizar um user
func UpdateUser(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido, deve ser um número.",
		})
		return
	}

	var updateUser entity.User

	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	updatedUser := service.UpdateUser(id, updateUser)
	if updatedUser == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"user":   updatedUser,
	})
}

// Function para deletar o user
func DeleteUser(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido, deve ser um número.",
		})
		return
	}

	deleted := service.DeleteUser(id)
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"message": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Usuário deletado com sucesso",
	})
}
