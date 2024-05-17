/*
Copyright © 2024 Roberto Antolin <rantolin DOT geo AT gmail DOT com>
*/
package list

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/bigquery"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/api/iterator"
)

// listDatasets demonstrates iterating through the collection of datasets in a project.
// func listDatasets(projectID string, w io.Writer) error {
func listDatasets(projectID string, w io.Writer) error {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)

	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %v", err)
	}
	defer client.Close()

	it := client.Datasets(ctx)

	for {
		dataset, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "%s\n", dataset.DatasetID)
	}
	return nil
}

// datasetsCmd represents the datasets command
var datasetsCmd = &cobra.Command{
    Use:   "datasets",
    Short: "Lists all datasets in a BigQuery project",
    Long: `The 'datasets' subcommand lists all datasets in a specified BigQuery project.
You need to provide the project ID as an argument.
The command connects to the BigQuery client, retrieves all datasets in the project, and prints their names.
It's a convenient way to quickly inspect the datasets in your BigQuery project without leaving your terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		listDatasets(
			viper.GetString("project_id"),
			cmd.OutOrStdout(),
		)
	},
}

func init() {
	ListCmd.AddCommand(datasetsCmd)
}
