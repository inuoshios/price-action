/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ixxiv/price-action/utils"
	"github.com/ixxiv/price-action/utils/theme"
	"github.com/scylladb/termtables"
	"github.com/spf13/cobra"
)

// nftCmd represents the nft command
var nftCmd = &cobra.Command{
	Use:     "nft",
	Aliases: []string{"n"},
	Short:   "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		nftDetails, err := utils.GetNFTData()
		if err != nil {
			fmt.Printf("%v", err)
		}

		table := termtables.CreateTable()
		table.SetAlign(termtables.AlignCenter, 12)

		table.AddHeaders("Idx", "Slug", "Volume7D", "Sales7D", "VolumeAllTime", "TotalSupply", "MarketCap", "FloorPrice")
		for index, rows := range nftDetails {
			index++
			table.AddRow(
				theme.Red(fmt.Sprintf("%d", index)),
				theme.Blue(rows.Slug),
				theme.Green(fmt.Sprintf("%.2f", rows.Volume7D)),
				theme.Yellow(fmt.Sprintf("%v", rows.Stats.Sales7D)),
				theme.Green(fmt.Sprintf("%.2f", rows.Stats.VolumeAllTime)),
				theme.Green(fmt.Sprintf("%v", rows.Stats.TotalSupply)),
				theme.Yellow(fmt.Sprintf("%.2f", rows.Stats.MarketCap)),
				theme.Blue(fmt.Sprintf("%.2f", rows.Stats.FloorPrice)),
			)
		}

		fmt.Println(table.Render())
	},
}

func init() {
	rootCmd.AddCommand(nftCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nftCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nftCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
