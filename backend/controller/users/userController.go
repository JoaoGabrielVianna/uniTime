package users

import "github.com/joaogabrielvianna/config"

var logger *config.Logger

func InitializeUserController() {
	logger = config.GetLogger("Users")
}
