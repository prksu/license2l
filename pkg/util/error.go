package util

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DefaultError default error exit code and error prompt
var (
	DefaultErrorExitCode = 1
	DefaultErrorPrompt   = "\033[1;31merror\033[0m"
)

// CommandErr check and print an error
func CommandErr(err error) {
	if err == nil {
		return
	}

	msg := err.Error()
	if !strings.HasSuffix(msg, "\n") {
		msg += "\n"
	}
	if !strings.HasPrefix(msg, "error: ") {
		msg = fmt.Sprintf("%s: %s", DefaultErrorPrompt, msg)
	}

	fmt.Fprint(os.Stderr, msg)
	os.Exit(DefaultErrorExitCode)
}

// UsageErrorf usage error handler
func UsageErrorf(cmd *cobra.Command, format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	return fmt.Errorf("%s\nRun '%s -h' for help", msg, cmd.CommandPath())
}
