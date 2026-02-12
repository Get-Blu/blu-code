package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// GithubLightTheme implements the Theme interface with Github Light colors.
type GithubLightTheme struct {
	BaseTheme
}

// NewGithubLightTheme creates a new instance of the Github Light theme.
func NewGithubLightTheme() *GithubLightTheme {
	// Github Light color palette
	background := "#ffffff"
	foreground := "#24292e"
	selection := "#0366d61a" // approximation
	border := "#e1e4e8"
	comment := "#6a737d"
	red := "#d73a49"
	green := "#28a745"
	yellow := "#dbab09"
	orange := "#e36209"
	blue := "#0366d6"
	purple := "#6f42c1"
	
	theme := &GithubLightTheme{}

	// Base colors
	theme.PrimaryColor = lipgloss.AdaptiveColor{Light: blue, Dark: "#58a6ff"}
	theme.SecondaryColor = lipgloss.AdaptiveColor{Light: purple, Dark: "#bc8cff"}
	theme.AccentColor = lipgloss.AdaptiveColor{Light: blue, Dark: "#58a6ff"}

	// Status colors
	theme.ErrorColor = lipgloss.AdaptiveColor{Light: red, Dark: "#f85149"}
	theme.WarningColor = lipgloss.AdaptiveColor{Light: orange, Dark: "#d29922"}
	theme.SuccessColor = lipgloss.AdaptiveColor{Light: green, Dark: "#3fb950"}
	theme.InfoColor = lipgloss.AdaptiveColor{Light: blue, Dark: "#58a6ff"}

	// Text colors
	theme.TextColor = lipgloss.AdaptiveColor{Light: foreground, Dark: "#c9d1d9"}
	theme.TextMutedColor = lipgloss.AdaptiveColor{Light: comment, Dark: "#8b949e"}
	theme.TextEmphasizedColor = lipgloss.AdaptiveColor{Light: "#000000", Dark: "#ffffff"}

	// Background colors
	theme.BackgroundColor = lipgloss.AdaptiveColor{Light: background, Dark: "#0d1117"}
	theme.BackgroundSecondaryColor = lipgloss.AdaptiveColor{Light: "#f6f8fa", Dark: "#161b22"}
	theme.BackgroundDarkerColor = lipgloss.AdaptiveColor{Light: "#f1f2f4", Dark: "#010409"}

	// Border colors
	theme.BorderNormalColor = lipgloss.AdaptiveColor{Light: border, Dark: "#30363d"}
	theme.BorderFocusedColor = lipgloss.AdaptiveColor{Light: blue, Dark: "#58a6ff"}
	theme.BorderDimColor = lipgloss.AdaptiveColor{Light: "#e1e4e8", Dark: "#21262d"}

	// Diff view colors
	theme.DiffAddedColor = lipgloss.AdaptiveColor{Light: green, Dark: "#3fb950"}
	theme.DiffRemovedColor = lipgloss.AdaptiveColor{Light: red, Dark: "#f85149"}
	theme.DiffContextColor = lipgloss.AdaptiveColor{Light: comment, Dark: "#8b949e"}
	theme.DiffHunkHeaderColor = lipgloss.AdaptiveColor{Light: blue, Dark: "#58a6ff"}
	theme.DiffHighlightAddedColor = lipgloss.AdaptiveColor{Light: "#acf2bd", Dark: "#2ea043"}
	theme.DiffHighlightRemovedColor = lipgloss.AdaptiveColor{Light: "#fdb8c0", Dark: "#f85149"}
	theme.DiffAddedBgColor = lipgloss.AdaptiveColor{Light: "#e6ffed", Dark: "#121d12"}
	theme.DiffRemovedBgColor = lipgloss.AdaptiveColor{Light: "#ffeef0", Dark: "#1d1212"}
	theme.DiffContextBgColor = lipgloss.AdaptiveColor{Light: background, Dark: "#0d1117"}
	theme.DiffLineNumberColor = lipgloss.AdaptiveColor{Light: "#babbbc", Dark: "#6e7681"}
	theme.DiffAddedLineNumberBgColor = lipgloss.AdaptiveColor{Light: "#cdffd8", Dark: "#152415"}
	theme.DiffRemovedLineNumberBgColor = lipgloss.AdaptiveColor{Light: "#ffdce0", Dark: "#241515"}

	// Markdown colors
	theme.MarkdownTextColor = theme.TextColor
	theme.MarkdownHeadingColor = lipgloss.AdaptiveColor{Light: "#24292e", Dark: "#c9d1d9"}
	theme.MarkdownLinkColor = lipgloss.AdaptiveColor{Light: "#0366d6", Dark: "#58a6ff"}
	theme.MarkdownLinkTextColor = lipgloss.AdaptiveColor{Light: "#0366d6", Dark: "#58a6ff"}
	theme.MarkdownCodeColor = lipgloss.AdaptiveColor{Light: "#24292e", Dark: "#c9d1d9"}
	theme.MarkdownBlockQuoteColor = lipgloss.AdaptiveColor{Light: "#6a737d", Dark: "#8b949e"}
	theme.MarkdownEmphColor = lipgloss.AdaptiveColor{Light: "#24292e", Dark: "#c9d1d9"}
	theme.MarkdownStrongColor = lipgloss.AdaptiveColor{Light: "#24292e", Dark: "#c9d1d9"}
	theme.MarkdownHorizontalRuleColor = lipgloss.AdaptiveColor{Light: border, Dark: "#30363d"}
	theme.MarkdownListItemColor = lipgloss.AdaptiveColor{Light: "#24292e", Dark: "#c9d1d9"}
	theme.MarkdownListEnumerationColor = lipgloss.AdaptiveColor{Light: "#24292e", Dark: "#c9d1d9"}
	theme.MarkdownImageColor = lipgloss.AdaptiveColor{Light: "#24292e", Dark: "#c9d1d9"}
	theme.MarkdownImageTextColor = lipgloss.AdaptiveColor{Light: "#0366d6", Dark: "#58a6ff"}
	theme.MarkdownCodeBlockColor = theme.TextColor

	// Syntax highlighting colors
	theme.SyntaxCommentColor = lipgloss.AdaptiveColor{Light: comment, Dark: "#8b949e"}
	theme.SyntaxKeywordColor = lipgloss.AdaptiveColor{Light: red, Dark: "#ff7b72"}
	theme.SyntaxFunctionColor = lipgloss.AdaptiveColor{Light: purple, Dark: "#d2a8ff"}
	theme.SyntaxVariableColor = lipgloss.AdaptiveColor{Light: orange, Dark: "#ffa657"}
	theme.SyntaxStringColor = lipgloss.AdaptiveColor{Light: "#032f62", Dark: "#a5d6ff"}
	theme.SyntaxNumberColor = lipgloss.AdaptiveColor{Light: "#005cc5", Dark: "#79c0ff"}
	theme.SyntaxTypeColor = lipgloss.AdaptiveColor{Light: red, Dark: "#ff7b72"}
	theme.SyntaxOperatorColor = lipgloss.AdaptiveColor{Light: blue, Dark: "#58a6ff"}
	theme.SyntaxPunctuationColor = lipgloss.AdaptiveColor{Light: foreground, Dark: "#c9d1d9"}

	return theme
}

func init() {
	RegisterTheme("github-light", NewGithubLightTheme())
}
