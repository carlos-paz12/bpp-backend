package routes

import (
	"spe/controllers"
	"spe/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigurarRotas(r *gin.Engine) {
	ctrls := controllers.NovoControlador()

	api := r.Group("/spe/api/v1")
	{
		//== Rotas gerais.
		// Documentação.
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		// Autenticação.
		api.POST("/login", ctrls.AutenticacaoCtrl.Autenticar)

		// == Rotas protegidas por autenticação.
		autenticacao := api.Group("/")
		autenticacao.Use(middlewares.RequerAutenticacao())
		{
			// Rotas de bolsistas (/bolsistas).
			bolsistas := autenticacao.Group("/bolsistas")
			{
				tecnico := bolsistas.Group("/")
				tecnico.Use(middlewares.RequerCargo("Técnico Administrativo"))
				{
					// Retorna todos os bolsistas.
					tecnico.GET("/", ctrls.BolsistaCtrl.Get)
					// Retorna os dados de um bolsista específico.
					tecnico.GET("/:id", ctrls.BolsistaCtrl.GetByID)
					// Cria novo bolsista.
					tecnico.POST("/", ctrls.BolsistaCtrl.Post)
					// Atualiza um bolsista existente.
					tecnico.PATCH("/:id", ctrls.BolsistaCtrl.PatchByID)
					// Remove um bolsista existente.
					tecnico.DELETE("/:id", ctrls.BolsistaCtrl.DeleteByID)
				}

				bolsista := bolsistas.Group("/me")
				bolsista.Use(middlewares.RequerCargo("Bolsista"))
				{
					// Retorna todos os dados associados ao bolsista autenticado.
					bolsista.GET("/", ctrls.BolsistaCtrl.GetMe)
				}
			}

			// Rotas de tecnicos (/tecnicos).
			tecnicos := autenticacao.Group("/tecnicos")
			{
				tecnico := tecnicos.Group("/")
				tecnico.Use(middlewares.RequerCargo("Técnico Administrativo"))
				{
					tecnico.GET("/")
					tecnico.GET("/:id")
					tecnico.POST("/", ctrls.TecnicoAdministrativoCtrl.Post)
					tecnico.PATCH("/:id")
					tecnico.DELETE("/:id")
				}
			}

			// Rotas de registros de ponto (/registrosPonto).
			registrosPonto := autenticacao.Group("/registrosPonto")
			{
				tecnico := registrosPonto.Group("/")
				tecnico.Use(middlewares.RequerCargo("Técnico Administrativo"))
				{
					tecnico.GET("/")
					tecnico.GET("/:id")
					tecnico.PATCH("/:id")
					tecnico.DELETE("/:id")
				}

				bolsista := registrosPonto.Group("/me")
				bolsista.Use(middlewares.RequerCargo("Bolsista"))
				{
					// Retorna todos os registros de ponto associados ao bolsista autenticado.
					bolsista.GET("/")
					// Cria novo registro de ponto associado ao bolsista autenticado.
					bolsista.POST("/")
					// Retorna o último registro de ponto associado ao bolsista autenticado.
					bolsista.GET("/ultimo")
				}
			}

			// justificativas := aut.Group("/justificativas")
			// {
			// 	tecnico := justificativas.Group("/")
			// 	tecnico.Use(middlewares.RequerCargo("Técnico Administrativo"))
			// 	{
			// 		tecnico.GET("/")
			// 		tecnico.GET("/:id")
			// 		tecnico.POST("/")
			// 		tecnico.PATCH("/:id")
			// 		tecnico.DELETE("/:id")
			// 	}
			// }

			// cargos := aut.Group("/cargos")
			// {
			// 	tecnico := cargos.Group("/")
			// 	tecnico.Use(middlewares.RequerCargo("Técnico Administrativo"))
			// 	{
			// 		tecnico.GET("/")
			// 		tecnico.GET("/:id")
			// 		tecnico.POST("/")
			// 		tecnico.PATCH("/:id")
			// 		tecnico.DELETE("/:id")
			// 	}
			// }

			// setores := aut.Group("/setores")
			// {
			// 	tecnico := setores.Group("/")
			// 	tecnico.Use(middlewares.RequerCargo("Técnico Administrativo"))
			// 	{
			// 		tecnico.GET("/")
			// 		tecnico.GET("/:id")
			// 		tecnico.POST("/")
			// 		tecnico.PATCH("/:id")
			// 		tecnico.DELETE("/:id")
			// 	}
			// }
		}
	}
}
