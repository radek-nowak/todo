package storage

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	todo "github.com/radek-nowak/go_todo_app/todo/model"
)

func TestReadData(t *testing.T) {

	validJson := `[
		{"ID": 1, "Task": "Task 1", "Done": false},
        {"ID": 2, "Task": "Task 2", "Done": true}
	]`

	dataStorageFilePath = "data.json"

	os.WriteFile(dataStorageFilePath, []byte(validJson), 0644)
	defer os.Remove(dataStorageFilePath)

	tests := []struct {
		name    string
		args    string
		want    *todo.Tasks
		wantErr bool
	}{
		{
			name: "Valid JSON",
			args: dataStorageFilePath,
			want: todo.FromTodos([]todo.Todo{
				{Task: "Task 1", Done: false},
				{Task: "Task 2", Done: true},
			}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadData(All)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteData(t *testing.T) {

	dataStorageFilePath = "test_data.json"

	os.WriteFile(dataStorageFilePath, []byte(""), 0644)
	defer os.Remove(dataStorageFilePath)

	tests := []struct {
		name    string
		args    string
		want    *todo.Tasks
		wantErr bool
	}{
		{
			name: "Valid JSON",
			args: dataStorageFilePath,
			want: todo.FromTodos([]todo.Todo{
				{Task: "Task 1", Done: false},
			}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeData(tt.want); (err != nil) != tt.wantErr {
				t.Errorf("WriteData() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, _ := os.ReadFile(dataStorageFilePath)
			var actualTodos []todo.Todo
			json.Unmarshal(got, &actualTodos)
			actualTodoList := todo.FromTodos(actualTodos)

			if !reflect.DeepEqual(actualTodoList, tt.want) {
				t.Errorf("ReadData() = %v, want %v", got, tt.want)
			}
		})
	}
}
