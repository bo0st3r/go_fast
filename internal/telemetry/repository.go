package telemetry

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db_connection *sql.DB) *Repository {
	return &Repository{db: db_connection}
}

func (repository *Repository) getAll() ([]Entry, error) {
	rows, err := repository.db.Query("SELECT * FROM telemetry")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []Entry
	for rows.Next() {
		var entry Entry
		err := rows.Scan(&entry.ID, &entry.Metric, &entry.Value, &entry.CreatedAt)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

func (repository *Repository) create(entry Entry) error {
	_, err := repository.db.Exec("INSERT INTO telemetry (metric, value) VALUES ($1, $2)", entry.Metric, entry.Value)
	return err
}
