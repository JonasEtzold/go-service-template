package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/JonasEtzold/go-service-template/pkg/crypto"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("authorization")
		if !crypto.ValidateToken(authorizationHeader) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		} else {
			c.Next()
		}
	}
}
