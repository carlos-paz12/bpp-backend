package services

import (
	"spe/models"
	"spe/models/auth"

	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("7YcpY9oM56xRZ444ynlFS/khnm5LPCa/ktUgpPUzom0=")

var admins = map[string]*models.Admin{}
var bolsistas = map[string]*models.Bolsista{}

func init() {
	leonardo := &models.Admin{
		User: models.User{
			Id:       1,
			Name:     "Leonardo",
			Username: "leonardo",
			Email:    "leonardo@dimap.ufrn.br",
		},
	}
	leonardo.SetPassword("12345")
	admins[leonardo.Username] = leonardo

	leandro := &models.Bolsista{
		User: models.User{
			Id:       2,
			Name:     "Leandro",
			Username: "leandro",
			Email:    "leandro@dimap.ufrn.br",
		},
		Matricula:    "20240000000",
		HorasMensais: 80,
	}
	leandro.SetPassword("qwerty@12345")
	bolsistas[leandro.Username] = leandro
}

func Login(username, password string) (string, error) {
	var role string
	var userID int

	if admin, exists := admins[username]; exists && admin.CheckPassword(password) {
		userID = admin.Id
		role = "admin"
	} else if bolsista, exists := bolsistas[username]; exists && bolsista.CheckPassword(password) {
		userID = bolsista.Id
		role = "bolsista"
	} else {
		return "", errors.New("usuário ou senha inválidos.")
	}

	claims := &auth.Claims{
		UserId: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
