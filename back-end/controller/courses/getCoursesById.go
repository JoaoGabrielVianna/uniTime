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
		c.JSON(http.StatusNotFound, gin.H{"message": "Curso não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"course": course})
}
