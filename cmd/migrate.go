package cmd

import (
	"github.com/lnk2005/td_read/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrates the database schema to the latest version",
	Long:  `migrates the database schema to the latest version.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return db.CreateTables()
	},
}
