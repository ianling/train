package org

import (
	"fmt"
	"os"

	"github.com/gomicro/train/config"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	dryRun bool
	client *github.Client
)

// OrgCmd represents the root of the org command
var OrgCmd = &cobra.Command{
	Use:              "org [flags]",
	Short:            "Org specific release train commands",
	PersistentPreRun: setupCommand,
}

func setupCommand(cmd *cobra.Command, args []string) {
	var err error
	client, err = config.GetClient()
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		os.Exit(1)
	}

	dryRun = viper.GetBool("dryRun")
}
