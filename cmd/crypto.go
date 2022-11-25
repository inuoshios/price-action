/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/ixxiv/price-action/utils"
	"github.com/ixxiv/price-action/utils/theme"
	"github.com/scylladb/termtables"
	"github.com/spf13/cobra"
)

// cryptoCmd represents the crypto command
var cryptoCmd = &cobra.Command{
	Use:     "crypto",
	Aliases: []string{"c"},
	Short:   "gives you information on real-time crypto prices 🚀",
	Run: func(_ *cobra.Command, _ []string) {
		coinDetails, err := utils.GetCryptoData()
		if err != nil {
			fmt.Printf("%v", err)
		}

		table := termtables.CreateTable()
		table.SetAlign(termtables.AlignCenter, 12)

		table.AddHeaders("Index", "Name", "Symbol", "Price", "PriceChangePercentage24h", "MarketCapChange24h", "ATH")
		for index, rows := range coinDetails {
			index++
			table.AddRow(
				theme.Blue(fmt.Sprintf("%d", index)),
				theme.Blue(rows.Name),
				theme.Gray(strings.ToUpper(rows.Symbol)),
				theme.Green(fmt.Sprintf("%.2f", rows.Price)),
				theme.Green(fmt.Sprintf("%.2f", rows.PriceChangePercentage24h)),
				theme.Red(fmt.Sprintf("%.2f", rows.MarketCapChange24h)),
				theme.Blue(fmt.Sprintf("%.2f", rows.ATH)),
			)
		}

		fmt.Println(table.Render())
	},
}

func init() {
	rootCmd.AddCommand(cryptoCmd)
}
