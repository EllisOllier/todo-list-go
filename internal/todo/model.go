package todo

// declares the structure for todos slice
type Todo struct {
	ID     int    `json:"id"`
	Task   string `json:"task"`
	Done   bool   `json:"done"`
	UserId int    `json:"user_id"`
}

// delcares the structure for the TodoService helper
type TodoService struct {
	todoRepository *TodoRepository
}

// a helper function to initialise and handle database interaction with TodoRepository
// which allows the data from the database to be saved in memory
// and accessed from both /cmd/server/main.go and /internal/todo/handler.go
// with the same dataset syncing across
func NewTodoService(givenTodoRepository *TodoRepository) *TodoService {
	return &TodoService{
		todoRepository: givenTodoRepository,
	}
}
