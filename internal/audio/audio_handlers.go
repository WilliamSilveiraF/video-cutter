package audio

import (
    "net/http"
    "os"
    "context"

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
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the file"})
        return
    }

    ctx := context.Background()
    client, err := createSpeechClient(ctx)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create speech client"})
        return
    }
    defer client.Close()

    transcription, err := transcribeAudio(ctx, client, tempFilePath)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to transcribe the audio"})
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
        Transcription: transcription,
    }

    _, err = InsertAudio(audioRecord)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save audio data"})
        return
    }

    os.Remove(tempFilePath)

    c.JSON(http.StatusOK, gin.H{"message": "File uploaded and transcribed successfully", "transcription": transcription})
}
