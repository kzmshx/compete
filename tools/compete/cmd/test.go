package cmd

import (
	"github.com/spf13/cobra"
)

func newTestCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "test [dir]",
		Short: "Test a solution",
	}
}
