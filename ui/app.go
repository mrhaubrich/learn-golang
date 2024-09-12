package ui

import (
	"github.com/rivo/tview"
)

// AppUI represents the terminal UI for the app
type AppUI struct {
	App      *tview.Application
	TextView *tview.TextView
	Pages    *tview.Pages
}

// NewAppUI is a constructor that initializes the UI elements
func NewAppUI() *AppUI {
	app := tview.NewApplication()

	// Create a simple text view to display information
	textView := tview.NewTextView().
		SetText("Welcome to Go OOP Terminal App!").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	// Create pages to manage different views
	pages := tview.NewPages().
		AddPage("main", textView, true, true)

	return &AppUI{
		App:      app,
		TextView: textView,
		Pages:    pages,
	}
}

// Run starts the application
func (ui *AppUI) Run() error {
	return ui.App.SetRoot(ui.Pages, true).Run()
}

// AddText allows dynamic updates to the TextView and queues updates to the UI
func (ui *AppUI) AddText(text string) {
	// Ensure that the update is performed within the application's event loop
	ui.App.QueueUpdateDraw(func() {
		ui.TextView.SetText(text)
	})
}
