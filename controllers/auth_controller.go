package controllers

import (
	"net/http"
	"spe/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (AuthController) Login(c *gin.Context) {
	var body struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Corpo de requisição inválido.",
			"status": http.StatusBadRequest,
		})
		return
	}

	token, err := services.Login(body.Username, body.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":  "Usuário ou senha inválidos.",
			"status": http.StatusUnauthorized,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":  token,
		"status": http.StatusOK,
	})
}
