package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func initDatabase() {
	var err error
	db, err = sql.Open("sqlite", "./expenses.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS expenses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		amount REAL NOT NULL,
		category TEXT NOT NULL,
		date TEXT NOT NULL
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	log.Println("Database initialized successfully")
}
