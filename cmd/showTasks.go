package cmd

import (
	"go_todo/todo/storage"
	"go_todo/ui"

	"github.com/spf13/cobra"
)

var showTaskCmd = &cobra.Command {
	Use: "show",
	Short: "shows tasks",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := storage.ReadData("./tasks.json")
		if err != nil {
			panic("Error occured in show task command")
		}

		ui.Display(*tasks)
	},
}

func init() {
	rootCmd.AddCommand(showTaskCmd)
}
