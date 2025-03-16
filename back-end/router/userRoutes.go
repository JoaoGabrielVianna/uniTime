package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/controller"
)

func UserRoutes(r *gin.Engine) {
	// Grupo de rotas para users
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/create", controller.CreateUser)
		userRoutes.GET("/", controller.GetUsers)
		userRoutes.GET("/:id", controller.GetUserById)
		userRoutes.PUT("/:id", controller.UpdateUser)
		userRoutes.DELETE("/:id", controller.DeleteUser)
	}
}
