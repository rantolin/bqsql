/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

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
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
}

func initViper() {

	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName(".bqsqlrc")
		viper.SetConfigType("env")
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

	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file (default is $HOME/.bqsqlrc)")
	rootCmd.PersistentFlags().StringVarP(&project_id, "project_id", "p", "", "Project ID")
	rootCmd.PersistentFlags().StringVarP(&dataset, "dataset", "d", "", "Dataset name")

	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	viper.BindPFlag("project_id", rootCmd.PersistentFlags().Lookup("project_id"))
	viper.BindPFlag("dataset", rootCmd.PersistentFlags().Lookup("dataset"))

	addSubcommandPalletes()
}
