package api

import (
	"github.com/shurcooL/githubql"
	"context"
	"github.com/crocus7724/korat/model"
)

func GetViewerIssues(repository githubql.String, states []githubql.IssueState) (chan []model.Issue, chan error) {
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
		"states":       states,
	}

	iChan := make(chan []model.Issue)
	errChan := make(chan error)
	go func() {
		defer close(iChan)
		defer close(errChan)
		for {
			if err := client.Query(context.Background(), &query, variables); err != nil {
				errChan <- err
				break
			}
			iChan <- query.Viewer.Repository.Issues.Nodes

			if ! query.Viewer.Repository.Issues.PageInfo.HasNextPage {
				break
			}

			variables["issuesCursor"] = query.Viewer.Repository.Issues.PageInfo.EndCursor
		}
	}()

	return iChan, errChan
}
