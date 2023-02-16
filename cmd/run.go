package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run the reader to read data into database",
	Long:  `run the reader to read data into database`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
