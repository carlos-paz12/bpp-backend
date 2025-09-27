package controllers

import (
	"net/http"
	"spe/models/dto"
	"spe/services"

	"github.com/gin-gonic/gin"
)

type SetorControlador interface {
	Get(c *gin.Context)
}

type SetorControladorImpl struct {
	setorServ services.SetorServico
}

func NovoSetorControlador(setorServ services.SetorServico) SetorControlador {
	return &SetorControladorImpl{setorServ: setorServ}
}

func (ctrl *SetorControladorImpl) Get(c *gin.Context) {
	setores, err := ctrl.setorServ.PegarTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErroDTO{
			Mensagem: err.Error(),
			Codigo:   http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"setores": setores,
		"codigo":  http.StatusOK,
	})
}
