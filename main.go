package main

import (
	"go_todo/cmd"
	"go_todo/todo/storage"
)

func main() {
	storage.Init()

	cmd.Execute()
}
