package models

import (
	"time"
)

type PointRecordType string

const (
	Entry PointRecordType = "entry" //!< Entrada.
	Exit  PointRecordType = "exit"  //!< Saída.
)

type PointRecord struct {
	ID            int64           `json:"id"`             //!< ID do registro de ponto.
	ScholarshipID int64           `json:"scholarship_id"` //!< ID do bolsista dono do registro.
	RecordedAt    time.Time       `json:"recorded_at"`    //!< Momento do registro.
	Type          PointRecordType `json:"type"`           //!< Entrada ou saída.
}
