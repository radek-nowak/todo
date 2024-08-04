package cmd

import (
	"fmt"
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
		err := storage.PersistChanges(func(t todo.Tasks) (*todo.Tasks, error) {
			taskId, err := strconv.Atoi(args[0])
			if err != nil {
				return nil, err
			}
			err = t.UpdateTask(taskId, args[1])
			if err != nil {
				return nil, fmt.Errorf("unable to update the task, %v", err)
			}
			return &t, nil
		})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateTaskCmd)
}
