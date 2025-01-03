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

const (
	DefaultProfileName = "default"
)

var project_id string
var dataset string
var configProfile string

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

type Configuration struct {
	Project string `mapstructure:"project"`
	Dataset string `mapstructure:"dataset"`
}
type Config struct {
	Configurations map[string]Configuration `mapstructure:"configurations"`
	Default        string                   `mapstructure:"default"`
}

func initViper() {
	cmd, _, _ := rootCmd.Find(os.Args[1:])
	isSetCommand := cmd != nil && cmd.Name() == "set"

	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.AddConfigPath(".")
	viper.SetConfigName("bqsql.yaml")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if !isSetCommand {
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		} else {
			fmt.Printf("\nWarning: Configuration file not found\n")
			fmt.Printf("  - Expected locations: ./bqsql.yaml and %s/bqsql.yaml\n", home)
			fmt.Printf("  - Error: %v\n", err)
			fmt.Printf("\nTo create a configuration file, run:\n")
			fmt.Printf("  bqsql set --project_id <your-project> --dataset <your-dataset> --profile %s\n\n", DefaultProfileName)
			fmt.Printf("Meanwhile, using values from flags and environment variables\n")
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		return
	}

	Profile := config.Configurations[config.Default]
	if configProfile != "" {
		Profile = config.Configurations[configProfile]
	}

	viper.Set("project_id", Profile.Project)
	viper.Set("dataset", Profile.Dataset)
}

func init() {
	cobra.OnInitialize(initViper)

	rootCmd.PersistentFlags().StringVarP(&configProfile, "profile", "", "", "Configuration profile (If not set, the 'default' profile will be used)")
	rootCmd.PersistentFlags().StringVarP(&project_id, "project_id", "", "", "Project ID")
	rootCmd.PersistentFlags().StringVarP(&dataset, "dataset", "", "", "Dataset name")

	viper.BindPFlag("project_id", rootCmd.PersistentFlags().Lookup("project_id"))
	viper.BindPFlag("dataset", rootCmd.PersistentFlags().Lookup("dataset"))
	viper.BindPFlag("profile", rootCmd.PersistentFlags().Lookup("profile"))

	addSubcommandPalletes()
}
