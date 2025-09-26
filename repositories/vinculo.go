package repositories

import (
	"spe/models/domain"

	"gorm.io/gorm"
)

type VinculoRepositorio interface {
	Salvar(v *domain.Vinculo) (*domain.Vinculo, error)
	BuscarTodos() ([]*domain.Vinculo, error)
	BuscarPeloID(id uint) (*domain.Vinculo, error)
	Atualizar(v *domain.Vinculo) (*domain.Vinculo, error)
	Deletar(id uint) error
}

type VinculoRepositorioImpl struct {
	db *gorm.DB
}

func NovoVinculoRepositorio(db *gorm.DB) VinculoRepositorio {
	return &VinculoRepositorioImpl{db: db}
}

func (repo *VinculoRepositorioImpl) Salvar(v *domain.Vinculo) (*domain.Vinculo, error) {
	err := repo.db.Create(v).Error
	return v, err
}

func (repo *VinculoRepositorioImpl) BuscarTodos() ([]*domain.Vinculo, error) {
	var vinculos []*domain.Vinculo
	err := repo.db.Find(&vinculos).Error
	return vinculos, err
}

func (repo *VinculoRepositorioImpl) BuscarPeloID(id uint) (*domain.Vinculo, error) {
	var vinculo domain.Vinculo
	err := repo.db.First(&vinculo, id).Error
	if err != nil {
		return nil, err
	}
	return &vinculo, nil
}

func (repo *VinculoRepositorioImpl) Atualizar(v *domain.Vinculo) (*domain.Vinculo, error) {
	err := repo.db.Save(v).Error
	return v, err
}

func (repo *VinculoRepositorioImpl) Deletar(id uint) error {
	return repo.db.Delete(&domain.Vinculo{}, id).Error
}
