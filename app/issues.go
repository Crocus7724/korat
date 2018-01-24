package app

import (
	"github.com/crocus7724/korat/api"
	"github.com/crocus7724/korat/model"
	"github.com/crocus7724/korat/view"
	"github.com/gdamore/tcell"
	"github.com/crocus7724/tview"
	"github.com/shurcooL/githubql"
)

func ViewerIssues(r *model.Repository) {
	iChan, errChan := api.GetViewerIssues(r.Name, []githubql.IssueState{
		githubql.IssueStateOpen,
		githubql.IssueStateClosed,
	})

	var issues []model.Issue
	v := view.NewIssuesView(func(selected int, event *tcell.EventKey) {
		i := issues[selected]
		if event.Key() == tcell.KeyEnter {

		} else if event.Rune() == 'b' {
			OpenUrl(string(i.Url))
		}
	})

	go func() {
		for {
			select {
			case i, ok := <-iChan:
				if ok {
					issues = append(issues, i...)
					v.AddIssueCells(i)
				} else {
					if len(issues) == 0 {
						view.SetEmptyCell(v.View().(*tview.Table), "There aren't issues")
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
