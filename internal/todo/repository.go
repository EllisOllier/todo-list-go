package todo

import (
	"database/sql"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(givenDb *sql.DB) *TodoRepository {
	return &TodoRepository{
		db: givenDb,
	}
}

func (r *TodoRepository) GetAllTodos(userId int) (*[]Todo, error) {
	todos := []Todo{}
	rows, err := r.db.Query("SELECT id, task, done, user_id FROM tasks WHERE user_id=$1", userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close() // [https://gobyexample.com/defer] essential, similar to closing an I/O stream in Java

	for rows.Next() {
		var temp Todo
		rows.Scan(&temp.ID, &temp.Task, &temp.Done, &temp.UserId)
		todos = append(todos, temp)
	}

	return &todos, nil
}

func (r *TodoRepository) GetTodoById(id int) (*Todo, error) {
	var temp Todo
	err := r.db.QueryRow("SELECT id, task FROM tasks WHERE id = $1", id).Scan(&temp.ID, &temp.Task)
	if err != nil {
		return nil, err
	}

	return &temp, nil
}

func (r *TodoRepository) AddTodo(todo Todo, userId int) (*int, error) {
	var entryId int
	err := r.db.QueryRow("INSERT INTO tasks (task, done, user_id) VALUES ($1, $2, $3) RETURNING id", todo.Task, todo.Done, userId).Scan(&entryId)
	if err != nil {
		return nil, err
	}
	return &entryId, nil
}

func (r *TodoRepository) UpdateTodo(todo Todo, id int, userId int) (*Todo, error) {
	var updatedTodo Todo

	err := r.db.QueryRow("UPDATE tasks SET task = COALESCE($1, task), done = COALESCE($2, done) WHERE id = $3 AND user_id = $4 RETURNING id, task, done, user_id", todo.Task, todo.Done, id, userId).Scan(&updatedTodo.ID, &updatedTodo.Task, &updatedTodo.Done, &updatedTodo.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &updatedTodo, nil
}

func (r *TodoRepository) DeleteTodo(id int, userId int) error {
	res, err := r.db.Exec("DELETE FROM tasks WHERE id=$1 AND user_id=$2", id, userId)
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
