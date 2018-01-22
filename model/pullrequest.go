package model

import "github.com/shurcooL/githubql"

type PullRequest struct {
	Title githubql.String
	Url githubql.String
	Number githubql.Int
	Author struct {
		Login githubql.String
	}
}
