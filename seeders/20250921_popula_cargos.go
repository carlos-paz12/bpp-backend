package seeders

import (
	"spe/models/domain"

	"gorm.io/gorm"
)

func PopulaCargos(db *gorm.DB) error {
	cargos := []domain.Cargo{
		{Nome: "Bolsista"},
		{Nome: "Técnico Administrativo"},
	}
	for _, c := range cargos {
		db.FirstOrCreate(&c, domain.Cargo{Nome: c.Nome})
	}
	return nil
}
