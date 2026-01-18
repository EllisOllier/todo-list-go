package main

import (
	"net/http"

	// imported with command: "go mod init github.com/EllisOllier/todo-list-go"
	"github.com/EllisOllier/todo-list-go/internal/database"
	"github.com/EllisOllier/todo-list-go/internal/todo" // uses repo to import /internal/todo code as they are private
)

func main() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	todoRepository := todo.NewTodoRepository(db)
	todoService := todo.NewTodoService(todoRepository) // references NewTodoServer in /internal/todo/model.go
	mux := http.NewServeMux()

	// uses todoService to access routes from /internal/todo/handler.go
	mux.HandleFunc("GET /todos", todoService.GetTodos)
	mux.HandleFunc("GET /todos/{id}", todoService.GetTodoById)
	mux.HandleFunc("POST /todos", todoService.PostTodo)
	mux.HandleFunc("PATCH /todos/{id}", todoService.PatchTodo)
	// mux.HandleFunc("DELETE /todos/{id}", todoService.DeleteTodo)

	http.ListenAndServe(":8080", mux)
}
