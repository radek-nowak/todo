package cmd

import (
	"bufio"
	"fmt"
	todo "go_todo/todo/model"
	"go_todo/todo/storage"
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

		err := storage.PersistChanges(func(tl todo.Tasks) (*todo.Tasks, error) {

			if rangeFlagChanged {

				lowerRange, err := cmd.Flags().GetInt(lowerRangeFlagName)
				if err != nil {
					panic(err.Error())
				}

				upperRange, err := cmd.Flags().GetInt(upperRangeFlagName)
				if err != nil {
					panic(err.Error())
				}

				err = tl.DeleteRange(lowerRange, upperRange)
				if err != nil {
					return nil, err
				}

				return &tl, nil

			} else {
				taskId := getTaskId(args)
				err := tl.Delete(taskId)
				if err != nil {
					return nil, fmt.Errorf("unable to delete the task %v", err)
				}

				return &tl, nil
			}

		})
		if err != nil {
			fmt.Println(err)
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
