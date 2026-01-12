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
			{ID: 1, Task: "Add sqlite file to project"},
			{ID: 2, Task: "Add sqlite.go database connection logic"},
			{ID: 3, Task: "Move any data logic from handler.go to repository.go which interacts directly with database"},
		},
	}
}
