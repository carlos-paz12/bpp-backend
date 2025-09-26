package domain

import (
	"time"

	"gorm.io/gorm"
)

type Vinculo struct {
	gorm.Model
	MembroID    uint
	Membro      Membro `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	SetorID     uint
	Setor       Setor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CargoID     uint
	Cargo       Cargo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	IniciadoEm  time.Time
	Ativo       bool
	Observacoes string
}
