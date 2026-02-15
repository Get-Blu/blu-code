package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// RosePineTheme implements the Theme interface with Rose Pine colors.
type RosePineTheme struct {
	BaseTheme
}

// NewRosePineTheme creates a new instance of the Rose Pine theme.
func NewRosePineTheme() *RosePineTheme {
	theme := &RosePineTheme{}

	// Rose Pine Colors (Dark)
	const (
		rpBase    = "#191724"
		rpSurface = "#1f1d2e"
		rpOverlay = "#26233a"
		rpMuted   = "#6e6a86"
		rpSubtle  = "#908caa"
		rpText    = "#e0def4"
		rpLove    = "#eb6f92"
		rpGold    = "#f6c177"
		rpRose    = "#ebbcba"
		rpPine    = "#31748f"
		rpFoam    = "#9ccfd8"
		rpIris    = "#c4a7e7"
	)

	// Rose Pine Dawn Colors (Light)
	const (
		rpdBase    = "#faf4ed"
		rpdSurface = "#fffaf3"
		rpdOverlay = "#f2e9de"
		rpdMuted   = "#9893a5"
		rpdSubtle  = "#797593"
		rpdText    = "#575279"
		rpdLove    = "#b4637a"
		rpdGold    = "#ea9d34"
		rpdRose    = "#d7827e"
		rpdPine    = "#286983"
		rpdFoam    = "#56949f"
		rpdIris    = "#907aa9"
	)

	// Base colors
	theme.PrimaryColor = lipgloss.AdaptiveColor{Dark: rpIris, Light: rpdIris}
	theme.SecondaryColor = lipgloss.AdaptiveColor{Dark: rpPine, Light: rpdPine}
	theme.AccentColor = lipgloss.AdaptiveColor{Dark: rpLove, Light: rpdLove}

	// Status colors
	theme.ErrorColor = lipgloss.AdaptiveColor{Dark: rpLove, Light: rpdLove}
	theme.WarningColor = lipgloss.AdaptiveColor{Dark: rpGold, Light: rpdGold}
	theme.SuccessColor = lipgloss.AdaptiveColor{Dark: rpPine, Light: rpdPine}
	theme.InfoColor = lipgloss.AdaptiveColor{Dark: rpFoam, Light: rpdFoam}

	// Text colors
	theme.TextColor = lipgloss.AdaptiveColor{Dark: rpText, Light: rpdText}
	theme.TextMutedColor = lipgloss.AdaptiveColor{Dark: rpMuted, Light: rpdMuted}
	theme.TextEmphasizedColor = lipgloss.AdaptiveColor{Dark: rpRose, Light: rpdRose}

	// Background colors
	theme.BackgroundColor = lipgloss.AdaptiveColor{Dark: rpBase, Light: rpdBase}
	theme.BackgroundSecondaryColor = lipgloss.AdaptiveColor{Dark: rpSurface, Light: rpdSurface}
	theme.BackgroundDarkerColor = lipgloss.AdaptiveColor{Dark: "#12101a", Light: "#f2e9de"}

	// Border colors
	theme.BorderNormalColor = lipgloss.AdaptiveColor{Dark: rpOverlay, Light: rpdOverlay}
	theme.BorderFocusedColor = lipgloss.AdaptiveColor{Dark: rpIris, Light: rpdIris}
	theme.BorderDimColor = lipgloss.AdaptiveColor{Dark: rpSurface, Light: rpdSurface}

	// Diff view colors
	theme.DiffAddedColor = lipgloss.AdaptiveColor{Dark: rpPine, Light: rpdPine}
	theme.DiffRemovedColor = lipgloss.AdaptiveColor{Dark: rpLove, Light: rpdLove}
	theme.DiffContextColor = lipgloss.AdaptiveColor{Dark: rpMuted, Light: rpdMuted}
	theme.DiffHunkHeaderColor = lipgloss.AdaptiveColor{Dark: rpIris, Light: rpdIris}
	theme.DiffHighlightAddedColor = lipgloss.AdaptiveColor{Dark: rpFoam, Light: rpdFoam}
	theme.DiffHighlightRemovedColor = lipgloss.AdaptiveColor{Dark: rpRose, Light: rpdRose}
	theme.DiffAddedBgColor = lipgloss.AdaptiveColor{Dark: "#223943", Light: "#d9e5e9"}
	theme.DiffRemovedBgColor = lipgloss.AdaptiveColor{Dark: "#3a2530", Light: "#f2e2e7"}
	theme.DiffContextBgColor = lipgloss.AdaptiveColor{Dark: rpBase, Light: rpdBase}
	theme.DiffLineNumberColor = lipgloss.AdaptiveColor{Dark: rpSubtle, Light: rpdSubtle}
	theme.DiffAddedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#263f4a", Light: "#cce0e6"}
	theme.DiffRemovedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#452d3a", Light: "#ecd3da"}

	// Markdown colors
	theme.MarkdownTextColor = lipgloss.AdaptiveColor{Dark: rpText, Light: rpdText}
	theme.MarkdownHeadingColor = lipgloss.AdaptiveColor{Dark: rpIris, Light: rpdIris}
	theme.MarkdownLinkColor = lipgloss.AdaptiveColor{Dark: rpFoam, Light: rpdFoam}
	theme.MarkdownLinkTextColor = lipgloss.AdaptiveColor{Dark: rpLove, Light: rpdLove}
	theme.MarkdownCodeColor = lipgloss.AdaptiveColor{Dark: rpGold, Light: rpdGold}
	theme.MarkdownBlockQuoteColor = lipgloss.AdaptiveColor{Dark: rpMuted, Light: rpdMuted}
	theme.MarkdownEmphColor = lipgloss.AdaptiveColor{Dark: rpRose, Light: rpdRose}
	theme.MarkdownStrongColor = lipgloss.AdaptiveColor{Dark: rpGold, Light: rpdGold}
	theme.MarkdownHorizontalRuleColor = lipgloss.AdaptiveColor{Dark: rpOverlay, Light: rpdOverlay}
	theme.MarkdownListItemColor = lipgloss.AdaptiveColor{Dark: rpIris, Light: rpdIris}
	theme.MarkdownListEnumerationColor = lipgloss.AdaptiveColor{Dark: rpPine, Light: rpdPine}
	theme.MarkdownImageColor = lipgloss.AdaptiveColor{Dark: rpRose, Light: rpdRose}
	theme.MarkdownImageTextColor = lipgloss.AdaptiveColor{Dark: rpIris, Light: rpdIris}
	theme.MarkdownCodeBlockColor = lipgloss.AdaptiveColor{Dark: rpText, Light: rpdText}

	// Syntax highlighting colors
	theme.SyntaxCommentColor = lipgloss.AdaptiveColor{Dark: rpMuted, Light: rpdMuted}
	theme.SyntaxKeywordColor = lipgloss.AdaptiveColor{Dark: rpPine, Light: rpdPine}
	theme.SyntaxFunctionColor = lipgloss.AdaptiveColor{Dark: rpRose, Light: rpdRose}
	theme.SyntaxVariableColor = lipgloss.AdaptiveColor{Dark: rpText, Light: rpdText}
	theme.SyntaxStringColor = lipgloss.AdaptiveColor{Dark: rpGold, Light: rpdGold}
	theme.SyntaxNumberColor = lipgloss.AdaptiveColor{Dark: rpGold, Light: rpdGold}
	theme.SyntaxTypeColor = lipgloss.AdaptiveColor{Dark: rpFoam, Light: rpdFoam}
	theme.SyntaxOperatorColor = lipgloss.AdaptiveColor{Dark: rpSubtle, Light: rpdSubtle}
	theme.SyntaxPunctuationColor = lipgloss.AdaptiveColor{Dark: rpSubtle, Light: rpdSubtle}

	return theme
}

func init() {
	RegisterTheme("rosepine", NewRosePineTheme())
}
