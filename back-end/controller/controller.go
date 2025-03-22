package controller

import (
	"github.com/joaogabrielvianna/controller/courses"
	"github.com/joaogabrielvianna/controller/users"
)

func InitializeController() {
	users.InitializeUserController()
	courses.InitializeCoursesController()
}
