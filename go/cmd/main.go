package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yana1205-test/sample-app/go/cmd/hello"
)

var (
	version = "none"
	commit  = "none"
	date    = "unknown"
)

func newVersionCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "version",
		Short: "Display version",
		RunE: func(cmd *cobra.Command, args []string) error {
			message := fmt.Sprintf("version: %s, commit: %s, date: %s", version, commit, date)
			fmt.Fprintln(os.Stdout, message)
			return nil
		},
	}
	return command
}

func main() {
	command := &cobra.Command{
		Use: "yana1205-test-sample-app",
	}

	command.AddCommand(newVersionCmd())
	command.AddCommand(hello.New())
	err := command.Execute()
	if err != nil {
		os.Exit(1)
	}
}
