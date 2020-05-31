package repository

import (
	"../entities"
	"../datastore"
)

type TodoRepository struct {
}

//GetAll get all Todo items from DB 
func (repo *TodoRepository) GetAll() ([]*entities.Todo, error) {
	database, err := datastore.NewDatabase()
	if err != nil {
		return nil, err
	}
	
	rows, error2 := database.Query("SELECT id, description FROM todos")
	if err != nil {
		return nil, error2
	}

	var todos []*entities.Todo
	for rows.Next() {
		var todo entities.Todo
		rows.Scan(&todo.Id, &todo.Description);
		todos = append(todos, &todo)
	}

	return todos, nil
}

//GetById get Todo item by Id
func (repo *TodoRepository) GetById(id string) (*entities.Todo, error) {
	database, err := datastore.NewDatabase()
	if err != nil {
		return nil, err
	}
	
	println("id", id)
	var todo entities.Todo
	row := database.QueryRow("SELECT id, description FROM todos WHERE id=?", id)

	err = row.Scan(&todo.Id, &todo.Description);
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

//DeleteById delete Todo item by Id
func (repo *TodoRepository) DeleteById(id *string) (bool, error) {
	database, err := datastore.NewDatabase()
	if err != nil {
		return false, err
	}

	result, err := database.Exec("DELETE FROM todos WHERE id=?", id)
	if err != nil {
		return false, err
	}

	rowsAffected , err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

//Create create todo item by passing description in body
func (repo *TodoRepository) Create(todo *entities.Todo) (bool) {
	database, err := datastore.NewDatabase()
	if err != nil {
		return false
	}
	
	result, err := database.Exec("INSERT INTO todos(Description) values(?)", todo.Description)
	if err != nil {
		return false
	}
	

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false
	}

	return rowsAffected > 0
}

//UpdateById update description on Todo item by id
func (repo *TodoRepository) UpdateById(id *string, todo *entities.Todo) (bool) {
	database, err := datastore.NewDatabase()
	if err != nil {
		return false
	}
	
	result, err := database.Exec("UPDATE todos SET description=? where id=?", todo.Description, id)
	if err != nil {
		return false
	}
	

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false
	}

	return rowsAffected > 0
}