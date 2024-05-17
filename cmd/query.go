/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/bigquery"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/api/iterator"
)

// QueryBasic demonstrates issuing a query and reading results.
func QueryBasic(w io.Writer, projectID string, query string) error {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %v", err)
	}
	defer client.Close()

	fmt.Println("query: ", query)
	q := client.Query(query)

	// Location must match that of the dataset(s) referenced in the query.
	q.Location = "US"

	// Run the query and print results when the query job is completed.
	job, err := q.Run(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}
	status, err := job.Wait(ctx)
	if err != nil {
		return err
	}
	if err := status.Err(); err != nil {
		return err
	}
	it, err := job.Read(ctx)
	for {
		var row []bigquery.Value
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Fprintln(w, row)
	}
	return nil
}

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
    Short: "Executes a BigQuery SQL query",
    Long: `The 'query' command allows you to execute a BigQuery SQL query directly from the command line.
You need to provide the project ID and the SQL query as arguments.
The command connects to the BigQuery client, runs the query, and prints the results.
It's a convenient way to interact with your BigQuery datasets without leaving your terminal.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		QueryBasic(
			cmd.OutOrStdout(),
			viper.GetString("project_id"),
			query,
		)
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
	queryCmd.SetArgs([]string{"query"})
}
