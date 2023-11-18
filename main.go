package main

import (
	"log"
	"os"
	"strings"
	"tracking/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.ConnectDatabase()

	router := gin.Default()

	proxyStr := os.Getenv("PROXIES")

	if proxyStr != ""{
		trustedProxies := strings.Split(proxyStr, ",")
		err = router.SetTrustedProxies(trustedProxies)

		if err != nil {
			log.Fatal("Error setting trusted proxies")
		}
	} else {
		log.Fatal("No trusted proxies set")
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	host := os.Getenv("HOST")
	err = router.Run(host)
	
	if err != nil {
		log.Fatal("Error running router")
	}
}