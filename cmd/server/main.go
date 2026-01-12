package main

import (
	"net/http"

	"github.com/EllisOllier/todo-list-go/internal/todo" // uses repo to import todo code as they are private
)

func main() {
	todoService := todo.NewTodoService() // references NewTodoServer in /internal/model.go
	mux := http.NewServeMux()

	// uses todoService to access routes from /internal/handler.go
	mux.HandleFunc("GET /todos", todoService.GetTodos)
	mux.HandleFunc("POST /todos", todoService.PostTodo)
	mux.HandleFunc("PUT /todos/{id}", todoService.PutTodo)
	mux.HandleFunc("PATCH /todos/{id}", todoService.PatchTodo)
	mux.HandleFunc("DELETE /todos/{id}", todoService.DeleteTodo)

	http.ListenAndServe(":8080", mux)
}
