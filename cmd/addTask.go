package cmd

import (
	model "github.com/radek-nowak/go_todo_app/todo/model"
	"github.com/radek-nowak/go_todo_app/todo/storage"

	"github.com/spf13/cobra"
)

var addTask = &cobra.Command{
	Use:     "add",
	Short:   "add new task",
	Aliases: []string{"a"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		storage.PersistChanges(func(t model.Tasks) (*model.Tasks, error) {
			t.Add(args[0])
			return &t, nil
		})
	},
}

func init() {
	rootCmd.AddCommand(addTask)
}
