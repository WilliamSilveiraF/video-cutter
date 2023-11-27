package person

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func CurrentPersonHandler(c *gin.Context) {
	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	person, err := RetrievePerson(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve person"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"person": person})
}