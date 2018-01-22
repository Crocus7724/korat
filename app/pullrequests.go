package app

import (
	"github.com/crocus7724/korat/model"
	"github.com/crocus7724/korat/api"
	"github.com/shurcooL/githubql"
	"github.com/crocus7724/korat/view"
	"github.com/gdamore/tcell"
)

func ViewerPullRequests(repository *model.Repository) {
	prs, err := api.GetViewerPullRequests(repository.Name, []githubql.PullRequestState{
		//githubql.PullRequestStateOpen,
	})

	if err != nil {
		view.ShowError(err)
		return
	}

	v := view.NewPullRequestsView(prs, func(pr *model.PullRequest, event *tcell.EventKey) {
		if event.Key() == tcell.KeyEnter {

		} else if event.Rune() == 'b' {
			OpenUrl(string(pr.Url))
		}
	})

	view.PushPage(v)
}
