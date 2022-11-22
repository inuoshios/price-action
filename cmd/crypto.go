/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/scylladb/termtables"
	"github.com/spf13/cobra"
)

const (
	ColorDefault = "\x1b[39m"

	ColorRed   = "\x1b[91m"
	ColorGreen = "\x1b[32m"
	ColorBlue  = "\x1b[94m"
	ColorGray  = "\x1b[90m"
)

func red(s string) string {
	return fmt.Sprintf("%s%s%s", ColorRed, s, ColorDefault)
}

func green(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGreen, s, ColorDefault)
}

func blue(s string) string {
	return fmt.Sprintf("%s%s%s", ColorBlue, s, ColorDefault)
}

func gray(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGray, s, ColorDefault)
}

type CoinDetails struct {
	Name                     string  `json:"name"`
	Symbol                   string  `json:"symbol"`
	Price                    float64 `json:"current_price"`
	PriceChangePercentage24h float64 `json:"price_change_percentage_24h"`
	MarketCapChange24h       float64 `json:"market_cap_change_24h"`
	ATH                      float64 `json:"ath"`
}

// cryptoCmd represents the crypto command
var cryptoCmd = &cobra.Command{
	Use:     "crypto",
	Aliases: []string{"c"},
	Short:   "gives you information on real-time crypto prices 🚀",
	Run: func(_ *cobra.Command, _ []string) {
		coinDetails, err := GetCryptoData()
		if err != nil {
			fmt.Printf("%v", err)
		}

		table := termtables.CreateTable()
		table.SetAlign(termtables.AlignCenter, 12)

		table.AddHeaders("Index", "Name", "Symbol", "Price", "PriceChangePercentage24h", "MarketCapChange24h", "ATH")
		for index, rows := range coinDetails {
			index++
			table.AddRow(
				blue(fmt.Sprintf("%d", index)),
				blue(rows.Name),
				gray(strings.ToUpper(rows.Symbol)),
				green(fmt.Sprintf("%.2f", rows.Price)),
				red(fmt.Sprintf("%.2f", rows.PriceChangePercentage24h)),
				red(fmt.Sprintf("%.2f", rows.MarketCapChange24h)),
				blue(fmt.Sprintf("%.2f", rows.ATH)),
			)
		}

		fmt.Println(table.Render())
	},
}

func GetCryptoData() ([]CoinDetails, error) {
	client := &http.Client{}

	var coinDetails []CoinDetails

	coingecko := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=100&page=1&sparkline=false&price_change_percentage=1h"
	req, err := http.NewRequest(http.MethodGet, coingecko, nil)
	if err != nil {
		return nil, fmt.Errorf("error making a request %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error getting a response %w", err)
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			return
		}
	}()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(response, &coinDetails); err != nil {
		return nil, fmt.Errorf("umarshall errror: %w", err)
	}

	return coinDetails, nil

}

func init() {
	rootCmd.AddCommand(cryptoCmd)
}
