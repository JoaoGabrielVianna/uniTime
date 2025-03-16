package main

import (
	"fmt"

	"github.com/joaogabrielvianna/router"
)

func main() {
	fmt.Println("UniTime API rodando...")

	r := router.SetupRouter()
	r.Run("127.0.0.1:8000")
}
