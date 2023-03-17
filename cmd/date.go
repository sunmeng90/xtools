/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sunmeng90/go/xtools/pkg/date"
	"time"
)

var tz []string
var tzDefault = []string{"Local", "", "EST"}

// dateCmd represents the date command
var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		for _, t := range tz {
			fmt.Printf("%5s: %s\n", t, date.Format(now, "Local"))
		}
	},
}

func init() {
	rootCmd.AddCommand(dateCmd)
	dateCmd.Flags().StringSliceVarP(&tz, "timezone", "z", tzDefault, "-z='Local,,EST'")
}
