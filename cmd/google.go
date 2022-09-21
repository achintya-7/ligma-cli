/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// googleCmd represents the google command
var googleCmd = &cobra.Command{
	Use:   "google",
	Short: "Give search results",
	Long:  `This command can give search results on the given query`,
	Run: func(cmd *cobra.Command, args []string) {
		queryTerm, _ := cmd.Flags().GetString("q")

		if queryTerm != "" {
			getResultQuery(queryTerm)
		} else {
			fmt.Println("Empty query")
		}
	},
}

func init() {
	rootCmd.AddCommand(googleCmd)
	rootCmd.PersistentFlags().String("q", "", "Add a query for google to search")
}

type Result struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
}

type SearchResults struct {
	Results json.RawMessage `json:"results"`
}

func getResultQuery(query string) {
	responseBytes := getResultData(query)

	results := SearchResults{}
	if err := json.Unmarshal(responseBytes, &results); err != nil {
		log.Printf("Could not unmarshal data : %v", err)
	}

	finalResults := []Result{}
	if err := json.Unmarshal(results.Results, &finalResults); err != nil {
		log.Printf("Could not unmarshal list of data : %v", err)
	}

	max := len(finalResults)
	if max > 2 {
		for i := 0; i < 2; i++ {
			fmt.Println("----------------------------------------------------")
			fmt.Println("Title : ")
			fmt.Println(finalResults[i].Title)
			fmt.Println()
			fmt.Println("Desc : ")
			fmt.Println(finalResults[i].Description)
			fmt.Println()
			fmt.Println("Link : ")
			fmt.Println(finalResults[i].Link)
			fmt.Println()
			fmt.Println("----------------------------------------------------")
		}
	} else {
			fmt.Println("----------------------------------------------------")
			fmt.Println("Title : ")
			fmt.Println(finalResults[0].Title)
			fmt.Println()
			fmt.Println("Desc : ")
			fmt.Println(finalResults[0].Description)
			fmt.Println()
			fmt.Println("Link : ")
			fmt.Println(finalResults[0].Link)
			fmt.Println()
			fmt.Println("----------------------------------------------------")
		
	}
}

func getResultData(query string) []byte {
	url := "https://google-search3.p.rapidapi.com/api/v1/search/q="
	queryUrl := fmt.Sprintf("%s%s", url, query)

	request, err := http.NewRequest(
		http.MethodGet,
		queryUrl,
		nil,
	)
	if err != nil {
		log.Printf("Could not request Rapid API : %v", err)
	}

	request.Header.Add("X-User-Agent", "desktop")
	request.Header.Add("X-Proxy-Location", "EU")
	request.Header.Add("X-RapidAPI-Key", "##############")
	request.Header.Add("X-RapidAPI-Host", "google-search3.p.rapidapi.com")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request : %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body : %v", err)
	}

	return responseBytes
}
