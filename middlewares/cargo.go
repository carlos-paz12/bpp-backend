package middlewares

import (
	"spe/models/dto"

	"net/http"

	"github.com/gin-gonic/gin"
)

func RequerCargo(cargoRequerido string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Verifica se o cargo do membro foi injetado no contexto da requisição.
		cargo, existe := c.Get("cargo")
		if !existe || cargo != cargoRequerido {
			c.JSON(http.StatusForbidden, dto.ErroDTO{
				Mensagem: "acesso negado",
				Codigo:   http.StatusForbidden,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
