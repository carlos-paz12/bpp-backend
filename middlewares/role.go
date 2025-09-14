package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequireRole garante que o usuário tenha uma role específica.
func RequireRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		/*!<
		 * Busca a role no contexto e compara com a requerida `requiredRole`.
		 * Se a role for diferente ou não existir, returna 403 Forbidden para o cliente.
		 * Se for a role correta, deixa a requisição seguir.
		 */
		role, exists := c.Get("role")
		if !exists || role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{
				"error":  "Acesso negado.",
				"status": http.StatusForbidden,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
