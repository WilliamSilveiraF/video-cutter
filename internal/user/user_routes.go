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
