package storage

import (
	"encoding/json"
	"errors"
	model "go_todo/todo/model"
	"os"
	"path/filepath"
)

const dataStorageLocationEnvVar = "TODO_DATA"

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

func ReadData() (*model.Tasks, error) {
	data, err := os.ReadFile(dataStorageFilePath)
	if err != nil {
		return nil, err
	}

	var todos []model.Todo

	json.Unmarshal(data, &todos)

	return model.FromTodos(todos), nil
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
	tasks, err := ReadData()
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
