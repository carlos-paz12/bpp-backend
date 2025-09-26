package services

import (
	"spe/models/domain"
	"spe/repositories"
)

type BolsistaServico interface {
	Salvar(b *domain.Bolsista) (*domain.Bolsista, error)
	PegarTodos() ([]*domain.Bolsista, error)
	PegarPeloID(id uint) (*domain.Bolsista, error)
	Atualizar(b *domain.Bolsista) (*domain.Bolsista, error)
	Deletar(id uint) error
}

type BolsistaServicoImpl struct {
	bolsistaRepo repositories.BolsistaRepositorio
}

func NovoBolsistaServico(repo repositories.BolsistaRepositorio) BolsistaServico {
	return &BolsistaServicoImpl{bolsistaRepo: repo}
}

func (serv *BolsistaServicoImpl) Salvar(b *domain.Bolsista) (*domain.Bolsista, error) {
	return serv.bolsistaRepo.Salvar(b)
}

func (serv *BolsistaServicoImpl) PegarTodos() ([]*domain.Bolsista, error) {
	return serv.bolsistaRepo.BuscarTodos()
}

func (serv *BolsistaServicoImpl) PegarPeloID(id uint) (*domain.Bolsista, error) {
	return serv.bolsistaRepo.BuscarPeloID(id)
}

func (serv *BolsistaServicoImpl) Atualizar(b *domain.Bolsista) (*domain.Bolsista, error) {
	return serv.bolsistaRepo.Atualizar(b)
}

func (serv *BolsistaServicoImpl) Deletar(id uint) error {
	return serv.bolsistaRepo.Deletar(id)
}
