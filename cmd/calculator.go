package cmd

import (
	"github.com/andres15alvarez/learning_go/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(calculator)
}

var calculator = &cobra.Command{
	Use:     "calculator",
	Aliases: []string{"calculadora", "cal"},
	Short:   "Simple calculator that receive two integers and use the arithmetic operators",
	Long:    `Simple calculator that receive two integers and use the arithmetic operatos.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.Calculate()
	},
}
