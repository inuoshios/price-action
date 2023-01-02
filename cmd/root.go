/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/inuoshios/price-action/utils/theme"
	f "github.com/mbndr/figlet4go"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "price",
	Short: theme.AsciiRender("Price", "larry3d", []f.Color{f.ColorBlue, f.ColorYellow}, f.Parser{Name: "terminal", NewLine: "\n"}),
	// Run: func(cmd *cobra.Command, args []string) {

	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
