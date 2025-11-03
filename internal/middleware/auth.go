package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(secret string, roles ...string) gin.HandlerFunc {
	roleSet := map[string]bool{}
	for _, r := range roles {
		roleSet[r] = true
	}
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if !strings.HasPrefix(h, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token requerido"})
			return
		}
		tok := strings.TrimPrefix(h, "Bearer ")
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(tok, claims, func(t *jwt.Token) (interface{}, error) { return []byte(secret), nil })
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}
		// opcional: validar rol
		if len(roleSet) > 0 {
			if role, _ := claims["Role"].(string); !roleSet[role] {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "permiso denegado"})
				return
			}
		}
		c.Set("claims", claims)
		c.Next()
	}
}
