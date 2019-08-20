package main

import (
	"os"

	"github.com/prksu/license2l/pkg/cmd"
)

func main() {
	commnad := cmd.NewLicense2lCommand()
	if err := commnad.Execute(); err != nil {
		os.Exit(1)
	}
}
