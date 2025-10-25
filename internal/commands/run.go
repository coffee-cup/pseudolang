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
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"v"},
			Usage:   "Print the generated Python code before execution",
		},
	},
	Action: runAction,
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

	verbose := cmd.Bool("verbose")
	return core.ExecuteWithLLM(ctx, string(content), verbose)
}
