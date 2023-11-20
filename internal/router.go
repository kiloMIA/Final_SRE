package internal

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(pool *pgxpool.Pool) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	controller := NewTaskController(pool)
	handler := NewTaskHandler(controller)

	r.Post("/tasks", handler.AddTaskHandler(pool))
	r.Get("/tasks", handler.GetTasksHandler(pool))
	r.Put("/tasks/{taskID}", handler.UpdateTaskHandler(pool))
	r.Delete("/tasks/{taskID}", handler.DeleteTaskHandler(pool))

	return r
}
