package commands

import (
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a task",

	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			color.Red("invalid id format")
			return
		}
		if err := model.Task.DeleteTaks(int64(id)); err != nil {
			color.Red("error deleting task:", err)
			return
		}

		model.Task.ListAllUnfinished()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
