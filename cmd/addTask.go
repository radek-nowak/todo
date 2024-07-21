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
		storage.PersistChanges(path, func(t todo.Tasks) (*todo.Tasks, error) {
			t.Add(args[0])
			return &t, nil
		})
	},
}

func init() {
	rootCmd.AddCommand(addTask)
}
