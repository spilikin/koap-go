package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	koap "github.com/spilikin/koap-go"
)

func newGetCertificatesCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "certificates <card-handle>",
		Short: "List certificates of a card",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			config, err := loadDotkon()
			if err != nil {
				return err
			}
			return runGetCertificates(config, args[0])
		},
	}
}

func runGetCertificates(config *koap.Dotkon, cardHandle string) error {
	client, err := loadClient(config)
	if err != nil {
		return err
	}

	certs, err := client.ReadAllCardCertificates(cardHandle)
	if err != nil {
		return err
	}

	if outputFlag == "json" {
		return printJSON(certs)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "REF\tSUBJECT\tISSUER\tNOT BEFORE\tNOT AFTER\tKEY")
	for _, c := range certs {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n",
			c.CertRef,
			c.X509.Subject.CommonName,
			c.X509.Issuer.CommonName,
			c.X509.NotBefore.Format("2006-01-02"),
			c.X509.NotAfter.Format("2006-01-02"),
			describePublicKey(c.X509.PublicKey),
		)
	}
	return w.Flush()
}
