package main

import (
	"encoding/json"
	"net/http"
	"sync"
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
	enc.Encode(todos)         // Encode todos to the output streamÍÍ
}

func postTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo
	var mu sync.Mutex // sync.Mutex is a type not a value
	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(&newTodo); err != nil {
		return
	}

	mu.Lock()
	todos = append(todos, newTodo)
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var todos = []Todo{
	{ID: 1, Task: "Make my other end points for todo-list-go"},
}
