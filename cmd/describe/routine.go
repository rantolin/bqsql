/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package describe

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/bigquery"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func printRoutineInfo(w io.Writer, projectID, datasetID, routineID string) error {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %w", err)
	}
	defer client.Close()

	meta, err := client.Dataset(datasetID).Routine(routineID).Metadata(ctx)
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "\nRoutine: %s\n", routineID)

	if viper.GetBool("full_name") {
		fullName := client.Dataset(datasetID).Routine(routineID).FullyQualifiedName()
		fmt.Fprintf(w, "Full Name: %s\n", fullName)
	}

	if viper.GetBool("body") {
		fmt.Fprintf(w, "Routine: %s\n", meta.Body)
	}

	fmt.Fprintf(w, "\nDescription: %s\n", meta.Description)

	if viper.GetBool("arguments") {
		var args_format string = "   %s: %-10s\n"
		fmt.Fprintf(w, "\nArguments:\n")
		for _, input := range meta.Arguments {
			fmt.Fprintf(w, args_format, input.Name, input.DataType.TypeKind)
		}
	}

	if viper.GetBool("return_type") {
		var returnType = "NULL"
		if meta.ReturnType != nil {
			returnType = meta.ReturnType.TypeKind
		}
		fmt.Fprintf(w, "Return Type: %s\n", returnType)
	}

	return nil
}

// routineCmd represents the routine command
var routineCmd = &cobra.Command{
	Use:   "routine",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("routine called")

		projectID := viper.GetString("project_id")
		datasetID := viper.GetString("dataset")
		routineID := args[0]
		if err := printRoutineInfo(cmd.OutOrStdout(), projectID, datasetID, routineID); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	DescribeCmd.AddCommand(routineCmd)

	routineCmd.Flags().BoolP("body", "b", false, "Print routine body")
	routineCmd.Flags().BoolP("arguments", "a", false, "Print routine arguments")
	routineCmd.Flags().BoolP("return_type", "r", false, "Print routine return type")
	routineCmd.Flags().BoolP("full_name", "f", false, "Print full qualified routine name")

	if err := viper.BindPFlag("body", routineCmd.Flags().Lookup("body")); err != nil {
		fmt.Println(err)
	}
	if err := viper.BindPFlag("arguments", routineCmd.Flags().Lookup("arguments")); err != nil {
		fmt.Println(err)
	}
	if err := viper.BindPFlag("return_type", routineCmd.Flags().Lookup("return_type")); err != nil {
		fmt.Println(err)
	}
	if err := viper.BindPFlag("full_name", routineCmd.Flags().Lookup("full_name")); err != nil {
		fmt.Println(err)
	}
}
