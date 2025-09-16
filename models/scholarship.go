package models

type Scholarship struct {
	ID            int64  `json:"id"`              //!< ID do registro de bolsista.
	UserID        int64  `json:"user_id"`         //!< ID do bolsista dono da bolsa.
	Register      string `json:"register"`        //!< Matrícula do bolsista.
	HoursPerMonth int    `json:"hours_per_month"` //!< Carga horária mensal.
}
