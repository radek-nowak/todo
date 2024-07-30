package cmd

import (
	"fmt"
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
		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			panic("Failed to parse argument as an integer" + err.Error())
		}
		err = storage.PersistChanges(func(tl todo.Tasks) (*todo.Tasks, error) {
			err := tl.Delete(taskId)
			if err != nil {
				return nil, fmt.Errorf("Unable to delete the task %v", err)
			}
			return &tl, nil
		})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteTaskCmd)
}
