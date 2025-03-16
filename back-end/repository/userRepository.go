package repository

import (
	"slices"

	"github.com/joaogabrielvianna/entity"
)

var Users []entity.User

// Função para adicionar um usuário ao repositório
func CreateUser(user entity.User) entity.User {
	user.ID = len(Users) + 1
	Users = append(Users, user)
	return user
}

// Função para listar todos os usuários
func GetUsers() []entity.User {
	return Users
}

// Função para buscar um usuário por ID
func GetUserById(id int) *entity.User {
	for _, user := range Users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}

// Função para atualizar um usuário
func UpdateUser(id int, updateUser entity.User) *entity.User {
	for i, user := range Users {
		if user.ID == id {
			Users[i] = updateUser
			return &Users[i]
		}
	}
	return nil
}

// Função para deletar um usuário
func DeleteUser(id int) bool {
	for i, user := range Users {
		if user.ID == id {
			Users = slices.Delete(Users, i, i+1)
			return true
		}
	}
	return false
}
