package controllers

import (
	"spe/models/dto"
	"spe/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

type AutenticacaoControlador interface {
	Autenticar(c *gin.Context)
}

type AutenticacaoControladorImpl struct {
	autenticacaoServ services.AutenticacaoServico
}

func NovoAutenticacaoControlador(autenticacaoServ services.AutenticacaoServico) AutenticacaoControlador {
	return &AutenticacaoControladorImpl{
		autenticacaoServ: autenticacaoServ,
	}
}

// Autenticar godoc
// @Summary      Autenticação de membro.
// @Description  Realiza autenticação de membro na API e retorna token de autenticação.
// @Tags         autenticação
// @Accept       json
// @Produce      json
// @Param        request	body		dto.ReqAutenticacaoDTO	true	"Credenciais de login"
// @Success      200		{object}	dto.ResAutenticacaoDTO
// @Failure      400		{object}	dto.ErroDTO
// @Router       /login		[post]
func (ctrl *AutenticacaoControladorImpl) Autenticar(c *gin.Context) {
	req := &dto.ReqAutenticacaoDTO{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErroDTO{
			Mensagem: err.Error(),
			Codigo:   http.StatusBadRequest,
		})
		return
	}

	res, err := ctrl.autenticacaoServ.Autenticar(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErroDTO{
			Mensagem: err.Error(),
			Codigo:   http.StatusBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
