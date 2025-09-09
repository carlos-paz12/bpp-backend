package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
)

var registros = map[string][]string{}

func MarcarPonto(c *gin.Context) {
    username := c.GetString("username")
    horario := time.Now().Format("02/01/2006 15:04:05")
    registros[username] = append(registros[username], horario)
    c.JSON(http.StatusOK, gin.H{"message": "Ponto registrado!", "horario": horario})
}

func ListaPontos(c *gin.Context) {
    username := c.GetString("username")
    c.JSON(http.StatusOK, gin.H{"registros": registros[username]})
}
