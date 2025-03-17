package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/repository"
)

// Function para obter o user por ID
func GetUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	user, err := repository.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Curso não encontrado"})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"user": user})

}
