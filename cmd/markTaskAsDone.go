package cmd

import (
	"fmt"
	"strconv"

	model "github.com/radek-nowak/go_todo_app/todo/model"
	"github.com/radek-nowak/go_todo_app/todo/storage"
	"github.com/spf13/cobra"
)

var completeTaskCmd = &cobra.Command{
	Use:     "complete",
	Short:   "Marks task as complete",
	Aliases: []string{"c"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := storage.PersistChanges(func(t model.Tasks) (*model.Tasks, error) {
			arg, err := strconv.Atoi(args[0])
			if err != nil {
				return nil, err
			}
			err = t.CompleteTask(arg)
			if err != nil {
				return nil, fmt.Errorf("unable to complete the task, %v", err)
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
