package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/entity"
	"github.com/joaogabrielvianna/repository"
)

// Function para atualizar um user
func UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	var updatedUser *entity.User

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	updatedUser, err := repository.UpdateUser(id, *updatedUser)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Curso não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"course": updatedUser,
	})
}
