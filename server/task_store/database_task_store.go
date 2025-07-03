package task_store

import (
	"context"
	"errors"
	"github.com/iamdanhart/godoyourtasks/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseTaskStore struct {
	conn *pgxpool.Pool
}

func NewDatabaseTaskStore(conn *pgxpool.Pool) TaskStore {
	return DatabaseTaskStore{conn: conn}
}

func (d DatabaseTaskStore) GetTasks() ([]model.Task, error) {
	query := "SELECT * FROM tasks"
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

func (d DatabaseTaskStore) AddTask(task *model.Task) error {
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
