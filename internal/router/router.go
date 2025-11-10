package router

import (
	"net/http"
	"taskify/internal/handlers"
	"taskify/internal/middleware"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	// Middlewares globais
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	// Rotas públicas
	r.Post("/api/login", handlers.Login)

	// Rotas protegidas
	r.Route("/api/tasks", func(rt chi.Router) {
		rt.Use(middleware.AuthMiddleware)
		rt.Get("/", handlers.GetTasks)
		rt.Post("/", handlers.CreateTask)
		rt.Get("/{id}", handlers.GetTask)
		rt.Put("/{id}", handlers.UpdateTask)
		rt.Delete("/{id}", handlers.DeleteTask)
	})

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	return r
}
