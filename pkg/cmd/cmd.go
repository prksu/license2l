package cmd

import (
	"github.com/spf13/cobra"

	"github.com/prksu/license2l/pkg/cmd/initialize"
)

// NewLicense2lCommand root command
func NewLicense2lCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "license2l",
		Short: "Simple command line tool for managing LICENSE",
	}

	cmd.AddCommand(initialize.NewInitializeCommand())
	return cmd
}
