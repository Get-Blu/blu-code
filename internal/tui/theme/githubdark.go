package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// GithubDarkTheme implements the Theme interface with Github Dark colors.
type GithubDarkTheme struct {
	BaseTheme
}

// NewGithubDarkTheme creates a new instance of the Github Dark theme.
func NewGithubDarkTheme() *GithubDarkTheme {
	theme := &GithubDarkTheme{}

	// Github Dark Colors
	const (
		ghdBg      = "#0d1117"
		ghdFg      = "#c9d1d9"
		ghdText    = "#c9d1d9"
		ghdMuted   = "#8b949e"
		ghdBlue    = "#58a6ff"
		ghdGreen   = "#3fb950"
		ghdRed     = "#ff7b72"
		ghdOrange  = "#ffa657"
		ghdPurple  = "#bc8cff"
		ghdBorder  = "#30363d"
		ghdSelect  = "#161b22"
	)

	// Base colors
	theme.PrimaryColor = lipgloss.AdaptiveColor{Dark: ghdBlue, Light: "#0969da"}
	theme.SecondaryColor = lipgloss.AdaptiveColor{Dark: ghdPurple, Light: "#8250df"}
	theme.AccentColor = lipgloss.AdaptiveColor{Dark: ghdOrange, Light: "#bc4c00"}

	// Status colors
	theme.ErrorColor = lipgloss.AdaptiveColor{Dark: ghdRed, Light: "#cf222e"}
	theme.WarningColor = lipgloss.AdaptiveColor{Dark: ghdOrange, Light: "#9a6700"}
	theme.SuccessColor = lipgloss.AdaptiveColor{Dark: ghdGreen, Light: "#1a7f37"}
	theme.InfoColor = lipgloss.AdaptiveColor{Dark: ghdBlue, Light: "#0969da"}

	// Text colors
	theme.TextColor = lipgloss.AdaptiveColor{Dark: ghdText, Light: "#1f2328"}
	theme.TextMutedColor = lipgloss.AdaptiveColor{Dark: ghdMuted, Light: "#656d76"}
	theme.TextEmphasizedColor = lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"}

	// Background colors
	theme.BackgroundColor = lipgloss.AdaptiveColor{Dark: ghdBg, Light: "#ffffff"}
	theme.BackgroundSecondaryColor = lipgloss.AdaptiveColor{Dark: "#161b22", Light: "#f6f8fa"}
	theme.BackgroundDarkerColor = lipgloss.AdaptiveColor{Dark: "#010409", Light: "#eff1f3"}

	// Border colors
	theme.BorderNormalColor = lipgloss.AdaptiveColor{Dark: ghdBorder, Light: "#d0d7de"}
	theme.BorderFocusedColor = lipgloss.AdaptiveColor{Dark: ghdBlue, Light: "#0969da"}
	theme.BorderDimColor = lipgloss.AdaptiveColor{Dark: "#21262d", Light: "#afb8c1"}

	// Diff view colors
	theme.DiffAddedColor = lipgloss.AdaptiveColor{Dark: ghdGreen, Light: "#1a7f37"}
	theme.DiffRemovedColor = lipgloss.AdaptiveColor{Dark: ghdRed, Light: "#cf222e"}
	theme.DiffContextColor = lipgloss.AdaptiveColor{Dark: ghdMuted, Light: "#656d76"}
	theme.DiffHunkHeaderColor = lipgloss.AdaptiveColor{Dark: ghdBlue, Light: "#0969da"}
	theme.DiffHighlightAddedColor = lipgloss.AdaptiveColor{Dark: "#aff5b4", Light: "#acf2bd"}
	theme.DiffHighlightRemovedColor = lipgloss.AdaptiveColor{Dark: "#ffdcd7", Light: "#ffd0d0"}
	theme.DiffAddedBgColor = lipgloss.AdaptiveColor{Dark: "#121d13", Light: "#dafbe1"}
	theme.DiffRemovedBgColor = lipgloss.AdaptiveColor{Dark: "#1d1212", Light: "#ffebe9"}
	theme.DiffContextBgColor = lipgloss.AdaptiveColor{Dark: ghdBg, Light: "#ffffff"}
	theme.DiffLineNumberColor = lipgloss.AdaptiveColor{Dark: ghdMuted, Light: "#6e7781"}
	theme.DiffAddedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#192a1c", Light: "#cce8d4"}
	theme.DiffRemovedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#2a1919", Light: "#ffd7d5"}

	// Markdown colors
	theme.MarkdownTextColor = lipgloss.AdaptiveColor{Dark: ghdText, Light: "#1f2328"}
	theme.MarkdownHeadingColor = lipgloss.AdaptiveColor{Dark: ghdBlue, Light: "#0969da"}
	theme.MarkdownLinkColor = lipgloss.AdaptiveColor{Dark: ghdBlue, Light: "#0969da"}
	theme.MarkdownLinkTextColor = lipgloss.AdaptiveColor{Dark: ghdPurple, Light: "#8250df"}
	theme.MarkdownCodeColor = lipgloss.AdaptiveColor{Dark: ghdGreen, Light: "#1a7f37"}
	theme.MarkdownBlockQuoteColor = lipgloss.AdaptiveColor{Dark: ghdMuted, Light: "#656d76"}
	theme.MarkdownEmphColor = lipgloss.AdaptiveColor{Dark: ghdOrange, Light: "#9a6700"}
	theme.MarkdownStrongColor = lipgloss.AdaptiveColor{Dark: ghdOrange, Light: "#9a6700"}
	theme.MarkdownHorizontalRuleColor = lipgloss.AdaptiveColor{Dark: ghdBorder, Light: "#d0d7de"}
	theme.MarkdownListItemColor = lipgloss.AdaptiveColor{Dark: ghdBlue, Light: "#0969da"}
	theme.MarkdownListEnumerationColor = lipgloss.AdaptiveColor{Dark: ghdBlue, Light: "#0969da"}
	theme.MarkdownImageColor = lipgloss.AdaptiveColor{Dark: ghdPurple, Light: "#8250df"}
	theme.MarkdownImageTextColor = lipgloss.AdaptiveColor{Dark: ghdBlue, Light: "#0969da"}
	theme.MarkdownCodeBlockColor = lipgloss.AdaptiveColor{Dark: ghdText, Light: "#1f2328"}

	// Syntax highlighting colors
	theme.SyntaxCommentColor = lipgloss.AdaptiveColor{Dark: ghdMuted, Light: "#656d76"}
	theme.SyntaxKeywordColor = lipgloss.AdaptiveColor{Dark: ghdRed, Light: "#cf222e"}
	theme.SyntaxFunctionColor = lipgloss.AdaptiveColor{Dark: ghdPurple, Light: "#8250df"}
	theme.SyntaxVariableColor = lipgloss.AdaptiveColor{Dark: ghdOrange, Light: "#bc4c00"}
	theme.SyntaxStringColor = lipgloss.AdaptiveColor{Dark: ghdBlue, Light: "#0a3069"}
	theme.SyntaxNumberColor = lipgloss.AdaptiveColor{Dark: ghdBlue, Light: "#0550ae"}
	theme.SyntaxTypeColor = lipgloss.AdaptiveColor{Dark: ghdRed, Light: "#cf222e"}
	theme.SyntaxOperatorColor = lipgloss.AdaptiveColor{Dark: ghdBlue, Light: "#0969da"}
	theme.SyntaxPunctuationColor = lipgloss.AdaptiveColor{Dark: ghdText, Light: "#1f2328"}

	return theme
}

func init() {
	RegisterTheme("githubdark", NewGithubDarkTheme())
}
