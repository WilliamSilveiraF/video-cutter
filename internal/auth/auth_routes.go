package auth

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)



func SetupAuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", registerHandler)
	}
}



func registerHandler(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := RegisterUser(newUser.Email, newUser.Password)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}