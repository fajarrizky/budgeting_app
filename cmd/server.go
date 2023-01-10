package cmd

import (
	"budgetapp/app"
	"github.com/spf13/cobra"
)

var apiServerCmd = &cobra.Command{
	Use:     "api-server",
	Aliases: []string{"api"},
	Short:   "budgetapp API server",
	Run: func(cmd *cobra.Command, args []string) {
		app.StartApi()
	},
}

func init() {
	rootCmd.AddCommand(apiServerCmd)
}
