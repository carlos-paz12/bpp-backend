package dto

import "github.com/golang-jwt/jwt/v5"

// ErroDTO representa um erro retornado pela API
type ErroDTO struct {
	Mensagem string `json:"mensagem"`
	Codigo   int    `json:"codigo,omitempty"`
	Detalhes string `json:"detalhes,omitempty"`
}

type ReivindicacoesJWT struct {
	jwt.RegisteredClaims
	MembroID uint   `json:"membro_id"`
	Cargo    string `json:"cargo"`
}

// ReqAutenticacaoDTO representa as credenciais enviadas para autenticação.
type ReqAutenticacaoDTO struct {
	NomeUsuario string `json:"nome_usuario" binding:"required"`
	Senha       string `json:"senha" binding:"required"`
}

// ResAutenticacaoDTO representa a resposta ao autenticar com sucesso.
type ResAutenticacaoDTO struct {
	Token  string    `json:"token"`
	Cargo  string    `json:"cargo"`
	Membro MembroDTO `json:"membro"`
}
