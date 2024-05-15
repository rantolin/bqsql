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

// listDatasets demonstrates iterating through the collection of datasets in a project.
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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
