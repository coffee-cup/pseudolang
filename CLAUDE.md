# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with
code in this repository.

## Build Commands

- Build the project:
  - `mise run build`
  - Outputs the binary to `out/pseudolang`
- Tidy dependencies:
  - `mise run tidy`
- Check code:
  - `mise run check`
  - Use this after making big changes to the code.
- Run tests:
  - `mise run test`

## Architecture

This is a Go-based CLI interpreter for a pseudolang language using
urfave/cli/v3 with integrated LLM provider management.

- `cmd/pseudolang/main.go`: Entry point that sets up the CLI
- `internal/commands/`: Command implementations (model, provider, run, exec)
- `internal/config/`: Configuration management and model-to-provider mapping

The interpreter logic itself is not yet implemented. The LLM provider/model
configuration system is fully functional.

The shorthand command name for the interpreter is `ps`.

### Commands

- `ps provider <provider> <token>`: Set API token for a provider
- `ps model <model>`: Switch to a model (auto-detects provider)
- `ps model <model> --token <token>`: Set model and token in one command
- `ps run <file>`: Run a pseudolang file (TODO)
- `ps exec <code>`: Execute pseudolang code (TODO)

## Notes

- Lookup the latest library documentation using context7
- When running the code, prefer to use `go run cmd/pseudolang/main.go` instead of building the binary.
