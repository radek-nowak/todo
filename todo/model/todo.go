package todo

import "fmt"

type Todo struct {
	ID   int    `json:"id,omitempty"`
	Task string `json:"task,omitempty"`
	Done bool   `json:"done,omitempty"`
}

type TodoList struct {
	todos  []Todo
	nextID int
}

func (TodoList) Schema() []string {
	return []string{"ID", "TASK", "DONE"}
}

func (tl TodoList) Data() [][]interface{} {
	var data [][]interface{}

	for _, todo := range tl.todos {
		data = append(data, []interface{}{todo.ID, todo.Task, todo.Done})
	}

	return data
}

func (TodoList) ColumnWidths() []int {
	return []int{3, 55, 5}
}

func NewTodoList() *TodoList {
	return &TodoList{nextID: 1}
}

func FromTodos(todos []Todo) *TodoList {
	return &TodoList{todos: todos, nextID: len(todos) + 1}
}

func (tl *TodoList) GetTodos() []Todo {
	return tl.todos
}

func (tl *TodoList) Add(task string) {
	newTodo := Todo{
		ID:   tl.nextID,
		Task: task,
		Done: false,
	}

	tl.todos = append(tl.todos, newTodo)
	tl.nextID++
}

func (tl *TodoList) CompleteTask(id int) {
	if id < 1 || id >= tl.nextID {
		panic("id out of range")
	}

	if tl.todos[id-1].Done == true {
		fmt.Println("Already completed")
		return
	}

	tl.todos[id-1].Done = true
}
