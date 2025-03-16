package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CourseRoutes(r *gin.Engine) {
	// Grupo de rotas para users
	courseRoutes := r.Group("/courses")
	{
		courseRoutes.POST("/create", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "CreateCourse",
			})
		})
		courseRoutes.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "GetCourses",
			})
		})
		courseRoutes.GET("/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "GetCourseByID",
			})
		})
		courseRoutes.PUT("/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "UpdateCourse",
			})
		})
		courseRoutes.DELETE("/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "DeleteCourse",
			})
		})
	}
}
