package main

import (
	todo "go_todo/todo/model"
	"go_todo/ui"
)

func main() {
	// reverser.Execute()
	todoList := todo.NewTodoList()
	todoList.Add("first task")
	todoList.Add("asdnladlasjldkjalskjdlkjasldkjlaskjdlkjasldkjajsldkjlajsdlkj")
	todoList.Add("first task")
	todoList.CompleteTask(2)
	ui.Display(*todoList)
}
