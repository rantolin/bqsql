/*
Copyright © 2024 Roberto Antolin <rantolin DOT geo AT gmail DOT com>
*/
package describe

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// printTableInfo demonstrates fetching metadata from a table and printing some basic information
// to an io.Writer.
func printTableInfo(cmd *cobra.Command, tableID string) error {
	ctx := context.Background()
	w := cmd.OutOrStdout()

	projectID := viper.GetString("project_id")
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %w", err)
	}
	defer client.Close()

	datasetID := viper.GetString("dataset")
	meta, err := client.Dataset(datasetID).Table(tableID).Metadata(ctx)
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "\nTable: %s\n", meta.FullID)
	fmt.Fprintf(w, "Description: %s\n", meta.Description)

	numRows, _ := cmd.Flags().GetBool("num_rows")
	if numRows {
		fmt.Fprintf(w, "Rows in managed storage: %d\n", meta.NumRows)
	}

	// Print basic information about the table's columns.
	schema, _ := cmd.Flags().GetBool("schema")
	if schema {
		var columns_format string = "%-20s %-10s\n"
		fmt.Fprintf(w, "\n"+columns_format, "column_name", "type")
		for _, fieldSchema := range meta.Schema {
			fmt.Fprintf(w, columns_format, fieldSchema.Name, fieldSchema.Type)
		}
	}
	return nil
}

// tableCmd represents the table command
var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "Describes a BigQuery table",
	Long: `The 'table' subcommand provides detailed information about a specified BigQuery table.
You need to provide the project ID, dataset, and table as arguments.
The command connects to the BigQuery client, retrieves the metadata of the table, and prints it.
You can also use flags to get specific information like the number of rows (--num_rows) or the schema (--schema).
It's a convenient way to quickly inspect the details of your BigQuery tables without leaving your terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("table called")
		table := args[0]
		printTableInfo(cmd, table)
	},
}

func init() {
	DescribeCmd.AddCommand(tableCmd)

	tableCmd.Flags().BoolP("num_rows", "n", false, "Show number of rows")
	tableCmd.Flags().BoolP("schema", "s", false, "Show schema")
}
