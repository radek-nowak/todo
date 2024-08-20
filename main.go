package main

import (
	"github.com/radek-nowak/go_todo_app/cmd"
	"github.com/radek-nowak/go_todo_app/todo/storage"
)

var defaultDataStorageLocation = "/.todo_app/data/todo_data.json"

func main() {

	c := storage.Config{
		FileName: "todo_data.json",
		FilePath: "/.todo_app/data/",
	}

	storage.Init(c, true)

	cmd.Execute()
}
