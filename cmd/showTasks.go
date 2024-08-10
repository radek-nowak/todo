package cmd

import (
	"fmt"

	"github.com/radek-nowak/go_todo_app/ui"
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
		maxItems, err := cmd.Flags().GetInt(maxItemsFlagName)
		if err != nil {
			fmt.Println(err)
		}

		err = showTasks(maxItems)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(showTaskCmd)
	showTaskCmd.PersistentFlags().IntP(maxItemsFlagName, maxItemsFlahShortName, 30, "Shows top x tasks.")
}

func showTasks(maxItems int) error {
	tasks, err := taskStorage.FindTop(maxItems)
	if err != nil {
		return err
	}

	ui.Display(tasks)
	return nil
}
