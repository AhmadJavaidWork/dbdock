package db

import (
	"database/sql"
	"log"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init(appDir string) {
	dbPath := filepath.Join(appDir, "app.db")

	var err error
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	createTables()
}

func createTables() {
	query := `
		CREATE TABLE IF NOT EXISTS database_drivers (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL UNIQUE,
			label TEXT NOT NULL,
			default_port INTEGER NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP
		)
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	query = `
    CREATE TABLE IF NOT EXISTS connections (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			database_driver_id INTEGER,
			name TEXT NOT NULL,
			host TEXT NOT NULL,
			port INTEGER NOT NULL,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			database TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP,
			FOREIGN KEY(database_driver_id) REFERENCES database_drivers(id)
		);
	`
	_, err = DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	seedDrivers()
}

func seedDrivers() {
	query := `
    INSERT OR IGNORE INTO database_drivers (name, label, default_port)
    VALUES
			('postgres', 'PostgreSQL', 5432),
			('mysql', 'MySQL', 3306)
    `
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
