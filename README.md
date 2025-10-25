# Pseudolang

A forgiving programming _language_ that allows you to write code in any
pseudocode style that you want. The only downside is that it costs money to run
and it is super slow. Oh, and the code is not guaranteed to be correct.

<img width="792" height="608" alt="screenshot-2025-10-25-03 03 01" src="https://github.com/user-attachments/assets/b1063a2a-0096-407a-beb0-5e5cbc61ec55" />


## Usage

```bash
# Set the active model to Claude Haiku 4.5. Provide your API token.
pseudo model claude-haiku-4-5-20251001 --token sk-ant-...

# Execute some psudocode
# Use any pseudocode style you want.
pseudo exec "print('Hello, world!')"
pseudo exec "cout << 'Hello, world!\n'"

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
