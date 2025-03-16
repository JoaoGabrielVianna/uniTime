package repository

import (
	"slices"

	"github.com/joaogabrielvianna/entity"
)

var Courses []entity.Course

// Função para adicionar um curso ao repositório
func CreateCourse(course entity.Course) entity.Course {
	Courses = append(Courses, course)
	return course
}

// Função para listar todos os usuários
func GetCourses() []entity.Course {
	return Courses
}

// Função para buscar um curso por ID
func GetCoursesById(id string) *entity.Course {
	for _, course := range Courses {
		if course.ID == id {
			return &course
		}
	}
	return nil
}

// Função para atualizar um curso
func UpdateCourse(id string, updateCourse entity.Course) *entity.Course {
	for i, course := range Courses {
		if course.ID == id {
			Courses[i] = updateCourse
			return &course
		}
	}
	return nil
}

// Função para deletar um curso
func DeleteCourse(id string) bool {
	for i, course := range Courses {
		if course.ID == id {
			Courses = slices.Delete(Courses, i, i+1)
			return true
		}
	}
	return false
}
