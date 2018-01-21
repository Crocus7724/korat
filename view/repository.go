package view

import (
	"github.com/crocus7724/korat/api"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var categories = []string{
	"Code",
	"Issues",
	"PullRequests",
}

func NewRepositoryView(r *api.Repository, ch func(c string, event *tcell.EventKey)) *tview.Table {
	table := tview.NewTable().
		Select(0, 0).
		SetSelectable(true, false)

	for i, category := range categories {
		table.SetCellSimple(i, 0, category)
	}

	table.SetTitle(string(r.Name)).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		row, _ := table.GetSelection()
		category := categories[row]
		ch(category, event)
		return event
	})
	return table
}
