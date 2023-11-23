package cmd

import (
	"github.com/spf13/cobra"
)

func newSubmitCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "submit [dir]",
		Short: "Submit a solution",
	}
}
