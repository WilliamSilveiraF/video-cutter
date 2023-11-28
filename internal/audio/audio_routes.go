package audio

import (
    "github.com/gin-gonic/gin"
)

func SetupAuthenticatedAddressRoutes(authenticatedGroup *gin.RouterGroup) {
    authenticatedGroup.POST("/upload", UploadAudioHandler)
}
