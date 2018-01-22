package api

import (
	"context"
	"github.com/shurcooL/githubql"
	"github.com/crocus7724/korat/model"
)

type Repositories struct {
}

func GetViewerRepositories() ([]model.Repository, error) {
	var query struct {
		Viewer struct {
			Repositories struct {
				Nodes []model.Repository
				PageInfo struct {
					HasNextPage githubql.Boolean
					EndCursor   githubql.String
				}
			} `graphql:"repositories(first:50, orderBy: {field: UPDATED_AT, direction: DESC}, after: $repositoriesCursor)"`
		}
	}
	variables := map[string]interface{}{
		"repositoriesCursor": (*githubql.String)(nil),
	}

	var repositories []model.Repository

	for {
		if err := client.Query(context.Background(), &query, variables); err != nil {
			return nil, err
		}

		repositories = append(repositories, query.Viewer.Repositories.Nodes...)

		if !query.Viewer.Repositories.PageInfo.HasNextPage {
			break
		}
		variables["repositoriesCursor"] = githubql.NewString(query.Viewer.Repositories.PageInfo.EndCursor)
	}

	return repositories, nil
}
