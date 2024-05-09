/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runList,
}

func init() {
	// Here you will define your flags and configuration settings.
}
