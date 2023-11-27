package user

import (
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", RegisterHandler)
		userGroup.POST("/login", LoginHandler)
	}
}

func SetupAuthenticatedUserRoutes(authenticatedUserGroup *gin.RouterGroup) {
	authenticatedUserGroup.POST("/update-password", UpdatePasswordHandler)
	authenticatedUserGroup.GET("/current", CurrentUserHandler)
}
