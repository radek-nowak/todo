package todo

type Todo struct {
	ID   int    `json:"id,omitempty"`
	Task string `json:"task,omitempty"`
	Done bool   `json:"done,omitempty"`
}

type TodoList struct {
	todos  []Todo
	nextID int
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

	tl.todos[id-1].Done = true
}

// func (tl *TodoList) AsBytes()  {
// 	return []byte(tl.todos)
// 	
// }
