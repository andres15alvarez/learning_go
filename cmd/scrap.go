package cmd

import (
	"github.com/andres15alvarez/learning_go/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(calculator)
}

var scrap = &cobra.Command{
	Use:   "scrap",
	Short: "Get the html content of a page",
	Long:  `Get the html content of a page given a full url. Example: https://www.google.com.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.Scrap(args[0])
	},
}
