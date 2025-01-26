package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	log.Println("Warming up the machine")

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Abusing HTTP 302 to grow my network 😎
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://github.com/bo0st3r", http.StatusFound)
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I'm fine sir!!"))
	})

	http.ListenAndServe(":8080", r)
	log.Println("She's ready!")
}
