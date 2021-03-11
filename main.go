package main

import (
	"context"
	"fmt"
	"os"

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

	app := tui.NewAppWithWidget().SetTableFrame().SetUsers(users)
	if err := app.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
