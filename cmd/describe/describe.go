/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package describe

import (
	"fmt"

	"github.com/spf13/cobra"
)

// describeCmd represents the describe command
var DescribeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Provides detailed information about BigQuery resources",
	Long: `The 'describe' command is a parent command that provides detailed information about various BigQuery resources.
It doesn't directly perform any operations but serves as a root for other subcommands.
Each subcommand corresponds to a specific resource type (e.g., table, dataset) and retrieves detailed information about it.
Use 'bqsql describe [subcommand]' to get information about a specific resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("describe called")
	},
}

func init() {
	// Here you will define your flags and configuration settings.
}
