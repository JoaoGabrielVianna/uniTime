package repository

import (
	"github.com/joaogabrielvianna/db"
	"github.com/joaogabrielvianna/entity"
)

var Courses []entity.Course

// Função para adicionar um curso ao repositório
func CreateCourse(course entity.Course) (entity.Course, error) {
	if err := db.DB.Create(&course).Error; err != nil {
		return entity.Course{}, err
	}
	return course, nil
}

// Função para listar todos os usuários
func GetCourses() ([]entity.Course, error) {
	var courses []entity.Course
	if err := db.DB.Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

// Função para buscar um curso por ID
func GetCourseById(id string) (*entity.Course, error) {
	var course entity.Course
	if err := db.DB.Where("id = ?", id).First(&course).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

// Função para atualizar um curso
func UpdateCourse(id string, updateCourse entity.Course) (*entity.Course, error) {
	var course entity.Course
	if err := db.DB.Where("id = ?", id).First(&course).Error; err != nil {
		return nil, err
	}

	if updateCourse.Name != "" {
		course.Name = updateCourse.Name
	}

	if updateCourse.Description != "" {
		course.Description = updateCourse.Description
	}

	if err := db.DB.Save(&course).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

// Função para deletar um curso
func DeleteCourse(id string) (bool, error) {
	var course entity.Course
	if err := db.DB.First(&course, "id=?", id).Error; err != nil {
		return false, err
	}

	if err := db.DB.Delete(&course).Error; err != nil {
		return false, err
	}
	return true, nil
}
