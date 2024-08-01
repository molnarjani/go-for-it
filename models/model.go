package models

import (
	"fmt"
	"log/slog"
)

type Todo struct {
	Id    string
	Title string
	Done  bool
}

type TodoStore struct {
	todos []Todo
}

func (ts *TodoStore) Add(todo Todo) {
	ts.todos = append(ts.todos, todo)
}

func (ts *TodoStore) List() []Todo {
	return ts.todos
}

func (ts *TodoStore) Get(todoId string) (int, Todo, error) {
	for i, todo := range ts.todos {
		if todo.Id == todoId {
			return i, todo, nil
		}
	}
	return -1, Todo{}, fmt.Errorf("todo not found")
}

func (ts *TodoStore) Update(todoId string, todo Todo) error {
	slog.Info("Updating todo", "id", todoId, "todo", todo)
	i, currentTodo, err := ts.Get(todoId)
	if err != nil {
		return err
	}
	ts.todos[i] = Todo{
		Id:    currentTodo.Id,
		Title: todo.Title,
		Done:  todo.Done,
	}
	return nil
}

func (ts *TodoStore) Delete(todoId string) error {
	slog.Info("Deleting todo", "id", todoId)
	index, _, err := ts.Get(todoId)
	if err != nil {
		return err
	}

	ts.todos = append(ts.todos[:index], ts.todos[index+1:]...)

	return nil
}

func NewTodoStore() *TodoStore {
	return &TodoStore{}
}
