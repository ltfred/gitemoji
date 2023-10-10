package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/manifoldco/promptui"

	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Interactively commit using the prompts.",
	Run:   commit,
}

func init() {
	rootCmd.AddCommand(commitCmd)
}

func commit(cmd *cobra.Command, args []string) {
	var data Data
	if err := json.Unmarshal(emojis, &data); err != nil {
		Error(cmd, args, err)
	}

	idx, err := fuzzyfinder.FindMulti(
		data.GitEmojis,
		func(i int) string {
			return strings.Join([]string{data.GitEmojis[i].Emoji, data.GitEmojis[i].Description},
				" - ")
		})
	if err != nil {
		Error(cmd, args, err)
	}

	moji := data.GitEmojis[idx[0]]
	blue := color.New(color.FgBlue).SprintFunc()
	fmt.Println(fmt.Sprintf("Selected： %v - %v", moji.Emoji, blue(moji.Description)))

	prompt := promptui.Prompt{
		Label: "Enter the commit title",
		Validate: func(s string) error {
			if s == "" {
				return errors.New("invalid commit title")
			}
			return nil
		},
	}

	title, err := prompt.Run()
	if err != nil {
		Error(cmd, args, err)
	}

	prompt = promptui.Prompt{
		Label: "Enter the commit message",
	}

	msg, err := prompt.Run()
	if err != nil {
		Error(cmd, args, err)
	}

	commandArgs := []string{"commit", "-a", "-m", moji.Emoji + title, "-m", msg}
	if err = exec.Command("git", commandArgs...).Run(); err != nil {
		Error(cmd, args, err)
	}
}
