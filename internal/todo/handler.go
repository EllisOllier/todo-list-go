package todo

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// runs GET request for all todos slice stored in-memory
func (s *TodoService) GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w) // set NewEncoder to use ResponseWriter as the output stream
	rows, err := s.todoRepository.GetAllTodos()
	if err != nil {
		http.Error(w, "Server Error: 500", http.StatusInternalServerError)
		return
	}
	if rows == nil { // test this as it doesnt really need to return NoResults
		http.Error(w, "No results found", http.StatusNoContent) // might be wrong status
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
	type NewTodoRequest struct {
		Task *string `json:"task"`
	}

	var req NewTodoRequest

	dec := json.NewDecoder(r.Body) // := requires a specified value on the right side
	if err := dec.Decode(&req); err != nil {
		http.Error(w, "Bad Request: 400", http.StatusBadRequest)
		return
	}
	if req.Task == nil {
		http.Error(w, "Missing task field", http.StatusBadRequest)
		return
	}
	newTodo := Todo{Task: *req.Task}

	todoId, err := s.todoRepository.AddTodo(newTodo)
	if err != nil {
		http.Error(w, "Server Error: 500", http.StatusInternalServerError)
		return
	}
	newTodo.ID = todoId

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

// runs PATCH request to replace a todo item matching the given id
func (s *TodoService) PatchTodo(w http.ResponseWriter, r *http.Request) {
	type UpdateTodoRequest struct {
		Task *string `json:"task"`
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var req UpdateTodoRequest
	dec := json.NewDecoder(r.Body)           // decodes the body
	if err := dec.Decode(&req); err != nil { // fetches the updated todo from the request body
		return
	}

	updateTodo := Todo{ID: id, Task: *req.Task}
	repoErr := s.todoRepository.UpdateTodo(id, updateTodo)
	if repoErr != nil {
		if errors.Is(repoErr, sql.ErrNoRows) {
			http.Error(w, "Not Found: 404", http.StatusNotFound)
			return
		}
		http.Error(w, "Server Error: 500", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updateTodo)
}

// // runs DELETE request to delete a todo item task matching the given id
// func (s *TodoService) DeleteTodo(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(r.PathValue("id"))
// 	if err != nil {
// 		http.Error(w, "Invalid ID format", http.StatusBadRequest)
// 		return
// 	}

// 	for i, v := range s.Todos {
// 		if v.ID == id {
// 			s.Mu.Lock()
// 			s.Todos = append(s.Todos[:i], s.Todos[i+1:]...) // ... unpacks the slice into individual items so append can handle
// 			s.Mu.Unlock()
// 			w.WriteHeader(http.StatusNoContent)
// 			return
// 		}
// 	}
// 	w.WriteHeader(http.StatusNotFound)
// }
