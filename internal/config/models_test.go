package config

import (
	"strings"
	"testing"
)

func TestDetermineProvider(t *testing.T) {
	tests := []struct {
		name         string
		modelName    string
		wantProvider string
		wantErr      bool
	}{
		{
			name:         "anthropic claude model",
			modelName:    "claude-3-opus",
			wantProvider: "anthropic",
			wantErr:      false,
		},
		{
			name:         "anthropic claude uppercase",
			modelName:    "CLAUDE-3-SONNET",
			wantProvider: "anthropic",
			wantErr:      false,
		},
		{
			name:         "openai gpt model",
			modelName:    "gpt-4",
			wantProvider: "openai",
			wantErr:      false,
		},
		{
			name:         "openai gpt-3.5",
			modelName:    "gpt-3.5-turbo",
			wantProvider: "openai",
			wantErr:      false,
		},
		{
			name:         "openai o1 model",
			modelName:    "o1-preview",
			wantProvider: "openai",
			wantErr:      false,
		},
		{
			name:         "openai o3 model",
			modelName:    "o3-mini",
			wantProvider: "openai",
			wantErr:      false,
		},
		{
			name:         "groq llama model",
			modelName:    "llama-3-70b",
			wantProvider: "groq",
			wantErr:      false,
		},
		{
			name:         "groq mixtral model",
			modelName:    "mixtral-8x7b",
			wantProvider: "groq",
			wantErr:      false,
		},
		{
			name:         "groq gemma model",
			modelName:    "gemma-7b",
			wantProvider: "groq",
			wantErr:      false,
		},
		{
			name:         "ollama llama",
			modelName:    "llama",
			wantProvider: "ollama",
			wantErr:      false,
		},
		{
			name:         "ollama mistral exact match",
			modelName:    "mistral",
			wantProvider: "ollama",
			wantErr:      false,
		},
		{
			name:         "ollama codellama",
			modelName:    "codellama",
			wantProvider: "ollama",
			wantErr:      false,
		},
		{
			name:         "ollama phi",
			modelName:    "phi",
			wantProvider: "ollama",
			wantErr:      false,
		},
		{
			name:         "ollama qwen",
			modelName:    "qwen",
			wantProvider: "ollama",
			wantErr:      false,
		},
		{
			name:         "openrouter model",
			modelName:    "openrouter/anthropic/claude-3-opus",
			wantProvider: "openrouter",
			wantErr:      false,
		},
		{
			name:         "azure openai model",
			modelName:    "azure/gpt-4",
			wantProvider: "azure-openai",
			wantErr:      false,
		},
		{
			name:         "unknown model",
			modelName:    "unknown-model-123",
			wantProvider: "",
			wantErr:      true,
		},
		{
			name:         "empty string",
			modelName:    "",
			wantProvider: "",
			wantErr:      true,
		},
		{
			name:         "case insensitive mixed case",
			modelName:    "GPT-4-Turbo",
			wantProvider: "openai",
			wantErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DetermineProvider(tt.modelName)

			if tt.wantErr {
				if err == nil {
					t.Errorf("DetermineProvider(%q) expected error but got none", tt.modelName)
					return
				}
				if !strings.Contains(err.Error(), "unable to determine provider") {
					t.Errorf("DetermineProvider(%q) error = %v, want error containing 'unable to determine provider'", tt.modelName, err)
				}
				return
			}

			if err != nil {
				t.Errorf("DetermineProvider(%q) unexpected error = %v", tt.modelName, err)
				return
			}

			if got != tt.wantProvider {
				t.Errorf("DetermineProvider(%q) = %q, want %q", tt.modelName, got, tt.wantProvider)
			}
		})
	}
}

func TestDetermineProviderPrefixMatching(t *testing.T) {
	tests := []struct {
		name         string
		modelName    string
		wantProvider string
	}{
		{
			name:         "claude prefix match",
			modelName:    "claude-custom-variant",
			wantProvider: "anthropic",
		},
		{
			name:         "gpt prefix match",
			modelName:    "gpt-custom-5",
			wantProvider: "openai",
		},
		{
			name:         "llama- prefix should match groq not ollama",
			modelName:    "llama-special",
			wantProvider: "groq",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DetermineProvider(tt.modelName)
			if err != nil {
				t.Errorf("DetermineProvider(%q) unexpected error = %v", tt.modelName, err)
				return
			}

			if got != tt.wantProvider {
				t.Errorf("DetermineProvider(%q) = %q, want %q", tt.modelName, got, tt.wantProvider)
			}
		})
	}
}
