package cmd

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

//go:embed emojis.json
var emojis []byte

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the available git emojis.",
	Run:   list,
}

func list(cmd *cobra.Command, args []string) {
	var data Data
	if err := json.Unmarshal(emojis, &data); err != nil {
		Error(cmd, args, err)
	}

	blue := color.New(color.FgBlue).SprintFunc()
	for _, v := range data.GitEmojis {
		fmt.Println(strings.Join([]string{v.Emoji, blue(v.Code), v.Description}, " - "))
	}
}
