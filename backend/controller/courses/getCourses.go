package courses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/repository"
)

// Função para listar todos os usuários
func GetCourses(c *gin.Context) {
	courses, err := repository.GetCourses()
	if err != nil {
		logger.ErrorLog("Erro ao tentar listar os cursos: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.SuccessLog("Cursos listados com sucesso")

	c.JSON(http.StatusOK, gin.H{
		"courses": courses,
	})

}
