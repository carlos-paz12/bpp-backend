package repositories

import (
	"spe/models/domain"

	"gorm.io/gorm"
)

type BolsistaRepositorio interface {
	Salvar(b *domain.Bolsista) (*domain.Bolsista, error)
	BuscarTodos() ([]*domain.Bolsista, error)
	BuscarPeloID(id uint) (*domain.Bolsista, error)
	Atualizar(b *domain.Bolsista) (*domain.Bolsista, error)
	Deletar(id uint) error
}

type BolsistaRepositorioImpl struct {
	db *gorm.DB
}

func NovoBolsistaRepositorio(db *gorm.DB) BolsistaRepositorio {
	return &BolsistaRepositorioImpl{db: db}
}

func (repo *BolsistaRepositorioImpl) Salvar(b *domain.Bolsista) (*domain.Bolsista, error) {
	err := repo.db.Create(b).Error
	return b, err
}

func (repo *BolsistaRepositorioImpl) BuscarTodos() ([]*domain.Bolsista, error) {
	var bolsistas []*domain.Bolsista
	err := repo.db.Find(&bolsistas).Error
	return bolsistas, err
}

func (repo *BolsistaRepositorioImpl) BuscarPeloID(id uint) (*domain.Bolsista, error) {
	var bolsista domain.Bolsista
	err := repo.db.First(&bolsista, id).Error
	if err != nil {
		return nil, err
	}
	return &bolsista, nil
}

func (repo *BolsistaRepositorioImpl) Atualizar(b *domain.Bolsista) (*domain.Bolsista, error) {
	err := repo.db.Save(b).Error
	return b, err
}

func (repo *BolsistaRepositorioImpl) Deletar(id uint) error {
	return repo.db.Delete(&domain.Bolsista{}, id).Error
}
