package controllers

import (
	"spe/models/domain"
	"spe/models/dto"
	"spe/services"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TecnicoAdministrativoControlador interface {
	Get(c *gin.Context)
	GetByID(c *gin.Context)
	GetMe(c *gin.Context)
	Post(c *gin.Context)
	PatchByID(c *gin.Context)
	DeleteByID(c *gin.Context)
}

type TecnicoAdministrativoControladorImpl struct {
	tecnicoAdministrativoServ services.TecnicoAdministrativoServico
	membroServ                services.MembroServico
	cargoServ                 services.CargoServico
	setorServ                 services.SetorServico
	vinculoServ               services.VinculoServico
}

func NovoTecnicoAdministrativoControlador(
	tecnicoAdministrativoServ services.TecnicoAdministrativoServico,
	membroServ services.MembroServico,
	cargoServ services.CargoServico,
	setorServ services.SetorServico,
	vinculoServ services.VinculoServico,
) TecnicoAdministrativoControlador {
	return &TecnicoAdministrativoControladorImpl{
		tecnicoAdministrativoServ: tecnicoAdministrativoServ,
		membroServ:                membroServ,
		cargoServ:                 cargoServ,
		setorServ:                 setorServ,
		vinculoServ:               vinculoServ,
	}
}

func (ctrl *TecnicoAdministrativoControladorImpl) Get(c *gin.Context) {

}

func (ctrl *TecnicoAdministrativoControladorImpl) GetByID(c *gin.Context) {

}

func (ctrl *TecnicoAdministrativoControladorImpl) GetMe(c *gin.Context) {

}

func (ctrl *TecnicoAdministrativoControladorImpl) Post(c *gin.Context) {
	req := &dto.ReqCriacaoTecnicoAdministrativoDTO{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErroDTO{
			Mensagem: err.Error(),
			Codigo:   http.StatusBadRequest,
		})
		return
	}

	novoMembro := &domain.Membro{
		NomeCompleto: req.NomeCompleto,
		NomeUsuario:  req.NomeUsuario,
		SenhaHash:    req.Senha,
		Email:        req.Email,
		Celular:      req.Celular,
		Matricula:    req.Matricula,
	}

	novoMembro, err := ctrl.membroServ.Salvar(novoMembro)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErroDTO{
			Mensagem: err.Error(),
			Codigo:   http.StatusBadRequest,
		})
		return
	}

	novoTecnicoAdministrativo := &domain.TecnicoAdministrativo{
		MembroID: novoMembro.ID,
	}

	novoTecnicoAdministrativo, err = ctrl.tecnicoAdministrativoServ.Salvar(novoTecnicoAdministrativo)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErroDTO{
			Mensagem: err.Error(),
			Codigo:   http.StatusBadRequest,
		})

		err = ctrl.membroServ.Deletar(novoMembro.ID) // <- Rollback do membro criado.
		return
	}

	cargoTecnicoAdministrativo, err := ctrl.cargoServ.PegarPeloNome("Técnico Administrativo")
	setor, err := ctrl.setorServ.PegarPeloNome(req.SetorNome)

	novoVinculo := &domain.Vinculo{
		MembroID:   novoMembro.ID,
		CargoID:    cargoTecnicoAdministrativo.ID,
		IniciadoEm: time.Now(),
		Ativo:      true,
		SetorID:    setor.ID,
	}

	novoVinculo, err = ctrl.vinculoServ.Salvar(novoVinculo)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErroDTO{
			Mensagem: err.Error(),
			Codigo:   http.StatusBadRequest,
		})

		err = ctrl.membroServ.Deletar(novoMembro.ID) // <- Rollback do membro criado, o registro de bolsista já vai embora junto.
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"tecnico_administrativo": dto.TecnicoAdministrativoDTO{
			MembroID:                novoMembro.ID,
			TecnicoAdministrativoID: novoTecnicoAdministrativo.ID,
			Matricula:               novoMembro.Matricula,
			NomeCompleto:            novoMembro.NomeCompleto,
			NomeUsuario:             novoMembro.NomeUsuario,
			Email:                   novoMembro.Email,
			Celular:                 novoMembro.Celular,
			SetorNome:               setor.Nome,
		},
		"codigo": http.StatusCreated,
	})
}

func (ctrl *TecnicoAdministrativoControladorImpl) PatchByID(c *gin.Context) {

}

func (ctrl *TecnicoAdministrativoControladorImpl) DeleteByID(c *gin.Context) {

}
