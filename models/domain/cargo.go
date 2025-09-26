package domain

import "gorm.io/gorm"

type Cargo struct {
	gorm.Model
	Nome string
}
