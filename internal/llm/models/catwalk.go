package models

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"

	"charm.land/catwalk/pkg/catwalk"
)

var (
	catwalkModels     map[ModelID]Model
	catwalkProviders  map[ModelProvider]bool
	catwalkOnce       *sync.Once = &sync.Once{}
	catwalkMu         sync.RWMutex
	catwalkErr        error
	catwalkURL        = "https://catwalk.charm.sh"
)

// FetchCatwalkModels fetches models from the Catwalk API
// This runs once and caches the results
func FetchCatwalkModels() error {
	catwalkOnce.Do(func() {
		// Allow override via environment variable
		if url := os.Getenv("CATWALK_URL"); url != "" {
			catwalkURL = url
		}

		slog.Info("Fetching models from Catwalk", "url", catwalkURL)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		client := catwalk.NewWithURL(catwalkURL)
		providers, err := client.GetProviders(ctx, "")
		if err != nil {
			catwalkErr = fmt.Errorf("failed to fetch providers from Catwalk: %w", err)
			slog.Warn("Failed to fetch Catwalk models, using static models only", "error", err)
			return
		}

		slog.Debug("Processing Catwalk providers", "count", len(providers))

		tempModels := make(map[ModelID]Model)
		tempProviders := make(map[ModelProvider]bool)

		// Convert Catwalk models to our Model type
		for _, provider := range providers {
			providerID := ModelProvider(provider.ID)
			tempProviders[providerID] = true

			for _, model := range provider.Models {
				modelID := ModelID(model.ID)
				
				tempModels[modelID] = Model{
					ID:                  modelID,
					Name:                model.Name,
					Provider:            providerID,
					APIModel:            model.ID,
					CostPer1MIn:         model.CostPer1MIn,
					CostPer1MOut:        model.CostPer1MOut,
					CostPer1MInCached:   model.CostPer1MInCached,
					CostPer1MOutCached:  model.CostPer1MOutCached,
					ContextWindow:       int64(model.ContextWindow),
					DefaultMaxTokens:    int64(model.DefaultMaxTokens),
					CanReason:           model.CanReason,
					DisableTools:        false, // Catwalk doesn't provide this
				}
			}
		}

		catwalkMu.Lock()
		catwalkModels = tempModels
		catwalkProviders = tempProviders
		catwalkMu.Unlock()

		slog.Info("Successfully fetched Catwalk models", "providers", len(providers), "models", len(catwalkModels))
	})

	return catwalkErr
}

// GetAllModels returns all available models (static + dynamic)
// Static models are used as fallback if Catwalk fetch fails
func GetAllModels() map[ModelID]Model {
	// Try to fetch Catwalk models (cached after first call)
	_ = FetchCatwalkModels()

	allModels := make(map[ModelID]Model)

	// Start with static models as base
	for k, v := range SupportedModels {
		allModels[k] = v
	}

	// Override/add with Catwalk models if available
	catwalkMu.RLock()
	defer catwalkMu.RUnlock()
	
	if len(catwalkModels) > 0 {
		for k, v := range catwalkModels {
			// If model already exists in static models, merge fields that Catwalk doesn't provide
			if staticModel, exists := allModels[k]; exists {
				// Keep static settings if they are more specific
				if staticModel.SupportsAttachments && !v.SupportsAttachments {
					v.SupportsAttachments = true
				}
				if staticModel.DisableTools && !v.DisableTools {
					v.DisableTools = true
				}
				// Keep other static fields if needed
			}
			allModels[k] = v
		}
	}

	return allModels
}

// IsCatwalkProvider checks if a provider is from Catwalk
func IsCatwalkProvider(provider ModelProvider) bool {
	_ = FetchCatwalkModels() // Ensure models are fetched
	catwalkMu.RLock()
	defer catwalkMu.RUnlock()
	return catwalkProviders[provider]
}

// GetCatwalkModels returns only the models fetched from Catwalk
func GetCatwalkModels() map[ModelID]Model {
	_ = FetchCatwalkModels()
	catwalkMu.RLock()
	defer catwalkMu.RUnlock()
	
	// Return a copy to be safe
	res := make(map[ModelID]Model, len(catwalkModels))
	for k, v := range catwalkModels {
		res[k] = v
	}
	return res
}

// RefreshCatwalkModels forces a refresh of Catwalk models
// This resets the sync.Once so models will be fetched again
func RefreshCatwalkModels() error {
	catwalkMu.Lock()
	catwalkOnce = &sync.Once{}
	catwalkMu.Unlock()
	return FetchCatwalkModels()
}
