package view

import ui "github.com/gizak/termui"

func Layout() error {
	if err := ui.Init(); err != nil {
		return err
	}

	defer ui.Close()

	list := ui.NewList()
	list.Border = true
	list.Items = []string{
		"hoge",
		"foo",
	}
	list.Height = ui.TermHeight() / 2
	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(12, 0, list),
		),
	)

	ui.Body.Align()

	ui.Render(ui.Body)
	ui.Handle("/sys/kbd/q", func(event ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("sys/wnd/resize", func(event ui.Event) {
		list.Height = ui.TermHeight() / 2
		ui.Body.Width = ui.TermWidth()
		ui.Body.Align()
		ui.Clear()
		ui.Render(ui.Body)
	})

	ui.Loop()
	return nil
}
