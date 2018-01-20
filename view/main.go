package view

import "github.com/rivo/tview"

func Layout() error {
	app := tview.NewApplication()
	list := tview.NewList().
		AddItem("hoge", "foo", 'a', nil).
		AddItem("foo", "hoge", 'b', nil).
		AddItem("Quit", "exit", 'q', func() {
		app.Stop()
	})

	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
	return nil
}
