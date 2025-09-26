package repositories

import (
	"spe/models/domain"

	"gorm.io/gorm"
)

type CargoRepositorio interface {
	Salvar(c *domain.Cargo) (*domain.Cargo, error)
	BuscarTodos() ([]*domain.Cargo, error)
	BuscarPeloID(id uint) (*domain.Cargo, error)
	BuscarPeloNome(nome string) (*domain.Cargo, error)
	Atualizar(c *domain.Cargo) (*domain.Cargo, error)
	Deletar(id uint) error
}

type CargoRepositorioImpl struct {
	db *gorm.DB
}

func NovoCargoRepositorio(db *gorm.DB) CargoRepositorio {
	return &CargoRepositorioImpl{db: db}
}

func (repo *CargoRepositorioImpl) Salvar(c *domain.Cargo) (*domain.Cargo, error) {
	err := repo.db.Create(c).Error
	return c, err
}

func (repo *CargoRepositorioImpl) BuscarTodos() ([]*domain.Cargo, error) {
	var cargos []*domain.Cargo
	err := repo.db.Find(&cargos).Error
	return cargos, err
}

func (repo *CargoRepositorioImpl) BuscarPeloID(id uint) (*domain.Cargo, error) {
	var cargo domain.Cargo
	err := repo.db.First(&cargo, id).Error
	if err != nil {
		return nil, err
	}
	return &cargo, nil
}

func (repo *CargoRepositorioImpl) BuscarPeloNome(nome string) (*domain.Cargo, error) {
	var cargo domain.Cargo
	err := repo.db.Where("nome = ?", nome).First(&cargo).Error
	if err != nil {
		return nil, err
	}
	return &cargo, nil
}

func (repo *CargoRepositorioImpl) Atualizar(b *domain.Cargo) (*domain.Cargo, error) {
	err := repo.db.Save(b).Error
	return b, err
}

func (repo *CargoRepositorioImpl) Deletar(id uint) error {
	return repo.db.Delete(&domain.Cargo{}, id).Error
}
