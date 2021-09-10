package cmd

import (
	"github.com/pterm/pterm"
	"github.com/r-c-correa/go-design-patterns/creational/abstract_factory"
	"github.com/spf13/cobra"
)

var abstractFactoryCommand = &cobra.Command{
	Use:   "abstract_factory",
	Short: "Abstract Factory",
	RunE: func(cmd *cobra.Command, args []string) error {
		pterm.Info.Println("Abstract Factory")

		appWin := abstract_factory.NewApp(&abstract_factory.WinFactory{})
		appWin.CreateUI()

		appLinux := abstract_factory.NewApp(&abstract_factory.LinuxFactory{})
		appLinux.CreateUI()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(abstractFactoryCommand)
}
