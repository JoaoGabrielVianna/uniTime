package courses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/repository"
)

// Function para deletar o curso
func DeleteCourse(c *gin.Context) {
	id := c.Param("id")

	_, err := repository.DeleteCourse(id)
	if err != nil {
		logger.ErrorLog("Erro ao tentar deletar o curso com ID " + id + ": " + err.Error())
		c.JSON(http.StatusNotFound, gin.H{"message": "Curso não encontrado"})
		return
	}

	logger.SuccessLog("Curso com ID " + id + " deletado com sucesso")

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Curso deletado com sucesso",
	})
}
