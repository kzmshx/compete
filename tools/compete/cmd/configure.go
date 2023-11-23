package cmd

import (
	"github.com/spf13/cobra"
)

func newConfigureCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "configure",
		Short: "Configure the environment",
	}
}
