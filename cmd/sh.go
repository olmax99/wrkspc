package cmd

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(shCmd)
}

var shCmd = &cobra.Command{
	Use:   "sh",
	Short: "The Ape Machine shell (amsh).",
	Long:  runtxt,
	RunE: func(_ *cobra.Command, _ []string) error {
		for {
			var prefix string

			switch prompt.Input("", completer) {
			case "test":
				prefix = "test"
			case "exit":
				return nil
			default:
				prefix = ""
			}

			fmt.Println(prefix)
		}
	},
}

func completer(d prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{
		{Text: "test", Description: "A test suggestion/completion."},
	}
}
