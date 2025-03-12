package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/entity"
)

// Simulando uma lista de usuários
var users = []entity.User{
	{ID: 1, Name: "John Doe", Email: "john@example.com", Password: "1234", CreatedAt: time.Now(), Course: "Go Programming"},
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}
