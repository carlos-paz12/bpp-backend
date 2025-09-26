package services

import (
	"spe/models/domain"
	"spe/repositories"
)

type CargoServico interface {
	Salvar(c *domain.Cargo) (*domain.Cargo, error)
	PegarTodos() ([]*domain.Cargo, error)
	PegarPeloID(id uint) (*domain.Cargo, error)
	PegarPeloNome(nome string) (*domain.Cargo, error)
	Atualizar(c *domain.Cargo) (*domain.Cargo, error)
	Deletar(id uint) error
}

type CargoServicoImpl struct {
	cargoRepo repositories.CargoRepositorio
}

func NovoCargoServico(repo repositories.CargoRepositorio) CargoServico {
	return &CargoServicoImpl{cargoRepo: repo}
}

func (serv *CargoServicoImpl) Salvar(c *domain.Cargo) (*domain.Cargo, error) {
	return serv.cargoRepo.Salvar(c)
}

func (serv *CargoServicoImpl) PegarTodos() ([]*domain.Cargo, error) {
	return serv.cargoRepo.BuscarTodos()
}

func (serv *CargoServicoImpl) PegarPeloID(id uint) (*domain.Cargo, error) {
	return serv.cargoRepo.BuscarPeloID(id)
}

func (serv *CargoServicoImpl) PegarPeloNome(nome string) (*domain.Cargo, error) {
	return serv.cargoRepo.BuscarPeloNome(nome)
}

func (serv *CargoServicoImpl) Atualizar(c *domain.Cargo) (*domain.Cargo, error) {
	return serv.cargoRepo.Atualizar(c)
}

func (serv *CargoServicoImpl) Deletar(id uint) error {
	return serv.cargoRepo.Deletar(id)
}
