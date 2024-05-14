/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("head called")

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
