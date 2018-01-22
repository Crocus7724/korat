package view

import (
	"github.com/crocus7724/korat/model"
	"github.com/rivo/tview"
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/crocus7724/korat/util"
)

func NewPullRequestsView(prs []model.PullRequest, ch func(pr *model.PullRequest, event *tcell.EventKey)) *tview.Table {
	if len(prs) > 0 {
		table := tview.NewTable().
			SetSelectable(true, false).
			Select(0, 2)

		for i, pr := range prs {
			numberCell := tview.NewTableCell(fmt.Sprintf("#%d", pr.Number))
			table.SetCell(i, 0, numberCell)

			authorCell := tview.NewTableCell(string(pr.Author.Login))
			authorCell.SetTextColor(tcell.ColorGreen)
			table.SetCell(i, 1, authorCell)

			titleCell := tview.NewTableCell(util.ReplaceBrackets(pr.Title))

			table.SetCell(i, 2, titleCell)
		}

		table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			row, _ := table.GetSelection()
			pr := prs[row]

			ch(&pr, event)
			return event
		})

		return table
	} else {
		return tview.NewTable().
			SetSelectable(false, false).
			SetCellSimple(0, 0, "There aren't pullrequests")
	}

}
