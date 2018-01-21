package app

import (
	"github.com/crocus7724/korat/api"
	"github.com/crocus7724/korat/view"
	"log"
	"github.com/gdamore/tcell"
)

func ViewerRepositories() {
	repositories, err := api.GetViewerRepositories()

	if err != nil {
		log.Fatalf("error: %v", err)
	}
	v := view.NewRepositoriesView(repositories,
		func(repository *api.Repository, event *tcell.EventKey) {
			if event.Key() == tcell.KeyEnter {
				ViewerRepository(repository)
			} else if event.Rune() == 'b' {
				OpenUrl(string(repository.Url))
			} else if event.Rune() == 'q' {
				view.PopPage()
			}
		})
	view.PushPage(v)
}
