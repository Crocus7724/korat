package api

import (
	"github.com/crocus7724/korat/model"
	"context"
	"github.com/shurcooL/githubql"
)


func GetViewerPullRequests(repository githubql.String, states []githubql.PullRequestState) ([]model.PullRequest, error) {
	var query struct {
		Viewer struct {
			Repository struct {
				PullRequests struct {
					Nodes    []model.PullRequest
					PageInfo PageInfo
				} `graphql:"pullRequests(first: 100, orderBy:{field:UPDATED_AT,direction:DESC}, states: $states, after: $cursor)"`
			} `graphql:"repository(name: $name)"`
		}
	}

	variables := map[string]interface{}{
		"name":   repository,
		"cursor": (*githubql.String)(nil),
		"states": states,
	}

	var prs []model.PullRequest
	for {
		if err := client.Query(context.Background(), &query, variables); err != nil {
			return nil, err
		}

		prs = append(prs, query.Viewer.Repository.PullRequests.Nodes...)
		if ! query.Viewer.Repository.PullRequests.PageInfo.HasNextPage {
			break
		}

		variables["cursor"] = query.Viewer.Repository.PullRequests.PageInfo.EndCursor
	}

	return prs, nil
}
