package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func InitDB() {

	//dbFile := "events.db"
	/*if err := os.Remove(dbFile); err != nil && !os.IsNotExist(err) {
		log.Fatalf("Failed to delete existing database: %v", err)
	}*/
	var err error
	DB, err = sql.Open("sqlite3", "events.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Set database connection properties
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// Create necessary tables
	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    email TEXT NOT NULL UNIQUE,
	    password TEXT NOT NULL
)`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    name TEXT NOT NULL,
	    description TEXT NOT NULL,
	    date_time TEXT NOT NULL,  -- Store as ISO 8601 string (YYYY-MM-DD HH:MM:SS)
	    user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}
