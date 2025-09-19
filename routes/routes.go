package routes

import (
	"spe/controllers"
	"spe/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRoutes define todas as rotas da API SPE.
func SetupRoutes(r *gin.Engine) {
	api := r.Group("/spe/api/v1")
	{
		// == Rotas públicas (não requerem autenticação).

		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		api.POST("/authenticate", controllers.AuthController{}.Authenticate)

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
					scholarshipsMe.POST("/pointRecords", controllers.PointRecordController{}.RegisterMyPoint)
					scholarshipsMe.GET("/pointRecords/last", controllers.PointRecordController{}.GetMyLastPoint)
					scholarshipsMe.GET("/justifications", controllers.JustificationController{}.GetMyJustifications)
					scholarshipsMe.POST("/justifications", controllers.JustificationController{}.CreateMyJustification)
				}
			}

			// === Rotas de admins (/admins).

			admins := auth.Group("/admins")
			admins.Use(middlewares.RequireRole("admin"))
			{
				admins.POST("/", controllers.AdminController{}.CreateAdmin)
				admins.GET("/", controllers.AdminController{}.GetAllAdmins)
				admins.DELETE("/:id", controllers.AdminController{}.DeleteAdminByID)
				admins.GET("/me", controllers.AdminController{}.Me)

				adminScholarships := admins.Group("/scholarships")
				{
					adminScholarships.POST("/", controllers.ScholarshipController{}.CreateScholarship)
					adminScholarships.GET("/", controllers.ScholarshipController{}.GetAllScholarships)
					adminScholarships.DELETE("/:id", controllers.ScholarshipController{}.DeleteScholarshipByID)
					adminScholarships.GET("/:id/pointRecords", controllers.PointRecordController{}.GetPointsByScholarshipID)
					adminScholarships.GET("/:id/justifications", controllers.JustificationController{}.GetJustificationsByScholarshipID)
				}

				adminJustifications := admins.Group("/justifications")
				{
					adminJustifications.GET("/", controllers.JustificationController{}.GetAllJustifications)
					adminJustifications.PATCH("/:id", controllers.JustificationController{}.UpdateJustification)
				}
			}
		}
	}
}
