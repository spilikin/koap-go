package main

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/alecthomas/chroma/v2/quick"
	"github.com/spf13/cobra"
	koap "github.com/spilikin/koap-go"
)

var outputFlag string

func newGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get resources from the Konnektor",
	}

	cmd.PersistentFlags().StringVarP(&outputFlag, "output", "o", "table", "output format: table, json")

	cmd.AddCommand(newGetInfoCmd())
	cmd.AddCommand(newGetServicesCmd())

	return cmd
}

func newGetInfoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "Show Konnektor product information",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			config, err := loadDotkon()
			if err != nil {
				return err
			}
			return runGetInfo(config)
		},
	}
}

func newGetServicesCmd() *cobra.Command {
	var raw bool

	cmd := &cobra.Command{
		Use:   "services",
		Short: "List available services",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			config, err := loadDotkon()
			if err != nil {
				return err
			}
			if raw {
				return runGetServicesRaw(config)
			}
			return runGetServices(config)
		},
	}

	cmd.Flags().BoolVar(&raw, "raw", false, "show raw service directory XML")

	return cmd
}

func loadServices(config *koap.Dotkon) (*koap.ConnectorServices, error) {
	httpClient, baseURL, err := koap.NewHTTPClient(config)
	if err != nil {
		return nil, err
	}
	services, err := koap.LoadConnectorServices(context.Background(), httpClient, baseURL)
	if err != nil {
		return nil, err
	}
	if config.RewriteServiceEndpoints {
		services.RewriteEndpoints(baseURL)
	}
	return services, nil
}

func runGetInfo(config *koap.Dotkon) error {
	services, err := loadServices(config)
	if err != nil {
		return err
	}

	if outputFlag == "json" {
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

func runGetServices(config *koap.Dotkon) error {
	services, err := loadServices(config)
	if err != nil {
		return err
	}

	if outputFlag == "json" {
		return printJSON(services.ServiceInformation.Service)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "SERVICE\tVERSION\tENDPOINT")
	for _, svc := range services.ServiceInformation.Service {
		for _, v := range svc.Versions {
			endpoint := ""
			if v.EndpointTLS != nil {
				endpoint = v.EndpointTLS.Location
			}
			fmt.Fprintf(w, "%s\t%s\t%s\n", svc.Name, v.Version, endpoint)
		}
	}
	return w.Flush()
}

func runGetServicesRaw(config *koap.Dotkon) error {
	services, err := loadServices(config)
	if err != nil {
		return err
	}

	indented, err := indentXML(services.Raw)
	if err != nil {
		fmt.Print(string(services.Raw))
		return nil
	}

	if isTerminal() {
		return quick.Highlight(os.Stdout, indented, "xml", "terminal256", "monokai")
	}
	fmt.Print(indented)
	return nil
}

func indentXML(data []byte) (string, error) {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	var buf bytes.Buffer
	encoder := xml.NewEncoder(&buf)
	encoder.Indent("", "  ")

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		if err := encoder.EncodeToken(token); err != nil {
			return "", err
		}
	}
	if err := encoder.Flush(); err != nil {
		return "", err
	}
	buf.WriteByte('\n')
	return buf.String(), nil
}
