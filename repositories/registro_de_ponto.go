package repositories

import (
	"spe/models/domain"

	"gorm.io/gorm"
)

type RegistroDePontoRepositorio interface {
	Salvar(rp *domain.RegistroDePonto) (*domain.RegistroDePonto, error)
	BuscarTodos() ([]*domain.RegistroDePonto, error)
	BuscarPeloID(id uint) (*domain.RegistroDePonto, error)
	Atualizar(rp *domain.RegistroDePonto) (*domain.RegistroDePonto, error)
	Deletar(id uint) error
}

type RegistroDePontoRepositorioImpl struct {
	db *gorm.DB
}

func NovoRegistroDePontoRepositorio(db *gorm.DB) RegistroDePontoRepositorio {
	return &RegistroDePontoRepositorioImpl{db: db}
}

func (repo *RegistroDePontoRepositorioImpl) Salvar(rp *domain.RegistroDePonto) (*domain.RegistroDePonto, error) {
	err := repo.db.Create(rp).Error
	return rp, err
}

func (repo *RegistroDePontoRepositorioImpl) BuscarTodos() ([]*domain.RegistroDePonto, error) {
	var registrosDePonto []*domain.RegistroDePonto
	err := repo.db.Find(&registrosDePonto).Error
	return registrosDePonto, err
}

func (repo *RegistroDePontoRepositorioImpl) BuscarPeloID(id uint) (*domain.RegistroDePonto, error) {
	var registroDePonto domain.RegistroDePonto
	err := repo.db.First(&registroDePonto, id).Error
	if err != nil {
		return nil, err
	}
	return &registroDePonto, nil
}

func (repo *RegistroDePontoRepositorioImpl) Atualizar(rp *domain.RegistroDePonto) (*domain.RegistroDePonto, error) {
	err := repo.db.Save(rp).Error
	return rp, err
}

func (repo *RegistroDePontoRepositorioImpl) Deletar(id uint) error {
	return repo.db.Delete(&domain.RegistroDePonto{}, id).Error
}
