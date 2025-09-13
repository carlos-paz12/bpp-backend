package auth

import (
	"spe/models"
	"spe/models/auth"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("7YcpY9oM56xRZ444ynlFS/khnm5LPCa/ktUgpPUzom0=")

var bolsistas = map[string]*models.User{}

func init() {
	user1 := &models.User{ID: 1, Username: "carlos"}
	user1.SetPassword("1234")
	bolsistas[user1.Username] = user1
}

func Register(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	if _, exists := bolsistas[body.Username]; exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Usuário já existe"})
		return
	}

	u := &models.User{Username: body.Username}
	u.SetPassword(body.Password)
	bolsistas[body.Username] = u

	c.JSON(http.StatusCreated, gin.H{"message": "Usuário criado com sucesso"})
}

func Login(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	user, exists := bolsistas[body.Username]
	if !exists || !user.CheckPassword(body.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha inválidos"})
		return
	}

	claims := &auth.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 8)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// Para logout, no lugar de termos um método do tipo `func Logout(c *gin.Context)` e
// criar um endpoint dedicado para tratar o logout servidor, é mais prática (já que
// estamos usando JWT) simplesmente apagar o token de acesso armazenado no frontend
// e redirecionar o usuário para o endpoint de login.
