/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// stocksCmd represents the stocks command
var stocksCmd = &cobra.Command{
	Use:     "stocks",
	Short:   "WIP - Coming soon...",
	Aliases: []string{"stks"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Coming soon...")
	},
}

func init() {
	rootCmd.AddCommand(stocksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stocksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stocksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
