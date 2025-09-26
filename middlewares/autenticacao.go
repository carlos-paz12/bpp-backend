package middlewares

import (
	"spe/models/dto"

	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var segredo = []byte(os.Getenv("API_SECRET"))

func RequerAutenticacao() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Tenta recuperar o cabeçalho 'Authorization' da requisição.
		headerAutorizacao := c.GetHeader("Authorization")
		if headerAutorizacao == "" {
			c.JSON(http.StatusBadRequest, dto.ErroDTO{
				Mensagem: "token não informado",
				Codigo:   http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		// Verifica se o cabeçalho veio no formato esperado.
		partes := strings.Split(headerAutorizacao, " ")
		if len(partes) != 2 || partes[0] != "Bearer" {
			c.JSON(http.StatusBadRequest, dto.ErroDTO{
				Mensagem: "token inválido",
				Codigo:   http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		// Verifica tipo de assinatura do token e injeta payload em `claims`.
		claims := &dto.ReivindicacoesJWT{}
		token, err := jwt.ParseWithClaims(partes[1], claims, func(token *jwt.Token) (any, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("token inválido")
			}
			return segredo, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusBadRequest, dto.ErroDTO{
				Mensagem: err.Error(),
				Codigo:   http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		// Injeta id do membro e cargo do membro no contexto da requisição.
		c.Set("membro_id", claims.MembroID)
		c.Set("cargo", claims.Cargo)

		// Continua a requisição.
		c.Next()
	}
}
