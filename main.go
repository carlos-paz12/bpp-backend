package main

import (
    "github.com/gin-gonic/gin"
    "spe/routes"
    "github.com/gin-contrib/cors"
)

func main() {
    r := gin.Default()
    r.Use(cors.Default()) // Allow frontend requests.

    routes.SetupRoutes(r)
    r.Run(":8080")
}
