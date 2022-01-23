package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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
		Scrap(args[0])
	},
}

func printScrap(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error getting response. ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}

	fmt.Printf("%s\n", body)
}

func Scrap(url string) {
	printScrap(url)
}
