package models

import "time"

enum StatusJustificativa {
	PENDENTE = "pendente"
	APROVADO = "aprovado"
	REJEITADO = "rejeitado"
}

type HorasJustificadas struct {
	Id        int       `json:"id"`
	BolsistaID int       `json:"bolsista_id"`
	Data      time.Time `json:"data"`
	Motivo    string    `json:"motivo"`
	Horas     int       `json:"horas"`
	Status    StatusJustificativa `json:"status"`
}
