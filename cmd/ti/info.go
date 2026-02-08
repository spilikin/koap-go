package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	koap "github.com/spilikin/koap-go"
)

func newInfoCmd() *cobra.Command {
	var jsonOutput bool

	cmd := &cobra.Command{
		Use:   "info",
		Short: "Show Konnektor product information",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			config, err := loadDotkon()
			if err != nil {
				return err
			}
			return runInfo(config, jsonOutput)
		},
	}

	cmd.Flags().BoolVar(&jsonOutput, "json", false, "output as JSON")

	return cmd
}

func runInfo(config *koap.Dotkon, jsonOutput bool) error {
	services, err := loadServices(config)
	if err != nil {
		return err
	}

	if jsonOutput {
		return printJSON(services.ProductInformation)
	}

	pi := services.ProductInformation
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintf(w, "Product Type\t%s\n", pi.ProductTypeInformation.ProductType)
	fmt.Fprintf(w, "Product Type Version\t%s\n", pi.ProductTypeInformation.ProductTypeVersion)
	fmt.Fprintf(w, "Vendor\t%s\n", pi.ProductIdentification.ProductVendorID)
	fmt.Fprintf(w, "Product Code\t%s\n", pi.ProductIdentification.ProductCode)
	fmt.Fprintf(w, "Hardware Version\t%s\n", pi.ProductIdentification.ProductVersion.Local.HWVersion)
	fmt.Fprintf(w, "Firmware Version\t%s\n", pi.ProductIdentification.ProductVersion.Local.FWVersion)
	return w.Flush()
}
