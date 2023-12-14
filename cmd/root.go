package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use: "obsidian-bot",
}

func Execute() {
	initCommands()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initCommands() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yaml", "config file (default is config.yaml)")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(runCmd)
}
