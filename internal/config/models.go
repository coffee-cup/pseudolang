package config

import (
	"fmt"
	"strings"
)

var modelPatterns = map[string][]string{
	"anthropic":    {"claude-"},
	"openai":       {"gpt-", "o1-", "o3-"},
	"groq":         {"llama-", "mixtral-", "gemma-"},
	"ollama":       {"llama", "mistral", "codellama", "phi", "qwen"},
	"mistral":      {"mistral-"},
	"openrouter":   {"openrouter/"},
	"azure-openai": {"azure/"},
}

func DetermineProvider(modelName string) (string, error) {
	modelLower := strings.ToLower(modelName)

	for provider, patterns := range modelPatterns {
		for _, pattern := range patterns {
			if strings.HasPrefix(modelLower, pattern) {
				return provider, nil
			}
		}
	}

	return "", fmt.Errorf("unable to determine provider for model: %s", modelName)
}
