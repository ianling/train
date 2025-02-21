package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/gosuri/uiprogress"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:               "create [org_name|user_name]",
	Short:             "Create release PRs for an org or user's repos",
	Args:              cobra.ExactArgs(1),
	PersistentPreRun:  setupClient,
	Run:               createFunc,
	ValidArgsFunction: createCmdValidArgsFunc,
}

func createFunc(cmd *cobra.Command, args []string) {
	ctx := context.Background()

	uiprogress.Start()

	fmt.Printf("Entity: %v\n", args[0])
	fmt.Printf("Base: %v\n", base)

	if dryRun {
		fmt.Println()
		fmt.Println("===============")
		fmt.Println("Doing a dry run")
		fmt.Println("===============")
	}

	fmt.Println()

	repos, err := clt.GetRepos(ctx, args[0])
	if err != nil {
		fmt.Printf("repos: %v\n", err.Error())
		os.Exit(1)
	}

	if len(repos) < 1 {
		fmt.Printf("github: no repos found\n")
		return
	}

	urls, err := clt.ProcessRepos(ctx, repos, dryRun)
	if err != nil {
		fmt.Printf("process repos: %v\n", err.Error())
		os.Exit(1)
	}

	uiprogress.Stop()

	if len(urls) > 0 {
		fmt.Println()
		if dryRun {
			fmt.Println("(Dryrun) Release PRs Created:")
		} else {
			fmt.Println("Release PRs Created:")
		}

		for _, url := range urls {
			fmt.Println(url)
		}

		return
	}
}

func createCmdValidArgsFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	setupClient(cmd, args)

	valid, err := clt.GetLogins(context.Background())
	if err != nil {
		valid = []string{"error fetching"}
	}

	return valid, cobra.ShellCompDirectiveNoFileComp
}
