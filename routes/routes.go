package routes

import (
	"spe/controllers"
	"spe/middlewares"

	"github.com/gin-gonic/gin"
)

// SetupRoutes define todas as rotas da API SPE.
func SetupRoutes(r *gin.Engine) {
	api := r.Group("/spe/api/v1")
	{
		// == Rotas públicas (não requerem autenticação).

		/* @Brief	Realiza autenticação de usuário na API e retorna token de autenticação.
		 * @Method	POST
		 * @URL		/spe/api/v1/auth/login
		 * @Body	{
		 *				"username": "",
		 *				"password": ""
		 *			}
		 * @Return	{
		 *				"data": {
		 *							"token": "",
		 *							"role": "",
		 *							"user": {
		 *										"id": n,
		 *										"name": ""
		 *										"username": ""
		 *										"email": ""
		 *							}
		 *				},
		 *				"error": "",
		 *				"status": n
		 *			}
		 */
		api.POST("/auth/login", controllers.AuthController{}.Login)

		// == Rotas protegidas (requerem autenticação).

		/*
		 * Para qualquer uma das rotas a seguir, deve ter enviar, obrigatoriamente, no
		 * header da requisição o campo "Authorization" contêndo o token recebido na
		 * autenticação
		 */

		auth := api.Group("/")
		auth.Use(middlewares.AuthMiddleware())
		{
			// === Rotas de scholarships (/scholarships).

			scholarships := auth.Group("/scholarships")
			scholarships.Use(middlewares.RequireRole("scholarship"))
			{
				scholarshipsMe := scholarships.Group("/me")
				{
					scholarshipsMe.GET("/", controllers.ScholarshipController{}.Me)

					/* @Brief     Retorna todos os registros de ponto do bolsista autenticado.
					 * @Method    GET
					 * @URL       /spe/api/v1/scholarships/me/pointRecords
					 */
					scholarshipsMe.GET("/pointRecords", controllers.PointRecordController{}.GetMyPoints)

					/* @Brief     Insere novo registro de ponto.
					 * @Method    POST
					 * @URL       /spe/api/v1/scholarships/me/pointRecords
					 */
					scholarshipsMe.POST("/pointRecords", controllers.PointRecordController{}.RegisterMyPoint)

					/* @Brief     Retorna último registro de ponto do bolsista autenticado.
					 * @Method    GET
					 * @URL       /spe/api/v1/scholarships/me/pointRecords/last
					 */
					scholarshipsMe.GET("/pointRecords/last", controllers.PointRecordController{}.GetMyLastPoint)

					/* @Brief     Retorna todas as justifications do bolsista autenticado.
					 * @Method    GET
					 * @URL       /spe/api/v1/scholarships/me/justifications
					 */
					scholarshipsMe.GET("/justifications", controllers.JustificationController{}.GetMyJustifications)

					/* @Brief     Insere nova justificativa.
					 * @Method    POST
					 * @URL       /spe/api/v1/scholarships/me/justifications
					 */
					scholarshipsMe.POST("/justifications", controllers.JustificationController{}.CreateMyJustification)
				}
			}

			// === Rotas de admins (/admins).

			admins := auth.Group("/admins")
			admins.Use(middlewares.RequireRole("admin"))
			{
				/* @Brief     Cadastra um novo usuário do tipo admin.
				 * @Method    POST
				 * @URL       /spe/api/v1/admins
				 */
				admins.POST("/", controllers.AdminController{}.CreateAdmin)

				/* @Brief     Lista todos os usuários do tipo admin.
				 * @Method    GET
				 * @URL       /spe/api/v1/admins
				 */
				admins.GET("/", controllers.AdminController{}.GetAllAdmins)

				/* @Brief     Deleta um admin pelo ID.
				 * @Method    DELETE
				 * @URL       /spe/api/v1/admins/:id
				 */
				admins.DELETE("/:id", controllers.AdminController{}.DeleteAdminByID)

				admins.GET("/me", controllers.AdminController{}.Me)

				adminScholarships := admins.Group("/scholarships")
				{
					/* @Brief     Cria um novo usuário do tipo bolsista.
					 * @Method    POST
					 * @URL       /spe/api/v1/admins/scholarships
					 */
					adminScholarships.POST("/", controllers.ScholarshipController{}.CreateScholarship)

					/* @Brief     Lista todos os usuários do tipo bolsista.
					 * @Method    GET
					 * @URL       /spe/api/v1/admins/scholarships
					 */
					adminScholarships.GET("/", controllers.ScholarshipController{}.GetAllScholarships)

					/* @Brief     Deleta um bolsista pelo ID.
					 * @Method    DELETE
					 * @URL       /spe/api/v1/admins/scholarships/:id
					 */
					adminScholarships.DELETE("/:id", controllers.ScholarshipController{}.DeleteScholarshipByID)

					/* @Brief     Lista todos os records de um bolsista específico.
					 * @Method    GET
					 * @URL       /spe/api/v1/admins/scholarships/:id/pointRecords
					 */
					adminScholarships.GET("/:id/pointRecords", controllers.PointRecordController{}.GetPointsByScholarshipID)

					/* @Brief     Lista todas as justifications de um bolsista específico.
					 * @Method    GET
					 * @URL       /spe/api/v1/admins/scholarships/:id/justifications
					 */
					adminScholarships.GET("/:id/justifications", controllers.JustificationController{}.GetJustificationsByScholarshipID)
				}

				adminJustifications := admins.Group("/justifications")
				{
					/* @Brief     Lista todas as justifications.
					 * @Method    GET
					 * @URL       /spe/api/v1/admins/justifications
					 */
					adminJustifications.GET("/", controllers.JustificationController{}.GetAllJustifications)

					/* @Brief     Atualiza uma justificativa (aprovar ou rejeitar).
					 * @Method    PATCH
					 * @URL       /spe/api/v1/admins/justifications/:id
					 */
					adminJustifications.PATCH("/:id", controllers.JustificationController{}.UpdateJustification)
				}
			}
		}
	}
}
