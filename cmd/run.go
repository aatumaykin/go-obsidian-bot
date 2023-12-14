package cmd

import (
	"github.com/aatumaykin/go-obsidian-bot/internal/app"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use: "run",
	RunE: func(cmd *cobra.Command, args []string) error {
		application, err := app.NewApp(cfgFile)
		if err != nil {
			return err
		}

		if err := application.NotifyRun(); err != nil {
			return err
		}

		return application.Run()
	},
}
