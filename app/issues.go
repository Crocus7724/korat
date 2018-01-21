package app

import (
	"github.com/crocus7724/korat/api"
	"github.com/crocus7724/korat/view"
	"github.com/gdamore/tcell"
)

func ViewerIssues(r *api.Repository) {
	issues, err := api.GetViewerIssues(r.Name)

	if err != nil {
		view.ShowError(err)
		return
	}

	v := view.NewIssuesView(issues, func(issue *api.Issue, event *tcell.EventKey) {
		if event.Key() == tcell.KeyEnter {

		} else if event.Rune() == 'b' {
			OpenUrl(string(issue.Url))
		}
	})
	view.PushPage(v)
}
