package todo

import (
	"errors"
)

var ErrInvalidTaskId = errors.New("invaid task id")
var ErrTaskAlreadyCompleted = errors.New("already completed")

type Todo struct {
	Task string `json:"task,omitempty"`
	Done bool   `json:"done,omitempty"`
}

type Tasks struct {
	todos []Todo
}

func (Tasks) Schema() []string {
	return []string{"ID", "TASK", "DONE"}
}

func (t Tasks) Data() [][]any {
	var data [][]any

	for i, todo := range t.todos {
		data = append(data, []any{i + 1, todo.Task, todo.Done})
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
	if err := t.taskIdIsWithinBounds(id); err != nil {
		return ErrInvalidTaskId
	}

	t.todos = append(t.todos[:id-1], t.todos[id:]...)
	return nil
}

func (t *Tasks) CompleteTask(id int) error {
	if err := t.taskIdIsWithinBounds(id); err != nil {
		return ErrInvalidTaskId
	}

	if t.todos[id-1].Done == true {
		return ErrTaskAlreadyCompleted
	}

	t.todos[id-1].Done = true
	return nil
}

func (t *Tasks) UpdateTask(id int, task string) error {
	if err := t.taskIdIsWithinBounds(id); err != nil {
		return ErrInvalidTaskId
	}

	t.todos[id-1].Task = task
	return nil
}

func (t *Tasks) taskIdIsWithinBounds(id int) error {
	if id < 1 || id > len(t.todos) {
		return ErrInvalidTaskId
	}
	return nil
}
