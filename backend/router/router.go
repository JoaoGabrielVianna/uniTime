package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/controller"
)

func SetupRouter() {
	r := gin.Default()

	// Habilita CORS para qualquer origem (ajuste conforme necessário)
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:4000", // Front-end
			"http://localhost:8000", // API
		},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Authorization", "Content-Type", "User-Agent"},
	}))
	// Inicializar o controlador (para configurar os logs específicos do controller)
	controller.InitializeController()

	// Adicionando as rotas
	UserRoutes(r)
	CourseRoutes(r)

	r.Run("127.0.0.1:8000")
}
