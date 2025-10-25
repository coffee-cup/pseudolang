package commands

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

var ExecCommand = &cli.Command{
	Name:      "exec",
	Usage:     "Execute a pseudolang string",
	ArgsUsage: "<string>",
	Action:    execAction,
}

func execAction(ctx context.Context, cmd *cli.Command) error {
	code := cmd.Args().First()
	if code == "" {
		return fmt.Errorf("code string is required")
	}

	// TODO: Implement string execution logic
	fmt.Printf("Executing: %s\n", code)

	return nil
}
