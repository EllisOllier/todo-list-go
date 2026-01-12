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

// a helper function to initialise and handle Todos slice
func NewTodoService() *TodoService {
	return &TodoService{
		Todos: []Todo{
			{ID: 1, Task: "Make my other end points for todo-list-go"},
			{ID: 2, Task: "Did this one delete or stay?"},
		},
	}
}
