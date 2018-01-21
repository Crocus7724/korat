package view

import (
	"github.com/crocus7724/korat/api"
	"github.com/rivo/tview"
	"github.com/gdamore/tcell"
)

func NewRepositoriesView(rs []api.Repository, ch func(repository *api.Repository, event *tcell.EventKey)) *tview.Table {
	table := tview.NewTable()

	table.SetSelectable(true, false)
	table.SetTitle("Repositories")

	table.Select(0, 0)

	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		row, _ := table.GetSelection()
		r := rs[row]

		ch(&r, event)

		return event
	})
	for i, repository := range rs {
		nameCell := tview.NewTableCell(string(repository.Name))
		table.SetCell(i, 0, nameCell)
		descriptionCell := tview.NewTableCell(string(repository.Description))
		descriptionCell.SetSelectable(false)
		descriptionCell.SetTextColor(tcell.ColorGray)
		table.SetCell(i, 1, descriptionCell)
	}
	return table
}
