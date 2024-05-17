/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
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

func listRoutines(w io.Writer, projectID string, datasetID string) error {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %w", err)
	}
	defer client.Close()

	it := client.Dataset(datasetID).Routines(ctx)
	for {
		routine, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "%s\n", routine.RoutineID)
	}
	return nil
}

// routinesCmd represents the routines command
var routinesCmd = &cobra.Command{
    Use:   "routines",
    Short: "Lists all routines in a BigQuery dataset",
    Long: `The 'routines' subcommand lists all routines in a specified BigQuery dataset.
You need to provide the project ID and dataset as arguments.
The command connects to the BigQuery client, retrieves all routines in the dataset, and prints their names.
It's a convenient way to quickly inspect the routines in your BigQuery dataset without leaving your terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("routines called")

		projectID := viper.GetString("project_id")
		datasetID := viper.GetString("dataset")
		listRoutines(cmd.OutOrStdout(), projectID, datasetID)
	},
}

func init() {
	ListCmd.AddCommand(routinesCmd)
}
