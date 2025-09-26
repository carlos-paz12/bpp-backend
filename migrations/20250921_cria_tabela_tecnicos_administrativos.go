package migrations

import (
	"spe/models/domain"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CriaTabelaTecnicosAdministrativos() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20250921_cria_tabela_tecnicos_administrativos",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&domain.TecnicoAdministrativo{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&domain.TecnicoAdministrativo{})
		},
	}
}
