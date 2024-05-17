/*
Copyright © 2024 Roberto Antolin <rantolin DOT geo AT gmail DOT com>
*/
package list

import (
	"fmt"

	"github.com/spf13/cobra"
)

func runList(cmd *cobra.Command, args []string) {
	fmt.Printf("project: %s\n", cmd.Flag("project_id").Value)
	fmt.Println("list called")
}

// listCmd represents the list command
var ListCmd = &cobra.Command{
    Use:   "list",
    Short: "Lists BigQuery resources",
    Long: `The 'list' command is a parent command that lists various BigQuery resources.
It doesn't directly perform any operations but serves as a root for other subcommands.
Each subcommand corresponds to a specific resource type (e.g., tables, datasets) and lists all instances of that resource.
Use 'bqsql list [subcommand]' to list all instances of a specific resource.`,
	Run: runList,
}

func init() {
	// Here you will define your flags and configuration settings.
}
