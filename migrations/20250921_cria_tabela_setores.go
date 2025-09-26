package migrations

import (
	"spe/models/domain"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CriaTabelaSetores() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20250921_cria_tabela_setores",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&domain.Setor{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&domain.Setor{})
		},
	}
}
