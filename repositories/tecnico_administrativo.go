package repositories

import (
	"spe/models/domain"

	"gorm.io/gorm"
)

type TecnicoAdministrativoRepositorio interface {
	Salvar(ta *domain.TecnicoAdministrativo) (*domain.TecnicoAdministrativo, error)
	BuscarTodos() ([]*domain.TecnicoAdministrativo, error)
	BuscarPeloID(id uint) (*domain.TecnicoAdministrativo, error)
	Atualizar(ta *domain.TecnicoAdministrativo) (*domain.TecnicoAdministrativo, error)
	Deletar(id uint) error
}

type TecnicoAdministrativoRepositorioImpl struct {
	db *gorm.DB
}

func NovoTecnicoAdministrativoRepositorio(db *gorm.DB) TecnicoAdministrativoRepositorio {
	return &TecnicoAdministrativoRepositorioImpl{db: db}
}

func (repo *TecnicoAdministrativoRepositorioImpl) Salvar(ta *domain.TecnicoAdministrativo) (*domain.TecnicoAdministrativo, error) {
	err := repo.db.Create(ta).Error
	return ta, err
}

func (repo *TecnicoAdministrativoRepositorioImpl) BuscarTodos() ([]*domain.TecnicoAdministrativo, error) {
	var tecnicosAdministrativos []*domain.TecnicoAdministrativo
	err := repo.db.Find(&tecnicosAdministrativos).Error
	return tecnicosAdministrativos, err
}

func (repo *TecnicoAdministrativoRepositorioImpl) BuscarPeloID(id uint) (*domain.TecnicoAdministrativo, error) {
	var tecnicoAdministrativo domain.TecnicoAdministrativo
	err := repo.db.First(&tecnicoAdministrativo, id).Error
	if err != nil {
		return nil, err
	}
	return &tecnicoAdministrativo, nil
}

func (repo *TecnicoAdministrativoRepositorioImpl) Atualizar(ta *domain.TecnicoAdministrativo) (*domain.TecnicoAdministrativo, error) {
	err := repo.db.Save(ta).Error
	return ta, err
}

func (repo *TecnicoAdministrativoRepositorioImpl) Deletar(id uint) error {
	return repo.db.Delete(&domain.TecnicoAdministrativo{}, id).Error
}
