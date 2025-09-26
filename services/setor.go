package services

import (
	"spe/models/domain"
	"spe/repositories"
)

type SetorServico interface {
	Salvar(s *domain.Setor) (*domain.Setor, error)
	PegarTodos() ([]*domain.Setor, error)
	PegarPeloID(id uint) (*domain.Setor, error)
	PegarPeloNome(nome string) (*domain.Setor, error)
	Atualizar(s *domain.Setor) (*domain.Setor, error)
	Deletar(id uint) error
}

type SetorServicoImpl struct {
	setorRepo repositories.SetorRepositorio
}

func NovoSetorServico(repo repositories.SetorRepositorio) SetorServico {
	return &SetorServicoImpl{setorRepo: repo}
}

func (serv *SetorServicoImpl) Salvar(s *domain.Setor) (*domain.Setor, error) {
	return serv.setorRepo.Salvar(s)
}

func (serv *SetorServicoImpl) PegarTodos() ([]*domain.Setor, error) {
	return serv.setorRepo.BuscarTodos()
}

func (serv *SetorServicoImpl) PegarPeloID(id uint) (*domain.Setor, error) {
	return serv.setorRepo.BuscarPeloID(id)
}

func (serv *SetorServicoImpl) PegarPeloNome(nome string) (*domain.Setor, error) {
	return serv.setorRepo.BuscarPeloNome(nome)
}

func (serv *SetorServicoImpl) Atualizar(s *domain.Setor) (*domain.Setor, error) {
	return serv.setorRepo.Atualizar(s)
}

func (serv *SetorServicoImpl) Deletar(id uint) error {
	return serv.setorRepo.Deletar(id)
}
