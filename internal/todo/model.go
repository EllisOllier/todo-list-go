package todo

// declares the structure for todos slice
type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

// delcares the structure for the TodoService helper
type TodoService struct {
	todoRepository *TodoRepository
}

// a helper function to initialise and handle Todos slice which
// allows a single instance to be saved in memory and accessed from
// both /cmd/server/main.go and /internal/todo/handler.go
// with the same dataset syncing across
func NewTodoService(givenTodoRepository *TodoRepository) *TodoService {
	return &TodoService{
		todoRepository: givenTodoRepository,
	}
}
