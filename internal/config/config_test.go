package config

import (
	"strings"
	"testing"
)

func TestIsValidProvider(t *testing.T) {
	tests := []struct {
		name     string
		provider string
		want     bool
	}{
		{
			name:     "openai is valid",
			provider: "openai",
			want:     true,
		},
		{
			name:     "anthropic is valid",
			provider: "anthropic",
			want:     true,
		},
		{
			name:     "groq is valid",
			provider: "groq",
			want:     true,
		},
		{
			name:     "ollama is valid",
			provider: "ollama",
			want:     true,
		},
		{
			name:     "mistral is valid",
			provider: "mistral",
			want:     true,
		},
		{
			name:     "openrouter is valid",
			provider: "openrouter",
			want:     true,
		},
		{
			name:     "azure-openai is valid",
			provider: "azure-openai",
			want:     true,
		},
		{
			name:     "invalid provider",
			provider: "invalid-provider",
			want:     false,
		},
		{
			name:     "empty string",
			provider: "",
			want:     false,
		},
		{
			name:     "case sensitive check",
			provider: "OpenAI",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidProvider(tt.provider)
			if got != tt.want {
				t.Errorf("IsValidProvider(%q) = %v, want %v", tt.provider, got, tt.want)
			}
		})
	}
}

func TestConfig_GetToken(t *testing.T) {
	tests := []struct {
		name      string
		config    *Config
		provider  string
		wantToken string
		wantOk    bool
	}{
		{
			name: "get existing token",
			config: &Config{
				Providers: map[string]ProviderConfig{
					"openai": {Token: "sk-test123"},
				},
			},
			provider:  "openai",
			wantToken: "sk-test123",
			wantOk:    true,
		},
		{
			name: "get non-existent provider",
			config: &Config{
				Providers: map[string]ProviderConfig{
					"openai": {Token: "sk-test123"},
				},
			},
			provider:  "anthropic",
			wantToken: "",
			wantOk:    false,
		},
		{
			name: "empty providers map",
			config: &Config{
				Providers: map[string]ProviderConfig{},
			},
			provider:  "openai",
			wantToken: "",
			wantOk:    false,
		},
		{
			name: "nil providers map",
			config: &Config{
				Providers: nil,
			},
			provider:  "openai",
			wantToken: "",
			wantOk:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotToken, gotOk := tt.config.GetToken(tt.provider)
			if gotToken != tt.wantToken {
				t.Errorf("Config.GetToken(%q) token = %q, want %q", tt.provider, gotToken, tt.wantToken)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Config.GetToken(%q) ok = %v, want %v", tt.provider, gotOk, tt.wantOk)
			}
		})
	}
}

func TestConfig_SetProviderToken(t *testing.T) {
	tests := []struct {
		name     string
		config   *Config
		provider string
		token    string
	}{
		{
			name: "set token on existing providers map",
			config: &Config{
				Providers: map[string]ProviderConfig{},
			},
			provider: "openai",
			token:    "sk-test123",
		},
		{
			name: "set token on nil providers map",
			config: &Config{
				Providers: nil,
			},
			provider: "anthropic",
			token:    "sk-ant-test456",
		},
		{
			name: "override existing token",
			config: &Config{
				Providers: map[string]ProviderConfig{
					"openai": {Token: "old-token"},
				},
			},
			provider: "openai",
			token:    "new-token",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.config.SetProviderToken(tt.provider, tt.token)

			got, ok := tt.config.GetToken(tt.provider)
			if !ok {
				t.Errorf("Config.SetProviderToken(%q, %q) did not set token", tt.provider, tt.token)
				return
			}
			if got != tt.token {
				t.Errorf("Config.SetProviderToken(%q, %q) set token = %q, want %q", tt.provider, tt.token, got, tt.token)
			}
		})
	}
}

func TestConfig_SetActiveProvider(t *testing.T) {
	tests := []struct {
		name     string
		config   *Config
		provider string
		wantErr  bool
	}{
		{
			name: "set active provider with token",
			config: &Config{
				Providers: map[string]ProviderConfig{
					"openai": {Token: "sk-test123"},
				},
			},
			provider: "openai",
			wantErr:  false,
		},
		{
			name: "set active provider without token",
			config: &Config{
				Providers: map[string]ProviderConfig{},
			},
			provider: "openai",
			wantErr:  true,
		},
		{
			name: "set active provider with nil map",
			config: &Config{
				Providers: nil,
			},
			provider: "openai",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.SetActiveProvider(tt.provider)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Config.SetActiveProvider(%q) expected error but got none", tt.provider)
					return
				}
				if !strings.Contains(err.Error(), "no token configured") {
					t.Errorf("Config.SetActiveProvider(%q) error = %v, want error containing 'no token configured'", tt.provider, err)
				}
				return
			}

			if err != nil {
				t.Errorf("Config.SetActiveProvider(%q) unexpected error = %v", tt.provider, err)
				return
			}

			if tt.config.ActiveProvider != tt.provider {
				t.Errorf("Config.SetActiveProvider(%q) set ActiveProvider = %q, want %q", tt.provider, tt.config.ActiveProvider, tt.provider)
			}
		})
	}
}

func TestConfig_SetActiveModel(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		model   string
		wantErr bool
	}{
		{
			name: "set model with matching provider token",
			config: &Config{
				Providers: map[string]ProviderConfig{
					"openai": {Token: "sk-test123"},
				},
			},
			model:   "gpt-4",
			wantErr: false,
		},
		{
			name: "set model without matching provider token",
			config: &Config{
				Providers: map[string]ProviderConfig{
					"anthropic": {Token: "sk-ant-test"},
				},
			},
			model:   "gpt-4",
			wantErr: true,
		},
		{
			name: "set unknown model",
			config: &Config{
				Providers: map[string]ProviderConfig{},
			},
			model:   "unknown-model",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.SetActiveModel(tt.model)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Config.SetActiveModel(%q) expected error but got none", tt.model)
				}
				return
			}

			if err != nil {
				t.Errorf("Config.SetActiveModel(%q) unexpected error = %v", tt.model, err)
				return
			}

			if tt.config.ActiveModel != tt.model {
				t.Errorf("Config.SetActiveModel(%q) set ActiveModel = %q, want %q", tt.model, tt.config.ActiveModel, tt.model)
			}
		})
	}
}

func TestConfig_SetModelWithToken(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		model   string
		token   string
		wantErr bool
	}{
		{
			name: "set model and token for valid model",
			config: &Config{
				Providers: map[string]ProviderConfig{},
			},
			model:   "gpt-4",
			token:   "sk-test123",
			wantErr: false,
		},
		{
			name: "set model and token for claude",
			config: &Config{
				Providers: map[string]ProviderConfig{},
			},
			model:   "claude-3-opus",
			token:   "sk-ant-test",
			wantErr: false,
		},
		{
			name: "set unknown model with token",
			config: &Config{
				Providers: map[string]ProviderConfig{},
			},
			model:   "unknown-model",
			token:   "some-token",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.SetModelWithToken(tt.model, tt.token)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Config.SetModelWithToken(%q, %q) expected error but got none", tt.model, tt.token)
				}
				return
			}

			if err != nil {
				t.Errorf("Config.SetModelWithToken(%q, %q) unexpected error = %v", tt.model, tt.token, err)
				return
			}

			if tt.config.ActiveModel != tt.model {
				t.Errorf("Config.SetModelWithToken(%q, %q) set ActiveModel = %q, want %q", tt.model, tt.token, tt.config.ActiveModel, tt.model)
			}

			provider, _ := DetermineProvider(tt.model)
			gotToken, ok := tt.config.GetToken(provider)
			if !ok {
				t.Errorf("Config.SetModelWithToken(%q, %q) did not set provider token", tt.model, tt.token)
				return
			}
			if gotToken != tt.token {
				t.Errorf("Config.SetModelWithToken(%q, %q) set provider token = %q, want %q", tt.model, tt.token, gotToken, tt.token)
			}
		})
	}
}
