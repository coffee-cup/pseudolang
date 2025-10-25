package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/urfave/cli/v3"
	"github.com/username/pseudolang/internal/core"
)

var ExecCommand = &cli.Command{
	Name:      "exec",
	Usage:     "Execute a pseudolang string",
	ArgsUsage: "<string>",
	Action:    execAction,
}

func execAction(ctx context.Context, cmd *cli.Command) error {
	args := cmd.Args().Slice()
	if len(args) == 0 {
		return fmt.Errorf("input string is required")
	}

	userInput := strings.Join(args, " ")

	return core.ExecuteWithLLM(ctx, userInput)
}
