package commands

import (
	"github.com/spf13/cobra"
)

var all bool
var finished bool
var sourceTitle string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",

	Run: func(cmd *cobra.Command, args []string) {
		if sourceTitle != "" {
			model.Task.ListAllTasksByTitle(sourceTitle)
			return
		}

		if all {
			model.Task.ListAll()
			return
		}

		if finished {
			model.Task.ListAllFinished()
			return
		}

		model.Task.ListAllUnfinished()
	},
}

func init() {
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "show all tasks")
	listCmd.Flags().BoolVarP(&finished, "finished", "f", false, "show all finished tasks")
	listCmd.Flags().StringVarP(&sourceTitle, "title", "t", "", "show all tasks with matching title")
	rootCmd.AddCommand(listCmd)
}
