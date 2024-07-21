package storage

import (
	"encoding/json"
	"fmt"
	todo "go_todo/todo/model"
	"os"
	"reflect"
	"testing"
)

func TestReadData(t *testing.T) {

	validJson := `[
		{"ID": 1, "Task": "Task 1", "Done": false},
        {"ID": 2, "Task": "Task 2", "Done": true}
	]`

	validJsonPath := "data.json"

	os.WriteFile(validJsonPath, []byte(validJson), 0644)
	defer os.Remove(validJsonPath)

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    string
		want    *todo.Tasks
		wantErr bool
	}{
		{
			name: "Valid JSON",
			args: validJsonPath,
			want: todo.FromTodos([]todo.Todo{
				{Task: "Task 1", Done: false},
				{Task: "Task 2", Done: true},
			}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadData(tt.args)
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

	validJsonPath := "data.json"

	os.WriteFile(validJsonPath, []byte(""), 0644)
	defer os.Remove(validJsonPath)

	type args struct {
		path     string
		todoList *todo.Tasks
	}
	tests := []struct {
		name    string
		args    string
		want    *todo.Tasks
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Valid JSON",
			args: validJsonPath,
			want: todo.FromTodos([]todo.Todo{
				{Task: "Task 1", Done: false},
			}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeData(tt.args, tt.want); (err != nil) != tt.wantErr {
				t.Errorf("WriteData() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, _ := os.ReadFile(validJsonPath)
			fmt.Println(got)
			var actualTodos []todo.Todo
			json.Unmarshal(got, &actualTodos)
			actualTodoList := todo.FromTodos(actualTodos)

			if !reflect.DeepEqual(actualTodoList, tt.want) {
				t.Errorf("ReadData() = %v, want %v", got, tt.want)
			}
		})
	}
}
