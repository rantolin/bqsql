/*
Copyright © 2024 Roberto Antolin <rantolin DOT geo AT gmail DOT com>
*/
package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func queryHead(w io.Writer, projectID string, dataset string, table string) error {
	query := fmt.Sprintf("SELECT * FROM `%s.%s.%s` LIMIT 10", projectID, dataset, table)
	QueryBasic(w, projectID, query)
	return nil
}

// headCmd represents the head command
var headCmd = &cobra.Command{
	Use:   "head",
	Short: "Fetches the first few records from a BigQuery table",
	Long: `The 'head' command fetches the first few records from a specified BigQuery table.
You need to provide the project ID, dataset, and table as arguments.
The command connects to the BigQuery client, runs a SELECT * query with a LIMIT of 10 to fetch the first few records, and prints the results.
It's a convenient way to quickly inspect the data in your BigQuery tables without leaving your terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectID := viper.GetString("project_id")
		dataset := viper.GetString("dataset")
		table := args[0]
		queryHead(
			cmd.OutOrStdout(),
			projectID,
			dataset,
			table,
		)
	},
}

func init() {
	rootCmd.AddCommand(headCmd)
}
