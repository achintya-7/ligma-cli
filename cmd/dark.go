/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/achintya-7/ligma-cli/helperFunc"

	"github.com/spf13/cobra"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

// darkCmd represents the dark command
var darkCmd = &cobra.Command{
	Use:   "dark",
	Short: "get a dark joke :)",
	Long:  `Use this under parental guidance :)`,
	Run: func(cmd *cobra.Command, args []string) {
		err := displayPost()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(darkCmd)

}

func displayPost() error {
	httpClient := &http.Client{Timeout: time.Second * 30}
	apiCredentials, err := helperFunc.GetApiCredentials()
	if err != nil {
		fmt.Printf("Something went wrong while using .ENV file %v", err)
		return nil
	}

	credentials := reddit.Credentials{
		ID:       apiCredentials.ID,
		Password: apiCredentials.Password,
		Secret:   apiCredentials.Secret,
		Username: "achintya22",
	}

	client, _ := reddit.NewClient(credentials, reddit.WithHTTPClient(httpClient))

	posts, _, err := client.Subreddit.TopPosts(context.Background(), "darkjokes", &reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{
			Limit: 200,
		},
		Time: "new",
	})

	if err != nil {
		return err
	}

	rand.Seed(time.Now().Unix())

	min, max := 0, len(posts)
	randomNum := min + rand.Intn(max)

	post := posts[randomNum]

	fmt.Println(post.Title)
	time.Sleep(2 * time.Second)
	fmt.Println(post.Body)
	time.Sleep(2 * time.Second)

	return nil
}
