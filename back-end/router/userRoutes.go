package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	// Grupo de rotas para users
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/create", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "CreateUser",
			})
		})
		userRoutes.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "GetUsers",
			})
		})
		userRoutes.GET("/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "GetUserByID",
			})
		})
		userRoutes.PUT("/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "UpdateUser",
			})
		})
		userRoutes.DELETE("/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "DeleteUser",
			})
		})
	}
}
