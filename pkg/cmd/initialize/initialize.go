package initialize

import (
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/spf13/cobra"

	"github.com/prksu/license2l/pkg/license"
	"github.com/prksu/license2l/pkg/util"
)

// Options for init command
type Options struct {
	License            *template.Template
	IsLicenseAvailable bool
	LicenseFile        *os.File
	LicensePath        string
	LicenseType        string
	LicenseData        license.Data
}

var defaultYear = time.Now().Year()

// NewInitializeCommand create `init` command
func NewInitializeCommand() *cobra.Command {
	opts := new(Options)
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize new LICENSE file",
		Run: func(cmd *cobra.Command, args []string) {
			util.CommandErr(opts.Complete())
			util.CommandErr(opts.Validate())
			util.CommandErr(opts.Run())
		},
	}

	cmd.Flags().StringVar(&opts.LicenseType, "type", opts.LicenseType, "License Type")
	cmd.Flags().StringVar(&opts.LicenseData.Holder, "holder", opts.LicenseData.Holder, "Copyright holder")
	cmd.Flags().IntVarP(&opts.LicenseData.Year, "year", "y", defaultYear, "Year")
	cmd.MarkFlagRequired("type")
	cmd.MarkFlagRequired("holder")
	return cmd
}

// Complete fill all the required options struct
func (o *Options) Complete() error {
	var err error
	o.LicensePath = "LICENSE"
	o.LicenseFile, err = os.Create(o.LicensePath)
	if err != nil {
		return err
	}

	o.License, o.IsLicenseAvailable = license.License[o.LicenseType]

	return nil
}

// Validate makes sure provided values from options struct are valid
func (o *Options) Validate() error {
	if !o.IsLicenseAvailable {
		return fmt.Errorf("Unsupported %s license", o.LicenseType)
	}

	return nil
}

// Run run init command
func (o *Options) Run() error {
	o.License.Execute(o.LicenseFile, o.LicenseData)
	o.LicenseFile.Close()

	return nil
}
