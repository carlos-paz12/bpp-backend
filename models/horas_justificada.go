package models

import "time"

type StatusJustificativa string

const (
	Pendente  StatusJustificativa = "pendente"
	Aprovado  StatusJustificativa = "aprovado"
	Rejeitado StatusJustificativa = "rejeitado"
)

type HorasJustificadas struct {
	Id         int                 `json:"id"`
	BolsistaID int                 `json:"bolsista_id"`
	Data       time.Time           `json:"data"`
	Motivo     string              `json:"motivo"`
	Horas      int                 `json:"horas"`
	Minutos    int                 `json:"minutos"`
	Status     StatusJustificativa `json:"status"`
}
