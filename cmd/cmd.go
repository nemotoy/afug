package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	gh "github.com/nemotoy/afug/github"
	"github.com/nemotoy/afug/tui"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	FollowingUsers int `short:"u" long:"users" description:"number of displaying following users"`
	StarredRepos   int `short:"r" long:"repos" description:"number of displaying starred repositories"`
}

func init() {
	parser := flags.NewParser(&opts, flags.Default)
	parser.Name, parser.Usage = "afug", "[options]"
	_, err := parser.Parse()
	if err != nil {
		os.Exit(1)
	}
}

func Execute() {
	cli := gh.NewClient(context.Background(), os.Getenv("GITHUB_TOKEN"))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	users, err := cli.GetUsersStarredRepos(ctx, opts.FollowingUsers, opts.StarredRepos)
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
