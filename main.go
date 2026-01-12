// Package main is a simple Todo API using the standard library
package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /todos", getTodos)
	mux.HandleFunc("POST /todos", postTodo)
	mux.HandleFunc("PUT /todos/{id}", putTodo)
	mux.HandleFunc("PATCH /todos/{id}", patchTodo)
	mux.HandleFunc("DELETE /todos/{id}", deleteTodo)

	http.ListenAndServe(":8080", mux)
}

// runs GET request for all todos slice stored in-memory
func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w) // set NewEncoder to use ResponseWriter as the output stream
	enc.Encode(todos)         // encode todos to the output stream
}

// runs POST request to add new a task to todos slice stored in-memory
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

// runs PUT request to replace a todo item matching the given id
func putTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var updatedTodo Todo
	dec := json.NewDecoder(r.Body)                   // decodes the body
	if err := dec.Decode(&updatedTodo); err != nil { // fetches the updated todo from the request body
		return
	}

	json.NewDecoder(r.Body)
	for _, v := range todos {
		if v.ID == id {
			mu.Lock()
			todos[v.ID] = updatedTodo
			mu.Unlock()
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

// runs PATCH request to replace a todo item matching the given id
func patchTodo(w http.ResponseWriter, r *http.Request) {
	type UpdateTodoRequest struct {
		Task *string `json:"task"`
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var updateReq UpdateTodoRequest
	dec := json.NewDecoder(r.Body)                 // decodes the body
	if err := dec.Decode(&updateReq); err != nil { // fetches the updated todo from the request body
		return
	}

	for i, v := range todos {
		if v.ID == id {
			if updateReq.Task != nil {
				mu.Lock()
				todos[i].Task = *updateReq.Task
				mu.Unlock()
			}

			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

// runs DELETE request to delete a todo item task matching the given id
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	for i, v := range todos {
		if v.ID == id {
			mu.Lock()
			todos = append(todos[:i], todos[i+1:]...) // ... unpacks the slice into individual items so append can handle
			mu.Unlock()
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
