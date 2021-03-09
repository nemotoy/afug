package main

import (
	"context"
	"fmt"
	"os"

	gh "github.com/nemotoy/afug/github"
)

func main() {

	cli := gh.NewClient(context.Background(), os.Getenv("GITHUB_TOKEN"))
	users, err := cli.GetUsersStarredRepos(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("users: %+v", users)
}
