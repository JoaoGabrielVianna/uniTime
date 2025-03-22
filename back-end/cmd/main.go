package main

import (
	"github.com/joaogabrielvianna/config"
	"github.com/joaogabrielvianna/controller/users"
	"github.com/joaogabrielvianna/db"
	"github.com/joaogabrielvianna/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("Main")
	users.InitializeController()
	logger.InfoLog("UniTime API rodando...")

	// Conectare ao banco de dados
	db.ConnectDataBase()

	r := router.SetupRouter()
	r.Run("127.0.0.1:8000")
}
