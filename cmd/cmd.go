package cmd

import (
	"log"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "Go Design Patterns",
}

func Execute() {
	pterm.DefaultHeader.Println("Go Design Patterns")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
