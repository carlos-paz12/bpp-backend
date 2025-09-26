package migrations

import (
	"spe/models/domain"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CriaTabelaBolsistas() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20250921_cria_tabela_bolsistas",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&domain.Bolsista{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&domain.Bolsista{})
		},
	}
}
