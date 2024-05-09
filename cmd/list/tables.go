/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package list

import (
	"context"
	"fmt"
	"io"
	"strings"

	"cloud.google.com/go/bigquery"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/api/iterator"
)

// listTables demonstrates iterating through the collection of tables in a given dataset.
func listTables(w io.Writer, projectID, datasetID string) error {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %w", err)
	}
	defer client.Close()

	ts := client.Dataset(datasetID).Tables(ctx)
	for {
		t, err := ts.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		if like_flag == "" || strings.Contains(t.TableID, like_flag) {
			fmt.Fprintf(w, "%s\n", t.TableID)
		}
	}
	return nil
}

// tablesCmd represents the tables command
var tablesCmd = &cobra.Command{
	Use:   "tables",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		listTables(
			cmd.OutOrStdout(),
			viper.GetString("project_id"),
			viper.GetString("dataset"),
		)
	},
}

var like_flag string

func init() {
	ListCmd.AddCommand(tablesCmd)

	tablesCmd.Flags().StringVarP(&like_flag, "like", "l", "", "String that table name should contain")
}
