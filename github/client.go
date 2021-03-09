package client

import (
	"context"
	"os"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type User struct {
	Name                githubv4.String
	StarredRepositories struct {
		Edges []starredRepositoritoryEdge
		Nodes []repository
	} `graphql:"starredRepositories(last: $starredRepositoriesLast)"`
}

type starredRepositoritoryEdge struct {
	StarredAt githubv4.DateTime
}

type repository struct {
	URL             githubv4.URI
	PrimaryLanguage struct {
		Color githubv4.String
		Name  githubv4.String
	}
	StargazerCount githubv4.Int
}

type language struct {
	Color githubv4.String
	Name  githubv4.String
}

type Client struct {
	c *githubv4.Client
}

// TODO: add optinal settings to a http client
func NewClient(ctx context.Context, tokenKey string) *Client {
	httpClient := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv(tokenKey)},
	))
	return &Client{c: githubv4.NewClient(httpClient)}
}

func (c *Client) GetUsersStarredRepos(ctx context.Context) ([]User, error) {

	var query struct {
		Viewer struct {
			Following struct {
				Nodes []User
			} `graphql:"following(last: $followingLast)"`
		}
	}

	// TODO: DI
	variables := map[string]interface{}{
		"followingLast":           githubv4.Int(5),
		"starredRepositoriesLast": githubv4.Int(5),
	}

	err := c.c.Query(ctx, &query, variables)
	if err != nil {
		return nil, err
	}

	return query.Viewer.Following.Nodes, nil
}
