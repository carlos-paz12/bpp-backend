package dto

type TecnicoAdministrativoDTO struct {
	MembroID                uint   `json:"membro_id,omitempty"`
	TecnicoAdministrativoID uint   `json:"tecnico_administrativo_id,omitempty"`
	NomeCompleto            string `json:"nome_completo,omitempty"`
	Matricula               string `json:"matricula,omitempty"`
	NomeUsuario             string `json:"nome_usuario,omitempty"`
	Email                   string `json:"email,omitempty"`
	Celular                 string `json:"celular,omitempty"`
	SetorNome               string `json:"setor_nome,omitempty"`
}

type ReqCriacaoTecnicoAdministrativoDTO struct {
	NomeCompleto       string `json:"nome_completo" binding:"required"`
	Matricula          string `json:"matricula" binding:"required"`
	NomeUsuario        string `json:"nome_usuario" binding:"required"`
	Senha              string `json:"senha" binding:"required"`
	Email              string `json:"email" binding:"required"`
	Celular            string `json:"celular" binding:"required"`
	SetorNome          string `json:"setor_nome" binding:"required"`
}

type ReqAtualizacaoTecnicoAdministrativoDTO struct {
	NomeUsuario string `json:"nome_usuario,omitempty"`
	Email       string `json:"email,omitempty"`
	Celular     string `json:"celular,omitempty"`
}
