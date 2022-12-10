/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ixxiv/price-action/utils"
	"github.com/ixxiv/price-action/utils/theme"
	"github.com/mbndr/figlet4go"
	"github.com/scylladb/termtables"
	"github.com/spf13/cobra"
)

// stocksCmd represents the stocks command
var stocksCmd = &cobra.Command{
	Use:     "stocks",
	Short:   theme.Yellow("WIP - Coming soon..."),
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		// out, err := exec.Command("stocks", args...).Output()
		// if err != nil {
		// 	log.Fatal(err)
		// }

		fmt.Println(theme.AsciiRender("Crypto Prices", "standard", []figlet4go.Color{figlet4go.ColorYellow, figlet4go.ColorBlue}, figlet4go.Parser{Name: "terminal", NewLine: "\n"}))

		stockDetails, err := utils.GetStocksdata()
		if err != nil {
			fmt.Printf("%v", err)
		}

		table := termtables.CreateTable()
		table.SetAlign(termtables.AlignCenter, 12)

		table.AddHeaders("Index", "Currency", "Symbol", "Timezone", "RegularMarketPrice", "ChartPreviousClose", "PreviousClose")
		for index, element := range stockDetails.Chart.Result {
			table.AddRow(
				theme.Green(fmt.Sprintf("%d", index+1)),
				theme.Blue(element.Meta.Currency),
				theme.Yellow(element.Meta.Symbol),
				theme.Green(element.Meta.Timezone),
				theme.Green(fmt.Sprintf("%.2f", element.Meta.RegularMarketPrice)),
				theme.Red(fmt.Sprintf("%.2f", element.Meta.ChartPreviousClose)),
				theme.Blue(fmt.Sprintf("%.2f", element.Meta.PreviousClose)),
			)
		}

		fmt.Println(table.Render())
	},
}

func init() {
	rootCmd.AddCommand(stocksCmd)
}
