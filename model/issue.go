package model

import "github.com/shurcooL/githubql"

type Issue struct {
	Title githubql.String
	Url githubql.String
	Number githubql.Int
}
