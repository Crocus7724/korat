package view

import (
	"github.com/crocus7724/korat/api"
	"github.com/rivo/tview"
	"github.com/gdamore/tcell"
	"github.com/crocus7724/korat/util"
)

func NewRepositoriesView(rs []api.Repository, ch func(repository *api.Repository, event *tcell.EventKey)) *tview.Table {
	count := len(rs)

	if count > 0 {
		table := tview.NewTable()

		table.SetSelectable(true, false)
		table.SetTitle("Repositories")

		table.Select(0, 0)

		for i, repository := range rs {
			updatedAtCell := tview.NewTableCell(util.GetTimeString(string(repository.UpdatedAt)))
			updatedAtCell.SetTextColor(tcell.ColorRoyalBlue)
			table.SetCell(i, 0, updatedAtCell)

			nameCell := tview.NewTableCell(string(repository.Name))
			table.SetCell(i, 1, nameCell)
		}

		table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			row, _ := table.GetSelection()
			r := rs[row]

			ch(&r, event)

			return event
		})

		table.SetSelectionChangedFunc(func(row, column int) {
			d := rs[row].Description
			WriteMessageBox(string(d))
		})

		d := rs[0].Description
		WriteMessageBox(string(d))

		return table
	} else {
		table := tview.NewTable()
		table.SetSelectable(false, false)
		table.SetCellSimple(0, 0, "There aren't repositories.")

		return table
	}
}
