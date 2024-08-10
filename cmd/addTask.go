package cmd

import (
	"github.com/spf13/cobra"
)

var addTask = &cobra.Command{
	Use:     "add",
	Short:   "add new task",
	Aliases: []string{"a"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskStorage.AddNew(args[0])
	},
}

func init() {
	rootCmd.AddCommand(addTask)
}
