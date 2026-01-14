package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not get user home directory: %v", err)
	}

	dbPath := filepath.Join(homeDir, "dream-journal.db")
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return fmt.Errorf("could not open database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return fmt.Errorf("could not ping database: %v", err)
	}

	DB = db
	return createTables()
}

func createTables() error {
	createDreamsTable := `
	CREATE TABLE IF NOT EXISTS dreams (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	if _, err := DB.Exec(createDreamsTable); err != nil {
		return fmt.Errorf("could not create dreams table: %v", err)
	}

	return nil
}
