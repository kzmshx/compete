package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newNewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "new <url> <lang>",
		Short: "Create a new directory for a workspace",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			url, lang := args[0], args[1]
			fmt.Println(url, lang)
			return nil
		},
	}
}
