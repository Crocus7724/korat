package view

import (
	"github.com/crocus7724/tview"
	"fmt"
	"github.com/gdamore/tcell"
)

var (
	messageView *tview.TextView
	footerView  *tview.TextView
)

func init() {
	messageView = tview.NewTextView().
		SetRegions(true).
		SetScrollable(true).
		SetDynamicColors(true)

	footerView = tview.NewTextView().
		SetRegions(true).
		SetScrollable(true).
		SetDynamicColors(true)

	messageView.SetBackgroundColor(tcell.ColorDeepSkyBlue)
	messageView.SetTextColor(tcell.ColorLightGray)
}

func WriteMessageBox(message string) {
	messageView.Clear()
	fmt.Fprint(messageView, message)
	app.Draw()
}

func WriteFooterBox(message string) {
	footerView.Clear()
	fmt.Fprint(footerView, message)
	app.Draw()
}
