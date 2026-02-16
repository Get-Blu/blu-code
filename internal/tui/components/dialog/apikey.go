package dialog

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/Get-Blu/blu-code/internal/llm/models"
	"github.com/Get-Blu/blu-code/internal/tui/theme"
	"github.com/Get-Blu/blu-code/internal/tui/util"
)

type APIKeySelectedMsg struct {
	Provider models.ModelProvider
	APIKey   string
}

type CloseAPIKeyDialogMsg struct{}

type APIKeyDialogCmp struct {
	width, height int
	input         textinput.Model
	provider      models.ModelProvider
}

func NewAPIKeyDialogCmp(provider models.ModelProvider) APIKeyDialogCmp {
	t := theme.CurrentTheme()
	ti := textinput.New()
	ti.Placeholder = "Enter API Key..."
	ti.EchoMode = textinput.EchoPassword
	ti.EchoCharacter = '•'
	ti.Focus()
	ti.Width = 40
	ti.PromptStyle = ti.PromptStyle.Foreground(t.Primary())
	ti.TextStyle = ti.TextStyle.Foreground(t.Text())

	return APIKeyDialogCmp{
		input:    ti,
		provider: provider,
	}
}

func (m APIKeyDialogCmp) Init() tea.Cmd {
	return textinput.Blink
}

func (m APIKeyDialogCmp) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.input.Value() != "" {
				return m, util.CmdHandler(APIKeySelectedMsg{
					Provider: m.provider,
					APIKey:   m.input.Value(),
				})
			}
		case "esc":
			return m, util.CmdHandler(CloseAPIKeyDialogMsg{})
		}
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m APIKeyDialogCmp) View() string {
	t := theme.CurrentTheme()
	baseStyle := styles.BaseStyle()

	title := baseStyle.Copy().
		Foreground(t.Primary()).
		Bold(true).
		Padding(0, 0, 1).
		Render("Configure API Key")

	providerLabel := baseStyle.Copy().
		Foreground(t.TextMuted()).
		Padding(0, 0, 1).
		Render("Provider: " + string(m.provider))

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		providerLabel,
		m.input.View(),
	)

	// Calculate content width to ensure background fills
	contentWidth := lipgloss.Width(content)

	return baseStyle.
		Padding(1, 2).
		Width(contentWidth + 4).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(t.BorderNormal()).
		Render(content)
}

func (m *APIKeyDialogCmp) SetSize(width, height int) {
	m.width = width
	m.height = height
}

func (m APIKeyDialogCmp) BindingKeys() []key.Binding {
	return []key.Binding{
		key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "save")),
		key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "cancel")),
	}
}
