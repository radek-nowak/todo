package storage

import (
	"encoding/json"
	todo "go_todo/todo/model"
	"os"
)

func ReadData(path string) (*todo.Tasks, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var todos []todo.Todo

	json.Unmarshal(data, &todos)

	return todo.FromTodos(todos), nil
}

func WriteData(path string, todoList *todo.Tasks) error {
	bytes, err := json.Marshal(todoList.GetTodos())
	err = os.WriteFile(path, bytes, 0644)

	if err != nil {
		return err
	}

	return nil
}

func PersistChanges(path string, operation func(todo.Tasks) (*todo.Tasks, error)) error {
	tasks, err := ReadData(path)
	if err != nil {
		return err
	}

	tasks, err = operation(*tasks)
	if err != nil {
		return err
	}

	err = WriteData(path, tasks)
	if err != nil {
		return err
	}

	return nil
}
