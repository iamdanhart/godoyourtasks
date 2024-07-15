package task_store

import (
	"github.com/iamdanhart/godoyourtasks/model"
)

type InMemTaskStore struct {
	tasks []model.Task
}

func NewInMemTaskStore() *InMemTaskStore {
	return &InMemTaskStore{
		tasks: make([]model.Task, 0),
	}
}

func (taskStore *InMemTaskStore) GetTasks() []model.Task {
	return taskStore.tasks
}

func (taskStore *InMemTaskStore) AddTask(task model.Task) error {
	taskStore.tasks = append(taskStore.tasks, task)
	return nil
}
