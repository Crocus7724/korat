package api

import (
	"context"
	"github.com/shurcooL/githubql"
	"golang.org/x/oauth2"
)

type PageInfo struct {
	HasNextPage githubql.Boolean
	EndCursor   githubql.String
}

var client *githubql.Client

func Init(url string, token string) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: token,
		},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	if url == "" {
		client = githubql.NewClient(httpClient)
	} else {
		client = githubql.NewEnterpriseClient(url, httpClient)
	}
}
