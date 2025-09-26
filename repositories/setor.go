package repositories

import (
	"spe/models/domain"

	"gorm.io/gorm"
)

type SetorRepositorio interface {
	Salvar(s *domain.Setor) (*domain.Setor, error)
	BuscarTodos() ([]*domain.Setor, error)
	BuscarPeloID(id uint) (*domain.Setor, error)
	BuscarPeloNome(nome string) (*domain.Setor, error)
	Atualizar(s *domain.Setor) (*domain.Setor, error)
	Deletar(id uint) error
}

type SetorRepositorioImpl struct {
	db *gorm.DB
}

func NovoSetorRepositorio(db *gorm.DB) SetorRepositorio {
	return &SetorRepositorioImpl{db: db}
}

func (repo *SetorRepositorioImpl) Salvar(s *domain.Setor) (*domain.Setor, error) {
	err := repo.db.Create(s).Error
	return s, err
}

func (repo *SetorRepositorioImpl) BuscarTodos() ([]*domain.Setor, error) {
	var setores []*domain.Setor
	err := repo.db.Find(&setores).Error
	return setores, err
}

func (repo *SetorRepositorioImpl) BuscarPeloID(id uint) (*domain.Setor, error) {
	var setor domain.Setor
	err := repo.db.First(&setor, id).Error
	if err != nil {
		return nil, err
	}
	return &setor, nil
}

func (repo *SetorRepositorioImpl) BuscarPeloNome(nome string) (*domain.Setor, error) {
	var setor domain.Setor
	err := repo.db.Where("nome = ?", nome).First(&setor).Error
	if err != nil {
		return nil, err
	}
	return &setor, nil
}

func (repo *SetorRepositorioImpl) Atualizar(s *domain.Setor) (*domain.Setor, error) {
	err := repo.db.Save(s).Error
	return s, err
}

func (repo *SetorRepositorioImpl) Deletar(id uint) error {
	return repo.db.Delete(&domain.Setor{}, id).Error
}
