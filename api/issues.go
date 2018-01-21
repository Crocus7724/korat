package api

import (
	"github.com/shurcooL/githubql"
	"context"
)

type Issue struct {
	Title githubql.String
	Url githubql.String
	Number githubql.Int
}

type Issues struct {
	Nodes []Issue
	PageInfo struct {
		HasNextPage githubql.Boolean
		EndCursor   githubql.String
	}
}

func GetViewerIssues(repository githubql.String) ([]Issue, error) {
	var query struct {
		Viewer struct {
			Repository struct {
				Issues Issues `graphql:"issues(first: 50, orderBy: {field: CREATED_AT, direction: DESC}, after: $issuesCursor, states: OPEN)"`
			} `graphql:"repository(name: $name)"`
		}
	}

	variables := map[string]interface{}{
		"issuesCursor": (*githubql.String)(nil),
		"name":         repository,
	}

	var issues []Issue

	for {
		if err := client.Query(context.Background(), &query, variables); err != nil {
			return nil, err
		}

		issues = append(issues, query.Viewer.Repository.Issues.Nodes...)

		if ! query.Viewer.Repository.Issues.PageInfo.HasNextPage {
			break
		}

		variables["issuesCursor"] = query.Viewer.Repository.Issues.PageInfo.EndCursor
	}

	return issues, nil
}
