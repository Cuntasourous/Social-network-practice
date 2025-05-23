package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Global variable
var DB *sql.DB

func InitDB() {
	// dbPath := filepath.Join("backend", "db", "Database", "socialdb.db")
	// log.Printf("Database path: %s", dbPath)

	dbPath := "backend/db/Database/socialdb.db"

	var err error
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	// Test database connection
	if err = DB.Ping(); err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// Verify table creation
	var tableName string
	err = DB.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='users'").Scan(&tableName)
	if err != nil {
		log.Fatal("Table was not created:", err)
	}

	log.Println("Database and table initialized successfully")
}
