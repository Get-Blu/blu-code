package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// NordTheme implements the Theme interface with Nord colors.
type NordTheme struct {
	BaseTheme
}

// NewNordTheme creates a new instance of the Nord theme.
func NewNordTheme() *NordTheme {
	// Nord color palette
	// https://www.nordtheme.com/docs/colors-and-palettes
	nord0 := "#2e3440"
	nord1 := "#3b4252"
	nord2 := "#434c5e"
	nord3 := "#4c566a"
	nord4 := "#d8dee9"
	nord5 := "#e5e9f0"
	nord6 := "#eceff4"
	nord7 := "#8fbcbb"
	nord8 := "#88c0d0"
	nord9 := "#81a1c1"
	nord10 := "#5e81ac"
	nord11 := "#bf616a" // red
	nord12 := "#d08770" // orange
	nord13 := "#ebcb8b" // yellow
	nord14 := "#a3be8c" // green
	nord15 := "#b48ead" // purple

	theme := &NordTheme{}

	// Base colors
	theme.PrimaryColor = lipgloss.AdaptiveColor{Dark: nord8, Light: nord10}
	theme.SecondaryColor = lipgloss.AdaptiveColor{Dark: nord9, Light: nord9}
	theme.AccentColor = lipgloss.AdaptiveColor{Dark: nord7, Light: nord8}

	// Status colors
	theme.ErrorColor = lipgloss.AdaptiveColor{Dark: nord11, Light: nord11}
	theme.WarningColor = lipgloss.AdaptiveColor{Dark: nord12, Light: nord12}
	theme.SuccessColor = lipgloss.AdaptiveColor{Dark: nord14, Light: nord14}
	theme.InfoColor = lipgloss.AdaptiveColor{Dark: nord8, Light: nord10}

	// Text colors
	theme.TextColor = lipgloss.AdaptiveColor{Dark: nord4, Light: nord0}
	theme.TextMutedColor = lipgloss.AdaptiveColor{Dark: nord3, Light: nord3}
	theme.TextEmphasizedColor = lipgloss.AdaptiveColor{Dark: nord6, Light: nord1}

	// Background colors
	theme.BackgroundColor = lipgloss.AdaptiveColor{Dark: nord0, Light: nord6}
	theme.BackgroundSecondaryColor = lipgloss.AdaptiveColor{Dark: nord1, Light: nord5}
	theme.BackgroundDarkerColor = lipgloss.AdaptiveColor{Dark: "#242933", Light: "#e5e9f0"}

	// Border colors
	theme.BorderNormalColor = lipgloss.AdaptiveColor{Dark: nord2, Light: nord4}
	theme.BorderFocusedColor = lipgloss.AdaptiveColor{Dark: nord8, Light: nord10}
	theme.BorderDimColor = lipgloss.AdaptiveColor{Dark: nord1, Light: nord5}

	// Diff view colors
	theme.DiffAddedColor = lipgloss.AdaptiveColor{Dark: nord14, Light: "#4f7a38"}
	theme.DiffRemovedColor = lipgloss.AdaptiveColor{Dark: nord11, Light: "#a33e3e"}
	theme.DiffContextColor = lipgloss.AdaptiveColor{Dark: nord3, Light: nord3}
	theme.DiffHunkHeaderColor = lipgloss.AdaptiveColor{Dark: nord15, Light: nord15}
	theme.DiffHighlightAddedColor = lipgloss.AdaptiveColor{Dark: "#b5cea1", Light: "#a5d6a7"}
	theme.DiffHighlightRemovedColor = lipgloss.AdaptiveColor{Dark: "#f44747", Light: "#ef9a9a"}
	theme.DiffAddedBgColor = lipgloss.AdaptiveColor{Dark: "#3b4b3b", Light: "#e8f5e9"}
	theme.DiffRemovedBgColor = lipgloss.AdaptiveColor{Dark: "#4b3535", Light: "#ffebee"}
	theme.DiffContextBgColor = lipgloss.AdaptiveColor{Dark: nord0, Light: nord6}
	theme.DiffLineNumberColor = lipgloss.AdaptiveColor{Dark: nord3, Light: nord3}
	theme.DiffAddedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#334033", Light: "#c8e6c9"}
	theme.DiffRemovedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#403333", Light: "#ffcdd2"}

	// Markdown colors
	theme.MarkdownTextColor = theme.TextColor
	theme.MarkdownHeadingColor = lipgloss.AdaptiveColor{Dark: nord8, Light: nord10}
	theme.MarkdownLinkColor = lipgloss.AdaptiveColor{Dark: nord9, Light: nord9}
	theme.MarkdownLinkTextColor = lipgloss.AdaptiveColor{Dark: nord7, Light: nord8}
	theme.MarkdownCodeColor = lipgloss.AdaptiveColor{Dark: nord14, Light: nord14}
	theme.MarkdownBlockQuoteColor = lipgloss.AdaptiveColor{Dark: nord13, Light: nord13}
	theme.MarkdownEmphColor = lipgloss.AdaptiveColor{Dark: nord13, Light: nord13}
	theme.MarkdownStrongColor = lipgloss.AdaptiveColor{Dark: nord12, Light: nord12}
	theme.MarkdownHorizontalRuleColor = lipgloss.AdaptiveColor{Dark: nord3, Light: nord3}
	theme.MarkdownListItemColor = lipgloss.AdaptiveColor{Dark: nord15, Light: nord15}
	theme.MarkdownListEnumerationColor = lipgloss.AdaptiveColor{Dark: nord7, Light: nord7}
	theme.MarkdownImageColor = lipgloss.AdaptiveColor{Dark: nord15, Light: nord15}
	theme.MarkdownImageTextColor = lipgloss.AdaptiveColor{Dark: nord7, Light: nord7}
	theme.MarkdownCodeBlockColor = theme.TextColor

	// Syntax highlighting colors
	theme.SyntaxCommentColor = lipgloss.AdaptiveColor{Dark: nord3, Light: nord3}
	theme.SyntaxKeywordColor = lipgloss.AdaptiveColor{Dark: nord9, Light: nord9}
	theme.SyntaxFunctionColor = lipgloss.AdaptiveColor{Dark: nord8, Light: nord8}
	theme.SyntaxVariableColor = lipgloss.AdaptiveColor{Dark: nord4, Light: nord0}
	theme.SyntaxStringColor = lipgloss.AdaptiveColor{Dark: nord14, Light: "#4f7a38"}
	theme.SyntaxNumberColor = lipgloss.AdaptiveColor{Dark: nord15, Light: nord15}
	theme.SyntaxTypeColor = lipgloss.AdaptiveColor{Dark: nord7, Light: nord7}
	theme.SyntaxOperatorColor = lipgloss.AdaptiveColor{Dark: nord9, Light: nord9}
	theme.SyntaxPunctuationColor = lipgloss.AdaptiveColor{Dark: nord6, Light: nord1}

	return theme
}

func init() {
	RegisterTheme("nord", NewNordTheme())
}
