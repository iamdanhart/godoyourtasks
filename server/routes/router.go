package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/iamdanhart/godoyourtasks/model"
	"github.com/iamdanhart/godoyourtasks/task_store"
	"io"
	"log"
	"net/http"
	"strings"
)

var taskStore task_store.TaskStore

func NewRouter(store task_store.TaskStore) *http.ServeMux {
	if store == nil {
		log.Fatalln("Need a non-nil task store")
	}
	taskStore = store

	router := http.NewServeMux()
	router.HandleFunc("GET /tasks", getTasksHandler)
	router.HandleFunc("POST /tasks", addTaskHandler)

	return router
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	//var err error
	tasks, err := taskStore.GetTasks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error getting tasks: %v\n", err)
		return
	}
	tasksJson, _ := json.Marshal(tasks)
	w.Header().Set("Content-Type", "application/json")
	w.Write(tasksJson)
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	newTask, err := parseTask(w, r)
	if err != nil {
		// parseTask handles writing the error response
		// so we don't need to dig into the error here
		return
	}
	_ = taskStore.AddTask(&newTask)
	//w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	//w.Write([]byte("POST tasks placeholder"))
}

// reference for parsing body https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body
func parseTask(w http.ResponseWriter, r *http.Request) (model.Task, error) {
	// Use http.MaxBytesReader to enforce a maximum read of 1MB from the
	// response body. A request body larger than that will now result in
	// Decode() returning a "http: request body too large" error.
	maxBytesReader := http.MaxBytesReader(w, r.Body, 1048576)

	// Setup the decoder and call the DisallowUnknownFields() method on it.
	// This will cause Decode() to return a "json: unknown field ..." error
	// if it encounters any extra unexpected fields in the JSON. Strictly
	// speaking, it returns an error for "keys which do not match any
	// non-ignored, exported fields in the destination".
	dec := json.NewDecoder(maxBytesReader)
	dec.DisallowUnknownFields()

	var newTask model.Task
	err := dec.Decode(&newTask)
	if err != nil {
		handleParseError(err, w)
		return model.Task{}, errors.New("failed to decode task from input")
	}

	// Call decode again, using a pointer to an empty anonymous struct as
	// the destination. If the request body only contained a single JSON
	// object this will return an io.EOF error. So if we get anything else,
	// we know that there is additional data in the request body.
	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		msg := "Request body must only contain a single JSON object"
		http.Error(w, msg, http.StatusBadRequest)
		return model.Task{}, errors.New("")
	}

	return newTask, nil
}

func handleParseError(err error, w http.ResponseWriter) {
	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError

	switch {
	// Catch any syntax errors in the JSON and send an error message
	// which interpolates the location of the problem to make it
	// easier for the client to fix.
	case errors.As(err, &syntaxError):
		msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
		http.Error(w, msg, http.StatusBadRequest)

	// In some circumstances Decode() may also return an
	// io.ErrUnexpectedEOF error for syntax errors in the JSON. There
	// is an open issue regarding this at
	// https://github.com/golang/go/issues/25956.
	case errors.Is(err, io.ErrUnexpectedEOF):
		msg := fmt.Sprintf("Request body contains badly-formed JSON")
		http.Error(w, msg, http.StatusBadRequest)

	// Catch any type errors, like trying to assign a string in the
	// JSON request body to a int field in our Person struct. We can
	// interpolate the relevant field name and position into the error
	// message to make it easier for the client to fix.
	case errors.As(err, &unmarshalTypeError):
		msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
		http.Error(w, msg, http.StatusBadRequest)

	// Catch the error caused by extra unexpected fields in the request
	// body. We extract the field name from the error message and
	// interpolate it in our custom error message. There is an open
	// issue at https://github.com/golang/go/issues/29035 regarding
	// turning this into a sentinel error.
	case strings.HasPrefix(err.Error(), "json: unknown field "):
		fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
		msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
		http.Error(w, msg, http.StatusBadRequest)

	// An io.EOF error is returned by Decode() if the request body is
	// empty.
	case errors.Is(err, io.EOF):
		msg := "Request body must not be empty"
		http.Error(w, msg, http.StatusBadRequest)

	// Catch the error caused by the request body being too large. Again
	// there is an open issue regarding turning this into a sentinel
	// error at https://github.com/golang/go/issues/30715.
	case err.Error() == "http: request body too large":
		msg := "Request body must not be larger than 1MB"
		http.Error(w, msg, http.StatusRequestEntityTooLarge)

	// Otherwise default to logging the error and sending a 500 Internal
	// Server Error response.
	default:
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

}
