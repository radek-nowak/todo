package main

import (
	"github.com/radek-nowak/go_todo_app/cmd"
	"github.com/radek-nowak/go_todo_app/todo/storage"
)

func main() {
	storage.Init()

	cmd.Execute()
}
