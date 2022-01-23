package cmd

import (
	"github.com/andres15alvarez/learning_go/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ssr)
}

var ssr = &cobra.Command{
	Use:     "ssr",
	Aliases: []string{"html"},
	Short:   "Run a server in port 8000 to display html content.",
	Long:    `Run a server in port 8000 to display html content using templates and http.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.RunServer()
	},
}
