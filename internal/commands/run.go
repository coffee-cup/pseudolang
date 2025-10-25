package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
	"github.com/username/pseudolang/internal/core"
)

var RunCommand = &cli.Command{
	Name:      "run",
	Usage:     "Run a pseudolang file",
	ArgsUsage: "<file>",
	Action:    runAction,
}

func runAction(ctx context.Context, cmd *cli.Command) error {
	filePath := cmd.Args().First()
	if filePath == "" {
		return fmt.Errorf("file path is required")
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	return core.ExecuteWithLLM(ctx, string(content))
}
