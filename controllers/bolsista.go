package controllers

import (
	"spe/models/domain"
	"spe/models/dto"
	"spe/services"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type BolsistaControlador interface {
	Get(c *gin.Context)
	GetByID(c *gin.Context)
	GetMe(c *gin.Context)
	Post(c *gin.Context)
	PatchByID(c *gin.Context)
	DeleteByID(c *gin.Context)
}

type BolsistaControladorImpl struct {
	bolsistaServ services.BolsistaServico
	membroServ   services.MembroServico
	cargoServ    services.CargoServico
	setorServ    services.SetorServico
	vinculoServ  services.VinculoServico
}

func NovoBolsistaControlador(
	bolsistaServ services.BolsistaServico,
	membroServ services.MembroServico,
	cargoServ services.CargoServico,
	setorServ services.SetorServico,
	vinculoServ services.VinculoServico,
) BolsistaControlador {
	return &BolsistaControladorImpl{
		bolsistaServ: bolsistaServ,
		membroServ:   membroServ,
		cargoServ:    cargoServ,
		setorServ:    setorServ,
		vinculoServ:  vinculoServ,
	}
}

func (ctrl *BolsistaControladorImpl) Get(c *gin.Context) {

}

func (ctrl *BolsistaControladorImpl) GetByID(c *gin.Context) {

}

func (ctrl *BolsistaControladorImpl) GetMe(c *gin.Context) {

}

func (ctrl *BolsistaControladorImpl) Post(c *gin.Context) {
	req := &dto.ReqCriacaoBolsistaDTO{}
	if err := c.ShouldBindJSON(&req); err != nil {
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

	novoBolsista := &domain.Bolsista{
		MembroID:           novoMembro.ID,
		CargaHorariaMensal: req.CargaHorariaMensal,
	}

	novoBolsista, err = ctrl.bolsistaServ.Salvar(novoBolsista)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErroDTO{
			Mensagem: err.Error(),
			Codigo:   http.StatusBadRequest,
		})

		err = ctrl.membroServ.Deletar(novoMembro.ID) // <- Rollback do membro criado.
		return
	}

	cargoBolsista, err := ctrl.cargoServ.PegarPeloNome("Bolsista")
	setor, err := ctrl.setorServ.PegarPeloNome(req.SetorNome)

	novoVinculo := &domain.Vinculo{
		MembroID:   novoMembro.ID,
		CargoID:    cargoBolsista.ID,
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
		"bolsista": dto.BolsistaDTO{
			MembroID:           novoMembro.ID,
			BolsistaID:         novoBolsista.ID,
			Matricula:          novoMembro.Matricula,
			NomeCompleto:       novoMembro.NomeCompleto,
			NomeUsuario:        novoMembro.NomeUsuario,
			Email:              novoMembro.Email,
			Celular:            novoMembro.Celular,
			SetorNome:          setor.Nome,
			CargaHorariaMensal: novoBolsista.CargaHorariaMensal,
		},
		"codigo": http.StatusCreated,
	})
}

func (ctrl *BolsistaControladorImpl) PatchByID(c *gin.Context) {

}

func (ctrl *BolsistaControladorImpl) DeleteByID(c *gin.Context) {

}
