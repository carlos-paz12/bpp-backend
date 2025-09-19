package models

// AuthRequest representa o corpo da requisição para autenticação de um usuário.
type AuthRequest struct {
	Username string `json:"username" binding:"required"` // Username do usuário que será autenticado.
	Password string `json:"password" binding:"required"` // Senha do usuário.
}
