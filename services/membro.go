package services

import (
	"spe/models/domain"
	"spe/repositories"

	"golang.org/x/crypto/bcrypt"
)

type MembroServico interface {
	Salvar(m *domain.Membro) (*domain.Membro, error)
	PegarTodos() ([]*domain.Membro, error)
	PegarPeloID(id uint) (*domain.Membro, error)
	PegarPeloNomeUsuario(nomeUsuario string) (*domain.Membro, error)
	PegarCargosPeloMembroID(id uint) ([]domain.Membro, error)
	Atualizar(m *domain.Membro) (*domain.Membro, error)
	Deletar(id uint) error
}

type MembroServicoImpl struct {
	membroRepo repositories.MembroRepositorio
}

func NovoMembroServico(repo repositories.MembroRepositorio) MembroServico {
	return &MembroServicoImpl{membroRepo: repo}
}

func (serv *MembroServicoImpl) Salvar(m *domain.Membro) (*domain.Membro, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(m.SenhaHash), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	m.SenhaHash = string(hashed)

	return serv.membroRepo.Salvar(m)
}

func (serv *MembroServicoImpl) PegarTodos() ([]*domain.Membro, error) {
	return serv.membroRepo.BuscarTodos()
}

func (serv *MembroServicoImpl) PegarPeloID(id uint) (*domain.Membro, error) {
	return serv.membroRepo.BuscarPeloID(id)
}

func (serv *MembroServicoImpl) PegarPeloNomeUsuario(nomeUsuario string) (*domain.Membro, error) {
	return serv.membroRepo.BuscarPeloNomeUsuario(nomeUsuario)
}

func (serv *MembroServicoImpl) PegarCargosPeloMembroID(id uint) ([]domain.Membro, error) {
	return serv.PegarCargosPeloMembroID(id)
}

func (serv *MembroServicoImpl) Atualizar(m *domain.Membro) (*domain.Membro, error) {
	return serv.membroRepo.Atualizar(m)
}

func (serv *MembroServicoImpl) Deletar(id uint) error {
	return serv.membroRepo.Deletar(id)
}
