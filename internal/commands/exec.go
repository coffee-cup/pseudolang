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
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"v"},
			Usage:   "Print the generated Python code before execution",
		},
	},
	Action: execAction,
}

func execAction(ctx context.Context, cmd *cli.Command) error {
	args := cmd.Args().Slice()
	if len(args) == 0 {
		return fmt.Errorf("input string is required")
	}

	userInput := strings.Join(args, " ")

	verbose := cmd.Bool("verbose")
	return core.ExecuteWithLLM(ctx, userInput, verbose)
}
