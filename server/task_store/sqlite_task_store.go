package task_store

import (
	"database/sql"
	"errors"
	"github.com/iamdanhart/godoyourtasks/model"
)

type SqliteTaskStore struct {
	conn *sql.DB
}

func NewSqliteTaskStore(conn *sql.DB) TaskStore {
	return SqliteTaskStore{conn: conn}
}

func (d SqliteTaskStore) GetTasks() ([]model.Task, error) {
	query := "SELECT * FROM tasks"
	rows, err := d.conn.Query(query)
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

func (d SqliteTaskStore) AddTask(task *model.Task) error {
	query := "INSERT INTO tasks (task) VALUES (?)"
	_, err := d.conn.Exec(query, task.Description)
	if err != nil {
		return errors.New("failed to add task: " + err.Error())
	}
	return nil
}
