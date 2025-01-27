package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/bo0st3r/go-fast/internal/config"
	"github.com/bo0st3r/go-fast/internal/db"
	"github.com/bo0st3r/go-fast/internal/telemetry"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	log.Println("Warming up the machine")

	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Loaded .env file")

	config, err := config.Load()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}
	log.Println("Config loaded")

	db_connection, err := db.Connect(config.DatabaseDSN)
	if err != nil {
		log.Fatal("Error connecting to DB!", err)
	}
	defer db_connection.Close()
	log.Println("Connected to the SOURCE")

	telemetryRepository := telemetry.NewRepository(db_connection)
	telemetryService := telemetry.NewService(telemetryRepository)
	telemetryHandler := telemetry.NewHandler(telemetryService)

	r := setupRouter(telemetryHandler)

	log.Printf("Guys at %s are about to be ready to serve you", config.Port)
	http.ListenAndServe(":"+config.Port, r)
}

func setupRouter(telemetryHandler *telemetry.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Abusing HTTP 302 to grow my network 😎
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://github.com/bo0st3r", http.StatusFound)
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I'm fine sir!!"))
	})

	r.Get("/telemetry", telemetryHandler.GetAll)
	r.Post("/telemetry", telemetryHandler.Create)

	return r
}
