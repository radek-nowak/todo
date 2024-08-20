package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	model "github.com/radek-nowak/go_todo_app/todo/model"
)

const (
	dataStorageLocationEnvVar = "TODO_DATA"
	All                       = -1
)

var dataStorageFilePath string

func Init(config Config, fromHomeDir bool) {
	homeDir, err := getHome(fromHomeDir)
	if err != nil {
		panic(err)
	}

	dataStoragePath := path.Join(homeDir, config.FilePath, config.FileName)

	dataStorageDir := filepath.Dir(dataStoragePath)
	createDir(dataStorageDir)

	dataStorageFile := dataStoragePath

	createFile(dataStorageFile)
	dataStorageFilePath = dataStorageFile
}

// todo hide to config
func getHome(get bool) (string, error) {
	if !get {
		return "", nil
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir, nil
}

type JsonFileStorage struct {
	path *string
}

func NewJsonFileStorage() *JsonFileStorage {
	return &JsonFileStorage{path: &dataStorageFilePath}
}

func (j *JsonFileStorage) FindAll() (*model.Tasks, error) {
	tasks, err := j.FindTop(All)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (j *JsonFileStorage) FindTop(maxItems int) (*model.Tasks, error) {
	data, err := os.ReadFile(dataStorageFilePath)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return model.FromTodos([]model.Todo{}), nil
	}

	var todos []model.Todo

	err = json.Unmarshal(data, &todos)
	if err != nil {
		return nil, err
	}

	todos = topTasks(maxItems, todos)

	return model.FromTodos(todos), nil
}

func (j *JsonFileStorage) AddNew(task string) {
	_ = persistChanges(func(t model.Tasks) (*model.Tasks, error) {
		t.Add(task)
		return &t, nil
	})
}

func (j *JsonFileStorage) Delete(taskId int) error {
	return persistChanges(func(t model.Tasks) (*model.Tasks, error) {
		err := t.Delete(taskId)
		if err != nil {
			return nil, fmt.Errorf("unable to delete the task %w", err)
		}

		return &t, nil
	})
}

func (j *JsonFileStorage) DeleteRange(from int, to int) error {
	return persistChanges(func(t model.Tasks) (*model.Tasks, error) {
		err := t.DeleteRange(from, to)
		if err != nil {
			return nil, fmt.Errorf("unable to delete the task from range %w", err)
		}
		return &t, nil
	})
}

func (j *JsonFileStorage) Complete(taksId int) error {
	return persistChanges(func(t model.Tasks) (*model.Tasks, error) {
		err := t.CompleteTask(taksId)
		if err != nil {
			return nil, fmt.Errorf("unable to complete the task, %v", err)
		}
		return &t, nil
	})
}

func (j *JsonFileStorage) Update(taskId int, task string) error {
	return persistChanges(func(t model.Tasks) (*model.Tasks, error) {
		err := t.UpdateTask(taskId, task)
		if err != nil {
			return nil, fmt.Errorf("unable to update the task, %v", err)
		}
		return &t, nil
	})
}

func readData(maxItems int) (*model.Tasks, error) {
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

func persistChanges(operation func(model.Tasks) (*model.Tasks, error)) error {
	tasks, err := readData(All)
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
