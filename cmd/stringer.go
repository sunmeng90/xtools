/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/sunmeng90/go/cobra/pkg/stringer"

	"github.com/spf13/cobra"
)

var onlyDigits bool

// stringerCmd represents the stringer command
var stringerCmd = &cobra.Command{
	Use:     "stringer",
	Aliases: []string{"str"},
	Short:   "A tool for string: reverse, count char etc.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stringer called, please add a subcommand")
	},
}

func init() {
	stringerCmd.AddCommand(reverseCmd)
	stringerCmd.AddCommand(lenCmd)
	rootCmd.AddCommand(stringerCmd)
	lenCmd.Flags().BoolVarP(&onlyDigits, "digits", "d", false, "Count only digits")
}

var reverseCmd = &cobra.Command{
	Use:     "reverse [string to reverse]",
	Aliases: []string{"r", "rev"},
	Short:   "Reverse a string",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := stringer.Reverse(args[0])
		fmt.Println(res)
	},
}

var lenCmd = &cobra.Command{
	Use:     "length [string]",
	Aliases: []string{"l", "len"},
	Short:   "Get length of a string",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cnt, kind := stringer.Length(args[0], onlyDigits)
		pluralS := "s"
		if cnt < 2 {
			pluralS = ""
		}
		fmt.Printf("Input [%s] has %d %s%s\n", args[0], cnt, kind, pluralS)
	},
}
