/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string = "0.0.1"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "A brief description of your command",
	Long: `This command provides the current version of the bqsql project.
The version follows semantic versioning standards.
Each release increments this version.
Use this command to check which version of bqsql you are currently using.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version: ", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
