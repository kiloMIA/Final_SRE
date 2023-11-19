package internal

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(pool *pgxpool.Pool) *chi.Mux {
    r := chi.NewRouter()
	r.Use(middleware.Logger)
	
    r.Post("/tasks", AddTaskHandler(pool))
    r.Get("/tasks", GetTasksHandler(pool))
    r.Put("/tasks/{taskID}", UpdateTaskHandler(pool))
    r.Delete("/tasks/{taskID}", DeleteTaskHandler(pool))

    return r
}
