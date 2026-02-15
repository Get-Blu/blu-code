package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// AyuDarkTheme implements the Theme interface with Ayu Dark colors.
type AyuDarkTheme struct {
	BaseTheme
}

// NewAyuDarkTheme creates a new instance of the Ayu Dark theme.
func NewAyuDarkTheme() *AyuDarkTheme {
	theme := &AyuDarkTheme{}

	// Ayu Dark Colors
	const (
		ayuBg      = "#0a0e14"
		ayuFg      = "#b3b1ad"
		ayuText    = "#e6e1cf"
		ayuKeyword = "#ff8f40"
		ayuTag     = "#39adb5"
		ayuFunc    = "#ffb454"
		ayuString  = "#c2d94c"
		ayuComment = "#5c6773"
		ayuAccent  = "#e6b450"
		ayuError   = "#ff3333"
	)

	// Base colors
	theme.PrimaryColor = lipgloss.AdaptiveColor{Dark: ayuAccent, Light: "#e6b450"}
	theme.SecondaryColor = lipgloss.AdaptiveColor{Dark: ayuKeyword, Light: "#ff8f40"}
	theme.AccentColor = lipgloss.AdaptiveColor{Dark: ayuTag, Light: "#39adb5"}

	// Status colors
	theme.ErrorColor = lipgloss.AdaptiveColor{Dark: ayuError, Light: "#d32f2f"}
	theme.WarningColor = lipgloss.AdaptiveColor{Dark: ayuAccent, Light: "#f57c00"}
	theme.SuccessColor = lipgloss.AdaptiveColor{Dark: ayuString, Light: "#388e3c"}
	theme.InfoColor = lipgloss.AdaptiveColor{Dark: ayuTag, Light: "#1976d2"}

	// Text colors
	theme.TextColor = lipgloss.AdaptiveColor{Dark: ayuText, Light: "#0a0e14"}
	theme.TextMutedColor = lipgloss.AdaptiveColor{Dark: ayuComment, Light: "#5c6773"}
	theme.TextEmphasizedColor = lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"}

	// Background colors
	theme.BackgroundColor = lipgloss.AdaptiveColor{Dark: ayuBg, Light: "#ffffff"}
	theme.BackgroundSecondaryColor = lipgloss.AdaptiveColor{Dark: "#0d1219", Light: "#f5f5f5"}
	theme.BackgroundDarkerColor = lipgloss.AdaptiveColor{Dark: "#070a0e", Light: "#e0e0e0"}

	// Border colors
	theme.BorderNormalColor = lipgloss.AdaptiveColor{Dark: "#151b23", Light: "#bdbdbd"}
	theme.BorderFocusedColor = lipgloss.AdaptiveColor{Dark: ayuAccent, Light: "#e6b450"}
	theme.BorderDimColor = lipgloss.AdaptiveColor{Dark: "#0d1219", Light: "#eeeeee"}

	// Diff view colors
	theme.DiffAddedColor = lipgloss.AdaptiveColor{Dark: ayuString, Light: "#2e7d32"}
	theme.DiffRemovedColor = lipgloss.AdaptiveColor{Dark: ayuError, Light: "#c62828"}
	theme.DiffContextColor = lipgloss.AdaptiveColor{Dark: ayuComment, Light: "#757575"}
	theme.DiffHunkHeaderColor = lipgloss.AdaptiveColor{Dark: ayuTag, Light: "#1976d2"}
	theme.DiffHighlightAddedColor = lipgloss.AdaptiveColor{Dark: "#d6f5d6", Light: "#c8e6c9"}
	theme.DiffHighlightRemovedColor = lipgloss.AdaptiveColor{Dark: "#f5d6d6", Light: "#ffcdd2"}
	theme.DiffAddedBgColor = lipgloss.AdaptiveColor{Dark: "#152015", Light: "#e8f5e9"}
	theme.DiffRemovedBgColor = lipgloss.AdaptiveColor{Dark: "#201515", Light: "#ffebee"}
	theme.DiffContextBgColor = lipgloss.AdaptiveColor{Dark: ayuBg, Light: "#ffffff"}
	theme.DiffLineNumberColor = lipgloss.AdaptiveColor{Dark: ayuComment, Light: "#9e9e9e"}
	theme.DiffAddedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#1a2a1a", Light: "#c8e6c9"}
	theme.DiffRemovedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#2a1a1a", Light: "#ffcdd2"}

	// Markdown colors
	theme.MarkdownTextColor = lipgloss.AdaptiveColor{Dark: ayuText, Light: "#0a0e14"}
	theme.MarkdownHeadingColor = lipgloss.AdaptiveColor{Dark: ayuAccent, Light: "#e6b450"}
	theme.MarkdownLinkColor = lipgloss.AdaptiveColor{Dark: ayuTag, Light: "#39adb5"}
	theme.MarkdownLinkTextColor = lipgloss.AdaptiveColor{Dark: ayuKeyword, Light: "#ff8f40"}
	theme.MarkdownCodeColor = lipgloss.AdaptiveColor{Dark: ayuString, Light: "#388e3c"}
	theme.MarkdownBlockQuoteColor = lipgloss.AdaptiveColor{Dark: ayuComment, Light: "#5c6773"}
	theme.MarkdownEmphColor = lipgloss.AdaptiveColor{Dark: ayuFunc, Light: "#f57c00"}
	theme.MarkdownStrongColor = lipgloss.AdaptiveColor{Dark: ayuFunc, Light: "#f57c00"}
	theme.MarkdownHorizontalRuleColor = lipgloss.AdaptiveColor{Dark: "#151b23", Light: "#bdbdbd"}
	theme.MarkdownListItemColor = lipgloss.AdaptiveColor{Dark: ayuAccent, Light: "#e6b450"}
	theme.MarkdownListEnumerationColor = lipgloss.AdaptiveColor{Dark: ayuAccent, Light: "#e6b450"}
	theme.MarkdownImageColor = lipgloss.AdaptiveColor{Dark: ayuKeyword, Light: "#ff8f40"}
	theme.MarkdownImageTextColor = lipgloss.AdaptiveColor{Dark: ayuAccent, Light: "#e6b450"}
	theme.MarkdownCodeBlockColor = lipgloss.AdaptiveColor{Dark: ayuText, Light: "#0a0e14"}

	// Syntax highlighting colors
	theme.SyntaxCommentColor = lipgloss.AdaptiveColor{Dark: ayuComment, Light: "#5c6773"}
	theme.SyntaxKeywordColor = lipgloss.AdaptiveColor{Dark: ayuKeyword, Light: "#ff8f40"}
	theme.SyntaxFunctionColor = lipgloss.AdaptiveColor{Dark: ayuFunc, Light: "#ffb454"}
	theme.SyntaxVariableColor = lipgloss.AdaptiveColor{Dark: ayuText, Light: "#0a0e14"}
	theme.SyntaxStringColor = lipgloss.AdaptiveColor{Dark: ayuString, Light: "#c2d94c"}
	theme.SyntaxNumberColor = lipgloss.AdaptiveColor{Dark: ayuAccent, Light: "#e6b450"}
	theme.SyntaxTypeColor = lipgloss.AdaptiveColor{Dark: ayuTag, Light: "#39adb5"}
	theme.SyntaxOperatorColor = lipgloss.AdaptiveColor{Dark: ayuKeyword, Light: "#ff8f40"}
	theme.SyntaxPunctuationColor = lipgloss.AdaptiveColor{Dark: ayuFg, Light: "#b3b1ad"}

	return theme
}

func init() {
	RegisterTheme("ayudark", NewAyuDarkTheme())
}
