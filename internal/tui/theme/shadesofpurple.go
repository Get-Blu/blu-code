package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// ShadesOfPurpleTheme implements the Theme interface with Shades of Purple colors.
type ShadesOfPurpleTheme struct {
	BaseTheme
}

// NewShadesOfPurpleTheme creates a new instance of the Shades of Purple theme.
func NewShadesOfPurpleTheme() *ShadesOfPurpleTheme {
	theme := &ShadesOfPurpleTheme{}

	// Shades of Purple Colors
	const (
		sopBg       = "#2d2b55"
		sopSelection = "#b362ff"
		sopKeyword   = "#ff9d00"
		sopFunction  = "#fad000"
		sopString    = "#3ad900"
		sopComment   = "#b362ff"
		sopNumber    = "#ff628c"
		sopText      = "#ffffff"
		sopMuted     = "#a599e9"
		sopAccent    = "#fb9e00"
		sopError     = "#ec3a37"
		sopSuccess   = "#3ad900"
		sopInfo      = "#00a6ff"
	)

	// Base colors
	theme.PrimaryColor = lipgloss.AdaptiveColor{Dark: sopFunction, Light: "#5f43e9"}
	theme.SecondaryColor = lipgloss.AdaptiveColor{Dark: sopKeyword, Light: "#e91e63"}
	theme.AccentColor = lipgloss.AdaptiveColor{Dark: sopAccent, Light: "#ff9800"}

	// Status colors
	theme.ErrorColor = lipgloss.AdaptiveColor{Dark: sopError, Light: "#d32f2f"}
	theme.WarningColor = lipgloss.AdaptiveColor{Dark: sopKeyword, Light: "#f57c00"}
	theme.SuccessColor = lipgloss.AdaptiveColor{Dark: sopSuccess, Light: "#388e3c"}
	theme.InfoColor = lipgloss.AdaptiveColor{Dark: sopInfo, Light: "#1976d2"}

	// Text colors
	theme.TextColor = lipgloss.AdaptiveColor{Dark: sopText, Light: "#2d2b55"}
	theme.TextMutedColor = lipgloss.AdaptiveColor{Dark: sopMuted, Light: "#5d5d5d"}
	theme.TextEmphasizedColor = lipgloss.AdaptiveColor{Dark: sopFunction, Light: "#fad000"}

	// Background colors
	theme.BackgroundColor = lipgloss.AdaptiveColor{Dark: sopBg, Light: "#f0f0f0"}
	theme.BackgroundSecondaryColor = lipgloss.AdaptiveColor{Dark: "#1e1e3f", Light: "#e0e0e0"}
	theme.BackgroundDarkerColor = lipgloss.AdaptiveColor{Dark: "#181830", Light: "#d5d5d5"}

	// Border colors
	theme.BorderNormalColor = lipgloss.AdaptiveColor{Dark: "#4d4d80", Light: "#b0b0b0"}
	theme.BorderFocusedColor = lipgloss.AdaptiveColor{Dark: sopFunction, Light: "#5f43e9"}
	theme.BorderDimColor = lipgloss.AdaptiveColor{Dark: "#3d3d66", Light: "#cccccc"}

	// Diff view colors
	theme.DiffAddedColor = lipgloss.AdaptiveColor{Dark: sopSuccess, Light: "#2e7d32"}
	theme.DiffRemovedColor = lipgloss.AdaptiveColor{Dark: sopError, Light: "#c62828"}
	theme.DiffContextColor = lipgloss.AdaptiveColor{Dark: sopMuted, Light: "#757575"}
	theme.DiffHunkHeaderColor = lipgloss.AdaptiveColor{Dark: sopFunction, Light: "#1976d2"}
	theme.DiffHighlightAddedColor = lipgloss.AdaptiveColor{Dark: "#a6ff00", Light: "#4caf50"}
	theme.DiffHighlightRemovedColor = lipgloss.AdaptiveColor{Dark: "#ff628c", Light: "#ef5350"}
	theme.DiffAddedBgColor = lipgloss.AdaptiveColor{Dark: "#233323", Light: "#e8f5e9"}
	theme.DiffRemovedBgColor = lipgloss.AdaptiveColor{Dark: "#332323", Light: "#ffebee"}
	theme.DiffContextBgColor = lipgloss.AdaptiveColor{Dark: sopBg, Light: "#f0f0f0"}
	theme.DiffLineNumberColor = lipgloss.AdaptiveColor{Dark: sopMuted, Light: "#9e9e9e"}
	theme.DiffAddedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#2a3d2a", Light: "#c8e6c9"}
	theme.DiffRemovedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#3d2a2a", Light: "#ffcdd2"}

	// Markdown colors
	theme.MarkdownTextColor = lipgloss.AdaptiveColor{Dark: sopText, Light: "#2d2b55"}
	theme.MarkdownHeadingColor = lipgloss.AdaptiveColor{Dark: sopFunction, Light: "#5f43e9"}
	theme.MarkdownLinkColor = lipgloss.AdaptiveColor{Dark: sopInfo, Light: "#00a6ff"}
	theme.MarkdownLinkTextColor = lipgloss.AdaptiveColor{Dark: sopKeyword, Light: "#ff9d00"}
	theme.MarkdownCodeColor = lipgloss.AdaptiveColor{Dark: sopString, Light: "#3ad900"}
	theme.MarkdownBlockQuoteColor = lipgloss.AdaptiveColor{Dark: sopMuted, Light: "#a599e9"}
	theme.MarkdownEmphColor = lipgloss.AdaptiveColor{Dark: sopFunction, Light: "#fad000"}
	theme.MarkdownStrongColor = lipgloss.AdaptiveColor{Dark: sopKeyword, Light: "#ff9d00"}
	theme.MarkdownHorizontalRuleColor = lipgloss.AdaptiveColor{Dark: "#4d4d80", Light: "#b0b0b0"}
	theme.MarkdownListItemColor = lipgloss.AdaptiveColor{Dark: sopFunction, Light: "#5f43e9"}
	theme.MarkdownListEnumerationColor = lipgloss.AdaptiveColor{Dark: sopFunction, Light: "#5f43e9"}
	theme.MarkdownImageColor = lipgloss.AdaptiveColor{Dark: sopKeyword, Light: "#ff9d00"}
	theme.MarkdownImageTextColor = lipgloss.AdaptiveColor{Dark: sopFunction, Light: "#fad000"}
	theme.MarkdownCodeBlockColor = lipgloss.AdaptiveColor{Dark: sopText, Light: "#2d2b55"}

	// Syntax highlighting colors
	theme.SyntaxCommentColor = lipgloss.AdaptiveColor{Dark: sopComment, Light: "#b362ff"}
	theme.SyntaxKeywordColor = lipgloss.AdaptiveColor{Dark: sopKeyword, Light: "#e91e63"}
	theme.SyntaxFunctionColor = lipgloss.AdaptiveColor{Dark: sopFunction, Light: "#fad000"}
	theme.SyntaxVariableColor = lipgloss.AdaptiveColor{Dark: "#9effff", Light: "#00acc1"}
	theme.SyntaxStringColor = lipgloss.AdaptiveColor{Dark: sopString, Light: "#3ad900"}
	theme.SyntaxNumberColor = lipgloss.AdaptiveColor{Dark: sopNumber, Light: "#ff628c"}
	theme.SyntaxTypeColor = lipgloss.AdaptiveColor{Dark: "#9effff", Light: "#00acc1"}
	theme.SyntaxOperatorColor = lipgloss.AdaptiveColor{Dark: sopKeyword, Light: "#ff9d00"}
	theme.SyntaxPunctuationColor = lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#2d2b55"}

	return theme
}

func init() {
	RegisterTheme("shadesofpurple", NewShadesOfPurpleTheme())
}
