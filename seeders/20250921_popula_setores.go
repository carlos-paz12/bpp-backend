package seeders

import (
	"spe/models/domain"

	"gorm.io/gorm"
)

func PopulaSetores(db *gorm.DB) error {
	setores := []domain.Setor{
		{Nome: "Gerência de Redes"},
		{Nome: "Secretária"},
		{Nome: "Secretária de Pós-graduação"},
		{Nome: "Coordenação de Ciência da Computação"},
	}
	for _, s := range setores {
		db.FirstOrCreate(&s, domain.Cargo{Nome: s.Nome})
	}
	return nil
}
