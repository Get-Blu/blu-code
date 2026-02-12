package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// EverforestTheme implements the Theme interface with Everforest colors.
type EverforestTheme struct {
	BaseTheme
}

// NewEverforestTheme creates a new instance of the Everforest theme.
func NewEverforestTheme() *EverforestTheme {
	// Everforest color palette (Hard Dark)
	background := "#2b3339"
	foreground := "#d3c6aa"
	comment := "#859289"
	red := "#e67e80"
	orange := "#e69875"
	yellow := "#dbbc7f"
	green := "#a7c080"
	blue := "#7fbbba"
	purple := "#d699b6"
	aqua := "#83c092"
	grey := "#475258"

	theme := &EverforestTheme{}

	// Base colors
	theme.PrimaryColor = lipgloss.AdaptiveColor{Dark: green, Light: "#5b8c32"}
	theme.SecondaryColor = lipgloss.AdaptiveColor{Dark: blue, Light: "#3a94c5"}
	theme.AccentColor = lipgloss.AdaptiveColor{Dark: aqua, Light: "#35a775"}

	// Status colors
	theme.ErrorColor = lipgloss.AdaptiveColor{Dark: red, Light: "#e67e80"}
	theme.WarningColor = lipgloss.AdaptiveColor{Dark: orange, Light: "#e69875"}
	theme.SuccessColor = lipgloss.AdaptiveColor{Dark: green, Light: "#a7c080"}
	theme.InfoColor = lipgloss.AdaptiveColor{Dark: blue, Light: "#7fbbba"}

	// Text colors
	theme.TextColor = lipgloss.AdaptiveColor{Dark: foreground, Light: "#5c6a72"}
	theme.TextMutedColor = lipgloss.AdaptiveColor{Dark: comment, Light: "#939f91"}
	theme.TextEmphasizedColor = lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#2d353b"}

	// Background colors
	theme.BackgroundColor = lipgloss.AdaptiveColor{Dark: background, Light: "#fdf6e3"}
	theme.BackgroundSecondaryColor = lipgloss.AdaptiveColor{Dark: grey, Light: "#f3ead3"}
	theme.BackgroundDarkerColor = lipgloss.AdaptiveColor{Dark: "#232a2e", Light: "#eee8d5"}

	// Border colors
	theme.BorderNormalColor = lipgloss.AdaptiveColor{Dark: "#323c41", Light: "#e9dfc1"}
	theme.BorderFocusedColor = lipgloss.AdaptiveColor{Dark: green, Light: "#5b8c32"}
	theme.BorderDimColor = lipgloss.AdaptiveColor{Dark: "#3a454a", Light: "#eee8d5"}

	// Diff view colors
	theme.DiffAddedColor = lipgloss.AdaptiveColor{Dark: green, Light: "#5b8c32"}
	theme.DiffRemovedColor = lipgloss.AdaptiveColor{Dark: red, Light: "#e67e80"}
	theme.DiffContextColor = lipgloss.AdaptiveColor{Dark: comment, Light: "#939f91"}
	theme.DiffHunkHeaderColor = lipgloss.AdaptiveColor{Dark: purple, Light: "#d699b6"}
	theme.DiffHighlightAddedColor = lipgloss.AdaptiveColor{Dark: "#a7c080", Light: "#7db669"}
	theme.DiffHighlightRemovedColor = lipgloss.AdaptiveColor{Dark: "#e67e80", Light: "#e67e80"}
	theme.DiffAddedBgColor = lipgloss.AdaptiveColor{Dark: "#344034", Light: "#e8f5e9"}
	theme.DiffRemovedBgColor = lipgloss.AdaptiveColor{Dark: "#403434", Light: "#ffebee"}
	theme.DiffContextBgColor = lipgloss.AdaptiveColor{Dark: background, Light: "#fdf6e3"}
	theme.DiffLineNumberColor = lipgloss.AdaptiveColor{Dark: comment, Light: "#939f91"}
	theme.DiffAddedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#2e3b2e", Light: "#c8e6c9"}
	theme.DiffRemovedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#3b2e2e", Light: "#ffcdd2"}

	// Markdown colors
	theme.MarkdownTextColor = theme.TextColor
	theme.MarkdownHeadingColor = lipgloss.AdaptiveColor{Dark: green, Light: "#5b8c32"}
	theme.MarkdownLinkColor = lipgloss.AdaptiveColor{Dark: blue, Light: "#3a94c5"}
	theme.MarkdownLinkTextColor = lipgloss.AdaptiveColor{Dark: blue, Light: "#3a94c5"}
	theme.MarkdownCodeColor = lipgloss.AdaptiveColor{Dark: aqua, Light: "#35a775"}
	theme.MarkdownBlockQuoteColor = lipgloss.AdaptiveColor{Dark: yellow, Light: "#8f6f05"}
	theme.MarkdownEmphColor = lipgloss.AdaptiveColor{Dark: yellow, Light: "#8f6f05"}
	theme.MarkdownStrongColor = lipgloss.AdaptiveColor{Dark: orange, Light: "#9a4300"}
	theme.MarkdownHorizontalRuleColor = lipgloss.AdaptiveColor{Dark: "#323c41", Light: "#e9dfc1"}
	theme.MarkdownListItemColor = lipgloss.AdaptiveColor{Dark: purple, Light: "#d699b6"}
	theme.MarkdownListEnumerationColor = lipgloss.AdaptiveColor{Dark: blue, Light: "#3a94c5"}
	theme.MarkdownImageColor = lipgloss.AdaptiveColor{Dark: purple, Light: "#d699b6"}
	theme.MarkdownImageTextColor = lipgloss.AdaptiveColor{Dark: blue, Light: "#3a94c5"}
	theme.MarkdownCodeBlockColor = theme.TextColor

	// Syntax highlighting colors
	theme.SyntaxCommentColor = lipgloss.AdaptiveColor{Dark: comment, Light: "#939f91"}
	theme.SyntaxKeywordColor = lipgloss.AdaptiveColor{Dark: red, Light: "#e67e80"}
	theme.SyntaxFunctionColor = lipgloss.AdaptiveColor{Dark: blue, Light: "#3a94c5"}
	theme.SyntaxVariableColor = lipgloss.AdaptiveColor{Dark: foreground, Light: "#5c6a72"}
	theme.SyntaxStringColor = lipgloss.AdaptiveColor{Dark: green, Light: "#5b8c32"}
	theme.SyntaxNumberColor = lipgloss.AdaptiveColor{Dark: purple, Light: "#d699b6"}
	theme.SyntaxTypeColor = lipgloss.AdaptiveColor{Dark: yellow, Light: "#8f6f05"}
	theme.SyntaxOperatorColor = lipgloss.AdaptiveColor{Dark: aqua, Light: "#35a775"}
	theme.SyntaxPunctuationColor = lipgloss.AdaptiveColor{Dark: foreground, Light: "#5c6a72"}

	return theme
}

func init() {
	RegisterTheme("everforest", NewEverforestTheme())
}
