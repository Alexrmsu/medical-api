package db

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func Connect() {
	var err error
	connStr := "user=youruser password=yourpassword dbname=yourdb sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected")
}
