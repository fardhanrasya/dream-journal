package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"fardhan.dev/dreamjournal/internal/db"
	"fardhan.dev/dreamjournal/internal/repository"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all dreams",
	Run: func(cmd *cobra.Command, args []string) {
		repo := repository.NewDreamRepository(db.DB)
		dreams, err := repo.GetDreams()
		if err != nil {
			fmt.Printf("Error fetching dreams: %v\n", err)
			return
		}

		if len(dreams) == 0 {
			fmt.Println("No dreams recorded yet.")
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tDate\tTitle")
		fmt.Fprintln(w, "--\t----\t-----")
		
		for _, d := range dreams {
			dateStr := d.CreatedAt.Format("2006-01-02")
			fmt.Fprintf(w, "%d\t%s\t%s\n", d.ID, dateStr, d.Title)
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
