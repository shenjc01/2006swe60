package internal

import (
	"database/sql"
	"log"

	//_ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

// Initialize the database connection
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "./db/data.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}
