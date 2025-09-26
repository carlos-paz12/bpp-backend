package domain

import (
	"time"

	"gorm.io/gorm"
)

type TipoDeRegistro string

const (
	Entrada TipoDeRegistro = "entrada"
	Saida   TipoDeRegistro = "saida"
)

type RegistroDePonto struct {
	gorm.Model
	RegistradoEm         time.Time
	TipoDeRegistro       TipoDeRegistro
	EhRegistroAutomatico bool
	BolsistaID           uint
	Bolsista             Bolsista `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
