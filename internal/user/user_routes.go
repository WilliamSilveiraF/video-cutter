package user

import (
	"github.com/gin-gonic/gin"
)


func SetupUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", RegisterHandler)
		userGroup.POST("/login", LoginHandler)

		authenticatedUserGroup := userGroup.Group("/").Use(UserMiddleware())
		{
			authenticatedUserGroup.POST("/update-password", UpdatePasswordHandler)
			authenticatedUserGroup.GET("/current", CurrentUserHandler)
		}
	}
}
