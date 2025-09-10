package routes

import (
	"spe/controllers"
	"spe/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		//=== Rotas públicas.
		api.POST("/login", controllers.Login)
		api.POST("/register", controllers.Register)

		//=== Rotas protegidas por autenticação.
		auth := api.Group("/")
		auth.Use(middlewares.AuthMiddleware())
		{
			//=== Rotas para bolsistas.
			bolsista := auth.Group("/")
			bolsista.Use(middlewares.RequireRole("bolsista"))
			{
				bolsista.GET("/ponto", controllers.ListaPontos)
				bolsista.POST("/ponto/marcar", controllers.MarcarPonto)
				/// @todo: bolsista.POST("/ponto/justificar", controllers.JustificarPonto)
			}

			//=== Rotas para admin.
			admin := auth.Group("/admin")
			admin.Use(middlewares.RequireRole("admin"))
			{
				/// @todo: admin.GET("/bolsista/:id/pontos", controllers.VerPontosBolsista)
				/// @todo: admin.GET("/bolsistas", controllers.ListarBolsistas)
				/// @todo: admin.POST("/ponto/aceitar-justificativa", controllers.AceitarJustificativa)
			}
		}
	}
}
