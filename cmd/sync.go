package cmd

import (
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(syncCmd)
}

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync emoji list with the repo.",
	Run:   sync,
}

type Data struct {
	GitEmojis []emoji `json:"gitmojis"`
}

type emoji struct {
	Emoji       string `json:"emoji"`
	Entity      string `json:"entity"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

func sync(cmd *cobra.Command, args []string) {
	request, err := http.NewRequest("GET", "https://gitmoji.dev/api/gitmojis", nil)
	if err != nil {
		Error(cmd, args, err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		Error(cmd, args, err)
	}

	defer response.Body.Close()

	all, err := io.ReadAll(response.Body)
	if err != nil {
		Error(cmd, args, err)
	}

	if err = os.WriteFile("./cmd/emojis.json", all, os.ModePerm); err != nil {
		Error(cmd, args, err)
	}
}
