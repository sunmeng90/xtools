/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init x tools config",
	RunE: func(cmd *cobra.Command, args []string) error {
		return initX()
	},
}

func initX() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	configFile := filepath.Join(home, ".xtools.yaml")
	_, err = os.Lstat(configFile)
	if err != nil {
		_, err := os.OpenFile(configFile, os.O_CREATE, 0644)
		if err != nil {
			log.Errorf("failed to create configuration for x tools, error: %s", err)
			return err
		} else {
			log.Infof("configuration file %s created for x tools", configFile)
		}
	}
	rootPath := filepath.Join(home, ".xtools")
	log.Infof("creating root folder %s", rootPath)
	err = os.MkdirAll(rootPath, 0700)
	if err != nil {
		log.Errorf("failed to create root folder for x tools, error: %s", err)
		return err
	} else {
		log.Infof("root folder %s created for x tools", rootPath)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	initCmd.Flags().String("path", "p", "init x tool configuration")
}
