package cmd

import (
	"fmt"
	"strings"

	"fardhan.dev/dreamjournal/internal/db"
	"fardhan.dev/dreamjournal/internal/model"
	"fardhan.dev/dreamjournal/internal/repository"
	"fardhan.dev/dreamjournal/internal/utils"
	"github.com/spf13/cobra"
)

var titleFlag string

var addCmd = &cobra.Command{
	Use:   "add [content]",
	Short: "Add a new dream entry",
	Long:  `Add a new dream entry. You can provide content as an argument, or leave it empty to open your default editor.`,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		if len(args) > 0 {
			content = strings.Join(args, " ")
		} else {
			var err error
			content, err = utils.OpenEditor("")
			if err != nil {
				fmt.Printf("Error opening editor: %v\n", err)
				return
			}
		}

		if strings.TrimSpace(content) == "" {
			fmt.Println("Dream content cannot be empty.")
			return
		}

		title := titleFlag
		if title == "" {
			title = utils.GenerateAutoTitle(content)
		}

		repo := repository.NewDreamRepository(db.DB)
		dream := &model.Dream{
			Title:   title,
			Content: content,
		}

		if err := repo.CreateDream(dream); err != nil {
			fmt.Printf("Error adding dream: %v\n", err)
			return
		}

		fmt.Printf("Dream added successfully! ID: %d\n", dream.ID)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&titleFlag, "title", "t", "", "Title of the dream")
}
