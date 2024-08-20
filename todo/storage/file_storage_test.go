package storage

import (
	"encoding/json"
	"os"
	"path"
	"reflect"
	"testing"

	"github.com/radek-nowak/go_todo_app/tests"
	todo "github.com/radek-nowak/go_todo_app/todo/model"
	"github.com/stretchr/testify/assert"
)

func newIntegrationTest(t *testing.T, testData []byte) string {
	tempDir := t.TempDir()

	config := Config{
		FileName: "test_tasks_data.json",
		FilePath: tempDir,
	}

	Init(config, false)
	testDataPath := path.Join(config.FilePath, config.FileName)
	os.WriteFile(testDataPath, []byte(testData), 0644)

	return testDataPath
}

func TestFindAll(t *testing.T) {

	validJson := `[
		{"ID": 1, "Task": "Task 1", "Done": false},
        {"ID": 2, "Task": "Task 2", "Done": true}
	]`

	tests := []struct {
		name                  string
		testData              []byte
		expectedNumberOfTasks int
	}{
		{
			name:                  "find 0 tasks",
			testData:              nil,
			expectedNumberOfTasks: 0,
		},

		{
			name:                  "find 2 tasks",
			testData:              []byte(validJson),
			expectedNumberOfTasks: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			newIntegrationTest(t, test.testData)
			storage := NewJsonFileStorage()

			// when
			tasks, err := storage.FindAll()

			// then
			assert.NoError(t, err)
			assert.NotNil(t, tasks)
			assert.Len(t, tasks.GetTodos(), test.expectedNumberOfTasks)
		})
	}
}

func TestFindTop(t *testing.T) {

	validJson := `[
		{"ID": 1, "Task": "Task 1", "Done": false},
        {"ID": 2, "Task": "Task 2", "Done": true}
	]`

	tests := []struct {
		name                  string
		testData              []byte
		top                   int
		expectedNumberOfTasks int
	}{
		{
			name:                  "find top 1 task, 0 tasks saved",
			testData:              nil,
			top:                   1,
			expectedNumberOfTasks: 0,
		},

		{
			name:                  "find top 1 task, 2 tasks saved",
			testData:              []byte(validJson),
			top:                   1,
			expectedNumberOfTasks: 1,
		},
		{
			name:                  "find top 3 task, 2 tasks saved",
			testData:              []byte(validJson),
			top:                   3,
			expectedNumberOfTasks: 2,
		},
		{
			name:                  "find top -1 task, 2 tasks saved",
			testData:              []byte(validJson),
			top:                   -1,
			expectedNumberOfTasks: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			newIntegrationTest(t, test.testData)
			storage := NewJsonFileStorage()

			// when
			tasks, err := storage.FindTop(test.top)

			// then
			assert.NoError(t, err)
			assert.NotNil(t, tasks)
			assert.Len(t, tasks.GetTodos(), test.expectedNumberOfTasks)
		})
	}
}

func TestAddNew(t *testing.T) {

	validJson := `[
		{"ID": 1, "Task": "Task 1", "Done": false},
        {"ID": 2, "Task": "Task 2", "Done": true}
	]`

	tests := []struct {
		name          string
		testData      []byte
		newTasks      []string
		expectedTasks *todo.Tasks
	}{
		{
			name:     "",
			testData: nil,
			newTasks: []string{},
			expectedTasks: todo.FromTodos([]todo.Todo{{
				Task: "",
				Done: false,
			}}),
		},

		{
			name:     "first taks",
			testData: nil,
			newTasks: []string{"first task"},
			expectedTasks: todo.FromTodos([]todo.Todo{{
				Task: "first task",
				Done: false,
			}}),
		},

		{
			name:     "third task",
			testData: []byte(validJson),
			newTasks: []string{"third task"},
			expectedTasks: todo.FromTodos([]todo.Todo{
				{
					Task: "Task 1",
					Done: false,
				},
				{
					Task: "Task 2",
					Done: true,
				},
				{
					Task: "third task",
					Done: false,
				},
			}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testDataPath := newIntegrationTest(t, test.testData)
			storage := NewJsonFileStorage()

			// when
			for _, newTask := range test.newTasks {
				storage.AddNew(newTask)
			}

			// then
			data, _ := os.ReadFile(testDataPath)
			var todos []todo.Todo
			json.Unmarshal(data, &todos)

			tasks := todo.FromTodos(todos)

			assert.ObjectsAreEqualValues(test.expectedTasks, tasks)
		})
	}
}

func TestDelete(t *testing.T) {

	tests := []struct {
		name          string
		testData      []byte
		id            int
		expectedTasks *todo.Tasks
		expectedError error
	}{
		{
			name:          "invalid task id err is returned when non existing task is deleted",
			testData:      nil,
			id:            1,
			expectedTasks: nil,
			expectedError: todo.ErrInvalidTaskId,
		},

		{
			name: "task is deleted",
			testData: []byte(
				`[
					{"ID": 1, "Task": "Task 1", "Done": false},
				    {"ID": 2, "Task": "Task 2", "Done": true}
				]`,
			),
			id: 2,
			expectedTasks: todo.FromTodos([]todo.Todo{{
				Task: "Task 1",
				Done: false,
			}}),
			expectedError: nil,
		},

		{
			name: "first task is deleted",
			testData: []byte(
				`[
					{"ID": 1, "Task": "Task 1", "Done": false},
				    {"ID": 2, "Task": "Task 2", "Done": true}
				]`,
			),
			id: 1,
			expectedTasks: todo.FromTodos([]todo.Todo{{
				Task: "Task 2",
				Done: true,
			}}),
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testDataPath := newIntegrationTest(t, test.testData)
			storage := NewJsonFileStorage()

			// when
			err := storage.Delete(test.id)

			// then
			if test.expectedError != nil {
				assert.Error(t, err)
				assert.ErrorIs(t, err, test.expectedError)
			} else {
				assert.NoError(t, err)

				data, _ := os.ReadFile(testDataPath)
				var todos []todo.Todo
				json.Unmarshal(data, &todos)

				tasks := todo.FromTodos(todos)

				assert.ObjectsAreEqualValues(test.expectedTasks, tasks)
			}
		})
	}
}

func TestDeleteRange(t *testing.T) {
	testCases := []struct {
		name          string
		idFrom        int
		idTo          int
		testData      []byte
		expectedTasks *todo.Tasks
		expectedError error
	}{

		{
			name:   "deletes one task when id from and to are equal",
			idFrom: 2,
			idTo:   2,
			testData: []byte(
				`[
					{"ID": 1, "Task": "Task 1", "Done": false},
				    {"ID": 2, "Task": "Task 2", "Done": true}
				]`,
			),
			expectedTasks: &todo.Tasks{},
			expectedError: nil,
		},

		{
			name:   "deletes tasks from a given range",
			idFrom: 1,
			idTo:   2,
			testData: []byte(
				`[
					{"ID": 1, "Task": "Task 1", "Done": false},
					{"ID": 2, "Task": "Task 2", "Done": false},
				    {"ID": 3, "Task": "Task 3", "Done": true}
				]`,
			),
			expectedTasks: todo.FromTodos([]todo.Todo{{
				Task: "Task 3",
				Done: false,
			}}),
			expectedError: nil,
		},

		{
			name:   "deletes all tasks to given 'to id' when non positive 'from id' passed",
			idFrom: 0,
			idTo:   2,
			testData: []byte(
				`[
					{"ID": 1, "Task": "Task 1", "Done": false},
					{"ID": 2, "Task": "Task 2", "Done": false},
				    {"ID": 3, "Task": "Task 3", "Done": true}
				]`,
			),
			expectedTasks: todo.FromTodos([]todo.Todo{{
				Task: "Task 3",
				Done: false,
			}}),
			expectedError: nil,
		},

		{
			name:   "deletes all tasks from given 'from id' when too large 'to id' passed",
			idFrom: 2,
			idTo:   100,
			testData: []byte(
				`[
					{"ID": 1, "Task": "Task 1", "Done": false},
					{"ID": 2, "Task": "Task 2", "Done": false},
				    {"ID": 3, "Task": "Task 3", "Done": true}
				]`,
			),
			expectedTasks: todo.FromTodos([]todo.Todo{{
				Task: "Task 1",
				Done: false,
			}}),
			expectedError: nil,
		},

		{
			name:          "",
			idFrom:        1,
			idTo:          1,
			testData:      nil,
			expectedTasks: &todo.Tasks{},
			expectedError: &todo.OutOfRangeError{
				Value: 1,
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			testDataPath := newIntegrationTest(t, test.testData)
			storage := NewJsonFileStorage()

			// when
			err := storage.DeleteRange(test.idFrom, test.idTo)

			// then
			if test.expectedError != nil {
				assert.Error(t, err)
				assert.ErrorIs(t, err, test.expectedError)
			} else {
				assert.NoError(t, err)

				data, _ := os.ReadFile(testDataPath)
				var todos []todo.Todo
				json.Unmarshal(data, &todos)

				tasks := todo.FromTodos(todos)

				assert.ObjectsAreEqualValues(test.expectedTasks, tasks)
			}
		})
	}
}

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
			got, err := readData(All)
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

var c = Config{
	FileName: "test_data.json",
	FilePath: "./todo_storage_test/",
}

func TestJsonFileStorageFind(t *testing.T) {

	Init(c, false)

	defer os.RemoveAll(c.FilePath)

	storage := NewJsonFileStorage()

	t.Run("FindAll should not reurn error when file is empty", func(t *testing.T) {
		_, err := storage.FindAll()
		assertErrorIsNil(t, err)
	})

	t.Run("FindAll", func(t *testing.T) {
		validJson := `[
			{"ID": 1, "Task": "Task 1", "Done": false},
    	    {"ID": 2, "Task": "Task 2", "Done": true}
		]`

		os.WriteFile(c.fullPath(), []byte(validJson), 0644)

		tasks, err := storage.FindAll()
		assertErrorIsNil(t, err)

		if len(tasks.GetTodos()) != 2 {
			t.Fatalf("expected")
		}
	})

}

func TestJsonFIleStorageFindTop(t *testing.T) {

	validJson := `[
			{"ID": 1, "Task": "Task 1", "Done": false},
    	    {"ID": 2, "Task": "Task 2", "Done": true}
		]`

	testCases := []tests.TestCase{
		{
			Name:     "Find top",
			FileName: "test_data.json",
			FilePath: "./todo_storage_test/",
			Data:     validJson,
			Err:      nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			c := Config{
				FileName: test.FileName,
				FilePath: test.FilePath,
			}

			Init(c, false)

			teardown := tests.NewIntegrationTest(test)
			defer teardown()

			storage := NewJsonFileStorage()
			tasks, err := storage.FindTop(1)

			assertErrorIsNil(t, err)

			if tasksLen := len(tasks.GetTodos()); tasksLen != 1 {
				t.Fatalf("expected to return 1 task, got %d", tasksLen)
			}

		})
	}

}

func assertErrorIsNil(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatalf("expected not to return error, got %q", got)
	}
}
