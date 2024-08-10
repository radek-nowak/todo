package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

const (
	lowerRangeFlagName      = "from"
	lowerRangeFlagShortName = "f"

	upperRangeFlagName      = "to"
	upperRangeFlagShortName = "t"
)

var deleteTaskCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "deletes a task",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		rangeFlagChanged := cmd.Flags().Changed(lowerRangeFlagName) || cmd.Flags().Changed(upperRangeFlagName)

		if len(args) == 1 && rangeFlagChanged {
			fmt.Println("please provide either task id or range of ids")
			return
		}

		var cmdErr error

		if rangeFlagChanged {

			lowerRange, err := cmd.Flags().GetInt(lowerRangeFlagName)
			if err != nil {
				panic(err)
			}

			upperRange, err := cmd.Flags().GetInt(upperRangeFlagName)
			if err != nil {
				panic(err)
			}

			cmdErr = taskStorage.DeleteRange(lowerRange, upperRange)

		} else {
			taskId := getTaskId(args)
			cmdErr = taskStorage.Delete(taskId)
		}

		if cmdErr != nil {
			panic(cmdErr)
		}

	},
}

func init() {
	rootCmd.AddCommand(deleteTaskCmd)
	deleteTaskCmd.PersistentFlags().IntP(lowerRangeFlagName, lowerRangeFlagShortName, -1, "Deletes tasks from index")
	deleteTaskCmd.PersistentFlags().IntP(upperRangeFlagName, upperRangeFlagShortName, -1, "Deletes tasks to index")
}

func getTaskId(args []string) (taskId int) {
	if len(args) == 0 {
		fmt.Print("Delete id:")
		reader := bufio.NewReader(os.Stdin)
		taskIdString, err := reader.ReadString('\n')
		if err != nil {
			panic(err.Error())
		}
		taskIdString = strings.TrimSpace(taskIdString)
		taskId, err = strconv.Atoi(taskIdString)
		if err != nil {
			panic(err.Error())
		}

	} else {
		var convErr error
		taskId, convErr = strconv.Atoi(args[0])
		if convErr != nil {
			panic("Failed to parse argument as an integer" + convErr.Error())
		}
	}

	return taskId
}
