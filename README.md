# Pseudolang

A forgiving programming _language_ that allows you to write code in any
pseudocode style that you want. The only downside is that it costs money to run
and it is super slow. Oh, and the code is not guaranteed to be correct.

<img width="1002" height="690" alt="screenshot-2025-10-25-03 13 57" src="https://github.com/user-attachments/assets/baf5a830-74ff-4e3f-8462-9262db5f20de" />

## Usage

```bash
# Set the active model to Claude Haiku 4.5 and provide your API token.
pseudo model claude-haiku-4-5-20251001 --token sk-ant-...

# Execute some psudocode in any style
pseudo exec "let x = sort([5,2,8,1,9]); filter x where n > 4 then print"
# [5, 8, 9]

# Execute a file
pseudo run tests/quicksort.pseudo
```

## Installation

- Checkout repo
- Run `mise run build` to build the project
- `./out/pseudo --help` to see the available commands

## Configuring models

pseudolang uses
[gollm](https://github.com/teilomillet/gollm?tab=readme-ov-file#supported-providers)
to interface with LLMs. You can use any supported provider and model.

```bash
# Set the model and token in one command
pseudo model claude-haiku-4-5-20251001 --token <token>

# Change the model
pseudo model claude-sonnet-4-5-20250929

# Change the token
pseudo provider openai sk-proj-...
```

## Running Code

```bash
pseudo run <file>     # Run a pseudolang file
pseudo exec <code>    # Execute pseudolang code directly
```

Verbose mode can be enabled with the `--verbose` flag. This will print the
generated Python code before execution.

## Development

- `mise run build`: Build the project (outputs to `out/ps`)
- `mise run test`: Run tests
- `mise run check`: Check code with linters and formatters
- `mise run tidy`: Clean up Go module dependencies

Configuration is stored at `~/.config/pseudolang/config.json`.
