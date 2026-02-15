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

	headerStyle := styles.Bold().Foreground(t.Primary())
	linkStyle := styles.Regular().Foreground(t.MarkdownLink())
	mutedStyle := baseStyle.Foreground(t.TextMuted())

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		headerStyle.Render("Documentation"),
		"",
		"GitHub Repository: ",
		linkStyle.Render("https://github.com/Get-Blu/blu-code"),
		"",
		"Local Documentation Folder: ",
		linkStyle.Render("./docs/"),
		"",
		"Key Documentation Files:",
		" - Introduction: "+linkStyle.Render("docs/introduction.md"),
		" - Getting Started: "+linkStyle.Render("docs/getting-started.md"),
		" - Configuration: "+linkStyle.Render("docs/configuration.md"),
		" - Best Practices: "+linkStyle.Render("docs/guides/best-practices.md"),
		"",
		mutedStyle.Render("Press esc or q to close"),
	)

	return baseStyle.Padding(1, 2).
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
