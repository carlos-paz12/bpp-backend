package models

import "user.go"
type Bolsista struct {
	user
	Matricula string `json:"matricula"`
	Horas     int64  `json:"horas"`
}