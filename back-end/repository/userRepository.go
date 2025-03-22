package repository

import (
	"github.com/joaogabrielvianna/db"
	"github.com/joaogabrielvianna/entity"
)

var Users []entity.User

// Função para adicionar um usuário ao repositório
func CreateUser(user entity.User) (entity.User, error) {
	// Contar o número de usuários existentes no banco
	var count int64
	if err := db.DB.Model(&entity.User{}).Count(&count).Error; err != nil {
		return entity.User{}, err
	}

	// Definir o ID do usuário com base na quantidade atual de usuários
	user.ID = int(count + 1) // Conversão explícita de int64 para int

	if err := db.DB.Create(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

// Função para listar todos os usuários
func GetUsers() ([]entity.User, error) {
	var users []entity.User
	if err := db.DB.Find(&users).Error; err != nil {
		return []entity.User{}, err
	}
	return users, nil
}

// Função para buscar um usuário por ID
func GetUserById(id string) (*entity.User, error) {
	var user *entity.User

	if err := db.DB.First(&user, "id=?", id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Função para atualizar um usuário
func UpdateUser(id string, updateUser entity.User) (*entity.User, error) {
	var user entity.User

	// Buscar o usuário no banco de dados usando o ID
	result := db.DB.First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error // Se o usuário não for encontrado, retorna erro
	}

	if updateUser.Name != "" {
		user.Name = updateUser.Name
	}
	if updateUser.Email != "" {
		user.Email = updateUser.Email
	}
	if updateUser.Role != "" {
		user.Role = updateUser.Role
	}
	if updateUser.Course != nil { // Verifica se o campo "course" não é nil
		user.Course = updateUser.Course
	}
	if updateUser.CourseID != "" {
		user.CourseID = updateUser.CourseID
	}

	if err := db.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Função para deletar um usuário
func DeleteUser(id string) (bool, error) {
	var user entity.User

	if err := db.DB.First(&user, "id=?", id).Error; err != nil {
		return false, err
	}
	if err := db.DB.Delete(&user).Error; err != nil {
		return false, err
	}

	return true, nil
}
