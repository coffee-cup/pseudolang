package core

import (
	"context"
	"fmt"

	"github.com/teilomillet/gollm"
	"github.com/username/pseudolang/internal/config"
)

func ExecuteWithLLM(ctx context.Context, input string) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if cfg.ActiveModel == "" {
		return fmt.Errorf("no active model configured. Use 'ps model <model>' to set one")
	}

	if cfg.ActiveProvider == "" {
		return fmt.Errorf("no active provider configured")
	}

	token, ok := cfg.GetToken(cfg.ActiveProvider)
	if !ok || token == "" {
		return fmt.Errorf("no API token configured for provider: %s", cfg.ActiveProvider)
	}

	llm, err := gollm.NewLLM(
		gollm.SetProvider(cfg.ActiveProvider),
		gollm.SetModel(cfg.ActiveModel),
		gollm.SetAPIKey(token),
		gollm.SetMaxTokens(1000),
	)
	if err != nil {
		return fmt.Errorf("failed to initialize LLM: %w", err)
	}

	systemPrompt := "you are a robot. respond like one"
	promptText := fmt.Sprintf("%s\n\nquery: %s", systemPrompt, input)

	prompt := gollm.NewPrompt(promptText)

	response, err := llm.Generate(ctx, prompt)
	if err != nil {
		return fmt.Errorf("failed to generate response: %w", err)
	}

	fmt.Println(response)

	return nil
}
