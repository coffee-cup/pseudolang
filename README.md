# Pseudolang

A command-line interpreter for the pseudolang language with integrated LLM provider management. Pseudolang allows you to configure and switch between multiple LLM providers (OpenAI, Anthropic, Groq, etc.) and models seamlessly.

## Usage

### Installation

Build the project using mise:

```bash
mise run build
```

The binary will be output to `out/ps`. The shorthand command is `ps`.

### Configuring Providers

Before using LLM features, configure API tokens for your providers:

```bash
ps provider anthropic sk-ant-...
ps provider openai sk-proj-...
ps provider groq gsk-...
```

This stores the API tokens without changing your active model.

### Switching Models

Switch between models using the model command. The provider is automatically detected from the model name:

```bash
ps model claude-3-5-sonnet-20240620
ps model gpt-4o
ps model llama-3.1-70b-versatile
```

You can also set a model and its provider token in one command:

```bash
ps model claude-3-5-sonnet --token sk-ant-...
```

### Supported Providers

- **OpenAI**: gpt-4, gpt-4o, gpt-3.5-turbo, o1, o3
- **Anthropic**: claude-3-5-sonnet, claude-3-opus, claude-3-haiku
- **Groq**: llama-3, mixtral, gemma
- **Ollama**: Local models (llama, mistral, codellama, phi, qwen)
- **Mistral**: mistral-large, mistral-medium
- **OpenRouter**: Models via OpenRouter API
- **Azure OpenAI**: Azure-hosted OpenAI models

### Running Code

```bash
ps run <file>     # Run a pseudolang file
ps exec <code>    # Execute pseudolang code directly
```

Note: Interpreter functionality is currently in development.

## Development

### Building

Build the project:

```bash
mise run build
```

The binary will be created at `out/pseudolang`.

### Testing

Run tests:

```bash
mise run test
```

### Code Quality

Check code with linters and formatters:

```bash
mise run check
```

Use this after making significant changes.

### Tidy Dependencies

Clean up Go module dependencies:

```bash
mise run tidy
```

### Project Structure

```
pseudolang/
├── cmd/pseudolang/       # Main entry point
│   └── main.go
├── internal/
│   ├── commands/         # CLI command implementations
│   │   ├── model.go      # Model switching command
│   │   ├── provider.go   # Provider configuration command
│   │   ├── run.go        # Run interpreter command
│   │   └── exec.go       # Execute code command
│   └── config/           # Configuration management
│       ├── config.go     # Config file handling
│       └── models.go     # Model-to-provider mapping
└── mise.toml             # Build configuration
```

### Running During Development

Instead of building, you can run directly:

```bash
go run cmd/pseudolang/main.go <command> [args]
```

### Configuration Location

Configuration is stored at `~/.config/pseudolang/config.json`.
