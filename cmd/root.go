package cmd

import (
	"fmt"
	"os"

	"fardhan.dev/dreamjournal/internal/db"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dream",
	Short: "A CLI tool for journaling your dreams",
	Long:  `Dream Journal is a CLI application that allows you to quickly record, manage, and search your dreams.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if err := db.InitDB(); err != nil {
		fmt.Printf("Error initializing database: %v\n", err)
		os.Exit(1)
	}
}
