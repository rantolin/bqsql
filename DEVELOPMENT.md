# bqsql Developer Tutorial

This guide provides a reference for developers who want to extend `bqsql` by adding new features, commands, or modules.

## Project Structure Recap

- `main.go`: The entry point. It calls `cmd.Execute()`.
- `cmd/`: Contains all command logic.
  - `root.go`: Defines the base `bqsql` command and handles global flags and configuration (Viper).
  - `[command].go`: Top-level commands (e.g., `query.go`, `head.go`).
  - `[sub-package]/`: Groups of sub-commands (e.g., `list/`, `describe/`).
- `utils/`: Helper functions, specifically for output formatting.

---

## 1. Adding a New Top-Level Command

To add a new command like `bqsql mycommand`:

1.  **Create a new file** in `cmd/`, e.g., `cmd/mycommand.go`.
2.  **Define the command** using Cobra:

```go
package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var myCommand = &cobra.Command{
    Use:   "mycommand",
    Short: "Brief description of mycommand",
    Long:  `A longer description of what mycommand does.`,
    Run: func(cmd *cobra.Command, args []string) {
        // Your logic here
        fmt.Println("Running mycommand...")
    },
}

func init() {
    rootCmd.AddCommand(myCommand)
    // Define local flags here if needed:
    // myCommand.Flags().StringP("name", "n", "", "Description")
}
```

---

## 2. Adding a Sub-Command

To add a sub-command to an existing group, like `bqsql list myresource`:

1.  **Create a file** in the sub-package directory, e.g., `cmd/list/myresource.go`.
2.  **Define and register** the command:

```go
package list

import (
    "fmt"
    "github.com/spf13/cobra"
)

var myResourceCmd = &cobra.Command{
    Use:   "myresource",
    Short: "Lists my resources",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Listing my resources...")
    },
}

func init() {
    // Register it to the parent command defined in cmd/list/list.go
    ListCmd.AddCommand(myResourceCmd)
}
```

---

## 3. Working with Configuration (Viper)

`bqsql` uses Viper to manage settings like `project_id` and `dataset`. These can come from a config file, environment variables, or persistent flags.

Always retrieve these values using `viper`:

```go
import "github.com/spf13/viper"

projectID := viper.GetString("project_id")
datasetID := viper.GetString("dataset")
```

---

## 4. Interacting with BigQuery

Standard pattern for BigQuery operations:

```go
import (
    "context"
    "cloud.google.com/go/bigquery"
    "github.com/spf13/viper"
)

func myBigQueryLogic() error {
    ctx := context.Background()
    projectID := viper.GetString("project_id")
    
    client, err := bigquery.NewClient(ctx, projectID)
    if err != nil {
        return err
    }
    defer client.Close()

    // Use the client...
    return nil
}
```

---

## 5. Formatting Table Output

If your command returns BigQuery rows, use the utilities in `utils/formats.go` to ensure a consistent look:

```go
import (
    "github.com/rantolin/bqsql/utils"
    "cloud.google.com/go/bigquery"
)

// 1. Calculate widths (requires iterating through the results once)
widths, err := formats.CalculateRowWidths(it, schema)

// 2. Reset or re-run the iterator and print
for {
    var row []bigquery.Value
    err := it.Next(&row)
    if err == iterator.Done { break }
    
    formats.PrintFormatedRow(cmd.OutOrStdout(), row, widths)
}
```

---

## 6. Best Practices

- **Errors**: Return errors from functions and let Cobra or your `Run` block handle them (usually with `fmt.Println` or `os.Exit`).
- **Context**: Use `context.Background()` or pass context down if necessary for cancellation.
- **Flags**: Use `PersistentFlags()` on the `rootCmd` only for global settings. Use local `Flags()` for command-specific options.
- **Documentation**: Always provide `Short` and `Long` descriptions for your commands.
