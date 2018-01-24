package view

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell"
	"fmt"
	"github.com/crocus7724/korat/model"
	"github.com/crocus7724/korat/util"
	"sync"
)

type IssuesView struct {
	view  *tview.Table
	count int
	mutex *sync.Mutex
}

func NewIssuesView(ch func(selected int, event *tcell.EventKey)) *IssuesView {
	iv := IssuesView{
		view: tview.NewTable().
			SetSelectable(true, false),
		count: 0,
		mutex: &sync.Mutex{},
	}

	iv.view.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if iv.count != 0 {
			row, _ := iv.view.GetSelection()
			ch(row, event)
		}
		return event
	})
	return &iv
}

func (i *IssuesView) AddIssueCells(issues []model.Issue) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	for _, issue := range issues {
		numberCell := tview.NewTableCell(fmt.Sprintf("#%d", issue.Number))
		i.view.SetCell(i.count, 0, numberCell)

		titleCell := tview.NewTableCell(util.ReplaceBrackets(issue.Title))
		i.view.SetCell(i.count, 1, titleCell)

		i.count++
	}
	app.Draw()
}

func (i *IssuesView) View() tview.Primitive {
	return i.view
}

func (i *IssuesView) Count() int {
	return i.count
}
