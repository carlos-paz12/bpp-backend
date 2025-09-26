package seeders

import (
	"spe/models/domain"

	"os"
	"time"

	"gorm.io/gorm"
)

func CriaMembroRedesInicial(db *gorm.DB) error {
	redes := domain.Membro{
		NomeCompleto: os.Getenv("REDES_MEMBER_NAME"),
		NomeUsuario:  os.Getenv("REDES_MEMBER_USERNAME"),
		SenhaHash:    os.Getenv("REDES_MEMBER_PASSWORD"),
		Email:        os.Getenv("REDES_MEMBER_EMAIL"),
		Celular:      os.Getenv("REDES_MEMBER_PHONE"),
		Matricula:    os.Getenv("REDES_MEMBER_REGISTER"),
	}
	db.FirstOrCreate(&redes, domain.Membro{NomeUsuario: redes.NomeUsuario})

	var setor domain.Setor
	db.Where("nome = ?", "Gerência de Redes").First(&setor)

	var cargo domain.Cargo
	db.Where("nome = ?", "Técnico Administrativo").First(&cargo)

	vinculo := domain.Vinculo{
		MembroID:    redes.ID,
		SetorID:     setor.ID,
		CargoID:     cargo.ID,
		IniciadoEm:  time.Now(),
		Ativo:       true,
		Observacoes: "Membro administrador padrão",
	}
	db.FirstOrCreate(&vinculo, domain.Vinculo{
		MembroID: redes.ID,
		SetorID:  setor.ID,
		CargoID:  cargo.ID,
	})

	tecnicoAdmin := domain.TecnicoAdministrativo{
		MembroID: redes.ID,
	}
	db.FirstOrCreate(&tecnicoAdmin, domain.TecnicoAdministrativo{MembroID: redes.ID})

	return nil
}
