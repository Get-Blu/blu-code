package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// CyberpunkTheme implements the Theme interface with Cyberpunk colors.
type CyberpunkTheme struct {
	BaseTheme
}

// NewCyberpunkTheme creates a new instance of the Cyberpunk theme.
func NewCyberpunkTheme() *CyberpunkTheme {
	theme := &CyberpunkTheme{}

	// Cyberpunk Colors
	const (
		cpBg      = "#000b1e"
		cpFg      = "#00fff9"
		cpYellow  = "#fdf500"
		cpPink    = "#ff00a0"
		cpBlue    = "#00c2ff"
		cpGreen   = "#00ff41"
		cpPurple  = "#cc00ff"
		cpComment = "#3b3b3b"
		cpWhite   = "#ffffff"
	)

	// Base colors
	theme.PrimaryColor = lipgloss.AdaptiveColor{Dark: cpPink, Light: "#e91e63"}
	theme.SecondaryColor = lipgloss.AdaptiveColor{Dark: cpPurple, Light: "#9c27b0"}
	theme.AccentColor = lipgloss.AdaptiveColor{Dark: cpYellow, Light: "#fbc02d"}

	// Status colors
	theme.ErrorColor = lipgloss.AdaptiveColor{Dark: cpPink, Light: "#d32f2f"}
	theme.WarningColor = lipgloss.AdaptiveColor{Dark: cpYellow, Light: "#f57c00"}
	theme.SuccessColor = lipgloss.AdaptiveColor{Dark: cpGreen, Light: "#388e3c"}
	theme.InfoColor = lipgloss.AdaptiveColor{Dark: cpBlue, Light: "#1976d2"}

	// Text colors
	theme.TextColor = lipgloss.AdaptiveColor{Dark: cpFg, Light: "#000b1e"}
	theme.TextMutedColor = lipgloss.AdaptiveColor{Dark: cpComment, Light: "#5d5d5d"}
	theme.TextEmphasizedColor = lipgloss.AdaptiveColor{Dark: cpWhite, Light: "#000b1e"}

	// Background colors
	theme.BackgroundColor = lipgloss.AdaptiveColor{Dark: cpBg, Light: "#ffffff"}
	theme.BackgroundSecondaryColor = lipgloss.AdaptiveColor{Dark: "#001235", Light: "#f5f5f5"}
	theme.BackgroundDarkerColor = lipgloss.AdaptiveColor{Dark: "#00050d", Light: "#e0e0e0"}

	// Border colors
	theme.BorderNormalColor = lipgloss.AdaptiveColor{Dark: "#003a66", Light: "#bdbdbd"}
	theme.BorderFocusedColor = lipgloss.AdaptiveColor{Dark: cpPink, Light: "#e91e63"}
	theme.BorderDimColor = lipgloss.AdaptiveColor{Dark: "#001a33", Light: "#eeeeee"}

	// Diff view colors
	theme.DiffAddedColor = lipgloss.AdaptiveColor{Dark: cpGreen, Light: "#2e7d32"}
	theme.DiffRemovedColor = lipgloss.AdaptiveColor{Dark: cpPink, Light: "#c62828"}
	theme.DiffContextColor = lipgloss.AdaptiveColor{Dark: cpComment, Light: "#757575"}
	theme.DiffHunkHeaderColor = lipgloss.AdaptiveColor{Dark: cpBlue, Light: "#1976d2"}
	theme.DiffHighlightAddedColor = lipgloss.AdaptiveColor{Dark: "#d6f5d6", Light: "#c8e6c9"}
	theme.DiffHighlightRemovedColor = lipgloss.AdaptiveColor{Dark: "#f5d6d6", Light: "#ffcdd2"}
	theme.DiffAddedBgColor = lipgloss.AdaptiveColor{Dark: "#0a2a0a", Light: "#e8f5e9"}
	theme.DiffRemovedBgColor = lipgloss.AdaptiveColor{Dark: "#2a0a0a", Light: "#ffebee"}
	theme.DiffContextBgColor = lipgloss.AdaptiveColor{Dark: cpBg, Light: "#ffffff"}
	theme.DiffLineNumberColor = lipgloss.AdaptiveColor{Dark: cpComment, Light: "#9e9e9e"}
	theme.DiffAddedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#0f330f", Light: "#c8e6c9"}
	theme.DiffRemovedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#330f0f", Light: "#ffcdd2"}

	// Markdown colors
	theme.MarkdownTextColor = lipgloss.AdaptiveColor{Dark: cpFg, Light: "#000b1e"}
	theme.MarkdownHeadingColor = lipgloss.AdaptiveColor{Dark: cpPink, Light: "#e91e63"}
	theme.MarkdownLinkColor = lipgloss.AdaptiveColor{Dark: cpBlue, Light: "#00b0ff"}
	theme.MarkdownLinkTextColor = lipgloss.AdaptiveColor{Dark: cpYellow, Light: "#fbc02d"}
	theme.MarkdownCodeColor = lipgloss.AdaptiveColor{Dark: cpGreen, Light: "#388e3c"}
	theme.MarkdownBlockQuoteColor = lipgloss.AdaptiveColor{Dark: cpComment, Light: "#5d5d5d"}
	theme.MarkdownEmphColor = lipgloss.AdaptiveColor{Dark: cpYellow, Light: "#fbc02d"}
	theme.MarkdownStrongColor = lipgloss.AdaptiveColor{Dark: cpYellow, Light: "#fbc02d"}
	theme.MarkdownHorizontalRuleColor = lipgloss.AdaptiveColor{Dark: "#003a66", Light: "#bdbdbd"}
	theme.MarkdownListItemColor = lipgloss.AdaptiveColor{Dark: cpPink, Light: "#e91e63"}
	theme.MarkdownListEnumerationColor = lipgloss.AdaptiveColor{Dark: cpPink, Light: "#e91e63"}
	theme.MarkdownImageColor = lipgloss.AdaptiveColor{Dark: cpPurple, Light: "#9c27b0"}
	theme.MarkdownImageTextColor = lipgloss.AdaptiveColor{Dark: cpPink, Light: "#e91e63"}
	theme.MarkdownCodeBlockColor = lipgloss.AdaptiveColor{Dark: cpFg, Light: "#000b1e"}

	// Syntax highlighting colors
	theme.SyntaxCommentColor = lipgloss.AdaptiveColor{Dark: cpComment, Light: "#5d5d5d"}
	theme.SyntaxKeywordColor = lipgloss.AdaptiveColor{Dark: cpPink, Light: "#e91e63"}
	theme.SyntaxFunctionColor = lipgloss.AdaptiveColor{Dark: cpBlue, Light: "#1976d2"}
	theme.SyntaxVariableColor = lipgloss.AdaptiveColor{Dark: cpFg, Light: "#000b1e"}
	theme.SyntaxStringColor = lipgloss.AdaptiveColor{Dark: cpYellow, Light: "#fbc02d"}
	theme.SyntaxNumberColor = lipgloss.AdaptiveColor{Dark: cpPurple, Light: "#9c27b0"}
	theme.SyntaxTypeColor = lipgloss.AdaptiveColor{Dark: cpYellow, Light: "#fbc02d"}
	theme.SyntaxOperatorColor = lipgloss.AdaptiveColor{Dark: cpPurple, Light: "#9c27b0"}
	theme.SyntaxPunctuationColor = lipgloss.AdaptiveColor{Dark: cpFg, Light: "#000b1e"}

	return theme
}

func init() {
	RegisterTheme("cyberpunk", NewCyberpunkTheme())
}
