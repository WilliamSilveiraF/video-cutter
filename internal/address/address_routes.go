package address

import (
    "github.com/gin-gonic/gin"
)

func SetupAuthenticatedAddressRoutes(authenticatedGroup *gin.RouterGroup) {
    authenticatedGroup.POST("/create", CreateAddressHandler)
    authenticatedGroup.PATCH("/update", UpdateAddressHandler)
    authenticatedGroup.GET("/current", GetCurrentAddressHandler)
}
