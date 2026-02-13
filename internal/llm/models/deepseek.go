package models

const (
	ProviderDeepSeek ModelProvider = "deepseek"

	DeepSeekV3     ModelID = "deepseek-v3"
	DeepSeekV3_2   ModelID = "deepseek-v3.2"
	DeepSeekR1     ModelID = "deepseek-r1"
	DeepSeekV4     ModelID = "deepseek-v4"
)

var DeepSeekModels = map[ModelID]Model{
	DeepSeekV3: {
		ID:                 DeepSeekV3,
		Name:               "DeepSeek V3",
		Provider:           ProviderDeepSeek,
		APIModel:           "deepseek-chat",
		CostPer1MIn:        0.14,
		CostPer1MInCached:  0.07,
		CostPer1MOut:       0.28,
		CostPer1MOutCached: 0,
		ContextWindow:      128_000,
		DefaultMaxTokens:   8192,
	},
	DeepSeekV3_2: {
		ID:                 DeepSeekV3_2,
		Name:               "DeepSeek V3.2",
		Provider:           ProviderDeepSeek,
		APIModel:           "deepseek-chat-v3.2",
		CostPer1MIn:        0.10,
		CostPer1MInCached:  0.05,
		CostPer1MOut:       0.20,
		CostPer1MOutCached: 0,
		ContextWindow:      128_000,
		DefaultMaxTokens:   16384,
	},
	DeepSeekR1: {
		ID:                 DeepSeekR1,
		Name:               "DeepSeek R1",
		Provider:           ProviderDeepSeek,
		APIModel:           "deepseek-reasoner",
		CostPer1MIn:        0.55,
		CostPer1MInCached:  0.14,
		CostPer1MOut:       2.19,
		CostPer1MOutCached: 0,
		ContextWindow:      128_000,
		DefaultMaxTokens:   32768,
		CanReason:          true,
	},
	DeepSeekV4: {
		ID:                 DeepSeekV4,
		Name:               "DeepSeek V4 (Coming Soon)",
		Provider:           ProviderDeepSeek,
		APIModel:           "deepseek-v4",
		CostPer1MIn:        0.08,
		CostPer1MInCached:  0.04,
		CostPer1MOut:       0.16,
		CostPer1MOutCached: 0,
		ContextWindow:      1_000_000,
		DefaultMaxTokens:   64_000,
		CanReason:          true,
	},
}
