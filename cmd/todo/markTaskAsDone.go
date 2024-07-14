package todo

import (
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
		storage.PersistChanges(path, func(tl todo.TodoList) (*todo.TodoList, error) {
			arg, err := strconv.Atoi(args[0])
			if err != nil {
				panic("Failed to parse argument as an integer" + err.Error())
			}
			tl.CompleteTask(arg)
			return &tl, nil
		})
	},
}

func init() {
	rootCmd.AddCommand(completeTaskCmd)
}
