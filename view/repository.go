package view

import (
	"github.com/gdamore/tcell"
	"github.com/crocus7724/tview"
	"github.com/crocus7724/korat/model"
)

type RepositoryView struct {
	view *tview.Table
}

var categories = []string{
	"Code",
	"Issues",
	"PullRequests",
}

func NewRepositoryView(r *model.Repository, ch func(c string, event *tcell.EventKey)) *RepositoryView {
	rv := RepositoryView{
		view: tview.NewTable().
			Select(0, 0).
			SetSelectable(true, false),
	}
	for i, category := range categories {
		rv.view.SetCellSimple(i, 0, category)
	}

	rv.view.SetTitle(string(r.Name)).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		row, _ := rv.view.GetSelection()
		category := categories[row]
		ch(category, event)
		return event
	})

	return &rv
}

func (r *RepositoryView) View() tview.Primitive {
	return r.view
}
