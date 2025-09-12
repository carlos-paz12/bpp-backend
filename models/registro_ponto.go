package models

import "time"

type RegistroPonto struct {
	Id        int       `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Tipo      bool      `json:"tipo"` // true para "entrada" ou false para "saida"
}
