package db

import (
	"io/ioutil"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to open database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Successfully connected to database")
}


func SetupTables() {
	paths := []string{
		"internal/user/sql/create_table.sql",
		"internal/person/sql/create_table.sql",
	}

	for _, path := range paths {
		sqlQuery, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalf("SetupTables: failed to read %s (%s)", path, err)
			return
		}

		_, err = db.Exec(string(sqlQuery))
		if err != nil {
			log.Fatalf("SetupTables: failed to execute SQL from %s (%v)", path, err)
			return
		}

		log.Printf("Successfully executed SQL from %s\n", path)
	}
}

func GetDB() *sql.DB {
	return db
}

