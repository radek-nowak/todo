package todo

import (
	"errors"
)

var ErrInvalidTaskId = errors.New("Invaid task id")
var ErrTaskAlreadyCompleted = errors.New("Already completed")

type Todo struct {
	// ID   int    `json:"id,omitempty"`
	Task string `json:"task,omitempty"`
	Done bool   `json:"done,omitempty"`
}

type Tasks struct {
	todos  []Todo
}

func (Tasks) Schema() []string {
	return []string{"ID", "TASK", "DONE"}
}

func (t Tasks) Data() [][]interface{} {
	var data [][]interface{}

	for i, todo := range t.todos {
		data = append(data, []interface{}{i+1, todo.Task, todo.Done})
	}

	return data
}

func (Tasks) ColumnWidths() []int {
	return []int{3, 55, 5}
}

func NewTodoList() *Tasks {
	return &Tasks{}
}

func FromTodos(todos []Todo) *Tasks {
	return &Tasks{todos: todos}
}

func (t *Tasks) GetTodos() []Todo {
	return t.todos
}

func (t *Tasks) Add(task string) {
	newTodo := Todo{
		Task: task,
		Done: false,
	}

	t.todos = append(t.todos, newTodo)
}

func (t *Tasks) Delete(id int) error {
	if id < 1 || id > len(t.todos) {
		return ErrInvalidTaskId
	}

	// todo add get metod that checks if task exists, and returns error if not
	t.todos = append(t.todos[:id-1], t.todos[id:]...)
	return nil
}

func (t *Tasks) CompleteTask(id int) error {
	if id < 1 || id > len(t.todos) {
		return ErrInvalidTaskId
	}

	if t.todos[id-1].Done == true {
		return ErrTaskAlreadyCompleted
	}

	t.todos[id-1].Done = true
	return nil
}

func (t *Tasks) UpdateTask(id int, task string) error {
	if id < 1 || id > len(t.todos) {
		return ErrInvalidTaskId
	}

	t.todos[id-1].Task = task
	return nil
}
