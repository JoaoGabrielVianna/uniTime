package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/controller/users"
)

func UserRoutes(r *gin.Engine) {
	// Grupo de rotas para users
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/create", users.CreateUser)
		userRoutes.GET("/", users.GetUsers)
		userRoutes.GET("/:id", users.GetUserById)
		userRoutes.PUT("/:id", users.UpdateUser)
		userRoutes.DELETE("/:id", users.DeleteUser)
	}
}
