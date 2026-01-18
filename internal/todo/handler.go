package todo

import (
	"encoding/json"
	"net/http"
)

// runs GET request for all todos slice stored in-memory
func (s *TodoService) GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w) // set NewEncoder to use ResponseWriter as the output stream
	rows, err := s.todoRepository.GetAllTodos()
	if err != nil {
		return
	}
	if rows == nil {
		enc.Encode("Server Error: 500")
	}

	enc.Encode(rows) // encode todos to the output stream
}

// // runs POST request to add new a task to todos slice stored in-memory
// func (s *TodoService) PostTodo(w http.ResponseWriter, r *http.Request) {
// 	var newTodo Todo

// 	dec := json.NewDecoder(r.Body) // := requires a specified value on the right side

// 	if err := dec.Decode(&newTodo); err != nil {
// 		return
// 	}

// 	s.Mu.Lock()                        // locks the goroutine so todos slice can only be modified at this point
// 	s.Todos = append(s.Todos, newTodo) // appends newTodo from body to todos
// 	s.Mu.Unlock()                      // unlocks the goroutine and allows todos slice to be modified

// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(newTodo)
// }

// // runs PUT request to replace a todo item matching the given id
// func (s *TodoService) PutTodo(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(r.PathValue("id"))
// 	if err != nil {
// 		http.Error(w, "Invalid ID format", http.StatusBadRequest)
// 		return
// 	}

// 	var updatedTodo Todo
// 	dec := json.NewDecoder(r.Body)                   // decodes the body
// 	if err := dec.Decode(&updatedTodo); err != nil { // fetches the updated todo from the request body
// 		return
// 	}

// 	json.NewDecoder(r.Body)
// 	for i, v := range s.Todos {
// 		if v.ID == id {
// 			s.Mu.Lock()
// 			s.Todos[i] = updatedTodo
// 			s.Mu.Unlock()
// 			w.WriteHeader(http.StatusOK)
// 			return
// 		}
// 	}
// 	w.WriteHeader(http.StatusNotFound)
// }

// // runs PATCH request to replace a todo item matching the given id
// func (s *TodoService) PatchTodo(w http.ResponseWriter, r *http.Request) {
// 	type UpdateTodoRequest struct {
// 		Task *string `json:"task"`
// 	}

// 	id, err := strconv.Atoi(r.PathValue("id"))
// 	if err != nil {
// 		http.Error(w, "Invalid ID format", http.StatusBadRequest)
// 		return
// 	}

// 	var updateReq UpdateTodoRequest
// 	dec := json.NewDecoder(r.Body)                 // decodes the body
// 	if err := dec.Decode(&updateReq); err != nil { // fetches the updated todo from the request body
// 		return
// 	}

// 	for i, v := range s.Todos {
// 		if v.ID == id {
// 			if updateReq.Task != nil {
// 				s.Mu.Lock()
// 				s.Todos[i].Task = *updateReq.Task
// 				s.Mu.Unlock()
// 			}

// 			w.WriteHeader(http.StatusOK)
// 			return
// 		}
// 	}
// 	w.WriteHeader(http.StatusNotFound)
// }

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
