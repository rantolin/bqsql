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

## 7. Commit Standard (Conventional Commits)

`bqsql` follows the [Conventional Commits 1.0.0](https://www.conventionalcommits.org/en/v1.0.0/) specification for commit messages. This allows for automated changelog generation and easier project tracking.

### Commit Format

Each commit message consists of a **header**, a **body**, and a **footer**.

```text
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

### Common Types

- `feat`: A new feature.
- `fix`: A bug fix.
- `docs`: Documentation only changes.
- `style`: Changes that do not affect the meaning of the code (white-space, formatting, etc).
- `refactor`: A code change that neither fixes a bug nor adds a feature.
- `perf`: A code change that improves performance.
- `test`: Adding missing tests or correcting existing tests.
- `chore`: Changes to the build process or auxiliary tools and libraries.

### Example

```text
feat(query): add support for dry-run mode

This allows users to estimate the number of bytes processed before running the query.

BREAKING CHANGE: the '--estimate' flag has been replaced by '--dry-run'.
```

### Automation

A `CHANGELOG.md` is automatically generated and updated via GitHub Actions whenever a new version tag (e.g., `v1.2.3`) is pushed to the repository.
