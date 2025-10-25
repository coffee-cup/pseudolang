package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
	"github.com/username/pseudolang/internal/commands"
)

func main() {
	cmd := &cli.Command{
		Name:    "pseudolang",
		Version: "0.1.0",
		Usage:   "A pseudolang interpreter",
		Commands: []*cli.Command{
			commands.RunCommand,
			commands.ExecCommand,
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
