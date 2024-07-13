package todo

import (
	"testing"
)

func TestNewTodoList(t *testing.T) {
	todoList := NewTodoList()
	if todoList == nil {
		t.Fatal("expected non-nil TodoList")
	}

	if len(todoList.todos) != 0 {
		t.Errorf("expected empty todo list, got %d", len(todoList.todos))
	}

	if todoList.nextID != 1 {
		t.Errorf("expected nextID to be 1, got %d", todoList.nextID)
	}
}

func TestAdd(t *testing.T) {
	todoList := NewTodoList()

	task := "first task"

	todoList.Add(task)

	if todoList.nextID != 2 {
		t.Errorf("expected nextID to be 2 after task was added, got %d", todoList.nextID)
	}

	if todoList.todos[0].Task != task {
		t.Errorf("expected task to be %s, got %s", task, todoList.todos[0].Task)
	}

}
