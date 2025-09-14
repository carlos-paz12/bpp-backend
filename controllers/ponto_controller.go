package controllers

import (
	"spe/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

type PontoController struct{}

var pontoService = services.PontoService{}

func (PontoController) Create(c *gin.Context) {
	bolsistaID := c.GetInt("user_id")
	ponto := pontoService.Create(bolsistaID)

	c.JSON(http.StatusOK, gin.H{
		"novo_ponto": ponto,
		"error":      "",
		"status":     http.StatusOK,
	})
}

func (PontoController) RetrieveAllWhereBolsistaId(c *gin.Context) {
	bolsistaID := c.GetInt("user_id")
	pontos := pontoService.RetrieveAll(bolsistaID)
	c.JSON(http.StatusOK, gin.H{
		"pontos": pontos,
		"error":  "",
		"status": http.StatusOK,
	})
}

func (PontoController) RetrieveLastWhereBolsistaId(c *gin.Context) {
	bolsistaID := c.GetInt("user_id")
	ultimo := pontoService.RetrieveLast(bolsistaID)
	c.JSON(http.StatusOK, gin.H{
		"ultimo": ultimo,
		"error":  "",
		"status": http.StatusOK,
	})
}
