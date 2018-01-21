package view

import "fmt"

func ShowError(err error) {
	footerView.Clear()
	fmt.Fprintf(footerView, "error: [red]%v", err)
}
