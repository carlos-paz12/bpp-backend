package models

import "time"

type Ponto struct {
	Id         int       `json:"id"`
	BolsistaID int       `json:"bolsista_id"`
	Timestamp  time.Time `json:"timestamp"`
	Tipo       bool      `json:"tipo"`
}
