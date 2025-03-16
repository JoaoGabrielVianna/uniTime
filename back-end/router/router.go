package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Adicionando as rotas
	UserRoutes(r)
	CourseRoutes(r)

	return r
}
