package cmd

import "github.com/spf13/cobra"

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Database commands to create migration and seeding",
	Run:   func(cmd *cobra.Command, args []string) {},
}
