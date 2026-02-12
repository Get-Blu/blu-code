package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// SolarizedLightTheme implements the Theme interface with Solarized Light colors.
type SolarizedLightTheme struct {
	BaseTheme
}

// NewSolarizedLightTheme creates a new instance of the Solarized Light theme.
func NewSolarizedLightTheme() *SolarizedLightTheme {
	// Solarized Light color palette
	// https://ethanschoonover.com/solarized/
	base03 := "#002b36"
	base02 := "#073642"
	base01 := "#586e75"
	base00 := "#657b83"
	base0 := "#839496"
	base1 := "#93a1a1"
	base2 := "#eee8d5"
	base3 := "#fdf6e3"
	yellow := "#b58900"
	orange := "#cb4b16"
	red := "#dc322f"
	magenta := "#d33682"
	violet := "#6c71c4"
	blue := "#268bd2"
	cyan := "#2aa198"
	green := "#859900"

	theme := &SolarizedLightTheme{}

	// Base colors
	theme.PrimaryColor = lipgloss.AdaptiveColor{Light: blue, Dark: blue}
	theme.SecondaryColor = lipgloss.AdaptiveColor{Light: magenta, Dark: magenta}
	theme.AccentColor = lipgloss.AdaptiveColor{Light: cyan, Dark: cyan}

	// Status colors
	theme.ErrorColor = lipgloss.AdaptiveColor{Light: red, Dark: red}
	theme.WarningColor = lipgloss.AdaptiveColor{Light: orange, Dark: orange}
	theme.SuccessColor = lipgloss.AdaptiveColor{Light: green, Dark: green}
	theme.InfoColor = lipgloss.AdaptiveColor{Light: cyan, Dark: cyan}

	// Text colors
	theme.TextColor = lipgloss.AdaptiveColor{Light: base00, Dark: base0}
	theme.TextMutedColor = lipgloss.AdaptiveColor{Light: base1, Dark: base01}
	theme.TextEmphasizedColor = lipgloss.AdaptiveColor{Light: base01, Dark: base1}

	// Background colors
	theme.BackgroundColor = lipgloss.AdaptiveColor{Light: base3, Dark: base03}
	theme.BackgroundSecondaryColor = lipgloss.AdaptiveColor{Light: base2, Dark: base02}
	theme.BackgroundDarkerColor = lipgloss.AdaptiveColor{Light: "#eee8d5", Dark: "#001e26"}

	// Border colors
	theme.BorderNormalColor = lipgloss.AdaptiveColor{Light: base2, Dark: base02}
	theme.BorderFocusedColor = lipgloss.AdaptiveColor{Light: blue, Dark: blue}
	theme.BorderDimColor = lipgloss.AdaptiveColor{Light: "#eee8d5", Dark: "#001e26"}

	// Diff view colors
	theme.DiffAddedColor = lipgloss.AdaptiveColor{Light: green, Dark: green}
	theme.DiffRemovedColor = lipgloss.AdaptiveColor{Light: red, Dark: red}
	theme.DiffContextColor = lipgloss.AdaptiveColor{Light: base1, Dark: base01}
	theme.DiffHunkHeaderColor = lipgloss.AdaptiveColor{Light: violet, Dark: violet}
	theme.DiffHighlightAddedColor = lipgloss.AdaptiveColor{Light: "#859900", Dark: "#a3be8c"}
	theme.DiffHighlightRemovedColor = lipgloss.AdaptiveColor{Light: "#dc322f", Dark: "#bf616a"}
	theme.DiffAddedBgColor = lipgloss.AdaptiveColor{Light: "#e3f2e1", Dark: "#2a3b2a"}
	theme.DiffRemovedBgColor = lipgloss.AdaptiveColor{Light: "#f9e1e1", Dark: "#3b2a2a"}
	theme.DiffContextBgColor = lipgloss.AdaptiveColor{Light: base3, Dark: base03}
	theme.DiffLineNumberColor = lipgloss.AdaptiveColor{Light: base1, Dark: base01}
	theme.DiffAddedLineNumberBgColor = lipgloss.AdaptiveColor{Light: "#d7ebd5", Dark: "#243224"}
	theme.DiffRemovedLineNumberBgColor = lipgloss.AdaptiveColor{Light: "#f4d7d7", Dark: "#322424"}

	// Markdown colors
	theme.MarkdownTextColor = theme.TextColor
	theme.MarkdownHeadingColor = lipgloss.AdaptiveColor{Light: blue, Dark: blue}
	theme.MarkdownLinkColor = lipgloss.AdaptiveColor{Light: cyan, Dark: cyan}
	theme.MarkdownLinkTextColor = lipgloss.AdaptiveColor{Light: cyan, Dark: cyan}
	theme.MarkdownCodeColor = lipgloss.AdaptiveColor{Light: magenta, Dark: magenta}
	theme.MarkdownBlockQuoteColor = lipgloss.AdaptiveColor{Light: violet, Dark: violet}
	theme.MarkdownEmphColor = lipgloss.AdaptiveColor{Light: yellow, Dark: yellow}
	theme.MarkdownStrongColor = lipgloss.AdaptiveColor{Light: orange, Dark: orange}
	theme.MarkdownHorizontalRuleColor = lipgloss.AdaptiveColor{Light: base2, Dark: base02}
	theme.MarkdownListItemColor = lipgloss.AdaptiveColor{Light: violet, Dark: violet}
	theme.MarkdownListEnumerationColor = lipgloss.AdaptiveColor{Light: blue, Dark: blue}
	theme.MarkdownImageColor = lipgloss.AdaptiveColor{Light: magenta, Dark: magenta}
	theme.MarkdownImageTextColor = lipgloss.AdaptiveColor{Light: cyan, Dark: cyan}
	theme.MarkdownCodeBlockColor = theme.TextColor

	// Syntax highlighting colors
	theme.SyntaxCommentColor = lipgloss.AdaptiveColor{Light: base1, Dark: base01}
	theme.SyntaxKeywordColor = lipgloss.AdaptiveColor{Light: green, Dark: green}
	theme.SyntaxFunctionColor = lipgloss.AdaptiveColor{Light: blue, Dark: blue}
	theme.SyntaxVariableColor = lipgloss.AdaptiveColor{Light: base00, Dark: base0}
	theme.SyntaxStringColor = lipgloss.AdaptiveColor{Light: cyan, Dark: cyan}
	theme.SyntaxNumberColor = lipgloss.AdaptiveColor{Light: violet, Dark: violet}
	theme.SyntaxTypeColor = lipgloss.AdaptiveColor{Light: yellow, Dark: yellow}
	theme.SyntaxOperatorColor = lipgloss.AdaptiveColor{Light: base01, Dark: base1}
	theme.SyntaxPunctuationColor = lipgloss.AdaptiveColor{Light: base00, Dark: base0}

	return theme
}

func init() {
	RegisterTheme("solarized-light", NewSolarizedLightTheme())
}
