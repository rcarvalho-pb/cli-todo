package commands

import (
	"github.com/rcarvalho-pb/cli-todo/internal/models"
	"github.com/spf13/cobra"
)

var model *models.Models

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A Task management cli tool",
}

func Execute(m *models.Models) {
	model = m
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
}
