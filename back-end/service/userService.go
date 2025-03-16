package service

import (
	"github.com/joaogabrielvianna/entity"
	"github.com/joaogabrielvianna/repository"
)

// Função para criar um usuário
func CreateUser(user entity.User) entity.User {
	return repository.CreateUser(user)
}

// Função para listar todos os usuários
func GetUsers() []entity.User {
	return repository.GetUsers()
}

// Função para buscar um usuário por ID
func GetUserById(id int) *entity.User {
	return repository.GetUserById(id)
}

// Função para atualizar um usuário
func UpdateUser(id int, updateUser entity.User) *entity.User {
	return repository.UpdateUser(id, updateUser)
}

// Função para deletar um usuário
func DeleteUser(id int) bool {
	return repository.DeleteUser(id)
}
