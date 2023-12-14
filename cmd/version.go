package cmd

import (
	"fmt"

	"github.com/aatumaykin/go-obsidian-bot/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Application Version %s\n", version.Version)
	},
}
