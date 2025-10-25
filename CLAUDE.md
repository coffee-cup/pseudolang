# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with
code in this repository.

## Build Commands

Build the project:

```bash
mise run build
```

This outputs the binary to `out/pseudolang`.

Tidy dependencies:

```bash
mise run tidy
```

## Architecture

This is a Go-based CLI interpreter for a pseudolang language using
urfave/cli/v3.

- `cmd/pseudolang/main.go`: Entry point that sets up the CLI with two commands
- `internal/commands/`: Command implementations
  - `run.go`: Executes pseudolang from a file (not yet implemented)
  - `exec.go`: Executes pseudolang from a string argument (not yet implemented)

The interpreter logic itself is not yet implemented (marked with TODOs in both
command files).
