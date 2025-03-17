package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/controller/courses"
)

func CourseRoutes(r *gin.Engine) {
	// Grupo de rotas para users
	CourseRoutes := r.Group("/courses")
	{
		CourseRoutes.POST("/create", courses.CreateCourse)
		CourseRoutes.GET("/", courses.GetCourses)
		CourseRoutes.GET("/:id", courses.GetCoursesById)
		CourseRoutes.PUT("/:id", courses.UpdateCourse)
		CourseRoutes.DELETE("/:id", courses.DeleteCourse)
	}
}
