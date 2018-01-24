package app

import (
	"github.com/crocus7724/korat/model"
	"github.com/crocus7724/korat/api"
	"github.com/shurcooL/githubql"
	"github.com/crocus7724/korat/view"
	"github.com/gdamore/tcell"
	"github.com/crocus7724/tview"
)

func ViewerPullRequests(repository *model.Repository) {
	pChan, errChan := api.GetViewerPullRequests(repository.Name, []githubql.PullRequestState{
		githubql.PullRequestStateOpen,
		githubql.PullRequestStateClosed,
	})
	var prs []model.PullRequest
	v := view.NewPullRequestsView(func(selected int, event *tcell.EventKey) {
		p := prs[selected]
		if event.Key() == tcell.KeyEnter {

		} else if event.Rune() == 'b' {
			OpenUrl(string(p.Url))
		}
	})
	go func() {
		for {
			select {
			case ps, ok := <-pChan:
				if ok {
					v.AddCells(ps)
					prs = append(prs, ps...)
				} else {
					if len(prs) == 0 {
						view.SetEmptyCell(v.View().(*tview.Table), "There aren't PullRequests")
					}

					return
				}
			case err, ok := <-errChan:
				if ok {
					view.ShowError(err)
					return
				}
			}
		}
	}()

	view.PushPage(v)
}
