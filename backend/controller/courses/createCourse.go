package courses

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joaogabrielvianna/entity"
	"github.com/joaogabrielvianna/repository"
)

// Função para criar um usuário
func CreateCourse(c *gin.Context) {
	var course entity.Course

	if err := c.ShouldBindJSON(&course); err != nil {
		logger.ErrorLog("Erro ao fazer o bind dos dados do curso: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Gerando o UUID para o ID do curso (se não estiver sendo gerado automaticamente)
	course.ID = uuid.New() // Gerando o UUID automaticamente
	// Adicionando a data
	course.CreatedAt = time.Now()

	// Criando o curso no banco de dados usando GORM
	createCourse, err := repository.CreateCourse(course)
	if err != nil {
		logger.ErrorLog("Erro ao criar o curso: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.SuccessLog("Curso criado com sucesso com ID " + createCourse.ID.String())

	c.JSON(http.StatusOK, gin.H{
		"message": "Curso criado com sucesso",
		"course":  createCourse,
	})
}
