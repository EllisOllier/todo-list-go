# todo-list-go
**A lightweight Todo API built using only the Go standard library**

### Features
- GET route which returns all todos in the slice
- GET route which returns the todo matching a specified ID
- POST route which appends a new todo at the end of the slice
- PUT route which replaces a specficied todo index
- PATCH route which replaces the todo Task of a specified todo ID
- DELETE route which deletes a specified todo ID

### About
This project was created primarily as a reference for designing and making Go api routes using the standard library.  

Functionally this project is a todo list api which allows the user to:
- Check all of their todos
- Add new todos to the list
- Update todos in the list (Renaming a task)
- Delete todos from the list

### Multi-file Structure
/cmd/server/main.go:  
Handles the setup of the server and calling the /internal/todo/handler.go routes via importing routes from github repo

/internal/todo/handler.go:  
Handles all logic and data handling for the routes

/internal/todo/model.go:  
Declares structure for Todo and TodoService struct which is implemented with NewTodoService helper function

### Documentation Used
https://pkg.go.dev/net/http  
https://pkg.go.dev/encoding/json  
https://pkg.go.dev/sync  
https://pkg.go.dev/strconv  
https://go.dev/blog/slices-intro  
https://github.com/golang-standards/project-layout  