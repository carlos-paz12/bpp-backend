package models

import (
	"time"
)

type StatusJustification string

const (
	Pending  StatusJustification = "Pending"  //!< Pendente.
	Deferred StatusJustification = "Deferred" //!< Aprovada/Deferida.
	Rejected StatusJustification = "Rejected" //!< Rejeitada.
)

type Justification struct {
	ID              int64               `json:"id"`               //!< ID da justificativa.
	ScholarshipID   int64               `json:"scholarship_id"`   //!< ID do bolsista dono da justificativa.
	Date            time.Time           `json:"date"`             //!< Dia que está sendo justificado.
	Hours           int8                `json:"hours"`            //!< Horas justificadas.
	Minutes         int8                `json:"minutes"`          //!< Minutos justificados.
	Reason          string              `json:"reason"`           //!< Motivo (atestado, compromisso acadêmico, etc.).
	Status          StatusJustification `json:"status"`           //!< Estado da justificativa.
	ApproverID      int64               `json:"approver_id"`      //!< ID do admin que aprovou/rejeitou.
	ApprovedAt      time.Time           `json:"approved_at"`      //!< Data de aprovação/rejeição.
	ApprovalComment string              `json:"approval_comment"` //!< Comentário de quem aprova/rejeita.
}
