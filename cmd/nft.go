/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// nftCmd represents the nft command
var nftCmd = &cobra.Command{
	Use:     "nft",
	Aliases: []string{"n"},
	Short:   "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("nft called")
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
