package storage

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	model "github.com/radek-nowak/go_todo_app/todo/model"
)

const (
	dataStorageLocationEnvVar = "TODO_DATA"
	All                       = -1
)

var dataStorageFilePath string
var defaultDataStorageLocation = "/.data/todo_data.json"

func Init() {
	storageLocationEnvVar, exists := os.LookupEnv(dataStorageLocationEnvVar)
	if exists {
		dataStorageFilePath = storageLocationEnvVar
		return
	}

	// if env var is not set, then use the default location starting at home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Unable to get home directory. " + err.Error())
	}
	defaultDataStorageLocation = homeDir + defaultDataStorageLocation

	dataStorageDir := filepath.Dir(defaultDataStorageLocation)
	createDir(dataStorageDir)

	createFile(defaultDataStorageLocation)
	dataStorageFilePath = defaultDataStorageLocation
}

func ReadData(maxItems int) (*model.Tasks, error) {
	data, err := os.ReadFile(dataStorageFilePath)
	if err != nil {
		return nil, err
	}

	var todos []model.Todo
	json.Unmarshal(data, &todos)

	todos = topTasks(maxItems, todos)

	return model.FromTodos(todos), nil
}

func topTasks(maxItems int, todos []model.Todo) []model.Todo {
	if maxItems > 0 {
		if maxItems > len(todos) {
			maxItems = len(todos)
		}
		todos = todos[:maxItems]
	}
	return todos
}

func writeData(todoList *model.Tasks) error {
	bytes, err := json.Marshal(todoList.GetTodos())
	if err != nil {
		return err
	}

	err = os.WriteFile(dataStorageFilePath, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func PersistChanges(operation func(model.Tasks) (*model.Tasks, error)) error {
	tasks, err := ReadData(All)
	if err != nil {
		return err
	}

	tasks, err = operation(*tasks)
	if err != nil {
		return err
	}

	err = writeData(tasks)
	if err != nil {
		return err
	}

	return nil
}

func createFile(name string) {
	if _, err := os.Stat(name); errors.Is(err, os.ErrNotExist) {
		err = os.WriteFile(name, []byte(""), 0644)
		if err != nil {
			panic("Unable to set up storage file." + err.Error())
		}
	}
}

func createDir(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			panic("Unable to create strage directory. " + err.Error())
		}
	}
}
