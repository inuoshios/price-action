/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/inuoshios/price-action/utils"
	"github.com/inuoshios/price-action/utils/theme"
	"github.com/mbndr/figlet4go"
	"github.com/scylladb/termtables"
	"github.com/spf13/cobra"
)

var stockName string

// stocksCmd represents the stocks command
var stocksCmd = &cobra.Command{
	Use:     "stocks",
	Short:   theme.Yellow("Generate details of a particular stock price based on the input given to it"),
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(theme.AsciiRender("Stock Prices", "standard", []figlet4go.Color{figlet4go.ColorYellow, figlet4go.ColorBlue}, figlet4go.Parser{Name: "terminal", NewLine: "\n"}))

		stockDetails, err := utils.GetStocksdata(stockName)
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
	stocksCmd.Flags().StringVarP(&stockName, "abbreviation", "a", "TSLA", "Get a single stock by passing in the stock Abbrevation as a cmd flag")
}
