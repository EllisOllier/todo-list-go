package todo

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/EllisOllier/todo-list-go/internal/middleware"
)

type TodoRequest struct {
	Task *string `json:"task"`
	Done *bool   `json:"done"`
}

// runs GET request for all todos slice stored in-memory
func (s *TodoService) GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rawId := r.Context().Value(middleware.UserIdKey)
	userId, ok := rawId.(int)
	if !ok {
		http.Error(w, "Could not find user ID", http.StatusUnauthorized)
		return
	}

	enc := json.NewEncoder(w) // set NewEncoder to use ResponseWriter as the output stream
	rows, err := s.todoRepository.GetAllTodos(userId)
	if err != nil {
		http.Error(w, "Server Error: 500", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	enc.Encode(rows) // encode todos to the output stream
}

func (s *TodoService) GetTodoById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	row, err := s.todoRepository.GetTodoById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Not Found: 404", http.StatusNotFound)
			return
		}
		http.Error(w, "Server Error: 500", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	enc.Encode(row)
}

// runs POST request to add new a task to todos slice stored in-memory
func (s *TodoService) PostTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rawId := r.Context().Value(middleware.UserIdKey)
	userId, ok := rawId.(int)
	if !ok {
		http.Error(w, "Could not find user ID", http.StatusUnauthorized)
		return
	}

	var req TodoRequest

	dec := json.NewDecoder(r.Body) // := requires a specified value on the right side
	if err := dec.Decode(&req); err != nil {
		http.Error(w, "Bad Request: 400", http.StatusBadRequest)
		return
	}
	if req.Task == nil {
		http.Error(w, "Missing task field", http.StatusBadRequest)
		return
	}

	isDone := false // always false as new todo
	newTodo := Todo{Task: req.Task, Done: &isDone, UserId: userId}

	todoId, err := s.todoRepository.AddTodo(newTodo, userId)
	if err != nil {
		http.Error(w, "Server Error: 500", http.StatusInternalServerError)
		return
	}
	newTodo.ID = *todoId

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

// runs PATCH request to replace a todo item matching the given id
func (s *TodoService) PatchTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rawId := r.Context().Value(middleware.UserIdKey)
	userId, ok := rawId.(int)
	if !ok {
		http.Error(w, "Could not find user ID", http.StatusUnauthorized)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var req TodoRequest
	dec := json.NewDecoder(r.Body)           // decodes the body
	if err := dec.Decode(&req); err != nil { // fetches the updated todo from the request body
		return
	}

	updateTodo := Todo{ID: id, Task: req.Task, Done: req.Done}
	todo, repoErr := s.todoRepository.UpdateTodo(updateTodo, id, userId)
	if repoErr != nil {
		if errors.Is(repoErr, sql.ErrNoRows) {
			http.Error(w, "Not Found: 404", http.StatusNotFound)
			return
		}
		http.Error(w, "Server Error: 500", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}

// runs DELETE request to delete a todo item task matching the given id
func (s *TodoService) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	rawId := r.Context().Value(middleware.UserIdKey)
	userId, ok := rawId.(int)
	if !ok {
		http.Error(w, "Could not find user ID", http.StatusUnauthorized)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	repoErr := s.todoRepository.DeleteTodo(id, userId)
	if repoErr != nil {
		if errors.Is(repoErr, sql.ErrNoRows) {
			http.Error(w, "Not Found: 404", http.StatusNotFound)
			return
		}
		http.Error(w, "Server Error: 500", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
