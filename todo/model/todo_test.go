package todo

import (
	"reflect"
	"testing"
)

func TestDisplayable(t *testing.T) {
	t.Run("Schema", func(t *testing.T) {
		// given
		tasks := Tasks{}
		want := []string{"ID", "TASK", "DONE"}

		// when
		schema := tasks.Schema()

		// then
		if schema == nil {
			t.Error("expected schema not to be nil")
		}

		if !reflect.DeepEqual(want, schema) {
			t.Errorf("expeced schema to be %v, got %v", want, schema)
		}
	})
}

func TestTasks(t *testing.T) {

	t.Run("Add", func(t *testing.T) {
		// given
		task := "new task"
		tasks := Tasks{}

		// when
		tasks.Add(task)

		// then
		if len(tasks.todos) != 1 {
			t.Error("expected to add one task, none added")
		}
		want := Todo{Task: task, Done: false}
		if tasks.todos[0] != want {
			t.Errorf("expected task to be %v, got %v", want, tasks.todos[0].Task)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		// given
		tasks := Tasks{
			todos: []Todo{{
				Task: "task 1",
				Done: false,
			}},
		}
		// when
		err := tasks.Delete(1)
		// then
		assertErrorIsNil(t, err)
		if len(tasks.todos) != 0 {
			t.Error("expected to delete task")
		}
	})

	t.Run("Delete non existing task", func(t *testing.T) {
		// given
		tasks := Tasks{
			todos: []Todo{{
				Task: "task 1",
				Done: false,
			}},
		}
		// when
		err := tasks.Delete(2)
		// then
		if len(tasks.todos) != 1 {
			t.Error("expected not to delete task")
		}
		assertError(t, err, ErrInvalidTaskId)
	})

	t.Run("Complete task", func(t *testing.T) {
		// given
		tasks := Tasks{
			todos: []Todo{{
				Task: "task 1",
				Done: false,
			}},
		}
		// when
		err := tasks.CompleteTask(1)
		// then
		assertErrorIsNil(t, err)
		if tasks.todos[0].Done != true {
			t.Error("expected to complete task")
		}
	})

	t.Run("Complete non existing task", func(t *testing.T) {
		// given
		tasks := Tasks{
			todos: []Todo{{
				Task: "task 1",
				Done: false,
			}},
		}
		// when
		err := tasks.CompleteTask(2)
		// then
		assertError(t, err, ErrInvalidTaskId)
	})

	t.Run("Complete already completed task", func(t *testing.T) {
		// given
		tasks := Tasks{
			todos: []Todo{{
				Task: "task 1",
				Done: true,
			}},
		}
		// when
		err := tasks.CompleteTask(1)
		// then
		assertError(t, err, ErrTaskAlreadyCompleted)
	})

	t.Run("Update task", func(t *testing.T) {
		// given
		tasks := Tasks{
			todos: []Todo{{
				Task: "task 1",
				Done: false,
			}},
		}
		// when
		newTask := "new task 1"
		err := tasks.UpdateTask(1, newTask)
		// then
		assertErrorIsNil(t, err)
		if tasks.todos[0].Task != newTask {
			t.Error("expected to update task")
		}
	})

	t.Run("Update non existing task", func(t *testing.T) {
		// given
		tasks := Tasks{
			todos: []Todo{{
				Task: "task 1",
				Done: false,
			}},
		}
		// when
		newTask := "new task 1"
		err := tasks.UpdateTask(2, newTask)
		// then
		assertError(t, err, ErrInvalidTaskId)
	})
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected error, but did not get one")
	}
	if got.Error() != want.Error() {
		t.Errorf("expected error message: %q, got: %q", got, want)
	}
}

func assertErrorIsNil(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("expected not to return error")
	}
}
