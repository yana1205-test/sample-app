package hello

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	command := &cobra.Command{
		Use:   "hello",
		Short: "hello world",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := Complete(); err != nil {
				return err
			}
			if err := Validate(); err != nil {
				return err
			}
			return Run()
		},
	}
	AddFlags(command.Flags())
	return command
}
