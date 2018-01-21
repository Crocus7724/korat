package view

import (
	"github.com/rivo/tview"
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

	messageView.SetBackgroundColor(tcell.ColorGreen)
}

func WriteMessageBox(message string) {
	messageView.Clear()
	fmt.Fprintf(messageView, message)
	app.Draw()
}

func WriteFooterBox(message string) {
	footerView.Clear()
	fmt.Fprintf(footerView, message)
	app.Draw()
}
