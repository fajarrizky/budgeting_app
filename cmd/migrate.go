package cmd

import (
	"budgetapp/config"
	"budgetapp/module/db"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:     "migration",
	Aliases: []string{"migrate"},
	Short:   "nanocard migration",
	Run: func(cmd *cobra.Command, args []string) {
		configModule := config.NewConfigModule()

		configService := configModule.GetConfigService()

		db.RunMigrations(configService.GetDbConfig(), "migrations")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
