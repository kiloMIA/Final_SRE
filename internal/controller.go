package internal

import (
    "context"
    "github.com/jackc/pgx/v5/pgxpool"
)

func AddTaskToDB(pool *pgxpool.Pool, description, priority string, dueDate string) error {
    _, err := pool.Exec(context.Background(), "INSERT INTO tasks (description, priority, due_date, completed) VALUES ($1, $2, $3, $4)", description, priority, dueDate, false)
    return err
}

func ViewTasksFromDB(pool *pgxpool.Pool) ([]Task, error) {
    rows, err := pool.Query(context.Background(), "SELECT id, description, priority, to_char(due_date, 'YYYY-MM-DD'), completed FROM tasks")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tasks []Task
    for rows.Next() {
        var task Task
        err := rows.Scan(&task.ID, &task.Description, &task.Priority, &task.DueDate, &task.Completed)
        if err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }
    return tasks, nil
}



func UpdateTaskInDB(pool *pgxpool.Pool, id int, description, priority string, dueDate string, completed bool) error {
    _, err := pool.Exec(context.Background(), "UPDATE tasks SET description = $1, priority = $2, due_date = $3, completed = $4 WHERE id = $5", description, priority, dueDate, completed, id)
    return err
}

func DeleteTaskFromDB(pool *pgxpool.Pool, id int) error {
    _, err := pool.Exec(context.Background(), "DELETE FROM tasks WHERE id = $1", id)
    return err
}
