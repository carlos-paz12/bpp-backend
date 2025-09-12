package models

type Bolsista struct {
	User
	Matricula    string `json:"matricula"`
	HorasMensais int64  `json:"horas_mensais"`
}
