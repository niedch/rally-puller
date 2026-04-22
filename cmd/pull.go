package cmd

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/niedch/rally-puller/internal/branch"
	"github.com/niedch/rally-puller/internal/conf"
	"github.com/niedch/rally-puller/internal/markdown"
	"github.com/niedch/rally-puller/internal/rallyapi"
	"github.com/spf13/cobra"
)

var (
	filename string
	defect   string
	cwd      string
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		config := conf.Load()

		client := rallyapi.NewRestClient(config)

		ticket, err := branch.ResolveTicket(defect, cwd)
		if err != nil {
			log.Fatal(err)
		}

		query := rallyapi.NewQueryBuilder().WithFormattedId(ticket.ID)
		defects, err := client.FindTickets(ctx, string(ticket.Type), *query)
		if err != nil {
			log.Fatal(err)
		}

		out, err := markdown.ConvertToMarkdown(defects[0].Description)
		if err != nil {
			log.Fatal(err)
		}

		outPath := filepath.Join(cwd, filename)
		if err := os.WriteFile(outPath, []byte(out), 0o644); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)

	pullCmd.Flags().StringVarP(&filename, "filename", "f", "task.md", "Where the markdown file should be written to!")
	pullCmd.Flags().StringVarP(&defect, "defect", "d", "", "Rally defect FormattedID (e.g. DE123); if omitted, the current git branch is scanned for DE<number>")

	pullCmd.PersistentFlags().StringVarP(&cwd, "cwd", "c", "", "Working directory")
}
