package domain

import "gorm.io/gorm"

type TecnicoAdministrativo struct {
	gorm.Model
	MembroID uint
	Membro   Membro `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
