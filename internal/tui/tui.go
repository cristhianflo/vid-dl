package tui

import (
	"cristhianflo/vid-dl/internal/downloader"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	title    string
	url      string
	formats  []downloader.Format
	choices  []string
	cursor   int
	selected int
}

func NewModel(result *downloader.Video) model {
	var model model

	for _, choice := range result.Formats {
		formatString := fmt.Sprintf("%s | %s | %s", choice.Resolution, choice.Ext, humanFileSize(choice.Filesize))
		model.choices = append(model.choices, formatString)
	}
	model.formats = result.Formats
	return model
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			m.selected = m.cursor
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	// The header
	s := " Video downloader\n\n"

	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = " ▶" // cursor!
		}

		// Is this choice selected?
		checked := "○" // not selected
		if m.selected == i {
			checked = "●" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s %s %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
