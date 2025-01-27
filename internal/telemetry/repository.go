package telemetry

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db_connection *sql.DB) *Repository {
	return &Repository{db: db_connection}
}

func (repository *Repository) getHighestValuePerMetric() ([]Metric, error) {
	rows, err := repository.db.Query("SELECT metric, MAX(value) FROM telemetry GROUP BY metric")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var metrics []Metric

	for rows.Next() {
		var metric Metric
		err := rows.Scan(&metric.Metric, &metric.Value)
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, metric)
	}

	return metrics, nil
}

func (repository *Repository) getAll() ([]Metric, error) {
	rows, err := repository.db.Query("SELECT * FROM telemetry")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []Metric
	for rows.Next() {
		var entry Metric
		err := rows.Scan(&entry.ID, &entry.Metric, &entry.Value, &entry.CreatedAt)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

func (repository *Repository) create(entry Metric) error {
	_, err := repository.db.Exec("INSERT INTO telemetry (metric, value) VALUES ($1, $2)", entry.Metric, entry.Value)
	return err
}
