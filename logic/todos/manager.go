package todos

import (
	"fmt"
	. "sourav.kabiraj/goboilerplate/logic/todos/models"
)

type Manager struct {
	todos  []Todo
	nextID int
}

func NewManager() *Manager {
	return &Manager{
		todos:  []Todo{},
		nextID: 1,
	}
}

func (tm *Manager) GetTodos() []Todo {
	return tm.todos
}

func (tm *Manager) AddTodo(todo Todo) Todo {
	todo.ID = tm.nextID
	tm.nextID++
	tm.todos = append(tm.todos, todo)
	return todo
}

func (tm *Manager) UpdateTodo(updatedTodo Todo) (*Todo, error) {
	for i, todo := range tm.todos {
		if todo.ID == updatedTodo.ID {
			tm.todos[i] = updatedTodo
			return &updatedTodo, nil
		}
	}
	return nil, fmt.Errorf("Todo not found")
}

func (tm *Manager) DeleteTodoByID(id int) error {
	for i, todo := range tm.todos {
		if todo.ID == id {
			tm.todos = append(tm.todos[:i], tm.todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Todo not found")
}
