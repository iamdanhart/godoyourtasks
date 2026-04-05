package task_store

import (
	"github.com/iamdanhart/godoyourtasks/server/model"
)

type TaskStore interface {
	GetTasks() ([]model.Task, error)
	AddTask(task *model.Task) error
}
