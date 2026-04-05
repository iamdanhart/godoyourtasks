package task_store

import (
	"context"
	"errors"
	"strings"

	"github.com/iamdanhart/godoyourtasks/server/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresTaskStore struct {
	conn *pgxpool.Pool
}

func NewPostgresTaskStore(conn *pgxpool.Pool) TaskStore {
	return PostgresTaskStore{conn: conn}
}

func (d PostgresTaskStore) GetTasks() ([]model.Task, error) {
	query := "SELECT id, task FROM tasks"
	rows, err := d.conn.Query(context.Background(), query)
	if err != nil {
		return []model.Task{}, errors.New("failed to query for tasks: " + err.Error())
	}
	defer rows.Close()

	tasks := make([]model.Task, 0, 8)
	for rows.Next() {
		var task model.Task
		rows.Scan(&task.Id, &task.Description)
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (d PostgresTaskStore) AddTask(task *model.Task) error {
	if strings.TrimSpace(task.Description) == "" {
		return errors.New("task description must not be blank")
	}
	query := "INSERT INTO tasks (task) VALUES (@taskDescription)"
	args := pgx.NamedArgs{
		"taskDescription": task.Description,
	}
	_, err := d.conn.Exec(context.Background(), query, args)
	if err != nil {
		return errors.New("failed to add task: " + err.Error())
	}
	return nil
}
