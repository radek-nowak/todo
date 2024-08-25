package main

import (
	"github.com/radek-nowak/todo/cmd"
	"github.com/radek-nowak/todo/tasks/storage"
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
