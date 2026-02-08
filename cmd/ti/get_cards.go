package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	koap "github.com/spilikin/koap-go"
)

func newGetCardsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "cards",
		Short: "List inserted cards",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			config, err := loadDotkon()
			if err != nil {
				return err
			}
			return runGetCards(config)
		},
	}
}

func runGetCards(config *koap.Dotkon) error {
	client, err := loadClient(config)
	if err != nil {
		return err
	}

	cards, err := client.GetCards()
	if err != nil {
		return err
	}

	if outputFlag == "json" {
		return printJSON(cards)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "HANDLE\tTYPE\tICCSN\tCT\tSLOT\tHOLDER")
	for _, c := range cards {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%d\t%s\n",
			c.CardHandle, c.CardType, c.Iccsn, c.CtId, c.SlotId, c.CardHolderName)
	}
	return w.Flush()
}
