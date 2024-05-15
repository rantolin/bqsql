/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package describe

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/bigquery"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// printTableInfo demonstrates fetching metadata from a table and printing some basic information
// to an io.Writer.
func printTableInfo(w io.Writer, projectID, datasetID, tableID string) error {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %w", err)
	}
	defer client.Close()

	meta, err := client.Dataset(datasetID).Table(tableID).Metadata(ctx)
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "\nTable: %s\n", meta.FullID)
	fmt.Fprintf(w, "Description: %s\n", meta.Description)

	if viper.GetBool("num_rows") {
		fmt.Fprintf(w, "Rows in managed storage: %d\n", meta.NumRows)
	}

	// Print basic information about the table's columns.
	if viper.GetBool("schema") {
		var columns_format string = "%-20s %-10s\n"
		fmt.Fprintf(w, "\n"+columns_format, "column_name", "type")
		for _, fieldSchema := range meta.Schema {
			fmt.Fprintf(w, columns_format, fieldSchema.Name, fieldSchema.Type)
			// fmt.Fprintf(w, "Description: %s\n", fieldSchema.Description)
		}
	}
	return nil
}

// tableCmd represents the table command
var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("table called")
		table := args[0]
		projectID := viper.GetString("project_id")
		datasetID := viper.GetString("dataset")
		if err := printTableInfo(cmd.OutOrStdout(), projectID, datasetID, table); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	DescribeCmd.AddCommand(tableCmd)

	tableCmd.Flags().BoolP("num_rows", "n", false, "Show number of rows")
	tableCmd.Flags().BoolP("schema", "s", false, "Show schema")

	if err := tableCmd.MarkFlagRequired("num_rows"); err != nil {
		fmt.Println(err)
	}
	if err := tableCmd.MarkFlagRequired("schema"); err != nil {
		fmt.Println(err)
	}
}
