package repositories

import (
	"spe/models/domain"

	"gorm.io/gorm"
)

type MembroRepositorio interface {
	Salvar(m *domain.Membro) (*domain.Membro, error)
	BuscarTodos() ([]*domain.Membro, error)
	BuscarPeloID(id uint) (*domain.Membro, error)
	BuscarPeloNomeUsuario(nomeUsuario string) (*domain.Membro, error)
	BuscarCargosPeloMembroID(id uint) ([]domain.Cargo, error)
	Atualizar(m *domain.Membro) (*domain.Membro, error)
	Deletar(id uint) error
}

type MembroRepositorioImpl struct {
	db *gorm.DB
}

func NovoMembroRepositorio(db *gorm.DB) MembroRepositorio {
	return &MembroRepositorioImpl{db: db}
}

func (repo *MembroRepositorioImpl) Salvar(m *domain.Membro) (*domain.Membro, error) {
	err := repo.db.Create(m).Error
	return m, err
}

func (repo *MembroRepositorioImpl) BuscarTodos() ([]*domain.Membro, error) {
	var membros []*domain.Membro
	err := repo.db.Find(&membros).Error
	return membros, err
}

func (repo *MembroRepositorioImpl) BuscarPeloID(id uint) (*domain.Membro, error) {
	var membro domain.Membro
	err := repo.db.First(&membro, id).Error
	if err != nil {
		return nil, err
	}
	return &membro, nil
}

func (repo *MembroRepositorioImpl) BuscarPeloNomeUsuario(nomeUsuario string) (*domain.Membro, error) {
	var membro domain.Membro
	err := repo.db.Where("nome_usuario = ?", nomeUsuario).First(&membro).Error
	if err != nil {
		return nil, err
	}
	return &membro, nil
}

func (repo *MembroRepositorioImpl) BuscarCargosPeloMembroID(id uint) ([]domain.Cargo, error) {
	var vinculos []domain.Vinculo

	err := repo.db.Preload("Cargo").Where("membro_id = ? AND ativo = ?", id, true).Find(&vinculos).Error
	if err != nil {
		return nil, err
	}

	cargos := make([]domain.Cargo, len(vinculos))
	for i, v := range vinculos {
		cargos[i] = v.Cargo
	}

	return cargos, nil
}

func (repo *MembroRepositorioImpl) Atualizar(m *domain.Membro) (*domain.Membro, error) {
	err := repo.db.Save(m).Error
	return m, err
}

func (repo *MembroRepositorioImpl) Deletar(id uint) error {
	return repo.db.Delete(&domain.Membro{}, id).Error
}
