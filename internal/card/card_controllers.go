package card

import (
	"log"
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)


func CreateCardHandler(c *gin.Context) {
    var newCard Card
    if err := c.BindJSON(&newCard); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    newCard.UserID = userID.(int)
    cardID, err := InsertCard(newCard)
    if err != nil {
		log.Println(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert card"})
        return
    }

    newCard.ID = cardID
    c.JSON(http.StatusCreated, gin.H{"message": "Card created successfully", "card": newCard})
}

func ListCardsHandler(c *gin.Context) {
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    cards, err := RetrieveCardsByUserID(userID.(int))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cards"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"cards": cards})
}

func DeleteCardHandler(c *gin.Context) {
    cardIDStr := c.Param("id")
    cardID, err := strconv.Atoi(cardIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card ID"})
        return
    }

    if err := DeleteCardByID(cardID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete card"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Card deleted successfully"})
}
