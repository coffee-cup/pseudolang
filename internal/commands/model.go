package commands

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
	"github.com/username/pseudolang/internal/config"
)

var ModelCommand = &cli.Command{
	Name:      "model",
	Usage:     "Switch to a specific model (auto-detects provider)",
	ArgsUsage: "<model>",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "token",
			Usage: "API token for the model's provider",
		},
	},
	Action: modelAction,
}

func modelAction(ctx context.Context, cmd *cli.Command) error {
	if cmd.Args().Len() != 1 {
		return fmt.Errorf("expected exactly 1 argument: <model>")
	}

	model := cmd.Args().Get(0)
	token := cmd.String("token")

	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if token != "" {
		if err := cfg.SetModelWithToken(model, token); err != nil {
			return fmt.Errorf("failed to set model with token: %w", err)
		}
		if err := cfg.Save(); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}
		provider, _ := config.DetermineProvider(model)
		fmt.Printf("Successfully configured %s (provider: %s) with new token\n", model, provider)
		return nil
	}

	if err := cfg.SetActiveModel(model); err != nil {
		return fmt.Errorf("failed to switch to model: %w", err)
	}

	if err := cfg.Save(); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	provider, _ := config.DetermineProvider(model)
	fmt.Printf("Switched to model %s (provider: %s)\n", model, provider)

	return nil
}
