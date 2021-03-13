package main

import (
	"os"

	"github.com/nemotoy/afug/cmd"

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

func main() {
	cmd.Execute()
}
