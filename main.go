package main

import (
	"context"
	"fmt"
	"os"
	"time"

	gh "github.com/nemotoy/afug/github"
	"github.com/nemotoy/afug/tui"
)

func main() {

	cli := gh.NewClient(context.Background(), os.Getenv("GITHUB_TOKEN"))
	users, err := cli.GetUsersStarredRepos(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var id int
	for _, user := range users {
		for _, repo := range user.StarredRepositories.Nodes {
			id++
			fmt.Printf("id: %d, name: %s, url: %s, lang: %s\n", id, user.Name, repo.URL, repo.PrimaryLanguage.Name)
		}
	}

	time.Sleep(3 * time.Second)

	app := tui.NewAppWithWidget()
	app.SetUsers(users)
	if err := app.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
