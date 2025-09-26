package domain

import "gorm.io/gorm"

type Setor struct {
	gorm.Model
	Nome string
}
