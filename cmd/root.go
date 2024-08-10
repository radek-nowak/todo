package cmd

import (
	"fmt"
	"os"

	"github.com/radek-nowak/go_todo_app/todo/storage"
	"github.com/spf13/cobra"
)

var taskStorage storage.Storage = storage.New()

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "todo cli",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "an error occured when executing root command %s", err)
		os.Exit(1)
	}
}
