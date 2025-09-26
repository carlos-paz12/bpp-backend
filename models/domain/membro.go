package domain

import "gorm.io/gorm"

type Membro struct {
	gorm.Model
	NomeCompleto string `gorm:"not null"`
	Matricula    string `gorm:"unique;not null"`
	NomeUsuario  string `gorm:"unique;not null"`
	SenhaHash    string `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
	Celular      string `gorm:"unique;not null"`
}
