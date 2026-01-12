// Package main is a simple Todo API using the standard library
package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /todos", getTodos)
	mux.HandleFunc("POST /todos", postTodo)

	http.ListenAndServe(":8080", mux)
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w) // set NewEncoder to use ResponseWriter as the output stream
	enc.Encode(todos)         // encode todos to the output stream
}

func postTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo

	dec := json.NewDecoder(r.Body) // := requires a specified value on the right side

	if err := dec.Decode(&newTodo); err != nil {
		return
	}

	mu.Lock()                      // locks the goroutine so todos slice can only be modified at this point
	todos = append(todos, newTodo) // appends newTodo from body to todos
	mu.Unlock()                    // unlocks the goroutine and allows todos slice to be modified

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {

}

func updateTodo(w http.ResponseWriter, r *http.Request) {

}
