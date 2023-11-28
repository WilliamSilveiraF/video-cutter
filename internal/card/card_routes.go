package card

import (
    "github.com/gin-gonic/gin"
)

func SetupAuthenticatedCardRoutes(authenticatedGroup *gin.RouterGroup) {
    authenticatedGroup.POST("/cards", CreateCardHandler)
    authenticatedGroup.GET("/cards", ListCardsHandler)
    authenticatedGroup.DELETE("/cards/:id", DeleteCardHandler)
}
