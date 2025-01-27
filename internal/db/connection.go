package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Connect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	err = createTables(db)
	if err != nil {
		log.Fatalf("Could not create the table(s): %s", err)
		return nil, err
	}

	return db, nil
}

// Function to create necessary tables
func createTables(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS telemetry (
		id SERIAL PRIMARY KEY,
		metric VARCHAR(255) NOT NULL,
		value DOUBLE PRECISION NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(query)
	return err
}
