package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/controller"
)

func SetupRouter() {
	r := gin.Default()

	// Inicializar o controlador (para configurar os logs específicos do controller)
	controller.InitializeController()

	// Adicionando as rotas
	UserRoutes(r)
	CourseRoutes(r)

	r.Run("127.0.0.1:8000")
}
