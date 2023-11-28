package audio

import (
    "net/http"
    "log"
    "strconv"
    "github.com/gin-gonic/gin"
)

func UploadAudioHandler(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
        return
    }

    tempFilePath := "temp/" + file.Filename
    if err := c.SaveUploadedFile(file, tempFilePath); err != nil {
        log.Println(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the file"})
        return
    }

    userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

    audioRecord := Audio{
        UserID:        userID.(int),
        Filename:      file.Filename,
    }

    _, err = InsertAudio(audioRecord)
    if err != nil {
        log.Println(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save audio data"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "File uploaded and transcribed successfully", })
}

func ListAudiosHandler(c *gin.Context) {
    userID, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

    audios, err := RetrieveAudiosByUserID(userID.(int))
    if err != nil {
        log.Println(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve audios"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"audios": audios})
}

func DeleteAudioHandler(c *gin.Context) {
    audioIDStr := c.Param("id")
    audioID, err := strconv.Atoi(audioIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid audio ID"})
        return
    }

    err = DeleteAudioByID(audioID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete audio"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Audio deleted successfully"})
}

func DownloadAudioHandler(c *gin.Context) {
    log.Println("BAteUUU")
    audioIDStr := c.Param("id")
    audioID, err := strconv.Atoi(audioIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid audio ID"})
        return
    }
    log.Println("TESTE 1")
    audio, err := RetrieveAudioByID(audioID) // Implement this function
    if err != nil {
        log.Println(err)
        c.JSON(http.StatusNotFound, gin.H{"error": "Audio not found"})
        return
    }
    log.Println(audio.FilePath)
    c.File(audio.FilePath) // Send the file
}
