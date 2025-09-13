package routes

import (
	"spe/controllers/admin"
	"spe/controllers/auth"
	"spe/controllers/bolsista"
	"spe/controllers/justificativa"
	"spe/controllers/ponto"
	"spe/middlewares"

	"github.com/gin-gonic/gin"
)

// SetupRoutes define todas as rotas da API SPE.
// Aqui documentamos:
// - Quem pode acessar cada rota (público, bolsista, admin)
// - Propósito da rota
// - Método HTTP
// - URL
// - Expectativa de body
// - Retorno
func SetupRoutes(r *gin.Engine) {
	api := r.Group("/spe/api/v1")
	{
		// == Rotas públicas (não requerem autenticação).

		// @Brief	Realiza autenticação de usuário na API e retorna token de autenticação.
		// @Method	POST
		// @URL		/auth/login
		// @Body	{ "username": "", "senha": "" }
		// @Return	{ "token": "", "error": "", "status": "" }
		api.POST("/auth/login", auth.Login)

		// == Rotas protegidas (requerem autenticação).
		// Isto é, para qualquer uma das rotas a seguir, deve ter obrigatoriamente no corpo da requisição o token JWT.
		auth := api.Group("/")
		auth.Use(middlewares.AuthMiddleware())
		{
			// === Rotas do bolsista autenticado (/me).
			bolsistas := auth.Group("/me")
			bolsistas.Use(middlewares.RequireRole("bolsista"))
			{
				// @Brief     Retorna todos os registros de ponto do bolsista autenticado.
				// @Method    GET
				// @URL       /me/pontos
				// @Return    { "pontos": [], "error": "", "status": "" }
				bolsistas.GET("/pontos", ponto.RetrieveAllWhereBolsistaId)

				// @Brief     Retorna último registro de ponto do bolsista autenticado.
				// @Method    GET
				// @URL       /me/pontos/ultimo
				// @Return    { "ultimo_ponto": {}, "error": "", "status": "" }
				bolsistas.GET("/pontos/ultimo", ponto.RetrieveLastWhereBolsistaId)

				// @Brief     Registra novo ponto (entrada | saída).
				// @Method    POST
				// @URL       /me/pontos
				// @Return    { "error": "", "status": "" }
				bolsistas.POST("/pontos", ponto.Create)

				// @Brief     Retorna todas as justificativas do bolsista autenticado.
				// @Method    GET
				// @URL       /me/justificativas
				// @Return    { "justificativas": [], "error": "", "status": "" }
				bolsistas.GET("/justificativas", justificativa.RetrieveAllWhereBolsistaId)

				// @Brief     Registra nova justificativa.
				// @Method    POST
				// @URL       /me/justificativas
				// @Body      { "data": "DD-MM-AAAA", "horas": "", "minutos": "", "motivo": "..." }
				// @Return    { "error": "", "status": "" }
				bolsistas.POST("/justificativas", justificativa.Create)
			}

			// === Rotas do admin (/me).
			admins := auth.Group("/admins")
			admins.Use(middlewares.RequireRole("admin"))
			{
				// @Brief     Cadastra um novo usuário do tipo admin.
				// @Method    POST
				// @URL       /admins
				// @Body      { "name": "", "username": "", "password": "", "email": "" }
				// @Return    { "error": "", "status": "" }
				admins.POST("/", admin.Create)

				// @Brief     Lista todos os usuários do tipo admin.
				// @Method    GET
				// @URL       /admins
				// @Return    { "admins": [], "error": "", "status": "" }
				admins.GET("/", admin.RetrieveAll)

				// @Brief     Deleta um admin pelo ID.
				// @Method    DELETE
				// @URL       /admins/:id
				// @Return    { "error": "", "status": "" }
				admins.DELETE("/:id", admin.Delete)

				// @Brief     Cria um novo usuário do tipo bolsista.
				// @Method    POST
				// @URL       /admins/bolsistas
				// @Body      { "name": "", "username": "", "password": "", "email": "", "matricula": "", "horas_mensais": "" }
				// @Return    { "error": "", "status": "" }
				admins.POST("/bolsistas", bolsista.Create)

				// @Brief     Lista todos os usuários do tipo bolsista.
				// @Method    GET
				// @URL       /admins/bolsistas
				// @Return    { "bolsistas": [], "error": "", "status": "" }
				admins.GET("/bolsistas", bolsista.RetrieveAll)

				// @Brief     Deleta um bolsista pelo ID.
				// @Method    DELETE
				// @URL       /admins/bolsistas/:id
				// @Return    { "error": "", "status": "" }
				admins.DELETE("/bolsistas/:id", bolsista.Delete)

				// @Brief     Lista todos os pontos de um bolsista específico.
				// @Method    GET
				// @URL       /admins/bolsistas/:id/pontos
				// @Return    { "pontos": [], "error": "", "status": "" }
				admins.GET("/bolsistas/:id/pontos", ponto.RetrieveAllWhereBolsistaId)

				// @Brief     Lista todas as justificativas de um bolsista específico.
				// @Method    GET
				// @URL       /admins/bolsistas/:id/justificativas
				// @Return    { "justificativas": [], "error": "", "status": "" }
				admins.GET("/bolsistas/:id/justificativas", justificativa.RetrieveAllWhereBolsistaId)

				// @Brief     Lista todas as justificativas.
				// @Method    GET
				// @URL       /admins/justificativas
				// @Return    { "justificativas": [], "error": "", "status": "" }
				admins.GET("/justificativas", justificativa.RetrieveAll)

				// @Brief     Atualiza uma justificativa (aprovar ou rejeitar).
				// @Method    PATCH
				// @URL       /admins/justificativas/:id
				// @Body      { "status": "aprovado" | "rejeitado", "comentario": "" }
				// @Return    { "error": "", "status": "" }
				admins.PATCH("/justificativas/:id", justificativa.Update)
			}
		}
	}
}
