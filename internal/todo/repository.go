package todo

import (
	"database/sql"
)

// add database integration with sqlite
// will run sql queries which are handled by the handler.go
// works very similar to controller (handler.go) and model (repository.go) architecture

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(givenDb *sql.DB) *TodoRepository {
	return &TodoRepository{
		db: givenDb,
	}
}

func (r *TodoRepository) GetAllTodos() ([]Todo, error) {
	var todos []Todo
	rows, err := r.db.Query("SELECT * FROM tasks")
	if err != nil {
		return todos, err
	}

	defer rows.Close() // [https://gobyexample.com/defer] essential, similar to closing an I/O stream in Java

	for rows.Next() {
		var temp Todo
		rows.Scan(&temp.ID, &temp.Task)
		todos = append(todos, temp)
	}

	return todos, nil
}

func (r *TodoRepository) GetTodoById(id int) (Todo, error) {
	var temp Todo
	err := r.db.QueryRow("SELECT id, task FROM tasks WHERE id = $1", id).Scan(&temp.ID, &temp.Task)
	if err != nil {
		return temp, err
	}

	return temp, nil
}

func (r *TodoRepository) AddTodo(todo Todo) (int, error) {
	var entryId int
	err := r.db.QueryRow("INSERT INTO tasks (task) VALUES ($1) RETURNING id", todo.Task).Scan(&entryId)
	if err != nil {
		return -1, err
	}
	return entryId, nil
}

func (r *TodoRepository) UpdateTodo(id int, todo Todo) error {
	res, err := r.db.Exec("UPDATE tasks set task=$1 WHERE id=$2", todo.Task, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *TodoRepository) DeleteTodo(id int) error {
	res, err := r.db.Exec("DELETE FROM tasks WHERE id=$1", id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return sql.ErrNoRows
	}

	return nil
}
