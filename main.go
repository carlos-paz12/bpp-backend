// @title SPE API
// @version 1.0
// @description API do Sistema de Ponto Eletrônico - SPE
// @host localhost:8080
// @BasePath /spe/api/v1
package main

import (
	_ "spe/docs"
	"spe/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	routes.SetupRoutes(r)

	r.Run(":8080")
}
