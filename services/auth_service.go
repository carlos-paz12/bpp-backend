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

type LoginResponse struct {
	Token string      `json:"token"`
	Role  string      `json:"role"`
	User  models.User `json:"user"`
}

func (AuthService) Login(username, password string) (*LoginResponse, error) {
	user, exists := repository.UserRepository{}.FindByUsername(username)
	if !exists || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New("invalid username or password.")
	}
	role := resolveRole(user.ID)

	claims := models.JwtCustomClaims{
		UserID: user.ID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString(jwtKey)

	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token: tokenSigned,
		Role:  role,
		User:  *user,
	}, nil
}

func resolveRole(uid int64) string {
	_, exists := repository.AdminRepository{}.FindByUserID(uid)
	if exists {
		return "admin"
	} else {
		_, exists := repository.ScholarshipRepository{}.FindByUserID(uid)
		if exists {
			return "scholarship"
		}
	}
	return "user"
}
