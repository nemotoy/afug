package tui

import (
	"github.com/rivo/tview"
)

type TUI struct {
	app   *tview.Application
	table *tview.Table
}

func NewAppWithWidget() *TUI {
	app := tview.NewApplication()

	table := tview.NewTable().
		SetBorders(false).
		SetCell(0, 0,
			tview.NewTableCell("ID").
				SetAlign(tview.AlignLeft)).
		SetCell(0, 1,
			tview.NewTableCell("Language").
				SetAlign(tview.AlignLeft)).
		SetCell(0, 2,
			tview.NewTableCell("URL").
				SetAlign(tview.AlignLeft))

	return &TUI{
		app:   app,
		table: table,
	}
}
