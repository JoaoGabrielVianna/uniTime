package users

import "github.com/joaogabrielvianna/config"

var logger *config.Logger

func InitializeController() {
	logger = config.GetLogger("Users")
}
