package models

import "time"

type StatusJustificativa string

const (
	Pendente  StatusJustificativa = "pendente"
	Aprovado  StatusJustificativa = "aprovado"
	Rejeitado StatusJustificativa = "rejeitado"
)

type Justificativa struct {
	Id         int                 `json:"id"`
	BolsistaID int                 `json:"bolsista_id"`
	Data       time.Time           `json:"data"`
	Horas      int                 `json:"horas"`
	Minutos    int                 `json:"minutos"`
	Motivo     string              `json:"motivo"`
	Status     StatusJustificativa `json:"status"`
	Comentario string              `json:"comentario"`
}
