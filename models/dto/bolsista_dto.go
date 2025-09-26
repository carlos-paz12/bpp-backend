package dto

type BolsistaDTO struct {
	MembroID           uint   `json:"membro_id,omitempty"`
	BolsistaID         uint   `json:"bolsista_id,omitempty"`
	NomeCompleto       string `json:"nome_completo,omitempty"`
	Matricula          string `json:"matricula,omitempty"`
	NomeUsuario        string `json:"nome_usuario,omitempty"`
	Email              string `json:"email,omitempty"`
	Celular            string `json:"celular,omitempty"`
	SetorNome          string `json:"setor_nome,omitempty"`
	CargaHorariaMensal uint   `json:"ch_mensal,omitempty"`
}

type ReqCriacaoBolsistaDTO struct {
	NomeCompleto       string `json:"nome_completo" binding:"required"`
	Matricula          string `json:"matricula" binding:"required"`
	NomeUsuario        string `json:"nome_usuario" binding:"required"`
	Senha              string `json:"senha" binding:"required"`
	Email              string `json:"email" binding:"required"`
	Celular            string `json:"celular" binding:"required"`
	SetorNome          string `json:"setor_nome" binding:"required"`
	CargaHorariaMensal uint   `json:"ch_mensal" binding:"required"`
}

type ReqAtualizacaoBolsistaDTO struct {
	NomeUsuario string `json:"nome_usuario,omitempty"`
	Email       string `json:"email,omitempty"`
	Celular     string `json:"celular,omitempty"`
}
