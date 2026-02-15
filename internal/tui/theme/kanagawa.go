package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// KanagawaTheme implements the Theme interface with Kanagawa colors.
type KanagawaTheme struct {
	BaseTheme
}

// NewKanagawaTheme creates a new instance of the Kanagawa theme.
func NewKanagawaTheme() *KanagawaTheme {
	theme := &KanagawaTheme{}

	// Kanagawa Colors
	const (
		knBg      = "#1f1f28"
		knFg      = "#dcd7ba"
		knText    = "#dcd7ba"
		knMuted   = "#727169"
		knRed     = "#c34043"
		knGreen   = "#76946a"
		knYellow  = "#c0a36e"
		knBlue    = "#7e9cd8"
		knMagenta = "#957fb8"
		knCyan    = "#6a9589"
		knOrange  = "#ffa066"
		knBorder  = "#16161d"
	)

	// Base colors
	theme.PrimaryColor = lipgloss.AdaptiveColor{Dark: knBlue, Light: "#43649b"}
	theme.SecondaryColor = lipgloss.AdaptiveColor{Dark: knMagenta, Light: "#6b4e96"}
	theme.AccentColor = lipgloss.AdaptiveColor{Dark: knOrange, Light: "#b3643b"}

	// Status colors
	theme.ErrorColor = lipgloss.AdaptiveColor{Dark: knRed, Light: "#8a2f31"}
	theme.WarningColor = lipgloss.AdaptiveColor{Dark: knYellow, Light: "#7e622e"}
	theme.SuccessColor = lipgloss.AdaptiveColor{Dark: knGreen, Light: "#4a5d43"}
	theme.InfoColor = lipgloss.AdaptiveColor{Dark: knCyan, Light: "#3c5049"}

	// Text colors
	theme.TextColor = lipgloss.AdaptiveColor{Dark: knText, Light: "#1f1f28"}
	theme.TextMutedColor = lipgloss.AdaptiveColor{Dark: knMuted, Light: "#727169"}
	theme.TextEmphasizedColor = lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"}

	// Background colors
	theme.BackgroundColor = lipgloss.AdaptiveColor{Dark: knBg, Light: "#ffffff"}
	theme.BackgroundSecondaryColor = lipgloss.AdaptiveColor{Dark: "#2a2a37", Light: "#f5f5f5"}
	theme.BackgroundDarkerColor = lipgloss.AdaptiveColor{Dark: "#16161d", Light: "#e0e0e0"}

	// Border colors
	theme.BorderNormalColor = lipgloss.AdaptiveColor{Dark: knBorder, Light: "#bdbdbd"}
	theme.BorderFocusedColor = lipgloss.AdaptiveColor{Dark: knBlue, Light: "#43649b"}
	theme.BorderDimColor = lipgloss.AdaptiveColor{Dark: "#2d2d2d", Light: "#eeeeee"}

	// Diff view colors
	theme.DiffAddedColor = lipgloss.AdaptiveColor{Dark: knGreen, Light: "#2e7d32"}
	theme.DiffRemovedColor = lipgloss.AdaptiveColor{Dark: knRed, Light: "#c62828"}
	theme.DiffContextColor = lipgloss.AdaptiveColor{Dark: knMuted, Light: "#757575"}
	theme.DiffHunkHeaderColor = lipgloss.AdaptiveColor{Dark: knBlue, Light: "#1976d2"}
	theme.DiffHighlightAddedColor = lipgloss.AdaptiveColor{Dark: "#d6f5d6", Light: "#c8e6c9"}
	theme.DiffHighlightRemovedColor = lipgloss.AdaptiveColor{Dark: "#f5d6d6", Light: "#ffcdd2"}
	theme.DiffAddedBgColor = lipgloss.AdaptiveColor{Dark: "#252e25", Light: "#e8f5e9"}
	theme.DiffRemovedBgColor = lipgloss.AdaptiveColor{Dark: "#2e2525", Light: "#ffebee"}
	theme.DiffContextBgColor = lipgloss.AdaptiveColor{Dark: knBg, Light: "#ffffff"}
	theme.DiffLineNumberColor = lipgloss.AdaptiveColor{Dark: knMuted, Light: "#9e9e9e"}
	theme.DiffAddedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#2a332a", Light: "#c8e6c9"}
	theme.DiffRemovedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#332a2a", Light: "#ffcdd2"}

	// Markdown colors
	theme.MarkdownTextColor = lipgloss.AdaptiveColor{Dark: knText, Light: "#1f1f28"}
	theme.MarkdownHeadingColor = lipgloss.AdaptiveColor{Dark: knBlue, Light: "#43649b"}
	theme.MarkdownLinkColor = lipgloss.AdaptiveColor{Dark: knCyan, Light: "#3c5049"}
	theme.MarkdownLinkTextColor = lipgloss.AdaptiveColor{Dark: knMagenta, Light: "#6b4e96"}
	theme.MarkdownCodeColor = lipgloss.AdaptiveColor{Dark: knGreen, Light: "#4a5d43"}
	theme.MarkdownBlockQuoteColor = lipgloss.AdaptiveColor{Dark: knMuted, Light: "#727169"}
	theme.MarkdownEmphColor = lipgloss.AdaptiveColor{Dark: knYellow, Light: "#7e622e"}
	theme.MarkdownStrongColor = lipgloss.AdaptiveColor{Dark: knOrange, Light: "#b3643b"}
	theme.MarkdownHorizontalRuleColor = lipgloss.AdaptiveColor{Dark: knBorder, Light: "#bdbdbd"}
	theme.MarkdownListItemColor = lipgloss.AdaptiveColor{Dark: knBlue, Light: "#43649b"}
	theme.MarkdownListEnumerationColor = lipgloss.AdaptiveColor{Dark: knBlue, Light: "#43649b"}
	theme.MarkdownImageColor = lipgloss.AdaptiveColor{Dark: knMagenta, Light: "#6b4e96"}
	theme.MarkdownImageTextColor = lipgloss.AdaptiveColor{Dark: knBlue, Light: "#43649b"}
	theme.MarkdownCodeBlockColor = lipgloss.AdaptiveColor{Dark: knText, Light: "#1f1f28"}

	// Syntax highlighting colors
	theme.SyntaxCommentColor = lipgloss.AdaptiveColor{Dark: knMuted, Light: "#727169"}
	theme.SyntaxKeywordColor = lipgloss.AdaptiveColor{Dark: knMagenta, Light: "#6b4e96"}
	theme.SyntaxFunctionColor = lipgloss.AdaptiveColor{Dark: knBlue, Light: "#43649b"}
	theme.SyntaxVariableColor = lipgloss.AdaptiveColor{Dark: knText, Light: "#1f1f28"}
	theme.SyntaxStringColor = lipgloss.AdaptiveColor{Dark: knGreen, Light: "#4a5d43"}
	theme.SyntaxNumberColor = lipgloss.AdaptiveColor{Dark: knOrange, Light: "#b3643b"}
	theme.SyntaxTypeColor = lipgloss.AdaptiveColor{Dark: knYellow, Light: "#7e622e"}
	theme.SyntaxOperatorColor = lipgloss.AdaptiveColor{Dark: knCyan, Light: "#3c5049"}
	theme.SyntaxPunctuationColor = lipgloss.AdaptiveColor{Dark: knFg, Light: "#dcd7ba"}

	return theme
}

func init() {
	RegisterTheme("kanagawa", NewKanagawaTheme())
}
