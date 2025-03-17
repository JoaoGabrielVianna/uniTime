package repository

import (
	"github.com/joaogabrielvianna/db"
	"github.com/joaogabrielvianna/entity"
)

var Users []entity.User

// Função para adicionar um usuário ao repositório
func CreateUser(user entity.User) (entity.User, error) {
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
func GetUserById(id int) (*entity.User, error) {
	var user *entity.User

	if err := db.DB.Where("id = ?").First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Função para atualizar um usuário
func UpdateUser(id int, updateUser entity.User) (*entity.User, error) {
	var user entity.User

	if err := db.DB.Where("id = ?").First(&user).Error; err != nil {
		return nil, err
	}

	user.Name = updateUser.Name
	user.Email = updateUser.Email

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
