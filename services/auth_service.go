package services

import (
	"errors"
	"spe/models"
	"spe/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("7YcpY9oM56xRZ444ynlFS/khnm5LPCa/ktUgpPUzom0=")

type AuthService struct{}

func (AuthService) Login(username, password string) (*models.LoginResponse, error) {
	user, err := repository.UserRepository{}.FindByUsername(username)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New("Usuário ou senha inválidos")
	}

	role := "user"
	adminRepo := repository.AdminRepository{}
	if _, err := adminRepo.FindByUserID(user.ID); err == nil {
		role = "admin"
	} else {
		scholarRepo := repository.ScholarshipRepository{}
		if _, err := scholarRepo.FindByUserID(user.ID); err == nil {
			role = "bolsista"
		}
	}

	claims := models.JwtCustomClaims{
		Role:   role,
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, errors.New("Erro ao gerar token.")
	}

	return &models.LoginResponse{
		Token: tokenString,
		Role:  role,
		User:  *user,
	}, nil
}
