package auth

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}
