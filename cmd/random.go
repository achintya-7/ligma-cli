/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random joke",
	Long:  `This gives out a random joke in your terminal`,
	Run: func(cmd *cobra.Command, args []string) {
		
		jokeTerm, _ := cmd.Flags().GetString("term")
		
		if jokeTerm != "" {
			getRandomJokeWithTerm(jokeTerm)
		} else {
			getRandomJoke()
		}
		
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
	rootCmd.PersistentFlags().String("term", "", "Add a term to get jokes realted to it")
}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

type SearchResult struct {
	Results json.RawMessage `json:"results"`
	SearchTerm string `json:"search_term"`
	Status int `json:"status"`
	TotalJokes int `json:"total_jokes"`
}

func getRandomJoke() {
	url := "https://icanhazdadjoke.com/"
	responseBytes := getJokeData(url)
	joke := Joke{}
	
	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		log.Panicf("Could not unmarshal the response : %v", err)
	}

	fmt.Println(string(joke.Joke))
}

func getRandomJokeWithTerm(term string) {
	total, results := getJokeDataWithTerm(term)
	randomiseJokeList(total, results)
}

func getJokeDataWithTerm(term string) (totalJokes int, jokeList []Joke) {
	url := fmt.Sprintf("https://icanhazdadjoke.com/search?term=%s", term)
	responseBytes := getJokeData(url)

	jokeListRaw := SearchResult{}

	if err := json.Unmarshal(responseBytes, &jokeListRaw); err != nil {
		log.Printf("Could not unmarshal response : %v", err)
	}

	jokes := []Joke{}
	if err := json.Unmarshal(jokeListRaw.Results, &jokes); err != nil {
		log.Printf("Could not unmarshal list of results : %v", err)
	} 

	return jokeListRaw.TotalJokes, jokes
}

func getJokeData(baseApi string) []byte {
	// it takes 3 values
	request, err := http.NewRequest(
		http.MethodGet,
		baseApi,
		nil,
	)
	if err != nil {
		log.Printf("Could not request Dad joke : %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dadjoke CLI (github.com/achintya-7/ligma-cli)")

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

func randomiseJokeList(length int, jokeList []Joke) {
	rand.Seed(time.Now().Unix())

	min, max := 0, length

	if length <= 0 {
		err := fmt.Errorf("no jokes found with this term \nPs. This a family friendly cli")
		fmt.Println(err)
	} else {
		randomNum := min + rand.Intn(max)
		fmt.Println(jokeList[randomNum].Joke)
	}
}