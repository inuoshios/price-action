/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/inuoshios/price-action/utils"
	"github.com/inuoshios/price-action/utils/theme"
	"github.com/mbndr/figlet4go"
	"github.com/scylladb/termtables"
	"github.com/spf13/cobra"
)

// cryptoCmd represents the crypto command
var cryptoCmd = &cobra.Command{
	Use:     "crypto",
	Aliases: []string{"c"},
	Short:   theme.Yellow("Generates real-time information on different crypto prices 🚀"),
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println(theme.AsciiRender("Crypto Prices", "standard", []figlet4go.Color{figlet4go.ColorYellow, figlet4go.ColorBlue}, figlet4go.Parser{Name: "terminal", NewLine: "\n"}))

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
				theme.Gray(fmt.Sprintf("%d", index)),
				theme.Blue(rows.Name),
				theme.Yellow(strings.ToUpper(rows.Symbol)),
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
