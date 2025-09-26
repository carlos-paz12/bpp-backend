package services

import (
	"spe/models/domain"
	"spe/repositories"
)

type TecnicoAdministrativoServico interface {
	Salvar(ta *domain.TecnicoAdministrativo) (*domain.TecnicoAdministrativo, error)
	PegarTodos() ([]*domain.TecnicoAdministrativo, error)
	PegarPeloID(id uint) (*domain.TecnicoAdministrativo, error)
	Atualizar(ta *domain.TecnicoAdministrativo) (*domain.TecnicoAdministrativo, error)
	Deletar(id uint) error
}

type TecnicoAdministrativoImpl struct {
	tecnicoAdministrativoRepo repositories.TecnicoAdministrativoRepositorio
}

func NovoTecnicoAdministrativoServico(repo repositories.TecnicoAdministrativoRepositorio) TecnicoAdministrativoServico {
	return &TecnicoAdministrativoImpl{tecnicoAdministrativoRepo: repo}
}

func (serv *TecnicoAdministrativoImpl) Salvar(ta *domain.TecnicoAdministrativo) (*domain.TecnicoAdministrativo, error) {
	return serv.tecnicoAdministrativoRepo.Salvar(ta)
}

func (serv *TecnicoAdministrativoImpl) PegarTodos() ([]*domain.TecnicoAdministrativo, error) {
	return serv.tecnicoAdministrativoRepo.BuscarTodos()
}

func (serv *TecnicoAdministrativoImpl) PegarPeloID(id uint) (*domain.TecnicoAdministrativo, error) {
	return serv.tecnicoAdministrativoRepo.BuscarPeloID(id)
}

func (serv *TecnicoAdministrativoImpl) Atualizar(ta *domain.TecnicoAdministrativo) (*domain.TecnicoAdministrativo, error) {
	return serv.tecnicoAdministrativoRepo.Atualizar(ta)
}

func (serv *TecnicoAdministrativoImpl) Deletar(id uint) error {
	return serv.tecnicoAdministrativoRepo.Deletar(id)
}
