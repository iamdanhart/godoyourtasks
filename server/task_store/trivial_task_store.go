package task_store

import (
	"github.com/iamdanhart/godoyourtasks/model"
)

type TrivialTaskStore struct {
	tasks []model.Task
}

func NewTrivialTaskStore() *TrivialTaskStore {
	return &TrivialTaskStore{
		tasks: []model.Task{
			{
				Id:          1,
				Description: "Trivial task"},
		},
	}
}

func (store *TrivialTaskStore) GetTasks() ([]model.Task, error) {
	return store.tasks, nil
}

func (store *TrivialTaskStore) AddTask(_ *model.Task) error {
	return nil
}
