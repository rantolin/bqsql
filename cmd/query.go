/*
Copyright © 2024 Roberto Antolin <rantolin DOT geo AT gmail DOT com>
*/
package cmd

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/bigquery"
	formats "github.com/rantolin/bqsql/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/api/iterator"
)

func getFinalQuery(query string, maxRows int) string {
	queryIsLimited := strings.Contains(strings.ToLower(query), "limit")
	queryNeedsLimit := !queryIsLimited && maxRows != 0
	if queryNeedsLimit {
		if strings.Contains(query, ";") {
			query = strings.TrimSuffix(query, ";")
		}
		query = fmt.Sprintf("%s LIMIT %d;", query, maxRows)
	}
	return query
}

// QueryBasic demonstrates issuing a query and reading results.
func QueryBasic(cmd *cobra.Command, query string) error {
	w := cmd.OutOrStdout()
	ctx := context.Background()

	projectID := viper.GetString("project_id")
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %v", err)
	}
	defer client.Close()

	maxRows, _ := cmd.Flags().GetInt("max_rows")
	query = getFinalQuery(query, maxRows)

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
	if err != nil {
		return err
	}

	schema := it.Schema.Relax()
	widths, err := formats.CalculateRowWidths(it, schema)
	if err != nil {
		return err
	}

	fmt.Fprintln(w)

	for i, field := range schema {
		fieldName := field.Name
		if i != 0 {
			fmt.Fprint(w, " | ")
		}
		fmt.Fprintf(w, "%-*v", widths[i], fieldName)
	}
	fmt.Fprintln(w)

	it, err = job.Read(ctx)
	if err != nil {
		return err
	}

	for {
		var row []bigquery.Value
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("iterator.Next: %v", err)
		}

		formats.PrintFormatedRow(w, row, widths)
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
		QueryBasic(cmd, query)
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
	queryCmd.SetArgs([]string{"query"})
	queryCmd.Flags().StringP("max_rows", "m", "100", "Maximum number of rows to return. If value is 0, all rows are returned. Default is 100. LIMIT clause takes precedence.")
}
