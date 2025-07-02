package task_store

import (
	"github.com/iamdanhart/godoyourtasks/model"
)

type InMemTaskStore struct {
	tasks []model.Task
}

func NewInMemTaskStore() *InMemTaskStore {
	return &InMemTaskStore{
		tasks: make([]model.Task, 0, 8),
	}
}

func (taskStore *InMemTaskStore) GetTasks() ([]model.Task, error) {
	return taskStore.tasks, nil
}

func (taskStore *InMemTaskStore) AddTask(task *model.Task) error {
	task.Id = len(taskStore.tasks)
	taskStore.tasks = append(taskStore.tasks, *task)
	return nil
}
