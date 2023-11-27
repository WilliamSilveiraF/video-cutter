package main

import (
	"log"
	"os"
	
	"workflow-editor/db"
	"workflow-editor/middleware"
	"workflow-editor/internal/user"
	"workflow-editor/internal/person"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	dir, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Current directory:", dir)

	err = godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.ConnectDatabase()
	db.SetupTables()

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	
	user.SetupUserRoutes(router)
	authenticatedUserGroup := router.Group("/user").Use(middleware.UserMiddleware())
	if authGroup, ok := authenticatedUserGroup.(*gin.RouterGroup); ok {
		user.SetupAuthenticatedUserRoutes(authGroup)
	} else {
		log.Fatal("Failed to assert type of authenticatedUserGroup")
	}

	authenticatedPersonGroup := router.Group("/person").Use(middleware.UserMiddleware())
	if personGroup, ok := authenticatedPersonGroup.(*gin.RouterGroup); ok {
		person.SetupAuthenticatedPersonRoutes(personGroup)
	} else {
		log.Fatal("Failed to assert type of authenticated person group")
	}
	

	host := os.Getenv("HOST")
	err = router.Run(host)
	
	if err != nil {
		log.Fatal("Error running router")
	}
}