package audio

import (
    "github.com/gin-gonic/gin"
)

func SetupAuthenticatedAudioRoutes(authenticatedGroup *gin.RouterGroup) {
    authenticatedGroup.POST("/upload", UploadAudioHandler)
    authenticatedGroup.GET("/list", ListAudiosHandler)
    authenticatedGroup.DELETE("/delete/:id", DeleteAudioHandler)
    authenticatedGroup.GET("/download/:id", DownloadAudioHandler)

}
