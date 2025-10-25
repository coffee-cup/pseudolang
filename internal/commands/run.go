package commands

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
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

	// TODO: Implement file execution logic
	fmt.Printf("Running file: %s\n", filePath)

	return nil
}
