package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// AuraTheme implements the Theme interface with Aura colors.
type AuraTheme struct {
	BaseTheme
}

// NewAuraTheme creates a new instance of the Aura theme.
func NewAuraTheme() *AuraTheme {
	theme := &AuraTheme{}

	// Aura Colors
	const (
		auraBg      = "#151515"
		auraFg      = "#edecee"
		auraBlack   = "#110f0f"
		auraRed     = "#ff6767"
		auraGreen   = "#61ffca"
		auraYellow  = "#ffca85"
		auraBlue    = "#82e2ff"
		auraMagenta = "#a277ff"
		auraCyan    = "#61ffca"
		auraWhite   = "#edecee"
		auraOrange  = "#ffca85"
		auraPurple  = "#a277ff"
		auraComment = "#6d6d6d"
	)

	// Base colors
	theme.PrimaryColor = lipgloss.AdaptiveColor{Dark: auraPurple, Light: "#6200ee"}
	theme.SecondaryColor = lipgloss.AdaptiveColor{Dark: auraMagenta, Light: "#9c27b0"}
	theme.AccentColor = lipgloss.AdaptiveColor{Dark: auraGreen, Light: "#00c853"}

	// Status colors
	theme.ErrorColor = lipgloss.AdaptiveColor{Dark: auraRed, Light: "#d50000"}
	theme.WarningColor = lipgloss.AdaptiveColor{Dark: auraYellow, Light: "#ffab00"}
	theme.SuccessColor = lipgloss.AdaptiveColor{Dark: auraGreen, Light: "#00c853"}
	theme.InfoColor = lipgloss.AdaptiveColor{Dark: auraBlue, Light: "#00b0ff"}

	// Text colors
	theme.TextColor = lipgloss.AdaptiveColor{Dark: auraFg, Light: "#151515"}
	theme.TextMutedColor = lipgloss.AdaptiveColor{Dark: auraComment, Light: "#757575"}
	theme.TextEmphasizedColor = lipgloss.AdaptiveColor{Dark: auraCyan, Light: "#00bfa5"}

	// Background colors
	theme.BackgroundColor = lipgloss.AdaptiveColor{Dark: auraBg, Light: "#ffffff"}
	theme.BackgroundSecondaryColor = lipgloss.AdaptiveColor{Dark: "#1c1c1c", Light: "#f5f5f5"}
	theme.BackgroundDarkerColor = lipgloss.AdaptiveColor{Dark: "#101010", Light: "#e0e0e0"}

	// Border colors
	theme.BorderNormalColor = lipgloss.AdaptiveColor{Dark: "#333333", Light: "#bdbdbd"}
	theme.BorderFocusedColor = lipgloss.AdaptiveColor{Dark: auraPurple, Light: "#6200ee"}
	theme.BorderDimColor = lipgloss.AdaptiveColor{Dark: "#222222", Light: "#eeeeee"}

	// Diff view colors
	theme.DiffAddedColor = lipgloss.AdaptiveColor{Dark: auraGreen, Light: "#2e7d32"}
	theme.DiffRemovedColor = lipgloss.AdaptiveColor{Dark: auraRed, Light: "#c62828"}
	theme.DiffContextColor = lipgloss.AdaptiveColor{Dark: auraComment, Light: "#757575"}
	theme.DiffHunkHeaderColor = lipgloss.AdaptiveColor{Dark: auraBlue, Light: "#1976d2"}
	theme.DiffHighlightAddedColor = lipgloss.AdaptiveColor{Dark: "#b2ffd9", Light: "#c8e6c9"}
	theme.DiffHighlightRemovedColor = lipgloss.AdaptiveColor{Dark: "#ffb2b2", Light: "#ffcdd2"}
	theme.DiffAddedBgColor = lipgloss.AdaptiveColor{Dark: "#1e2a24", Light: "#e8f5e9"}
	theme.DiffRemovedBgColor = lipgloss.AdaptiveColor{Dark: "#2a1e1e", Light: "#ffebee"}
	theme.DiffContextBgColor = lipgloss.AdaptiveColor{Dark: auraBg, Light: "#ffffff"}
	theme.DiffLineNumberColor = lipgloss.AdaptiveColor{Dark: auraComment, Light: "#9e9e9e"}
	theme.DiffAddedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#24332b", Light: "#c8e6c9"}
	theme.DiffRemovedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#332424", Light: "#ffcdd2"}

	// Markdown colors
	theme.MarkdownTextColor = lipgloss.AdaptiveColor{Dark: auraFg, Light: "#151515"}
	theme.MarkdownHeadingColor = lipgloss.AdaptiveColor{Dark: auraPurple, Light: "#6200ee"}
	theme.MarkdownLinkColor = lipgloss.AdaptiveColor{Dark: auraBlue, Light: "#00b0ff"}
	theme.MarkdownLinkTextColor = lipgloss.AdaptiveColor{Dark: auraMagenta, Light: "#9c27b0"}
	theme.MarkdownCodeColor = lipgloss.AdaptiveColor{Dark: auraGreen, Light: "#00c853"}
	theme.MarkdownBlockQuoteColor = lipgloss.AdaptiveColor{Dark: auraComment, Light: "#757575"}
	theme.MarkdownEmphColor = lipgloss.AdaptiveColor{Dark: auraYellow, Light: "#ffab00"}
	theme.MarkdownStrongColor = lipgloss.AdaptiveColor{Dark: auraOrange, Light: "#ffab00"}
	theme.MarkdownHorizontalRuleColor = lipgloss.AdaptiveColor{Dark: "#333333", Light: "#bdbdbd"}
	theme.MarkdownListItemColor = lipgloss.AdaptiveColor{Dark: auraPurple, Light: "#6200ee"}
	theme.MarkdownListEnumerationColor = lipgloss.AdaptiveColor{Dark: auraPurple, Light: "#6200ee"}
	theme.MarkdownImageColor = lipgloss.AdaptiveColor{Dark: auraMagenta, Light: "#9c27b0"}
	theme.MarkdownImageTextColor = lipgloss.AdaptiveColor{Dark: auraPurple, Light: "#6200ee"}
	theme.MarkdownCodeBlockColor = lipgloss.AdaptiveColor{Dark: auraFg, Light: "#151515"}

	// Syntax highlighting colors
	theme.SyntaxCommentColor = lipgloss.AdaptiveColor{Dark: auraComment, Light: "#757575"}
	theme.SyntaxKeywordColor = lipgloss.AdaptiveColor{Dark: auraMagenta, Light: "#9c27b0"}
	theme.SyntaxFunctionColor = lipgloss.AdaptiveColor{Dark: auraPurple, Light: "#6200ee"}
	theme.SyntaxVariableColor = lipgloss.AdaptiveColor{Dark: auraFg, Light: "#151515"}
	theme.SyntaxStringColor = lipgloss.AdaptiveColor{Dark: auraGreen, Light: "#00c853"}
	theme.SyntaxNumberColor = lipgloss.AdaptiveColor{Dark: auraOrange, Light: "#ffab00"}
	theme.SyntaxTypeColor = lipgloss.AdaptiveColor{Dark: auraYellow, Light: "#ffab00"}
	theme.SyntaxOperatorColor = lipgloss.AdaptiveColor{Dark: auraCyan, Light: "#00bfa5"}
	theme.SyntaxPunctuationColor = lipgloss.AdaptiveColor{Dark: auraFg, Light: "#151515"}

	return theme
}

func init() {
	RegisterTheme("aura", NewAuraTheme())
}
