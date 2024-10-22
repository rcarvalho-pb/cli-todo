package commands

import (
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var title string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",

	Run: func(cmd *cobra.Command, args []string) {
		title := strings.Join(args, " ")
		if err := model.Task.AddTask(title); err != nil {
			color.Red("error adding new task:", err)
			return
		}

		model.Task.ListAllUnfinished()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
