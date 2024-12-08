/*
Copyright © 2024 Roberto Antolin <rantolin DOT geo AT gmail DOT com>
*/
package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func WriteConfigFile(w io.Writer, isDefault bool, profile string, projectID string, datasetID string) {
    // Create a new Viper instance just for writing
    v := viper.New()
    v.SetConfigFile("bqsql.yaml")

    // Read existing config
    if err := v.ReadInConfig(); err != nil {
        // If file doesn't exist, create a new config
        config := Config{
            Configurations: make(map[string]Configuration),
        }
        v.Set("configurations", config.Configurations)
        v.Set("default", DefaultProfileName)
    }

    // Get the entire existing configuration
    var config Config
    if err := v.Unmarshal(&config); err != nil {
        fmt.Fprintf(w, "Error parsing config: %v\n", err)
        return
    }

    // If no profile specified, use the default one
    if profile == "" {
        profile = config.Default
    }

    // Get or create the profile configuration
    if config.Configurations == nil {
        config.Configurations = make(map[string]Configuration)
    }
    profileConfig := config.Configurations[profile]

    // Only update values that are provided
    if projectID != "" {
        profileConfig.Project = projectID
    }
    if datasetID != "" {
        profileConfig.Dataset = datasetID
    }

    // Update the configuration
    config.Configurations[profile] = profileConfig
    if isDefault {
        config.Default = profile
    }

    // Write back the entire configuration
    v.Set("configurations", config.Configurations)
    v.Set("default", config.Default)

    if err := v.WriteConfig(); err != nil {
        fmt.Fprintf(w, "Error writing config file: %v\n", err)
        return
    }

    fmt.Fprintf(w, "Configuration written for profile %s in file ./%s\n", profile, v.ConfigFileUsed())
}

// setCmd represents the set command
var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the project and dataset",
	Long: `The 'set' command sets the project and dataset to be used by the bqsql tool.
These values are stored in a configuration file and are used by other commands to determine the project and dataset to operate on.
Use 'bqsql set --project_id my-project --dataset my_dataset'
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get values directly from flags
		projectID, _ := cmd.Flags().GetString("project_id")
		datasetID, _ := cmd.Flags().GetString("dataset")
		profile, _ := cmd.Flags().GetString("profile")

		isDefault, _ := cmd.Flags().GetBool("default")
		WriteConfigFile(cmd.OutOrStdout(), isDefault, profile, projectID, datasetID)
	},
}

func init() {
	rootCmd.AddCommand(SetCmd)

	SetCmd.Flags().BoolP("default", "d", true, "Set the default profile")
}
