package tui

import (
	"fmt"
	"github.com/cristhianflo/vid-dl/internal/downloader"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type downloadDoneMsg struct{}
type downloadErrMsg struct {
	err error
}

type model struct {
	formats    []downloader.Format
	downloader downloader.Downloader
	form       *huh.Form
	status     string
}

func NewTui(model model) *tea.Program {
	return tea.NewProgram(model)
}

func NewModel(d downloader.Downloader) (*model, error) {
	result, err := d.GetFormats()
	if err != nil {
		return nil, err
	}

	var m model
	var opts []huh.Option[int]
	for i, format := range result.Formats {
		label := fmt.Sprintf("%s | %s | %s", format.Resolution, format.Ext, humanFileSize(format.Filesize))
		opts = append(opts, huh.NewOption(label, i))
	}

	// Check if we have any formats
	if len(opts) == 0 {
		return nil, fmt.Errorf("no video formats available")
	}

	m.form = huh.NewForm(
		huh.NewGroup(
			// Display title
			huh.NewNote().Title("Title").Description(result.Title),
			// Display video ID
			huh.NewNote().Title("ID").Description(result.ID),
			// Display format selector
			huh.NewSelect[int]().
				Key("format").
				Title("Format").
				Options(opts...),
		),
	)

	m.formats = result.Formats
	m.downloader = d
	return &m, nil
}

func (m model) Init() tea.Cmd {
	return m.form.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		// If download is done, pressing enter/esc/q exits
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Interrupt
		case "esc", "q":
			return m, tea.Quit
		}

	case downloadErrMsg:
		m.status = fmt.Sprintf("Download failed: %v", msg.err)
		return m, tea.Quit

	case downloadDoneMsg:
		m.status = "Download completed successfully!"
		return m, tea.Quit
	}

	var cmds []tea.Cmd

	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		selectedFormat := m.form.GetInt("format")

		if selectedFormat >= 0 && selectedFormat < len(m.formats) {
			format := &m.formats[selectedFormat]
			// Set status to show download starting
			m.status = "Downloading..."
			cmds = append(cmds, func() tea.Msg {
				err := m.downloader.DownloadVideo(format)
				if err != nil {
					return downloadErrMsg{err: err}
				}
				return downloadDoneMsg{}
			})
		}
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var out string
	if m.status != "" {
		out += SectionTitleStyle.Render(m.status) + "\n\n"
	}
	// Render the huh form as the body
	out += m.form.View()
	// Footer help actions
	out += "\n" + HelpKeyStyle.Render("ctrl") + HelpDescStyle.Render("↵ submit")
	out += HelpKeyStyle.Render("tab") + HelpDescStyle.Render("navigate")
	out += HelpKeyStyle.Render("^k") + HelpDescStyle.Render("actions")
	return out
}
