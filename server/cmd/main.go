package main

import (
	"context"
	"fmt"
	"github.com/iamdanhart/godoyourtasks/routes"
	"github.com/iamdanhart/godoyourtasks/task_store"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
	"os"
)

func main() {
	// TODO this is just for dev, a real app would never list connection creds in code
	pool, err :=
		pgxpool.New(
			context.Background(),
			"postgres://tasksuser:tasks@localhost:5432/tasks")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	//defer dbpool.Close()

	taskStore := task_store.NewDatabaseTaskStore(pool)
	mux := routes.NewRouter(taskStore)

	err = http.ListenAndServe(":8081", mux)
	log.Fatal(err)
}
