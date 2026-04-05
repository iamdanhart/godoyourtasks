package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	godoyourtasks "github.com/iamdanhart/godoyourtasks"
	"github.com/iamdanhart/godoyourtasks/server/routes"
	"github.com/iamdanhart/godoyourtasks/server/task_store"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	mode := os.Args[1]

	var taskStore task_store.TaskStore
	switch mode {
	case "postgres":
		// TODO this is just for dev, a real app would never list connection creds in code
		pool, err :=
			pgxpool.New(
				context.Background(),
				"postgres://tasksuser:tasks@localhost:5432/tasks")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
			os.Exit(1)
		}
		taskStore = task_store.NewPostgresTaskStore(pool)
		defer pool.Close()
	case "sqlite":
		connStr := "file:tasks.sqlite"
		db, err := sql.Open("sqlite3", connStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to establish sqlite connection: %v\n", err)
			os.Exit(1)
		}
		taskStore = task_store.NewSqliteTaskStore(db)
	default:
		taskStore = task_store.NewInMemTaskStore()
	}

	mux := routes.NewRouter(taskStore, godoyourtasks.ClientFiles)
	err := http.ListenAndServe(":8081", mux)
	log.Fatal(err)
}
