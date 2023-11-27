package middleware

import (
	"net/http"
	"log"
	"strings"

	"workflow-editor/internal/user"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"

)

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		if strings.HasPrefix(tokenString, "Bearer ") {
            tokenString = strings.TrimPrefix(tokenString, "Bearer ")
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
            c.Abort()
            return
        }

		claims := &user.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return user.JwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		email := claims.Email

		user, err := user.RetrieveUser(email)
		
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
			c.Abort()
			return
		}
		
		c.Set("userID", user.ID)
		c.Set("user", user)
		c.Set("email", email)
		c.Next()
	}
}
