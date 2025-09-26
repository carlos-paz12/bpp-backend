package services

import (
	"spe/models/domain"
	"spe/repositories"
)

type VinculoServico interface {
	Salvar(v *domain.Vinculo) (*domain.Vinculo, error)
	PegarTodos() ([]*domain.Vinculo, error)
	PegarPeloID(id uint) (*domain.Vinculo, error)
	Atualizar(v *domain.Vinculo) (*domain.Vinculo, error)
	Deletar(id uint) error
}

type VinculoServicoImpl struct {
	vinculoRepo repositories.VinculoRepositorio
}

func NovoVinculoServico(repo repositories.VinculoRepositorio) VinculoServico {
	return &VinculoServicoImpl{vinculoRepo: repo}
}

func (serv *VinculoServicoImpl) Salvar(v *domain.Vinculo) (*domain.Vinculo, error) {
	return serv.vinculoRepo.Salvar(v)
}

func (serv *VinculoServicoImpl) PegarTodos() ([]*domain.Vinculo, error) {
	return serv.vinculoRepo.BuscarTodos()
}

func (serv *VinculoServicoImpl) PegarPeloID(id uint) (*domain.Vinculo, error) {
	return serv.vinculoRepo.BuscarPeloID(id)
}

func (serv *VinculoServicoImpl) Atualizar(v *domain.Vinculo) (*domain.Vinculo, error) {
	return serv.vinculoRepo.Atualizar(v)
}

func (serv *VinculoServicoImpl) Deletar(id uint) error {
	return serv.vinculoRepo.Deletar(id)
}
