package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lits-06/ecommerce/tokens"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		ClientToken := c.Request.Header.Get("token")
		if ClientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No authorization header provied"})
			c.Abort()
			return
		}

		claims, err := tokens.ValidateToken(ClientToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("password", claims.Password)
		c.Next()
	}
}