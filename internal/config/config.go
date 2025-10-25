package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type ProviderConfig struct {
	Token string `json:"token"`
}

type Config struct {
	ActiveProvider string                    `json:"active_provider,omitempty"`
	ActiveModel    string                    `json:"active_model,omitempty"`
	Providers      map[string]ProviderConfig `json:"providers"`
}

var validProviders = map[string]bool{
	"openai":       true,
	"anthropic":    true,
	"groq":         true,
	"ollama":       true,
	"mistral":      true,
	"openrouter":   true,
	"azure-openai": true,
}

func IsValidProvider(provider string) bool {
	return validProviders[provider]
}

func configPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}

	configDir := filepath.Join(home, ".config", "pseudolang")
	return filepath.Join(configDir, "config.json"), nil
}

func Load() (*Config, error) {
	path, err := configPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{
				Providers: make(map[string]ProviderConfig),
			}, nil
		}
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	if cfg.Providers == nil {
		cfg.Providers = make(map[string]ProviderConfig)
	}

	return &cfg, nil
}

func (c *Config) Save() error {
	path, err := configPath()
	if err != nil {
		return err
	}

	configDir := filepath.Dir(path)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

func (c *Config) GetToken(provider string) (string, bool) {
	if providerCfg, ok := c.Providers[provider]; ok {
		return providerCfg.Token, true
	}
	return "", false
}

func (c *Config) SetProviderToken(provider, token string) {
	if c.Providers == nil {
		c.Providers = make(map[string]ProviderConfig)
	}
	c.Providers[provider] = ProviderConfig{Token: token}
}

func (c *Config) SetActiveProvider(provider string) error {
	if _, ok := c.Providers[provider]; !ok {
		return fmt.Errorf("no token configured for provider: %s", provider)
	}
	c.ActiveProvider = provider
	return nil
}

func (c *Config) SetActiveModel(model string) error {
	provider, err := DetermineProvider(model)
	if err != nil {
		return err
	}

	if _, ok := c.Providers[provider]; !ok {
		return fmt.Errorf("no token configured for provider: %s (required for model: %s)", provider, model)
	}

	c.ActiveProvider = provider
	c.ActiveModel = model
	return nil
}

func (c *Config) SetModelWithToken(model, token string) error {
	provider, err := DetermineProvider(model)
	if err != nil {
		return err
	}

	c.SetProviderToken(provider, token)
	c.ActiveProvider = provider
	c.ActiveModel = model
	return nil
}
