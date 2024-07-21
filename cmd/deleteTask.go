package cmd

import (
	todo "go_todo/todo/model"
	"go_todo/todo/storage"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteTaskCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "deletes a task",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := "./tasks.json"
		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			panic("Failed to parse argument as an integer" + err.Error())
		}
		storage.PersistChanges(path, func(tl todo.Tasks) (*todo.Tasks, error) {
			tl.Delete(taskId)
			return &tl, nil
		})
	},
}

func init() {
	rootCmd.AddCommand(deleteTaskCmd)
}
