package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/entity"
	"github.com/joaogabrielvianna/service"
)

// Função para criar um usuário
func CreateCourse(c *gin.Context) {
	var course entity.Course

	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Adicionando a data
	course.CreatedAt = time.Now()

	// Chamando o servico para criar o curso
	createCourse := service.CreateCourse(course)

	c.JSON(http.StatusOK, gin.H{
		"message": "Usuário criado com sucesso",
		"course":  createCourse,
	})
}

// Função para listar todos os usuários
func GetCourses(c *gin.Context) {
	courses := service.GetCourses()
	c.JSON(http.StatusOK, gin.H{
		"courses": courses,
	})

}

// Function para obter o user por ID
func GetCoursesById(c *gin.Context) {
	id := c.Param("id")

	course := service.GetCourseById(id)
	if course == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Curso não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"course": course})
}

// Function para atualizar um user
func UpdateCourse(c *gin.Context) {
	id := c.Param("id")

	var updateCourse entity.Course

	if err := c.ShouldBindJSON(&updateCourse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	updatedCoursed := service.UpdateCourse(id, updateCourse)
	if updatedCoursed == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Curso não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"course": updateCourse,
	})
}

func AddYearToCourse(c *gin.Context) {
	id := c.Param("id")
	var yearDetails entity.Year

	yearDetails.ID = fmt.Sprintf("%d%s", yearDetails.Year+1, id)

	if err := c.ShouldBindJSON(&yearDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	updatedCourse, err := service.AddYearToCourse(id, yearDetails)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Curso não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"course": updatedCourse,
	})
}

// Function para deletar o curso
func DeleteCourse(c *gin.Context) {
	id := c.Param("id")

	deleted := service.DeleteCourse(id)
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"message": "Curso não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Curso deletado com sucesso",
	})
}
