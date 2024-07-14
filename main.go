package main

import (
	todo "go_todo/todo/model"
)

func main() {
	// reverser.Execute()
	todoList := todo.NewTodoList()
	todoList.Add("first task")
	todoList.Add("asdnladlasjldkjalskjdlkjasldkjlaskjdlkjasldkjajsldkjlajsdlkj")
	todoList.Add("first task")
	todoList.CompleteTask(2)
	// ui.Display(*todoList)
}
