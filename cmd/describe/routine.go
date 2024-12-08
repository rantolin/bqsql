/*
Copyright © 2024 Roberto Antolin <rantolin DOT geo AT gmail DOT com>
*/
package describe

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func printRoutineInfo(cmd *cobra.Command, routineID string) error {
	w := cmd.OutOrStdout()
	ctx := context.Background()
	projectID := viper.GetString("project_id")
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %w", err)
	}
	defer client.Close()

	datasetID := viper.GetString("dataset")
	meta, err := client.Dataset(datasetID).Routine(routineID).Metadata(ctx)
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "\nRoutine: %s\n", routineID)

	fullName, _ := cmd.Flags().GetBool("full_name")
	if fullName {
		fullName := client.Dataset(datasetID).Routine(routineID).FullyQualifiedName()
		fmt.Fprintf(w, "Full Name: %s\n", fullName)
	}

	body, _ := cmd.Flags().GetBool("body")
	if body {
		fmt.Fprintf(w, "Routine: %s\n", meta.Body)
	}

	fmt.Fprintf(w, "\nDescription: %s\n", meta.Description)

	argumnets, _ := cmd.Flags().GetBool("arguments")
	if argumnets {
		var args_format string = "   %s: %-10s\n"
		fmt.Fprintf(w, "\nArguments:\n")
		for _, input := range meta.Arguments {
			fmt.Fprintf(w, args_format, input.Name, input.DataType.TypeKind)
		}
	}

	returnType, _ := cmd.Flags().GetBool("return_type")
	if returnType {
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
	Short: "Describes a BigQuery routine",
	Long: `The 'routine' subcommand provides detailed information about a specified BigQuery routine.
You need to provide the project ID, dataset, and routine as arguments.
The command connects to the BigQuery client, retrieves the metadata of the routine, and prints it.
You can also use flags to get specific information like the body (--body), arguments (--arguments), return type (--return_type), or the fully qualified routine name (--full_name).
It's a convenient way to quickly inspect the details of your BigQuery routines without leaving your terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("routine called")

		routineID := args[0]
		printRoutineInfo(cmd, routineID)
	},
}

func init() {
	DescribeCmd.AddCommand(routineCmd)

	routineCmd.Flags().BoolP("body", "b", false, "Print routine body")
	routineCmd.Flags().BoolP("arguments", "a", false, "Print routine arguments")
	routineCmd.Flags().BoolP("return_type", "r", false, "Print routine return type")
	routineCmd.Flags().BoolP("full_name", "f", false, "Print full qualified routine name")
}
