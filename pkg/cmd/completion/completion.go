package completion

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/prksu/license2l/pkg/util"
)

// Options options for completion command
type Options struct {
	Shell string
}

// NewCompletionCommand create `completion` command
func NewCompletionCommand() *cobra.Command {
	opts := new(Options)
	cmd := &cobra.Command{
		Use:   "completion SHELL",
		Short: "Generate shell completion",
		Run: func(cmd *cobra.Command, args []string) {
			util.CommandErr(opts.Complete(cmd, args))
			util.CommandErr(opts.Validate())
			util.CommandErr(opts.Run(cmd))
		},
		ValidArgs: []string{"bash"},
	}

	return cmd
}

// Complete fill all the required options struct
func (o *Options) Complete(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return util.UsageErrorf(cmd, "Shell not specified")
	}

	o.Shell = args[0]
	return nil
}

// Validate makes sure provided values from options struct are valid
func (o *Options) Validate() error { return nil }

// Run run completion command
func (o *Options) Run(cmd *cobra.Command) error {
	cmdParrent := cmd.Parent()
	switch o.Shell {
	case "bash":
		return cmdParrent.GenBashCompletion(os.Stdout)
	}

	return fmt.Errorf("%s shell is unsupported yet", o.Shell)
}
