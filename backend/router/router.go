package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/controller"
)

func SetupRouter() {
	// Definir o Gin para rodar no modo 'release' (silencioso)
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Habilita CORS para qualquer origem (ajuste conforme necessário)
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:4000", // Front-end
			"http://localhost:8000", // API
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "User-Agent"},
		AllowCredentials: true,           // Permitir enviar cookies com requisições
		MaxAge:           12 * time.Hour, // O tempo que o navegador deve armazenar as permissões CORS
	}))
	// Inicializar o controlador (para configurar os logs específicos do controller)
	controller.InitializeController()

	// Adicionando as rotas
	UserRoutes(r)
	CourseRoutes(r)

	r.Run("0.0.0.0:8000")
}
