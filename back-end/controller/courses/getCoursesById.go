package courses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/repository"
)

// Function para obter o user por ID
func GetCoursesById(c *gin.Context) {
	id := c.Param("id")

	course, err := repository.GetCourseById(id)
	if err != nil {
		logger.ErrorLog("Erro ao tentar obter o curso com ID " + id + ": " + err.Error())
		c.JSON(http.StatusNotFound, gin.H{"message": "Curso não encontrado"})
		return
	}

	logger.SuccessLog("Curso com ID " + id + " encontrado com sucesso")

	c.JSON(http.StatusOK, gin.H{"course": course})
}
