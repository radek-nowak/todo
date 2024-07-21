package cmd

import (
	todo "go_todo/todo/model"
	"go_todo/todo/storage"
	"strconv"

	"github.com/spf13/cobra"
)

var updateTaskCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"u"},
	Short:   "update task",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		path := "./tasks.json"
		storage.PersistChanges(path, func(tl todo.Tasks) (*todo.Tasks, error) {
			taskId, err := strconv.Atoi(args[0])
			if err != nil {
				panic("Failed to parse argument as an integer" + err.Error())
			}
			tl.UpdateTask(taskId, args[1])
			return &tl, nil
		})
	},
}

func init() {
	rootCmd.AddCommand(updateTaskCmd)
}
