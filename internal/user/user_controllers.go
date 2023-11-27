package user

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"

	"workflow-editor/internal/person"
)

func RegisterHandler(c *gin.Context) {
    var request RegisterUserRequest

    if err := c.BindJSON(&request); err != nil {
		log.Println(err);
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    userID, err := RegisterUser(request.User.Email, request.User.Password)
    if err != nil {
        log.Println(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
        return
    }

    request.Person.UserID = userID

    err = person.InsertPerson(request.Person)
    if err != nil {
        log.Println(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save person details"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}


func LoginHandler(c *gin.Context) {
	var loginUser User

	if err := c.BindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	authSuccess, err := LoginUser(loginUser.Email, loginUser.Password)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Login failed"})
		return
	}

	if !authSuccess {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := GenerateJWT(loginUser.Email)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}


func UpdatePasswordHandler(c *gin.Context) {
	var updatePasswordRequest UpdatePasswordRequest

	if err := c.BindJSON(&updatePasswordRequest); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := UpdateUserPassword(updatePasswordRequest.Email, updatePasswordRequest.OldPassword, updatePasswordRequest.NewPassword)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}


func CurrentUserHandler(c *gin.Context) {
	userObj, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, ok := userObj.(*User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	user.Password = ""

	c.JSON(http.StatusOK, gin.H{"user": user})
}