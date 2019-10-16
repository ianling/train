package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var (
	// Version is the current version of train, made available for use through
	// out the application.
	Version = "0.0.6"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the version",
	Long:  `Display the version of the CLI.`,
	Run:   versionFunc,
}

func versionFunc(cmd *cobra.Command, args []string) {
	printf("Train version %v", Version)
}
