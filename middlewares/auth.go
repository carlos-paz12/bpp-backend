package middlewares

import (
	"errors"
	"spe/models"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("7YcpY9oM56xRZ444ynlFS/khnm5LPCa/ktUgpPUzom0=")

// AuthMiddleware valida o token JWT e injeta claims no contexto.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*!<
		 * Busca o header "Authorization".
		 * Espera algo como: Authorization: Bearer <token>.
		 * Se estiver vazio, retorna 401 Unauthorized para o cliente.
		 */
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusBadRequest, models.ApiResponse{
				Message:  "",
				Error:    "token not provided.",
				HttpCode: http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		/*!<
		 * Divide o header em duas partes.
		 * Espera algo como: ["Bearer", "<token>"].
		 * Se não tiver esse formato, retorna 401 Unauthorized para o cliente.
		 */
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusBadRequest, models.ApiResponse{
				Message:  "",
				Error:    "invalid token.",
				HttpCode: http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		tokenSigned := parts[1]             //!< Token JWT enviado na requisição.
		claims := &models.JwtCustomClaims{} //!< Claims personalizadas.

		/*!<
		 * Valida o token com as claims personalizadas.
		 * Garante que o algoritmo de assinatura é HMAC (ex: HS256).
		 * Usa jwtKey para validar a assinatura.
		 * Decodifica o payload no struct `claims`.
		 */
		token, err := jwt.ParseWithClaims(tokenSigned, claims, func(token *jwt.Token) (any, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("invalid token.")
			}
			return jwtKey, nil
		})

		/*!<
		 * Se o token for inválido...
		 * Retorna 401 Unauthorized para o cliente.
		 */
		if err != nil || !token.Valid {
			c.JSON(http.StatusBadRequest, models.ApiResponse{
				Message:  "",
				Error:    err.Error(),
				HttpCode: http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		/*!<
		 * Se o token for válido...
		 * Injeta o ID e a role do usuário no contexto e continua.
		 */
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)
		c.Next()
	}
}
