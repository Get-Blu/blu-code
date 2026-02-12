package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// AyuMirageTheme implements the Theme interface with Ayu Mirage colors.
type AyuMirageTheme struct {
	BaseTheme
}

// NewAyuMirageTheme creates a new instance of the Ayu Mirage theme.
func NewAyuMirageTheme() *AyuMirageTheme {
	// Ayu Mirage color palette
	background := "#212733"
	foreground := "#d9d7ce"
	comment := "#5c6773"
	accent := "#ffcc66"
	red := "#f07178"
	green := "#bae67e"
	yellow := "#ffcc66"
	orange := "#ffae57"
	cyan := "#95e6cb"
	purple := "#d4bfff"
	selection := "#343f4c"
	border := "#272d38"

	theme := &AyuMirageTheme{}

	// Base colors
	theme.PrimaryColor = lipgloss.AdaptiveColor{Dark: accent, Light: "#fa8d3e"}
	theme.SecondaryColor = lipgloss.AdaptiveColor{Dark: purple, Light: "#a626a4"}
	theme.AccentColor = lipgloss.AdaptiveColor{Dark: cyan, Light: "#0184bc"}

	// Status colors
	theme.ErrorColor = lipgloss.AdaptiveColor{Dark: red, Light: "#e45649"}
	theme.WarningColor = lipgloss.AdaptiveColor{Dark: orange, Light: "#986801"}
	theme.SuccessColor = lipgloss.AdaptiveColor{Dark: green, Light: "#50a14f"}
	theme.InfoColor = lipgloss.AdaptiveColor{Dark: cyan, Light: "#0184bc"}

	// Text colors
	theme.TextColor = lipgloss.AdaptiveColor{Dark: foreground, Light: "#383a42"}
	theme.TextMutedColor = lipgloss.AdaptiveColor{Dark: comment, Light: "#a0a1a7"}
	theme.TextEmphasizedColor = lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"}

	// Background colors
	theme.BackgroundColor = lipgloss.AdaptiveColor{Dark: background, Light: "#fafafa"}
	theme.BackgroundSecondaryColor = lipgloss.AdaptiveColor{Dark: selection, Light: "#f0f0f1"}
	theme.BackgroundDarkerColor = lipgloss.AdaptiveColor{Dark: "#191e2a", Light: "#eaeaeaff"}

	// Border colors
	theme.BorderNormalColor = lipgloss.AdaptiveColor{Dark: border, Light: "#dbdbdc"}
	theme.BorderFocusedColor = lipgloss.AdaptiveColor{Dark: accent, Light: "#fa8d3e"}
	theme.BorderDimColor = lipgloss.AdaptiveColor{Dark: "#2d333f", Light: "#eaeaeaff"}

	// Diff view colors
	theme.DiffAddedColor = lipgloss.AdaptiveColor{Dark: green, Light: "#50a14f"}
	theme.DiffRemovedColor = lipgloss.AdaptiveColor{Dark: red, Light: "#e45649"}
	theme.DiffContextColor = lipgloss.AdaptiveColor{Dark: comment, Light: "#a0a1a7"}
	theme.DiffHunkHeaderColor = lipgloss.AdaptiveColor{Dark: purple, Light: "#a626a4"}
	theme.DiffHighlightAddedColor = lipgloss.AdaptiveColor{Dark: "#bae67e", Light: "#4caf50"}
	theme.DiffHighlightRemovedColor = lipgloss.AdaptiveColor{Dark: "#f07178", Light: "#ff5252"}
	theme.DiffAddedBgColor = lipgloss.AdaptiveColor{Dark: "#2a3b2a", Light: "#e8f5e9"}
	theme.DiffRemovedBgColor = lipgloss.AdaptiveColor{Dark: "#3b2a2a", Light: "#ffebee"}
	theme.DiffContextBgColor = lipgloss.AdaptiveColor{Dark: background, Light: "#fafafa"}
	theme.DiffLineNumberColor = lipgloss.AdaptiveColor{Dark: comment, Light: "#a0a1a7"}
	theme.DiffAddedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#243224", Light: "#c8e6c9"}
	theme.DiffRemovedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#322424", Light: "#ffcdd2"}

	// Markdown colors
	theme.MarkdownTextColor = theme.TextColor
	theme.MarkdownHeadingColor = lipgloss.AdaptiveColor{Dark: accent, Light: "#fa8d3e"}
	theme.MarkdownLinkColor = lipgloss.AdaptiveColor{Dark: cyan, Light: "#4078f2"}
	theme.MarkdownLinkTextColor = lipgloss.AdaptiveColor{Dark: cyan, Light: "#4078f2"}
	theme.MarkdownCodeColor = lipgloss.AdaptiveColor{Dark: green, Light: "#50a14f"}
	theme.MarkdownBlockQuoteColor = lipgloss.AdaptiveColor{Dark: orange, Light: "#986801"}
	theme.MarkdownEmphColor = lipgloss.AdaptiveColor{Dark: yellow, Light: "#986801"}
	theme.MarkdownStrongColor = lipgloss.AdaptiveColor{Dark: orange, Light: "#986801"}
	theme.MarkdownHorizontalRuleColor = lipgloss.AdaptiveColor{Dark: border, Light: "#dbdbdc"}
	theme.MarkdownListItemColor = lipgloss.AdaptiveColor{Dark: purple, Light: "#a626a4"}
	theme.MarkdownListEnumerationColor = lipgloss.AdaptiveColor{Dark: cyan, Light: "#4078f2"}
	theme.MarkdownImageColor = lipgloss.AdaptiveColor{Dark: purple, Light: "#a626a4"}
	theme.MarkdownImageTextColor = lipgloss.AdaptiveColor{Dark: cyan, Light: "#4078f2"}
	theme.MarkdownCodeBlockColor = theme.TextColor

	// Syntax highlighting colors
	theme.SyntaxCommentColor = lipgloss.AdaptiveColor{Dark: comment, Light: "#a0a1a7"}
	theme.SyntaxKeywordColor = lipgloss.AdaptiveColor{Dark: orange, Light: "#e45649"}
	theme.SyntaxFunctionColor = lipgloss.AdaptiveColor{Dark: accent, Light: "#4078f2"}
	theme.SyntaxVariableColor = lipgloss.AdaptiveColor{Dark: foreground, Light: "#383a42"}
	theme.SyntaxStringColor = lipgloss.AdaptiveColor{Dark: green, Light: "#50a14f"}
	theme.SyntaxNumberColor = lipgloss.AdaptiveColor{Dark: purple, Light: "#d75f00"}
	theme.SyntaxTypeColor = lipgloss.AdaptiveColor{Dark: purple, Light: "#c18401"}
	theme.SyntaxOperatorColor = lipgloss.AdaptiveColor{Dark: accent, Light: "#0184bc"}
	theme.SyntaxPunctuationColor = lipgloss.AdaptiveColor{Dark: foreground, Light: "#383a42"}

	return theme
}

func init() {
	RegisterTheme("ayu-mirage", NewAyuMirageTheme())
}
