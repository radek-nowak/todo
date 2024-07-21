package cmd

import (
	todo "go_todo/todo/model"
	"go_todo/todo/storage"

	"github.com/spf13/cobra"
)

var addTask = &cobra.Command{
	Use:     "add",
	Short:   "add new task",
	Aliases: []string{"a"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := "./tasks.json"
		data, err := storage.ReadData(path)
		if err != nil {
			panic("Error in write data cmd during reading" + err.Error())
		}

		todos := todo.FromTodos(data.GetTodos())
		todos.Add(args[0])

		err = storage.WriteData(path, todos)
		if err != nil {
			panic("Error in write data cmd during writing" + err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(addTask)
}
