package models

const (
	ProviderMistral ModelProvider = "mistral"

	MistralLarge3   ModelID = "mistral-large-3"
	Ministral3_8B   ModelID = "ministral-3-8b"
	Ministral3_3B   ModelID = "ministral-3-3b"
	MistralSmall3_1 ModelID = "mistral-small-3.1"
)

var MistralModels = map[ModelID]Model{
	MistralLarge3: {
		ID:                 MistralLarge3,
		Name:               "Mistral Large 3",
		Provider:           ProviderMistral,
		APIModel:           "mistral-large-2411", // Latest as of search
		CostPer1MIn:        2.0,
		CostPer1MInCached:  1.0,
		CostPer1MOut:       6.0,
		CostPer1MOutCached: 0,
		ContextWindow:      128_000,
		DefaultMaxTokens:   32000,
	},
	Ministral3_8B: {
		ID:                 Ministral3_8B,
		Name:               "Ministral 3 8B",
		Provider:           ProviderMistral,
		APIModel:           "ministral-8b-latest",
		CostPer1MIn:        0.1,
		CostPer1MInCached:  0.05,
		CostPer1MOut:       0.3,
		CostPer1MOutCached: 0,
		ContextWindow:      128_000,
		DefaultMaxTokens:   32000,
	},
	Ministral3_3B: {
		ID:                 Ministral3_3B,
		Name:               "Ministral 3 3B",
		Provider:           ProviderMistral,
		APIModel:           "ministral-3b-latest",
		CostPer1MIn:        0.04,
		CostPer1MInCached:  0.02,
		CostPer1MOut:       0.04,
		CostPer1MOutCached: 0,
		ContextWindow:      128_000,
		DefaultMaxTokens:   32000,
	},
	MistralSmall3_1: {
		ID:                 MistralSmall3_1,
		Name:               "Mistral Small 3.1",
		Provider:           ProviderMistral,
		APIModel:           "mistral-small-latest",
		CostPer1MIn:        0.2,
		CostPer1MInCached:  0.1,
		CostPer1MOut:       0.6,
		CostPer1MOutCached: 0,
		ContextWindow:      128_000,
		DefaultMaxTokens:   32000,
	},
}
