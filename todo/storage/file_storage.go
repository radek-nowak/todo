package storage

import (
	"encoding/json"
	todo "go_todo/todo/model"
	"os"
)

func ReadData(path string) (*todo.TodoList, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var todos []todo.Todo
	
	json.Unmarshal(data, &todos)

	return todo.FromTodos(todos), nil
}

func WriteData(path string, todoList *todo.TodoList) error {
	bytes, err := json.Marshal(todoList.GetTodos())
	err = os.WriteFile(path, bytes, 0644)

	if err != nil {
		return err
	}

	return nil
}
