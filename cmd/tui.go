package cmd

import (
	"fmt"
	"os"

	"fardhan.dev/dreamjournal/internal/db"
	"fardhan.dev/dreamjournal/internal/tui"
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Launch the interactive TUI",
	Run: func(cmd *cobra.Command, args []string) {
		if err := tui.Start(db.DB); err != nil {
			fmt.Printf("Error running TUI: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
