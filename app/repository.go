package app

import (
	"github.com/crocus7724/korat/view"
	"github.com/gdamore/tcell"
	"path"
	"github.com/crocus7724/korat/model"
)

func ViewerRepository(r *model.Repository) {
	view.PushPage(view.NewRepositoryView(r, func(c string, event *tcell.EventKey) {
		if event.Rune() == 'b' {
			switch c {
			case "Issues":
				OpenUrl(path.Join(string(r.Url), "issues"))
			case "PullRequests":
				OpenUrl(path.Join(string(r.Url), "pulls"))
			}
		} else if event.Key() == tcell.KeyEnter {
			switch c {
			case "Issues":
				ViewerIssues(r)
			case "PullRequests":
				ViewerPullRequests(r)
			}
		}
	}))
}
