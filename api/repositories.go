package api

import (
	"context"
	"github.com/shurcooL/githubql"
	"github.com/crocus7724/korat/model"
)

func GetViewerRepositories() (chan []model.Repository, chan error) {
	var query struct {
		Viewer struct {
			Repositories struct {
				Nodes []model.Repository
				PageInfo PageInfo
			} `graphql:"repositories(first:50, orderBy: {field: UPDATED_AT, direction: DESC}, after: $repositoriesCursor)"`
		}
	}
	variables := map[string]interface{}{
		"repositoriesCursor": (*githubql.String)(nil),
	}

	rChan := make(chan []model.Repository)
	eChan := make(chan error)

	go func() {
		defer close(rChan)
		defer close(eChan)
		for {
			if err := client.Query(context.Background(), &query, variables); err != nil {
				eChan <- err
				break
			}

			rChan <- query.Viewer.Repositories.Nodes
			if !query.Viewer.Repositories.PageInfo.HasNextPage {
				break
			}
			variables["repositoriesCursor"] = githubql.NewString(query.Viewer.Repositories.PageInfo.EndCursor)
		}
	}()

	return rChan, eChan
}
