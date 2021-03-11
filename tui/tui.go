package tui

import (
	gh "github.com/nemotoy/afug/github"
	"github.com/rivo/tview"
)

type TUI struct {
	app    *tview.Application
	table  *tview.Table
	layout *tview.Grid
}

func NewAppWithWidget() *TUI {
	app := tview.NewApplication()

	table := tview.NewTable().
		SetBorders(false)

	layout := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0, 30).
		SetBorders(false).
		AddItem(table, 1, 0, 1, 3, 0, 100, false)

	return &TUI{
		app:    app,
		table:  table,
		layout: layout,
	}
}

func (t *TUI) Run() error {
	return t.app.SetRoot(t.layout, true).SetFocus(t.table).Run()
}

func (t *TUI) SetUsers(users []gh.User) *TUI {

	var id int = 1
	for _, user := range users {
		for _, repo := range user.StarredRepositories.Nodes {
			t.table.SetCell(id, 0,
				tview.NewTableCell(string(user.Name)).
					SetAlign(tview.AlignLeft))
			t.table.SetCell(id, 1,
				tview.NewTableCell(string(repo.PrimaryLanguage.Name)).
					SetAlign(tview.AlignLeft))
			t.table.SetCell(id, 2,
				tview.NewTableCell(repo.URL.String()).
					SetAlign(tview.AlignLeft))
		}
	}
	return t
}
