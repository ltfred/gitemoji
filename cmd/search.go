package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ktr0731/go-fuzzyfinder"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search git emojis.",
	Run:   search,
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

func search(cmd *cobra.Command, args []string) {
	var data Data
	if err := json.Unmarshal(emojis, &data); err != nil {
		Error(cmd, args, err)
	}

	_, err := fuzzyfinder.FindMulti(
		data.GitEmojis,
		func(i int) string {
			return strings.Join([]string{data.GitEmojis[i].Emoji, data.GitEmojis[i].Description},
				" - ")
		}, fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("Emoji: %s (%s)\nDescription: %s",
				data.GitEmojis[i].Emoji,
				data.GitEmojis[i].Code,
				data.GitEmojis[i].Description)
		}))
	if err != nil {
		Error(cmd, args, err)
	}
}
