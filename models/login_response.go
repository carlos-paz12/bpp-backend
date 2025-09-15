package models

import (
	"github.com/golang-jwt/jwt/v5"
)

// JwtCustomClaims define os claims personalizados para JWT.
// Contém informações adicionais sobre o usuário, como ID e papel (role).
type JwtCustomClaims struct {
	jwt.RegisteredClaims        //!< Claims padrão do JWT (exp, iat, etc.).
	UserID               int64  `json:"user_id"` //!< ID do usuário.
	Role                 string `json:"role"`    //!< Papel do usuário.
}

// LoginResponse representa a resposta de login da API.
type LoginResponse struct {
	Token string `json:"token"` //!< JWT gerado para autenticação.
	Role  string `json:"role"`  //!< Papel do usuário retornado no login.
	User  User   `json:"user"`  //!< Informações do usuário.
}
