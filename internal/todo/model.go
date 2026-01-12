package todo

import "sync"

// declares the structure for todos slice
type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

// delcares the structure for the TodoService helper
type TodoService struct {
	Mu    sync.Mutex
	Todos []Todo
}

// todos holds the stored in memory list of tasks
var Todos = []Todo{
	{ID: 1, Task: "Make my other end points for todo-list-go"},
	{ID: 2, Task: "Did this one delete or stay?"},
}

// mu is used to arrange access to the todos slice across goroutines
var Mu sync.Mutex // sync.Mutex is a type not a value

// a helper function to initialise and handle Todos slice
func NewTodoService() *TodoService {
	return &TodoService{
		Todos: []Todo{
			{ID: 1, Task: "Make my other end points for todo-list-go"},
			{ID: 2, Task: "Did this one delete or stay?"},
		},
	}
}
