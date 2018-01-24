package view

import (
	"github.com/crocus7724/tview"
	"github.com/gdamore/tcell"
	"github.com/crocus7724/korat/util"
	"github.com/crocus7724/korat/model"
	"sync"
)

type RepositoriesView struct {
	count int
	mutex *sync.Mutex
	view  *tview.Table
}

func NewRepositoriesView(ch func(selected int, event *tcell.EventKey)) *RepositoriesView {
	r := RepositoriesView{
		view: tview.NewTable().
			SetSelectable(true, false),
		mutex: &sync.Mutex{},
		count: 0,
	}
	r.view.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if r.count != 0 {
			r, _ := r.view.GetSelection()
			ch(r, event)
		}
		return event
	})

	return &r
}

func (r *RepositoriesView) AddRepositoriesCells(rs []model.Repository) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for _, repository := range rs {
		uc := tview.NewTableCell(util.GetTimeString(string(repository.UpdatedAt)))
		uc.SetTextColor(tcell.ColorRoyalBlue)
		r.view.SetCell(r.count, 0, uc)

		nc := tview.NewTableCell(string(repository.Name))
		r.view.SetCell(r.count, 1, nc)

		r.count++
	}
	app.Draw()
}

func (r *RepositoriesView) Count() int {
	return r.count
}

func (r *RepositoriesView) View() tview.Primitive {
	return r.view
}

func (r *RepositoriesView) SetSelectionChangedFunc(changed func(row int)) {
	r.view.SetSelectionChangedFunc(func(row, column int) {
		if r.count > 0 {
			changed(row)
		}
	})

}
