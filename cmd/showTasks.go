package cmd

import (
	"go_todo/todo/storage"
	"go_todo/ui"

	"github.com/spf13/cobra"
)

const (
	maxItemsFlagName      = "top"
	maxItemsFlahShortName = "t"
)

var showTaskCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"s"},
	Short:   "shows tasks",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// todo error handling
		maxItems, _ := cmd.Flags().GetInt(maxItemsFlagName)
		showTasks(maxItems)
	},
}

func init() {
	rootCmd.AddCommand(showTaskCmd)
	showTaskCmd.PersistentFlags().IntP(maxItemsFlagName, maxItemsFlahShortName, 30, "Shows top x tasks.")
}

func showTasks(maxItems int) {
	tasks, err := storage.ReadData(maxItems)
	if err != nil {
		panic("Error occured in show task command")
	}

	ui.Display(tasks)
}
