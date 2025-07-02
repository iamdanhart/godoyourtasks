package main

import (
	"github.com/iamdanhart/godoyourtasks/routes"
	"github.com/iamdanhart/godoyourtasks/task_store"
	"log"
	"net/http"
)

func main() {

	taskStore := task_store.NewInMemTaskStore()
	mux := routes.NewRouter(taskStore)

	err := http.ListenAndServe(":8081", mux)
	log.Fatal(err)
}
