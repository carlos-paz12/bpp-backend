package routes

import (
    "github.com/gin-gonic/gin"
    "spe/controllers"
    "spe/middlewares"
)

func SetupRoutes(r *gin.Engine) {
    api := r.Group("/api")
    {
        api.POST("/register", controllers.Register)
        api.POST("/login", controllers.Login)

        // Rotas protegidas
        api.Use(middlewares.AuthMiddleware())
        api.POST("/ponto/marcar", controllers.MarcarPonto)
        api.GET("/ponto", controllers.ListaPontos)
    }
}
