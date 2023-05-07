/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/sunmeng90/go/xtools/pkg/git"
	"time"
)

var timeout time.Duration

// gitCmd represents the git command
var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	PreRun: func(cmd *cobra.Command, args []string) {
		timeoutStr, _ := cmd.Flags().GetString("timeout")
		timeout, _ = time.ParseDuration(timeoutStr)
	},
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "fetch changes from remote",
	Run: func(cmd *cobra.Command, args []string) {
		var ctx context.Context
		if timeout <= 0 {
			ctx = context.Background()
			cmd.SetContext(ctx)
		} else {
			ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
			defer cancelFunc()
			cmd.SetContext(ctx)
		}

		path, _ := cmd.Flags().GetString("path")
		log.Infof("start fetching all repositories in %s, timeout: %s", path, timeout)
		git.FetchAllWithContext(cmd.Context(), path)
	},
}

func init() {
	rootCmd.AddCommand(gitCmd)
	gitCmd.AddCommand(fetchCmd)
	fetchCmd.Flags().StringP("timeout", "t", "", "--timeout 1s")
	fetchCmd.Flags().StringP("path", "p", ".", "-path .")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
