package view

import (
	"github.com/crocus7724/korat/model"
	"github.com/rivo/tview"
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/crocus7724/korat/util"
	"sync"
)

type PullRequestsView struct {
	view  *tview.Table
	count int
	mutex *sync.Mutex
}

func NewPullRequestsView(ch func(selected int, event *tcell.EventKey)) *PullRequestsView {
	p := PullRequestsView{
		view: tview.NewTable().
			SetSelectable(true, false),
		count: 0,
		mutex: &sync.Mutex{},
	}

	p.view.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if p.count != 0 {
			r, _ := p.view.GetSelection()
			ch(r, event)
		}
		return event
	})
	return &p

}

func (p *PullRequestsView) AddCells(prs []model.PullRequest) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for i, pr := range prs {
		numberCell := tview.NewTableCell(fmt.Sprintf("#%d", pr.Number))
		p.view.SetCell(i, 0, numberCell)

		authorCell := tview.NewTableCell(string(pr.Author.Login))
		authorCell.SetTextColor(tcell.ColorGreen)
		p.view.SetCell(i, 1, authorCell)

		titleCell := tview.NewTableCell(util.ReplaceBrackets(pr.Title))

		p.view.SetCell(i, 2, titleCell)
	}
	app.Draw()
}

func (p *PullRequestsView) View() tview.Primitive {
	return p.view
}

func (p *PullRequestsView) Count() int {
	return p.count
}
