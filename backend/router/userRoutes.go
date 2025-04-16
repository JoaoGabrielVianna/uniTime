package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/controller/users"
	"github.com/joaogabrielvianna/middleware"
)

func UserRoutes(r *gin.Engine) {
	// Grupo de rotas para users
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/create", users.CreateUser)
		userRoutes.POST("/login", users.Login) // Login e geração do JWT
		userRoutes.GET("/", users.GetUsers)
		userRoutes.GET("/:id", users.GetUserById)
		userRoutes.PUT("/:id", users.UpdateUser)
		userRoutes.DELETE("/:id", users.DeleteUser)
	}

	// Rotas protegidas com JWT
	protected := r.Group("/auth")
	protected.Use(middleware.AuthMiddleware()) // Aplica o middleware AQUI primeiro
	{
		protected.GET("/authorized", func(c *gin.Context) {
			email := c.MustGet("email").(string) // Se você setou o email no middleware
			c.JSON(200, gin.H{
				"message": "Authorized",
				"email":   email,
			})
		})
	}
}
