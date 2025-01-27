package telemetry

import "time"

type Entry struct {
	ID        int       `json:"id"`
	Metric    string    `json:"metric"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}
