package main

import (
	"log"
	"net/http"
	"os"

	// imported with command: "go mod init github.com/EllisOllier/todo-list-go"
	_ "github.com/EllisOllier/todo-list-go/docs"

	"github.com/EllisOllier/todo-list-go/internal/database"
	"github.com/EllisOllier/todo-list-go/internal/middleware"
	"github.com/EllisOllier/todo-list-go/internal/todo" // uses repo to import /internal/todo code as they are private
	"github.com/EllisOllier/todo-list-go/internal/user"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Todo List API
// @version 1.0
// @description A simple todo list api written with Golang using a postgresql docker and JWT authentication
// @host localhost:8080
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

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
	mux.Handle("PATCH /todos/{id}", middleware.Authenticate(http.HandlerFunc(todoService.PatchTodo)))
	mux.Handle("DELETE /todos/{id}", middleware.Authenticate(http.HandlerFunc(todoService.DeleteTodo)))

	// uses userService to access routes from /internal/user/handler.go
	mux.HandleFunc("POST /user", userService.CreateAccount)
	mux.HandleFunc("POST /user/login", userService.Login)

	mux.Handle("GET /swagger/", httpSwagger.WrapHandler)

	middlewareMux := middleware.LoggingMiddleware(mux) // all routes run through mux
	http.ListenAndServe(port, middlewareMux)
}
