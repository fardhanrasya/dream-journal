package cmd

import (
	"fmt"
	"strconv"

	"fardhan.dev/dreamjournal/internal/db"
	"fardhan.dev/dreamjournal/internal/repository"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a dream by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Invalid ID: %s\n", args[0])
			return
		}

		repo := repository.NewDreamRepository(db.DB)
		if err := repo.DeleteDream(id); err != nil {
			fmt.Printf("Error deleting dream: %v\n", err)
			return
		}

		fmt.Printf("Dream %d deleted successfully.\n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
