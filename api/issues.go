package api

import (
	"github.com/shurcooL/githubql"
	"context"
	"github.com/crocus7724/korat/model"
)

func GetViewerIssues(repository githubql.String, states []githubql.IssueState) ([]model.Issue, error) {
	var query struct {
		Viewer struct {
			Repository struct {
				Issues struct {
					Nodes []model.Issue

					PageInfo PageInfo
				} `graphql:"issues(first: 50, orderBy: {field: CREATED_AT, direction: DESC}, after: $issuesCursor, states: $states)"`
			} `graphql:"repository(name: $name)"`
		}
	}

	variables := map[string]interface{}{
		"issuesCursor": (*githubql.String)(nil),
		"name":         repository,
		"states": states,
	}

	var issues []model.Issue

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
