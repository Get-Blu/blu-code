package models

const (
	ProviderVertexAI ModelProvider = "vertexai"

	// Models
	VertexAIGemini25Flash ModelID = "vertexai.gemini-2.5-flash"
	VertexAIGemini25      ModelID = "vertexai.gemini-2.5"
	VertexAIGemini30Flash ModelID = "vertexai.gemini-3.0-flash"
	VertexAIGemini30Pro   ModelID = "vertexai.gemini-3.0-pro"
	VertexAIGemini3DeepThink ModelID = "vertexai.gemini-3-deep-think"
)

var VertexAIGeminiModels = map[ModelID]Model{
	VertexAIGemini25Flash: {
		ID:                  VertexAIGemini25Flash,
		Name:                "VertexAI: Gemini 2.5 Flash",
		Provider:            ProviderVertexAI,
		APIModel:            "gemini-2.5-flash-preview-04-17",
		CostPer1MIn:         GeminiModels[Gemini25Flash].CostPer1MIn,
		CostPer1MInCached:   GeminiModels[Gemini25Flash].CostPer1MInCached,
		CostPer1MOut:        GeminiModels[Gemini25Flash].CostPer1MOut,
		CostPer1MOutCached:  GeminiModels[Gemini25Flash].CostPer1MOutCached,
		ContextWindow:       GeminiModels[Gemini25Flash].ContextWindow,
		DefaultMaxTokens:    GeminiModels[Gemini25Flash].DefaultMaxTokens,
		SupportsAttachments: true,
	},
	VertexAIGemini25: {
		ID:                  VertexAIGemini25,
		Name:                "VertexAI: Gemini 2.5 Pro",
		Provider:            ProviderVertexAI,
		APIModel:            "gemini-2.5-pro-preview-03-25",
		CostPer1MIn:         GeminiModels[Gemini25].CostPer1MIn,
		CostPer1MInCached:   GeminiModels[Gemini25].CostPer1MInCached,
		CostPer1MOut:        GeminiModels[Gemini25].CostPer1MOut,
		CostPer1MOutCached:  GeminiModels[Gemini25].CostPer1MOutCached,
		ContextWindow:       GeminiModels[Gemini25].ContextWindow,
		DefaultMaxTokens:    GeminiModels[Gemini25].DefaultMaxTokens,
		SupportsAttachments: true,
	},
	VertexAIGemini30Flash: {
		ID:                  VertexAIGemini30Flash,
		Name:                "VertexAI: Gemini 3.0 Flash",
		Provider:            ProviderVertexAI,
		APIModel:            "gemini-3.0-flash",
		CostPer1MIn:         GeminiModels[Gemini30Flash].CostPer1MIn,
		CostPer1MInCached:   GeminiModels[Gemini30Flash].CostPer1MInCached,
		CostPer1MOut:        GeminiModels[Gemini30Flash].CostPer1MOut,
		CostPer1MOutCached:  GeminiModels[Gemini30Flash].CostPer1MOutCached,
		ContextWindow:       GeminiModels[Gemini30Flash].ContextWindow,
		DefaultMaxTokens:    GeminiModels[Gemini30Flash].DefaultMaxTokens,
		SupportsAttachments: true,
	},
	VertexAIGemini30Pro: {
		ID:                  VertexAIGemini30Pro,
		Name:                "VertexAI: Gemini 3.0 Pro",
		Provider:            ProviderVertexAI,
		APIModel:            "gemini-3.0-pro",
		CostPer1MIn:         GeminiModels[Gemini30Pro].CostPer1MIn,
		CostPer1MInCached:   GeminiModels[Gemini30Pro].CostPer1MInCached,
		CostPer1MOut:        GeminiModels[Gemini30Pro].CostPer1MOut,
		CostPer1MOutCached:  GeminiModels[Gemini30Pro].CostPer1MOutCached,
		ContextWindow:       GeminiModels[Gemini30Pro].ContextWindow,
		DefaultMaxTokens:    GeminiModels[Gemini30Pro].DefaultMaxTokens,
		SupportsAttachments: true,
	},
	VertexAIGemini3DeepThink: {
		ID:                  VertexAIGemini3DeepThink,
		Name:                "VertexAI: Gemini 3 Deep Think",
		Provider:            ProviderVertexAI,
		APIModel:            "gemini-3-deep-think",
		CostPer1MIn:         GeminiModels[Gemini3DeepThink].CostPer1MIn,
		CostPer1MInCached:   GeminiModels[Gemini3DeepThink].CostPer1MInCached,
		CostPer1MOut:        GeminiModels[Gemini3DeepThink].CostPer1MOut,
		CostPer1MOutCached:  GeminiModels[Gemini3DeepThink].CostPer1MOutCached,
		ContextWindow:       GeminiModels[Gemini3DeepThink].ContextWindow,
		DefaultMaxTokens:    GeminiModels[Gemini3DeepThink].DefaultMaxTokens,
		SupportsAttachments: true,
	},
}
