package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// NightOwlTheme implements the Theme interface with Night Owl colors.
type NightOwlTheme struct {
	BaseTheme
}

// NewNightOwlTheme creates a new instance of the Night Owl theme.
func NewNightOwlTheme() *NightOwlTheme {
	// Night Owl color palette
	background := "#011627"
	foreground := "#d6deeb"
	comment := "#637777"
	cyan := "#7fdbca"
	green := "#addb67"
	orange := "#f78c6c"
	pink := "#c792ea"
	purple := "#82aaff"
	red := "#ef5350"
	yellow := "#ecc48d"
	selection := "#1d3b53"
	uiBorder := "#1d3b53"

	theme := &NightOwlTheme{}

	// Base colors
	theme.PrimaryColor = lipgloss.AdaptiveColor{Dark: purple, Light: "#005cc5"}
	theme.SecondaryColor = lipgloss.AdaptiveColor{Dark: pink, Light: "#6f42c1"}
	theme.AccentColor = lipgloss.AdaptiveColor{Dark: cyan, Light: "#005cc5"}

	// Status colors
	theme.ErrorColor = lipgloss.AdaptiveColor{Dark: red, Light: "#d73a49"}
	theme.WarningColor = lipgloss.AdaptiveColor{Dark: orange, Light: "#e36209"}
	theme.SuccessColor = lipgloss.AdaptiveColor{Dark: green, Light: "#22863a"}
	theme.InfoColor = lipgloss.AdaptiveColor{Dark: cyan, Light: "#005cc5"}

	// Text colors
	theme.TextColor = lipgloss.AdaptiveColor{Dark: foreground, Light: "#24292e"}
	theme.TextMutedColor = lipgloss.AdaptiveColor{Dark: comment, Light: "#6a737d"}
	theme.TextEmphasizedColor = lipgloss.AdaptiveColor{Dark: yellow, Light: "#b08800"}

	// Background colors
	theme.BackgroundColor = lipgloss.AdaptiveColor{Dark: background, Light: "#ffffff"}
	theme.BackgroundSecondaryColor = lipgloss.AdaptiveColor{Dark: selection, Light: "#f6f8fa"}
	theme.BackgroundDarkerColor = lipgloss.AdaptiveColor{Dark: "#01111d", Light: "#f1f1f1"}

	// Border colors
	theme.BorderNormalColor = lipgloss.AdaptiveColor{Dark: uiBorder, Light: "#e1e4e8"}
	theme.BorderFocusedColor = lipgloss.AdaptiveColor{Dark: purple, Light: "#0366d6"}
	theme.BorderDimColor = lipgloss.AdaptiveColor{Dark: "#0b253a", Light: "#f1f1f1"}

	// Diff view colors
	theme.DiffAddedColor = lipgloss.AdaptiveColor{Dark: green, Light: "#22863a"}
	theme.DiffRemovedColor = lipgloss.AdaptiveColor{Dark: red, Light: "#d73a49"}
	theme.DiffContextColor = lipgloss.AdaptiveColor{Dark: comment, Light: "#6a737d"}
	theme.DiffHunkHeaderColor = lipgloss.AdaptiveColor{Dark: pink, Light: "#6f42c1"}
	theme.DiffHighlightAddedColor = lipgloss.AdaptiveColor{Dark: "#addb67", Light: "#34d058"}
	theme.DiffHighlightRemovedColor = lipgloss.AdaptiveColor{Dark: "#ef5350", Light: "#d73a49"}
	theme.DiffAddedBgColor = lipgloss.AdaptiveColor{Dark: "#1a3a1a", Light: "#e6ffed"}
	theme.DiffRemovedBgColor = lipgloss.AdaptiveColor{Dark: "#3a1a1a", Light: "#ffeef0"}
	theme.DiffContextBgColor = lipgloss.AdaptiveColor{Dark: background, Light: "#ffffff"}
	theme.DiffLineNumberColor = lipgloss.AdaptiveColor{Dark: comment, Light: "#959da5"}
	theme.DiffAddedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#152a15", Light: "#cdffd8"}
	theme.DiffRemovedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#2a1515", Light: "#ffdce0"}

	// Markdown colors
	theme.MarkdownTextColor = theme.TextColor
	theme.MarkdownHeadingColor = lipgloss.AdaptiveColor{Dark: purple, Light: "#0366d6"}
	theme.MarkdownLinkColor = lipgloss.AdaptiveColor{Dark: cyan, Light: "#005cc5"}
	theme.MarkdownLinkTextColor = lipgloss.AdaptiveColor{Dark: cyan, Light: "#005cc5"}
	theme.MarkdownCodeColor = lipgloss.AdaptiveColor{Dark: green, Light: "#22863a"}
	theme.MarkdownBlockQuoteColor = lipgloss.AdaptiveColor{Dark: yellow, Light: "#b08800"}
	theme.MarkdownEmphColor = lipgloss.AdaptiveColor{Dark: yellow, Light: "#b08800"}
	theme.MarkdownStrongColor = lipgloss.AdaptiveColor{Dark: orange, Light: "#e36209"}
	theme.MarkdownHorizontalRuleColor = lipgloss.AdaptiveColor{Dark: uiBorder, Light: "#e1e4e8"}
	theme.MarkdownListItemColor = lipgloss.AdaptiveColor{Dark: pink, Light: "#6f42c1"}
	theme.MarkdownListEnumerationColor = lipgloss.AdaptiveColor{Dark: cyan, Light: "#005cc5"}
	theme.MarkdownImageColor = lipgloss.AdaptiveColor{Dark: pink, Light: "#6f42c1"}
	theme.MarkdownImageTextColor = lipgloss.AdaptiveColor{Dark: cyan, Light: "#005cc5"}
	theme.MarkdownCodeBlockColor = theme.TextColor

	// Syntax highlighting colors
	theme.SyntaxCommentColor = lipgloss.AdaptiveColor{Dark: comment, Light: "#6a737d"}
	theme.SyntaxKeywordColor = lipgloss.AdaptiveColor{Dark: pink, Light: "#d73a49"}
	theme.SyntaxFunctionColor = lipgloss.AdaptiveColor{Dark: purple, Light: "#6f42c1"}
	theme.SyntaxVariableColor = lipgloss.AdaptiveColor{Dark: foreground, Light: "#24292e"}
	theme.SyntaxStringColor = lipgloss.AdaptiveColor{Dark: yellow, Light: "#032f62"}
	theme.SyntaxNumberColor = lipgloss.AdaptiveColor{Dark: orange, Light: "#e36209"}
	theme.SyntaxTypeColor = lipgloss.AdaptiveColor{Dark: cyan, Light: "#6f42c1"}
	theme.SyntaxOperatorColor = lipgloss.AdaptiveColor{Dark: cyan, Light: "#d73a49"}
	theme.SyntaxPunctuationColor = lipgloss.AdaptiveColor{Dark: foreground, Light: "#24292e"}

	return theme
}

func init() {
	RegisterTheme("night-owl", NewNightOwlTheme())
}
