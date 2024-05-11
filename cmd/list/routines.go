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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
