package commands

import (
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var newTitle string

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit an unfinished task",

	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			color.Red("invalid id format")
			return
		}
		newTitle := strings.Join(args[1:], " ")
		if err := model.Task.UpdateTask(int64(id), newTitle); err != nil {
			color.Red("error editing task: %s\n", err.Error())
			return
		}

		model.Task.ListAllUnfinished()
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
