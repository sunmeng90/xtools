/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	config2 "github.com/sunmeng90/go/xtools/config"
	"os"
)

var cfgFile string
var version = "v0.0.1"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "x",
	Version: version,
	Short:   "A powerful toolset for personal use",
	Long: `A powerful toolset for personal use. For example:

String functions, date calculation etc. `,
	PreRun: func(cmd *cobra.Command, args []string) {
		initLog()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	initLog()

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initLog() {
	log.SetLevel(log.TraceLevel)
	log.SetOutput(os.Stdout)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		parseSampleConfig()
	}
}

// not applicable for CLI
func watchSampleConfigChange() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		parseSampleConfig()
	})
}

func parseSampleConfig() {
	fmt.Printf("===============Get sample config by key: ===============\n")
	fmt.Printf("custom.name: %v\n", viper.Get("custom.name"))
	fmt.Printf("custom.name: %v\n", viper.GetString("custom.name"))
	fmt.Printf("custom.prop1: %v\n", viper.GetInt("custom.prop1"))
	fmt.Printf("custom.prop2: %v\n", viper.GetBool("custom.prop2"))
	fmt.Printf("custom.time1: %v\n", viper.GetDuration("custom.time1"))
	fmt.Printf("custom.time2: %v\n", viper.GetDuration("custom.time2"))
	fmt.Printf("custom.time3: %v\n", viper.GetDuration("custom.time3"))
	fmt.Printf("custom.time4: %v\n", viper.GetDuration("custom.time4"))
	fmt.Printf("custom.time5: %v\n", viper.GetDuration("custom.time5"))
	fmt.Printf("custom.str_arr1: %v\n", viper.GetStringSlice("custom.str_arr1"))
	fmt.Printf("custom.int_arr1: %v\n", viper.GetIntSlice("custom.int_arr1"))
	fmt.Printf("custom.map1: %v\n", viper.GetStringMap("custom.map1"))
	fmt.Printf("custom.map1.k2: %v\n", viper.GetStringMapString("custom.map1.k2"))

	fmt.Printf("===============Unmarshal sample config and get Config: ===============\n")
	var sampleConfig config2.SampleConfig
	// viper.UnmarshalKey("sample1", &sampleConfig)
	viper.Unmarshal(&sampleConfig, viper.DecoderConfigOption(func(config *mapstructure.DecoderConfig) {
		config.TagName = "yaml"
	}))
	fmt.Printf("custom.name: %v\n", sampleConfig.Custom.Name)
	fmt.Printf("custom.prop1: %v\n", sampleConfig.Custom.Prop1)
	fmt.Printf("custom.prop2: %v\n", sampleConfig.Custom.Prop2)
	fmt.Printf("custom.time1: %v\n", sampleConfig.Custom.Time1)
	fmt.Printf("custom.time2: %v\n", sampleConfig.Custom.Time2)
	fmt.Printf("custom.time3: %v\n", sampleConfig.Custom.Time3)
	fmt.Printf("custom.time4: %v\n", sampleConfig.Custom.Time4)
	fmt.Printf("custom.time5: %v\n", sampleConfig.Custom.Time5)
	fmt.Printf("custom.str_arr1: %v\n", sampleConfig.Custom.StrArr1)
	fmt.Printf("custom.int_arr1: %v\n", sampleConfig.Custom.IntArr1)
	fmt.Printf("custom.map1: %v\n", sampleConfig.Custom.Map1)
	fmt.Printf("custom.map1.k2: %v\n", sampleConfig.Custom.Map1["k2"])
	fmt.Printf("custom.db.redis.host: %v\n", sampleConfig.Db.Redis.Host)
	fmt.Printf("custom.db.redis.port: %v\n", sampleConfig.Db.Redis.Port)
	fmt.Printf("custom.db.mysql.host: %v\n", sampleConfig.Db.Mysql.Host)
	fmt.Printf("custom.db.mysql.port: %v\n", sampleConfig.Db.Mysql.Port)

	fmt.Printf("===============Get all settings: ===============\n")

	fmt.Printf("%v\n", viper.AllSettings())
	fmt.Printf("===============End of printing config ===============\n")
}
