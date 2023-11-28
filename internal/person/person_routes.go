package person

import (
	"github.com/gin-gonic/gin"
)


func SetupAuthenticatedPersonRoutes(authenticatedGroup *gin.RouterGroup) {
	authenticatedGroup.GET("/current", CurrentPersonHandler)
	authenticatedGroup.PUT("/update", UpdatePersonHandler)
}
