package provider

import (
	"errors"
	"net/http"
	"testing"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestAnthropicShouldRetry(t *testing.T) {
	client := &anthropicClient{}

	tests := []struct {
		name        string
		attempts    int
		err         error
		wantRetry   bool
		wantErr     bool
		expectedErr string
	}{
		{
			name:      "Retry on 429",
			attempts:  1,
			err:       &anthropic.Error{StatusCode: 429, Response: &http.Response{Header: make(http.Header)}},
			wantRetry: true,
			wantErr:   false,
		},
		{
			name:      "Retry on 529",
			attempts:  1,
			err:       &anthropic.Error{StatusCode: 529, Response: &http.Response{Header: make(http.Header)}},
			wantRetry: true,
			wantErr:   false,
		},
		{
			name:      "No retry on 400",
			attempts:  1,
			err:       &anthropic.Error{StatusCode: 400},
			wantRetry: false,
			wantErr:   true,
		},
		{
			name:      "No retry on random error",
			attempts:  1,
			err:       errors.New("random error"),
			wantRetry: false,
			wantErr:   true,
		},
		{
			name:        "Max retries reached",
			attempts:    maxRetries + 1,
			err:         &anthropic.Error{StatusCode: 429, Response: &http.Response{Header: make(http.Header)}},
			wantRetry:   false,
			wantErr:     true,
			expectedErr: "maximum retry attempts reached",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retry, _, err := client.shouldRetry(tt.attempts, tt.err)
			assert.Equal(t, tt.wantRetry, retry)
			if tt.wantErr {
				assert.Error(t, err)
				if tt.expectedErr != "" {
					assert.Contains(t, err.Error(), tt.expectedErr)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
