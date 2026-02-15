package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// ShadesOfGreyTheme implements the Theme interface with Monochrome/Grey colors.
type ShadesOfGreyTheme struct {
	BaseTheme
}

// NewShadesOfGreyTheme creates a new instance of the Shades of Grey theme.
func NewShadesOfGreyTheme() *ShadesOfGreyTheme {
	theme := &ShadesOfGreyTheme{}

	// Shades of Grey Colors
	const (
		sgBg      = "#1a1a1a"
		sgFg      = "#e0e0e0"
		sgText    = "#e0e0e0"
		sgMuted   = "#808080"
		sgDarker  = "#121212"
		sgLighter = "#333333"
		sgBorder  = "#404040"
	)

	// Base colors
	theme.PrimaryColor = lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"}
	theme.SecondaryColor = lipgloss.AdaptiveColor{Dark: "#cccccc", Light: "#333333"}
	theme.AccentColor = lipgloss.AdaptiveColor{Dark: "#999999", Light: "#666666"}

	// Status colors
	theme.ErrorColor = lipgloss.AdaptiveColor{Dark: "#ff5555", Light: "#aa0000"}
	theme.WarningColor = lipgloss.AdaptiveColor{Dark: "#ffff55", Light: "#aaaa00"}
	theme.SuccessColor = lipgloss.AdaptiveColor{Dark: sgFg, Light: sgBg}
	theme.InfoColor = lipgloss.AdaptiveColor{Dark: sgFg, Light: sgBg}

	// Text colors
	theme.TextColor = lipgloss.AdaptiveColor{Dark: sgText, Light: "#1a1a1a"}
	theme.TextMutedColor = lipgloss.AdaptiveColor{Dark: sgMuted, Light: "#808080"}
	theme.TextEmphasizedColor = lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"}

	// Background colors
	theme.BackgroundColor = lipgloss.AdaptiveColor{Dark: sgBg, Light: "#ffffff"}
	theme.BackgroundSecondaryColor = lipgloss.AdaptiveColor{Dark: sgLighter, Light: "#f5f5f5"}
	theme.BackgroundDarkerColor = lipgloss.AdaptiveColor{Dark: sgDarker, Light: "#e0e0e0"}

	// Border colors
	theme.BorderNormalColor = lipgloss.AdaptiveColor{Dark: sgBorder, Light: "#bdbdbd"}
	theme.BorderFocusedColor = lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"}
	theme.BorderDimColor = lipgloss.AdaptiveColor{Dark: sgLighter, Light: "#eeeeee"}

	// Diff view colors
	theme.DiffAddedColor = lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"}
	theme.DiffRemovedColor = lipgloss.AdaptiveColor{Dark: "#555555", Light: "#aaaaaa"}
	theme.DiffContextColor = lipgloss.AdaptiveColor{Dark: sgMuted, Light: "#757575"}
	theme.DiffHunkHeaderColor = lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"}
	theme.DiffHighlightAddedColor = lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"}
	theme.DiffHighlightRemovedColor = lipgloss.AdaptiveColor{Dark: "#888888", Light: "#777777"}
	theme.DiffAddedBgColor = lipgloss.AdaptiveColor{Dark: "#333333", Light: "#f0f0f0"}
	theme.DiffRemovedBgColor = lipgloss.AdaptiveColor{Dark: "#222222", Light: "#e0e0e0"}
	theme.DiffContextBgColor = lipgloss.AdaptiveColor{Dark: sgBg, Light: "#ffffff"}
	theme.DiffLineNumberColor = lipgloss.AdaptiveColor{Dark: sgMuted, Light: "#9e9e9e"}
	theme.DiffAddedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#444444", Light: "#e5e5e5"}
	theme.DiffRemovedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: sgBg, Light: "#d5d5d5"}

	// Markdown colors
	theme.MarkdownTextColor = lipgloss.AdaptiveColor{Dark: sgText, Light: "#1a1a1a"}
	theme.MarkdownHeadingColor = lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"}
	theme.MarkdownLinkColor = lipgloss.AdaptiveColor{Dark: "#cccccc", Light: "#333333"}
	theme.MarkdownLinkTextColor = lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"}
	theme.MarkdownCodeColor = lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"}
	theme.MarkdownBlockQuoteColor = lipgloss.AdaptiveColor{Dark: sgMuted, Light: "#808080"}
	theme.MarkdownEmphColor = lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"}
	theme.MarkdownStrongColor = lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"}
	theme.MarkdownHorizontalRuleColor = lipgloss.AdaptiveColor{Dark: sgBorder, Light: "#bdbdbd"}
	theme.MarkdownListItemColor = lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"}
	theme.MarkdownListEnumerationColor = lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"}
	theme.MarkdownImageColor = lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"}
	theme.MarkdownImageTextColor = lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"}
	theme.MarkdownCodeBlockColor = lipgloss.AdaptiveColor{Dark: sgText, Light: "#1a1a1a"}

	// Syntax highlighting colors
	theme.SyntaxCommentColor = lipgloss.AdaptiveColor{Dark: sgMuted, Light: "#808080"}
	theme.SyntaxKeywordColor = lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"}
	theme.SyntaxFunctionColor = lipgloss.AdaptiveColor{Dark: "#eeeeee", Light: "#111111"}
	theme.SyntaxVariableColor = lipgloss.AdaptiveColor{Dark: "#dddddd", Light: "#222222"}
	theme.SyntaxStringColor = lipgloss.AdaptiveColor{Dark: "#cccccc", Light: "#333333"}
	theme.SyntaxNumberColor = lipgloss.AdaptiveColor{Dark: "#bbbbbb", Light: "#444444"}
	theme.SyntaxTypeColor = lipgloss.AdaptiveColor{Dark: "#aaaaaa", Light: "#555555"}
	theme.SyntaxOperatorColor = lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"}
	theme.SyntaxPunctuationColor = lipgloss.AdaptiveColor{Dark: sgFg, Light: sgBg}

	return theme
}

func init() {
	RegisterTheme("shadesofgrey", NewShadesOfGreyTheme())
}
