package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PontoController struct{}

var registros = map[string][]string{}

func (PontoController) Create(c *gin.Context) {
	username := c.GetString("username")
	horario := time.Now().Format("02/01/2006 15:04:05")
	registros[username] = append(registros[username], horario)
	c.JSON(http.StatusOK, gin.H{"message": "Ponto registrado!", "horario": horario})
}

func (PontoController) RetrieveAllWhereBolsistaId(c *gin.Context) {
	username := c.GetString("username")
	c.JSON(http.StatusOK, gin.H{"registros": registros[username]})
}

func (PontoController) RetrieveLastWhereBolsistaId(c *gin.Context) {
	// Todo
}
