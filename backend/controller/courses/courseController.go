package courses

import "github.com/joaogabrielvianna/config"

var logger *config.Logger

func InitializeCoursesController() {
	logger = config.GetLogger("Courses")
}
