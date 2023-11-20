package internal

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskController struct {
	pool *pgxpool.Pool
}

type TaskControllerInterface interface {
	AddTaskToDB(description, priority, dueDate string) error
	ViewTasksFromDB() ([]Task, error)
	UpdateTaskInDB(id int, description, priority, dueDate string, completed bool) error
	DeleteTaskFromDB(id int) error
}

func NewTaskController(pool *pgxpool.Pool) *TaskController {
	return &TaskController{pool: pool}
}
func (tc *TaskController) AddTaskToDB(description, priority string, dueDate string) error {
	_, err := tc.pool.Exec(context.Background(), "INSERT INTO tasks (description, priority, due_date, completed) VALUES ($1, $2, $3, $4)", description, priority, dueDate, false)
	return err
}

func (tc *TaskController) ViewTasksFromDB() ([]Task, error) {
	rows, err := tc.pool.Query(context.Background(), "SELECT id, description, priority, to_char(due_date, 'YYYY-MM-DD'), completed FROM tasks")
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

func (tc *TaskController) UpdateTaskInDB(id int, description, priority string, dueDate string, completed bool) error {
	_, err := tc.pool.Exec(context.Background(), "UPDATE tasks SET description = $1, priority = $2, due_date = $3, completed = $4 WHERE id = $5", description, priority, dueDate, completed, id)
	return err
}

func (tc *TaskController) DeleteTaskFromDB(id int) error {
	_, err := tc.pool.Exec(context.Background(), "DELETE FROM tasks WHERE id = $1", id)
	return err
}
