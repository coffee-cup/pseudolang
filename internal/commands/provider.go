package commands

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
	"github.com/username/pseudolang/internal/config"
)

var ProviderCommand = &cli.Command{
	Name:      "provider",
	Usage:     "Set API token for a provider",
	ArgsUsage: "<provider> <token>",
	Action:    providerAction,
}

func providerAction(ctx context.Context, cmd *cli.Command) error {
	if cmd.Args().Len() != 2 {
		return fmt.Errorf("expected exactly 2 arguments: <provider> <token>")
	}

	provider := cmd.Args().Get(0)
	token := cmd.Args().Get(1)

	if !config.IsValidProvider(provider) {
		return fmt.Errorf("invalid provider: %s\nValid providers: openai, anthropic, groq, ollama, mistral, openrouter, azure-openai", provider)
	}

	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	cfg.SetProviderToken(provider, token)

	if err := cfg.Save(); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	fmt.Printf("Successfully saved API token for %s\n", provider)

	return nil
}
