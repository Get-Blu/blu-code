package page

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/Get-Blu/blu-code/internal/app"
	"github.com/Get-Blu/blu-code/internal/tui/components/dialog"
	"github.com/Get-Blu/blu-code/internal/tui/theme"
	"github.com/Get-Blu/blu-code/internal/tui/util"
)

var EditorPage PageID = "editor"

type editorPage struct {
	app      *app.App
	textarea textarea.Model
	width    int
	height   int
}

func NewEditorPage(app *app.App) tea.Model {
	p := &editorPage{
		app: app,
	}
	p.textarea = p.createTextArea(nil)
	return p
}

func (p *editorPage) Init() tea.Cmd {
	return textarea.Blink
}

func (p *editorPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case dialog.ThemeChangedMsg:
		p.textarea = p.createTextArea(&p.textarea)
		return p, nil
	case tea.WindowSizeMsg:
		p.width = msg.Width
		p.height = msg.Height
		p.textarea.SetWidth(msg.Width)
		p.textarea.SetHeight(msg.Height - 3) // Leave room for header and status bar
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return p, func() tea.Msg {
				return PageChangeMsg{ID: ChatPage}
			}
		case "ctrl+s":
			return p, util.ReportInfo("File saved! (Simulated)")
		}
	}

	p.textarea, cmd = p.textarea.Update(msg)
	return p, cmd
}

func (p *editorPage) createTextArea(existing *textarea.Model) textarea.Model {
	t := theme.CurrentTheme()
	ta := textarea.New()
	
	style := lipgloss.NewStyle().Foreground(t.Text()).Background(t.Background())
	cursorLineStyle := lipgloss.NewStyle().Background(t.BackgroundSecondary())
	lineNumberStyle := lipgloss.NewStyle().Foreground(t.TextMuted())
	
	ta.FocusedStyle.Base = style
	ta.FocusedStyle.CursorLine = cursorLineStyle
	ta.FocusedStyle.LineNumber = lineNumberStyle
	
	ta.BlurredStyle.Base = style
	ta.BlurredStyle.CursorLine = cursorLineStyle
	ta.BlurredStyle.LineNumber = lineNumberStyle
	
	ta.Placeholder = "Write code here..."
	ta.ShowLineNumbers = true
	ta.CharLimit = 0
	
	if existing != nil {
		ta.SetValue(existing.Value())
	}
	
	ta.Focus()
	ta.SetWidth(p.width)
	ta.SetHeight(p.height - 3)
	return ta
}

func (p *editorPage) View() string {
	t := theme.CurrentTheme()
	
	// Headers
	headerStyle := lipgloss.NewStyle().
		Background(t.BackgroundSecondary()).
		Foreground(t.Primary()).
		Padding(0, 1).
		Bold(true)
	
	header := headerStyle.Render(" \ufb44 Editor ")
	headerPadding := lipgloss.NewStyle().
		Background(t.BackgroundSecondary()).
		Width(p.width - lipgloss.Width(header)).
		Render("")
	
	fullHeader := lipgloss.JoinHorizontal(lipgloss.Top, header, headerPadding)

	// Sidebar (Mock)
	sidebarWidth := 20
	sidebarStyle := lipgloss.NewStyle().
		Width(sidebarWidth).
		Height(p.height - 3).
		Background(t.BackgroundDarker()).
		Foreground(t.TextMuted()).
		Padding(1, 1).
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(t.BorderNormal())

	files := []string{
		"\uf015 explorer",
		"",
		"\uf471 README.md",
		"\ue73c main.go",
		"\uf121 go.mod",
		"\ue7a8 config.yaml",
	}
	
	sidebarContent := ""
	for _, f := range files {
		sidebarContent += f + "\n"
	}
	sidebar := sidebarStyle.Render(sidebarContent)

	// Editor area
	p.textarea.SetWidth(p.width - sidebarWidth - 1)
	editorView := p.textarea.View()

	middle := lipgloss.JoinHorizontal(lipgloss.Top, sidebar, editorView)

	// Status bar
	statusStyle := lipgloss.NewStyle().
		Background(t.Primary()).
		Foreground(t.Background()).
		Padding(0, 1)
	
	line := p.textarea.Line() + 1
	col := 0 // Defaulting to 0 since Column() is undefined in this version
	
	statusInfo := statusStyle.Render(fmt.Sprintf("LN %d, COL %d", line, col))
	statusPadding := lipgloss.NewStyle().
		Background(t.BackgroundSecondary()).
		Foreground(t.TextMuted()).
		Width(p.width - lipgloss.Width(statusInfo)).
		Render(" ESC: Back | CTRL+S: Save")
	
	fullStatus := lipgloss.JoinHorizontal(lipgloss.Top, statusPadding, statusInfo)

	return lipgloss.JoinVertical(lipgloss.Left,
		fullHeader,
		middle,
		fullStatus,
	)
}

func (p *editorPage) SetSize(width, height int) tea.Cmd {
	p.width = width
	p.height = height
	sidebarWidth := 20
	p.textarea.SetWidth(width - sidebarWidth - 1)
	p.textarea.SetHeight(height - 3)
	return nil
}

func (p *editorPage) BindingKeys() []key.Binding {
	return []key.Binding{}
}
