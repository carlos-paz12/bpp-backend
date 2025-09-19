package controllers

import (
	"net/http"
	"spe/models"
	"spe/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

// Authenticate godoc
// @Summary Autenticação de usuário.
// @Description Realiza autenticação de usuário a partir de **username** e **password** e retorna token de acesso.
// @Tags auth
// @Accept json
// @Produce json
// @Param Body body models.AuthRequest true "Credenciais de login."
// @Success 200 {object} models.ApiResponse
// @Failure 400 {object} models.ApiResponse
// @Router /authenticate [post]
func (AuthController) Authenticate(c *gin.Context) {
	body := &models.AuthRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.ApiResponse{
			Message:  "",
			Error:    err.Error(),
			HttpCode: http.StatusBadRequest,
		})
		return
	}

	loginResponse, err := services.AuthService{}.Login(body.Username, body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ApiResponse{
			Message:  "",
			Error:    err.Error(),
			HttpCode: http.StatusBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, models.ApiResponse{
		Data:     loginResponse,
		Message:  "user logged in successfully.",
		Error:    "",
		HttpCode: http.StatusOK,
	})
}
