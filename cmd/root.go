/*
Copyright © 2024 Roberto Antolin <rantolin DOT geo AT gmail DOT com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/rantolin/bqsql/cmd/describe"
	"github.com/rantolin/bqsql/cmd/list"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var project_id string
var dataset string
var configFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bqsql",
	Short: "A CLI tool for interacting with BigQuery SQL",
	Long: `bqsql is a command-line interface (CLI) tool that simplifies the process of interacting with Google's BigQuery SQL.
It provides commands for describing, listing, and managing BigQuery datasets and tables.
You can use it to perform operations like listing all datasets in a project, describing a specific table, and more.
It is designed to be easy to use and flexible, making it a powerful tool for anyone working with BigQuery SQL.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubcommandPalletes() {
	rootCmd.AddCommand(list.ListCmd)
	rootCmd.AddCommand(describe.DescribeCmd)
}

func initViper() {

	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName("bqsql.yaml")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("Error reading config file:", err)
	}
	// viper.SetEnvPrefix("BQSQL")
	// viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func init() {
	cobra.OnInitialize(initViper)

	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file (default is $HOME/.bqsql/bqsql.yaml or ./bqsql.yaml)")
	rootCmd.PersistentFlags().StringVarP(&project_id, "project_id", "p", "", "Project ID")
	rootCmd.PersistentFlags().StringVarP(&dataset, "dataset", "d", "", "Dataset name")

	viper.BindPFlag("project_id", rootCmd.PersistentFlags().Lookup("project_id"))
	viper.BindPFlag("dataset", rootCmd.PersistentFlags().Lookup("dataset"))

	addSubcommandPalletes()
}
