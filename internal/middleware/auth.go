package middleware

import (
	"log"
	"net/http"
	"strings"

	utils "github.com/eliferdentr/finance-tracker-app/internal/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(secret string) gin.HandlerFunc {

	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header!"})
			return
		}
		// “Bearer ” ile başlıyor mu? → hayırsa 401
		if !strings.HasPrefix(authorizationHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}

		// Token’ı al → doğruysa 200, yanlışsa yine 401
		token := strings.TrimPrefix(authorizationHeader, "Bearer ")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		claims, err := utils.ValidateToken(token, secret)
		if err != nil {
			log.Println("token validation failed:", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
		return 
		}
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
