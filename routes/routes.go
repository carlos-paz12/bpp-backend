package routes

import (
	"spe/controllers"
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
		api.POST("/auth/login", controllers.AuthController{}.Login)

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
				bolsistas.GET("/pontos", controllers.PointRecordController{}.FindAllByScholarshipID)

				// @Brief     Retorna último registro de ponto do bolsista autenticado.
				// @Method    GET
				// @URL       /me/pontos/ultimo
				// @Return    { "ultimo_ponto": {}, "error": "", "status": "" }
				bolsistas.GET("/pontos/ultimo", controllers.PointRecordController{}.FindLastByScholarshipID)

				// @Brief     Registra novo ponto (entrada | saída).
				// @Method    POST
				// @URL       /me/pontos
				// @Return    { "error": "", "status": "" }
				bolsistas.POST("/pontos", controllers.PointRecordController{}.Create)

				// @Brief     Retorna todas as justificativas do bolsista autenticado.
				// @Method    GET
				// @URL       /me/justificativas
				// @Return    { "justificativas": [], "error": "", "status": "" }
				bolsistas.GET("/justificativas", controllers.JustificativaController{}.FindAllByScholarshipID)

				// @Brief     Registra nova justificativa.
				// @Method    POST
				// @URL       /me/justificativas
				// @Body      { "data": "DD-MM-AAAA", "horas": "", "minutos": "", "motivo": "..." }
				// @Return    { "error": "", "status": "" }
				bolsistas.POST("/justificativas", controllers.JustificativaController{}.Create)
			}

			// === Rotas do admin (/admin).
			admins := auth.Group("/admins")
			admins.Use(middlewares.RequireRole("admin"))
			{
				// @Brief     Cadastra um novo usuário do tipo admin.
				// @Method    POST
				// @URL       /admins
				// @Body      { "name": "", "username": "", "password": "", "email": "" }
				// @Return    { "error": "", "status": "" }
				admins.POST("/", controllers.AdminController{}.Create)

				// @Brief     Lista todos os usuários do tipo admin.
				// @Method    GET
				// @URL       /admins
				// @Return    { "admins": [], "error": "", "status": "" }
				admins.GET("/", controllers.AdminController{}.FindAll)

				// @Brief     Deleta um admin pelo ID.
				// @Method    DELETE
				// @URL       /admins/:id
				// @Return    { "error": "", "status": "" }
				admins.DELETE("/:id", controllers.AdminController{}.Delete)

				// @Brief     Cria um novo usuário do tipo bolsista.
				// @Method    POST
				// @URL       /admins/bolsistas
				// @Body      { "name": "", "username": "", "password": "", "email": "", "matricula": "", "horas_mensais": "" }
				// @Return    { "error": "", "status": "" }
				admins.POST("/bolsistas", controllers.BolsistaController{}.Create)

				// @Brief     Lista todos os usuários do tipo bolsista.
				// @Method    GET
				// @URL       /admins/bolsistas
				// @Return    { "bolsistas": [], "error": "", "status": "" }
				admins.GET("/bolsistas", controllers.BolsistaController{}.FindAll)

				// @Brief     Deleta um bolsista pelo ID.
				// @Method    DELETE
				// @URL       /admins/bolsistas/:id
				// @Return    { "error": "", "status": "" }
				admins.DELETE("/bolsistas/:id", controllers.BolsistaController{}.Delete)

				// @Brief     Lista todos os pontos de um bolsista específico.
				// @Method    GET
				// @URL       /admins/bolsistas/:id/pontos
				// @Return    { "pontos": [], "error": "", "status": "" }
				admins.GET("/bolsistas/:id/pontos", controllers.PointRecordController{}.FindAllByScholarshipID)

				// @Brief     Lista todas as justificativas de um bolsista específico.
				// @Method    GET
				// @URL       /admins/bolsistas/:id/justificativas
				// @Return    { "justificativas": [], "error": "", "status": "" }
				admins.GET("/bolsistas/:id/justificativas", controllers.JustificativaController{}.FindAllByScholarshipID)

				// @Brief     Lista todas as justificativas.
				// @Method    GET
				// @URL       /admins/justificativas
				// @Return    { "justificativas": [], "error": "", "status": "" }
				admins.GET("/justificativas", controllers.JustificativaController{}.FindAll)

				// @Brief     Atualiza uma justificativa (aprovar ou rejeitar).
				// @Method    PATCH
				// @URL       /admins/justificativas/:id
				// @Body      { "status": "aprovado" | "rejeitado", "comentario": "" }
				// @Return    { "error": "", "status": "" }
				admins.PATCH("/justificativas/:id", controllers.JustificativaController{}.Update)
			}
		}
	}
}
