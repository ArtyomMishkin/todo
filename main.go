package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"example.com/todo/internal/task"
	myMW "example.com/todo/pkg/middleware"
)

func main() {
	storage := task.NewFileStorage("tasks.json")
	repo := task.NewRepo(storage)
	h := task.NewHandler(repo)

	r := chi.NewRouter()
	r.Use(chimw.RequestID)
	r.Use(chimw.Recoverer)
	r.Use(myMW.Logger)
	r.Use(myMW.SimpleCORS)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Версионирование API
	r.Route("/api/v1", func(api chi.Router) {
		api.Mount("/tasks", h.Routes())
	})

	// Старая версия для обратной совместимости (опционально)
	r.Route("/api", func(api chi.Router) {
		api.Mount("/tasks", h.Routes())
	})

	addr := ":8080"
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
