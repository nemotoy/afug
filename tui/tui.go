package tui

import (
	tcell "github.com/gdamore/tcell/v2"
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
	app.SetInputCapture(func(e *tcell.EventKey) *tcell.EventKey {
		switch e.Key() {
		case tcell.KeyEnter:
			app.Stop()
			return nil
		}
		return e
	})

	table := tview.NewTable().
		SetBorders(false).
		SetFixed(1, 1)

	layout := tview.NewGrid().
		SetBorders(false).
		AddItem(table, 0, 0, 3, 3, 0, 100, false)

	return &TUI{
		app:    app,
		table:  table,
		layout: layout,
	}
}

func (t *TUI) Run() error {
	return t.app.SetRoot(t.layout, true).SetFocus(t.table).Run()
}

func (t *TUI) SetTableFrame() *TUI {
	t.table.SetCell(0, 1,
		tview.NewTableCell("ID").
			SetAlign(tview.AlignLeft))
	t.table.SetCell(0, 2,
		tview.NewTableCell("Language").
			SetAlign(tview.AlignLeft))
	t.table.SetCell(0, 3,
		tview.NewTableCell("URL").
			SetAlign(tview.AlignLeft))
	return t
}

func (t *TUI) SetUsers(users []gh.User) *TUI {

	var row int = 1
	for _, user := range users {
		for _, repo := range user.StarredRepositories.Nodes {
			t.table.SetCell(row, 1,
				tview.NewTableCell(string(user.Name)).
					SetAlign(tview.AlignLeft))
			t.table.SetCell(row, 2,
				tview.NewTableCell(string(repo.PrimaryLanguage.Name)).
					SetAlign(tview.AlignLeft))
			t.table.SetCell(row, 3,
				tview.NewTableCell(repo.URL.String()).
					SetAlign(tview.AlignLeft))
			row++
		}
	}
	return t
}

func (t *TUI) SetStub() *TUI {
	users := []struct {
		ID   string
		Name string
		URL  string
	}{
		{"aaa", "aaa", "https://google.com"},
		{"aaa", "aaa", "https://google.com"},
		{"aaa", "aaa", "https://google.com"},
		{"aaa", "aaa", "https://google.com"},
		{"aaa", "aaa", "https://google.com"},
	}
	var row int = 1
	for _, user := range users {
		t.table.SetCell(row, 1,
			tview.NewTableCell(user.ID).
				SetAlign(tview.AlignLeft))
		t.table.SetCell(row, 2,
			tview.NewTableCell(user.Name).
				SetAlign(tview.AlignLeft))
		t.table.SetCell(row, 3,
			tview.NewTableCell(user.URL).
				SetAlign(tview.AlignLeft))
		row++
	}
	return t
}
