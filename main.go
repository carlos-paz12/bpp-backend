// @title SPE API

// @version 1.0

// @description API do Sistema de Ponto Eletrônico - SPE

// @host localhost:8080

// @BasePath /spe/api/v1
package main

import (
	"spe/database"
	_ "spe/docs"
	"spe/routes"

	"log"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error to load '.env' file")
	}

	database.Conectar()

	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()

	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	origins := strings.Split(allowedOrigins, ",")
	r.Use(cors.New(cors.Config{
		AllowOrigins: origins,
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	routes.ConfigurarRotas(r)

	r.Run(":" + os.Getenv("SERVER_PORT"))
}
