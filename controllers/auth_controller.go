package controllers

import (
	"net/http"
	"spe/models"
	"spe/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (AuthController) Login(c *gin.Context) {
	var body struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Message:  "",
			Error:    err.Error(),
			HttpCode: http.StatusBadRequest,
		})
		return
	}

	loginResponse, err := services.AuthService{}.Login(body.Username, body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Message:  "",
			Error:    err.Error(),
			HttpCode: http.StatusBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Data:     loginResponse,
		Message:  "user logged in successfully.",
		Error:    "",
		HttpCode: http.StatusOK,
	})
}
