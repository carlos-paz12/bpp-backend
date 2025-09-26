package domain

import "gorm.io/gorm"

type Bolsista struct {
	gorm.Model
	CargaHorariaMensal uint
	MembroID           uint
	Membro             Membro `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
