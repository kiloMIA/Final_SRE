package internal

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/go-chi/chi/v5"
    "github.com/jackc/pgx/v5/pgxpool"
)

func AddTaskHandler(pool *pgxpool.Pool) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var task Task
        if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        err := AddTaskToDB(pool, task.Description, task.Priority, task.DueDate)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(task)
    }
}

func GetTasksHandler(pool *pgxpool.Pool) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tasks, err := ViewTasksFromDB(pool)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(tasks)
    }
}

func UpdateTaskHandler(pool *pgxpool.Pool) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        taskID, err := strconv.Atoi(chi.URLParam(r, "taskID"))
        if err != nil {
            http.Error(w, "Invalid task ID", http.StatusBadRequest)
            return
        }

        var task Task
        if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        err = UpdateTaskInDB(pool, taskID, task.Description, task.Priority, task.DueDate, task.Completed)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
    }
}

func DeleteTaskHandler(pool *pgxpool.Pool) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        taskID, err := strconv.Atoi(chi.URLParam(r, "taskID"))
        if err != nil {
            http.Error(w, "Invalid task ID", http.StatusBadRequest)
            return
        }

        err = DeleteTaskFromDB(pool, taskID)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
    }
}
