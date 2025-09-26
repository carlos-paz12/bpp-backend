package services

import (
	"spe/models/domain"
	"spe/repositories"
)

type RegistroDePontoServico interface {
	Salvar(rp *domain.RegistroDePonto) (*domain.RegistroDePonto, error)
	PegarTodos() ([]*domain.RegistroDePonto, error)
	PegarPeloID(id uint) (*domain.RegistroDePonto, error)
	Atualizar(rp *domain.RegistroDePonto) (*domain.RegistroDePonto, error)
	Deletar(id uint) error
}

type RegistroDePontoServicoImpl struct {
	registroDePontoRepo repositories.RegistroDePontoRepositorio
}

func NovoRegistroDePontoServico(repo repositories.RegistroDePontoRepositorio) RegistroDePontoServico {
	return &RegistroDePontoServicoImpl{registroDePontoRepo: repo}
}

func (serv *RegistroDePontoServicoImpl) Salvar(rp *domain.RegistroDePonto) (*domain.RegistroDePonto, error) {
	return serv.registroDePontoRepo.Salvar(rp)
}

func (serv *RegistroDePontoServicoImpl) PegarTodos() ([]*domain.RegistroDePonto, error) {
	return serv.registroDePontoRepo.BuscarTodos()
}

func (serv *RegistroDePontoServicoImpl) PegarPeloID(id uint) (*domain.RegistroDePonto, error) {
	return serv.registroDePontoRepo.BuscarPeloID(id)
}

func (serv *RegistroDePontoServicoImpl) Atualizar(rp *domain.RegistroDePonto) (*domain.RegistroDePonto, error) {
	return serv.registroDePontoRepo.Atualizar(rp)
}

func (serv *RegistroDePontoServicoImpl) Deletar(id uint) error {
	return serv.registroDePontoRepo.Deletar(id)
}
