package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addTask = &cobra.Command{
	Use:     "add",
	Short:   "add new task",
	Aliases: []string{"a"},
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := getTask(args)
		taskStorage.AddNew(task)
	},
}

func init() {
	rootCmd.AddCommand(addTask)
}

func getTask(args []string) string {
	if len(args) == 0 {

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Provide task: ")

		taskString, err := reader.ReadString('\n')

		if err != nil {
			panic(err)
		}

		return strings.TrimSpace(taskString)
	} else {
		return args[1]
	}

}
