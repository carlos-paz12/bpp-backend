package migrations

import (
	"spe/models/domain"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CriaTabelaMembros() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20250921_cria_tabela_membros",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&domain.Membro{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&domain.Membro{})
		},
	}
}
