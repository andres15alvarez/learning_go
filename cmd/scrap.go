package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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

func Scrap() {
	url := Input("Enter the url to scrap: ")
	printScrap(url)
}
