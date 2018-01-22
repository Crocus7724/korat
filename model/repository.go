package model

import "github.com/shurcooL/githubql"

type Repository struct {
	Name        githubql.String
	Description githubql.String
	Url         githubql.String
	UpdatedAt   githubql.String

}
