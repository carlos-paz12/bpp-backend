package migrations

import (
	"spe/models/domain"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CriaTabelaCargos() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20250921_cria_tabela_cargos",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&domain.Cargo{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&domain.Cargo{})
		},
	}
}
