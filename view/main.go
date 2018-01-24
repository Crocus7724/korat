package view

import (
	"github.com/rivo/tview"
	"fmt"
	"github.com/gdamore/tcell"
)

var (
	app       *tview.Application
	pages     *tview.Pages
	pageCount = 0
	wrapper   *tview.Flex
	//main      *tview.Flex
	footer *tview.Flex
)

type View interface {
	View() tview.Primitive
}

func Init() {
	app = tview.NewApplication()
	pages = tview.NewPages()
	wrapper = tview.NewFlex()

	footer = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(messageView, 0, 1, false).
		AddItem(footerView, 0, 1, false)

	wrapper.SetDirection(tview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(footer, 2, 1, false)

}

func Start() {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			PopPage()
		}

		return event
	})

	if err := app.SetRoot(wrapper, true).SetFocus(wrapper).Run(); err != nil {
		panic(err)
	}
}

func PushPage(view View) {
	pageCount++
	pages.AddAndSwitchToPage(fmt.Sprintf("page-%d", pageCount), view.View(),
		true)

	app.Draw()

	messageView.Clear()
	footerView.Clear()

	app.SetFocus(pages)
}

func PopPage() {
	if pageCount == 1 {
		app.Stop()
	}
	pages.SwitchToPage(fmt.Sprintf("page-%d", pageCount-1))
	pages.RemovePage(fmt.Sprintf("page-%d", pageCount))
	pageCount--
	messageView.Clear()
	footerView.Clear()
}

func SetEmptyCell(t *tview.Table, message string) {
	t.SetCellSimple(0, 0, message).
		SetSelectable(false, false)
	app.Draw()
}
