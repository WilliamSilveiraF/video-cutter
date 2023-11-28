package address

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "workflow-editor/db"
)

func InsertAddress(address Address) (int, error) {
    sqlQuery, err := db.ReadSQLFile("internal/address/sql/insert_address.sql")
    if err != nil {
        return 0, err
    }

    stmt, err := db.GetDB().Prepare(sqlQuery)
    if err != nil {
        return 0, err
    }
    defer stmt.Close()

    var addressID int
    err = stmt.QueryRow(address.UserID, address.Zip, address.Street, address.Unit, address.City, address.State).Scan(&addressID)
    if err != nil {
        return 0, err
    }

    return addressID, nil
}

func UpdateAddress(userID int, address Address) error {
    sqlQuery, err := db.ReadSQLFile("internal/address/sql/update_address.sql")
    if err != nil {
        return err
    }

    stmt, err := db.GetDB().Prepare(sqlQuery)
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(userID, address.Zip, address.Street, address.Unit, address.City, address.State)
    if err != nil {
        return err
    }

    return nil
}

func RetrieveAddress(userID int) (*Address, error) {
    var address Address
    
    sqlQuery, err := db.ReadSQLFile("internal/address/sql/retrieve_address.sql")
    if err != nil {
        return nil, err
    }

    row := db.GetDB().QueryRow(sqlQuery, userID)
    err = row.Scan(&address.ID, &address.UserID, &address.Zip, &address.Street, &address.Unit, &address.City, &address.State)
    if err != nil {
        return nil, err
    }

    return &address, nil
}


func CreateAddressHandler(c *gin.Context) {
    var newAddress Address

    if err := c.BindJSON(&newAddress); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }
    newAddress.UserID = userID.(int)

    addressID, err := InsertAddress(newAddress)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert address"})
        return
    }

    newAddress.ID = addressID
    c.JSON(http.StatusCreated, gin.H{"message": "Address created successfully", "address": newAddress})
}


func UpdateAddressHandler(c *gin.Context) {
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    var updatedAddress Address
    if err := c.BindJSON(&updatedAddress); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    if err := UpdateAddress(userID.(int), updatedAddress); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update address"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Address updated successfully"})
}


func GetCurrentAddressHandler(c *gin.Context) {
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    address, err := RetrieveAddress(userID.(int))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve address"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"address": address})
}
