package cmd

import (
	"fmt"
	todo "go_todo/todo/model"
	"go_todo/todo/storage"
	"strconv"

	"github.com/spf13/cobra"
)

var completeTaskCmd = &cobra.Command{
	Use:     "complete",
	Short:   "Marks task as complete",
	Aliases: []string{"c"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := "./tasks.json"
		err := storage.PersistChanges(path, func(t todo.Tasks) (*todo.Tasks, error) {
			arg, err := strconv.Atoi(args[0])
			if err != nil {
				return nil, err
			}
			err = t.CompleteTask(arg)
			if err != nil {
				return nil, fmt.Errorf("Unable to complete the task, %q", err)
			}
			return &t, nil
		})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(completeTaskCmd)
}
