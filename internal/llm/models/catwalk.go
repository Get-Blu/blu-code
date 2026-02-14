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
	catwalkOnce       sync.Once
	catwalkErr        error
	catwalkURL        = "https://catwalk.charm.land"
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

		catwalkModels = make(map[ModelID]Model)
		catwalkProviders = make(map[ModelProvider]bool)

		// Convert Catwalk models to our Model type
		for _, provider := range providers {
			providerID := ModelProvider(provider.ID)
			catwalkProviders[providerID] = true

			for _, model := range provider.Models {
				modelID := ModelID(model.ID)
				
				catwalkModels[modelID] = Model{
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
	if len(catwalkModels) > 0 {
		for k, v := range catwalkModels {
			allModels[k] = v
		}
	}

	return allModels
}

// IsCatwalkProvider checks if a provider is from Catwalk
func IsCatwalkProvider(provider ModelProvider) bool {
	_ = FetchCatwalkModels() // Ensure models are fetched
	return catwalkProviders[provider]
}

// GetCatwalkModels returns only the models fetched from Catwalk
func GetCatwalkModels() map[ModelID]Model {
	_ = FetchCatwalkModels()
	return catwalkModels
}

// RefreshCatwalkModels forces a refresh of Catwalk models
// This resets the sync.Once so models will be fetched again
func RefreshCatwalkModels() error {
	catwalkOnce = sync.Once{}
	return FetchCatwalkModels()
}
