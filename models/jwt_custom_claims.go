package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	jwt.RegisteredClaims        //!< Claims padrão do JWT (exp, iat, etc.).
	UserID               int64  `json:"user_id"` //!< ID do usuário.
	Role                 string `json:"role"`    //!< Papel do usuário.
}
