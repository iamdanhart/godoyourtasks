package task_store

import (
	"database/sql"
	"errors"
	"log"

	"github.com/iamdanhart/godoyourtasks/server/model"
)

type SqliteTaskStore struct {
	conn *sql.DB
}

func NewSqliteTaskStore(conn *sql.DB) TaskStore {
	return SqliteTaskStore{conn: conn}
}

func (d SqliteTaskStore) GetTasks() ([]model.Task, error) {
	query := "SELECT id, task FROM tasks"
	rows, err := d.conn.Query(query)
	if err != nil {
		return []model.Task{}, errors.New("failed to query for tasks: " + err.Error())
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}(rows)

	tasks := make([]model.Task, 0, 8)
	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.Id, &task.Description)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (d SqliteTaskStore) AddTask(task *model.Task) error {
	query := "INSERT INTO tasks (task) VALUES (?)"
	_, err := d.conn.Exec(query, task.Description)
	if err != nil {
		return errors.New("failed to add task: " + err.Error())
	}
	return nil
}
