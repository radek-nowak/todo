package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var updateTaskCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"u"},
	Short:   "update task",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
		}
		err = taskStorage.Update(taskId, args[1])
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateTaskCmd)
}
