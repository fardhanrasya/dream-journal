package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"fardhan.dev/dreamjournal/internal/db"
	"fardhan.dev/dreamjournal/internal/repository"
	"fardhan.dev/dreamjournal/internal/utils"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [id]",
	Short: "Edit an existing dream entry",
	Long:  `Edit an existing dream entry. This will open your default editor with the current content of the dream.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Invalid ID: %s\n", args[0])
			return
		}

		repo := repository.NewDreamRepository(db.DB)
		dream, err := repo.GetDreamByID(id)
		if err != nil {
			fmt.Printf("Error fetching dream: %v\n", err)
			return
		}

		newContent, err := utils.OpenEditor(dream.Content)
		if err != nil {
			fmt.Printf("Error opening editor: %v\n", err)
			return
		}

		if strings.TrimSpace(newContent) == "" {
			fmt.Println("Dream content cannot be empty. Edit cancelled.")
			return
		}
		
		dream.Content = newContent
		if err := repo.UpdateDream(dream); err != nil {
			fmt.Printf("Error updating dream: %v\n", err)
			return
		}

		fmt.Printf("Dream %d updated successfully.\n", dream.ID)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
