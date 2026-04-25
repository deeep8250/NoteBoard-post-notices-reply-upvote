package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Miiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		HeaderValue := c.GetHeader("Authorization")
		if HeaderValue == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token is empty",
			})
			c.Abort()
			return
		}

		parts := strings.Split(HeaderValue, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token format",
			})
			c.Abort()
			return
		}

		tokenstring := parts[1]

		token, err := jwt.Parse(tokenstring, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invlaid signature",
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid claims in token",
			})
			c.Abort()
			return
		}

		userID := int(claims["userID"].(float64))
		c.Set("userID", userID)
		c.Next()

	}
}
