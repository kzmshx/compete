package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "compete <command> [flags]",
	Short: "Compete is a helper tool for competitive programming",
}

func init() {
	rootCommand.AddCommand(newConfigureCommand())
	rootCommand.AddCommand(newNewCommand())
	rootCommand.AddCommand(newSubmitCommand())
	rootCommand.AddCommand(newTestCommand())
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
