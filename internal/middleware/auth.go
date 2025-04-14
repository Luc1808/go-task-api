package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var accessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")

// authMiddleware validaes access tokens
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required!"})
			ctx.Abort()
			return
		}

		// Extract token from Bearer prefix
		tokenString := ""
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			ctx.Abort()
			return
		}

		// Parse refresh token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(accessTokenSecret), nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token", "err": err})
			ctx.Abort()
			return
		}

		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse token claims"})
			ctx.Abort()
			return
		}

		userID := uint(claims["user_id"].(float64))
		ctx.Set("userID", userID)
		ctx.Next()
	}
}
