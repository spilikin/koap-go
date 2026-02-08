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

func newSDSCmd() *cobra.Command {
	sdsCmd := &cobra.Command{
		Use:   "sds",
		Short: "Service directory commands",
	}

	sdsCmd.AddCommand(newSDSRawCmd())
	sdsCmd.AddCommand(newSDSInfoCmd())
	sdsCmd.AddCommand(newSDSServicesCmd())

	return sdsCmd
}

func newSDSRawCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "raw",
		Short: "Show raw service directory XML",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			config, err := loadDotkon()
			if err != nil {
				return err
			}
			return runSDSRaw(config)
		},
	}
}

func newSDSInfoCmd() *cobra.Command {
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

func newSDSServicesCmd() *cobra.Command {
	var jsonOutput bool

	cmd := &cobra.Command{
		Use:   "services",
		Short: "List available services",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			config, err := loadDotkon()
			if err != nil {
				return err
			}
			return runSDSServices(config, jsonOutput)
		},
	}

	cmd.Flags().BoolVar(&jsonOutput, "json", false, "output as JSON")

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

func runSDSRaw(config *koap.Dotkon) error {
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

func runSDSServices(config *koap.Dotkon, jsonOutput bool) error {
	services, err := loadServices(config)
	if err != nil {
		return err
	}

	if jsonOutput {
		return printJSON(services.ServiceInformation.Service)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "SERVICE\tVERSION\tENDPOINT")
	for _, svc := range services.ServiceInformation.Service {
		for _, v := range svc.Versions {
			fmt.Fprintf(w, "%s\t%s\t%s\n", svc.Name, v.Version, v.EndpointTLS.Location)
		}
	}
	return w.Flush()
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
