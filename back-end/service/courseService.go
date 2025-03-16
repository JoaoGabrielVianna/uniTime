package service

import (
	"errors"

	"github.com/joaogabrielvianna/entity"
	"github.com/joaogabrielvianna/repository"
)

// Função para criar um curse
func CreateCourse(course entity.Course) entity.Course {
	return repository.CreateCourse(course)
}

// Função para listar todos os curses
func GetCourses() []entity.Course {
	return repository.GetCourses()
}

// Função para buscar um curse por ID
func GetCourseById(id string) *entity.Course {
	return repository.GetCoursesById(id)
}

// Função para atualizar um curse
func UpdateCourse(id string, updateCourse entity.Course) *entity.Course {
	return repository.UpdateCourse(id, updateCourse)
}

func AddYearToCourse(id string, yearDetails entity.Year) (*entity.Course, error) {
	for i, course := range repository.Courses {
		if course.ID == id {

			course.Years = append(course.Years, yearDetails)
			repository.Courses[i] = course

			return &course, nil
		}
	}

	return nil, errors.New("Curso não encontrado")
}

// Função para deletar um curse
func DeleteCourse(id string) bool {
	return repository.DeleteCourse(id)
}
