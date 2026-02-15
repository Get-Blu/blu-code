package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// MaterialDeepOceanTheme implements the Theme interface with Material Deep Ocean colors.
type MaterialDeepOceanTheme struct {
	BaseTheme
}

// NewMaterialDeepOceanTheme creates a new instance of the Material Deep Ocean theme.
func NewMaterialDeepOceanTheme() *MaterialDeepOceanTheme {
	theme := &MaterialDeepOceanTheme{}

	// Material Deep Ocean Colors
	const (
		mdoBg      = "#0f111a"
		mdoFg      = "#8f93a2"
		mdoText    = "#eeffff"
		mdoSelection = "#1f2233"
		mdoRed     = "#f07178"
		mdoGreen   = "#c3e88d"
		mdoYellow  = "#ffcb6b"
		mdoBlue    = "#82aaff"
		mdoPurple  = "#c792ea"
		mdoCyan    = "#89ddff"
		mdoOrange  = "#f78c6c"
		mdoComment = "#464b5d"
	)

	// Base colors
	theme.PrimaryColor = lipgloss.AdaptiveColor{Dark: mdoBlue, Light: "#303f9f"}
	theme.SecondaryColor = lipgloss.AdaptiveColor{Dark: mdoPurple, Light: "#7b1fa2"}
	theme.AccentColor = lipgloss.AdaptiveColor{Dark: mdoCyan, Light: "#0097a7"}

	// Status colors
	theme.ErrorColor = lipgloss.AdaptiveColor{Dark: mdoRed, Light: "#d32f2f"}
	theme.WarningColor = lipgloss.AdaptiveColor{Dark: mdoYellow, Light: "#f57c00"}
	theme.SuccessColor = lipgloss.AdaptiveColor{Dark: mdoGreen, Light: "#388e3c"}
	theme.InfoColor = lipgloss.AdaptiveColor{Dark: mdoCyan, Light: "#1976d2"}

	// Text colors
	theme.TextColor = lipgloss.AdaptiveColor{Dark: mdoText, Light: "#0f111a"}
	theme.TextMutedColor = lipgloss.AdaptiveColor{Dark: mdoFg, Light: "#464b5d"}
	theme.TextEmphasizedColor = lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"}

	// Background colors
	theme.BackgroundColor = lipgloss.AdaptiveColor{Dark: mdoBg, Light: "#ffffff"}
	theme.BackgroundSecondaryColor = lipgloss.AdaptiveColor{Dark: "#1a1c25", Light: "#f5f5f5"}
	theme.BackgroundDarkerColor = lipgloss.AdaptiveColor{Dark: "#090b10", Light: "#e0e0e0"}

	// Border colors
	theme.BorderNormalColor = lipgloss.AdaptiveColor{Dark: "#1f2233", Light: "#bdbdbd"}
	theme.BorderFocusedColor = lipgloss.AdaptiveColor{Dark: mdoBlue, Light: "#303f9f"}
	theme.BorderDimColor = lipgloss.AdaptiveColor{Dark: "#151720", Light: "#eeeeee"}

	// Diff view colors
	theme.DiffAddedColor = lipgloss.AdaptiveColor{Dark: mdoGreen, Light: "#2e7d32"}
	theme.DiffRemovedColor = lipgloss.AdaptiveColor{Dark: mdoRed, Light: "#c62828"}
	theme.DiffContextColor = lipgloss.AdaptiveColor{Dark: mdoComment, Light: "#757575"}
	theme.DiffHunkHeaderColor = lipgloss.AdaptiveColor{Dark: mdoBlue, Light: "#1976d2"}
	theme.DiffHighlightAddedColor = lipgloss.AdaptiveColor{Dark: "#d6f5d6", Light: "#c8e6c9"}
	theme.DiffHighlightRemovedColor = lipgloss.AdaptiveColor{Dark: "#f5d6d6", Light: "#ffcdd2"}
	theme.DiffAddedBgColor = lipgloss.AdaptiveColor{Dark: "#1a251a", Light: "#e8f5e9"}
	theme.DiffRemovedBgColor = lipgloss.AdaptiveColor{Dark: "#251a1a", Light: "#ffebee"}
	theme.DiffContextBgColor = lipgloss.AdaptiveColor{Dark: mdoBg, Light: "#ffffff"}
	theme.DiffLineNumberColor = lipgloss.AdaptiveColor{Dark: mdoComment, Light: "#9e9e9e"}
	theme.DiffAddedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#202c20", Light: "#c8e6c9"}
	theme.DiffRemovedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#2c2020", Light: "#ffcdd2"}

	// Markdown colors
	theme.MarkdownTextColor = lipgloss.AdaptiveColor{Dark: mdoText, Light: "#0f111a"}
	theme.MarkdownHeadingColor = lipgloss.AdaptiveColor{Dark: mdoBlue, Light: "#303f9f"}
	theme.MarkdownLinkColor = lipgloss.AdaptiveColor{Dark: mdoCyan, Light: "#0097a7"}
	theme.MarkdownLinkTextColor = lipgloss.AdaptiveColor{Dark: mdoPurple, Light: "#7b1fa2"}
	theme.MarkdownCodeColor = lipgloss.AdaptiveColor{Dark: mdoGreen, Light: "#388e3c"}
	theme.MarkdownBlockQuoteColor = lipgloss.AdaptiveColor{Dark: mdoComment, Light: "#464b5d"}
	theme.MarkdownEmphColor = lipgloss.AdaptiveColor{Dark: mdoYellow, Light: "#f57c00"}
	theme.MarkdownStrongColor = lipgloss.AdaptiveColor{Dark: mdoOrange, Light: "#f57c00"}
	theme.MarkdownHorizontalRuleColor = lipgloss.AdaptiveColor{Dark: "#1f2233", Light: "#bdbdbd"}
	theme.MarkdownListItemColor = lipgloss.AdaptiveColor{Dark: mdoBlue, Light: "#303f9f"}
	theme.MarkdownListEnumerationColor = lipgloss.AdaptiveColor{Dark: mdoBlue, Light: "#303f9f"}
	theme.MarkdownImageColor = lipgloss.AdaptiveColor{Dark: mdoPurple, Light: "#7b1fa2"}
	theme.MarkdownImageTextColor = lipgloss.AdaptiveColor{Dark: mdoBlue, Light: "#303f9f"}
	theme.MarkdownCodeBlockColor = lipgloss.AdaptiveColor{Dark: mdoText, Light: "#0f111a"}

	// Syntax highlighting colors
	theme.SyntaxCommentColor = lipgloss.AdaptiveColor{Dark: mdoComment, Light: "#464b5d"}
	theme.SyntaxKeywordColor = lipgloss.AdaptiveColor{Dark: mdoPurple, Light: "#7b1fa2"}
	theme.SyntaxFunctionColor = lipgloss.AdaptiveColor{Dark: mdoBlue, Light: "#303f9f"}
	theme.SyntaxVariableColor = lipgloss.AdaptiveColor{Dark: mdoText, Light: "#0f111a"}
	theme.SyntaxStringColor = lipgloss.AdaptiveColor{Dark: mdoGreen, Light: "#388e3c"}
	theme.SyntaxNumberColor = lipgloss.AdaptiveColor{Dark: mdoOrange, Light: "#f57c00"}
	theme.SyntaxTypeColor = lipgloss.AdaptiveColor{Dark: mdoYellow, Light: "#fbc02d"}
	theme.SyntaxOperatorColor = lipgloss.AdaptiveColor{Dark: mdoCyan, Light: "#0097a7"}
	theme.SyntaxPunctuationColor = lipgloss.AdaptiveColor{Dark: mdoCyan, Light: "#0097a7"}

	return theme
}

func init() {
	RegisterTheme("materialdeepocean", NewMaterialDeepOceanTheme())
}
