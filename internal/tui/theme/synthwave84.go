package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// SynthwaveTheme implements the Theme interface with Synthwave '84 colors.
type SynthwaveTheme struct {
	BaseTheme
}

// NewSynthwaveTheme creates a new instance of the Synthwave theme.
func NewSynthwaveTheme() *SynthwaveTheme {
	// Synthwave '84 color palette
	background := "#262335"
	foreground := "#ffffff"
	comment := "#848bb2"
	pink := "#ff7edb"
	yellow := "#fede5d"
	orange := "#f97e72"
	cyan := "#03edf9"
	purple := "#b893ff"
	green := "#72f1b8"
	selection := "#34294f"
	border := "#44475a"

	theme := &SynthwaveTheme{}

	// Base colors
	theme.PrimaryColor = lipgloss.AdaptiveColor{Dark: pink, Light: "#d310ea"}
	theme.SecondaryColor = lipgloss.AdaptiveColor{Dark: purple, Light: "#731ad2"}
	theme.AccentColor = lipgloss.AdaptiveColor{Dark: cyan, Light: "#00b1b1"}

	// Status colors
	theme.ErrorColor = lipgloss.AdaptiveColor{Dark: "#fe4450", Light: "#ff0000"}
	theme.WarningColor = lipgloss.AdaptiveColor{Dark: yellow, Light: "#ffcc00"}
	theme.SuccessColor = lipgloss.AdaptiveColor{Dark: green, Light: "#2ecc71"}
	theme.InfoColor = lipgloss.AdaptiveColor{Dark: cyan, Light: "#00b1b1"}

	// Text colors
	theme.TextColor = lipgloss.AdaptiveColor{Dark: foreground, Light: "#262335"}
	theme.TextMutedColor = lipgloss.AdaptiveColor{Dark: comment, Light: "#6b7294"}
	theme.TextEmphasizedColor = lipgloss.AdaptiveColor{Dark: "#fffa00", Light: "#000000"}

	// Background colors
	theme.BackgroundColor = lipgloss.AdaptiveColor{Dark: background, Light: "#fdfdfd"}
	theme.BackgroundSecondaryColor = lipgloss.AdaptiveColor{Dark: selection, Light: "#f0eff5"}
	theme.BackgroundDarkerColor = lipgloss.AdaptiveColor{Dark: "#201b2d", Light: "#eeeeee"}

	// Border colors
	theme.BorderNormalColor = lipgloss.AdaptiveColor{Dark: border, Light: "#d1d1d1"}
	theme.BorderFocusedColor = lipgloss.AdaptiveColor{Dark: pink, Light: "#d310ea"}
	theme.BorderDimColor = lipgloss.AdaptiveColor{Dark: "#2d2a3e", Light: "#eeeeee"}

	// Diff view colors
	theme.DiffAddedColor = lipgloss.AdaptiveColor{Dark: green, Light: "#2ecc71"}
	theme.DiffRemovedColor = lipgloss.AdaptiveColor{Dark: "#fe4450", Light: "#ff0000"}
	theme.DiffContextColor = lipgloss.AdaptiveColor{Dark: comment, Light: "#6b7294"}
	theme.DiffHunkHeaderColor = lipgloss.AdaptiveColor{Dark: purple, Light: "#731ad2"}
	theme.DiffHighlightAddedColor = lipgloss.AdaptiveColor{Dark: "#72f1b8", Light: "#4caf50"}
	theme.DiffHighlightRemovedColor = lipgloss.AdaptiveColor{Dark: "#fe4450", Light: "#ff5252"}
	theme.DiffAddedBgColor = lipgloss.AdaptiveColor{Dark: "#2d3b34", Light: "#e8f5e9"}
	theme.DiffRemovedBgColor = lipgloss.AdaptiveColor{Dark: "#3b2d2f", Light: "#ffebee"}
	theme.DiffContextBgColor = lipgloss.AdaptiveColor{Dark: background, Light: "#fdfdfd"}
	theme.DiffLineNumberColor = lipgloss.AdaptiveColor{Dark: comment, Light: "#8e8e8e"}
	theme.DiffAddedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#25332c", Light: "#c8e6c9"}
	theme.DiffRemovedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#332527", Light: "#ffcdd2"}

	// Markdown colors
	theme.MarkdownTextColor = theme.TextColor
	theme.MarkdownHeadingColor = lipgloss.AdaptiveColor{Dark: pink, Light: "#d310ea"}
	theme.MarkdownLinkColor = lipgloss.AdaptiveColor{Dark: cyan, Light: "#00b1b1"}
	theme.MarkdownLinkTextColor = lipgloss.AdaptiveColor{Dark: cyan, Light: "#00b1b1"}
	theme.MarkdownCodeColor = lipgloss.AdaptiveColor{Dark: yellow, Light: "#ffcc00"}
	theme.MarkdownBlockQuoteColor = lipgloss.AdaptiveColor{Dark: purple, Light: "#731ad2"}
	theme.MarkdownEmphColor = lipgloss.AdaptiveColor{Dark: yellow, Light: "#ffcc00"}
	theme.MarkdownStrongColor = lipgloss.AdaptiveColor{Dark: orange, Light: "#ff5e00"}
	theme.MarkdownHorizontalRuleColor = lipgloss.AdaptiveColor{Dark: border, Light: "#d1d1d1"}
	theme.MarkdownListItemColor = lipgloss.AdaptiveColor{Dark: pink, Light: "#d310ea"}
	theme.MarkdownListEnumerationColor = lipgloss.AdaptiveColor{Dark: cyan, Light: "#00b1b1"}
	theme.MarkdownImageColor = lipgloss.AdaptiveColor{Dark: pink, Light: "#d310ea"}
	theme.MarkdownImageTextColor = lipgloss.AdaptiveColor{Dark: cyan, Light: "#00b1b1"}
	theme.MarkdownCodeBlockColor = theme.TextColor

	// Syntax highlighting colors
	theme.SyntaxCommentColor = lipgloss.AdaptiveColor{Dark: comment, Light: "#6b7294"}
	theme.SyntaxKeywordColor = lipgloss.AdaptiveColor{Dark: pink, Light: "#d310ea"}
	theme.SyntaxFunctionColor = lipgloss.AdaptiveColor{Dark: cyan, Light: "#00b1b1"}
	theme.SyntaxVariableColor = lipgloss.AdaptiveColor{Dark: purple, Light: "#731ad2"}
	theme.SyntaxStringColor = lipgloss.AdaptiveColor{Dark: yellow, Light: "#ffcc00"}
	theme.SyntaxNumberColor = lipgloss.AdaptiveColor{Dark: orange, Light: "#ff5e00"}
	theme.SyntaxTypeColor = lipgloss.AdaptiveColor{Dark: purple, Light: "#731ad2"}
	theme.SyntaxOperatorColor = lipgloss.AdaptiveColor{Dark: pink, Light: "#d310ea"}
	theme.SyntaxPunctuationColor = lipgloss.AdaptiveColor{Dark: foreground, Light: "#262335"}

	return theme
}

func init() {
	RegisterTheme("synthwave84", NewSynthwaveTheme())
}
