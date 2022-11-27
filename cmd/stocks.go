/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ixxiv/price-action/utils/theme"
	"github.com/spf13/cobra"
)

// stocksCmd represents the stocks command
var stocksCmd = &cobra.Command{
	Use:     "stocks",
	Short:   theme.Yellow("WIP - Coming soon..."),
	Aliases: []string{"s"},
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println(theme.Green("WIP - Coming soon..."))
		fmt.Println(theme.Yellow("Dev is Cooking... 🧑‍🍳"))
	},
}

func init() {
	rootCmd.AddCommand(stocksCmd)
}
