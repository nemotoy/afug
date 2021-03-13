package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	gh "github.com/nemotoy/afug/github"
	"github.com/nemotoy/afug/tui"
)

func Execute() {
	cli := gh.NewClient(context.Background(), os.Getenv("GITHUB_TOKEN"))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	users, err := cli.GetUsersStarredRepos(ctx)
	if err != nil {
		fmt.Printf("fetch users failed: %v\n", err)
		os.Exit(1)
	}

	app := tui.NewAppWithWidget().SetTableFrame().SetUsers(users)
	if err := app.Run(); err != nil {
		fmt.Printf("app failed: %v\n", err)
		os.Exit(1)
	}
}
