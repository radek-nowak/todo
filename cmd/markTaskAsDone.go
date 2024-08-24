package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var completeTaskCmd = &cobra.Command{
	Use:     "complete",
	Short:   "Marks task as complete",
	Aliases: []string{"c"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
		}

		err = taskStorage.Complete(id)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(completeTaskCmd)
}
