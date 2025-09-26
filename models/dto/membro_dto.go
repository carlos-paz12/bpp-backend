package dto

type MembroDTO struct {
	MembroID     uint   `json:"membro_id,omitempty"`
	NomeCompleto string `json:"nome_completo,omitempty"`
	Matricula    string `json:"matricula,omitempty"`
	NomeUsuario  string `json:"nome_usuario,omitempty"`
	Email        string `json:"email,omitempty"`
	Celular      string `json:"celular,omitempty"`
}
