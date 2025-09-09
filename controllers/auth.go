package controllers

import (
    "net/http"
    "time"
    "spe/models"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("minha_chave_secreta")

var bolsistas = map[string]*models.Bolsista{}

func init() {
    user1 := &models.Bolsista{Username: "carlos"}
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

    u := &models.Bolsista{Username: body.Username}
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

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": body.Username,
        "exp":      time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString(jwtKey)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
