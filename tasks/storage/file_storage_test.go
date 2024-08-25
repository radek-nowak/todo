package storage_test

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"testing"

	model "github.com/radek-nowak/todo/tasks/model"
	"github.com/radek-nowak/todo/tasks/storage"
	"github.com/stretchr/testify/assert"
)

type integrationTest struct {
	storage      *storage.JsonFileStorage
	testFilePath string
}

func newIntegrationTest(t *testing.T) *integrationTest {
	tmpDir := t.TempDir()

	tmpFile := "todo_test.json"

	config := storage.Config{FilePath: tmpDir, FileName: tmpFile}
	storage.Init(config, false)

	return &integrationTest{
		storage:      storage.NewJsonFileStorage(),
		testFilePath: path.Join(tmpDir, tmpFile),
	}
}

func (it *integrationTest) readFile() ([]model.Todo, error) {
	data, err := os.ReadFile(it.testFilePath)
	if err != nil {
		return nil, err
	}

	var todos []model.Todo
	err = json.Unmarshal(data, &todos)
	return todos, err
}

func TestJsonFileStorage_AddNewAndPersist(t *testing.T) {
	it := newIntegrationTest(t)

	it.storage.AddNew("Task 1")
	it.storage.AddNew("Task 2")

	// Verify tasks persisted to file
	todos, err := it.readFile()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(todos))
	assert.Equal(t, "Task 1", todos[0].Task)
	assert.Equal(t, "Task 2", todos[1].Task)
}

func TestJsonFileStorage_DeleteAndPersist(t *testing.T) {
	it := newIntegrationTest(t)

	it.storage.AddNew("Task 1")
	it.storage.AddNew("Task 2")
	_, err := it.readFile()
	assert.NoError(t, err)

	err = it.storage.Delete(1)
	assert.NoError(t, err)

	// Verify tasks persisted to file
	todos, err := it.readFile()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(todos))
	assert.Equal(t, "Task 2", todos[0].Task)
}

func TestJsonFileStorage_UpdateAndPersist(t *testing.T) {
	it := newIntegrationTest(t)

	it.storage.AddNew("Task 1")
	todos, err := it.readFile()
	assert.NoError(t, err)

	err = it.storage.Update(1, "Updated Task 1")
	assert.NoError(t, err)

	// Verify tasks persisted to file
	todos, err = it.readFile()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(todos))
	assert.Equal(t, "Updated Task 1", todos[0].Task)
}

func TestJsonFileStorage_CompleteAndPersist(t *testing.T) {
	it := newIntegrationTest(t)

	it.storage.AddNew("Task 1")
	todos, err := it.readFile()
	assert.NoError(t, err)

	err = it.storage.Complete(1)
	assert.NoError(t, err)

	// Verify tasks persisted to file
	todos, err = it.readFile()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(todos))
	assert.True(t, todos[0].Done)
}

func TestJsonFileStorage_DeleteRangeAndPersist(t *testing.T) {
	it := newIntegrationTest(t)

	it.storage.AddNew("Task 1")
	it.storage.AddNew("Task 2")
	it.storage.AddNew("Task 3")

	err := it.storage.DeleteRange(1, 2)
	assert.NoError(t, err)

	// Verify tasks persisted to file
	todos, err := it.readFile()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(todos))
	assert.Equal(t, "Task 3", todos[0].Task)
}

func TestJsonFileStorage_DeleteInvalidID(t *testing.T) {
	it := newIntegrationTest(t)

	it.storage.AddNew("Task 1")
	err := it.storage.Delete(99) // Invalid ID
	assert.ErrorIs(t, err, model.ErrInvalidTaskId)

	// Verify no tasks were deleted
	todos, err := it.readFile()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(todos))
}

func TestJsonFileStorage_CompleteAlreadyCompletedTask(t *testing.T) {
	it := newIntegrationTest(t)

	it.storage.AddNew("Task 1")
	err := it.storage.Complete(1)
	assert.NoError(t, err)

	// Attempt to complete the task again
	err = it.storage.Complete(1)
	assert.ErrorIs(t, err, model.ErrTaskAlreadyCompleted)
}

func TestJsonFileStorage_UpdateInvalidID(t *testing.T) {
	it := newIntegrationTest(t)

	err := it.storage.Update(99, "Updated Task") // Invalid ID
	if assert.Error(t, err, "should return an error for invalid task ID") {
		var expectedError *model.OutOfRangeError
		assert.True(t, errors.As(err, &expectedError), "error should be of type OutOfRangeError")
		assert.Equal(t, 99, expectedError.Value, "the OutOfRangeError should have the value 99")
	}
}

func TestJsonFileStorage_CompleteInvalidID(t *testing.T) {
	it := newIntegrationTest(t)

	err := it.storage.Complete(99) // Invalid ID
	if assert.Error(t, err, "should return an error for invalid task ID") {
		var expectedError *model.OutOfRangeError
		assert.True(t, errors.As(err, &expectedError), "error should be of type OutOfRangeError")
		assert.Equal(t, 99, expectedError.Value, "the OutOfRangeError should have the value 99")
	}
}

