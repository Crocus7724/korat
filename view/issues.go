package view

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell"
	"fmt"
	"github.com/crocus7724/korat/model"
)

func NewIssuesView(issues []model.Issue, ch func(issue *model.Issue, event *tcell.EventKey)) *tview.Table {
	if len(issues) > 0 {
		table := tview.NewTable().
			SetSelectable(true, false)

		for i, issue := range issues {
			numberCell := tview.NewTableCell(fmt.Sprintf("#%d", issue.Number))
			table.SetCell(i, 0, numberCell)

			titleCell := tview.NewTableCell(string(issue.Title))
			table.SetCell(i, 1, titleCell)
		}
		table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			row, _ := table.GetSelection()
			issue := issues[row]
			ch(&issue, event)

			return event
		})
		return table
	} else {
		return tview.NewTable().
			SetSelectable(false, false).
			SetCellSimple(0, 0, "There aren't issues")
	}

}
