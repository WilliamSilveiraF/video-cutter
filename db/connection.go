package db

import _ "github.com/lib/pq"
import (
	"fmt"
	"log"
	"os"
    
)

var DB 

func ConnectDatabase() {
	dsn := fmt.Sprintf("postgres:// %s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	var err error
	

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
}