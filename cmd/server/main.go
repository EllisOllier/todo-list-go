package main

import (
	"log"
	"net/http"
	"os"

	// imported with command: "go mod init github.com/EllisOllier/todo-list-go"
	"github.com/EllisOllier/todo-list-go/internal/database"
	"github.com/EllisOllier/todo-list-go/internal/middleware"
	"github.com/EllisOllier/todo-list-go/internal/todo" // uses repo to import /internal/todo code as they are private
	"github.com/EllisOllier/todo-list-go/internal/user"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("SERVER_PORT")

	log.Println("Attempting to connect to database...")
	db, err := database.Connect()
	if err != nil {
		log.Println("Failed to connect to database: ", err)
		panic(err)
	}
	log.Println("Successfully connected to database")

	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)

	todoRepository := todo.NewTodoRepository(db)
	todoService := todo.NewTodoService(todoRepository) // references NewTodoServer in /internal/todo/model.go
	mux := http.NewServeMux()

	// uses todoService to access routes from /internal/todo/handler.go
	mux.Handle("GET /todos", middleware.Authenticate(http.HandlerFunc(todoService.GetTodos)))
	mux.HandleFunc("GET /todos/{id}", todoService.GetTodoById)
	mux.Handle("POST /todos", middleware.Authenticate(http.HandlerFunc(todoService.PostTodo)))
	mux.HandleFunc("PATCH /todos/{id}", todoService.PatchTodo)
	mux.HandleFunc("DELETE /todos/{id}", todoService.DeleteTodo)

	// uses userService to access routes from /internal/user/handler.go
	mux.HandleFunc("POST /user", userService.CreateAccount)
	mux.HandleFunc("POST /user/login", userService.Login)

	middlewareMux := middleware.LoggingMiddleware(mux) // all routes run through mux
	http.ListenAndServe(port, middlewareMux)
}
