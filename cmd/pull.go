package cmd

import (
	"context"
	"log"

	"com.github.niedch/internal/conf"
	"com.github.niedch/internal/markdown"
	"com.github.niedch/internal/rallyapi"
	"github.com/spf13/cobra"
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
		ctx := context.Background();
		config := conf.Load()
		
		client := rallyapi.NewRestClient(config);

		query := rallyapi.NewQueryBuilder().WithFormattedId("DE58840")
		defects, err := client.FindDefects(ctx, *query)
		if err != nil {
			log.Fatal(err);
		}

		out, err := markdown.ConvertToMarkdown(defects[0].Description)
		if err != nil {
			log.Fatal(err);
		}

		log.Println(out)
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	pullCmd.PersistentFlags().String("filename", "task.md", "Where the markdown file should be written to!")
	pullCmd.PersistentFlags().String("defect", "", "Where the markdown file should be written to!")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}
