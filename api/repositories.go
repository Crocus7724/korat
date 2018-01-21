package api

import (
	"context"
	"github.com/shurcooL/githubql"
)

type Repository struct {
	Name        githubql.String
	Description githubql.String
	Url         githubql.String
}

type Repositories struct {
	Nodes []Repository
	PageInfo struct {
		HasNextPage githubql.Boolean
		EndCursor   githubql.String
	}
}

func GetViewerRepositories() ([]Repository, error) {
	var query struct {
		Viewer struct {
			Repositories Repositories `graphql:"repositories(first:50, orderBy: {field: UPDATED_AT, direction: DESC}, after: $repositoriesCursor)"`
		}
	}
	variables := map[string]interface{}{
		"repositoriesCursor": (*githubql.String)(nil),
	}

	var repositories []Repository

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
