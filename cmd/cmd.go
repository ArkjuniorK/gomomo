package cmd

import "github.com/spf13/cobra"

var (
	name string

	host string
	port string

  ext string
)

var rootCmd = &cobra.Command{Use: "app"}

func init() {
	rootCmd.AddCommand(apiCmd, genCmd, dbCmd)

	apiCmd.Flags().StringVar(&port, "port", "8080", "Port to run the API server")
	apiCmd.Flags().StringVar(&host, "host", "0.0.0.0", "Host to run the API server")

	genCmd.AddCommand(modelGenCmd, pkgGenCmd)
	pkgGenCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the package")
	modelGenCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the model")

  dbCmd.AddCommand(dbMigrateCmd, dbSeedCmd)
  dbMigrateCmd.AddCommand(dbMigrateCreateCmd, dbMigrateUpCmd, dbMigrateDownCmd)
  dbMigrateCreateCmd.Flags().StringVarP(&ext, "ext", "e", "sql", "Migration extension")
  dbMigrateCreateCmd.Flags().StringVarP(&name, "name", "n", "", "Name or title of migration")

	rootCmd.Execute()
}
