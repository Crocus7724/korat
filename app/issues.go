package app

import (
	"github.com/crocus7724/korat/api"
	"github.com/crocus7724/korat/view"
	"github.com/gdamore/tcell"
	"github.com/crocus7724/korat/model"
	"github.com/shurcooL/githubql"
)

func ViewerIssues(r *model.Repository) {
	issues, err := api.GetViewerIssues(r.Name, []githubql.IssueState{})

	if err != nil {
		view.ShowError(err)
		return
	}

	v := view.NewIssuesView(issues, func(issue *model.Issue, event *tcell.EventKey) {
		if event.Key() == tcell.KeyEnter {

		} else if event.Rune() == 'b' {
			OpenUrl(string(issue.Url))
		}
	})
	view.PushPage(v)
}
