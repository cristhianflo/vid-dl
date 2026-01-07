package tui

import "github.com/charmbracelet/lipgloss"

var (
	TitleStyle        = lipgloss.NewStyle().Bold(true)
	SectionTitleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFC940")).Bold(true)
	IconStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFC940")).Bold(true)
	FaintStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#888888"))
	LabelStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(false)

	FormatStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("#D6D6D6"))
	HighlightStyle      = FormatStyle.Copy().Bold(true).Foreground(lipgloss.Color("#FFC940"))
	SelectorStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFC940")).Bold(true)
	BulletStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFC940")).Bold(true)
	SelectedBulletStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFC940")).Bold(true)

	HelpKeyStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#fff")).Background(lipgloss.Color("#222")).Padding(0, 1)
	HelpDescStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#888")).Padding(0, 1)
	HelpStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#b9b9b9")).Margin(1, 0, 0, 0)

	VertLineStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFC940"))
)
