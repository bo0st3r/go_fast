package telemetry

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (handler *Handler) GetHighestValuePerMetric(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	metrics, err := handler.service.GetHighestValuePerMetric()
	if err != nil {
		http.Error(w, "failed to get telemetry", http.StatusInternalServerError)
		return
	}

	if len(metrics) == 0 {
		metrics = []Metric{}
	}

	if err := json.NewEncoder(w).Encode(metrics); err != nil {
		http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}

func (handler *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	entries, err := handler.service.GetAll()
	if err != nil {
		http.Error(w, "failed to get telemetry", http.StatusInternalServerError)
		return
	}

	if len(entries) == 0 {
		entries = []Metric{}
	}

	if err := json.NewEncoder(w).Encode(entries); err != nil {
		http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var payload Metric

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if payload.Metric == "" || payload.Value == 0 {
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return
	}

	if err := handler.service.Create(Metric{Metric: payload.Metric, Value: payload.Value}); err != nil {
		http.Error(w, "failed to record telemetry", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
