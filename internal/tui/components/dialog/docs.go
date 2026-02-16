package dialog

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/Get-Blu/blu-code/internal/tui/layout"
	"github.com/Get-Blu/blu-code/internal/tui/styles"
	"github.com/Get-Blu/blu-code/internal/tui/theme"
	"github.com/Get-Blu/blu-code/internal/tui/util"
)

type DocsDialogCmp interface {
	tea.Model
	layout.Bindings
}

type CloseDocsDialogMsg struct{}

type docsDialogCmp struct{}

func (d *docsDialogCmp) Init() tea.Cmd {
	return nil
}

func (d *docsDialogCmp) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, key.NewBinding(key.WithKeys("esc", "q", "ctrl+d"))):
			return d, util.CmdHandler(CloseDocsDialogMsg{})
		}
	}
	return d, nil
}

func (d *docsDialogCmp) View() string {
	t := theme.CurrentTheme()
	baseStyle := styles.BaseStyle()

	headerStyle := baseStyle.Copy().Bold(true).Foreground(t.Primary())
	linkStyle := baseStyle.Copy().Foreground(t.MarkdownLink())
	mutedStyle := baseStyle.Copy().Foreground(t.TextMuted())
	labelStyle := baseStyle.Copy()

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		headerStyle.Render("Documentation"),
		labelStyle.Render(""),
		labelStyle.Render("GitHub Repository:"),
		linkStyle.Render("https://github.com/Get-Blu/blu-code"),
		labelStyle.Render(""),
		labelStyle.Render("Local Documentation Folder:"),
		linkStyle.Render("./docs/"),
		labelStyle.Render(""),
		labelStyle.Render("Key Documentation Files:"),
		labelStyle.Render(" - Introduction: "+linkStyle.Render("docs/introduction.md")),
		labelStyle.Render(" - Getting Started: "+linkStyle.Render("docs/getting-started.md")),
		labelStyle.Render(" - Configuration: "+linkStyle.Render("docs/configuration.md")),
		labelStyle.Render(" - Best Practices: "+linkStyle.Render("docs/guides/best-practices.md")),
		labelStyle.Render(""),
		mutedStyle.Render("Press esc or q to close"),
	)

	// Calculate the actual content width to ensure the background fills the entire box
	contentWidth := lipgloss.Width(content)

	return baseStyle.
		Padding(1, 2).
		Width(contentWidth + 4). // Add padding to the width
		Border(lipgloss.RoundedBorder()).
		BorderBackground(t.Background()).
		BorderForeground(t.TextMuted()).
		Render(content)
}

func (d *docsDialogCmp) BindingKeys() []key.Binding {
	return []key.Binding{
		key.NewBinding(
			key.WithKeys("esc", "q"),
			key.WithHelp("esc/q", "close"),
		),
	}
}

func NewDocsDialogCmp() DocsDialogCmp {
	return &docsDialogCmp{}
}
