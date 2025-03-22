package courses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/entity"
	"github.com/joaogabrielvianna/repository"
)

// Function para atualizar um user
func UpdateCourse(c *gin.Context) {
	id := c.Param("id")

	var updateCourse entity.Course

	if err := c.ShouldBindJSON(&updateCourse); err != nil {
		logger.ErrorLog("Erro ao tentar fazer bind no curso com ID " + id + ": " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	updatedCoursed, err := repository.UpdateCourse(id, updateCourse)
	if err != nil {
		logger.ErrorLog("Erro ao tentar atualizar o curso com ID " + id + ": " + err.Error())
		c.JSON(http.StatusNotFound, gin.H{"message": "Curso não encontrado"})
		return
	}

	logger.SuccessLog("Curso com ID " + id + " atualizado com sucesso")

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"course": updatedCoursed,
	})
}
