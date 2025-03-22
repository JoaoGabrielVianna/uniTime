package main

import (
	"github.com/joaogabrielvianna/config"
	"github.com/joaogabrielvianna/db"

	"github.com/joaogabrielvianna/router"
)

func main() {
	// Inicializar o logger
	logger := config.GetLogger("Main")

	// Conectar ao banco de dados
	db.ConnectDataBase()

	// Iniciar o servidor
	router.SetupRouter()

	// Iniciar a aplicação
	logger.InfoLog("UniTime API rodando...")

}
