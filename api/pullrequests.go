package api

import (
	"github.com/crocus7724/korat/model"
	"context"
	"github.com/shurcooL/githubql"
)

func GetViewerPullRequests(repository githubql.String, states []githubql.PullRequestState) (chan []model.PullRequest, chan error) {
	var query struct {
		Viewer struct {
			Repository struct {
				PullRequests struct {
					Nodes    []model.PullRequest
					PageInfo PageInfo
				} `graphql:"pullRequests(first: 100, orderBy:{field:CREATED_AT,direction:DESC}, states: $states, after: $cursor)"`
			} `graphql:"repository(name: $name)"`
		}
	}

	variables := map[string]interface{}{
		"name":   repository,
		"cursor": (*githubql.String)(nil),
		"states": states,
	}

	pChan := make(chan []model.PullRequest, 1)
	errChan := make(chan error)

	go func() {
		defer close(pChan)
		defer close(errChan)
		for {
			if err := client.Query(context.Background(), &query, variables); err != nil {
				errChan <- err
				return
			}

			pChan <- query.Viewer.Repository.PullRequests.Nodes
			if ! query.Viewer.Repository.PullRequests.PageInfo.HasNextPage {
				break
			}

			variables["cursor"] = query.Viewer.Repository.PullRequests.PageInfo.EndCursor
		}
	}()

	return pChan, errChan
}
