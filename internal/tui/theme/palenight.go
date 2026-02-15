package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// PalenightTheme implements the Theme interface with Palenight colors.
type PalenightTheme struct {
	BaseTheme
}

// NewPalenightTheme creates a new instance of the Palenight theme.
func NewPalenightTheme() *PalenightTheme {
	theme := &PalenightTheme{}

	// Palenight Colors
	const (
		pnBg      = "#292d3e"
		pnFg      = "#a6accd"
		pnComment = "#676e95"
		pnRed     = "#f07178"
		pnGreen   = "#c3e88d"
		pnYellow  = "#ffcb6b"
		pnBlue    = "#82aaff"
		pnMagenta = "#c792ea"
		pnCyan    = "#89ddff"
		pnWhite   = "#ffffff"
		pnAccent  = "#ab47bc"
	)

	// Base colors
	theme.PrimaryColor = lipgloss.AdaptiveColor{Dark: pnBlue, Light: "#303f9f"}
	theme.SecondaryColor = lipgloss.AdaptiveColor{Dark: pnMagenta, Light: "#7b1fa2"}
	theme.AccentColor = lipgloss.AdaptiveColor{Dark: pnCyan, Light: "#0097a7"}

	// Status colors
	theme.ErrorColor = lipgloss.AdaptiveColor{Dark: pnRed, Light: "#d32f2f"}
	theme.WarningColor = lipgloss.AdaptiveColor{Dark: pnYellow, Light: "#f57c00"}
	theme.SuccessColor = lipgloss.AdaptiveColor{Dark: pnGreen, Light: "#388e3c"}
	theme.InfoColor = lipgloss.AdaptiveColor{Dark: pnCyan, Light: "#1976d2"}

	// Text colors
	theme.TextColor = lipgloss.AdaptiveColor{Dark: pnFg, Light: "#292d3e"}
	theme.TextMutedColor = lipgloss.AdaptiveColor{Dark: pnComment, Light: "#546e7a"}
	theme.TextEmphasizedColor = lipgloss.AdaptiveColor{Dark: pnWhite, Light: "#000000"}

	// Background colors
	theme.BackgroundColor = lipgloss.AdaptiveColor{Dark: pnBg, Light: "#ffffff"}
	theme.BackgroundSecondaryColor = lipgloss.AdaptiveColor{Dark: "#32374d", Light: "#f5f5f5"}
	theme.BackgroundDarkerColor = lipgloss.AdaptiveColor{Dark: "#232635", Light: "#e0e0e0"}

	// Border colors
	theme.BorderNormalColor = lipgloss.AdaptiveColor{Dark: "#444b6a", Light: "#bdbdbd"}
	theme.BorderFocusedColor = lipgloss.AdaptiveColor{Dark: pnBlue, Light: "#303f9f"}
	theme.BorderDimColor = lipgloss.AdaptiveColor{Dark: "#32374d", Light: "#eeeeee"}

	// Diff view colors
	theme.DiffAddedColor = lipgloss.AdaptiveColor{Dark: pnGreen, Light: "#2e7d32"}
	theme.DiffRemovedColor = lipgloss.AdaptiveColor{Dark: pnRed, Light: "#c62828"}
	theme.DiffContextColor = lipgloss.AdaptiveColor{Dark: pnComment, Light: "#757575"}
	theme.DiffHunkHeaderColor = lipgloss.AdaptiveColor{Dark: pnBlue, Light: "#1976d2"}
	theme.DiffHighlightAddedColor = lipgloss.AdaptiveColor{Dark: "#d6f5d6", Light: "#c8e6c9"}
	theme.DiffHighlightRemovedColor = lipgloss.AdaptiveColor{Dark: "#f5d6d6", Light: "#ffcdd2"}
	theme.DiffAddedBgColor = lipgloss.AdaptiveColor{Dark: "#2d352d", Light: "#e8f5e9"}
	theme.DiffRemovedBgColor = lipgloss.AdaptiveColor{Dark: "#352d2d", Light: "#ffebee"}
	theme.DiffContextBgColor = lipgloss.AdaptiveColor{Dark: pnBg, Light: "#ffffff"}
	theme.DiffLineNumberColor = lipgloss.AdaptiveColor{Dark: pnComment, Light: "#9e9e9e"}
	theme.DiffAddedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#343c34", Light: "#c8e6c9"}
	theme.DiffRemovedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#3c3434", Light: "#ffcdd2"}

	// Markdown colors
	theme.MarkdownTextColor = lipgloss.AdaptiveColor{Dark: pnFg, Light: "#292d3e"}
	theme.MarkdownHeadingColor = lipgloss.AdaptiveColor{Dark: pnBlue, Light: "#303f9f"}
	theme.MarkdownLinkColor = lipgloss.AdaptiveColor{Dark: pnCyan, Light: "#0097a7"}
	theme.MarkdownLinkTextColor = lipgloss.AdaptiveColor{Dark: pnMagenta, Light: "#7b1fa2"}
	theme.MarkdownCodeColor = lipgloss.AdaptiveColor{Dark: pnGreen, Light: "#388e3c"}
	theme.MarkdownBlockQuoteColor = lipgloss.AdaptiveColor{Dark: pnComment, Light: "#546e7a"}
	theme.MarkdownEmphColor = lipgloss.AdaptiveColor{Dark: pnYellow, Light: "#f57c00"}
	theme.MarkdownStrongColor = lipgloss.AdaptiveColor{Dark: pnYellow, Light: "#f57c00"}
	theme.MarkdownHorizontalRuleColor = lipgloss.AdaptiveColor{Dark: "#444b6a", Light: "#bdbdbd"}
	theme.MarkdownListItemColor = lipgloss.AdaptiveColor{Dark: pnBlue, Light: "#303f9f"}
	theme.MarkdownListEnumerationColor = lipgloss.AdaptiveColor{Dark: pnBlue, Light: "#303f9f"}
	theme.MarkdownImageColor = lipgloss.AdaptiveColor{Dark: pnMagenta, Light: "#7b1fa2"}
	theme.MarkdownImageTextColor = lipgloss.AdaptiveColor{Dark: pnBlue, Light: "#303f9f"}
	theme.MarkdownCodeBlockColor = lipgloss.AdaptiveColor{Dark: pnFg, Light: "#292d3e"}

	// Syntax highlighting colors
	theme.SyntaxCommentColor = lipgloss.AdaptiveColor{Dark: pnComment, Light: "#546e7a"}
	theme.SyntaxKeywordColor = lipgloss.AdaptiveColor{Dark: pnMagenta, Light: "#7b1fa2"}
	theme.SyntaxFunctionColor = lipgloss.AdaptiveColor{Dark: pnBlue, Light: "#303f9f"}
	theme.SyntaxVariableColor = lipgloss.AdaptiveColor{Dark: pnFg, Light: "#292d3e"}
	theme.SyntaxStringColor = lipgloss.AdaptiveColor{Dark: pnGreen, Light: "#388e3c"}
	theme.SyntaxNumberColor = lipgloss.AdaptiveColor{Dark: pnYellow, Light: "#f57c00"}
	theme.SyntaxTypeColor = lipgloss.AdaptiveColor{Dark: pnYellow, Light: "#f57c00"}
	theme.SyntaxOperatorColor = lipgloss.AdaptiveColor{Dark: pnCyan, Light: "#0097a7"}
	theme.SyntaxPunctuationColor = lipgloss.AdaptiveColor{Dark: pnCyan, Light: "#0097a7"}

	return theme
}

func init() {
	RegisterTheme("palenight", NewPalenightTheme())
}
