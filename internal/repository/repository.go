package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sonikq/todo"
)

type Autorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (string, error)
}
type TodoList interface {
	Create(userId int, list todo.TodoList) (int,error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) ([]todo.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo.UpdateListInput)  error
}

type TodoItem interface {
	Create( listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error 
	Update(userId, itemId int, input todo.UpdateItemInput) error
}

type Repository struct {
	Autorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autorization: NewAuthPostgres(db),
		TodoList: NewTodoListPostgres(db),
		TodoItem: NewTodoItemPostgres(db),
	}
}