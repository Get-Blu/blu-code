package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// OneLightTheme implements the Theme interface with Atom One Light colors.
type OneLightTheme struct {
	BaseTheme
}

// NewOneLightTheme creates a new instance of the One Light theme.
func NewOneLightTheme() *OneLightTheme {
	theme := &OneLightTheme{}

	// One Light Colors
	const (
		olBg      = "#fafafa"
		olFg      = "#383a42"
		olMono1   = "#383a42"
		olMono2   = "#696c77"
		olMono3   = "#a0a1a7"
		olCyan    = "#0184bc"
		olBlue    = "#4078f2"
		olPurple  = "#a626a4"
		olGreen   = "#50a14f"
		olRed1    = "#e45649"
		olRed2    = "#ca1243"
		olOrange1 = "#d75f00"
		olOrange2 = "#986801"
	)

	// Base colors
	theme.PrimaryColor = lipgloss.AdaptiveColor{Dark: olBlue, Light: olBlue}
	theme.SecondaryColor = lipgloss.AdaptiveColor{Dark: olPurple, Light: olPurple}
	theme.AccentColor = lipgloss.AdaptiveColor{Dark: olOrange1, Light: olOrange1}

	// Status colors
	theme.ErrorColor = lipgloss.AdaptiveColor{Dark: olRed1, Light: olRed1}
	theme.WarningColor = lipgloss.AdaptiveColor{Dark: olOrange2, Light: olOrange2}
	theme.SuccessColor = lipgloss.AdaptiveColor{Dark: olGreen, Light: olGreen}
	theme.InfoColor = lipgloss.AdaptiveColor{Dark: olCyan, Light: olCyan}

	// Text colors
	theme.TextColor = lipgloss.AdaptiveColor{Dark: olFg, Light: olFg}
	theme.TextMutedColor = lipgloss.AdaptiveColor{Dark: olMono2, Light: olMono2}
	theme.TextEmphasizedColor = lipgloss.AdaptiveColor{Dark: olMono1, Light: olMono1}

	// Background colors
	theme.BackgroundColor = lipgloss.AdaptiveColor{Dark: olBg, Light: olBg}
	theme.BackgroundSecondaryColor = lipgloss.AdaptiveColor{Dark: "#f0f0f0", Light: "#f0f0f0"}
	theme.BackgroundDarkerColor = lipgloss.AdaptiveColor{Dark: "#e5e5e5", Light: "#e5e5e5"}

	// Border colors
	theme.BorderNormalColor = lipgloss.AdaptiveColor{Dark: "#dbdbdb", Light: "#dbdbdb"}
	theme.BorderFocusedColor = lipgloss.AdaptiveColor{Dark: olBlue, Light: olBlue}
	theme.BorderDimColor = lipgloss.AdaptiveColor{Dark: "#ececec", Light: "#ececec"}

	// Diff view colors
	theme.DiffAddedColor = lipgloss.AdaptiveColor{Dark: olGreen, Light: olGreen}
	theme.DiffRemovedColor = lipgloss.AdaptiveColor{Dark: olRed1, Light: olRed1}
	theme.DiffContextColor = lipgloss.AdaptiveColor{Dark: olMono2, Light: olMono2}
	theme.DiffHunkHeaderColor = lipgloss.AdaptiveColor{Dark: olPurple, Light: olPurple}
	theme.DiffHighlightAddedColor = lipgloss.AdaptiveColor{Dark: "#dff0df", Light: "#dff0df"}
	theme.DiffHighlightRemovedColor = lipgloss.AdaptiveColor{Dark: "#f7e8e8", Light: "#f7e8e8"}
	theme.DiffAddedBgColor = lipgloss.AdaptiveColor{Dark: "#f0fff0", Light: "#f0fff0"}
	theme.DiffRemovedBgColor = lipgloss.AdaptiveColor{Dark: "#fff0f0", Light: "#fff0f0"}
	theme.DiffContextBgColor = lipgloss.AdaptiveColor{Dark: olBg, Light: olBg}
	theme.DiffLineNumberColor = lipgloss.AdaptiveColor{Dark: olMono3, Light: olMono3}
	theme.DiffAddedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#e0ffe0", Light: "#e0ffe0"}
	theme.DiffRemovedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#ffe0e0", Light: "#ffe0e0"}

	// Markdown colors
	theme.MarkdownTextColor = lipgloss.AdaptiveColor{Dark: olFg, Light: olFg}
	theme.MarkdownHeadingColor = lipgloss.AdaptiveColor{Dark: olPurple, Light: olPurple}
	theme.MarkdownLinkColor = lipgloss.AdaptiveColor{Dark: olBlue, Light: olBlue}
	theme.MarkdownLinkTextColor = lipgloss.AdaptiveColor{Dark: olRed2, Light: olRed2}
	theme.MarkdownCodeColor = lipgloss.AdaptiveColor{Dark: olGreen, Light: olGreen}
	theme.MarkdownBlockQuoteColor = lipgloss.AdaptiveColor{Dark: olMono2, Light: olMono2}
	theme.MarkdownEmphColor = lipgloss.AdaptiveColor{Dark: olOrange2, Light: olOrange2}
	theme.MarkdownStrongColor = lipgloss.AdaptiveColor{Dark: olOrange1, Light: olOrange1}
	theme.MarkdownHorizontalRuleColor = lipgloss.AdaptiveColor{Dark: "#dbdbdb", Light: "#dbdbdb"}
	theme.MarkdownListItemColor = lipgloss.AdaptiveColor{Dark: olBlue, Light: olBlue}
	theme.MarkdownListEnumerationColor = lipgloss.AdaptiveColor{Dark: olBlue, Light: olBlue}
	theme.MarkdownImageColor = lipgloss.AdaptiveColor{Dark: olPurple, Light: olPurple}
	theme.MarkdownImageTextColor = lipgloss.AdaptiveColor{Dark: olPurple, Light: olPurple}
	theme.MarkdownCodeBlockColor = lipgloss.AdaptiveColor{Dark: olFg, Light: olFg}

	// Syntax highlighting colors
	theme.SyntaxCommentColor = lipgloss.AdaptiveColor{Dark: olMono3, Light: olMono3}
	theme.SyntaxKeywordColor = lipgloss.AdaptiveColor{Dark: olPurple, Light: olPurple}
	theme.SyntaxFunctionColor = lipgloss.AdaptiveColor{Dark: olBlue, Light: olBlue}
	theme.SyntaxVariableColor = lipgloss.AdaptiveColor{Dark: olRed1, Light: olRed1}
	theme.SyntaxStringColor = lipgloss.AdaptiveColor{Dark: olGreen, Light: olGreen}
	theme.SyntaxNumberColor = lipgloss.AdaptiveColor{Dark: olOrange1, Light: olOrange1}
	theme.SyntaxTypeColor = lipgloss.AdaptiveColor{Dark: olOrange2, Light: olOrange2}
	theme.SyntaxOperatorColor = lipgloss.AdaptiveColor{Dark: olCyan, Light: olCyan}
	theme.SyntaxPunctuationColor = lipgloss.AdaptiveColor{Dark: olMono2, Light: olMono2}

	return theme
}

func init() {
	RegisterTheme("onelight", NewOneLightTheme())
}
