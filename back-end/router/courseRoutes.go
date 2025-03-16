package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/controller"
)

func CourseRoutes(r *gin.Engine) {
	// Grupo de rotas para users
	CourseRoutes := r.Group("/courses")
	{
		CourseRoutes.POST("/create", controller.CreateCourse)
		CourseRoutes.GET("/", controller.GetCourses)
		CourseRoutes.GET("/:id", controller.GetCoursesById)
		CourseRoutes.PUT("/:id", controller.UpdateCourse)
		CourseRoutes.POST("/:id/add-year", controller.AddYearToCourse)
		CourseRoutes.DELETE("/:id", controller.DeleteCourse)
	}
}
