package task_store

import "github.com/iamdanhart/godoyourtasks/model"

type TrivialTaskStore struct{}

var trivialTask = model.Task{
	Id:          1,
	Description: "Trivial task",
}

func (store *TrivialTaskStore) GetTasks() ([]model.Task, error) {
	return []model.Task{trivialTask}, nil
}

func (store *TrivialTaskStore) AddTask(unused *model.Task) error {
	return nil
}
