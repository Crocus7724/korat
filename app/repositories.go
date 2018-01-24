package app

import (
	"github.com/crocus7724/korat/api"
	"github.com/crocus7724/korat/view"
	"log"
	"github.com/gdamore/tcell"
	"github.com/crocus7724/korat/model"
	"github.com/rivo/tview"
)

func ViewerRepositories() {
	rChan, errChan := api.GetViewerRepositories()
	var rs []model.Repository

	v := view.NewRepositoriesView(func(selected int, event *tcell.EventKey) {
		r := rs[selected]
		if event.Key() == tcell.KeyEnter {
			ViewerRepository(&r)
		} else if event.Rune() == 'b' {
			OpenUrl(string(r.Url))
		} else if event.Rune() == 'q' {
			view.PopPage()
		}
	})

	v.SetSelectionChangedFunc(func(row int) {
		r := rs[row]
		view.WriteMessageBox(string(r.Description))
	})

	go func() {
		for {
			select {
			case r, ok :=<-rChan:
				if ok {
					rs = append(rs, r...)
					v.AddRepositoriesCells(r)
				} else {
					if len(rs) == 0 {
						view.SetEmptyCell(v.View().(*tview.Table), "There aren't repositories")
					}
					return
				}
			case err, ok := <-errChan:
				if ok {
					log.Fatal(err)
				}
			}
		}
	}()
	view.PushPage(v)
}
